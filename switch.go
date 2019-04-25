// switch

n := 3
switch n {
case 1:
    fmt.Println("first") // default break this

case 2,3:
    fmt.Println("second, third") 
    fallthrough // change default behavior
case 4:
    fmt.Println("fourth") 
default:
     fmt.Println("unknown")
}
