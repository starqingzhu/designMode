package prototype

import (
	"encoding/json"
	"time"
)

/*
原型设计模式
定义：
	利用已有对象（原型）进行复制（拷贝）的方式来创建新对象，已达到节省创建时间的目的

使用场景：
	1.对象创建成本比较大，并且同一个类的不同对象之间差别不大。（大部分字段相同）
	2.对象的创建成本比较大
		1）对象数据需要经过复杂的计算，排序，hash等
		2）需要从rpc， 网络， 数据库 等非常慢的io中获取

实现方式：
	1.深拷贝
		完全复制，复制出的对象和原本的对象没关系
		实现方式
			1）递归复制
			2）序列化与反序列化
	2.浅拷贝
		1）仅复制对象的引用，不进行递归复制
		2）如果字段是一个地址，那么复制的对象，江河源对象共享相同的数据
		3）如果字段对象的是可变对象，那么复制的对象改动会导致原来对象的改动
		4）如果字段对象的是不可变对象，则没有什么问题
*/

// Keyword 搜索关键字
type KeyWord struct {
	Word      string
	Visit     int
	UpdatedAt *time.Time
}

// Clone 这里使用序列化与反序列化的方式深拷贝
func (k *KeyWord) Clone() *KeyWord {
	var newKeyWord KeyWord

	b, _ := json.Marshal(k)

	json.Unmarshal(b, &newKeyWord)
	return &newKeyWord
}

type KeyWords map[string]*KeyWord

// Clone 复制一个新的 keywords
// updatedWords: 需要更新的关键词列表，由于从数据库中获取数据常常是数组的方式

func (words KeyWords) Clone(updatedWords []*KeyWord) KeyWords {
	newKeyWords := KeyWords{}

	for k, v := range words {
		// 这里是浅拷贝，直接拷贝了地址
		newKeyWords[k] = v
	}
	// 替换掉需要更新的字段，这里用的是深拷贝
	for _, word := range updatedWords {
		newKeyWords[word.Word] = word.Clone()
	}

	return newKeyWords
}
