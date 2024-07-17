package proxy

// IUser IUser
type IBase interface {
	Login(username, password string) error
	Verify() (int, error)
}

// Base 用户
// @proxy IBase
type Base struct {
}

// Login 用户登录
func (b *Base) Login(username, password string) error {
	// 不实现细节
	return nil
}

// Verify 验权
func (b *Base) Verify() (int, error) {
	return 0, nil
}
