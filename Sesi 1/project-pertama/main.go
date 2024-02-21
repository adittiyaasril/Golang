package main

import "fmt"

func main() {
		//Println for new line print
    // fmt.Println("Hello, World!")
		// fmt.Println("Adittiya", "Jordan")

		//Print for same line print
		// fmt.Print("Hello")
		// fmt.Print("World")

		// var name string = "Adittiya"
		// var age int = 2000

		
		//reassign value 
		// var name string
		// var age int
		
		// name = "Adittiya"
		// age = 2000
		
		
		//variable without data declaration
		// var name = "Adittiya"
		// var age = 2000
		
		// fmt.Printf("%T, %T", name, age)
		
		//short declaration variables but u need to give a value 
		// name := "adittiya"
		// age := 2000
		
		// fmt.Println("This is my name ==> ", name)
		// fmt.Println("This is my age ==> ", age)

		//multiple declaration variable
		// var student1, student2, student3 string = "One", "Two", "Three"
		// var first, second, third int 

		// first, second, third = 1, 2, 3

		// fmt.Println(student1,student2,student3)
		// fmt.Println(first,second,third)

		//underscore variable to prevent errors when variables are not used
		// var firstVariable string
		// var name, age, address = "Adittiya", 2000, "West Sumatra"

		// _, _, _, _ = firstVariable, name, age, address

		//fmt.Printf usage
		// first, second := "1", 2

		// fmt.Printf("the data type of the first variable is %T \n", first)
		// fmt.Printf("the data type of the second variable is %T \n", second)

		var name = "Adittiya"
		var age = 2000
		var address = "West Sumatra"

		fmt.Printf("Hello my name is %s, I'm %d years old, i lived in %s", name, age, address)
}