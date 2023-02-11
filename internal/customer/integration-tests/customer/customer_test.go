package customer_integration_test

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	migration "tmprldemo/internal/customer/migrations"
// 	customerpb "tmprldemo/internal/pb/customer/v1"
// 	"tmprldemo/pkg/testutils"

// 	"github.com/google/uuid"
// 	"github.com/orlangure/gnomock"
// 	"github.com/stretchr/testify/suite"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// 	"google.golang.org/protobuf/proto"
// )

// type CustomerServiceIntegrationTestSuite struct {
// 	suite.Suite
// 	customerClient  customerpb.CustomerServiceClient
// 	postgres        *gnomock.Container
// 	temporalServer  *gnomock.Container
// 	temporalWorker  *gnomock.Container
// 	customerService *gnomock.Container
// }

// func (s *CustomerServiceIntegrationTestSuite) SetupSuite() {
// 	var err error
// 	s.postgres, _ = testutils.MustNewPostgresInstance(
// 		"customer",
// 		migration.Customer,
// 	)

// 	s.temporalServer, err = gnomock.StartCustom(
// 		"ianthpun/temporalite",
// 		gnomock.DefaultTCP(7233),
// 		gnomock.WithUseLocalImagesFirst(),
// 		gnomock.WithDebugMode(),
// 	)
// 	s.Require().NoError(err)

// 	s.temporalWorker, err = gnomock.StartCustom(
// 		"customer-worker",
// 		gnomock.DefaultTCP(8082),
// 		gnomock.WithUseLocalImagesFirst(),
// 		gnomock.WithDebugMode(),
// 		gnomock.WithEnv(fmt.Sprintf("TEMPORAL_ADDRESS=%s", s.temporalServer.DefaultAddress())),
// 	)
// 	s.Require().NoError(err)

// 	s.customerService, err = gnomock.StartCustom(
// 		"customer-service",
// 		gnomock.DefaultTCP(8080),
// 		gnomock.WithUseLocalImagesFirst(),
// 		gnomock.WithEnv("POSTGRES_DB=customer"),
// 		gnomock.WithEnv(fmt.Sprintf("POSTGRES_HOST=%s", s.postgres.Host)),
// 		gnomock.WithEnv(fmt.Sprintf("POSTGRES_PORT=%d", s.postgres.DefaultPort())),
// 		gnomock.WithEnv(fmt.Sprintf("TEMPORAL_ADDRESS=%s", s.temporalServer.DefaultAddress())),
// 	)
// 	s.Require().NoError(err)

// 	conn, err := grpc.Dial(
// 		s.customerService.DefaultAddress(),
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	)
// 	s.Require().NoError(err)
// 	s.customerClient = customerpb.NewCustomerServiceClient(conn)
// }

// func (s *CustomerServiceIntegrationTestSuite) TearDownSuite() {
// 	err := gnomock.Stop(
// 		s.customerService,
// 		s.postgres,
// 		s.temporalServer,
// 		s.temporalWorker,
// 	)
// 	s.Require().NoError(err)
// }

// func (s *CustomerServiceIntegrationTestSuite) TestCreateCustomer() {
// 	testCtx := context.Background()
// 	customerID := uuid.New()
// 	customerToCreate := &customerpb.Customer{
// 		Id:          customerID.String(),
// 		FirstName:   "Naruto",
// 		LastName:    "Uzumaki",
// 		Email:       "naruto@hiddenleaf.com",
// 		PhoneNumber: "112738914",
// 	}

// 	requestID := uuid.New()
// 	req := &customerpb.CreateCustomerRequest{
// 		RequestId: requestID.String(),
// 		Customer:  customerToCreate,
// 	}

// 	createResponse, err := s.customerClient.CreateCustomer(testCtx, req)
// 	s.Require().NoError(err)
// 	s.True(proto.Equal(customerToCreate, createResponse.Customer))

// 	getResponse, err := s.customerClient.GetCustomer(context.Background(), &customerpb.GetCustomerRequest{
// 		CustomerId: customerID.String(),
// 	})
// 	s.Require().NoError(err)
// 	s.True(proto.Equal(customerToCreate, getResponse.Customer))
// }

// func TestCustomerServiceIntegrationTestSuite(t *testing.T) {
// 	suite.Run(t, new(CustomerServiceIntegrationTestSuite))
// }
