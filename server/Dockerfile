# Версия Go
FROM golang:1.23
# Рабочая директория
WORKDIR /app
# Перенос содержимого файлов в рабочую директорию
COPY go.mod go.sum ./
# Запуск установки всех пакетов
RUN go mod download
# Копия всего содержимого проекта
COPY . .
# Run server
# Launch command
CMD [ "mirea-events" ]