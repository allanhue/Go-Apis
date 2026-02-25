package main

import (
    "fmt"
    "os"
)
// type Page struct {
//     Title string
//     Body  []byte
// }

// func (p *Page) save() error {
//     filename := p.Title + ".txt"
//     return os.WriteFile(filename, p.Body, 0600)
// }
// func main() {
// 	fmt.Println("Hello, World")
// 	fmt.Println("This is a simple Go program.")
// }


// package main
// import ("fmt")

// func main() {
//   var arr1 = [3]int{1,2,3}
//   arr2 := [5]int{4,5,6,7,8}
//  arr := [...]int{1,2,3,4,5}    // the best soultion for array declaration infreerign the length 

//   fmt.Println(arr1)
//   fmt.Println(arr2)
// }

// package main
// import ("fmt")

func main() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)
}