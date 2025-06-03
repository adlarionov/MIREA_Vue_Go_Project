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

### Docker

Запуск контейнера с зависимостями:

```bash
docker compose up
```

### Swagger

[Swagger](http://localhost:3000/swagger)

В Сваггер подробнее описаны все возможные запросы и ответы для создания событий и взаимодействие с API пользователя.
