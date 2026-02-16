package main

import (
	"fmt"
	"github.com/damakalshchikov/thank-go-1/basic2/interfaces"
)

// "github.com/damakalshchikov/thank-go-1/basic/arrays_and_maps"
// "github.com/damakalshchikov/thank-go-1/basic/basic_constructions"
// "github.com/damakalshchikov/thank-go-1/basic/funcs_and_points"
// "github.com/damakalshchikov/thank-go-1/basic/structs_and_methods"

func main() {
	// Задачи из раздела 1.2 "Базовые конструкции"
	// basic_constructions.DurationInMinutes()
	// basic_constructions.DistanceBetweenPoints()
	// basic_constructions.RepeatString("a", 5)
	// basic_constructions.LangByCode("en")

	// Задачи из раздела 1.3 "Массивы и карты"
	// arrays_and_maps.ToShortString("12345", 1)
	// arrays_and_maps.NumCounter(123)
	// arrays_and_maps.MakeAbbreviation()

	// Задачи из раздела 1.4 "Функции и указатели"
	// src := funcs_and_points.ReadInput()
	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`
	// res := funcs_and_points.Filter(func(x int) bool { return x%2 == 0 }, src)
	// fmt.Println(res)

	// a, b := 1.0, 3.0
	// funcs_and_points.Normalize(&a, &b)
	// fmt.Println(a, b)

	// a, b, c, d := 1.0, 2.0, 3.0, 4.0
	// funcs_and_points.Normalize(&a, &b, &c, &d)
	// fmt.Println(a, b, c, d)

	// Задачи из раздела 1.5 "Структуры и методы"
	// src := structs_and_methods.ReadString()
	// trn := structs_and_methods.ParseTournament(src)
	// rt := trn.CalcRating()
	// rt.Print()

	// var s string
	// fmt.Scan(&s)
	// // валидатор, который проверяет, что пароль содержит буквы и цифры,
	// // либо его длина не менее 10 символов
	// validator := structs_and_methods.Or(structs_and_methods.And(structs_and_methods.Digits, structs_and_methods.Letters, structs_and_methods.Minlen(10)))
	// p := structs_and_methods.NewPassword(s, validator)
	// fmt.Println(p.IsValid())

	// Задачи из раздела 2.1 "Интерфейсы"
	// "Универсальный итератор"

	nums := interfaces.ReadInput()
	it := interfaces.NewIntIterator(nums)
	weight := func(el interfaces.Element) int {
		return el.(int)
	}
	m := interfaces.Max(it, weight)
	fmt.Println(m)
}
