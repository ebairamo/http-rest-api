# FROM golang:1.23

# WORKDIR /app

# # Устанавливаем gcc, который нужен для CGO
# RUN apt update && apt install -y gcc

# # Копируем файлы проекта
# COPY . .

# # Включаем CGO и устанавливаем зависимости
# ENV CGO_ENABLED=1
# RUN go mod tidy

# # Собираем бинарник
# RUN go build -o hot-coffee

# CMD ["./hot-coffee"]

# FROM golang:1.23

# WORKDIR /app

# # Устанавливаем gcc и Delve
# RUN apt update && apt install -y gcc && go install github.com/go-delve/delve/cmd/dlv@latest

# # Копируем файлы проекта
# COPY . .

# # Включаем CGO и устанавливаем зависимости
# ENV CGO_ENABLED=1
# RUN go mod tidy

# # Собираем бинарник (но он не запустится сразу)
# RUN go build -gcflags "all=-N -l" -o hot-coffee

# # Запускаем Delve в режиме отладки
# CMD ["dlv", "exec", "/app/hot-coffee", "--headless", "--listen=:40000", "--api-version=2", "--accept-multiclient"]


FROM golang:1.23.2

# Установка Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Копирование исходного кода
COPY . /app
WORKDIR /app

# Сборка приложения
RUN go build -o hot-coffee .

# Команда для запуска Delve
CMD ["dlv", "exec", "/app/hot-coffee", "--headless", "--listen=:40000", "--api-version=2", "--log", "--"]