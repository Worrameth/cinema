package main

import (
	"fmt"

	"github.com/Worrameth/cinema/movie"
	"github.com/Worrameth/cinema/ticket"
)
func init() {
	fmt.Println("init : main")
}

func main() {
	movieName := movie.FindName("tt1825683")
	fmt.Println(movieName)
	movie.Review(movieName, 8.1)
	ticket.BuyTicket(movieName)
}
