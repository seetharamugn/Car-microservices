# Car API

This repository contains a simple Golang implementation of a RESTful API for managing car data. The API includes four HTTP endpoints:

1. GET endpoint to retrieve an existing car
2. GET endpoint to retrieve the list of cars
3. POST endpoint to create a new car
4. PUT endpoint to update an existing car

The API documentation is written using OpenAPI.

## Endpoints

### 1. GET `/GetCar`

Retrieve an existing car by providing its ID as a query parameter.

Example Request:
```
GET /GetCar?id=Pjw7kqKc
```

Example Response:
```json
{
  "StatusCode": 200,
  "Message": "success",
  "Data": {
    "Id": "",
    "Make": "Ford",
    "Model": "F10",
    "Package": "Base",
    "Color": "Silver",
    "Year": 2011,
    "Category": "Truck",
    "Mileage (ml)": 120123,
    "Price (cents)": 1999900
  }
}
```

### 2. GET `/GetCarList`

Retrieve a list of all cars.

Example Request:
```
GET /GetCarList
```

Example Response:
```json
{
  "StatusCode": 200,
  "Message": "success",
  "Data": {
    "JrqSa1f0": {
      "Id": "JrqSa1f0",
      "Make": "Toyoto",
      "Model": "camry",
      "Package": "SE",
      "Color": "White",
      "Year": 2019,
      "Category": "Sedan",
      "Mileage (ml)": 399123,
      "Price (cents)": 2999900
    },
    "q8WNl79p": {
      "Id": "q8WNl79p",
      "Make": "Ford",
      "Model": "F10",
      "Package": "Base",
      "Color": "Silver",
      "Year": 2010,
      "Category": "Truck",
      "Mileage (ml)": 120123,
      "Price (cents)": 1999900
    }
  }
}
```

### 3. POST `/CreateNewCar`

Create a new car by sending a JSON payload with car details.

Example Request:
```
POST /CreateNewCar
Content-Type: application/json

{
    "Make": "Toyoto",
    "Model": "camry",
    "Package": "SE",
    "Color": "White",
    "Year": 2019,
    "Category": "Sedan",
    "Mileage (ml)": 399123,
    "Price (cents)":2999900
}
```

Example Response:
```json
{
  "StatusCode": 200,
  "Message": "success",
  "Data": {
    "Id": "JrqSa1f0",
    "Make": "Toyoto",
    "Model": "camry",
    "Package": "SE",
    "Color": "White",
    "Year": 2019,
    "Category": "Sedan",
    "Mileage (ml)": 399123,
    "Price (cents)": 2999900
  }
}
```

### 4. PUT `/UpdateCar`

Update an existing car by providing its ID in the query parameter and sending a JSON payload with updated car details.

Example Request:
```
PUT /UpdateCar?id=13vUVf1J
Content-Type: application/json

{
    "Make": "Ford",
    "Model": "F10",
    "Package": "Base",
    "Color": "Silver",
    "Year": 2011,
    "Category": "Truck",
    "Mileage (ml)": 120123,
    "Price (cents)": 1999900
}
```

Example Response:
```json
{
  "StatusCode": 200,
  "Message": "success",
  "Data": {
    "Id": "13vUVf1J",
    "Make": "Ford",
    "Model": "F10",
    "Package": "Base",
    "Color": "Silver",
    "Year": 2011,
    "Category": "Truck",
    "Mileage (ml)": 120123,
    "Price (cents)": 1999900
  }
}
```

## Implementation Details

- This API is implemented in Golang using standard libraries, including `net/http`.
- It is assumed to be an internal API and does not include security measures like authentication and authorization.
- For persistence, an in-memory storage system is used, and car data is not persisted between application runs.
- Observability is implemented through logging. Logs are generated for each endpoint and can be configured based on your needs.
- Automated testing for endpoints is included to ensure their functionality.

## Sample Dataset

Below is an example dataset for your reference:

```json
[
  {
    "Id": "13vUVf1J",
    "Make": "Ford",
    "Model": "F10",
    "Package": "Base",
    "Color": "Silver",
    "Year": 2011,
    "Category": "Truck",
    "Mileage (ml)": 120123,
    "Price (cents)": 1999900
  },
  {
    "Id": "JrqSa1f0",
    "Make": "Toyoto",
    "Model": "camry",
    "Package": "SE",
    "Color": "White",
    "Year": 2019,
    "Category": "Sedan",
    "Mileage (ml)": 399123,
    "Price (cents)": 2999900
  },
  {
    "Id": "EvOjEIh2",
    "Make": "Toyoto",
    "Model": "Rav4",
    "Package": "XSE",
    "Color": "Red",
    "Year": 2018,
    "Category": "SUV",
    "Mileage (ml)": 24001,
    "Price (cents)": 227500
  },
  {
    "Id": "MMYmZLI8",
    "Make": "FORD",
    "Model": "BRNCO",
    "Package": "Badiands",
    "Color": "BrutOrange",
    "Year": 2022,
    "Category": "SUV",
    "Mileage (ml)": 1,
    "Price (cents)": 449900
  }
]
```

Feel free to clone this repository and use it as a starting point for your Golang web application that manages car data.