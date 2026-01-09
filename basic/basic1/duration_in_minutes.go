package basic1

import (
	"fmt"
	"time"
)

func DurationInMinutes() {
	// считываем временной отрезок из os.Stdin
	// гарантируется, что значение корректное
	// не меняйте этот блок
	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)

	// выведите исходное значение (s)
	// и количество минут в нем
	// в формате "исходное = X min"
	// используйте метод .Minutes() объекта d
	fmt.Println(s+" =", d.Minutes(), "min")
}
