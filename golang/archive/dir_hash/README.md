[![Go Report Card](https://goreportcard.com/badge/github.com/XanthusL/dir_hash)](https://goreportcard.com/report/github.com/XanthusL/dir_hash)
# ~~dir_hash~~
>desperated

### 
- 对整个目录下的所有文件进行遍历，获取所有文件的大小和计算文件的sha1哈希值 `done`
- 每一行一个文件，包含文件名称，哈希值，文件大小。结果格式：a.用逗号隔开(csv)，b.json `done`
- 代码实现简洁，能够最大限度利用上多核性能 `done`
- path参数, out参数，help参数，filter参数，format参数  `done`
- 可以指定忽略哪些目录、文件，需要支持通配符(正则) `done`
- 通过测试代码自我证明代码能够可靠运行并正确实现上述功能 `todo`
- 安全稳定可靠的，可读性强的代码 `todo`

---
#### -path
指定扫描的目录，默认为当前目录
#### -out
结果输出到指定文件，不指定则是标准输出
#### -format
指定输出格式，支持`line`和`json`，默认为`line`
#### -help
打印帮助信息

    $ ./dir_hash -help
    Usage of ./dir_hash:
      -filter string
          Regex of files to skip
      -format string
    	  Out put format. "line" and "json" supported.(default "line")
      -help
          Show this usage.
      -out string
          The file to save to. Prints to standard if not assigned.
      -path string
          The path to scan. Current path as default. (default ".")
          
#### -filter
通过正则指定需要跳过的文件
 
     $ ./dir_hash -filter \.git/?
    .idea/misc.xml,2e6d84f0e55f4186311c8b9493b155c1adc6dcf6,2713
    app.go,0043351289963f0e91a2e852a5fb6fa9d959bc75,2986
    .idea/modules.xml,757dd608b837c81618012046c95c9fbb99af405b,268
    .idea/workspace.xml,0fe0fd523f6d7d8a0a4d308852cbaac65e44646d,31735
    dir_hash,614d69f821ecb7647a203b5cab40025153e0b7fa,2299296
    .idea/dir_hash.iml,6084689ee61da613a5ad61c22f4c5d5e43b1c6f8,415
    .idea/libraries/GOPATH__dir_hash_.xml,9ff0b653d5b85c07336c2751bbf1a441d5bf2f38,1557
    .idea/vcs.xml,6f94fc1df9e8721673d47588ac444667dc9ded06,167
    .idea/inspectionProfiles/Project_Default.xml,ba839e093a4251d20e3266c0745142727b829f9d,155
    README.md,7f64baa52b95becda2d841adc5d85663f92efd6a,614

