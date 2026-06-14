# ALTER TABLE и UPDATE в PostgreSQL: подробное руководство

## Введение

При работе с базами данных необходимо понимать разницу между изменением структуры таблицы и изменением данных внутри таблицы.

Для этих целей в SQL используются разные операторы:

* `ALTER TABLE` — изменяет структуру таблицы.
* `UPDATE` — изменяет данные в существующих строках.

Новички часто путают эти команды, потому что обе что-то "изменяют". Однако они работают на разных уровнях базы данных.

---

# UPDATE

## Назначение

Оператор `UPDATE` используется для изменения данных в уже существующих строках таблицы.

Он не меняет структуру таблицы, количество столбцов, ограничения или типы данных.

Он изменяет только значения, которые хранятся в строках.

---

## Синтаксис

```sql
UPDATE table_name
SET column1 = value1,
    column2 = value2
WHERE condition;
```

### Порядок выполнения

1. PostgreSQL находит строки по условию `WHERE`.
2. Для каждой найденной строки изменяет указанные столбцы.
3. Проверяет ограничения (`NOT NULL`, `UNIQUE`, `CHECK`, `FOREIGN KEY`).
4. Если ограничения не нарушены — сохраняет изменения.

---

## Изменение одной строки

Исходные данные:

```sql
id | name
---+-------
1  | Ivan
2  | Petr
```

Запрос:

```sql
UPDATE users
SET name = 'Alex'
WHERE id = 1;
```

Результат:

```sql
id | name
---+-------
1  | Alex
2  | Petr
```

---

## Изменение нескольких полей

```sql
UPDATE users
SET name = 'Alex',
    is_active = false
WHERE id = 1;
```

---

## Изменение всех строк

Если убрать `WHERE`:

```sql
UPDATE users
SET is_active = false;
```

Будут изменены абсолютно все записи.

Это одна из самых распространённых ошибок начинающих разработчиков.

---

## Использование выражений

Можно использовать вычисления.

Пример:

```sql
UPDATE products
SET price = price * 1.1;
```

Все цены увеличатся на 10%.

---

## Использование функций

```sql
UPDATE users
SET name = UPPER(name);
```

Все имена будут переведены в верхний регистр.

---

## UPDATE и связанные таблицы

Рассмотрим таблицы:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE auth (
    user_id INTEGER REFERENCES users(id),
    password VARCHAR(255)
);
```

---

### Изменение данных в дочерней таблице

```sql
UPDATE auth
SET password = 'new_password'
WHERE user_id = 1;
```

Будет изменён пароль пользователя.

---

### UPDATE через JOIN

PostgreSQL позволяет использовать конструкцию `FROM`.

```sql
UPDATE auth a
SET password = 'blocked'
FROM users u
WHERE u.id = a.user_id
AND u.name = 'Ivan';
```

PostgreSQL:

1. Соединяет таблицы.
2. Находит нужные строки.
3. Изменяет записи в таблице `auth`.

---

## RETURNING

Позволяет получить обновлённые данные сразу после изменения.

```sql
UPDATE users
SET name = 'Alex'
WHERE id = 1
RETURNING *;
```

Вернёт обновлённую строку.

---

## Проверка ограничений при UPDATE

Предположим:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE
);
```

Если выполнить:

```sql
UPDATE users
SET email = 'test@mail.com'
WHERE id = 2;
```

И такой email уже существует, PostgreSQL выдаст ошибку.

Проверка ограничений выполняется после обновления строки.

---

# ALTER TABLE

## Назначение

Команда `ALTER TABLE` используется для изменения структуры таблицы.

Она работает с метаданными таблицы:

* столбцами;
* типами данных;
* ограничениями;
* индексами;
* внешними ключами.

---

## Добавление столбца

Исходная таблица:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);
```

Добавим email:

```sql
ALTER TABLE users
ADD COLUMN email VARCHAR(255);
```

Новая структура:

```sql
id
name
email
```

Все существующие строки получат значение NULL в новом столбце.

---

## Добавление столбца со значением по умолчанию

```sql
ALTER TABLE users
ADD COLUMN is_active BOOLEAN DEFAULT true;
```

---

## Удаление столбца

```sql
ALTER TABLE users
DROP COLUMN email;
```

Данные столбца будут полностью удалены.

---

## Переименование столбца

```sql
ALTER TABLE users
RENAME COLUMN name TO full_name;
```

---

## Переименование таблицы

```sql
ALTER TABLE users
RENAME TO customers;
```

---

## Изменение типа данных

Было:

```sql
age INTEGER
```

Станет:

```sql
ALTER TABLE users
ALTER COLUMN age TYPE BIGINT;
```

---

## Добавление NOT NULL

```sql
ALTER TABLE users
ALTER COLUMN name SET NOT NULL;
```

Теперь нельзя вставлять:

```sql
INSERT INTO users(name)
VALUES(NULL);
```

---

## Удаление NOT NULL

```sql
ALTER TABLE users
ALTER COLUMN name DROP NOT NULL;
```

---

## Добавление UNIQUE

```sql
ALTER TABLE users
ADD CONSTRAINT users_email_unique
UNIQUE(email);
```

Теперь значения email должны быть уникальными.

---

## Добавление CHECK

```sql
ALTER TABLE users
ADD CONSTRAINT age_check
CHECK(age >= 18);
```

Теперь PostgreSQL не позволит записать возраст меньше 18.

---

## Добавление внешнего ключа

Таблица:

```sql
CREATE TABLE auth (
    user_id INTEGER,
    password VARCHAR(255)
);
```

Добавим связь:

```sql
ALTER TABLE auth
ADD CONSTRAINT auth_user_fk
FOREIGN KEY(user_id)
REFERENCES users(id);
```

Теперь PostgreSQL будет контролировать целостность данных.

---

# ALTER TABLE и связанные таблицы

Предположим:

```sql
users
```

```sql
id
name
```

и

```sql
auth
```

```sql
user_id
password
```

Связь:

```sql
auth.user_id -> users.id
```

---

## Что происходит при изменении структуры

Если выполнить:

```sql
ALTER TABLE users
DROP COLUMN id;
```

PostgreSQL выдаст ошибку.

Причина:

Таблица `auth` зависит от столбца `users.id`.

Сначала необходимо удалить внешний ключ.

---

## CASCADE

Можно удалить зависимые объекты автоматически.

```sql
ALTER TABLE users
DROP COLUMN id CASCADE;
```

Однако это опасно.

PostgreSQL удалит все зависимости.

---

# Главное различие

## UPDATE

Изменяет данные.

Пример:

```sql
UPDATE users
SET name = 'Alex'
WHERE id = 1;
```

Меняется содержимое строки.

---

## ALTER TABLE

Изменяет структуру таблицы.

Пример:

```sql
ALTER TABLE users
ADD COLUMN email VARCHAR(255);
```

Меняется схема базы данных.

---

# Как запомнить

Если вопрос звучит:

"Я хочу изменить значение"

используется:

```sql
UPDATE
```

Если вопрос звучит:

"Я хочу изменить таблицу"

используется:

```sql
ALTER TABLE
```

Примеры:

Изменить имя пользователя:

```sql
UPDATE users
SET name = 'Alex';
```

Добавить новый столбец:

```sql
ALTER TABLE users
ADD COLUMN email VARCHAR(255);
```

Изменить пароль:

```sql
UPDATE auth
SET password = '12345';
```

Добавить уникальность:

```sql
ALTER TABLE auth
ADD UNIQUE(user_id);
```

Таким образом, UPDATE работает со строками таблицы, а ALTER TABLE работает со структурой самой таблицы.
