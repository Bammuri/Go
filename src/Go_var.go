package main

func main() {
	var a int
	var f float32 = 11.

	a = 0
	f = 12.0

	/*동일한 타입의 변수가 복수 개 있을 경우, 변수들을 나열하고 마지막에 타입을 한번만 지정할 수 있다.
	복수 변수들이 선언된 상황에서 초기값을 지정할 수 있다. 초기값은 순서대로 변수에 할당된다.
	예를 들어, 아래 코드의 경우 i는 1, j는 2, k는 3 이 할당된다.*/
	var i, j, k int = 1, 2, 3

	/*변수를 선언하면서 초기값을 지정하지 않으면, Go는 Zero Value를 기본적으로 할당한다.
	즉, 숫자형에는 0, bool 타입에는 false, 그리고 string 형에는 "" (빈문자열)을 할당한다.*/

	/*Go 에서는 할당되는 값을 보고 그 타입을 추론하는 기능이 자주 사용된다.
	즉, 아래 코드에서 i는 정수형으로 1이 할당되고, s는 문자열로 Hi가 할당된다.*/
	var i = 1
	var s = "Hi"

	const c int = 10
	const s string = "Hi"

	const c = 10
	const s = "Hi"

	/*여러 개의 상수들 묶어서 지정할 수 있는데, 아래와 같이 괄호 안에 상수들을 나열하여 사용할 수 있다.*/
	const (
		Visa = "Visa"
		Master = "MasterCard"
		Amex = "American Express"
	)

	/*한가지 유용한 팁으로 상수값을 0부터 순차적으로 부여하기 위해 iota 라는 identifier를 사용할 수 있다.
	 이 경우 iota가 지정된 Apple에는 0이 할당되고, 나머지 상수들을 순서대로 1씩 증가된 값을 부여받는다.*/
	const (
		Apple = iota //0
		Grape
		Orange
	)

	break default func interface select
	case defer go map struct
	chan else goto package switch 
	cosnt fallthrough if range type 
	continue for import return var


}
