// functions are passed as values

func isEvenNumber(n int) bool {
    if n % 2 == 0 {
        return true
    }
    return false
}

// declare a funcion type
type evenNumber func(int) bool

func filter(s []int, f evenNumber) []int {
    var res []int
    for _, v := range s {
        if f(v) {    // true
            res = append(res, v)
        }
    }
    return res
}

func main() {
    slice := []int {1, 2, 3, 4, 5, 6}
    ev := filter(slice, isEvenNumber)
    fmt.Println(ev)
}
