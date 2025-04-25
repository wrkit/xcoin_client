# XCoin Client

XCoin Client - Go библиотека для взаимодействия с XCoin Bot API. Библиотека предоставляет интерфейс для управления торговыми парами и выполнения торговых операций.

## Установка

```bash
go get github.com/wrkit/experiment/xcoin_client
```

## Быстрый старт

```go
package main

import (
    "context"
    "fmt"
    "time"
    "github.com/wrkit/xcoin_client"
)

func main() {
    // Создание нового клиента
    cfg := xcoin_client.NewConfig("http://your-bot-url:port")
    client := xcoin_client.NewClient(cfg)

    // Создание контекста с таймаутом
    ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
    defer cancel()

    // Получение версии бота
    version, err := client.GetVersion(ctx)
    if err != nil {
        fmt.Printf("Ошибка получения версии: %v\n", err)
        return
    }
    fmt.Printf("Версия: %+v\n", version)
}
```

## Основные возможности

### Конфигурация клиента

```go
// Базовая конфигурация
cfg := xcoin_client.NewConfig("http://bot-url:port")

// Настройка таймаута
cfg.WithTimeout(60 * time.Second)

// Использование собственного HTTP клиента
customClient := &http.Client{}
cfg.WithHTTPClient(customClient)
```

### Работа с торговыми парами

```go
// Получение списка торговых пар
pairs, err := client.GetPairList2(ctx)

// Добавление новой торговой пары
err = client.AddPair(ctx, "BTC", "USD", "binance", "mode")

// Получение настроек пары
settings, err := client.GetPairCurrentSettings(ctx, pairID)

// Получение ордеров пары
orders, err := client.GetPairOrders(ctx, pairID)

// Обновление настроек пары
err = client.UpdatePairSettings(ctx, pairID, newSettings)
```

## Архитектура

Библиотека построена на следующих принципах:

1. Конфигурируемый клиент с поддержкой настройки URL бота и HTTP клиента
2. Неблокирующие запросы с поддержкой контекстов
3. Типизированные структуры данных для всех ответов API

## Обработка ошибок

Все методы возвращают ошибку в качестве второго значения. Рекомендуется всегда проверять ошибки:

```go
result, err := client.GetVersion(ctx)
if err != nil {
    // Обработка ошибки
    return err
}
```

## Лицензия

MIT