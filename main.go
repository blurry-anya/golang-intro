package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var basicScore = 5

func main() {

	// greet("Anya")

	// types()

	// printAndFormat()

	// arraysAndSlices()

	// libraries()

	// loops()

	// boolAndConditions()

	// functions()

	// sayHello("Anya") // run with greetings.go

	// mapsDemo() // run with maps.go

	// passByValue()

	pointers()
}

func greet(name string) {
	fmt.Println("Yo,", name)
}

func types() {
	// strings
	var nameOne string = "Name One"
	nameTwo := "Name Two"

	fmt.Println(nameOne)
	fmt.Println(nameTwo)

	// integers
	var ageOne int = 20
	var ageTwo = 35
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var num1 int8 = 12
	var num2 uint8 = 255

	fmt.Println(num1, num2)

	// floats
	var scoreOne float32 = 3.2
	var scoreTwo float64 = 888888888328203402340234.3
	scoreThree := 444.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}

func printAndFormat() {
	// print
	fmt.Print("Printing ")
	fmt.Print("words")

	// print line
	fmt.Println("Sometimes I'm alone")
	fmt.Println("Sometimes I'm not")
	fmt.Println("Sometimes I'm alone")
	fmt.Println("Hello?")

	// formatted strings
	name := "Anya"
	age := 23
	score := 23.4
	// %_ - format specifier
	// %v - default one
	// %q - vars in quotes
	// %T - a typeof alternative
	// %0.Nf - float with N digits after the point
	fmt.Printf("Yo, my name is %q and I'm %v \n", name, age)
	fmt.Printf("Age variable is of type: %T \n", age)
	fmt.Printf("Score is: %0.2f !\n", score)

	// saving formatted strings into variables
	var greeting = fmt.Sprintf("Yo, my name is %q and I'm %v \n", name, age)
	fmt.Println("Saved string:", greeting)
}

func arraysAndSlices() {

	// arrays (always a fixed length)
	var numbers = [3]int{1, 2, 3} // OR var numbers [3]int = [3]int{1, 2, 3}

	names := [2]string{"Anya", "Loh"}
	names[1] = "Chmo"

	fmt.Println("Numbers:", numbers, "\nLength:", len(numbers))
	fmt.Println("Names:", names, "\nLength:", len(names))

	// slices (dynamic, arrays under the hood)
	var scores = []int{10, 20, 30}
	fmt.Println("Scores (slice)", scores)

	scores = append(scores, 40)
	fmt.Println("Scores (slice)", scores)

	// slice ranges
	rangeOne := scores[0:2]
	rangeTwo := scores[2:]
	rangeThree := scores[:3]

	fmt.Println("Range 1:", rangeOne)
	fmt.Println("Range 2:", rangeTwo)
	fmt.Println("Range 3:", rangeThree)

	rangeOne = append(rangeOne, 25)
	fmt.Println("Range 1:", rangeOne)

}

func libraries() {

	// strings
	fmt.Println("\tSTRINGS")
	greeting := "Hello, how are you?"
	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hey"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "you"))
	fmt.Println(strings.Split(greeting, " "), len(strings.Split(greeting, " ")))

	// sorts
	fmt.Println("\tSORT")
	ages := []int{5, 9, 1, 4, 5, 6, 1, 8, 9}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 9)

	fmt.Println("index:", index)

	names := []string{"anya", "abigail", "sebastian", "robin"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "robin"))
}

func loops() {
	i := 0
	for i < 5 {
		fmt.Println("i:", i)
		i++
	}

	fmt.Println()

	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	fmt.Println()

	names := []string{"Anya", "Abigail", "Robin", "Sebastian"}
	for i := 0; i < len(names); i++ {
		fmt.Printf("%v: %v\n", i, names[i])
	}

	fmt.Println()

	for index, value := range names {
		fmt.Printf("%v: %v\n", index, value)
	}

	fmt.Println()

	for _, value := range names {
		fmt.Printf("%v\n", value)
	}
}

func boolAndConditions() {
	age := 41

	fmt.Println(age >= 50)
	fmt.Println(age != 50)
	fmt.Println(age <= 50)
	fmt.Println(age == 23)

	if age < 30 {
		fmt.Println("Hooray, you are still young!")
	} else if age < 40 {
		fmt.Println("You are not THAT old.")
	} else {
		fmt.Println("Oh, you are old..")
	}

	fmt.Println()

	names := []string{"Anya", "Abigail", "Robin", "Sebastian"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("Continue at index", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at index", index)
			break
		}

		fmt.Printf("%v: %v \n", index, value)
	}

}

func functions() {
	names := []string{"Anya", "Abigail", "Robin", "Sebastian"}
	cycleStrings(names, greet)

	fmt.Printf("Area is: %0.3f\n", circleArea(4))

	firstName, secondName := getInitials("anna kononchenko")
	fmt.Println(firstName, secondName)
}

func cycleStrings(strings []string, function func(string)) {
	for _, value := range strings {
		function(value)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func getInitials(name string) (string, string) {
	name = strings.ToUpper(name)
	names := strings.Split(name, " ")

	var initials []string

	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}
