package slice

import (
	"fmt"
	"github.com/tao-xiaoxin/go-tools/errors"
	"strings"
)

// Join 将切片通过指定的字符串分隔符连接起来
func Join[T any](slice []T, separator string) (string, error) {
	length := len(slice)
	if length <= 0 {
		return "", errors.NewErrSliceIsEmpty()
	}
	var strSlice []string
	strSlice = make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = fmt.Sprintf("%v", v)
	}

	joinedString := strings.Join(strSlice, separator)
	return joinedString, nil
}
