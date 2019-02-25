package slices

type FunctionalIntSlice []int

func (i FunctionalIntSlice) ForEach(f func(int)) {
	for _, item := range i {
		f(item)
	}
}

func (i FunctionalIntSlice) Filter(f func(int) bool) FunctionalIntSlice {
	subSlice := FunctionalIntSlice{}
	for _, item := range i {
		if f(item) {
			subSlice = append(subSlice, item)
		}
	}
	return subSlice
}

func (i FunctionalIntSlice) Map(f func(int) int) FunctionalIntSlice {
	subSlice := FunctionalIntSlice{}
	for _, item := range i {
		subSlice = append(subSlice, f(item))
	}
	return subSlice
}

func (i FunctionalIntSlice) Reduce(f func(int, int) int, starter int) int {
	toPass := starter
	for _, item := range i {
		toPass = f(toPass, item)
	}
	return toPass
}

func (i FunctionalIntSlice) ToSlice() []int {
	subSlice := []int{}
	for _, item := range i {
		subSlice = append(subSlice, item)
	}
	return subSlice
}

func FromIntSlice(s []int) FunctionalIntSlice {
	f := FunctionalIntSlice{}
	for _, si := range s {
		f = append(f, si)
	}
	return f
}
