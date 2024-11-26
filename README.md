### Tech Stack
- Go
- Docker
- Postgres running on Docker
- Swagger for docs
- Golang migrate for migrations

### Folder Structure
- `bin`: Contains the binary file of the application
- `cmd`: Contains the entry point of the application
  - `api`: Contains anything related to the server
  - `migrate`: Contains the migration tool
    - `migrations`: Contains the migration files
- `internal`: Holds all the internal packages. Not to be exported for our api server. This package doesn't know about the outside world.
- `docs`: Contains the swagger docs
- `scripts`: Contains the scripts to run the application

### Commands to keep note of
- `go get -u github.com/go-chi/chi/v5` - To get the chi library for this project
- `go install xyz` - To install the package in the bin folder
- `migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_posts` - To create a new migration file
- `migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" up` - To run the migrations

### Steps to run the application
- `docker-compose up` - To start the postgres database
- `./bin/air` - To start the server
- `make migrate-up` - To run the migrations
- `make migrate-down <int?>` - To rollback the migrations
- `make migrate-create <string>` - To create a new migration file

### 3 layers
- Transport layer: This is the layer that handles the incoming requests and outgoing responses. It's responsible for parsing the incoming requests, validating them, and serializing the outgoing responses.
- Service layer: This is the layer that contains the business logic of your application. It's responsible for processing the incoming requests, interacting with the storage layer, and returning the results.
- Storage layer: This is the layer that interacts with the database. It's responsible for executing the queries, reading and writing the data, and returning the results.

### Thought process
- Start by thinking about the request lifetime
  - User makes a request from the phone/web
  - The request ```GET /v1/users/feed``` is sent to the server
  - Request reaches our mux and then goes to the HTTP handler
  - Handler calls the service layer
  - Service layer interacts with the storage layer
  - Storage layer interacts with the database
  - Database returns the data
  - Storage layer returns the data to the service layer
  - Service layer processes the data
  - Service layer returns the data to the handler
  - Handler serializes the data and sends it back to the user
  - User receives the data

### 3rd party libraries
- golang chi
- Hot reload: air
- golang migrate: To handle the migrations
- go-playground/validator: To validate the incoming requests

### Not included 3rd party libraries but worth mentioning
- GORM - Simplify the database interactions
- sqlx - Simplify the database interactions
- sqlboiler - Can be used to generate the models from the database
- goose - Can be used to handle the migrations

### Further reading
- Clean Architecture
- Twelve factor app - [Link](https://12factor.net)
- Repository pattern [Link](https://www.toptal.com/go/go-repository-tutorial)
- Software development principles (DRY, YAGNI, KISS, etc.)

### Uncertainties
- SAGA pattern
- What is middlewares used for?
- Repository Pattern
  - The repository pattern is a design pattern that abstracts the data access logic from the rest of the application. It provides a way to access the data without exposing the underlying database implementation.
  - Soft Delete vs Hard Delete?

### Principles to understand
- Separation of concerns:
  - Each level in your program should be separate by a clear barrier, the transport layer, the service layer, the storage layer...
- Dependency inversion principle:
  - You're injecting the dependencies in your layers. You don't directly call them.
  - Why? It promotes loose coupling and makes it easier to test your programs.
- Adaptability to change
  - By organizing your code in a modular and flexible way, you can more easily introduce new features, refactor existing code, and respond to evolving business requirements.
  - Your system should be easy to change, if you have to change a lot of existing code to add a new feature you're doing it wrong.
- Focus on business value
  - And finally, focus on delivering value to your users.