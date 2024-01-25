package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return "Hello, " + name
}

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

// func SumAllTails(numbersToSum ...[]int) []int {
// 	var sums []int
// 	for _, numbers := range numbersToSum {
// 		tail := numbers[1:]
// 		sums = append(sums, Sum(tail))
// 	}

//		return sums
//	}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}

type Dictionary map[string]string

func Search(dictionary Dictionary, word string) string {
	return dictionary[word]
}

func main() {
	sleeper := &ConfigurationSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func dangerous() {
	// panic: assignment to entry in nil map
	var m map[string]string
	m["key"] = "value"
}

func privateFunc(name string) string {
	// private
	fmt.Print("private")
	return name
}

// Stub / Mock
const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type ConfigurationSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurationSleeper) Sleep() {
	c.sleep(c.duration)
}
