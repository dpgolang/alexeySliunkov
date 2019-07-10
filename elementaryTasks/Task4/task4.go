package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
	/*var input string
	fmt.Println("Enter file path: ")
	fmt.Scanf("%s",&input)     //getting our input(file path)*/

	fmt.Println("Enter your file path and pattern: ")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	sl1 := strings.Split(strings.TrimSpace(input1), ";")

	validInput1:= validateInput(sl1,2)
	for !validInput1{
		fmt.Println("Enter your file path and pattern: ")
		reader1 := bufio.NewReader(os.Stdin)
		input1, _ := reader1.ReadString('\n')
		sl1 = strings.Split(strings.TrimSpace(input1), ";")
	}

	inputFile:=sl1[0]
	pattern := sl1[1]
	fmt.Println("Your pattern is",pattern)

	fmt.Println("Enter file path, string-to-change and string you want to paste instead: ")
	reader2 := bufio.NewReader(os.Stdin)
	input2, _ := reader2.ReadString('\n')
	sl2 := strings.Split(strings.TrimSpace(input2), ";")

	validInput2:= validateInput(sl2,3)
	for !validInput2{
		fmt.Println("Enter your file path and pattern: ")
		reader2 := bufio.NewReader(os.Stdin)
		input2, _ := reader2.ReadString('\n')
		sl1 = strings.Split(strings.TrimSpace(input2), ";")
	}

	substFile:=sl1[0]
	stringToReplace := sl1[1]
	substituteString := sl2[2]
	fmt.Println("Your substitute string is",substituteString)

	patternCount(inputFile,pattern)
	stringSubstitute(substFile,stringToReplace,substituteString)
}

//counts number of strings where pattern was found
func patternCount(file,str string)(int,error){
	data,err:=os.Open(file)
	if err != nil{
		fmt.Println(err)

		os.Exit(1)
	}
	defer data.Close()
	scanner:=bufio.NewScanner(data)
	count:=0
	for scanner.Scan(){
		if strings.Contains(scanner.Text(),str){
			count++
		}
	}
	if err:=scanner.Err();err!=nil{
		return 0, errors.New("Error occured")
	}
	fmt.Println("faced your pattern in this file",count,"times")
	return count,nil
}

func stringSubstitute(path, pattern, substitute string) error {
	read, err := ioutil.ReadFile(path)

	if err != nil {
		return errors.New(fmt.Sprintf("Can't read this file:\n %s", err))
	}

	replacedText := strings.Replace(string(read), pattern, substitute, -1)

	err = ioutil.WriteFile(path, []byte(replacedText), 0)
	if err != nil {
		return errors.New(fmt.Sprintf("Can't write to this file:\n %s", err))
	}
	return nil
}

func validateInput(arr []string,checkInput int) bool{
	if len(arr) != checkInput {
		fmt.Println("You need to enter 4 params")
		return false
	}else {
		return true
	}
}
