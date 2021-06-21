package main

import (
	"fmt"
)

type ErrNegativeSqrt struct {
	Code int32
	Message string
}

func (e ErrNegativeSqrt) Error() string{ //error is an interface with Error() method
	 return fmt.Sprintf("Code:%d , Message: %v",e.Code,e.Message)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt{404,"Cannot pass negative numbers."}
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
