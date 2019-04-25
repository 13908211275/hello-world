// process control

// if
if x := fetchInt(); x > 3 {
    fmt.Println("x is greater than 3")
}

// goto
// careful use
func demo1() {
    n := 0
Bgn:
    i++
    fmt.Println(i)
    goto Bgn
}
