## Запуск проекта:
```
    go run ./cmd/web
```

## Запуск миграции и seeder:
```
    go run ./cmd/migrate
```

## Генерация Swagger
```
    swag init -g ./cmd/web/main.go
```

## Debug
```
    dlv debug ./cmd/web/ --headless --listen=:40000 --api-version=2 --accept-multiclient
```

[Ссылка на шаблон](https://themewagon.com/themes/newsers/)