# Go API Project

This project is a RESTful API built with Go, Gin, and GORM. It provides endpoints for user authentication and book management.

## Prerequisites

- Go 1.23 or later
- Docker
- Docker Compose

## Getting Started

### Clone the repository

```sh
git clone https://github.com/yourusername/go-api.git
cd go-api
```

### Environment Variables

Create a `.env` file in the root directory and add the following environment variables:

```env
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
JWT_SECRET=
```

### Build and Run with Docker Compose

```sh
docker-compose up --build
```

This will start the API service and a PostgreSQL database.

### API Endpoints

#### Authentication

- **Register**: `POST /register`
    - Request Body: `{ "username": "your_username", "password": "your_password" }`
    - Response: `{ "message": "User created successfully" }`

- **Login**: `POST /login`
    - Request Body: `{ "username": "your_username", "password": "your_password" }`
    - Response: `{ "token": "your_jwt_token" }`

#### Books

- **Get All Books**: `GET /books`
    - Response: List of books

- **Get Book by ID**: `GET /books/:id`
    - Response: Book details

- **Create Book**: `POST /books`
    - Request Body: `{ "title": "Book Title", "author": "Author Name", "published_date": "YYYY-MM-DD", "isbn": "ISBN Number" }`
    - Response: Created book

- **Update Book**: `PUT /books/:id`
    - Request Body: `{ "title": "Updated Title", "author": "Updated Author", "published_date": "YYYY-MM-DD", "isbn": "Updated ISBN" }`
    - Response: Updated book

- **Delete Book**: `DELETE /books/:id`
    - Response: `{ "message": "Book deleted successfully" }`

## Project Structure

- `main.go`: Entry point of the application
- `database/`: Database connection and initialization
- `controllers/`: Handlers for API endpoints
- `models/`: Database models
- `routes/`: API route definitions
- `middleware/`: Middleware functions
- `Dockerfile`: Docker configuration for building the application
- `docker-compose.yml`: Docker Compose configuration
.