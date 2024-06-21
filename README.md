# IP Information Service

This is a simple service that provides detailed information about an IP address, including geographical location, Internet Service Provider (ISP), and other details. The service utilizes the [ipify](https://www.ipify.org/) and [ip-api](https://ip-api.com/) APIs to retrieve this information.

## Project Structure

The project is structured as follows:

```
.
├── main.go
├── ip.go
├── ip_test.go
└── handlers.go
```

- **main.go**: Entry point of the application. Configures and starts the HTTP server.
- **ip.go**: Contains the `IPResponse` structure and functions related to fetching and querying IP information.
- **handlers.go**: Includes the `handleIPRequest` function which handles HTTP requests to fetch IP details.
- **ip_test.go**: Contains tests for the `getClientIP` and `handleIPRequest` functions.

## How to Run

1. **Prerequisites**:
   - Go (Golang) installed on your machine.

2. **Installation**:
   - Clone the repository:
     ```sh
     git clone https://github.com/adormundo/ipexplorer.git
     cd ipexplorer
     ```

3. **Running the Server**:
   - Execute the command:
     ```sh
     go run main.go ip.go handlers.go
     ```
   - The server will be available at `http://<yourlocalhost>:8080`.

## Endpoints

- **GET /**: Returns detailed information about the client's IP.

## Testing

To run tests, use the command:

```sh
go test
```

## Contributions

Contributions are welcome! Please open an issue or submit a pull request to discuss what you would like to change.

## License

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/).