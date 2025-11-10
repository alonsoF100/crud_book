# Reading Diary

**REST API для ведения дневника прочитанных книг на Go.**

---

## Структура проекта

```
├── cmd
│   └── api
│       └── main.go 
├── deployments
│   └── dev 
│       ├── docker-compose.yml 
│       └── Dockerfile 
├── docs
│   └── postman
│       ├── crud_book_api.postman_collection.json
│       └── Local Development.postman_environment.json
├── go.mod
├── go.sum
├── internal
│   ├── dto
│   │   ├── mappers.go 
│   │   ├── request.go 
│   │   └── response.go
│   ├── handlers
│   │   ├── handlers.go
│   │   └── routing.go
│   ├── models
│   │   └── models.go
│   ├── services
│   │   ├── book_service.go
│   │   ├── interfaces.go
│   │   └── user_service.go
│   └── storage
│       ├── interfaces.go
│       ├── postgres
│       │   ├── book_storage.go 
│       │   ├── migrations
│       │   │   ├── 20251103174124_users.sql
│       │   │   └── 20251103174137_books.sql
│       │   ├── migrator.go 
│       │   ├── storage.go
│       │   └── user_storage.go 
│       └── redis
├── Makefile
├── pkg
│   ├── config
│   ├── logger
│   └── utils
└── README.md
```

---

## Описание основных директорий и файлов

### Корень проекта

* **go.mod**, **go.sum** — зависимости Go
* **Makefile** — автоматизация команд (run, test, build)
* **README.md** — документация проекта

### `cmd/api/`

* **main.go** — точка входа приложения, инициализация зависимостей

### `internal/dto/` (Data Transfer Objects)

* **request.go** — структуры входящих запросов с валидацией
* **response.go** — структуры исходящих ответов
* **mappers.go** — преобразователи моделей (models ↔ DTO)

### `internal/handlers/` (HTTP слой)

* **handlers.go** — обработчики HTTP-запросов (CRUD для книг и пользователей)
* **routing.go** — настройка маршрутов Gin

### `internal/services/` (Бизнес-логика)

* **interfaces.go** — интерфейсы сервисов для абстракции
* **user_service.go** — бизнес-логика пользователей
* **book_service.go** — бизнес-логика книг

### `internal/storage/` (Слой данных)

* **interfaces.go** — интерфейсы репозиториев

### `internal/storage/postgres/` (Реализация PostgreSQL)

* **storage.go** — подключение к БД
* **migrator.go** — управление миграциями
* **user_storage.go**, **book_storage.go** — реализация репозиториев
* **migrations/** — SQL файлы миграций

### `internal/models/`

* **models.go** — доменные модели (User, Book)

### `deployments/dev/`

* **docker-compose.yml** — контейнеризация PostgreSQL и приложения
* **Dockerfile** — сборка образа приложения

### `docs/postman/`

* **crud_book_api.postman_collection.json** — коллекция API для тестирования
* **Local Development.postman_environment.json** — переменные окружения

### `internal/storage/redis/`

* Зарезервировано под кэширование

---

## Поток данных внутри приложения

```
HTTP Request 
    → Gin Router 
    → handlers.CreateBook() 
    → dto.CreateBookRequest 
    → services.BookService.CreateBook() 
    → storage.BookRepository.CreateBook() 
    → PostgreSQL INSERT 
    → models.Book 
    → dto.BookToResponse() 
    → HTTP JSON Response
```

---

## Быстрый старт

1) Создай файл `.env` в корне проекта:

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres  
POSTGRES_DB=library

LOCAL_DATABASE_URL=postgres://postgres:postgres@localhost:5432/library?sslmode=disable
DOCKER_DATABASE_URL=postgres://postgres:postgres@db:5432/library?sslmode=disable
```

2) Подними контейнеры:

```bash
make docker-up
```

---

## Используемые технологии

* **Go + Gin** — бэкенд и HTTP-фреймворк
* **PostgreSQL + pgx/v5** — база данных и драйвер
* **Goose** — миграции базы данных
* **UUID** — уникальные идентификаторы

---

## API Endpoints

### Книги

#### POST `/books` — Создать книгу

**Тело запроса:**

```json
{
  "user_id": "uuid",
  "name": "string",
  "description": "string"
}
```

#### GET `/books/:id` — Получить книгу по ID

#### PUT `/books/:id/status` — Обновить статус книги

**Тело запроса:**

```json
{
  "status": "string"
}
```

#### PUT `/books/:id/rating` — Обновить рейтинг книги

**Тело запроса:**

```json
{
  "rating": 4.5
}
```

#### DELETE `/books/:id` — Удалить книгу

#### GET `/users/:id/books` — Получить книги пользователя

---

### Пользователи

#### POST `/users` — Создать пользователя

**Тело запроса:**

```json
{
  "name": "string", 
  "email": "string"
}
```

#### GET `/users` — Получить всех пользователей

#### GET `/users/:id` — Получить пользователя по ID

#### DELETE `/users/:id` — Удалить пользователя

---

## Модели данных

### Пользователь (User)

```json
{
  "id": "uuid",
  "name": "string", 
  "email": "string"
}
```

### Книга (Book)

```json
{
  "id": "uuid",
  "user_id": "uuid",
  "name": "string",
  "description": "string",
  "rating": 4.5,
  "status": "want|reading|finished"
}
```

---

## API Testing

* **`crud_book_api.postman_collection.json`** — коллекция API эндпоинтов
* **`Local Development.postman_environment.json`** — переменные окружения

Импортируй их в Postman для локального тестирования.

---

## Поддержка

Если проект был полезен — **поставь звёздочку ⭐ на GitHub!**

