
# Add Trading Pair

## Description
Добавляет новую торговую пару в систему для указанной биржи. После добавления пара становится доступной для торговли.

## Request
Method: POST 
URL: http://127.0.0.1:5678/XCoin/PairAdd/{val1}/{val2}/{birga}/{mode}
Headers:
  Content-Type: application/json

Parameters:
  val1 (string, required): Базовая валюта (например, "BTC")
  val2 (string, required): Котируемая валюта (например, "USDT")
  birga (string, required): Идентификатор биржи
  mode (string, required): Режим работы

Body: Empty
    
## Response
Headers:
  Content-Type: application/json

Body: {}

Status Codes:
  200: Пара успешно добавлена
  400: Некорректные параметры запроса
  409: Пара уже существует
  500: Внутренняя ошибка сервера

## Example Request
POST http://127.0.0.1:5678/XCoin/PairAdd/BTC/USDT/0/1

## Example Response
{
  "success": true,
  "message": "Pair BTC/USDT added successfully"
}