# Get Orders for pair

## Description
Получает список текущих ордеров для торговой пары по её уникальному идентификатору (keyWork). Возвращает массив объектов с детальной информацией о каждом ордере.

## Request
Method: GET 
URL: http://127.0.0.1:5678/XCoin/PairCurOrders/{keyWork}
Headers:
  Content-Type: application/json

Parameters:
  keyWork (integer, required): Уникальный идентификатор торговой пары
    
## Response
Headers:
  Content-Type: application/json

Body: 
[
  {
    "ClientOrderId": "string",     // Уникальный ID ордера на клиенте
    "Comission1": 0,             // Комиссия по базовой валюте
    "Comission2": 0,             // Комиссия по котируемой валюте
    "CreateTime": "string",      // Время создания ордера
    "DataUpdate": "string",      // Время последнего обновления
    "Level": 0,                  // Уровень ордера
    "Mode": 0,                   // Режим работы
    "OrderId": 0,                // ID ордера на бирже
    "Price": 0,                  // Цена ордера
    "Quantity": 0,               // Количество
    "QuantityFilled": 0,         // Исполненное количество
    "Side": 0,                   // Сторона ордера (0 - покупка, 1 - продажа)
    "Status": 0,                 // Статус ордера
    "Symbol": "string",          // Символ торговой пары
    "Type": 0,                   // Тип ордера
    "keySetka": 0,               // Ключ сетки
    "myManual": 0                // Флаг ручного ордера
  }
]

Status Codes:
  200: Успешное выполнение
  400: Некорректный запрос
  404: Пара не найдена
  500: Внутренняя ошибка сервера

## Example Response
{
  "ClientOrderId": "Tg4N18CMA36a0UGdEwqadKETnTvWPfzt",
  "CreateTime": "/Date(1703431889357+0200)/",
  "OrderId": 20770493,
  "Price": 11.5550000000,
  "Quantity": 1.7300000000,
  "QuantityFilled": 1.7300000000,
  "Side": 0,
  "Status": 2,
  "Symbol": "ATOMFDUSD"
}