package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)

		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

// TODO: privateのTestはどうやって書くのか？

func TestSum(t *testing.T) {
	t.Run("sum of 1,2,3,4,5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of 1,2,3,4,5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		numbers2 := []int{1, 2, 3, 4, 5}
		got := SumAll(numbers, numbers2)
		want := []int{15, 15}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 4})
	want := []int{5, 13}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSearch(t *testing.T) {
	dic := map[string]string{"test": "test is just a test"}
	got := Search(dic, "test")
	want := "test is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCallPrivateFunc(t *testing.T) {
	got := privateFunc("test")
	want := "test"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

type SpySleeper struct {
	Calls int
}

// spyは依存関係の使用方法を記録で切るモックの一種。送信された引数・呼び出された回数を記録 => Testできる
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurationSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
