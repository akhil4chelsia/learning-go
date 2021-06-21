package main
import "fmt"
func main(){
	var s Shape = &Rectangle{10,20} 
	s.Scale(0.5)
	fmt.Println(s.Area())
}

type Shape interface{
	Area() float64
	Scale(factor float64)
}

type Rectangle struct{
	length, breadth float64
}

func (r *Rectangle) Area() float64{
	return r.length * r.breadth
}

func (r *Rectangle) Scale(factor float64){
	r.length *= factor
	r.breadth *= factor
	
}