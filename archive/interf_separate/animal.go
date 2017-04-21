package interf_separate

import "lyh_demos/archive/interf_separate/internal"

// 只对外提供接口，具体实现在inner包，不导出，屏蔽实现细节
type Animal interface {
	Say() string
}

func init() {
	Pet = &internal.Dog{}
}

// 对外提供接口类型
var Pet Animal
