package main

import "time"




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

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (cs ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (st *SpyTime) Sleep(duration time.Duration) {
	st.durationSlept = duration
}