# Temporal Coffeeshop

This repository aims to demonstrate some of the features of the [Temporal](https://temporal.io) workflow engine and how it can alleviate some difficulties when developing distributed systems.

## High Level Architecture

![architecture](images/architecture.png)

## How to Run

In the root of this repository run `docker-compose up`\
Navigate to localhost:8233 to view the Temporal UI\

### Create a new customer

This will insert a new record into the DB and execute the VerifyPhoneWorkflow

```
curl --location 'http://localhost:8081/v1/customer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "customer": {
        "first_name": "Test",
        "last_name": "Customer",
        "email": "test@gmail.com",
        "phone_number": "07500514"
    },
    "request_id": "34dd76a6-230a-41a2-b603-7b99aad15de8"
}'
```

### Verifying a Customer's Phone Number

This sends a signals to the workflow containing the code you submitted in the request which will then perform the comparison logic. The handler then queries the state of the workflow to retrieve the verification result, which is then sent in the response back to the caller.

```
curl --location 'http://localhost:8081/v1/customer/<your-customer-id>/verify' \
--header 'Content-Type: application/json' \
--data '{
    "verification_code": "1233"
}'
```

The diagram below visualises the end to flow.

![architecture](images/workflow.png)
