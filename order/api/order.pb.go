// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: order/api/order.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	OrderName     string                 `protobuf:"bytes,2,opt,name=orderName,proto3" json:"orderName,omitempty"`
	OrderPrice    string                 `protobuf:"bytes,3,opt,name=orderPrice,proto3" json:"orderPrice,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_order_api_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_api_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetOrderName() string {
	if x != nil {
		return x.OrderName
	}
	return ""
}

func (x *Order) GetOrderPrice() string {
	if x != nil {
		return x.OrderPrice
	}
	return ""
}

type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_order_api_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_api_order_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	mi := &file_order_api_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_order_api_order_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_order_api_order_proto protoreflect.FileDescriptor

const file_order_api_order_proto_rawDesc = "" +
	"\n" +
	"\x15order/api/order.proto\x12\x1clingmou.simple.sdk.order.api\"_\n" +
	"\x05Order\x12\x18\n" +
	"\aorderId\x18\x01 \x01(\tR\aorderId\x12\x1c\n" +
	"\torderName\x18\x02 \x01(\tR\torderName\x12\x1e\n" +
	"\n" +
	"orderPrice\x18\x03 \x01(\tR\n" +
	"orderPrice\"+\n" +
	"\x0fGetOrderRequest\x12\x18\n" +
	"\aorderId\x18\x01 \x01(\tR\aorderId\"M\n" +
	"\x10GetOrderResponse\x129\n" +
	"\x05order\x18\x01 \x01(\v2#.lingmou.simple.sdk.order.api.OrderR\x05order2{\n" +
	"\fOrderService\x12k\n" +
	"\bGetOrder\x12-.lingmou.simple.sdk.order.api.GetOrderRequest\x1a..lingmou.simple.sdk.order.api.GetOrderResponse\"\x00B1Z/github.com/flyu518/dubbo-test-sdk/order/api;apib\x06proto3"

var (
	file_order_api_order_proto_rawDescOnce sync.Once
	file_order_api_order_proto_rawDescData []byte
)

func file_order_api_order_proto_rawDescGZIP() []byte {
	file_order_api_order_proto_rawDescOnce.Do(func() {
		file_order_api_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_api_order_proto_rawDesc), len(file_order_api_order_proto_rawDesc)))
	})
	return file_order_api_order_proto_rawDescData
}

var file_order_api_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_order_api_order_proto_goTypes = []any{
	(*Order)(nil),            // 0: lingmou.simple.sdk.order.api.Order
	(*GetOrderRequest)(nil),  // 1: lingmou.simple.sdk.order.api.GetOrderRequest
	(*GetOrderResponse)(nil), // 2: lingmou.simple.sdk.order.api.GetOrderResponse
}
var file_order_api_order_proto_depIdxs = []int32{
	0, // 0: lingmou.simple.sdk.order.api.GetOrderResponse.order:type_name -> lingmou.simple.sdk.order.api.Order
	1, // 1: lingmou.simple.sdk.order.api.OrderService.GetOrder:input_type -> lingmou.simple.sdk.order.api.GetOrderRequest
	2, // 2: lingmou.simple.sdk.order.api.OrderService.GetOrder:output_type -> lingmou.simple.sdk.order.api.GetOrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_api_order_proto_init() }
func file_order_api_order_proto_init() {
	if File_order_api_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_api_order_proto_rawDesc), len(file_order_api_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_api_order_proto_goTypes,
		DependencyIndexes: file_order_api_order_proto_depIdxs,
		MessageInfos:      file_order_api_order_proto_msgTypes,
	}.Build()
	File_order_api_order_proto = out.File
	file_order_api_order_proto_goTypes = nil
	file_order_api_order_proto_depIdxs = nil
}
