# Post BotCmd

## Description
Выполняет команду управления торговым ботом для указанной пары. Позволяет изменять состояние работы бота (запуск, остановка, пауза).

## Request
Method: POST 
URL: http://127.0.0.1:5678/XCoin/PairCmd/{keyWork}/{command}
Headers:
  Content-Type: application/json

Parameters:
  keyWork (integer, required): Уникальный идентификатор торговой пары
  command (string, required): Команда управления ботом. Допустимые значения:
    - Unset: Сброс состояния
    - StopWait: Остановка после завершения текущих операций
    - Stop: Немедленная остановка
    - Runned: Запуск работы

Body: {}

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

Status Codes:
  200: Успешное выполнение команды
  400: Некорректный запрос или неверная команда
  404: Пара не найдена
  500: Внутренняя ошибка сервера

## Example Request
POST http://127.0.0.1:5678/XCoin/PairCmd/1/Runned

## Example Response
{
  "Val": [
    {
      "Key": "Status",
      "Value": "Runned"
    }
  ]
}