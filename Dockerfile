# Этап сборки
FROM golang:1.22.4-alpine3.19 as build

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Копируем исходные файлы в контейнер
COPY . .

# Устанавливаем дополнительные пакеты, если необходимо
RUN apk --no-cache add build-base

# Сборка приложения
RUN GOOS=linux GOARCH=amd64 go build -o /app/cmd/app/main ./cmd/app

# Этап исполнения
FROM alpine:latest

# Устанавливаем необходимые библиотеки для работы с PostgreSQL и сертификаты
RUN apk --no-cache add ca-certificates

# Рабочая директория
WORKDIR /root/

# Копируем собранное приложение из этапа сборки
COPY --from=build /app/cmd/app/main .

# Экспонируем порт для приложения
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
