# Calculator API

A lightweight HTTP API for basic arithmetic operations built with Go.

## Description

This project provides a simple REST API server that performs basic mathematical operations including addition, subtraction, multiplication, and division. The server accepts parameters via query strings and returns formatted results. All requests and responses are logged with timestamps and client information.

## Features

- Four arithmetic endpoints: addition, subtraction, multiplication, and division
- Query parameter validation with error handling
- Request logging with IP address and status tracking
- Division by zero protection
- Plain text response format

## Installation

Ensure you have Go installed on your system (version 1.16 or later recommended).

Clone the repository:

```bash
git clone https://github.com/yourusername/calculator-api.git
cd calculator-api
```

## Usage

Start the server:

```bash
go run .
```

The server will start on `http://localhost:8080`.

## API Endpoints

All endpoints accept two query parameters: `a` and `b`. Both parameters must be valid numeric values.

### Addition

```
GET /add?a={number}&b={number}
```

Returns the sum of a and b.

### Subtraction

```
GET /sub?a={number}&b={number}
```

Returns the difference between a and b.

### Multiplication

```
GET /mul?a={number}&b={number}
```

Returns the product of a and b.

### Division

```
GET /div?a={number}&b={number}
```

Returns the quotient of a divided by b. Returns an error if b is zero.

## Example Requests

```bash
curl "http://localhost:8080/add?a=10&b=3"
# Result : 10.00 + 3.00 = 13.00

curl "http://localhost:8080/sub?a=10&b=3"
# Result : 10.00 - 3.00 = 7.00

curl "http://localhost:8080/mul?a=10&b=3"
# Result : 10.00 * 3.00 = 30.00

curl "http://localhost:8080/div?a=10&b=3"
# Result : 10.00 / 3.00 = 3.33
```

## Error Handling

The API returns appropriate HTTP status codes and error messages for invalid requests:

- `400 Bad Request`: Missing parameters, invalid number format, or division by zero
- `200 OK`: Successful operation

## Project Structure

```
.
├── main.go        # Server initialization and routing
├── handler.go     # HTTP request handlers and validation
└── services.go    # Arithmetic operation functions
```

## Logging

The server logs all requests with the following information:

- Timestamp
- HTTP status code
- Client IP address
- Request path

## License

This project is available under the MIT License.
