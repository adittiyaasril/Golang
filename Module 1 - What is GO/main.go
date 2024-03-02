package main

import "fmt"

func main() {
//------------------------------------------------------------------------------------------------------------------------------
	//VARIABLES
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

		// var name = "Adittiya"
		// var age = 2000
		// var address = "West Sumatra"

		// fmt.Printf("Hello my name is %s, I'm %d years old, i lived in %s", name, age, address)

//-----------------------------------------------------------------------------------------------------------------------------------------
		//DATA TYPES

		//Data types in golang
		//1.Basic Type: number, string, boolean.
		//2.Aggregate Type: array dan struct.
		//3.Reference Type: slice, pointer, map, function, channel
		//4.Interface Type: interface.

		//number integer
		// var first uint8 = 10
		// var second int8 = -23

		// fmt.Printf("First data types is : %T \n", first)
		// fmt.Printf("Second data types is : %T \n", second)

		//number float
		// var decimalNumber float32 = 3.63

		// fmt.Printf("Decimal Number: %f\n", decimalNumber)
		// fmt.Printf("Decimal Number: %.3f\n", decimalNumber)

		//Booleans
		// var condition bool = true
		// fmt.Printf("is permitted? %t \n", condition)

		//String
		// var message string = `Welcome to the "Golang" course!!!!`

		// fmt.Println(message)

//--------------------------------------------------------------------------------------------------------------------
		//CONSTANT AND OPERATOR
		
		//Constant
		// const fullName string = "Adittiya Asril"

		// fmt.Printf("Hello %s", fullName)

		//Operator
		//Arithmetic Operator
		// var value int = 2 + 2 * 3
		// fmt.Println(value)

		//Relational Operator
		// var firstCondition bool = 2 > 3
		// var secondCondition bool = "joey" == "Joey"
		// var thirdCondition bool = 10 != 2.3
		// var fourthCondition bool = 11 <= 11

		// fmt.Println("first condition", firstCondition)
		// fmt.Println("second condition", secondCondition)
		// fmt.Println("third condition", thirdCondition)
		// fmt.Println("forth condition", fourthCondition)

		//Logical Operator
		var right bool = true
		var wrong bool = false

		var rightAndWrong =(right && wrong)
		fmt.Printf("right and wrong \t(%t) \n", rightAndWrong)

		var rightOrWrong =(right || wrong)
		fmt.Printf("right or wrong \t(%t) \n", rightOrWrong)

		var wrongReverse =!wrong
		fmt.Printf("right and wrong \t(%t) \n", wrongReverse)

}