package reversearray

func ReverseArray(a []int32) []int32 {
	n := len(a)
	var v int32
	for i := 0; n/2 > i; i++ {
		v = a[i]
		a[i] = a[n-i-1]
		a[n-i-1] = v
		if n-i-1 < i+2 {
			return a
		}
	}
	return a

}

// 0
// 1 1
// 2 2
