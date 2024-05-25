package mergesort

// You can probably just use the standard sort libraries to do things better, but
// I just love the simplicity of setting up a merge sort with whatever you want
// to sort a list by. It's recursive, and that's all I need to hear when using it.
// Recursive functions are OP.

func MergeSort(u []int) []int {
	// You can substitute an integer for whatever you want to search by.
	// I'm using integers here for simplicity's sake.
	if len(u) == 1 {
		return u // If we're ever just one item, just return it.
	} else {
		// first, compute what half the size of the provided array is.
		// This will be a whole number, no remainders.
		half := len(u) / 2

		// now let's create two other arrays denoting the left
		// and right side of the array.
		left := u[:half]
		right := u[half:]

		// Now let's call the function again, remember this will
		// keep recursively repeating until it hits just one item.
		lsort := MergeSort(left)
		rsort := MergeSort(right)

		// Finally, let the helper function do the heavy lifting of
		// deciding the order of the items.
		return merge(lsort, rsort)
	}
}

// Note lowercase "merge", this is because this function should be
// considered internal. No need to call this from anywhere else.
func merge(left, right []int) []int {
	// create a blank return value
	var retval []int

	for {
		// We're going to loop until we break on a condition.

		if len(left) == 0 || len(right) == 0 {
			break
		}

		// this is where we decide _how_ this is going to be
		// ordered. For this, I'm just going go by ascending
		// order, so if left < right, push left onto the retval
		if left[0] < right[0] {
			retval = append(retval, left[0])
			// don't forget to pop off the left value
			left = left[1:]
		} else {
			// otherwise, right[0] is greater than left[0], or
			// they're equal in which case it doesn't matter anyway.
			retval = append(retval, right[0])
			right = right[1:]
		}
	}

	// Now we broke because either the left or right side were depleted.
	// We can just append the remainder to the rest of the retval.
	if len(left) > 0 {
		retval = append(retval, left...)
	}

	if len(right) > 0 {
		retval = append(retval, right...)
	}

	// remember that elipses in Go is the equivalent of writing out each
	// element in the array separated by commas. In this case, it's just
	// going to dump the rest of the values onto the rest of the return
	// value.
	return retval
}
