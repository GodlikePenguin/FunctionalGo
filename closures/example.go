package closures

import (
	"fmt"
	"time"
)

func CurryExample() {
	//Add is a function which takes an int and returns another function that takes the second int and returns the result.
	//The int passed to the first function is 'remembered' by the second function even though it is not a direct argument
	result := Add(5)(10)
	fmt.Println(result)
	// Prints 15
}

func ClosureExample() {
	result := LogLater("Hi")
	time.Sleep(1 * time.Second)
	fmt.Println(result())
	// Prints Hi with the time the original function call was made, not the time result was invoked

	logresult := AnnotatedLogger("example")
	logresult("log message")
	// Prints "example: log message"
}
