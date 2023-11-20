# gamedb 

A simple full-stack Go web application that displays a list of video games from a MySQL database. This app demonstrates basic full-stack development capabilities with Go, including server-side rendering and database interactions. It's designed to be adaptable for different database schemas and can serve as a foundation for more complex full-stack Go applications.

## Features

- Displays video games from a MySQL database.
- Converts numeric ratings to star symbols.
- Logs the IP address of each request.
- Serves as an example of a basic full-stack application using Go.

## Environment

- **Go Version**: 1.19
- **Database**: MariaDB 10.11
- **Operating System**: Debian 12

## Prerequisites

- Go (version 1.19 or later)
- MariaDB or MySQL
- A database named `gamedb` with a table `video_games` (refer to the schema provided in the repository)

## Flexibility and Adaptation

This application can be easily repurposed to work with different database schemas and structures. By modifying the database connection settings and SQL queries, it can serve various types of data, making it a versatile starting point for developing other full-stack applications with Go.

## Setup and Installation

1. **Clone the Repository**:
   ```sh
   git clone [repository-url]
   cd [repository-name]
   ```

2. **Configure Database Connection**: 
   - Open the `main.go` file.
   - Update the database connection string in the `initDB` function with your database credentials.

3. **Create the Database Schema**:
   - Refer to the SQL schema provided in the repository.
   - Create the `video_games` table in your `gamedb` database.

4. **Run the Application**:
   ```sh
   go run main.go
   ```
   - The application will start and listen on port `8085`.

## Accessing the Application

- Open a web browser and navigate to `http://localhost:8085`.
- You should see the list of video games retrieved from the database.

## Contributing

Feel free to fork the repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## Author

- [markmental](https://github.com/markmental)

---
