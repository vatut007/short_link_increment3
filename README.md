# URL Shortener

Сервис сокращения ссылок на Go. Принимает длинный URL и возвращает короткий, по которому выполняется редирект на оригинал.

## Стек

- Go 1.22+
- [chi v5](https://github.com/go-chi/chi) — HTTP-роутер
- `sync.Map` — хранилище в памяти (без персистентности)

## API

### Создать короткую ссылку

```
POST /
Content-Type: text/plain

https://www.example.com/some/long/path
```

**Ответ `201 Created`:**
```
http://localhost:8080/aB3dEf9Z
```

---

### Перейти по короткой ссылке

```
GET /{short_code}
```

**Ответ `307 Temporary Redirect`** — редирект на оригинальный URL.

Если код не найден — `404 Not Found`.

## Запуск

```bash
go run ./cmd/shortener
```

Сервер запустится на порту `:8080`.

## Тесты

```bash
go test ./...
```

## Структура проекта

```
cmd/shortener/   — точка входа
handlers/        — HTTP-обработчики и роутер
store/           — хранилище (sync.Map)
utils/           — генерация короткого кода (SHA-256 + base64)
```
