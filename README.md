# Дневник чтения (Reading Diary)

REST API для ведения дневника прочитанных книг на Go.

## Быстрый старт

1. Клонировать репозиторий
git clone https://github.com/alonsoF100/crud_book.git
cd crud_book

2. Установить зависимости
go mod tidy

3. Настроить базу данных
echo "DATABASE_URL=postgres://user:pass@localhost:5432/reading_diary" > .env

4. Запустить приложение
go run main.go

Приложение будет доступно по адресу: http://localhost:8080

## Технологии

- Go + Gin - бэкенд и HTTP фреймворк
- PostgreSQL + pgx/v5 - база данных и драйвер
- Goose - миграции базы данных
- UUID - уникальные идентификаторы

## API Endpoints

### Книги

POST /books - Создать книгу
Тело: {"user_id": "uuid", "name": "string", "description": "string"}

GET /books/:id - Получить книгу по ID

PUT /books/:id/status - Обновить статус книги
Тело: {"status": "string"}

PUT /books/:id/rating - Обновить рейтинг книги
Тело: {"rating": number}

DELETE /books/:id - Удалить книгу

GET /users/:id/books - Получить книги пользователя

### Пользователи

POST /users - Создать пользователя
Тело: {"name": "string", "email": "string"}

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

## Примеры использования

Создать пользователя:
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name":"Анна", "email":"anna@example.com"}'

Создать книгу:
curl -X POST http://localhost:8080/books -H "Content-Type: application/json" -d '{"user_id":"uuid", "name":"Книга", "description":"Описание"}'

Обновить статус книги:
curl -X PUT http://localhost:8080/books/uuid/status -H "Content-Type: application/json" -d '{"status":"finished"}'

## Структура проекта

crud_book/
├── handlers/     # HTTP обработчики
├── storage/      # Работа с базой данных
├── migrations/   # SQL миграции
├── main.go       # Точка входа
└── README.md

## Статусы книг

- want - запланировано к прочтению
- reading - в процессе чтения  
- finished - прочитано

---

⭐ Не забудьте поставить звезду репозиторию если проект был полезен!