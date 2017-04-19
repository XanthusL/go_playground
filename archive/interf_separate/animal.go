package interf_separate

import _ "lyh_demos/archive/interf_separate/inner"

// 只对外提供接口，具体实现在inner包，不导出，屏蔽实现细节
type Animal interface {
	Say() string
}

// 对外提供接口类型
var Pet Animal
