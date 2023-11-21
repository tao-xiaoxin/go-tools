package slice

import "github.com/tao-xiaoxin/go-tools/errors"

// Remove 按照切片中的某个值删除
func Remove[T comparable](slice []T, value T) ([]T, error) {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...), nil
		}
	}
	return nil, errors.NewErrValueNotFound(slice, value)
}
