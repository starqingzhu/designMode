/*
 * @Author: bsun
 * @Date: 2024-07-02 09:32:46
 * @Last Modified by: bsun
 * @Last Modified time: 2024-07-02 11:38:12
 */

package test

import (
	"designmode/behavioralDesign/iterator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayInt_Iterator(t *testing.T) {
	data := iterator.ArrayInt{1, 3, 5, 7, 8}
	iter := data.Iterator()
	// i 用于测试
	i := 0
	for iter.HasNext() {
		assert.Equal(t, data[i], iter.CurrentItem())
		iter.Next()
		i++
	}
}
