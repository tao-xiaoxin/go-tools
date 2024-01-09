package slice

/*
实现切片的缩容
*/

// Shrink 切片缩容
func Shrink[T any](src []T) []T {
	// 获取原始切片的容量和长度
	c, l := cap(src), len(src)
	// 计算新的容量和判断是否发生了变化
	n, changed := calCapacity(c, l)
	// 如果没有变化，直接返回原始切片
	if !changed {
		return src
	}
	// 创建一个新的切片，容量为新计算的容量，长度为 0
	s := make([]T, 0, n)
	// 将原始切片的元素追加到新的切片中
	s = append(s, src...)
	return s
}

// calCapacity 计算新的容量和判断是否需要缩容
func calCapacity(c, l int) (int, bool) {
	// 如果容量 <= 64，则不进行缩容，因为浪费内存很少可以忽略不计
	if c <= 64 {
		return c, false
	}
	// 如果容量大于 2048，但是元素个数超过容量的一半，
	// 则将容量降低为 0.625 倍，也就是 5/8
	// 这与正向扩容的 1.25 倍相呼应
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	// 如果容量在 2048 以内，并且元素个数不足容量的 1/4，
	// 直接将容量缩减为原来的一半
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	// 整个实现的核心是希望在后续少触发扩容的前提下，一次性释放尽可能多的内存
	return c, false
}
