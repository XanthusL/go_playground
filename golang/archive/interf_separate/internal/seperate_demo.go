package internal

// 不导出(internal)，屏蔽实现细节
type Dog struct {
}

// 实现接口
func (_ *Dog) Say() string {
	return "Woof!"
}
