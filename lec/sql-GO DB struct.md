# `database/sql.DB` в Go

`*sql.DB` — это не соединение с БД, как многие думают. Это **пул соединений**, который управляет созданием, переиспользованием и закрытием соединений.

```go
db, err := sql.Open("postgres", dsn)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

---

# Основные методы

## 1. Exec

Выполняет запрос, который **не возвращает строки**.

Подходит для:

* INSERT
* UPDATE
* DELETE
* CREATE TABLE
* DROP TABLE

```go
result, err := db.Exec(
    "INSERT INTO users(name) VALUES($1)",
    "John",
)
if err != nil {
    return err
}
```

Возвращает:

```go
sql.Result
```

Из него можно получить:

### LastInsertId()

Поддерживается не всеми драйверами.

```go
id, err := result.LastInsertId()
```

Для PostgreSQL обычно не работает.

Используют:

```sql
RETURNING id
```

---

### RowsAffected()

Количество измененных строк.

```go
count, err := result.RowsAffected()
```

Пример:

```go
res, _ := db.Exec(
    "UPDATE users SET active=false",
)

n, _ := res.RowsAffected()

fmt.Println(n)
```

---

# 2. Query

Выполняет запрос, который возвращает **много строк**.

Возвращает:

```go
*sql.Rows
```

Пример:

```go
rows, err := db.Query(
    "SELECT id, name FROM users",
)
if err != nil {
    return err
}
defer rows.Close()
```

Чтение:

```go
for rows.Next() {
    var id int
    var name string

    err := rows.Scan(&id, &name)
    if err != nil {
        return err
    }

    fmt.Println(id, name)
}
```

Проверка ошибок после цикла:

```go
if err := rows.Err(); err != nil {
    return err
}
```

---

# 3. QueryRow

Для получения **одной строки**.

Возвращает:

```go
*sql.Row
```

Пример:

```go
var name string

err := db.QueryRow(
    "SELECT name FROM users WHERE id=$1",
    1,
).Scan(&name)
```

Если запись не найдена:

```go
if errors.Is(err, sql.ErrNoRows) {
    fmt.Println("user not found")
}
```

---

# 4. Prepare

Создает подготовленный запрос.

```go
stmt, err := db.Prepare(
    "INSERT INTO users(name) VALUES($1)",
)
if err != nil {
    return err
}
defer stmt.Close()
```

Использование:

```go
stmt.Exec("John")
stmt.Exec("Mike")
stmt.Exec("Anna")
```

Полезно, когда один запрос выполняется много раз.

---

# 5. Begin

Начинает транзакцию.

```go
tx, err := db.Begin()
if err != nil {
    return err
}
```

Далее используем методы транзакции:

```go
tx.Exec(...)
tx.Query(...)
tx.QueryRow(...)
```

Подтверждение:

```go
err = tx.Commit()
```

Отмена:

```go
err = tx.Rollback()
```

Шаблон:

```go
tx, err := db.Begin()
if err != nil {
    return err
}

defer tx.Rollback()

_, err = tx.Exec(...)
if err != nil {
    return err
}

_, err = tx.Exec(...)
if err != nil {
    return err
}

return tx.Commit()
```

---

# 6. BeginTx

То же самое, но позволяет настроить транзакцию.

```go
tx, err := db.BeginTx(
    ctx,
    &sql.TxOptions{
        Isolation: sql.LevelSerializable,
        ReadOnly:  true,
    },
)
```

---

# 7. Ping

Проверяет доступность БД.

```go
err := db.Ping()
```

Часто используют при старте приложения.

```go
if err := db.Ping(); err != nil {
    log.Fatal(err)
}
```

---

# 8. PingContext

То же самое, но с контекстом.

```go
ctx, cancel := context.WithTimeout(
    context.Background(),
    5*time.Second,
)
defer cancel()

err := db.PingContext(ctx)
```

---

# 9. Close

Закрывает пул соединений.

```go
err := db.Close()
```

Обычно:

```go
defer db.Close()
```

После закрытия запросы выполнять нельзя.

---

# Методы с Context

Практически у каждого метода есть версия с контекстом.

## ExecContext

```go
db.ExecContext(ctx, query)
```

## QueryContext

```go
db.QueryContext(ctx, query)
```

## QueryRowContext

```go
db.QueryRowContext(ctx, query)
```

## PrepareContext

```go
db.PrepareContext(ctx, query)
```

Для HTTP-сервисов рекомендуется использовать именно версии с контекстом.

---

# Настройка пула соединений

## SetMaxOpenConns

Максимум открытых соединений.

```go
db.SetMaxOpenConns(20)
```

---

## SetMaxIdleConns

Сколько соединений держать свободными.

```go
db.SetMaxIdleConns(10)
```

---

## SetConnMaxLifetime

Максимальное время жизни соединения.

```go
db.SetConnMaxLifetime(time.Hour)
```

После часа соединение будет пересоздано.

---

## SetConnMaxIdleTime

Максимальное время простоя.

```go
db.SetConnMaxIdleTime(10 * time.Minute)
```

---

# Получение статистики

## Stats

```go
stats := db.Stats()

fmt.Println(stats.OpenConnections)
fmt.Println(stats.InUse)
fmt.Println(stats.Idle)
```

Полезно для мониторинга.

---

# Получение отдельного соединения

## Conn

Иногда нужен конкретный connection.

```go
conn, err := db.Conn(ctx)
if err != nil {
    return err
}
defer conn.Close()
```

Дальше:

```go
conn.ExecContext(...)
conn.QueryContext(...)
```

Используется редко.

---

# Что чаще всего используют в реальных проектах

Обычно 90% кода работают только с:

```go
db.ExecContext()
db.QueryContext()
db.QueryRowContext()
db.BeginTx()
db.PingContext()
db.Close()
```

И настройкой пула:

```go
db.SetMaxOpenConns(...)
db.SetMaxIdleConns(...)
db.SetConnMaxLifetime(...)
```

Этого достаточно для большинства REST API, микросервисов и CRUD-приложений.
