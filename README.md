# Gowash

A simple Go-based web application for booking time slots, designed for car wash services or similar appointment-based businesses.

## Features

- User registration and management
- Time slot booking system
- RESTful API endpoints
- In-memory storage (suitable for development/demo)
- Clean architecture with separated handlers, services, and storage layers

## Installation

1. Ensure you have Go 1.23.3 or later installed.
2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gowash.git
   cd gowash
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```

## Usage

Run the application:

```bash
go run cmd/app/main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

### User Registration
- **POST** `/user`
  - Registers a new user
  - Body: JSON with user details (email, password, username, address)

### Booking
- **POST** `/booking`
  - Books a time slot for a user
  - Body: JSON with user_id and time_slot_id

## Project Structure

```
gowash/
├── cmd/app/           # Application entry point
├── internal/
│   ├── handlers/      # HTTP request handlers and routing
│   ├── models/        # Data models (User, Booking, TimeSlot)
│   ├── services/      # Business logic layer
│   └── storage/       # Data storage layer (in-memory implementation)
├── go.mod             # Go module file
└── README.md          # This file
```

## Dependencies

- [github.com/google/uuid](https://github.com/google/uuid) - For generating unique IDs

## Development

This project uses a clean architecture approach:
- **Handlers**: Handle HTTP requests and responses
- **Services**: Contain business logic
- **Storage**: Abstract data persistence (currently in-memory)
- **Models**: Define data structures

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source. Please check the license file for details.