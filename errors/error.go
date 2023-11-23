package errors

import "fmt"

// NewErrIndexOutOfRange 创建一个代表下标超出范围的错误
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("go-tools: 下标超出范围，长度 %d, 下标 %d", length, index)
}

// NewErrValueNotFound 创建一个代表值不存在的错误
func NewErrValueNotFound[T any](args []T, value T) error {
	return fmt.Errorf("go-tools: %v not found in %v", value, args)
}

// NewErrSliceIsEmpty 创建一个代表切片为空的错误
func NewErrSliceIsEmpty() error {
	return fmt.Errorf("go-tools: slice is empty")
}
