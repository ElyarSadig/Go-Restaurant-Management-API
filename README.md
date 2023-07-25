# Restaurant Management RESTful API
This is a RESTful API for managing a restaurant's operations, including customers, restaurants, menu items, orders, delivery information, and customer reviews. It is built using the Gin framework for handling HTTP requests and GORM as the ORM (Object-Relational Mapping) library to interact with the database.

## Features
- Customers: Manage customer information such as name, email, phone, and address.
- Restaurants: Store information about different restaurants, including their name, location, and contact details.
- Menu Items: Add, update, and delete menu items for each restaurant, along with their details and pricing.
- Orders: Create, update, and delete orders for customers with the option for delivery information.
- Delivery Drivers: Manage delivery drivers who can be assigned to deliver orders.
- Reviews: Allow customers to leave reviews for specific restaurants.

## Prerequisites
- Go (1.16+)
- `github.com/gin-gonic/gin` (Gin framework)
- `gorm.io/gorm` (GORM library)
- `github.com/dgrijalva/jwt-go` (JWT library)
- `gorm.io/driver/sqlite`
- `github.com/joho/godotenv`

## Installation
1. Clone the repository
    ```gitbash
    git clone https://github.com/ElyarSadig/restaurant-management-api/
    cd restaurant-management-api
    ```

2. Install the required Go packages:
    ```bash
    go get -u github.com/gin-gonic/gin
    go get -u gorm.io/gorm
    go get -u github.com/dgrijalva/jwt-go
    go get -u github.com/joho/godotenv
    go get -u gorm.io/driver/sqlite
    ```

3. Run the application
    ```
    go run main.go
    ```

## API Endpoints
The API provides endpoints to perform various operations on customers, restaurants, menu items, orders, delivery drivers, delivery information, and customer reviews. Refer to the [API documentation](api-documentation.md) for detailed information on each endpoint, request, and response formats.

## Authentication
The API supports authentication using JSON Web Tokens (JWT). To access protected endpoints, clients must include a valid JWT token in the Authorization header of their requests.

## Contributing
Contributions are welcome! Please follow the standard guidelines for contributing to open-source projects.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgements
Special thanks to Gin and GORM for providing excellent tools to build this API.
