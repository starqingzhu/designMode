package memento

/*
如果输入：list 则显示当前保存的内容
如果输入：undo 则删除上一次的输入
如果输入其他内容则追加保存
*/

// InputText 用于保存数据
type InputText struct {
	content string
}

// 追加数据
func (in *InputText) Append(content string) {
	in.content += content
}

// 获取数据
func (in *InputText) GetText() string {
	return in.content
}

// 创建快照
func (in *InputText) Snapshot() *Snapshot {
	return &Snapshot{content: in.content}
}

// 恢复快照
func (in *InputText) Restore(s *Snapshot) {
	in.content = s.GetText()
}

// 快照，用于存储数据快照
type Snapshot struct {
	content string
}

func (s *Snapshot) GetText() string {
	return s.content
}
