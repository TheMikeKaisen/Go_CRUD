# Go CRUD Backend API

A simple and basic CRUD (Create, Read, Update, Delete) backend API built with Go (Golang). This project demonstrates how to use the `net/http` package and the `gorilla/mux` router to handle routes and perform CRUD operations on a list of movies.

## Features

- **GET /movies**: Retrieve a list of all movies in JSON format.
- **GET /movies/{id}**: Retrieve details of a specific movie by its ID.
- **POST /movies**: Create a new movie with a random unique ID.
- **PUT /movies/{id}**: Update an existing movie's details by its ID.
- **DELETE /movies/{id}**: Delete a movie from the list by its ID.

## Key Highlights

- Uses the `gorilla/mux` package for routing.
- Demonstrates working with JSON encoding/decoding for API requests and responses.
- Includes predefined sample movie data for testing.
- Lightweight and beginner-friendly backend implementation.

## How to Run

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_name>
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. Access the API at `http://localhost:8000` with tools like Postman, curl, or your browser.

## Example Endpoints

- **Retrieve all movies**:  
  `GET http://localhost:8000/movies`

- **Retrieve a movie by ID**:  
  `GET http://localhost:8000/movies/{id}`

- **Create a new movie**:  
  `POST http://localhost:8000/movies`  
  Body:  
  ```json
  {
    "isbn": "123456789",
    "title": "New Movie Title",
    "director": {
      "firstname": "First",
      "lastname": "Last"
    }
  }
  ```

- **Update a movie**:  
  `PUT http://localhost:8000/movies/{id}`  
  Body:  
  ```json
  {
    "isbn": "987654321",
    "title": "Updated Movie Title",
    "director": {
      "firstname": "Updated First",
      "lastname": "Updated Last"
    }
  }
  ```

- **Delete a movie**:  
  `DELETE http://localhost:8000/movies/{id}`

## Tech Stack

- **Language**: Go (Golang)
- **Router**: `gorilla/mux`

---
