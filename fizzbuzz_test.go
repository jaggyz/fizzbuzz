package fizzbuzz

import (
	//	"strconv"
	"testing"
)

func TestInput1ShouldBeDisplay1(t *testing.T) {
	v := fizzBuzz(1)
	if "1" != v {
		t.Error("fizzBuzz of 1 should be '1' but have '", v, "'")
	}
}

func TestInput2ShouldBeDisplay2(t *testing.T) {
	v := fizzBuzz(2)
	if "2" != v {
		t.Error("fizzBuzz of 2 should be '2' but have '", v, "'")
	}
}

func TestInput3ShouldBeDisplayFizz(t *testing.T) {
	v := fizzBuzz(3)
	if "Fizz" != v {
		t.Error("fizzBuzz of 3 should be 'Fizz' but have '", v, "'")
	}
}

func TestInput4ShouldBeDisplay4(t *testing.T) {
	v := fizzBuzz(4)
	if "4" != v {
		t.Error("fizzBuzz of 4 should be '4' but have '", v, "'")
	}
}

func TestInput5ShouldBeDisplayBuzz(t *testing.T) {
	v := fizzBuzz(5)
	if "Buzz" != v {
		t.Error("fizzBuzz of 5 should be 'Buzz' but have '", v, "'")
	}
}

func TestInput15ShouldBeDisplayFizzBuzz(t *testing.T) {
	v := fizzBuzz(15)
	if "FizzBuzz" != v {
		t.Error("fizzBuzz of 15 should be 'FizzBuzz' but have '", v, "'")
	}
}

//func TestFizzBuzz(t *testing.T) {
//	sNumber := ""
//	for i := 1; i <= 15; i++ {
//		if i%15 == 0 {
//			sNumber = "FizzBuzz"
//		} else if i%3 == 0 {
//			sNumber = "Fizz"
//		} else if i%5 == 0 {
//			sNumber = "Buzz"
//		} else {
//			sNumber = strconv.Itoa(i)
//		}
//		if v := fizzBuzz(i); v != sNumber {
//			t.Error("fizzBuzz of ", i, " should be '", sNumber, "' but have '", v, "'")
//		} //else {
//			//println("'", v, "'")
//		//}
//	}
//}
