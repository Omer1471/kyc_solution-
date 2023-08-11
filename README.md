# # Project1

Project1 is a simple web application written in Go that provides user management functionalities including user registration, login, and KYC (Know Your Customer) document submission. It utilizes a PostgreSQL database to store user data and KYC documents.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine.

### Prerequisites

You'll need Go installed on your machine. Additionally, you need PostgreSQL as the database. The project also uses several Go packages:

- `github.com/gorilla/mux`
- `github.com/lib/pq`
- `golang.org/x/crypto/bcrypt`
- `github.com/joho/godotenv`
- 'Go programming language'
- 'PostgreSQL database'

### Running Project1

1. Clone this repository to your machine.
2. Go to the project directory.
3. Run `go run main.go`.

## Key Files

- `main.go` : The main file of the application. It initializes the database connection, sets up the HTTP server and routes, and includes the register and login handlers.
 - `main.go` : The main file of the application. It initializes the database connection, sets up the HTTP server and routes, and includes the register and login handlers.

-`kyc/handlers.go` : Contains the KYC-related route handlers for uploading documents and checking status.
-`kyc/kyc.go` : Defines the KYC package, including the KYCHandler struct and its methods for handling KYC-related requests.

 `kyc/model.go` : Defines the data model for KYC documents.


## API Endpoints

- `/register` : A POST endpoint to register a new user. Expects a JSON body with `username` and `password`.
- `/login` : A POST endpoint for user login. Expects a JSON body with `username` and `password`.
- `/kyc/upload: A POST endpoint for uploading KYC documents.
- `/kyc/status/{document_id}: A GET endpoint for checking the status of a KYC document by its ID.

## Authors

Liban

## Acknowledgments

- Go community

