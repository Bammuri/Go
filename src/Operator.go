package main

func main() {
	var a, b, c int = 1, 2, 3
	c = (a + b) / 5

	if a == b {
	}
	if a != c {
	}
	if a >= b {
	}

	var A, B, C = true, false, true

	if A && B {

	}

	if A || !(C && B) {

	}

	c = (a & b) << 5

	a = 100
	a *= 10
	a >>= 2
	a |= 1

	var k int = 10
	var p = &k
	println(*p)

}
