## Что такое JOIN

`JOIN` — это операция, которая позволяет объединять данные из нескольких таблиц по определенному условию.

Без `JOIN` реляционные базы данных были бы практически бесполезны, потому что данные обычно нормализованы и хранятся в разных таблицах.

Например:

**customer**

| id | full_name |
| -- | --------- |
| 1  | Ivan      |
| 2  | Anna      |

**orders**

| id | customer_id |
| -- | ----------- |
| 1  | 1           |
| 2  | 2           |

В таблице `orders` хранится только `customer_id`, а имя клиента находится в таблице `customer`.

Чтобы получить заказ вместе с именем клиента:

```sql
SELECT *
FROM orders o
JOIN customer c ON c.id = o.customer_id;
```

---

# Как работает JOIN

Представь, что PostgreSQL берет строку из первой таблицы и ищет подходящие строки во второй таблице по условию после `ON`.

```sql
SELECT *
FROM orders o
JOIN customer c
ON c.id = o.customer_id;
```

Условие:

```sql
c.id = o.customer_id
```

Если совпадение найдено — строки объединяются.

---

# INNER JOIN

Самый популярный JOIN.

Возвращает только совпавшие записи.

### Данные

**customer**

| id | name |
| -- | ---- |
| 1  | Ivan |
| 2  | Anna |
| 3  | Bob  |

**orders**

| id | customer_id |
| -- | ----------- |
| 1  | 1           |
| 2  | 2           |

### Запрос

```sql
SELECT c.name, o.id
FROM customer c
INNER JOIN orders o
ON o.customer_id = c.id;
```

### Результат

| name | id |
| ---- | -- |
| Ivan | 1  |
| Anna | 2  |

Bob не попадет в результат, потому что у него нет заказов.

---

# LEFT JOIN

Возвращает ВСЕ строки из левой таблицы.

Если совпадения справа нет — подставляет NULL.

### Запрос

```sql
SELECT c.name, o.id
FROM customer c
LEFT JOIN orders o
ON o.customer_id = c.id;
```

### Результат

| name | id   |
| ---- | ---- |
| Ivan | 1    |
| Anna | 2    |
| Bob  | NULL |

---

## Где используют LEFT JOIN

Очень часто.

Например:

Показать всех пользователей и их заказы.

```sql
SELECT *
FROM users u
LEFT JOIN orders o
ON o.user_id = u.id;
```

Даже если заказов нет, пользователь будет показан.

---

# RIGHT JOIN

То же самое, но наоборот.

Возвращает все строки из правой таблицы.

```sql
SELECT *
FROM customer c
RIGHT JOIN orders o
ON o.customer_id = c.id;
```

На практике используется редко.

Почти всегда можно переписать через `LEFT JOIN`, просто поменяв таблицы местами.

---

# FULL JOIN

Возвращает:

* все строки слева
* все строки справа

Если совпадений нет — ставит NULL.

### Пример

**A**

| id |
| -- |
| 1  |
| 2  |

**B**

| id |
| -- |
| 2  |
| 3  |

```sql
SELECT *
FROM A
FULL JOIN B
ON A.id = B.id;
```

Результат:

| A.id | B.id |
| ---- | ---- |
| 1    | NULL |
| 2    | 2    |
| NULL | 3    |

---

# CROSS JOIN

Декартово произведение.

Каждая строка первой таблицы соединяется с каждой строкой второй.

### Данные

**colors**

| color |
| ----- |
| Red   |
| Blue  |

**sizes**

| size |
| ---- |
| S    |
| M    |

### Запрос

```sql
SELECT *
FROM colors
CROSS JOIN sizes;
```

### Результат

| color | size |
| ----- | ---- |
| Red   | S    |
| Red   | M    |
| Blue  | S    |
| Blue  | M    |

Количество строк:

```text
rowsA × rowsB
```

1000 × 1000 = 1 000 000 строк.

Поэтому с `CROSS JOIN` нужно быть осторожным.

---

# SELF JOIN

Таблица соединяется сама с собой.

Например, сотрудники и их начальники.

**employees**

| id | name | manager_id |
| -- | ---- | ---------- |
| 1  | Ivan | NULL       |
| 2  | Anna | 1          |
| 3  | Bob  | 1          |

```sql
SELECT
    e.name,
    m.name AS manager
FROM employees e
LEFT JOIN employees m
ON e.manager_id = m.id;
```

Результат:

| employee | manager |
| -------- | ------- |
| Ivan     | NULL    |
| Anna     | Ivan    |
| Bob      | Ivan    |

---

# JOIN нескольких таблиц

Обычно приходится соединять больше двух таблиц.

На примере твоих таблиц:

```sql
SELECT
    c.full_name,
    p.title,
    o.quantity
FROM orders o
JOIN customer c
    ON c.id = o.customer_id
JOIN products p
    ON p.id = o.product_id;
```

Получим:

| full_name | title  | quantity |
| --------- | ------ | -------- |
| Ivan      | Laptop | 1        |
| Anna      | Mouse  | 2        |

---

# JOIN и WHERE

Новички часто путают.

### Правильно

```sql
SELECT *
FROM customer c
JOIN orders o
ON o.customer_id = c.id
WHERE c.city = 'Moscow';
```

Сначала:

1. JOIN
2. WHERE

---

# Почему условие связи пишут в ON

Плохо:

```sql
SELECT *
FROM customer c, orders o
WHERE c.id = o.customer_id;
```

Это старый стиль SQL.

Лучше:

```sql
SELECT *
FROM customer c
JOIN orders o
ON c.id = o.customer_id;
```

Так код читается намного легче.

---

# Частая ошибка с LEFT JOIN

Есть запрос:

```sql
SELECT *
FROM customer c
LEFT JOIN orders o
ON o.customer_id = c.id
WHERE o.id IS NOT NULL;
```

Фактически это превращает `LEFT JOIN` в `INNER JOIN`.

Потому что строки с NULL будут отброшены в `WHERE`.

---

# JOIN и производительность

### Всегда создавай индексы для внешних ключей

Например:

```sql
CREATE INDEX idx_orders_customer
ON orders(customer_id);
```

```sql
CREATE INDEX idx_orders_product
ON orders(product_id);
```

Иначе на больших таблицах JOIN станет медленным.

---

# Как PostgreSQL выполняет JOIN

PostgreSQL может использовать несколько алгоритмов:

### Nested Loop Join

Ищет совпадения циклом.

Хорош для маленьких таблиц.

---

### Hash Join

Создает хеш-таблицу.

Очень быстрый вариант для большинства случаев.

---

### Merge Join

Работает на отсортированных данных.

Эффективен на больших выборках.

---

# Самые популярные сценарии

### Все заказы клиента

```sql
SELECT *
FROM customer c
JOIN orders o
ON o.customer_id = c.id;
```

---

### Все клиенты без заказов

```sql
SELECT *
FROM customer c
LEFT JOIN orders o
ON o.customer_id = c.id
WHERE o.id IS NULL;
```

---

### Сумма покупок каждого клиента

```sql
SELECT
    c.full_name,
    SUM(p.price * o.quantity) AS total
FROM customer c
JOIN orders o
    ON o.customer_id = c.id
JOIN products p
    ON p.id = o.product_id
GROUP BY c.id, c.full_name;
```

---

# Что нужно знать для собеседования

Минимальный набор:

* `INNER JOIN` — только совпавшие строки.
* `LEFT JOIN` — все строки слева + совпадения справа.
* `RIGHT JOIN` — все строки справа + совпадения слева.
* `FULL JOIN` — все строки обеих таблиц.
* `CROSS JOIN` — декартово произведение.
* `SELF JOIN` — таблица соединяется сама с собой.
* Разница между `ON` и `WHERE`.
* Как найти записи без связанных данных через `LEFT JOIN ... IS NULL`.
* Почему нужны индексы на поля, участвующие в JOIN (`customer_id`, `product_id`, и т.п.).

На практике около 90% запросов используют именно `INNER JOIN` и `LEFT JOIN`. Остальные виды встречаются значительно реже.
