# Get PairCurSetting

## Description
Получает текущие настройки торговой пары по её уникальному идентификатору (keyWork). Возвращает массив пар ключ-значение с параметрами пары.

## Request
Method: GET 
URL: http://127.0.0.1:5678/XCoin/PairCurSetting/{keyWork}
Headers:
  Content-Type: application/json

Parameters:
  keyWork (integer, required): Уникальный идентификатор торговой пары
    
## Response
Headers:
  Content-Type: application/json

Body:
{
  "Val": [
    {
      "Key": "string",    // Название параметра
      "Value": "string"  // Значение параметра
    }
  ]
}

Доступные параметры в ответе:
- BestAskPrice: Лучшая цена продажи
- BestBidPrice: Лучшая цена покупки
- Status: Текущий статус пары (StopWait, Runned, Stop)
- AllProfit: Общая прибыль
- AllCountSession: Количество сессий
- Birga: Идентификатор биржи
- Val1: Базовая валюта
- Val2: Котируемая валюта

Status Codes:
  200: Успешное выполнение
  400: Некорректный запрос
  404: Пара не найдена
  500: Внутренняя ошибка сервера

## Example Response
{
  "Val": [
    {
      "Key": "BestAskPrice",
      "Value": "4.28000000"
    },
    {
      "Key": "BestBidPrice",
      "Value": "4.27800000"
    },
    {
      "Key": "Status",
      "Value": "StopWait"
    },
    {
      "Key": "AllProfit",
      "Value": "17.0572"
    }
  ]
}