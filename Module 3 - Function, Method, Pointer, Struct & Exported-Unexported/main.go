package main

import (
	"fmt"
	"strings"
)

func main() {
	// greet("Adittiya", 25)
	// data("Indonesia", "Male")

	// names := []string{"Adittiya", "Stagnate"}
	// printGreat := greeting("Hello", names)

	// hobbies := []string{"Basketball", "Swimming"}
	// printHobbies := hobby("My hobby", hobbies)

	// game := []string{"Genshin Impact", "Honkai Impact 3rd", "Honkai Star Rails"}
	// printGames := games("I played games", game)

	// fmt.Println(printGreat)
	// fmt.Println(printHobbies)
	// fmt.Println(printGames)

	// diameter := 15
	// var area, circumference = calculate(float64(diameter))
	// fmt.Println("Area", area)
	// fmt.Println("Circumference", circumference)

	studentList := print("Adititya", "Stagnate", "Gracefeel", "Katzcelt", "RhtymKun")

	fmt.Printf("%v", studentList)

	arrayNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}

	fmt.Println(getTotal(arrayNumber...))

	profile("Adittiya Asril", "Dendeng", "Gulai Paruik", "Tambunsu", "Ayam Lado Ijau")

}

// func greet(name string, age int) {
// 	fmt.Println("Hello my name is", name, "Im", age, "years old")
// }

// func data(address, gender string) {
// 	fmt.Println("i'm a", gender, "I lived in", address)
// }

// func greeting(msg string, names []string) string {
// 	joinName := strings.Join(names, " ")

// 	result := fmt.Sprintf("%s %s", msg, joinName)

// 	return result
// }

// func hobby(msg string, hobbies []string) string {
// 	joinHobbies := strings.Join(hobbies, " And ")

// 	result := fmt.Sprintf("%s is %s", msg, joinHobbies)

// 	return result
// }

// func games(msg string, game []string) string {
// 	joinGames := strings.Join(game, ", ")

// 	result := fmt.Sprintln(msg, "like", joinGames)

// 	return result

// }

// func calculate(diameter float64) (float64, float64) {

// 	var area float64 = math.Pi * math.Pow(diameter/2, 2)
// 	var circumference = math.Pi * diameter

// 	return area, circumference
// }

// func calculate(diameter float64) (area float64, circumference float64) {
// 	area = math.Pi * math.Pow(diameter/2, 2)
// 	circumference = math.Pi * diameter

// 	return area, circumference
// }

func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, value := range names {
		key := fmt.Sprintf("Student %d", i+i)
		temp := map[string]string{
			key: value,
		}

		result = append(result, temp)
	}
	return result
}

func getTotal(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func profile(name string, favFood ...string) {
	mergeFavFood := strings.Join(favFood, ", ")

	fmt.Println("Hello My Name Is:", name)
	fmt.Println("My Favorite Food Is:", mergeFavFood)
}
