# Order Service gRPC

## О проекте

**Order Service** — это **pet-проект**, реализующий отдельный gRPC-микросервис для управления заказами в
интернет-магазине.

Проект написан на **Go** с использованием **gRPC**, **Protocol Buffers** и **PostgreSQL**. Основная цель проекта —
изучение разработки микросервисов, построения многослойной архитектуры и работы с gRPC.

Проект создавался исключительно в учебных целях и позволил на практике разобраться с построением backend-сервисов на Go.

---

# Стек технологий

- Go
- gRPC
- Protocol Buffers
- PostgreSQL
- sqlx
- UUID
- Zap
- Logrus

---

# Архитектура

Проект построен по классической многослойной архитектуре.

```
Client
   │
   ▼
gRPC Handler
   │
   ▼
Service
   │
   ▼
Repository
   │
   ▼
PostgreSQL
```

Каждый слой отвечает только за свою область ответственности:

- **Handler** — обработка gRPC-запросов
- **Service** — бизнес-логика
- **Repository** — работа с PostgreSQL
- **Database** — подключение к базе данных и выполнение миграций

---

# Возможности

Сервис предоставляет следующие возможности:

- Создание заказа
- Получение заказа по ID
- Получение списка заказов с фильтрацией
- Изменение статуса заказа
- Удаление заказа

Также реализованы:

- gRPC Reflection
- Panic Recovery Interceptor
- Централизованное логирование
- Загрузка конфигурации из `.env`
- Работа с PostgreSQL
- SQL-миграции

---

# Переменные окружения (.env)

```env
SERVER_PORT=50051

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=order_service
DB_SSLMODE=disable
```

---

# Запуск проекта

### Установка зависимостей

```bash
go mod download
```

### Выполнение миграции

```bash
psql -h localhost -U postgres -d order_service -f migrations/001_init.sql
```

### Запуск сервиса

```bash
go run main.go
```

После запуска gRPC-сервис будет доступен по адресу:

```
localhost:50051
```

---

# gRPC API

| Метод | Описание |
|--------|----------|
| CreateOrder | Создание нового заказа |
| GetOrderId | Получение заказа по ID |
| GetOrderList | Получение списка заказов |
| UpdateOrderStatus | Обновление статуса заказа |
| DeleteOrder | Удаление заказа |

---

# Модель заказа

```protobuf
Order {
    order_id
    product_id
    seller_id
    buyer_id
    status
    pickup_point_id
    estimated_delivery_time
    created_at
    updated_at
}
```

---

# Статусы заказа

Поддерживаются следующие состояния:

```
ORDER_STATUS_UNSPECIFIED
ORDER_STATUS_COLLECTING
ORDER_STATUS_DELIVERING
ORDER_STATUS_READY_FOR_PICKUP
ORDER_STATUS_DELIVERED
ORDER_STATUS_RETURNED
```

Если новый статус не передан при обновлении заказа, сервис автоматически переводит заказ на следующий этап.

---

# Фильтрация заказов

Метод `GetOrderList` поддерживает фильтрацию по:

- Product ID
- Seller ID
- Buyer ID
- Pickup Point ID
- Delivery Time
- Status

Все параметры являются необязательными.

---

# База данных

Используется PostgreSQL.

Таблица `orders` содержит:

- Order ID
- Product ID
- Seller ID
- Buyer ID
- Status
- Pickup Point ID
- Delivery Time
- Created At
- Updated At

Для ускорения поиска созданы индексы по:

- Buyer ID
- Seller ID
- Status

---

# Особенности реализации

В рамках проекта реализованы:

- многослойная архитектура (Handler → Service → Repository);
- gRPC API;
- Protocol Buffers;
- panic recovery interceptor;
- gRPC Reflection;
- централизованное логирование;
- генерация UUID для заказов;
- динамическое построение SQL-запросов при фильтрации;
- использование `context.Context` на всех уровнях приложения.

---

# Генерация gRPC-кода

```bash
protoc \
--proto_path=api/proto \
--go_out=pkg/generated/order \
--go_opt=paths=source_relative \
--go-grpc_out=pkg/generated/order \
--go-grpc_opt=paths=source_relative \
api/proto/order.proto
```

---

# Цель проекта

Данный сервис является **pet-проектом**, созданным для практического изучения:

- разработки gRPC-сервисов на Go;
- Protocol Buffers;
- многослойной архитектуры приложений;
- взаимодействия с PostgreSQL через sqlx;
- организации CRUD-операций;
- проектирования микросервисной архитектуры.

Проект не является production-решением и создавался как учебная работа для закрепления навыков backend-разработки и
знакомства с современным стеком Go.

