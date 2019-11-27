package gotuple

type IntTuple struct{
	First,Second,Third int
}

//QuicksortIntTupleFirst sort slice of IntTuple from low to high value of @IntTuple.First with quicksort algorithm
//returns sorted slice
func QuicksortIntTupleFirst(a []IntTuple) []IntTuple {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) >> 1

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].First < a[right].First {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuicksortIntTupleFirst(a[:left])
	QuicksortIntTupleFirst(a[left+1:])

	return a
}

//QuicksortIntTupleSecond sort slice of IntTuple from low to high value of @IntTuple.Second with quicksort algorithm
//returns sorted slice
func QuicksortIntTupleSecond(a []IntTuple) []IntTuple {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) >> 1

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].Second < a[right].Second {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuicksortIntTupleSecond(a[:left])
	QuicksortIntTupleSecond(a[left+1:])

	return a
}

//QuicksortIntTupleThird sort slice of IntTuple from low to high value of @IntTuple.Third with quicksort algorithm
//returns sorted slice
func QuicksortIntTupleThird(a []IntTuple) []IntTuple {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) >> 1

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].Third < a[right].Third {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuicksortIntTupleThird(a[:left])
	QuicksortIntTupleThird(a[left+1:])

	return a
}

//ReversIntTupleSlice returns reversed slice of @a IntTuple
func ReversIntTupleSlice(a []IntTuple) []IntTuple {
	res := make([]IntTuple, len(a))

	for i, num := range a {
		res[len(a)-i-1] = num
	}

	return res
}