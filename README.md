# Bookstore Application

The Bookstore application is a Go-based service designed to manage books and their data. It allows users to do a simple CRUD, all while storing the information in a MySQL database.

## Technologies Used

- **Go** - The programming language used to develop the application.
- **MySQL** - Relational database used to store the books.
- **Docker** - A containerization platform used to package and deploy the application in lightweight, isolated environments, ensuring consistency across different environments. muda isto para estar de acordo com o bookStore

## How to Run the Project Locally

### Installation Steps

1. **Clone the repository:**

   ```bash
   git clone https://github.com/LuisSilva7/go-bookstore.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd go-bookstore
   ```

3. **Run MySQL container:**

   ```bash
   docker compose up -d
   ```

4. **Navigate to the main directory:**

   ```bash
   cd cmd/main
   ```
   
5. **Compile the project:**

   ```bash
   go build
   ```
   
6. **Run the project:**

   ```bash
   go run main.go
   ```
