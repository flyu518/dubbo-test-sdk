# 项目微服务接口定义

## 微服务模块化处理，以 user 服务为例
- 每个微服务内部需要有一个 go.mod，里面是完整地址+模块，如：module github.com/flyu518/dubbo-test-sdk/user
- user.proto 里面的 option go_package 也需要对应，如：option go_package = "github.com/flyu518/dubbo-test-sdk/user/api;api";
- 每个微服务接口定义发生修改，需要执行类似：git tag user/v1.0.0 进行模块版本管理
- 使用示例：github.com/flyu518/dubbo-test-sdk/user@v1.0.0
