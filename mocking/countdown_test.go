package main

import (
	"bytes"
	"reflect"
	"testing"
)


func TestCountdown(t *testing.T) {
	t.Run("request a countdown and returns correct answere", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		spySleeper := &SpySleeper{}
	
	
		Countdown(buffer, spySleeper)
	
		got := buffer.String()
		want := `3
2
1
Go!`
	
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	
		if spySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want: 4, got :%d", spySleeper.Calls)
		}
	})
	
	t.Run("check the orden of calls to sleep and writes to buffer", func(t *testing.T) {
		callerSpy := &SpyCountDownOperations{}
		Countdown(callerSpy, callerSpy)

		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, callerSpy.Calls) {
			t.Errorf("wanted:calls %v got %v", want, callerSpy.Calls)
		}
	})
}