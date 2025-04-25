# GetListPair2

## Description
Get list of all trading pairs with their full information and settings, including current status, orders, and configuration.

## Request
Method: GET 
URL: http://127.0.0.1:5678/XCoin/GetListPair2
Headers:
  Content-Type: application/json

## Response
Headers:
  Content-Type: application/json

Body:
[
  {
    "Description": "string",       // Pair description (e.g. "ATOMFDUSD up FDUSD")
    "KeyWork": 0,                // Unique identifier for the pair
    "KolFill": 0,              // Number of filled orders
    "KolFillColor": {           // Color settings for filled orders display
      "A": 255,
      "B": 0,
      "G": 0,
      "R": 0
    },
    "KolOtkup": 0,             // Number of buyback orders
    "KolOtkupColor": {         // Color settings for buyback orders display
      "A": 255,
      "B": 0,
      "G": 69,
      "R": 255
    },
    "KolZakup": 0,             // Number of purchase orders
    "KolZakupColor": {         // Color settings for purchase orders display
      "A": 255,
      "B": 50,
      "G": 205,
      "R": 50
    },
    "Rasst": "string",          // Distance value
    "RasstColor": {            // Color settings for distance display
      "A": 255,
      "B": 0,
      "G": 0,
      "R": 255
    },
    "Setting": {                // Pair trading settings
      "Birga": 0,              // Exchange identifier
      "KolOrder": 99,          // Maximum number of orders
      "SVal": {                // Symbol information
        "BaseAsset": "string",  // Base currency (e.g. "ATOM")
        "QuoteAsset": "string", // Quote currency (e.g. "FDUSD")
        "Symbol": "string",      // Trading pair symbol
        "LotSizeFilter": {      // Order size restrictions
          "MaxQuantity": 0,
          "MinQuantity": 0,
          "StepSize": 0
        },
        "MinNotionalFilter": {  // Minimum order value
          "MinNotional": 0
        },
        "PriceFilter": {        // Price restrictions
          "MaxPrice": 0,
          "MinPrice": 0,
          "TickSize": 0
        }
      },
      "Val1": "string",         // Base currency code
      "Val2": "string",         // Quote currency code
      "depositOrder": 0,        // Deposit amount for orders
      "firstStep": 0,          // Initial step size
      "martingale": 0,         // Martingale coefficient
      "orderStep": 0,          // Order step size
      "profit": 0,             // Target profit percentage
      "reload": 0,             // Reload threshold
      "timeActivate": "string",  // Activation time
      "timeDeActivate": "string" // Deactivation time
    },
    "Status": "string",         // Current status (e.g. "Runned")
    "StatusInt": 0,            // Status code
    "Val1Color": {             // Color settings for base currency
      "A": 255,
      "B": 0,
      "G": 0,
      "R": 0
    },
    "Val2Color": {             // Color settings for quote currency
      "A": 255,
      "B": 0,
      "G": 0,
      "R": 0
    }
  }
]

Status Codes:
  200: Success
  400: Bad Request
  500: Internal Server Error

## Example Response
{
  "Description": "ATOMFDUSD up FDUSD",
  "KeyWork": 1,
  "KolFill": 7,
  "Status": "Runned",
  "StatusInt": 1,
  "Setting": {
    "Birga": 0,
    "Val1": "ATOM",
    "Val2": "FDUSD",
    "depositOrder": 60,
    "profit": 0.5
  }
}
