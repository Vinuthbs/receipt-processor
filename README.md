# receipt-processor
This project is a web service implemented in Go (Golang) that processes retail receipts and calculates points based on specific criteria. The service provides two endpoints as specified in the API documentation: one for processing receipts and another for retrieving points for a processed receipt.

# Description
The Receipt Processor Web Service allows users to submit retail receipts in JSON format, processes the receipt to calculate points based on predefined rules, and provides an endpoint to retrieve the points for a specific receipt.

# API Endpoints
POST /receipts/process
Processes a receipt and returns a unique ID for the receipt.
URL: /receipts/process
Method: POST
Request Body: (Example 1)
```json
{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}
```
Response:
```json
{
    "id": "9375be1a-4e4b-4d21-b1e1-4600348b3146"
}
```
GET /receipts/:id/points
Retrieves the points for a processed receipt by its unique ID.
URL: /receipts/:id/points
Method: GET
Response:
```json
{
    "points": 15
}
```

(Example 2)
Request Body:
```json
{
    "retailer": "Target",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "13:13",
    "total": "1.25",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
    ]
}
```
Response:
```json
{
    "id": "3bd56cb6-9841-4670-ae45-f74d8a0012d5"
}
```

GET response:
```json
{
    "points": 31
}
```

# Testing with Postman
1. Open Postman and create a new request.
2. Set the method to POST and the URL to http://localhost:8080/receipts/process
3. Set the request body to the example JSON provided above.
4. Send the request to receive the unique receipt ID.
5. Create another request with the method set to GET and the URL to http://localhost:8080/receipts/{id}/points, replacing {id} with the received receipt ID.
6. Send the request to get the points for the processed receipt.


