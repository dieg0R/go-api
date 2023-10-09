# go-api


## Folder structure

```
go-api/
|-- bin/
|-- cmd/
|   |-- server/
|       |-- main.go
|-- pkg/
|   |-- api/
|       |-- router.go
|       |-- bookings/
|       	|-- handler.go
|       	|-- handler_test.go
|       |-- classes/
|       	|-- handler.go
|       	|-- handler_test.go
|   |-- models/
|       |-- booking.go
|       |-- class.go
|   |-- mockDatabase/
|       |-- db.go
|-- go.mod
|-- go.sum
|-- README.md
```

### Explanation of Directories and Files:

1. **`bin/`**: Contains the compiled binaries.

2. **`cmd/`**: Main applications for this project. The directory name for each application should match the name of the executable.

    - **`main.go`**: The entry point.

3. **`pkg/`**: Libraries and packages that are okay to be used by applications from other projects. 

    - **`api/`**: API logic.
        - **`handler.go`**: HTTP handlers.
        - **`router.go`**: Routes.
    - **`models/`**: Data models.
    - **`mockDatabase/`**: fake Database.


## Getting Started

### Prerequisites

- Go 1.21.2+

### Installation

1. Clone the repository

```bash
git clone https://github.com/dieg0R/go-api.git
```

2. Navigate to the directory

```bash
cd go-api
```

### API Documentation

## Usage

### Endpoints

- `GET /api/classes`: Get all classes.
- `GET /api/classes/:id`: Get a class by ID.
- `POST /api/classes`: Create a new class.
- `PUT /api/classes/:id`: Update a class by ID.
- `DELETE /api/classes/:id`: Delete a class by ID.
- `GET /api/bookings`: Get all bookings.
- `GET /api/bookings/:id`: Get a booking by ID.
- `POST /api/bookings`: Create a new booking.
- `PUT /api/bookings/:id`: Update a booking by ID.
- `DELETE /api/bookings/:id`: Delete a booking by ID.
## Testing
To run the tests for this project, you can use the following command:
go test ./...


