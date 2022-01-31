package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3


type Sleeper interface {
	Sleep()
}

func Countdown(buffer io.Writer, waitTime Sleeper)  {
	for i := countDownStart; i > 0; i-- {
		waitTime.Sleep()
		fmt.Fprintln(buffer,  i)
	}

	waitTime.Sleep()
	fmt.Fprint(buffer, finalWord)

}


func main() {
	slp := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, slp)
}