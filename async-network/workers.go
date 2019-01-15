package main

import (
	"fmt"
	"strconv"
	"time"
)

// Count of initial numbers and goroutines in each layer
const count = 10

// Numbers produces numbers to create Sum numbers from
type Numbers struct {
	Out chan float64
}

// DoWork is standart function for do some work.
// In this case it produces slice of numbers
func (n Numbers) DoWork() []float64 {
	res := make([]float64, 0, 0)
	var i float64
	for i = 0; i < count; i++ {
		res = append(res, i)
	}
	return res
}

// Start is a standart function for receiving and sending data.
// In this case it sends numbers to channel
func (n Numbers) Start() {
	for _, i := range n.DoWork() {
		n.Out <- i
	}
}

// Sum produce sum of previous integers
type Sum struct {
	In    chan float64
	Out   chan float64
	ToLog chan string
}

// DoWork is standart function for do some work.
// In this case it calculates sum of previous integers
func (s Sum) DoWork(i float64) float64 {
	if i > 0 {
		return i + s.DoWork(i-1)
	}
	return 0
}

// Start is a standart function for receiving and sending data.
// In this case it gets number from in channel, calculates sum for it and sends it to out channel.
// Also it sends data to logger
func (s Sum) Start() {
	for nmbr := range s.In {
		sum := s.DoWork(nmbr)
		s.Out <- sum
		s.ToLog <- fmt.Sprintf("Sum number: %s",
			strconv.FormatFloat(sum, 'f', 0, 64))
	}
}

// Factorial calculates factorial
type Factorial struct {
	In    chan float64
	ToLog chan string
}

// DoWork is standart function for do some work.
// In this case it calculates factorial of number
func (f Factorial) DoWork(i float64) float64 {
	if i > 0 {
		return i * f.DoWork(i-1)
	}
	return 1
}

// Start is a standart function for receiving and sending data.
// In this case it gets number from in channel and calculates its factorial.
// Also it sends data to logger
func (f Factorial) Start() {
	for sum := range f.In {
		fctrl := f.DoWork(sum)
		f.ToLog <- fmt.Sprintf("Factorial of %s is: %s",
			strconv.FormatFloat(sum, 'f', 0, 64), strconv.FormatFloat(fctrl, 'e', 0, 64))
	}
}

// Logger prints incoming data
type Logger struct {
	In  chan string
	Out chan interface{}
}

// DoWork is standart function for do some work.
// In this case it prints incoming string
func (l Logger) DoWork(msg string) {
	fmt.Println(msg)
}

// Start is a standart function for receiving and sending data.
// In this case it prints incoming.
// Also it determines the finishing of calculations by timer
func (l Logger) Start() {
	timer := time.NewTimer(1 * time.Second)
	for {
		select {
		case msg := <-l.In:
			l.DoWork(msg)
			timer.Reset(1 * time.Second)
		case <-timer.C:
			l.Out <- ""
			return
		}

	}
}
