CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(100),
    age INT,
    phone VARCHAR(20),
    email VARCHAR(100),
    city VARCHAR(50)
);

INSERT INTO students (full_name, age, phone, email, city) VALUES
    ('Александр Иванов', 19, '89001112233', 'alex@mail.ru', 'Москва'),
    ('Мария Петрова', 20, '89002223344', 'maria@yandex.ru', 'Санкт-Петербург'),
    ('Дмитрий Сидоров', 18, '89003334455', 'dima@gmail.com', 'Новосибирск'),
    ('Елена Кузнецова', 22, '89004445566', 'elena@mail.ru', 'Екатеринбург'),
    ('Алексей Смирнов', 20, '89005556677', 'asmir@yandex.ru', 'Москва'),
    ('Ольга Новикова', 21, '89006667788', 'olga@gmail.com', 'Казань'),
    ('Игорь Федоров', 25, '89007778899', 'igor@mail.ru', 'Нижний Новгород'),
    ('Анна Морозова', 17, '89008889900', 'anna@yandex.ru', 'Новосибирск'),
    ('Максим Волков', 23, '89009990011', 'max@gmail.com', 'Самара'),
    ('Наталья Лебедева', 20, '89000001122', 'nat@mail.ru', 'Омск'),
    ('Артем Козлов', 19, '89111112233', 'artem@yandex.ru', 'Челябинск'),
    ('Юлия Павлова', 22, '89222223344', 'yulia@gmail.com', 'Ростов-на-Дону'),
    ('Сергей Степанов', 24, '89333334455', 'serg@mail.ru', 'Уфа'),
    ('Виктория Макарова', 18, '89444445566', 'vikt@yandex.ru', 'Волгоград'),
    ('Константин Орлов', 20, '89555556677', 'kostya@gmail.com', 'Красноярск');

SELECT * FROM students;
-- Задание 0
SELECT full_name, age FROM students;

-- Задание 1
-- Выведите всех студентов старше 20 лет.
SELECT * FROM students WHERE age > 18;

-- Задание 2
--Выведите только имена и города студентов.
SELECT full_name, city  FROM students;

-- Задание 3
--Получите список уникальных городов.
SELECT DISTINCT city FROM students;

-- Задание 4
-- Выведите студентов, возраст которых находится в диапазоне от 18 до 25 лет.
SELECT full_name, age  FROM students WHERE age BETWEEN 18 AND 25;

-- Задание 5
-- Выведите студентов, возраст которых равен 18, 20 или 22 годам.
SELECT full_name, age  FROM students WHERE age IN (18, 20, 22);

-- Задание 6
-- Отсортируйте студентов:
--     * по возрасту по возрастанию;
    SELECT * FROM students ORDER BY age;
--     * по возрасту по убыванию;
    SELECT * FROM students ORDER BY age DESC ;
--     * по имени в алфавитном порядке.
    SELECT * FROM students ORDER BY full_name ;

-- Задание 7
-- Получите первых 5 студентов.
SELECT * FROM students LIMIT 5;

-- Задание 8
-- Получите студентов с 6-й по 10-ю запись.
SELECT * FROM students LIMIT 5 OFFSET 5;

-- Сколько студентов старше 18 лет?
-- Используйте COUNT() и WHERE.
SELECT count(*) FROM students WHERE age > 18;

-- Задание 10
-- Сколько студентов проживает в каждом городе?
SELECT city, count(*) FROM students
GROUP BY city;