package main

import (
	"Task3/Triangle"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputDetails() []Triangle.Triangle {
	var triangles []Triangle.Triangle
	var input_next string

	for ok := true; ok; ok = (input_next == "y" || input_next == "yes") {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter triangle details: ")
		input, _ := reader.ReadString('\n')
		sl := strings.Split(strings.TrimSpace(input), ",")

		validInput,err:= Triangle.ValidateInput(sl)
		for !validInput{
			fmt.Print("Enter triangle details: ")
			input, _ := reader.ReadString('\n')
			sl = strings.Split(strings.TrimSpace(input), ",")
			validInput,err= Triangle.ValidateInput(sl)
			fmt.Println(err)
		}

		name := sl[0]
		side1, _ := strconv.ParseFloat(sl[1], 64)
		side2, _ := strconv.ParseFloat(sl[2], 64)
		side3, _ := strconv.ParseFloat(sl[3], 64)
		area := 0.0
		area = Triangle.GetTriangleArea(Triangle.Triangle{name, side1, side2, side3,area})
		check := Triangle.TriangleCheck(Triangle.Triangle{name, side1, side2, side3, area})
		if check{
			triangles = append(triangles, Triangle.Triangle{name, side1, side2, side3, area})
		}else{
			fmt.Println("Triangle cannot be constructed. Try again")
		}
		fmt.Print("Do you wish to add another one? Press 'y'/'yes' to continue: ")
		input_next, _ = reader.ReadString('\n')
		input_next = strings.TrimSpace(input_next)
	}

	return triangles
}

func main() {

	triangles := InputDetails()

	Triangle.TriangleOrder(triangles)

	fmt.Print("Input ended\n")
}