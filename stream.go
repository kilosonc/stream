package stream

func ForEach[S any](sources []S, fc func(S, int)) {
	for i := range sources {
		fc(sources[i], i)
	}
}

func Map[S any,R any](sources []S, fc func(S) R) []R {
	result := make([]R, 0,len(sources))
	ForEach(sources, func(source S, _ int) {
		result = append(result, fc(source))
	})
	return result
}

func Filter[T any](sources []T, fc func(T) bool) []T {
	result := make([]T,0)
	ForEach(sources, func(source T, _ int) {
		if fc(source) {
			result = append(result, source)
		}
	})
	return result
}

func FlatMap[S any, R any](sources [][]S, fc func([]S) []R) []R {
	return Flatten(Map(sources,fc))
}

func Flatten[S any](sources [][]S) []S {
	result := make([]S, 0)
	ForEach(sources, func(s []S, _ int) {
		result = append(result, s...)
	})
	return result
}

func Reduce[S any, R any](sources []S, fc func(R, S, int) R, accumulator R) R {
	ForEach(sources, func(source S, index int) {
		accumulator = fc(accumulator, source, index)
	})
	return accumulator;
}

// Fill fills elements of array with value from start up to, but not including, end.
// ***this method will cause `sources` changing***
func Fill[T any](sources []T, generator func() T) []T {
	for i := range sources {
		sources[i] = generator()
	}
	return sources
}



// Find iterates over elements of collection,
// returning the first element's index `fc` returns truthy for,
// -1 for not found
func Find[T any](sources []T, fc func(s T) bool) int {
	for i := range sources {
		if fc(sources[i]) {
			return i
		}
	}
	return -1
}

func Includes[T any](sources []T, equal func(s T) bool) bool {
	return Find(sources, equal) != -1
}
