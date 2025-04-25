# Post PairEditSetting

## Description
Обновляет настройки торговой пары по указанному keyWork. Позволяет изменять различные параметры работы пары, такие как профит, депозит, стратегия и другие настройки.

## Request
Method: POST 
URL: http://127.0.0.1:5678/XCoin/PairCurSetting/{keyWork}
Headers:
  Content-Type: application/json

Parameters:
  keyWork (integer, required): Уникальный идентификатор торговой пары

Body: 
{
  "Val": [
    {
      "Key": "string",    // Название параметра
      "Value": "string"  // Значение параметра
    }
  ]
}

Доступные параметры для изменения:
- profit: Целевой процент прибыли
- depositOrder: Размер депозита для ордеров
- firstStep: Начальный шаг
- orderStep: Шаг между ордерами
- martingale: Коэффициент мартингейла
- strateg: Номер стратегии
- limitDeposit: Лимит депозита
- trailing: Настройки трейлинга

## Response
Headers:
  Content-Type: application/json

Body: {}

Status Codes:
  200: Настройки успешно обновлены
  400: Некорректный запрос
  404: Пара не найдена
  500: Внутренняя ошибка сервера

## Example Request
POST http://127.0.0.1:5678/XCoin/PairCurSetting/1
{
  "Val": [
    {
      "Key": "profit",
      "Value": "0.5"
    },
    {
      "Key": "depositOrder",
      "Value": "100"
    }
  ]
}

## Example Response
{
  "success": true,
  "message": "Settings updated successfully"
}
        