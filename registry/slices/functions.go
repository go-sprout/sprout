package slices

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/spf13/cast"

	"github.com/go-sprout/sprout/internal/helpers"
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
// For an example of this function in a Go template, refer to [Sprout Documentation: list].
//
// [Sprout Documentation: list]: https://docs.atom.codes/sprout/registries/slices#list
func (sr *SlicesRegistry) List(values ...any) []any {
	return values
}

// Append appends an element to a slice or array, returning an error if the
// operation isn't applicable.
//
// Parameters:
//
//	v any - the element to append.
//	list any - the original list to append to.
//
// Returns:
//
//	[]any - the new list with the element appended.
//	error - error if the list is nil or not a slice/array.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: append].
//
// [Sprout Documentation: append]: https://docs.atom.codes/sprout/registries/slices#append
func (sr *SlicesRegistry) Append(v any, list any) ([]any, error) {
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
//	v any - the element to prepend.
//	list any - the original list to prepend to.
//
// Returns:
//
//	[]any - the new list with the element prepended.
//	error - error if the list is nil or not a slice/array.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: prepend].
//
// [Sprout Documentation: prepend]: https://docs.atom.codes/sprout/registries/slices#prepend
func (sr *SlicesRegistry) Prepend(v any, list any) ([]any, error) {
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
// For an example of this function in a Go template, refer to [Sprout Documentation: concat].
//
// [Sprout Documentation: concat]: https://docs.atom.codes/sprout/registries/slices#concat
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
// For an example of this function in a Go template, refer to [Sprout Documentation: chunk].
//
// [Sprout Documentation: chunk]: https://docs.atom.codes/sprout/registries/slices#chunk
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
// For an example of this function in a Go template, refer to [Sprout Documentation: uniq].
//
// [Sprout Documentation: uniq]: https://docs.atom.codes/sprout/registries/slices#uniq
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
// For an example of this function in a Go template, refer to [Sprout Documentation: compact].
//
// [Sprout Documentation: compact]: https://docs.atom.codes/sprout/registries/slices#compact
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

// Flatten flattens a nested list into a single list of elements.
//
// Parameters:
//
//	list any - the list to flatten.
//
// Returns:
//
//	[]any - the flattened list.
//	error - error if the list is nil or not a slice/array.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: flatten].
//
// [Sprout Documentation: flatten]: https://docs.atom.codes/sprout/registries/slices#flatten
func (sr *SlicesRegistry) Flatten(list any) ([]any, error) {
	return sr.FlattenDepth(-1, list)
}

// FlattenDepth flattens a nested list into a single list of elements up to a
// specified depth.
//
// Parameters:
//
//	deep int - the maximum depth to flatten to; -1 for infinite depth.
//	list any - the list to flatten.
//
// Returns:
//
//	[]any - the flattened list.
//	error - error if the list is nil or not a slice/array.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: flattenDepth].
//
// [Sprout Documentation: flattenDepth]: https://docs.atom.codes/sprout/registries/slices#flattendepth
func (sr *SlicesRegistry) FlattenDepth(deep int, list any) ([]any, error) {
	if list == nil {
		return nil, fmt.Errorf("cannot flatten nil")
	}

	valueOfList := reflect.ValueOf(list)
	tp := valueOfList.Kind()

	switch tp {
	case reflect.Slice, reflect.Array:
		return sr.flattenSlice(valueOfList, deep), nil
	default:
		return nil, fmt.Errorf("cannot flatten on type %s", tp)
	}
}

// Slice extracts a slice from a list between two indices.
//
// Parameters:
//
//	start int - the start index (inclusive).
//	end int - optional end index (exclusive); if omitted, slices to the end.
//	list any - the list to slice.
//
// Returns:
//
//	any - the sliced part of the list.
//	error - error if the list is nil or not a slice/array.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: slice].
//
// [Sprout Documentation: slice]: https://docs.atom.codes/sprout/registries/slices#slice
func (sr *SlicesRegistry) Slice(args ...any) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("slice requires at least one argument")
	}

	list := args[len(args)-1]
	indices := args[:len(args)-1]

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
		return nil, fmt.Errorf("last argument must be a slice but got %T", list)
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
// For an example of this function in a Go template, refer to [Sprout Documentation: has].
//
// [Sprout Documentation: has]: https://docs.atom.codes/sprout/registries/slices#has
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
//	omit ...any - elements to exclude from the new list.
//	list any - the original list (last argument).
//
// Returns:
//
//	[]any - the list excluding the specified elements.
//	error - error if the list is nil or not a slice/array.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: without].
//
// [Sprout Documentation: without]: https://docs.atom.codes/sprout/registries/slices#without
func (sr *SlicesRegistry) Without(args ...any) ([]any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("without requires at least two arguments")
	}

	list := args[len(args)-1]
	omit := args[:len(args)-1]

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
		return nil, fmt.Errorf("last argument must be a slice but got %T", list)
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
// For an example of this function in a Go template, refer to [Sprout Documentation: rest].
//
// [Sprout Documentation: rest]: https://docs.atom.codes/sprout/registries/slices#rest
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
// For an example of this function in a Go template, refer to [Sprout Documentation: initial].
//
// [Sprout Documentation: initial]: https://docs.atom.codes/sprout/registries/slices#initial
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
// For an example of this function in a Go template, refer to [Sprout Documentation: first].
//
// [Sprout Documentation: first]: https://docs.atom.codes/sprout/registries/slices#first
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
// For an example of this function in a Go template, refer to [Sprout Documentation: last].
//
// [Sprout Documentation: last]: https://docs.atom.codes/sprout/registries/slices#last
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
// For an example of this function in a Go template, refer to [Sprout Documentation: reverse].
//
// [Sprout Documentation: reverse]: https://docs.atom.codes/sprout/registries/slices#reverse
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
// For an example of this function in a Go template, refer to [Sprout Documentation: sortAlpha].
//
// [Sprout Documentation: sortAlpha]: https://docs.atom.codes/sprout/registries/slices#sortalpha
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
//	value string - the string to split.
//
// Returns:
//
//	[]string - a slice containing the substrings obtained from splitting the input string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: splitList].
//
// [Sprout Documentation: splitList]: https://docs.atom.codes/sprout/registries/slices#splitlist
func (sr *SlicesRegistry) SplitList(sep string, value string) []string {
	return strings.Split(value, sep)
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
// For an example of this function in a Go template, refer to [Sprout Documentation: strSlice].
//
// [Sprout Documentation: strSlice]: https://docs.atom.codes/sprout/registries/slices#strslice
func (sr *SlicesRegistry) StrSlice(value any) []string {
	return helpers.StrSlice(value)
}

// Until generates a slice of integers from 0 up to but not including 'count'.
// If 'count' is negative, it produces a descending slice from 0 down to 'count',
// inclusive, with a step of -1. The function leverages UntilStep to specify
// the range and step dynamically.
//
// Parameters:
//
//	count int - the endpoint (exclusive) of the range to generate.
//
// Returns:
//
//	[]int - a slice of integers from 0 to 'count' with the appropriate step
//	        depending on whether 'count' is positive or negative.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: until].
//
// [Sprout Documentation: until]: https://docs.atom.codes/sprout/registries/slices#until
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
// For an example of this function in a Go template, refer to [Sprout Documentation: untilStep].
//
// [Sprout Documentation: untilStep]: https://docs.atom.codes/sprout/registries/slices#untilstep
func (sr *SlicesRegistry) UntilStep(start, stop, step int) []int {
	return helpers.UntilStep(start, stop, step)
}
