package main

import (
	"fmt"
	"workWithCache/config"
	"workWithCache/server"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	server.Run(cfg)
}

// слой сервера
// слой бизнес-логики (мы его скипнули, но можешь добавить)
// слой базы данных (хранение)

// Сервер: chi
// - ручка GET /info?numbers=[1,2,3,4] -> [3,6,8,12] (максимум 10)
// - возвращает умноженное число

// Логика ручки:
// - поиск в мапе (кэш) значения
// - отправлю запрос к БД, ответ положу в мапу
// - map[int]int ; map[10] = 30 -> LinkedList LRU cache
// - кладёт в кеш (read about it)

// Слой "Базы данных":
// - "ищет" число, умножая его на три
// - при "поиске" ждет 0,2 сек

// Слой "Кеширования":
// - максимум N чисел, задаётся в конфигурации сервиса (var in main)
// - ожидания нет

// Этап 1.5 переписать на линкед лист

// Этап 2:
// - добавить ручку изменения размера кеша (указывается просто новый размер)
// POST /setCache, тело сам придумай + валидация (-500 handle exceptions (0 is ok))

// Этап 3:
// - каждые 5 секунд работы сервиса множитель увеличивается на 1
// - при увеличении множителя что делать с кэшем? (update it)

// Этап 4:
// - сделать простой! тестовый клиент для этого сервера, измеряя время запроса

// Этап 5 после проверки:
// Контрольные вопросы:
// - Что такое REST?
// - Что такое кэш и какие проблемы он решает?
// - Что такое контекст в разрезе сетевого взаимодействия, что с ним можно делать?
// - Long-live connection, TCP handshake
