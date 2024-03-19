# BookShop-gRPC

Это fullstack проект для онлайн магазина книг. Дизайн, фронт, бэк все было сделано мной. В этой ветке(main) находится весь проект, только фронт(вместе с просмотром через github pages) будет в ветке Front, а бэк в ветке Back. 

## Стек технологий
# Бэк:
- Go
- gRPC
- gRPC-Gateway
- MongoDB
- Redis
- Docker
# Фронт:
- React?
- TypeScript?
- AnimeJS?
- SCSS

## Структура проекта

- `api/proto/`: Этот каталог содержит файлы `.proto`, которые определяют сервисы gRPC и сообщения, используемые в этом проекте.
- `cmd/web/`: Здесь находится основной исполняемый файл проекта.
- `internal/`: Этот каталог содержит внутренний код проекта.
- `internal/models` : Здесь находятся модели данных.
- `pkg/grpc/`: Этот каталог содержит сгенерированный код gRPC.

## Запуск проекта

### 1. Установите зависимости
```sh
go mod download
```

### 2. Запустите/установите MongoDB и Redis
Создайте .env файл в корне проекта и добавьте туда следующие переменные:
```sh
MONGO_URI=<ваша ссылка>
```

### 3. Сгенерируйте код gRPC

## Генерация кода gRPC
```sh
protoc --proto_path=api/proto --go_out=pkg/grpc --go_opt=paths=source_relative --go-grpc_out=pkg/grpc --go-grpc_opt=paths=source_relative api/proto/bookshop_*.proto --grpc-gateway_out=pkg/grpc --grpc-gateway_opt=paths=source_relative
```

### 3. Запустите проект
```sh
go run ./cmd/web
```
