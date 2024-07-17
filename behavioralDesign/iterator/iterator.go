/*
 * @Author: bsun
 * @Date: 2024-07-02 09:25:50
 * @Last Modified by: bsun
 * @Last Modified time: 2024-07-02 09:31:24
 */
package iterator

type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{}
}

// 迭代器结构
type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

func (iter *ArrayIntIterator) HasNext() bool {
	return iter.index < len(iter.arrayInt)-1
}

func (iter *ArrayIntIterator) Next() {
	iter.index++
}

func (iter *ArrayIntIterator) CurrentItem() interface{} {
	return iter.arrayInt[iter.index]
}

// 数组结构
type ArrayInt []int

func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: a,
		index:    0,
	}
}
