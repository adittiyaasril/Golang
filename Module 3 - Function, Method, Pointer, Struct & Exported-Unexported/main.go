package main

import (
	"fmt"
	"strings"
)

func main() {
	greet("Adittiya", 25)
	data("Indonesia", "Male")

	names := []string{"Adittiya", "Stagnate"}
	printGreat := greeting("Hello", names)

	hobbies := []string{"Basketball", "Swimming"}
	printHobbies := hobby("My hobby is", hobbies)

	fmt.Println(printGreat)
	fmt.Println(printHobbies)

}

func greet(name string, age int) {
	fmt.Println("Hello my name is", name, "Im", age, "years old")
}

func data(address, gender string) {
	fmt.Println("i'm a", gender, "I lived in", address)
}

func greeting(msg string, names []string) string {
	joinName := strings.Join(names, " ")

	result := fmt.Sprintf("%s %s", msg, joinName)

	return result
}

func hobby(msg string, hobbies []string) string {
	joinHobbies := strings.Join(hobbies, " And ")

	result := fmt.Sprintln(msg, "is", joinHobbies)

	return result
}
