# KAI Backend: Indonesian Railway Ticket Booking API 🚆⚡

## ✨ Overview

Welcome to **KAI Backend**, a comprehensive REST API designed for managing Indonesian Railway (Kereta Api Indonesia) ticket booking system. Built with Go, Fiber framework, and MongoDB, this backend service provides a fast, scalable, and efficient foundation for a railway ticket booking application. It follows modern API design principles with comprehensive JWT authentication, clean architecture, and robust data validation, making it highly maintainable and performant.

## 🔋 Key Features

- 🔐 **JWT Authentication** — Secure endpoints using JSON Web Tokens (JWT) with Bearer token support, ensuring that only authenticated users can access protected resources.
- 🏗️ **Clean Architecture** — Organized into distinct layers (Handlers, Services, Repositories, Domain) for clear separation of concerns, making the codebase easy to understand, test, and scale.
- 📦 **Full CRUD Operations** — Comprehensive Create, Read, Update, and Delete functionality for all core entities:
  - **Users**: Complete user management with secure bcrypt password hashing and JWT-based authentication.
  - **Stations**: Manage railway stations with detailed information and geographic data.
  - **Trains**: Train fleet management with carriage and seat configuration.
  - **Schedules**: Route scheduling with origin-destination mapping and real-time availability.
  - **Bookings**: Complete booking lifecycle from creation to payment tracking and cancellation.
- 🎫 **Booking System** — Core booking functionality with:
  - Seat availability checking
  - Automated booking code generation
  - Basic booking status management
- 🔍 **Schedule Search** — Search schedules by origin, destination, and departure date.
- 🛡️ **Request Validation** — Built-in validation using Go validator with comprehensive error messages for data integrity.
- 🍃 **MongoDB Integration** — Utilizes MongoDB with official Go driver for flexible document-based data storage and efficient queries.
- 🚀 **High Performance** — Built on Fiber framework for automatic API documentation, middleware support, and blazing-fast HTTP performance.
- 📊 **Interactive Documentation** — Auto-generated Swagger/OpenAPI documentation with "Try it out" functionality.
- ⚙️ **Centralized Configuration** — Manages all environment-specific settings securely through environment variables with godotenv.
- 🔒 **Security Middleware** — CORS protection, request logging, and JWT middleware for comprehensive API security.

## 🧑‍💻 How It Works

1. **User registers** by sending their details to the `/api/auth/register` endpoint with email, password, and personal information.
2. **User authenticates** via `/api/auth/login` to receive JWT tokens for accessing protected endpoints.
3. **User searches schedules** by origin station, destination station, and departure date to find available trains.
4. **User creates booking** by selecting a schedule and seat, the system checks availability and creates the booking.
5. **JWT Middleware** validates tokens for protected endpoints and extracts user information.
6. **The system follows Clean Architecture**: Handler → Service → Repository → Database for clear separation of concerns.
7. **MongoDB stores all data** with proper indexing for efficient queries and data retrieval.
8. **Structured JSON responses** with consistent error handling are returned to the client.

## ⚙️ Tech Stack

- 🐹 **Go 1.21+** (Programming Language)
- ⚡ **Fiber v2** (High-performance HTTP Framework)
- 🍃 **MongoDB** (NoSQL Database)
- 🔗 **MongoDB Go Driver** (Official Database Driver)
- 🔐 **golang-jwt/jwt** (JWT Implementation)
- 🛡️ **bcrypt** (Password Hashing)
- ✅ **go-playground/validator** (Data Validation)
- 📝 **Swagger/OpenAPI** (API Documentation)
- 🔄 **godotenv** (Environment Configuration)
- 🌐 **CORS Middleware** (Cross-Origin Resource Sharing)
- 📊 **Request Logger** (HTTP Request Logging)

## 📚 KAI Backend Resources

- 🌐 **Go Backend**: [View Code](https://github.com/LouisFernando1204/kai-backend)
- 📖 **API Documentation**: `http://localhost:8080/docs` (when running locally)

## 🚀 Getting Started

Follow these steps to get KAI Backend up and running on your local machine.

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.21 or higher)
- [MongoDB](https://www.mongodb.com/try/download/community) (Local installation or MongoDB Atlas)
- A tool to interact with your database (e.g., MongoDB Compass, Studio 3T, or MongoDB Shell)

### Installation & Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/LouisFernando1204/kai-backend.git
   cd kai-backend
   ```

2. **Initialize Go modules:**

   ```bash
   go mod download
   ```

3. **Set up environment variables:**

   - Create a `.env` file in the root directory.
   - Add the following configuration variables:

   ```env
   # Server Configuration
   SERVER_HOST=localhost
   SERVER_PORT=8080

   # MongoDB Atlas Configuration
   MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/?retryWrites=true&w=majority
   MONGO_DATABASE=kai_db

   # JWT Configuration
   JWT_KEY=your_super_secret_jwt_key_here_at_least_32_characters
   JWT_EXP=60
   ```

4. **Generate JWT Secret Key (Optional):**

   ```bash
   go run cmd/jwt_key_generator/main.go
   ```

   Copy the generated key to your `.env` file as `JWT_KEY`.

5. **Set up the database:**

   - Start your MongoDB server (local) or ensure MongoDB Atlas cluster is running.
   - The application will automatically connect to the database specified in `MONGO_URI`.
   - Collections will be created automatically when first accessed.

6. **Build and run the application:**

   ```bash
   # Development mode
   go run main.go

   # Or build and run
   go build -o kai-backend
   ./kai-backend
   ```

   The server should now be running on `http://localhost:8080`.

7. **Access API Documentation:**
   - Swagger UI: `http://localhost:8080/docs`
   - Welcome Message: `http://localhost:8080/`

## 📋 API Endpoints

### Authentication

- `POST /api/auth/register` - Register new user with email, password, and personal details
- `POST /api/auth/login` - User login with email and password

### Station Management

- `GET /api/stations` - Get all railway stations
- `GET /api/stations/{id}` - Get station by ID
- `POST /api/stations` - Create new station (Admin only)
- `PUT /api/stations/{id}` - Update station (Admin only)
- `DELETE /api/stations/{id}` - Delete station (Admin only)

### Train Management

- `GET /api/trains` - Get all trains
- `GET /api/trains/{id}` - Get train by ID
- `POST /api/trains` - Create new train with carriages (Admin only)
- `PUT /api/trains/{id}` - Update train (Admin only)
- `DELETE /api/trains/{id}` - Delete train (Admin only)

### Schedule Management

- `GET /api/schedules/search` - Search schedules by origin, destination, and departure date
- `GET /api/schedules/{id}` - Get schedule by ID
- `GET /api/schedules/{id}/seats` - Get seat layout for schedule
- `GET /api/schedules` - Get all schedules (Admin only)
- `POST /api/schedules` - Create new schedule (Admin only)
- `PUT /api/schedules/{id}` - Update schedule (Admin only)
- `DELETE /api/schedules/{id}` - Delete schedule (Admin only)

### Booking Management (Protected)

- `POST /api/bookings` - Create new booking with seat selection
- `GET /api/bookings` - Get current user's bookings
- `GET /api/bookings/{id}` - Get booking by ID
- `PATCH /api/bookings/{id}/status` - Update booking status
- `POST /api/bookings/{id}/cancel` - Cancel booking

## 🤝 Contributor

- 🧑‍💻 **Louis Fernando** : [@LouisFernando1204](https://github.com/LouisFernando1204)