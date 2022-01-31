package main

import "time"

type DefaultSleeper struct {}

func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep()  {
	s.Calls++
}
