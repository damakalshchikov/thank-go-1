package main

import "github.com/damakalshchikov/thank-go-1/basic/arrays_and_maps"

// import "github.com/damakalshchikov/thank-go-1/basic/basic_constructions"

func main() {
	// Задачи из раздела 1.2 "Базовые конструкции"
	// basic_constructions.DurationInMinutes()
	// basic_constructions.DistanceBetweenPoints()
	// basic_constructions.RepeatString("a", 5)
	// basic_constructions.LangByCode("en")

	// Задачи из раздела 1.3 "Массивы и карты"
	arrays_and_maps.ToShortString("12345", 6)
	arrays_and_maps.ToShortString("12345", 5)
	arrays_and_maps.ToShortString("12345", 4)
	arrays_and_maps.ToShortString("12345", 1)
}
