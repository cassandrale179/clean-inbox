package main

// Libraries
import (
	"fmt"
  "sync"
)

var wg sync.WaitGroup

// Function one
go func count_loop() []int {
  int_arr := []int{}
  for i := 0; i < 10; i++ {
		int_arr = append(int_arr, i)
    fmt.Println("first loop running")
	}
  return int_arr
}


//Function two
go func double_value(int_arr []int){
  for i := 0; i < 10; i++ {
    int_arr[i] = int_arr[i]* 2
    fmt.Println("second loop running")
	}
  fmt.Println("after modify", int_arr)
}


// Main function
func main() {
  fmt.Println(“start”)
  wg.Add(1) // indicate we are going to wait for one thing

	int_arr := count_loop()
  fmt.Println("int_arr", int_arr)
  double_value(int_arr)


}
