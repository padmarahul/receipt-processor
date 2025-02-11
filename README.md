# receipt-processor

The application is built using **GoLang** programming language and **redis** has been used as in-memory database and application
has been running successfully on local machine as well as on **Docker**.

- The repository has project structure where **main.go** is the entry point of the code. 
- The utils folder contains the common reusable code which will be used in other files.
- The automated tests suite  are written in tests folder.
- The storage folder contains the code related to Redis database.
- Tge services folder contains the core business logic.
- The models folder contains the different JSON payload definitions and its datatypes.
- The handlers folder contains the code reladifferent API endpoints.
- The examples folder contains the payload requests for POST API request.
- The config folder contains the environment variables and other main configurations.
- The application can be run locally using go commands. 
- `go mod tidy` is the command to reinstall and build all dependencies.
- `go run main.go` is the command to run the application server.
- `go test ./tests -v` is the command to run the automated test suites.
- `Dockerfile` is created to build docker image for go application.
- `docker-compose.yml` is created to run all containers on same network.
- `docker compose up -d` is the command to build and run containers.
- Tested the APIs using postman. Sharing the API testing below.
- The keys are stored in redis database. The ids are stored in redis database. The ids would be
  coming in the POST API and the ids are retrieved from redis database to fetch the points.

1. `processReceipt API`- POST request to receipt processor.
- Endpoint: http://localhost:8080/receipts/process 
## Request Body

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },
    {
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },
    {
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },
    {
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },
    {
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}

```
## Response Body
```json
{
    "id": "57a0dcb2-122a-467c-97ac-480eb8c0f485"
}
```

2. `getPoints API`- GET request to receipt processor.
- Endpoint: http://localhost:8080/receipts/57a0dcb2-122a-467c-97ac-480eb8c0f485/points
## Response Body
```json
{
    "points": 109
}
```