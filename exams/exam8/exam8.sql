-- SQL ЗАДАЧА №1
CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    name VARCHAR(120) NOT NULL ,
    age INT CHECK ( age >0 AND age <120 )
);

INSERT INTO users(name, age)
VALUES ('Kate', 33),
       ('Bob', 12),
       ('Jake', 17),
       ('Yakob', 18),
       ('Grizly', 10),
       ('China boy', 42),
       ('Akbar', 24),
       ('Geralt', 104),
       ('Stive', 34),
       ('Stella', 16);

INSERT INTO users(name, age)
VALUES ('Drakula', 303); -- ERROR, check( age >0 AND age <120 )

-- Необходимо вывести всех пользователей, возраст которых больше 18 лет.
SELECT * FROM users
WHERE age > 18;


-- SQL ЗАДАЧА №2
CREATE TABLE orders (
    id SERIAL PRIMARY KEY ,
    user_id INT NOT NULL ,
    amount NUMERIC(10, 2) CHECK ( amount>0 ),

    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO orders (user_id, amount)
VALUES (8, 5132),
       (9, 120.99),
       (8, 3),
       (2, 12);

-- Вывести имя пользователя и количество его заказов.
SELECT u.name,
       count(*) AS orders_count
FROM users u
JOIN orders o on u.id = o.user_id
GROUP BY u.id;


-- SQL ЗАДАЧА №3

CREATE TABLE products (
    id SERIAL PRIMARY KEY ,
    title varchar(255) NOT NULL ,
    price NUMERIC(10, 2)  CHECK ( price>0 )
);

INSERT INTO products (title, price)
VALUES ('Lenovo notebook', 10500),
       ('Silver svord', 5132),
       ('Gaming mouse', 1200.99),
       ('Bag', 119.99);

-- Вывести товары, цена которых выше средней цены всех товаров.

SELECT *
FROM products
WHERE price > (SELECT avg(price) FROM products);

-- WITH newt AS (
--     SELECT avg(price) AS avg_price
--     FROM products
-- )
-- SELECT *
-- FROM products p
-- WHERE p.price > newt.avg_price;

-- SQL ЗАДАЧА №4

-- Вывести имя пользователя и сумму всех его заказов.
SELECT u.name,
       coalesce(sum(o.amount), 0) AS orders_sum
FROM users u
LEFT JOIN orders o on u.id = o.user_id
GROUP BY u.id
ORDER BY orders_sum DESC;


-- SQL ЗАДАЧА №5
-- Вывести пользователей, которые не сделали ни одного заказа.

SELECT u.name,
       coalesce(count(o.id), 0)
FROM users u
LEFT JOIN orders o on u.id = o.user_id
GROUP BY u.id
HAVING coalesce(count(o.id), 0) = 0;


SELECT u.name, count(o.amount) AS count
FROM users u
INNER JOIN orders o on u.id = o.user_id
WHERE count < 1
GROUP BY u.name;

SELECT id from users where id not in (
Select u.id from users u
INNER JOIN orders o on u.id = o.user_id);