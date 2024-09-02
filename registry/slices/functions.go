package slices

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/go-sprout/sprout/deprecated"
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

// Append appends an element to a slice or array, returning an error if the
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
//	error - protect against undesired behavior due to migration to new signature.
//
// Example:
//
//	{{ ["a", "b"] | append "c"  }} // Output: ["a", "b", "c"], nil
func (sr *SlicesRegistry) Append(args ...any) ([]any, error) {
	//! BACKWARDS COMPATIBILITY: deprecated in v1.0 and removed in v1.1
	//! Due to change in signature, this function still supports the old signature
	//! to let users transition to the new signature.
	//* Old signature: Append(list any, v any)
	//* New signature: Append(v any, list any)
	if len(args) != 2 {
		return []any{}, deprecated.ErrArgsCount(2, len(args))
	}

	switch reflect.ValueOf(args[0]).Kind() {
	case reflect.Array, reflect.Slice, reflect.Invalid:
		// Old signature
		deprecated.SignatureWarn(sr.handler.Logger(), "append", "{{ append list v }}", "{{ list | append v }}")
		return sr.Append(append(args[1:], args[0])...)
	}

	// New signature
	list := args[1]
	v := args[0]
	// ! END OF BACKWARDS COMPATIBILITY

	if list == nil {
		return nil, fmt.Errorf("cannot append to nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice:
		// If the list is already a slice, simply append the value
		result := make([]any, valueOfList.Len()+1)
		for i := 0; i < valueOfList.Len(); i++ {
			result[i] = valueOfList.Index(i).Interface()
		}
		result[len(result)-1] = v
		return result, nil

	case reflect.Array:
		// For arrays, we need to convert to a slice first
		length := valueOfList.Len()
		result := make([]any, length+1)
		for i := 0; i < length; i++ {
			result[i] = valueOfList.Index(i).Interface()
		}
		result[length] = v
		return result, nil

	default:
		return nil, fmt.Errorf("cannot append on type %s", tp)
	}
}

// Prepend prepends an element to a slice or array, returning an error if
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
//	error - protect against undesired behavior due to migration to new signature.
//
// Example:
//
//	{{ ["b", "c"] | prepend "a" }} // Output: ["a", "b", "c"], nil
func (sr *SlicesRegistry) Prepend(args ...any) ([]any, error) {
	//! BACKWARDS COMPATIBILITY: deprecated in v1.0 and removed in v1.1
	//! Due to change in signature, this function still supports the old signature
	//! to let users transition to the new signature.
	//* Old signature: Prepend(list any, v any)
	//* New signature: Prepend(v any, list any)
	if len(args) != 2 {
		return []any{}, deprecated.ErrArgsCount(2, len(args))
	}

	switch reflect.ValueOf(args[0]).Kind() {
	case reflect.Array, reflect.Slice, reflect.Invalid:
		// Old signature
		deprecated.SignatureWarn(sr.handler.Logger(), "prepend", "{{ prepend list v }}", "{{ list | prepend v }}")
		return sr.Prepend(append(args[1:], args[0])...)
	}

	// New signature
	list := args[1]
	v := args[0]
	// ! END OF BACKWARDS COMPATIBILITY

	if list == nil {
		return nil, fmt.Errorf("cannot prepend to nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
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
	// Estimate the total length to preallocate the result slice
	var totalLen int
	for _, list := range lists {
		if list == nil {
			continue
		}

		tp := reflect.TypeOf(list).Kind()
		if tp == reflect.Slice || tp == reflect.Array {
			totalLen += reflect.ValueOf(list).Len()
		}
	}

	// Preallocate the result slice
	res := make([]any, 0, totalLen)

	for _, list := range lists {
		if list == nil {
			continue
		}

		tp := reflect.TypeOf(list).Kind()
		if tp == reflect.Slice || tp == reflect.Array {
			valueOfList := reflect.ValueOf(list)
			for i := 0; i < valueOfList.Len(); i++ {
				res = append(res, valueOfList.Index(i).Interface())
			}
		}
	}

	return res
}

// Chunk divides a list into chunks of specified size, returning an error
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
//	{{ ["a", "b", "c", "d"] | chunk 2 }} // Output: [["a", "b"], ["c", "d"]], nil
func (sr *SlicesRegistry) Chunk(size int, list any) ([][]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot chunk nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()
	switch tp {
	case reflect.Slice, reflect.Array:
		length := valueOfList.Len()

		chunkCount := (length + size - 1) / size
		result := make([][]any, chunkCount)

		for i := 0; i < chunkCount; i++ {
			start := i * size
			end := start + size

			if end > length {
				end = length
			}

			chunkLength := end - start
			result[i] = make([]any, chunkLength)

			for j := 0; j < chunkLength; j++ {
				result[i][j] = valueOfList.Index(start + j).Interface()
			}
		}

		return result, nil

	default:
		return nil, fmt.Errorf("cannot chunk type %s", tp)
	}
}

// Uniq returns a new slice containing unique elements of the given list,
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
//	{{ ["a", "b", "a", "c"] | uniq }} // Output: ["a", "b", "c"], nil
func (sr *SlicesRegistry) Uniq(list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot uniq nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
		length := valueOfList.Len()
		result := make([]any, 0, length)
		// This allows for O(1) average-time complexity checks to see
		// if an item has already been encountered, which is much faster than
		// scanning a slice (O(n) time complexity).
		seen := make(map[any]bool, length)

		for i := 0; i < length; i++ {
			item := valueOfList.Index(i).Interface()
			if !seen[item] {
				seen[item] = true
				result = append(result, item)
			}
		}

		return result, nil
	default:
		return nil, fmt.Errorf("cannot find uniq on type %s", tp)
	}
}

// Compact removes nil or zero-value elements from a list.
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
//	{{ [0, 1, nil, 2, "", 3] | compact }} // Output: [1, 2, 3], nil
func (sr *SlicesRegistry) Compact(list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot compact nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
		length := valueOfList.Len()
		result := make([]any, 0, length)

		for i := 0; i < length; i++ {
			item := valueOfList.Index(i).Interface()
			if !helpers.Empty(item) {
				result = append(result, item)
			}
		}

		return result, nil
	default:
		return nil, fmt.Errorf("cannot compact on type %s", tp)
	}
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
//	error - protect against undesired behavior due to migration to new signature.
//
// Example:
//
//	{{ [1, 2, 3, 4, 5] | slice 1, 3 }} // Output: [2, 3], nil
func (sr *SlicesRegistry) Slice(args ...any) (any, error) {
	//! BACKWARDS COMPATIBILITY: deprecated in v1.0 and removed in v1.1
	//! Due to change in signature, this function still supports the old signature
	//! to let users transition to the new signature.
	//* Old signature: Slice(list any, indices ...any)
	//* New signature: Slice(indices ...any, list any)
	if len(args) < 1 {
		return []any{}, deprecated.ErrArgsCount(2, len(args))
	}

	if len(args) == 1 {
		return args[0], nil
	}

	switch reflect.ValueOf(args[0]).Kind() {
	case reflect.Array, reflect.Slice, reflect.Invalid:
		// Old signature
		deprecated.SignatureWarn(sr.handler.Logger(), "slice", "{{ slice list 1 2 }}", "{{ list | slice 1 2 }}")
		return sr.Slice(append(args[1:], args[0])...)
	}

	// New signature
	list := args[len(args)-1]
	indices := args[:len(args)-1]
	// ! END OF BACKWARDS COMPATIBILITY

	if list == nil {
		return nil, fmt.Errorf("cannot slice nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
		length := valueOfList.Len()
		if length == 0 {
			return nil, nil
		}

		start, end := 0, length

		// Handle start index
		if len(indices) > 0 {
			start = cast.ToInt(indices[0])
			if start < 0 || start > length {
				return nil, fmt.Errorf("start index out of bounds")
			}
		}

		// Handle end index
		if len(indices) > 1 {
			end = cast.ToInt(indices[1])
			if end < start || end > length {
				return nil, fmt.Errorf("end index out of bounds")
			}
		}

		return valueOfList.Slice(start, end).Interface(), nil
	default:
		return nil, fmt.Errorf("last argument must be a slice but got %T", args[len(args)-1])
	}
}

// Has checks if a specified element is present in a collection and handles
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
//	{{ [1, 2, 3, 4] | has 3 }} // Output: true, nil
func (sr *SlicesRegistry) Has(element any, list any) (bool, error) {
	if list == nil {
		return false, nil
	}
	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
		length := valueOfList.Len()

		for i := 0; i < length; i++ {
			item := valueOfList.Index(i).Interface()
			if reflect.DeepEqual(element, item) {
				return true, nil
			}
		}

		return false, nil
	default:
		return false, fmt.Errorf("cannot find has on type %s", tp)
	}
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
//	error - protect against undesired behavior due to migration to new signature.
//
// Example:
//
//	{{ [1, 2, 3, 4] | without 2, 4 }} // Output: [1, 3], nil
func (sr *SlicesRegistry) Without(args ...any) ([]any, error) {
	//! BACKWARDS COMPATIBILITY: deprecated in v1.0 and removed in v1.1
	//! Due to change in signature, this function still supports the old signature
	//! to let users transition to the new signature.
	//* Old signature: Without(list any, omit ...any)
	//* New signature: Without(omit ...any, list any)
	if len(args) < 2 {
		return []any{}, deprecated.ErrArgsCount(2, len(args))
	}

	switch reflect.ValueOf(args[0]).Kind() {
	case reflect.Array, reflect.Slice, reflect.Invalid:
		// Old signature
		deprecated.SignatureWarn(sr.handler.Logger(), "without", "{{ without list 1 2 }}", "{{ list | without 1 2 }}")
		return sr.Without(append(args[1:], args[0])...)
	}

	// New signature
	list := args[len(args)-1]
	omit := args[:len(args)-1]
	// ! END OF BACKWARDS COMPATIBILITY

	if list == nil {
		return nil, fmt.Errorf("cannot without nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
		length := valueOfList.Len()
		omitSet := make(map[any]struct{}, len(omit))

		// Populate the set of items to omit
		for _, o := range omit {
			omitSet[o] = struct{}{}
		}

		result := make([]any, 0, length)

		for i := 0; i < length; i++ {
			item := valueOfList.Index(i).Interface()
			if _, found := omitSet[item]; !found {
				result = append(result, item)
			}
		}

		return result, nil
	default:
		return nil, fmt.Errorf("last argument must be a slice but got %T", args[len(args)-1])
	}
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
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ [1, 2, 3, 4] | rest }} // Output: [2, 3, 4], nil
func (sr *SlicesRegistry) Rest(list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot rest nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
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

// Initial returns all elements of a list except the last.
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
//	{{ [1, 2, 3, 4] | initial }} // Output: [1, 2, 3], nil
func (sr *SlicesRegistry) Initial(list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot initial nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
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

// First returns the first element of a list.
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
//	{{ [1, 2, 3, 4] | first }} // Output: 1, nil
func (sr *SlicesRegistry) First(list any) (any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot first nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
		length := valueOfList.Len()
		if length == 0 {
			return nil, nil
		}

		return valueOfList.Index(0).Interface(), nil
	default:
		return nil, fmt.Errorf("cannot find first on type %s", tp)
	}
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
//	error - error if the list is nil, empty, or not a slice/array.
//
// Example:
//
//	{{ [1, 2, 3, 4] | last }} // Output: 4, nil
func (sr *SlicesRegistry) Last(list any) (any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot last nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
		length := valueOfList.Len()
		if length == 0 {
			return nil, nil
		}

		return valueOfList.Index(length - 1).Interface(), nil
	default:
		return nil, fmt.Errorf("cannot find last on type %s", tp)
	}
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
//	error - error if the list is nil or not a slice/array.
//
// Example:
//
//	{{ [1, 2, 3, 4] | reverse }} // Output: [4, 3, 2, 1], nil
func (sr *SlicesRegistry) Reverse(list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot reverse nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
		length := valueOfList.Len()

		// Create a new slice with the same length as the original
		nl := make([]any, length)
		for i := 0; i < length; i++ {
			nl[i] = valueOfList.Index(length - i - 1).Interface()
		}

		return nl, nil
	default:
		return nil, fmt.Errorf("cannot find reverse on type %s", tp)
	}
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
		strList := sr.StrSlice(list)
		sort.Strings(strList)
		return strList
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
