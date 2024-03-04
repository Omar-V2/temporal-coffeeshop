package integration_test

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"tmprldemo/internal/customer/api"
	customerdata "tmprldemo/internal/customer/data/customer"
	migration "tmprldemo/internal/customer/migrations"
	"tmprldemo/internal/customer/workflows/verifyphone"
	customerpb "tmprldemo/internal/pb/customer/v1"
	"tmprldemo/pkg/database"
	"tmprldemo/pkg/testutils"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type dockerCompose interface {
	Up(ctx context.Context, opts ...tc.StackUpOption) (err error)
	Down(ctx context.Context, opts ...tc.StackDownOption) (err error)
}

type CustomerIntegrationTestSuite struct {
	suite.Suite
	customerServiceAddress string
	customerServiceClient  customerpb.CustomerServiceClient
	dockerCompose          dockerCompose
	temporalClient         client.Client
	temporalWorker         worker.Worker
}

func TestCustomerIntegrationTestSuite(t *testing.T) {
	if os.Getenv("INTEGRATION") == "" {
		t.Skip("skipping customer integration test, to un-skip set the INTEGRATION environment variable")
	}

	suite.Run(t, new(CustomerIntegrationTestSuite))
}

func (s *CustomerIntegrationTestSuite) SetupSuite() {
	testCtx := context.Background()

	customerServiceAddress, err := testutils.GetFreeAddress("localhost")
	s.Require().NoError(err)
	s.customerServiceAddress = customerServiceAddress

	dockerCompose, err := tc.NewDockerCompose("docker-compose.yml")
	s.Require().NoError(err)
	s.dockerCompose = dockerCompose

	err = dockerCompose.
		WaitForService("temporal",
			wait.ForLog("temporal-sys-tq-scanner-workflow workflow successfully started").
				WithStartupTimeout(time.Second*10),
		).
		Up(testCtx, tc.Wait(true))
	s.Require().NoError(err)

	temporalContainer, err := dockerCompose.ServiceContainer(testCtx, "temporal")
	s.Require().NoError(err)

	temporalAddress, err := temporalContainer.Endpoint(testCtx, "")
	s.Require().NoError(err)

	postgresContainer, err := dockerCompose.ServiceContainer(testCtx, "postgres")
	s.Require().NoError(err)

	postgresAddress, err := postgresContainer.Endpoint(testCtx, "")
	s.Require().NoError(err)

	s.temporalClient, err = client.NewLazyClient(client.Options{
		HostPort:  temporalAddress,
		Namespace: "default",
	})
	s.Require().NoError(err)

	s.runCustomerServer(postgresAddress)
	s.runCustomerWorker(postgresAddress)

	conn, err := grpc.Dial(s.customerServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)

	s.customerServiceClient = customerpb.NewCustomerServiceClient(conn)
}

func (s *CustomerIntegrationTestSuite) TearDownSuite() {
	s.temporalWorker.Stop()

	err := s.dockerCompose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal)
	s.Require().NoError(err)
}

func (s *CustomerIntegrationTestSuite) runCustomerServer(postgresAddress string) {
	connectionString := fmt.Sprintf(
		"postgres://postgres:password@%s/customer?sslmode=disable", postgresAddress,
	)
	db, err := sql.Open("pgx", connectionString)
	s.Require().NoError(err)

	migrator, err := database.NewPostgresMigrator(migration.Customer, db)
	s.Require().NoError(err)

	err = migrator.Up()
	s.Require().NoError(err)

	customerDBCreator := customerdata.NewCustomerDBCreator(db)
	customerDBGetter := customerdata.NewCustomerDBGetter(db)
	customerServiceServer := api.NewCustomerServiceGRPCServer(
		customerDBCreator, customerDBGetter, s.temporalClient,
	)

	server := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(server, customerServiceServer)

	listener, err := net.Listen("tcp", s.customerServiceAddress)
	s.Require().NoError(err)

	go func() {
		err := server.Serve(listener)
		s.Require().NoError(err)
	}()
}

func (s *CustomerIntegrationTestSuite) runCustomerWorker(postgresAddress string) {
	connectionString := fmt.Sprintf(
		"postgres://postgres:password@%s/customer?sslmode=disable", postgresAddress,
	)
	db, err := sql.Open("pgx", connectionString)
	s.Require().NoError(err)

	s.temporalWorker = worker.New(s.temporalClient, "TEMPORAL_COFFEE_SHOP_TASK_QUEUE", worker.Options{})

	customerVerifier := customerdata.NewCustomerDBVerifier(db)
	codeGenerator := verifyphone.StaticCodeGenerator{}
	mockSMSSender := &verifyphone.MockSMSSender{}
	activities := verifyphone.NewActivities(mockSMSSender, customerVerifier, codeGenerator)

	s.temporalWorker.RegisterActivity(activities)
	s.temporalWorker.RegisterWorkflow(verifyphone.NewVerificationWorkflow)

	err = s.temporalWorker.Start()
	s.Require().NoError(err)
}
