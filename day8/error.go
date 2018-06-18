package main

import ("errors";
	"fmt")

func err_func (x int64) (int64, error) {
	if x > 5 {
		err := errors.New("this is error")
		return 0, err
	}
	return x , nil
}

func main() {
	_, err := err_func(9)
	if err != nil {
		fmt.Println(err)
	}
}
