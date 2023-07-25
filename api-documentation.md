# Restaurant Management API Documentation
The Restaurant Management API allows restaurants to manage their operations, customers, menu items, orders, deliveries, and customer reviews. This documentation provides details about the models and their URI endpoints.

# API Endpoints
### Users (Authentication)
- The User model represents a user (admin) of the restaurant.
- `POST /api/auth/register`: Register a user
- `POST /api/auth/login`: Login a user and get a JWT token

### Customers 
- `GET /api/customers`: Get all customers.
- `GET /api/customers/{customer_id}`: Get a specific customer.
- `POST /api/customers`: Create a new customer.
- `PUT /api/customers/{customer_id}`: Update a customer.
- `DELETE /api/customers/{customer_id}`: Delete a customer.

### Restaurants
- `GET /api/restaurants`: Get all restaurants.
- `GET /api/restaurants/{restaurant_id}`: Get a specific restaurant.
- `POST /api/restaurants`: Create a new restaurant.
- `PUT /api/restaurants/{restaurant_id}`: Update a restaurant.
- `DELETE /api/restaurants/{restaurant_id}`: Delete a restaurant.

### Menu Items
- `GET /api/restaurants/{restaurant_id}/menuitems`: Get all menu items for a restaurant.
- `GET /api/menuitems/{item_id}`: Get a specific menu item.
- `POST /api/restaurants/{restaurant_id}/menuitems`: Add a new menu item to a restaurant.
- `PUT /api/menuitems/{item_id}`: Update a menu item.
- `DELETE /api/menuitems/{item_id}`: Delete a menu item.

### Orders
- `GET /api/customers/{customer_id}/orders`: Get all orders for a customer.
- `GET /api/orders/{order_id}`: Get a specific order.
- `POST /api/customers/{customer_id}/orders`: Create a new order for a customer.
- `PUT /api/orders/{order_id}`: Update an order.
- `DELETE /api/orders/{order_id}`: Delete an order.

### Delivery Drivers
- `GET /api/drivers`: Get all delivery drivers.
- `GET /api/drivers/{driver_id}`: Get a specific delivery driver.
- `POST /api/drivers`: Create a new delivery driver.
- `PUT /api/drivers/{driver_id}`: Update a delivery driver.
- `DELETE /api/drivers/{driver_id}`: Delete a delivery driver.

### Delivery Information
- `GET /api/orders/{order_id}/delivery`: Get delivery information for an order.
- `POST /api/orders/{order_id}/delivery`: Assign a driver to an order.
- `PUT /api/delivery/{delivery_id}`: Update delivery information.
- `DELETE /api/delivery/{delivery_id}`: Delete delivery information.

### Reviews
- `GET /api/restaurants/{restaurant_id}/reviews`: Get all reviews for a restaurant.
- `GET /api/reviews/{review_id}`: Get a specific review.
- `POST /api/restaurants/{restaurant_id}/reviews`: Add a new review to a restaurant.
- `PUT /api/reviews/{review_id}`: Update a review.
- `DELETE /api/reviews/{review_id}`: Delete a review.
