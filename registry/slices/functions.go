package slices

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"

	"github.com/go-sprout/sprout/internal/helpers"
	"github.com/spf13/cast"
)

// List creates a list from the provided elements.
//
// Parameters:
//
//	values ...any - the elements to include in the list.
//
// Returns:
//
//	[]any - the created list containing the provided elements.
//
// Example:
//
//	{{ 1, 2, 3 | list }} // Output: [1, 2, 3]
func (sr *SlicesRegistry) List(values ...any) []any {
	return values
}

// Append adds an element to the end of the list.
//
// Parameters:
//
//	list any - the original list to append to.
//	v any - the element to append.
//
// Returns:
//
//	[]any - the new list with the element appended.
//
// Example:
//
//	{{ append ["a", "b"], "c" }} // Output: ["a", "b", "c"]
func (sr *SlicesRegistry) Append(list any, v any) []any {
	result, err := sr.MustAppend(list, v)
	if err != nil {
		return []any{}
		// panic(err)
	}

	return result
}

// Prepend adds an element to the beginning of the list.
//
// Parameters:
//
//	list any - the original list to prepend to.
//	v any - the element to prepend.
//
// Returns:
//
//	[]any - the new list with the element prepended.
//
// Example:
//
//	{{ prepend  ["b", "c"], "a" }} // Output: ["a", "b", "c"]
func (sr *SlicesRegistry) Prepend(list any, v any) []any {
	result, err := sr.MustPrepend(list, v)
	if err != nil {
		return []any{}
		// panic(err)
	}

	return result
}

// Concat merges multiple lists into a single list.
//
// Parameters:
//
//	lists ...any - the lists to concatenate.
//
// Returns:
//
//	any - a single concatenated list containing elements from all provided lists.
//
// Example:
//
//	{{ ["c", "d"] | concat ["a", "b"] }} // Output: ["a", "b", "c", "d"]
func (sr *SlicesRegistry) Concat(lists ...any) any {
	var res []any
	for _, list := range lists {
		if list == nil {
			continue
		}

		tp := reflect.TypeOf(list).Kind()
		switch tp {
		case reflect.Slice, reflect.Array:
			valueOfList := reflect.ValueOf(list)
			for i := 0; i < valueOfList.Len(); i++ {
				res = append(res, valueOfList.Index(i).Interface())
			}
		default:
			continue
			// panic(fmt.Sprintf("cannot concat type %s as list", tp))
		}
	}
	return res
}

// Chunk divides a list into chunks of specified size.
//
// Parameters:
//
//	size int - the size of each chunk.
//	list any - the list to divide.
//
// Returns:
//
//	[][]any - a list of chunks.
//
// Example:
//
//	{{ chunk 2, ["a", "b", "c", "d"] }} // Output: [["a", "b"], ["c", "d"]]
func (sr *SlicesRegistry) Chunk(size int, list any) [][]any {
	result, err := sr.MustChunk(size, list)
	if err != nil {
		return [][]any{}
		// panic(err)
	}

	return result
}

// Uniq removes duplicate elements from a list.
//
// Parameters:
//
//	list any - the list from which to remove duplicates.
//
// Returns:
//
//	[]any - a list containing only unique elements.
//
// Example:
//
//	{{ ["a", "b", "a", "c"] | uniq }} // Output: ["a", "b", "c"]
func (sr *SlicesRegistry) Uniq(list any) []any {
	result, err := sr.MustUniq(list)
	if err != nil {
		return []any{}
		// panic(err)
	}

	return result
}

// Compact removes nil and zero-value elements from a list.
//
// Parameters:
//
//	list any - the list to compact.
//
// Returns:
//
//	[]any - the list without nil or zero-value elements.
//
// Example:
//
//	{{ [0, 1, nil, 2, "", 3] | compact }} // Output: [1, 2, 3]
func (sr *SlicesRegistry) Compact(list any) []any {
	result, err := sr.MustCompact(list)
	if err != nil {
		return []any{}
		// panic(err)
	}

	return result
}

// Slice extracts a slice from a list between two indices.
//
// Parameters:
//
//	list any - the list to slice.
//	indices ...any - the start and optional end indices; if end is omitted,
//
// slices to the end.
//
// Returns:
//
//	any - the sliced part of the list.
//
// Example:
//
//	{{ slice [1, 2, 3, 4, 5], 1, 3 }} // Output: [2, 3]
func (sr *SlicesRegistry) Slice(list any, indices ...any) any {
	result, err := sr.MustSlice(list, indices...)
	if err != nil {
		return []any{}
		// panic(err)
	}

	return result
}

// Has checks if the specified element is present in the collection.
//
// Parameters:
//
//	element any - the element to search for.
//	list any - the collection to search.
//
// Returns:
//
//	bool - true if the element is found, otherwise false.
//
// Example:
//
//	{{ ["value", "other"] | has "value" }} // Output: true
func (sr *SlicesRegistry) Has(element any, list any) bool {
	result, _ := sr.MustHas(element, list)
	return result
}

// Without returns a new list excluding specified elements.
//
// Parameters:
//
//	list any - the original list.
//	omit ...any - elements to exclude from the new list.
//
// Returns:
//
//	[]any - the list excluding the specified elements.
//
// Example:
//
//	{{ without [1, 2, 3, 4], 2, 4 }} // Output: [1, 3]
func (sr *SlicesRegistry) Without(list any, omit ...any) []any {
	result, err := sr.MustWithout(list, omit...)
	if err != nil {
		return []any{}
		// panic(err)
	}

	return result
}

// Rest returns all elements of a list except the first.
//
// Parameters:
//
//	list any - the list to process.
//
// Returns:
//
//	[]any - the list without the first element.
//
// Example:
//
//	{{ [1, 2, 3, 4] | rest }} // Output: [2, 3, 4]
func (sr *SlicesRegistry) Rest(list any) []any {
	result, err := sr.MustRest(list)
	if err != nil {
		return []any{}
		// panic(err)
	}

	return result
}

// Initial returns all elements of a list except the last.
//
// Parameters:
//
//	list any - the list to process.
//
// Returns:
//
//	[]any - the list without the last element.
//
// Example:
//
//	{{ [1, 2, 3, 4] | initial }} // Output: [1, 2, 3]
func (sr *SlicesRegistry) Initial(list any) []any {
	result, err := sr.MustInitial(list)
	if err != nil {
		return []any{}
		// panic(err)
	}

	return result
}

// First returns the first element of a list.
//
// Parameters:
//
//	list any - the list from which to take the first element.
//
// Returns:
//
//	any - the first element of the list.
//
// Example:
//
//	{{ [1, 2, 3, 4] | first }} // Output: 1
func (sr *SlicesRegistry) First(list any) any {
	result, err := sr.MustFirst(list)
	if err != nil {
		return nil
		// panic(err)
	}

	return result
}

// Last returns the last element of a list.
//
// Parameters:
//
//	list any - the list from which to take the last element.
//
// Returns:
//
//	any - the last element of the list.
//
// Example:
//
//	{{ [1, 2, 3, 4] | last }} // Output: 4
func (sr *SlicesRegistry) Last(list any) any {
	result, err := sr.MustLast(list)
	if err != nil {
		return nil
		// panic(err)
	}

	return result
}

// Reverse returns a new list with the elements in reverse order.
//
// Parameters:
//
//	list any - the list to reverse.
//
// Returns:
//
//	[]any - the list in reverse order.
//
// Example:
//
//	{{ [1, 2, 3, 4] | reverse }} // Output: [4, 3, 2, 1]
func (sr *SlicesRegistry) Reverse(list any) []any {
	result, err := sr.MustReverse(list)
	if err != nil {
		return []any{}
		// panic(err)
	}

	return result
}

// SortAlpha sorts a list of strings in alphabetical order.
//
// Parameters:
//
//	list any - the list of strings to sort.
//
// Returns:
//
//	[]string - the sorted list.
//
// Example:
//
//	{{ ["d", "b", "a", "c"] | sortAlpha }} // Output: ["a", "b", "c", "d"]
func (sr *SlicesRegistry) SortAlpha(list any) []string {
	kind := reflect.Indirect(reflect.ValueOf(list)).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		sortedList := sort.StringSlice(sr.StrSlice(list))
		sortedList.Sort()
		return sortedList
	}

	return []string{helpers.ToString(list)}
}

// SplitList divides a string into a slice of substrings separated by the
// specified separator.
//
// ! FUTURE: Rename this function to be more explicit
//
// Parameters:
//
//	sep string - the delimiter used to split the string.
//	str string - the string to split.
//
// Returns:
//
//	[]string - a slice containing the substrings obtained from splitting the input string.
//
// Example:
//
//	{{ "one, two, three" | splitList ", " }} // Output: ["one", "two", "three"]
func (sr *SlicesRegistry) SplitList(sep string, str string) []string {
	return strings.Split(str, sep)
}

// StrSlice converts a value to a slice of strings, handling various types
// including []string, []any, and other slice types.
//
// Parameters:
//
//	value any - the value to convert to a slice of strings.
//
// Returns:
//
//	[]string - the converted slice of strings.
//
// Example:
//
//	{{ strSlice any["a", "b", "c"] }} // Output: ["a", "b", "c"]
func (sr *SlicesRegistry) StrSlice(value any) []string {
	return helpers.StrSlice(value)
}

// Until generates a slice of integers from 0 up to but not including 'count'.
// If 'count' is negative, it produces a descending slice from 0 down to 'count',
// inclusive, with a step of -1. The function leverages UntilStep to specify
// the range and step dynamically.
//
// Parameters:
//   count int - the endpoint (exclusive) of the range to generate.
//
// Returns:
//   []int - a slice of integers from 0 to 'count' with the appropriate step
//           depending on whether 'count' is positive or negative.
//
// Example:
//   {{ 5 | until }} // Output: [0 1 2 3 4]
//   {{ -3 | until }} // Output: [0 -1 -2]

func (sr *SlicesRegistry) Until(count int) []int {
	step := 1
	if count < 0 {
		step = -1
	}
	return sr.UntilStep(0, count, step)
}

// UntilStep generates a slice of integers from 'start' to 'stop' (exclusive),
// incrementing by 'step'. If 'step' is positive, the sequence increases; if
// negative, it decreases. The function returns an empty slice if the sequence
// does not make logical sense (e.g., positive step when start is greater than
// stop or vice versa).
//
// Parameters:
//
//	start int - the starting point of the sequence.
//	stop int - the endpoint (exclusive) of the sequence.
//	step int - the increment between elements in the sequence.
//
// Returns:
//
//	[]int - a dynamically generated slice of integers based on the input
//	        parameters, or an empty slice if the parameters are inconsistent
//	        with the desired range and step.
//
// Example:
//
//	{{ 0, 10, 2 | untilStep }} // Output: [0 2 4 6 8]
//	{{ 10, 0, -2 | untilStep }} // Output: [10 8 6 4 2]
func (sr *SlicesRegistry) UntilStep(start, stop, step int) []int {
	return helpers.UntilStep(start, stop, step)
}

// MustAppend appends an element to a slice or array, returning an error if the
// operation isn't applicable.
//
// Parameters:
//
//	list any - the original list to append to.
//	v any - the element to append.
//
// Returns:
//
//	[]any - the new list with the element appended.
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ mustAppend ["a", "b"], "c"  }} // Output: ["a", "b", "c"], nil
func (sr *SlicesRegistry) MustAppend(list any, v any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot append to nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		result := make([]any, length)
		for i := 0; i < length; i++ {
			result[i] = valueOfList.Index(i).Interface()
		}

		return append(result, v), nil

	default:
		return nil, fmt.Errorf("cannot append on type %s", tp)
	}
}

// MustPrepend prepends an element to a slice or array, returning an error if
// the operation isn't applicable.
//
// Parameters:
//
//	list any - the original list to prepend to.
//	v any - the element to prepend.
//
// Returns:
//
//	[]any - the new list with the element prepended.
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ mustPrepend ["b", "c"], "a" }} // Output: ["a", "b", "c"], nil
func (sr *SlicesRegistry) MustPrepend(list any, v any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot prepend to nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		result := make([]any, length)
		for i := 0; i < length; i++ {
			result[i] = valueOfList.Index(i).Interface()
		}

		return append([]any{v}, result...), nil

	default:
		return nil, fmt.Errorf("cannot prepend on type %s", tp)
	}
}

// MustChunk divides a list into chunks of specified size, returning an error
// if the list is nil or not a slice/array.
//
// Parameters:
//
//	size int - the maximum size of each chunk.
//	list any - the list to chunk.
//
// Returns:
//
//	[][]any - a list of chunks.
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ ["a", "b", "c", "d"] | mustChunk 2 }} // Output: [["a", "b"], ["c", "d"]], nil
func (sr *SlicesRegistry) MustChunk(size int, list any) ([][]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot chunk nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()

		chunkCount := int(math.Floor(float64(length-1)/float64(size)) + 1)
		result := make([][]any, chunkCount)

		for i := 0; i < chunkCount; i++ {
			chunkLength := size
			if i == chunkCount-1 {
				chunkLength = int(math.Floor(math.Mod(float64(length), float64(size))))
				if chunkLength == 0 {
					chunkLength = size
				}
			}

			result[i] = make([]any, chunkLength)

			for j := 0; j < chunkLength; j++ {
				ix := i*size + j
				result[i][j] = valueOfList.Index(ix).Interface()
			}
		}

		return result, nil

	default:
		return nil, fmt.Errorf("cannot chunk type %s", tp)
	}
}

// MustUniq returns a new slice containing unique elements of the given list,
// preserving order.
//
// Parameters:
//
//	list any - the list from which to remove duplicates.
//
// Returns:
//
//	[]any - a list containing only the unique elements.
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ ["a", "b", "a", "c"] | mustUniq }} // Output: ["a", "b", "c"], nil
func (sr *SlicesRegistry) MustUniq(list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot uniq nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		result := []any{}
		var item any
		for i := 0; i < length; i++ {
			item = valueOfList.Index(i).Interface()
			if !sr.inList(result, item) {
				result = append(result, item)
			}
		}

		return result, nil
	default:
		return nil, fmt.Errorf("cannot find uniq on type %s", tp)
	}
}

// MustCompact removes nil or zero-value elements from a list.
//
// Parameters:
//
//	list any - the list to compact.
//
// Returns:
//
//	[]any - the list without nil or zero-value elements.
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ [0, 1, nil, 2, "", 3] | mustCompact }} // Output: [1, 2, 3], nil
func (sr *SlicesRegistry) MustCompact(list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot compact nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		result := []any{}
		var item any
		for i := 0; i < length; i++ {
			item = valueOfList.Index(i).Interface()
			if !helpers.Empty(item) {
				result = append(result, item)
			}
		}

		return result, nil
	default:
		return nil, fmt.Errorf("cannot compact on type %s", tp)
	}
}

// MustSlice extracts a slice from a list between two indices.
//
// Parameters:
//
//	list any - the list to slice.
//	indices ...any - the start and optional end indices; if end is omitted,
//
// slices to the end.
//
// Returns:
//
//	any - the sliced part of the list.
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ mustSlice [1, 2, 3, 4, 5], 1, 3 }} // Output: [2, 3], nil
func (sr *SlicesRegistry) MustSlice(list any, indices ...any) (any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot slice nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		if length == 0 {
			return nil, nil
		}

		var start, end int
		if len(indices) > 0 {
			start = cast.ToInt(indices[0])
		}
		if len(indices) < 2 {
			end = length
		} else {
			end = cast.ToInt(indices[1])
		}

		return valueOfList.Slice(start, end).Interface(), nil
	default:
		return nil, fmt.Errorf("list should be type of slice or array but %s", tp)
	}
}

// MustHas checks if a specified element is present in a collection and handles
// type errors.
//
// Parameters:
//
//	element any - the element to search for in the collection.
//	list any - the collection in which to search for the element.
//
// Returns:
//
//	bool - true if the element is found, otherwise false.
//	error - error if the list is not a type that can be searched (not a slice or array).
//
// Example:
//
//	{{ [1, 2, 3, 4] | mustHas 3 }} // Output: true, nil
func (sr *SlicesRegistry) MustHas(element any, list any) (bool, error) {
	if list == nil {
		return false, nil
	}
	typeOfList := reflect.TypeOf(list).Kind()
	switch typeOfList {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)
		var item any
		length := valueOfList.Len()
		for i := 0; i < length; i++ {
			item = valueOfList.Index(i).Interface()
			if reflect.DeepEqual(element, item) {
				return true, nil
			}
		}

		return false, nil
	default:
		return false, fmt.Errorf("cannot find has on type %s", typeOfList)
	}
}

// MustWithout returns a new list excluding specified elements.
//
// Parameters:
//
//	list any - the original list.
//	omit ...any - elements to exclude from the new list.
//
// Returns:
//
//	[]any - the list excluding the specified elements.
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ mustWithout [1, 2, 3, 4], 2, 4 }} // Output: [1, 3], nil
func (sr *SlicesRegistry) MustWithout(list any, omit ...any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot without nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		result := []any{}
		var item any
		for i := 0; i < length; i++ {
			item = valueOfList.Index(i).Interface()
			if !sr.inList(omit, item) {
				result = append(result, item)
			}
		}

		return result, nil
	default:
		return nil, fmt.Errorf("cannot find without on type %s", tp)
	}
}

// MustRest returns all elements of a list except the first.
//
// Parameters:
//
//	list any - the list to process.
//
// Returns:
//
//	[]any - the list without the first element.
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ [1, 2, 3, 4] | mustRest }} // Output: [2, 3, 4], nil
func (sr *SlicesRegistry) MustRest(list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot rest nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		if length == 0 {
			return nil, nil
		}

		result := make([]any, length-1)
		for i := 1; i < length; i++ {
			result[i-1] = valueOfList.Index(i).Interface()
		}

		return result, nil
	default:
		return nil, fmt.Errorf("cannot find rest on type %s", tp)
	}
}

// MustInitial returns all elements of a list except the last.
//
// Parameters:
//
//	list any - the list to process.
//
// Returns:
//
//	[]any - the list without the last element.
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ [1, 2, 3, 4] | mustInitial }} // Output: [1, 2, 3], nil
func (sr *SlicesRegistry) MustInitial(list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot initial nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		if length == 0 {
			return nil, nil
		}

		result := make([]any, length-1)
		for i := 0; i < length-1; i++ {
			result[i] = valueOfList.Index(i).Interface()
		}

		return result, nil
	default:
		return nil, fmt.Errorf("cannot find initial on type %s", tp)
	}
}

// MustFirst returns the first element of a list.
//
// Parameters:
//
//	list any - the list from which to take the first element.
//
// Returns:
//
//	any - the first element of the list.
//	error - error if the list is nil, empty, or not a slice/array.
//
// Example:
//
//	{{ [1, 2, 3, 4] | mustFirst }} // Output: 1, nil
func (sr *SlicesRegistry) MustFirst(list any) (any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot first nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		if length == 0 {
			return nil, nil
		}

		return valueOfList.Index(0).Interface(), nil
	default:
		return nil, fmt.Errorf("cannot find first on type %s", tp)
	}
}

// MustLast returns the last element of a list.
//
// Parameters:
//
//	list any - the list from which to take the last element.
//
// Returns:
//
//	any - the last element of the list.
//	error - error if the list is nil, empty, or not a slice/array.
//
// Example:
//
//	{{ [1, 2, 3, 4] | mustLast }} // Output: 4, nil
func (sr *SlicesRegistry) MustLast(list any) (any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot last nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		if length == 0 {
			return nil, nil
		}

		return valueOfList.Index(length - 1).Interface(), nil
	default:
		return nil, fmt.Errorf("cannot find last on type %s", tp)
	}
}

// MustReverse returns a new list with the elements in reverse order.
//
// Parameters:
//
//	list any - the list to reverse.
//
// Returns:
//
//	[]any - the list in reverse order.
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ [1, 2, 3, 4] | mustReverse }} // Output: [4, 3, 2, 1], nil
func (sr *SlicesRegistry) MustReverse(list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot reverse nil")
	}

	tp := reflect.TypeOf(list).Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		valueOfList := reflect.ValueOf(list)

		length := valueOfList.Len()
		// We do not sort in place because the incoming array should not be altered.
		nl := make([]any, length)
		for i := 0; i < length; i++ {
			nl[length-i-1] = valueOfList.Index(i).Interface()
		}

		return nl, nil
	default:
		return nil, fmt.Errorf("cannot find reverse on type %s", tp)
	}
}
