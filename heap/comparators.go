package heap

// MinHeapInt64 can be used to create a min heap handling int64 datatype
// The function will panic for other types if casting fails
func MinHeapInt64(a, b interface{}) bool {
	return a.(int64) < b.(int64)
}

// MaxHeapInt64 can be used to create a min heap handling int64 datatype
// The function will panic for other types if casting fails
func MaxHeapInt64(a, b interface{}) bool {
	return a.(int64) > b.(int64)
}

var (
	_ AssertHeapProperty = MinHeapInt64
	_ AssertHeapProperty = MaxHeapInt64
)
