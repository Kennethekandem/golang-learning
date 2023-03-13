package main


import ("fmt")


func main() {

	var a int = 12
	// c := "test"

	var (
		c string = "test"
		name string = "Kenneth"
		age int = 24
		get_false bool = true
	)

	fmt.Println("Hello World")
	fmt.Println(a)
	fmt.Println(c)
	fmt.Printf("%T, %v", c, c)
	fmt.Println(c, name, age, get_false)
}