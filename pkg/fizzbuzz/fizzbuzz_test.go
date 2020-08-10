package fizzbuzz

import "testing"

func TestGiven1Get1(t *testing.T) {
    actual := FizzBuzz(1)
    expected := "1"

    if actual != expected {
        t.Errorf("actual %v\nexpected %v", actual, expected)
    }
}

func TestGiven2Get2(t *testing.T) {
    actual := FizzBuzz(2)
    expected := "2"

    if actual != expected {
        t.Errorf("actual %v\nexpected %v", actual, expected)
    }
}

func TestGiven3GetFizz(t *testing.T) {
    actual := FizzBuzz(3)
    expected := "Fizz"

    if actual != expected {
        t.Errorf("actual %v\nexpected %v", actual, expected)
    }
}

func TestGiven5GetBuzz(t *testing.T) {
    actual := FizzBuzz(5)
    expected := "Buzz"

    if actual != expected {
        t.Errorf("actual %v\nexpected %v", actual, expected)
    }
}

func TestGiven6GetFizz(t *testing.T) {
    actual := FizzBuzz(6)
    expected := "Fizz"

    if actual != expected {
        t.Errorf("actual %v\nexpected %v", actual, expected)
    }
}

func TestGiven10GetBuzz(t *testing.T) {
    actual := FizzBuzz(10)
    expected := "Buzz"

    if actual != expected {
        t.Errorf("actual %v\nexpected %v", actual, expected)
    }
}

func TestGiven15GetFizzBuzz(t *testing.T) {
    actual := FizzBuzz(15)
    expected := "FizzBuzz"

    if actual != expected {
        t.Errorf("actual %v\nexpected %v", actual, expected)
    }
}

func TestGiven30GetFizzBuzz(t *testing.T) {
    actual := FizzBuzz(30)
    expected := "FizzBuzz"

    if actual != expected {
        t.Errorf("actual %v\nexpected %v", actual, expected)
    }
}
