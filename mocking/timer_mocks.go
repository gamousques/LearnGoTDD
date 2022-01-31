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


type SpyCountDownOperations struct{
	Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"