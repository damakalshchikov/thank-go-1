package basic_constructions

import "fmt"

func LangByCode(code string) {
	var lang string

	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}

	fmt.Println(lang)
}
