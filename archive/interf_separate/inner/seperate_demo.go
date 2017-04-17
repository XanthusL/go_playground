package inner

import "lyh_demos/archive/interf_separate"

// 不导出，屏蔽实现细节
type dog struct {
}

// 实现接口
func (_ *dog) Say() string {
	return "Woof!"
}

func init() {
	interf_separate.Pet = &dog{}
}
