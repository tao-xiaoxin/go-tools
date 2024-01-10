package slice

import (
	"github.com/tao-xiaoxin/go-tools/errors"
)

/*
Delete 按照索引删除单个指定元素
*/
func Delete[T any](src []T, index int) ([]T, error) {
	length := len(src)
	if (index < 0) || (index >= length) {
		return nil, errors.NewErrIndexOutOfRange(length, index)
	}
	//从index位置开始，后面的元素依次往前挪1个位置
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	//去掉最后一个重复元素
	src = src[:length-1]
	return src, nil
}
