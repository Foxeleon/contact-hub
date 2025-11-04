package main

import (
	"fmt"
	"log"
	"contact-hub/backend/internal/parser"
)

func main() {
    fmt.Println("Запуск сервиса Contact Hub...")

    // Example of calling the parser stub
    persons, err := parser.LoadPersons("./data")
    if err != nil {
    	log.Fatalf("Ошибка при загрузке данных: %v", err)
    }

    log.Printf("Заглушка парсера успешно вызвана. Загружено %d записей.", len(persons))
    // TODO: Add storage initialization and HTTP server startup.
}