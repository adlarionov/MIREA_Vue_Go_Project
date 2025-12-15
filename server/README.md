# Mirea Events API

## Документация

### API

Для корректной работы API необходимо задать переменные окружения:

```bash
DB_NAME=<name>
DB_USER=<user>
DB_HOST=<host>
DB_PASSWORD=<password>

APP_JWT_SECRET=<secret>
```

Локальный запуск в режиме разработки:

```bash
go mod tidy
air
```

### Сервисы

**Auth**

Для работы с защищенными роутами необходим токен, который генерируется по ручке /auth/login.

```yml
POST /auth/login  # Получение токена
POST /auth/register # Регистрация
```

**Events**

Создание, обновление, удаление, получение событий.

```yml
GET /events
GET /events/{id}
POST /events # Обязательное поле OwnerEmail
PUT /events/{id}
DELETE /events/{id}
```

### Docker

Запуск контейнера с зависимостями:

```bash
docker compose up
```

### Swagger

[Swagger](http://localhost:3000/swagger)

В Сваггер подробнее описаны все возможные запросы и ответы для создания событий и взаимодействие с API пользователя.
