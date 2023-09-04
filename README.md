# Summary

This repository contains a basic API built using the Go programming language and the Gin web framework. The API is designed to manage student ("Aluno") records, providing CRUD operations.

## Features:

1. **Controllers** (`controllers/controller.go`):
   - Display all students.
   - Greet by name.
   - Create a new student.
   - Fetch a student by ID.
   - Delete a student by ID.
   - Edit student details.
   - Fetch a student by CPF (a Brazilian tax registration number).

2. **Database** (`database/db.go`):
   - Connection setup to a PostgreSQL database.
   - Automatic migration for the student model.

3. **Models** (`models/aluno.go`):
   - Defines the student model with fields: Name, CPF, and RG.

4. **Routes** (`routes/routes.go`):
   - Defines the API endpoints and their corresponding handlers.

5. **Main** (`main.go`):
   - Initializes the database connection and starts handling requests.

## Usage:

To run the API, ensure you have Go installed and the necessary dependencies. Then, simply execute the `main.go` file.

## License:

This project is under the [MIT](LICENSE) license.
