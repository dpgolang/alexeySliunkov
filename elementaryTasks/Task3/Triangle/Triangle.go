package Triangle

import (
	"fmt"
	"math"
)

type Triangle struct {
	Name  string
	Side1 float64
	Side2 float64
	Side3 float64
	Area float64
}

func GetTriangleArea(t Triangle) float64 {
	p := (t.Side1 + t.Side2 + t.Side3) / 2
	return math.Sqrt(p * (p - t.Side1) * (p - t.Side2) * (p - t.Side3))
}

func TriangleOrder(triangles []Triangle){
	for i:=0;i<len(triangles);i++{
		for j:=0;j<len(triangles)-1;j++{
			if triangles[j].Area < triangles[j+1].Area{
				tr:= triangles[j+1]
				triangles[j+1] = triangles[j]
				triangles[j] = tr
			}
		}}
	for _, t := range triangles {
		fmt.Println(t.Name,",",t.Area)
	}
}

func TriangleCheck(t Triangle) bool{
	check:= false
	if t.Side1+t.Side2>t.Side3 && t.Side1+t.Side3>t.Side2 && t.Side3+t.Side2>t.Side1{
		fmt.Println("Successful")
		check = true
	}
	return check
}

func ValidateInput(arr []string) (bool,error){
	var err error
	if len(arr) != 4 {
		fmt.Println("You need to enter 4 params")
		return false,fmt.Errorf("\"%s\"",err)
	}
	return true,err
}

