package slice

import (
	"github.com/tao-xiaoxin/go-tools/errors"
)

/*
按照索引删除单个指定元素
*/
func Delete[T any](slice []T, index int) ([]T, T, error) {
	length := len(slice)
	if (index < 0) || (index >= length) {
		var zero T
		return nil, zero, errors.NewErrIndexOutOfRange(length, index)
	}
	ele := slice[index]
	return append(slice[:index], slice[index+1:]...), ele, nil
}
