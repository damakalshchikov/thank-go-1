package interfaces

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// element - интерфейс элемента последовательности
type Element interface{}

// weightFunc - функция, которая возвращает вес элемента
type weightFunc func(Element) int

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
    // чтобы понять сигнатуры методов - посмотрите,
    // как они используются в функции max() ниже
    next() bool
    val()  Element
}

// intIterator - итератор по целым числам
// (реализует интерфейс iterator)
type intIterator struct {
    // поля структуры
    index int
    src   []int
}

// методы intIterator, которые реализуют интерфейс iterator
func (i *intIterator) next() bool {
    if indx := i.index + 1; indx < len(i.src) {
        i.index = indx
        return true
    }
    return false
}

func (i *intIterator) val() Element {
    return i.src[i.index]
}

// конструктор intIterator
func NewIntIterator(src []int) *intIterator {
    return &intIterator{index: -1, src: src}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
// func main() {
//     nums := readInput()
//     it := newIntIterator(nums)
//     weight := func(el element) int {
//         return el.(int)
//     }
//     m := max(it, weight)
//     fmt.Println(m)
// }

// max возвращает максимальный элемент в последовательности.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func Max(it iterator, weight weightFunc) Element {
    var maxEl Element
    for it.next() {
        curr := it.val()
        if maxEl == nil || weight(curr) > weight(maxEl) {
            maxEl = curr
        }
    }
    return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func ReadInput() []int {
    var nums []int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        nums = append(nums, num)
    }
    return nums
}
