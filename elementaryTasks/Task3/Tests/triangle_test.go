package main

import (
	"Task3/Triangle"
	"errors"
	"fmt"
	"testing"
)

/*func TestArea(t *testing.T){
	tr:= Triangle.Triangle{"tr1",4,5,6,0}
	tr.Area = Triangle.GetTriangleArea(tr)
	if tr.Area==0{
		t.Errorf("Something went wrong while getting triangle params! Area is\"%v\"",tr.Area)
	}else{
		fmt.Println("Test passed successfully! Area equals:",tr.Area)
	}
}*/

/*func TestInput(t *testing.T) {
	var tr1 Triangle.Triangle
	tr1.Name = "triangle one"
	tr1.Side1,tr1.Side2,tr1.Side3 = 4,5,6
	side1:= strconv.FormatFloat(tr1.Side1,'E',-1,64,)
	side2:= strconv.FormatFloat(tr1.Side1,'E',-1,64,)
	//side3:= strconv.FormatFloat(tr1.Side1,'E',-1,64,)
	name:=tr1.Name
	//var sl = []string{name,side1,side2,side3}
	var sl2 = []string{name,side1,side2}
	err := len(sl2)
	if err!=4{
		t.Errorf("Invalid params number! Expected 4, got \"%v\"",err)
	}
	/*if Triangle.ValidateInput(sl2)==true{
		fmt.Println("Test passed successfully! Params:",len(sl2))
	}else{
		t.Errorf("Invalid params number! Expected 4, got \"%v\"",len(sl2))
	}
}*/

func TestInput(t *testing.T) {
	type wanted struct {
		params bool
		err    error
	}
	var tests = []struct {
		input []string
		want wanted
	}{
		{[]string{"tr1","4","5","6"},wanted{true,nil}},
		{[]string{"tr1","4","5"},wanted{false,errors.New("you need to enter 4 params")}},
	}
	for _, test := range tests {
		inp1,err := Triangle.ValidateInput(test.input)
		var got = wanted{inp1,err}
		fmt.Println(got)
		fmt.Println(test.want)
		if  got != test.want {
			t.Errorf("Input params violation\"%v\", \"%s\"", len(test.input),test.want.err.Error())
		}
	}
}


