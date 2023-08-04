# # Project1

Project1 is a simple web application written in Go. It uses a PostgreSQL database for user management, and supports two core features - user registration and login.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine.

### Prerequisites

You'll need Go installed on your machine. Additionally, you need PostgreSQL as the database. The project also uses several Go packages:

- `github.com/gorilla/mux`
- `github.com/lib/pq`
- `golang.org/x/crypto/bcrypt`
- `github.com/joho/godotenv`

### Running Project1

1. Clone this repository to your machine.
2. Go to the project directory.
3. Run `go run main.go`.

## Key Files

- `main.go` : The main file of the application. It initializes the database connection, sets up the HTTP server and routes, and includes the register and login handlers.

## API Endpoints

- `/register` : A POST endpoint to register a new user. Expects a JSON body with `username` and `password`.
- `/login` : A POST endpoint for user login. Expects a JSON body with `username` and `password`.

## Authors

Liban

## Acknowledgments

- The Go community

