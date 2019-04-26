// defer

func readFile() bool {
    f.Open("../example.log")
    defer f.Close()
    
    if fail {
        return false
    }
    return true
}
