
USER_PROTO_DIR=./user/api
ORDER_PROTO_DIR=./order/api
# 安装基础插件
.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 安装 dubbo-go 插件
.PHONY: install-protoc-triple
install-protoc-triple:
	@go install github.com/dubbogo/protoc-gen-go-triple/v3@latest

# 执行所有 proto 生成任务
.PHONY: proto
proto: proto-user proto-order

# 生成 user-service 的 proto 文件
.PHONY: proto-user
proto-user:
	@protoc --go_out=./ --go_opt=paths=source_relative \
	--go-triple_out=./ --go-triple_opt=paths=source_relative \
	${USER_PROTO_DIR}/*.proto

# 生成 order-service 的 proto 文件
.PHONY: proto-order
proto-order:
	@protoc --go_out=./ --go_opt=paths=source_relative \
	--go-triple_out=./ --go-triple_opt=paths=source_relative \
	${ORDER_PROTO_DIR}/*.proto