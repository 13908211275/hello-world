// pointer is more efficient for large structures
// the implementation mechanism of string, slice and map is similar to pointer,
// which can be passed directly( Only slice length did not change )

func add(n *int) int {
    *n = *n + 1
    return *n
}

func main() {
    x := 1
    y := add(&x)
    
    fmt.Println(x)
    fmt.Println(y)
}
