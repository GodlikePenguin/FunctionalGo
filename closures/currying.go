package closures

func Add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
