
Задача 1 — Channel message
Условие
Создать channel типа string.
Требования
Запустить goroutine
Внутри goroutine отправить:
"hello from goroutine"
В main:
получить сообщение
вывести в консоль
Использовать
goroutine
channel
———————————————————————-———————————————————————-
Задача 2 — Sum через channel
(твоя подходящая задача)
Условие
nums := []int{1,2,3,4,5}
Требования
Запустить goroutine
Посчитать сумму
Отправить результат в channel
Получить и вывести результат
———————————————————————-———————————————————————-
Задача 3 — Pipeline строк
Условие
Есть slice:
words := []string{"go", "lang", "channel"}
Требования
Создать goroutine
Передавать слова через channel
В другой goroutine:
получать слова
переводить в upper case
отправлять в другой channel
В main вывести результат
Ожидаемый вывод
GO
LANG
CHANNEL
Использовать
2 channel
2 goroutine
