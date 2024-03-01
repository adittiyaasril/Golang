package main

import "fmt"

func main() {

	// CONDITION
	// If Else
	// currentYear := 2024

	// if age := currentYear - 1998; age < 17 {
	// 	fmt.Println("Kamu Tidak Boleh Membuat SIM")
	// 	} else{
	// 	fmt.Println("Kamu Boleh Membuat SIM")
	// }

	// switch case
	// score:= 8

	// switch score {
	// case 8:
	// 	fmt.Println("Perfect")
	// case 7:
	// 	fmt.Println("Awesome")
	// case 6:
	// 	fmt.Println("Not Bad")
	// default:
	// // 	fmt.Println("Bad")
	// }

	// switch case fallthrough
	// score:=6

	// switch {
	// case score == 8:
	// 	fmt.Println("Perfect")
	// case (score <= 8) && (score > 3):
	// 	fmt.Println("Not Bad")
	// 	fallthrough
	// case score >=5:
	// 	fmt.Println("Its Okay But Please Study Harder")
	// default:{
	// 	fmt.Println("Study Harder")
	// 	fmt.Println("You don't have a good score yet")
	// }
	// }

	// nested condition
	// score := 10

	// if score > 7 {
	// 	switch score {
	// 	case 10:
	// 		fmt.Println("Perfect")
	// 	default:
	// 		fmt.Println("Nice")
	// 	}
	// } else{
	// 	if score == 5 {
	// 		fmt.Println("Not Bad")
	// 	} else if score == 3 {
	// 		fmt.Println("Keep Trying")
	// 	} else{
	// 		fmt.Println("You Can Do It")
	// 		if score == 0 {
	// 			fmt.Println("Try Harder")
	// 		}
	// 	}
	// }

	// -----------------------------------------------------------------------------------------------------------
	// LOOPING
	// for i := 0; i <3; i++{
	// 	fmt.Println("Number", i)
	// }

	// i := 0
	// for i < 3 {
	// 	fmt.Println("Number", i)
	// 	i++
	// }

	// i := 0
	// for{
	// 	fmt.Println("Number", i)

	// 	i++
	// 	if i == 3{
	// 		break
	// 	}
	// }

	// for i := 0; i < 20; i++{
	// 	if(i%2 == 1){
	// 		continue
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println("Number", i)
	// }

	// for i := 0; i < 10; i++{
	// 	for j := i; j < 10; j++{
	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 9; i >= 0; i--{
	// 	for j := i; j < 10; j++{
	// 		fmt.Print(j," ")
	// 	}
	// 	fmt.Println()
	// }

	// outerLoop:
	// for i:= 0; i<5; i++ {
	// 	fmt.Println("Perulangan Ke", i)
	// 		for j := 0; j<5; j++{
	// 			if i == 4 {
	// 				break outerLoop
	// 			}
	// 			fmt.Print(j, " ")
	// 		}
	// 		fmt.Println()

	// }

	//--------------------------------------------------------------------------------------------------------------
	//ARRAY

	// numbers := [4]int{1, 2, 3, 4}

	// var string = [4]string{"Adittiya", "Stagnate", "Gracefeel", "Ramiris"}

	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", string)

	// for i := 0; i < len(string); i++{
	// 	fmt.Println(string[i])
	// }

	// fruits := [4]string{"Mango","Apple", "Banana","Orange"}

	// fruits[1] = "Coconut"

	// fmt.Println(fruits)
	// fmt.Printf("%#v\n", fruits)

	// for index, value := range fruits{
	// 	fmt.Println("Index ke", index, "=", value)
	// }

	// matrix := [2][3]int{{1,2,3}, {7,8,9}}

	//SLICE

	// var fruits = []string{"banana", "apple", "dragon fruit"}

	// _ = fruits

	// fruits := make([]string, 3)

	// fruits[0] = "Banana"
	// fruits[1] = "Mango"
	// fruits[2] = "Apple"

	// fruits = append(fruits, "Banana", "Mango", "Dragon Fruits")

	// fmt.Printf("%#v", fruits)

	// fruits1 := []string{"Mango", "Banana", "Lemon"}
	// fruits2 := []string{"Dragon Fruits", "Watermelon", "Coconut"}

	// fruits1 = append(fruits1, fruits2...)

	// fmt.Printf("%#v", fruits1)

	// number1 := []int{1, 2, 3}
	// number2 := []int{3, 4, 5}

	// nn := copy(number1, number2)

	// fmt.Println(number1, number2, nn)

	// fruits := []string{"Mango", "Banana", "Pineapple", "Lemon", "Watermelon", "Dragon Fruits", "Durian", "Starfruit"}

	// fruits1 := fruits[1:4]
	// fmt.Println(fruits1)

	// fruits2 := fruits[:2]
	// fmt.Println(fruits2)

	// fruits3 := fruits[3:]
	// fmt.Println(fruits3)

	// fruits4 := fruits[:]
	// fmt.Println(fruits4)

	// fruits := []string{"Mango", "Banana", "Pineapple", "Lemon", "Melon"}

	// fruits = append(fruits[:2], "Rambutan")

	// fmt.Printf("%v", fruits)

	// fruits := []string{"Mango", "Banana", "Pineapple", "Lemon", "Melon"}

	// fmt.Println("Fruit cap => ", cap(fruits))
	// fmt.Println("Fruit len => ", len(fruits))

	// fmt.Println(strings.Repeat("-", 100))

	// fruit1 := fruits[:4]

	// fmt.Println("Fruit cap => ", cap(fruit1))
	// fmt.Println("Fruit len => ", len(fruit1))

	// fmt.Println(strings.Repeat("-", 100))

	// fruit2 := fruits[1:]

	// fmt.Println("Fruit cap => ", cap(fruit2))
	// fmt.Println("Fruit len => ", len(fruit2))

	// fmt.Println(strings.Repeat("-", 100))

	// cars := []string{"Honda", "Ford", "Audi", "Range Rover"}
	// newCars := []string{}
	// newCars = append(newCars, cars...)

	// cars[0] = "Nissan"

	// for i := 0; i < len(cars); i++ {
	// 	fmt.Println("Cars Ke", i+1, cars[i])
	// }

	// for index, value := range newCars {
	// 	fmt.Println("New Cars ke ", index+1, " ", value)
	// }

	//MAP

	// person := map[string]string{}

	// person["name"] = "Adittiya"
	// person["age"] = "23"
	// person["address"] = "West Sumatra"

	// person := map[string]string{
	// 	"name":    "Adittiya",
	// 	"age":     "23",
	// 	"address": "West Sumatra",
	// }

	// fmt.Println("My name Is :", person["name"])
	// fmt.Println("Im ", person["age"], "Years old")
	// fmt.Println("I live in", person["address"])

	// for key, value := range person {
	// 	fmt.Println(key, ":", value)
	// }

	// fmt.Println("Before Deleting :", person)

	// delete(person, "age")

	// fmt.Println("After Deleting :", person)

	// value, exist := person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value cant be find")
	// }

	// delete(person, "age")

	// value, exist = person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value cant be find")
	// }

	employees := []map[string]string{
		{"name": "Adittiya", "age": "25", "address": "Indonesia"},
		{"name": "John", "age": "23", "address": "Singapore"},
		{"name": "Arnold", "age": "45", "address": "Malaysia"},
	}

	for i, employee := range employees {
		fmt.Printf("index: %v, name: %v, age: %v, address: %v\n", i, employee["name"], employee["age"], employee["address"])
	}
}
