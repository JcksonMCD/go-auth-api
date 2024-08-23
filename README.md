# Go User Management API

## Overview

This API provides endpoints for managing user data, including retrieving user information based on user IDs. It is built using Go and MongoDB.

## Technologies Used

- **Go**: Programming language used for implementing the API.
- **Gin**: Web framework for Go, used for routing and handling HTTP requests.
- **MongoDB**: NoSQL database used for storing user data.
- **Go Modules**: Dependency management for Go projects.

## Endpoints

### 1. **Get User by ID**

Retrieve a user's details by their unique user ID.

- **URL:** `/user/{user_id}`
- **Method:** `GET`
- **URL Params:**
  - `user_id` (required): The unique ID of the user you want to retrieve.
- **Success Response:**
  - **Code:** 200 OK
  - **Content:**
    ```json
    {
      "id": "66c8873ff32116292fe6ed03",
      "firstName": "John",
      "lastName": "Doe",
      "password": "$2a$14$ck73JJdzsOmQp2NCmaar1uQdn.FzpK9Z6qNYrRnNycIPv0Kmo7oha",
      "email": "john.doe@example.com",
      "phone": "1234567890",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "userType": "ADMIN",
      "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "createdAt": "2024-08-23T12:55:30Z",
      "updatedAt": "2024-08-23T13:01:19.537Z",
      "userId": "66c886c22f822334952ddbc9"
    }

### 2. **Get All Users**

Retrieve a list of all users.

- **URL:** `/users`
- **Method:** `GET`
- **Success Response:**
  - **Code:** 200 OK
  - **Content:**
    ```json
    {
      "total_count": 3,
      "user_items": [
        {
          "id": "66c886c22f822334952ddbc9",
          "createdAt": "2024-08-23T12:55:30Z",
          "email": "john.doe@example.com",
          "firstName": "John",
          "lastName": "Doe",
          "password": "$2a$14$ck73JJdzsOmQp2NCmaar1uQdn.FzpK9Z6qNYrRnNycIPv0Kmo7oha",
          "phone": "1234567890",
          "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
          "updatedAt": "2024-08-23T12:55:30Z",
          "userId": "66c886c22f822334952ddbc9",
          "userType": "ADMIN"
        },
        {
          "id": "66c886c22f822334952ddbc9",
          "createdAt": "2024-08-23T12:55:30Z",
          "email": "john.doe@example.com",
          "firstName": "John",
          "lastName": "Doe",
          "password": "$2a$14$ck73JJdzsOmQp2NCmaar1uQdn.FzpK9Z6qNYrRnNycIPv0Kmo7oha",
          "phone": "1234567890",
          "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
          "updatedAt": "2024-08-23T12:55:30Z",
          "userId": "66c886c22f822334952ddbc9",
          "userType": "ADMIN"
        },
        {
          "id": "66c886c22f822334952ddbc9",
          "createdAt": "2024-08-23T12:55:30Z",
          "email": "john.doe@example.com",
          "firstName": "John",
          "lastName": "Doe",
          "password": "$2a$14$ck73JJdzsOmQp2NCmaar1uQdn.FzpK9Z6qNYrRnNycIPv0Kmo7oha",
          "phone": "1234567890",
          "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
          "updatedAt": "2024-08-23T12:55:30Z",
          "userId": "66c886c22f822334952ddbc9",
          "userType": "ADMIN"
        }
      ]
    }

## Setup

### Prerequisites

- Go (1.16 or higher)
- MongoDB

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/your-repo/go-user-api.git
    ```
2. Navigate to the project directory:
    ```bash
    cd go-user-api
    ```
3. Install dependencies:
    ```bash
    go mod tidy
    ```

### Configuration

1. Create a `.env` file in the root directory of the project with the following content:
    ```
    MONGO_URI=mongodb://localhost:27017
    ```

2. Ensure MongoDB is running and accessible at the specified URI.

### Running the Server

Start the server with:
```bash
go run main.go
