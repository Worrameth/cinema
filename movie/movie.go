package movie

import "fmt"

func init() {
	fmt.Println("init : movie")
}

func Review(name string, rating float64) {
	fmt.Printf("I reviewed %s and it'a rating is %f\n", name, rating)
}
