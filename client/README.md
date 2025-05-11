# BMSTU UI

## X-Token

1. Создать файл .env.local
2. Получить токен [отсюда](https://oauth.yandex.ru/authorize?response_type=token&client_id=2fe7c34b1abd45028d39d2160655e0ce)
3. Добавить туда переменную `VITE_X_TOKEN=твой токен`

## [Back Swagger](http://130.193.59.30:8082/)

## Установка зависимостей

```sh
npm install
```

### Дев сборка с hot-reload

```sh
npm run dev
```

### Сборка в прод

```sh
npm run build
```

### Проверка линтером [ESLint](https://eslint.org/)

```sh
npm run lint
```
