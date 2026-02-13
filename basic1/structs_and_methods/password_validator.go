package structs_and_methods

import (
	"unicode"
	"unicode/utf8"
)

// validator проверяет строку на соответствие некоторому условию
// и возвращает результат проверки
type validator func(s string) bool

// digits возвращает true, если s содержит хотя бы одну цифру
// согласно unicode.IsDigit(), иначе false
func Digits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// letters возвращает true, если s содержит хотя бы одну букву
// согласно unicode.IsLetter(), иначе false
func Letters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// minlen возвращает валидатор, который проверяет, что длина
// строки согласно utf8.RuneCountInString() - не меньше указанной
func Minlen(length int) validator {
	return func(s string) bool {
		return utf8.RuneCountInString(s) >= length
	}
}

// and возвращает валидатор, который принимает строку и проверяет,
// что все funcs вернули true для этой строки
func And(funcs ...validator) validator {
	return func(s string) bool {
		for _, f := range funcs {
			if !f(s) {
				return false
			}
		}
		return true
	}
}

// or возвращает валидатор, который принимает строку и проверяет,
// что хотя бы одна из funcs вернула true для этой строки
func Or(funcs ...validator) validator {
	return func(s string) bool {
		for _, f := range funcs {
			if f(s) {
				return true
			}
		}
		return false
	}
}

// password содержит строку со значением пароля и валидатор
type password struct {
	value string
	validator
}

func NewPassword(v string, val validator) *password {
    p := password{value:  v, validator: val}
    return &p
}

// isValid() проверяет, что пароль корректный, согласно
// заданному для пароля валидатору
func (p *password) IsValid() bool {
	return p.validator(p.value)
}
