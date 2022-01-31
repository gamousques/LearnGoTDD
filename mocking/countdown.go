package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3

func Countdown(buffer io.Writer)  {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(buffer,  i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(buffer, finalWord)

}


func main() {
	Countdown(os.Stdout)
}