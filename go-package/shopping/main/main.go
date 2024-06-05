package main

import (
	"fmt"
	"time"
)

var counter = 0

func main() {
	//fmt.Println(shopping.PriceCheck(4043))
	/*
		if len(os.Args) != 2 {
			os.Exit(1)
		}

		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("not a valid number")
		} else {
			fmt.Println(n)
		}*/

	//Strings and Byte Arrays
	/*
		stra := "the spice must flow"
		byts := []byte(stra)
		strb := string(byts)
		fmt.Println(stra)
		fmt.Println(byts)
		fmt.Println(strb)
	*/
	//fmt.Println(len(""))
	/*
		fmt.Println(process(func(a int, b int) int {
			return a + b
		}))*/

	//Goroutines

	fmt.Println("Start")
	go process()
	time.Sleep(time.Millisecond * 10)
	go func() {
		fmt.Println("processing...")
	}()
	time.Sleep(time.Millisecond * 10)
	fmt.Println("Done")

}
func process() {
	fmt.Println("Processing ...")
}

/*func process(adder Add) int {
	return adder(1, 2)
}*/

// empty interface
/*func add(a interface{}, b interface{}) interface{} {
	//return a.(int) + b.(int)
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case float64:
		return a.(float64) + b.(float64)
	case string:
		return a.(string) + b.(string)
	}
	return a
}*/

type Add func(a int, b int) int

type error interface {
	Error() string
}
type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}
