package goprac
import "fmt"

func Square(num int) {
res := num*num
fmt.Println("Square = ",res)
}

func Cube(num int) {        // lower case name for  func ==> wont exec
res := num*num*num
fmt.Println("cube = ",res)
}

