# SkyServe-Assignment 
Overview
This project is an assignment for SkyServe. It implements an authentication system using Go (Golang) for backend development.

Features
Signup: Allows users to register by providing a username and password.
Login: Allows registered users to log in using their username and password.
Logout: Allows logged-in users to log out of their account.
Middleware: Implements authentication middleware to protect certain routes from unauthorized access.
Prerequisites
Before running this project, ensure you have the following installed:

Go (Golang)
PostgreSQL database
Required Go packages (can be installed using go get)
Installation
Clone the repository:
bash
Copy code
git clone https://github.com/your-username/SkyServe-Assignment.git
Navigate to the project directory:
bash
Copy code
cd SkyServe-Assignment
Install dependencies:
bash
Copy code
go mod tidy
Set up the PostgreSQL database and configure the connection string in main.go.
Run the application:
bash
Copy code
go run cmd/main.go
Usage
Use HTTP requests to interact with the authentication endpoints (signup, login, logout).
Example usage:

bash
Copy code
curl -X POST -d '{"username":"example", "password":"password123"}' http://localhost:8080/signup
Contributors
Your Name
License
This project is licensed under the MIT License.
