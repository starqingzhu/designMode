package builder

import "fmt"

/*
建造者

应用场景：
	1.类中属性比较多
	2.类的属性之间有一定依赖关系，或者是约束条件
	3.存在必选和非必选的属性
	4.希望创建不可变的对象

和工厂模式的区别：
	工厂模式：创建类型相关的不同对象
	建造者模式： 用于创建参数复杂的对象
*/

const (
	DefaultMaxTotal = 10
	DefaultMaxIdle  = 9
	DefaultMinIdle  = 1
)

type (
	ResourcePoolConfig struct {
		Name     string
		MaxTotal int
		MaxIdle  int
		MinIdle  int
	}

	ResourcePoolConfigBuilder struct {
		Name     string
		MaxTotal int
		MaxIdle  int
		MinIdle  int
	}
)

// SetName SetName
func (b *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("Name can not be empty")
	}
	b.Name = name
	return nil
}

// SetMaxIdle SetMaxIdle
func (b *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("max tatal cannot < 0, input: %d", maxIdle)
	}
	b.MaxIdle = maxIdle
	return nil
}

// SetMaxTotal SetMaxTotal
func (b *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max tatal cannot <= 0, input: %d", maxTotal)
	}
	b.MaxTotal = maxTotal
	return nil
}

// Build Build
func (b *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if b.Name == "" {
		return nil, fmt.Errorf("Name can not be empty")
	}

	// 设置默认值
	if b.MinIdle == 0 {
		b.MinIdle = DefaultMinIdle
	}

	if b.MaxIdle == 0 {
		b.MaxIdle = DefaultMaxIdle
	}

	if b.MaxTotal == 0 {
		b.MaxTotal = DefaultMaxTotal
	}

	if b.MaxTotal < b.MaxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", b.MaxTotal, b.MaxIdle)
	}

	if b.MinIdle > b.MaxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", b.MaxIdle, b.MinIdle)
	}

	return &ResourcePoolConfig{
		Name:     b.Name,
		MaxTotal: b.MaxTotal,
		MaxIdle:  b.MaxIdle,
		MinIdle:  b.MinIdle,
	}, nil
}
