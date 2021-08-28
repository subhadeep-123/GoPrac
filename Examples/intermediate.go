package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func star_pattern_loop() {
	// This is program prints the start pattern
	var length int
	fmt.Println("Enter Length of the Pattern")
	fmt.Scanf("%d", &length)

	for i := 0; i <= length; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func infinite_loop() {
	// This function shows the usage of infinite loop
	for i := 0; ; i++ {
		fmt.Println(i)
	}
}

func contd_based_loops() {
	// This function shows the usage of range based for loops
	i := 1
	for i != 1000 {
		fmt.Println(i)
		i += 1
	}
}

func range_based_loops() {
	// This function shows the usage of range in for loops
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for idx, val := range array {
		fmt.Println(idx, val)
	}
}

func example_goto() {
	// This function shows the usage of goto statement
	fmt.Println("This is an example of goto")
}

func guessName() {
	// This function shows the usage of break statement
	defaultName := "Subhadeep"
	var name string
	for i := 0; ; i++ {
		fmt.Println("Enter the name for the Game:- ")
		fmt.Scanln(&name)
		if name == defaultName {
			fmt.Println("Correct.")
			break
		} else {
			fmt.Println("Incorrect.")
		}
	}

}

func reverseEvenFinder() {
	// This function shows the usage of continue statement
	var range_ int
	fmt.Println("Enter the Range:- ")
	fmt.Scanln(&range_)

	for i := 0; i < range_; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}

func reverseStar() {
	// This function shows the usage const
	const LENGTH int = 5
	for i := LENGTH; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("")
	}
}

func showCastingBaby() {
	// This function shows the usage of typecasting
	var number float32 = 103.95959
	fmt.Printf("The float value is - %f\n", number)
	fmt.Printf("TypeCasted value is - %d\n", int(number))
}

func funcDemo(args ...int) (int, int) {
	// This function shows the usage of multiple parameters,
	// anonymous function, and multiple returns
	var sum int
	var mult int = 1
	for _, value := range args {
		func() {
			sum += value
		}()
		if value < 10 {
			func() {
				mult *= value
			}()
		}
	}
	return sum, mult
}

func factorial(num int) int {
	// Shows recursive factorial
	if num == 1 || num == 0 {
		return num
	} else {
		return num * factorial(num-1)
	}
}

func fibonacci(num int) int {
	// Shows recursize fibonacci\
	if num == 1 || num == 0 {
		return num
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func closures() {
	// shows how to make a closure
	number := 10
	squareNum := func() int {
		number *= number
		return number
	}
	fmt.Println(squareNum())
	fmt.Println(squareNum())
}

func print_array_1d(arr []int) {
	// shows how to make an 1d array also,
	// shows how to pass 1d array to a func.
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index - %d Value - %d\n", i, arr[i])
	}
}

func matrix(arr [5][5]int) {
	// shows 2d-arrays
	row, col := 5, 5
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Printf("\n")
	}
}

func printSlice(x []int) {
	fmt.Printf("Slice - %v, Length - %d, Capcity - %d\n", x, len(x), cap(x))
}

func showRegexBaby() {
	re := regexp.MustCompile(".com")
	fmt.Println(re.FindString("google.com"))
	fmt.Println(re.FindString("abc.org"))
	fmt.Println(re.FindString("fb.com"))
	fmt.Println(re.FindStringIndex("google.com"))
	fmt.Println(re.FindStringIndex("abc.org"))
	fmt.Println(re.FindStringIndex("fb.com"))

	re = regexp.MustCompile("f([a-z]+)ing")
	fmt.Println(re.FindStringSubmatch("flying"))
	fmt.Println(re.FindStringSubmatch("abcfloatingxyz"))
}

// Structure things - starts here
type Person struct {
	firstName string
	lastName  string
	age       int
}
type EmpInfo struct {
	Person
	empid int
}

func (person Person) details() {
	fmt.Println("Person Struct")
	fmt.Println(person.firstName)
	fmt.Println(person.lastName)
	fmt.Println(person.age)
}
func (emp EmpInfo) details() {
	fmt.Println("Employee Struct")
	fmt.Println(emp.firstName)
	fmt.Println(emp.lastName)
	fmt.Println(emp.age)
	fmt.Println(emp.empid)
}
func printStruct(p Person) {
	fmt.Printf("Person First Name:- %s\n", p.firstName)
	fmt.Printf("Person Last Name:- %s\n", p.lastName)
	fmt.Printf("Person Age:- %d\n", p.age)
}
func makeStruct() {
	person_1 := Person{"Subhadeep", "Banerjee", 21}
	person_2 := Person{"Iam", "Matrix", 1000}
	person_1.details()
	person_2.details()
	printStruct(person_1)
	printStruct(person_2)
	emp_1 := EmpInfo{Person: Person{"Riju", "SinghaRoy", 25}, empid: 1}
	emp_1.details()
}

// Structure things - ends here

func doPointers(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// Go Maps Starts here
func goMaps() {
	x := map[string]string{
		"Name":     "Subhadeep Banerjee",
		"Email Id": "subhadeeop762@gmail.com",
		"Phone No": "7980207055"}
	fmt.Printf("Map - %v \nType - %T\n", x, x)

	// Inserting and Updating the values
	x["Address"] = "86/11 k.k Roychowdhury Road"
	x["Phone No"] = "+91 7980207055"
	fmt.Println(x)

	// deleting values in the map
	delete(x, "Phone No")
	fmt.Println(x)

	// retriving values from the map
	// the ok will return true if the value is present false otherwise.
	val, ok := x["Address"]
	fmt.Println(val, ok)
	val, ok = x["Sex"]
	fmt.Println(val, ok)
}

type userInformation struct {
	Name  string
	Email string
	Age   float32
}

func goStructMaps() {
	var structMap = map[string]userInformation{
		"User 1": {
			"Subhadeep Banerjee",
			"subhadeep762@gmail.com",
			22.1},
		"User 2": {
			"Soumya Mitra",
			"mitrasoumya789@gmail.com",
			21.11}}
	fmt.Println(structMap)
}

// Go Maps Ends Here

// Defer - Panic - Recover Starts Here
// Error
func squareRoot(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Value cannot be less then 0")
	} else {
		return math.Sqrt(value), nil
	}
}

// Recover
func divideNumbers(value_1, value_2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	return value_1 / value_2
}

// Defer - Panis - Recover Ends Here

// Case 28 - Block
func executePanic() {
	defer func() {
		if recoveryMessage := recover(); recoveryMessage != nil {
			fmt.Println(recoveryMessage)
		}
		fmt.Println("This is recovery function...")
	}()
	panic("This is a very Panic Situation")
	fmt.Println("The Function executes completly")
}

// Case 28 - Block End
func main() {
	/*
		This is the main block of the programs
		and it uses switch case for menu based
		design.
	*/
	var input int

	cla, _ := strconv.Atoi(os.Args[1])
	if cla == 0 {
		input, _ = strconv.Atoi(os.Args[2])
	} else {
		fmt.Printf(`
		1. Star Pattern
		2. Infinite Loop
		3. Condition Based Loops
		4. Range Based Loops
		5. Goto Demo
		6. Guessing Game
		7. Reverse Even Finder
		8. Reverse Star Pattern
		9. Type Casting Baby
		10. Function Demo.
		11. Recursion
		12. Colosure
		13. 1D Array
		14. 2D Array
		15. Slice
		16. Command Line Arguments
		17. Strings
		18. Regex
		19. Structures
		20. Pointers
		21. Reflect
		22. Rune
		23. Maps
		24. Error
		25. Recover
		26. Defer
		27. Panic
		28. Defer - Recover - Panic
		`)
		fmt.Println("\nEnter Your Choice.")
		fmt.Scanln(&input)
	}
Demo:
	example_goto()

	switch input {
	case 1:
		star_pattern_loop()
	case 2:
		infinite_loop()
	case 3:
		contd_based_loops()
	case 4:
		range_based_loops()
	case 5:
		goto Demo
	case 6:
		guessName()
	case 7:
		reverseEvenFinder()
	case 8:
		reverseStar()
	case 9:
		showCastingBaby()
	case 10:
		sum, mult := funcDemo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
		fmt.Printf("The Sum of first 15 Natural no :- %d\n", sum)
		fmt.Printf("The Multiplication of first 10 Natural no :- %d\n", mult)
	case 11:
		var choice int
		fmt.Printf("1. Factorial\n2. Fibonacci\n")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var num int
			fmt.Printf("Enter the value for Factorial:- ")
			fmt.Scanln(&num)
			fmt.Printf("The Factorial of %d is %d", num, factorial(num))
		case 2:
			var num int
			fmt.Printf("Enter the value for Fibonacci:- ")
			fmt.Scanln(&num)
			fmt.Printf("The Fibonacci of %d is %d", num, fibonacci(num))
		}
	case 12:
		closures()
	case 13:
		var arr [10]int
		fmt.Printf("The length of the array - %d", len(arr))
		for i := 0; i < len(arr); i++ {
			arr[i] = i * i
		}
		print_array_1d(arr[:])
	case 14:
		var arr [5][5]int
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				arr[i][j] = j
			}
		}
		matrix(arr)
	case 15:
		nilSliceCheck := func(slice []int) {
			if slice == nil {
				fmt.Println("This is a new slice.")
			} else {
				fmt.Println(slice)
			}
		}

		// nil slice
		newSlice := make([]int, 10, 10)
		fmt.Printf("Length - %d, Capacity - %d\n", len(newSlice), cap(newSlice))
		nilSliceCheck(newSlice)

		// sub slicing
		tempSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var tempNewSlice []int
		fmt.Println(tempSlice[5:])
		fmt.Println(tempSlice[:5])
		fmt.Println(tempSlice[3:5])

		// appending to the nil slice
		i := 1
		for i != 6 {
			num := append(tempNewSlice, i)
			printSlice(num)
			i += 1
		}
		nilSliceCheck(tempNewSlice)
	case 16:
		fmt.Printf("The command line arguments :- \n")
		for _, args := range os.Args[:] {

			fmt.Println(args)
		}
	case 17:
		msg := "Hello World"
		fmt.Println(msg)
		fmt.Println(reflect.TypeOf(msg))
		fmt.Printf("The length of the string - %d\n", len(msg))
		fmt.Println("Ascii value of A is ", "A"[0])
		fmt.Printf("String to upper - %s\n", strings.ToUpper(msg))
		fmt.Printf("String to Lower - %s\n", strings.ToLower(msg))

		var arr = []string{"a", "b", "c", "d"}
		joinedString := strings.Join(arr, "*")
		fmt.Println(joinedString)
		fmt.Println(strings.Split(joinedString, "*"))
		fmt.Println(strings.ReplaceAll(joinedString, "*", ""))
		fmt.Println(strings.Count(joinedString, "*"))
		fmt.Println(strings.Contains(joinedString, "*"))
		fmt.Println(strings.Compare("a", "b"))
		fmt.Println(strings.Compare("a", "a"))
		fmt.Println(strings.Compare("b", "a"))
		fmt.Println(strings.ContainsAny(joinedString, "*"))
	case 18:
		showRegexBaby()
	case 19:
		makeStruct()
	case 20:
		var a, b int
		a = 20
		b = 30
		fmt.Printf("The Value of A - %d, B - %d\n", a, b)
		doPointers(&a, &b)
		fmt.Printf("The Value of A - %d, B - %d", a, b)
	case 21:
		age := 26.59889
		fmt.Printf("%T\n", age)
	case 22:
		charsToCheck := []rune{'♛', '♠', '♧', '♡', '♬'}
		for i, j := range charsToCheck {
			fmt.Printf("\nCharacter: %c, Unicode: %U, Position: %d ", j, j, i)
		}

		chars := "的是不 - ঊঋ"
		fmt.Println(chars)

		var s string = "This is a string"
		fmt.Printf("Value of = %v, Type = %T\n", s, s)
		fmt.Printf("Value of = %v, Type = %T\n", s[2], s[5])
	case 23:
		fmt.Printf(`
		1. Normal Go Maps
		2. Structure Based Go Maps
		`)
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			goMaps()
		case 2:
			goStructMaps()
		}
	case 24:
		defulatValue := -81.00
		result, err := squareRoot(defulatValue)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("The Square root of - %f - %f\n", defulatValue, result)
		}
	case 25:
		one := divideNumbers(100, 11)
		fmt.Println(one)
		two := divideNumbers(50, 0)
		fmt.Println(two)
	case 26:
		defer fmt.Println("This is the usage of defer - 1")
		fmt.Printf("Hey This block is for testing defer.\n")
	case 27:
		panic("Something is not right")
	case 28:
		executePanic()
		fmt.Println("Control returned ot this block")
	default:
		fmt.Println("Enter a valid option.")
	}

}
