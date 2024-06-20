# receipt-processor
This project is a web service implemented in Go (Golang) that processes retail receipts and calculates points based on specific criteria. The service provides two endpoints as specified in the API documentation: one for processing receipts and another for retrieving points for a processed receipt.

# Description
The Receipt Processor Web Service allows users to submit retail receipts in JSON format, processes the receipt to calculate points based on predefined rules, and provides an endpoint to retrieve the points for a specific receipt.

# API Endpoints
POST /receipts/process
Processes a receipt and returns a unique ID for the receipt.

URL: /receipts/process
Method: POST
Request Body:
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
