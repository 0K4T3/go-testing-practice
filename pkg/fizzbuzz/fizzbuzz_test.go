package fizzbuzz

import "testing"

func doTest(input int, expected string, t *testing.T) {
    actual := FizzBuzz(input)

    if actual != expected {
        t.Errorf("actual %v\nexpected %v", actual, expected)
    }
}

func TestGiven1Get1(t *testing.T) {
    doTest(1, "1", t)
}

func TestGiven2Get2(t *testing.T) {
    doTest(2, "2", t)
}

func TestGiven3GetFizz(t *testing.T) {
    doTest(3, "Fizz", t)
}

func TestGiven5GetBuzz(t *testing.T) {
    doTest(5, "Buzz", t)
}

func TestGiven6GetFizz(t *testing.T) {
    doTest(6, "Fizz", t)
}

func TestGiven10GetBuzz(t *testing.T) {
    doTest(10, "Buzz", t)
}

func TestGiven15GetFizzBuzz(t *testing.T) {
    doTest(15, "FizzBuzz", t)
}

func TestGiven30GetFizzBuzz(t *testing.T) {
    doTest(30, "FizzBuzz", t)
}
