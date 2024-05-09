package sprout

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
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
func (fh *FunctionHandler) List(values ...any) []any {
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
func (fh *FunctionHandler) Append(list any, v any) []any {
	result, err := fh.MustAppend(list, v)
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
func (fh *FunctionHandler) Prepend(list any, v any) []any {
	result, err := fh.MustPrepend(list, v)
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
func (fh *FunctionHandler) Concat(lists ...any) any {
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
func (fh *FunctionHandler) Chunk(size int, list any) [][]any {
	result, err := fh.MustChunk(size, list)
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
func (fh *FunctionHandler) Uniq(list any) []any {
	result, err := fh.MustUniq(list)
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
func (fh *FunctionHandler) Compact(list any) []any {
	result, err := fh.MustCompact(list)
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
func (fh *FunctionHandler) Slice(list any, indices ...any) any {
	result, err := fh.MustSlice(list, indices...)
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
func (fh *FunctionHandler) Has(element any, list any) bool {
	result, _ := fh.MustHas(element, list)
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
func (fh *FunctionHandler) Without(list any, omit ...any) []any {
	result, err := fh.MustWithout(list, omit...)
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
func (fh *FunctionHandler) Rest(list any) []any {
	result, err := fh.MustRest(list)
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
func (fh *FunctionHandler) Initial(list any) []any {
	result, err := fh.MustInitial(list)
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
func (fh *FunctionHandler) First(list any) any {
	result, err := fh.MustFirst(list)
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
func (fh *FunctionHandler) Last(list any) any {
	result, err := fh.MustLast(list)
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
func (fh *FunctionHandler) Reverse(list any) []any {
	result, err := fh.MustReverse(list)
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
func (fh *FunctionHandler) SortAlpha(list any) []string {
	kind := reflect.Indirect(reflect.ValueOf(list)).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		sortedList := sort.StringSlice(fh.StrSlice(list))
		sortedList.Sort()
		return sortedList
	}
	return []string{fh.ToString(list)}
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
//	{{ ", ", "one, two, three" | splitList }} // Output: ["one", "two", "three"]
func (fh *FunctionHandler) SplitList(sep string, str string) []string {
	return strings.Split(str, sep)
}

func (fh *FunctionHandler) StrSlice(value any) []string {
	if value == nil {
		return []string{}
	}

	// Handle []string type efficiently without reflection.
	if strs, ok := value.([]string); ok {
		return strs
	}

	// For slices of any, convert each element to a string, skipping nil values.
	if interfaces, ok := value.([]any); ok {
		var result []string
		for _, s := range interfaces {
			if s != nil {
				result = append(result, fh.ToString(s))
			}
		}
		return result
	}

	// Use reflection for other slice types to convert them to []string.
	reflectedValue := reflect.ValueOf(value)
	if reflectedValue.Kind() == reflect.Slice || reflectedValue.Kind() == reflect.Array {
		var result []string
		for i := 0; i < reflectedValue.Len(); i++ {
			value := reflectedValue.Index(i).Interface()
			if value != nil {
				result = append(result, fh.ToString(value))
			}
		}
		return result
	}

	// If it's not a slice, array, or nil, return a slice with the string representation of v.
	return []string{fh.ToString(value)}
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
func (fh *FunctionHandler) MustAppend(list any, v any) ([]any, error) {
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
func (fh *FunctionHandler) MustPrepend(list any, v any) ([]any, error) {
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
func (fh *FunctionHandler) MustChunk(size int, list any) ([][]any, error) {
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
func (fh *FunctionHandler) MustUniq(list any) ([]any, error) {
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
			if !fh.InList(result, item) {
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
func (fh *FunctionHandler) MustCompact(list any) ([]any, error) {
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
			if !fh.Empty(item) {
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
func (fh *FunctionHandler) MustSlice(list any, indices ...any) (any, error) {
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
			start = fh.ToInt(indices[0])
		}
		if len(indices) < 2 {
			end = length
		} else {
			end = fh.ToInt(indices[1])
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
func (fh *FunctionHandler) MustHas(element any, list any) (bool, error) {
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
func (fh *FunctionHandler) MustWithout(list any, omit ...any) ([]any, error) {
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
			if !fh.InList(omit, item) {
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
func (fh *FunctionHandler) MustRest(list any) ([]any, error) {
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
func (fh *FunctionHandler) MustInitial(list any) ([]any, error) {
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
func (fh *FunctionHandler) MustFirst(list any) (any, error) {
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
func (fh *FunctionHandler) MustLast(list any) (any, error) {
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
func (fh *FunctionHandler) MustReverse(list any) ([]any, error) {
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
