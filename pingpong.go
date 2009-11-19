package main

import "fmt"
import "time"

const nano = 1000000000

func player(in <-chan string, out chan<- string, name string) {
	for {
		fmt.Println(<-in);
		out <- name;
		time.Sleep(1 * nano);
	}
}

func main() {
	c0 := make(chan string);
	c1 := make(chan string);

	go player(c0, c1, "ping");
	go player(c1, c0, "pong");
	
  // cast the ball into the arena
	c0 <- "serve";
	
  time.Sleep(10 * nano);
}
