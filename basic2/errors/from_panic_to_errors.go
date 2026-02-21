package errors

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
)

// не удаляйте, они нужны для проверки
var _ = errors.As
var _ = reflect.Append
var _ = runtime.Gosched

// account представляет счет
type account struct {
    balance   int
    overdraft int
}

func main() {
    fmt.Print("-> ")
    acc, trans, err := parseInput()

    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }
    fmt.Println(acc, trans)
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {
    accSrc, transSrc := readInput()
    acc, aerr := parseAccount(accSrc)
    trans, terr := parseTransactions(transSrc)

    if aerr != nil {
        return acc, []int{}, aerr
    }
    if terr != nil {
        return acc, []int{}, terr
    }

    return acc, trans, nil
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()
    accSrc := scanner.Text()
    var transSrc []string
    for scanner.Scan() {
        transSrc = append(transSrc, scanner.Text())
    }
    return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) (account, error) {
    var acc account
    parts := strings.Split(src, "/")
    balance, berr := strconv.Atoi(parts[0])
    overdraft, oerr := strconv.Atoi(parts[1])
    if berr != nil {
        return acc, berr 
    }
    if oerr != nil {
        return acc, oerr
    }
    if overdraft < 0 {
        return acc, errors.New("expect overdraft >= 0")
    }
    if balance < -overdraft {
        return acc, errors.New("balance cannot exceed overdraft")
    }
    acc.balance = balance
    acc.overdraft = overdraft
    return acc, nil
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
    trans := make([]int, len(src))
    for idx, s := range src {
        t, err := strconv.Atoi(s)
        if err != nil {
            return trans, err
        }
        trans[idx] = t
    }
    return trans, nil
}
