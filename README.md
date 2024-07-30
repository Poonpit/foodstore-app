# Food Store Calculator

This project is a food store calculator that calculates the total price for customer orders, considering member card discounts and bundle discounts.

## Setup

1. Clone the Repository **

   ```go
   git clone https://github.com/Poonpit/foodstore-app.git
   cd foodstore-app-main
   ```
   ** Or download ZIP

## 1) Server
## Project Structure <server>

- **`main.go`**: Contains the Fiber application setup and endpoint for calculating totals.
- **`/repositories`**: Contains repository definitions and mock implementations.
  - **`menu.go`**: Defines the `Item` struct and `MenuRepository` interface.
  - **`menu_mock.go`**: Provides a mock implementation of the `MenuRepository`.
- **`/services`**: Contains the service logic for calculating totals and error definitions.
  - **`calculator.go`**: Implements the `CalculatorService` interface.
  - **`errors.go`**: Defines custom errors for the calculator service.
  - **`calculator_test.go`**: Contains unit tests for the `CalculatorService`.

## Setup

1. change directory

   ```
   cd server
   ```

2. Install Dependencies

   ```go
   go mod tidy
   ```

## Testing

  To run the unit tests:
  ```go
   go test server/services -v
   ```
  This will run the tests defined in calculator_test.go and output detailed results.

## Running the Application <server>

  ```
  go run main.go
  ```
  นNote: The server will start and listen on port 3000. It is configured to allow requests from http://localhost:5173 <client>

## API Endpoint Example
  **Request Body
  ```
  {
  "items": {
    "Red set": 1,
    "Green set": 1
  },
  "hasMemberCard": true
  }
  ```
  **Response Body
  ```
  {
    "total": 81
  }
  ```
## 2) Client - Frontend

## Project Structure
- **`src/components/FormComponent.tsx`**: A React component that provides a form for inputting item quantities and member card status.
- **`src/App.tsx`**: The main application component that integrates the `FormComponent` and handles form submission.
1. change directory

   ```
   cd client
   ```

2. Install Dependencies

   ```
   npm install
   ```

3. Run the Development Server

   ```
   npm run dev
   ```
