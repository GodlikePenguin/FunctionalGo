package slices

type FunctionalStringSlice []string

func (i FunctionalStringSlice) ForEach(f func(string)) {
	for _, item := range i {
		f(item)
	}
}

func (i FunctionalStringSlice) Filter(f func(string) bool) FunctionalStringSlice {
	subSlice := FunctionalStringSlice{}
	for _, item := range i {
		if f(item) {
			subSlice = append(subSlice, item)
		}
	}
	return subSlice
}

func (i FunctionalStringSlice) Map(f func(string) string) FunctionalStringSlice {
	subSlice := FunctionalStringSlice{}
	for _, item := range i {
		subSlice = append(subSlice, f(item))
	}
	return subSlice
}

func (i FunctionalStringSlice) Reduce(f func(string, string) string, starter string) string {
	toPass := starter
	for _, item := range i {
		toPass = f(toPass, item)
	}
	return toPass
}

func (i FunctionalStringSlice) ToSlice() []string {
	subSlice := []string{}
	for _, item := range i {
		subSlice = append(subSlice, item)
	}
	return subSlice
}

func FromStringSlice(s []string) FunctionalStringSlice {
	f := FunctionalStringSlice{}
	for _, si := range s {
		f = append(f, si)
	}
	return f
}
