# Дневник чтения (Reading Diary)

REST API для ведения дневника прочитанных книг на Go.

❯ tree
├── cmd
│   └── api
│       └── main.go 
├── deployments
│   └── dev 
│       ├── docker-compose.yml 
│       └── Dockerfile 
├── docs
│   └── postman
│       ├── crud_book_api.postman_collection.json
│       └── Local Development.postman_environment.json
├── go.mod
├── go.sum
├── internal
│   ├── dto
│   │   ├── mappers.go 
│   │   ├── request.go 
│   │   └── response.go
│   ├── handlers
│   │   ├── handlers.go
│   │   └── routing.go
│   ├── models
│   │   └── models.go
│   ├── services
│   │   ├── book_service.go
│   │   ├── interfaces.go
│   │   └── user_service.go
│   └── storage
│       ├── interfaces.go
│       ├── postgres
│       │   ├── book_storage.go 
│       │   ├── migrations
│       │   │   ├── 20251103174124_users.sql
│       │   │   └── 20251103174137_books.sql
│       │   ├── migrator.go 
│       │   ├── storage.go
│       │   └── user_storage.go 
│       └── redis
├── Makefile
├── pkg
│   ├── config
│   ├── logger
│   └── utils
└── README.md

Корневые файлы:
go.mod, go.sum          - зависимости Go
Makefile                - автоматизация команд (run, test, build)
README.md               - документация проекта

cmd/api/:
main.go                 - точка входа приложения, инициализация зависимостей

internal/dto/ (Data Transfer Objects):
request.go             - структуры входящих запросов с валидацией
response.go            - структуры исходящих ответов
mappers.go             - преобразователи models → DTO и обратно

internal/handlers/ (HTTP слой):
handlers.go            - обработчики HTTP запросов (CRUD для книг и пользователей)
routing.go             - настройка маршрутов Gin

internal/services/ (Бизнес-логика):
interfaces.go          - интерфейсы сервисов для абстракции
user_service.go        - бизнес-логика пользователей
book_service.go        - бизнес-логика книг

internal/storage/ (Слой данных):
interfaces.go          - интерфейсы репозиториев для абстракции хранилищ

internal/storage/postgres/ (Реализация PostgreSQL):
storage.go             - подключение к БД, управление соединением
user_storage.go        - реализация UserRepository для PostgreSQL
book_storage.go        - реализация BookRepository для PostgreSQL
migrator.go            - управление миграциями базы данных
migrations/            - SQL файлы миграций (up/down)

internal/models/:
models.go              - доменные модели (User, Book)

deployments/dev/:
docker-compose.yml     - контейнеризация PostgreSQL и приложения
Dockerfile             - образ приложения

docs/postman/
postman_collection.json - коллекция API endpoints для тестирования
postman_environment.json - environment переменные для Postman

internal/storage/redis/
(зарезервировано)      - будущая реализация кэширования

Поток данных внутри приложения:
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
    
## Быстрый старт

1) Заполнить
.env :
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres  
POSTGRES_DB=library

LOCAL_DATABASE_URL=postgres://postgres:postgres@localhost:5432/library?sslmode=disable
DOCKER_DATABASE_URL=postgres://postgres:postgres@db:5432/library?sslmode=disable

2) В корневой папке выполнить make docker-up

## Технологии

- Go + Gin - бэкенд и HTTP фреймворк
- PostgreSQL + pgx/v5 - база данных и драйвер
- Goose - миграции базы данных
- UUID - уникальные идентификаторы

## API Endpoints

### Книги

POST /books - Создать книгу
Тело: 
{
  "user_id": "uuid",
  "name": "string",
  "description": "string"
}

GET /books/:id - Получить книгу по ID

PUT /books/:id/status - Обновить статус книги
Тело: 
{
  "status": "string"
}

PUT /books/:id/rating - Обновить рейтинг книги
Тело: 
{
  "rating": number
}

DELETE /books/:id - Удалить книгу

GET /users/:id/books - Получить книги пользователя

### Пользователи

POST /users - Создать пользователя
Тело: 
{
  "name": "string", 
  "email": "string"
}

GET /users - Получить всех пользователей

GET /users/:id - Получить пользователя по ID

DELETE /users/:id - Удалить пользователя

## Модели данных

### Пользователь (User)

{
  "id": "uuid",
  "name": "string", 
  "email": "string"
}

### Книга (Book)

{
  "id": "uuid",
  "user_id": "uuid",
  "name": "string",
  "description": "string",
  "rating": 4.5,
  "status": "want|reading|finished"
}

## API Testing

postman_collection.json - коллекция API endpoints для тестирования
postman_environment.json - environment переменные для Postman

⭐ Не забудьте поставить звезду репозиторию если проект был полезен!