// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: v1/service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HostProcessReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *ProcessReport `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *HostProcessReportRequest) Reset() {
	*x = HostProcessReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostProcessReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostProcessReportRequest) ProtoMessage() {}

func (x *HostProcessReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostProcessReportRequest.ProtoReflect.Descriptor instead.
func (*HostProcessReportRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *HostProcessReportRequest) GetReport() *ProcessReport {
	if x != nil {
		return x.Report
	}
	return nil
}

type ActionKillProcess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PID    int32  `protobuf:"varint,1,opt,name=PID,proto3" json:"PID,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *ActionKillProcess) Reset() {
	*x = ActionKillProcess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionKillProcess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionKillProcess) ProtoMessage() {}

func (x *ActionKillProcess) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionKillProcess.ProtoReflect.Descriptor instead.
func (*ActionKillProcess) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ActionKillProcess) GetPID() int32 {
	if x != nil {
		return x.PID
	}
	return 0
}

func (x *ActionKillProcess) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ActionLogoutUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *ActionLogoutUser) Reset() {
	*x = ActionLogoutUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionLogoutUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionLogoutUser) ProtoMessage() {}

func (x *ActionLogoutUser) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionLogoutUser.ProtoReflect.Descriptor instead.
func (*ActionLogoutUser) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *ActionLogoutUser) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *ActionLogoutUser) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type HostActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processes []*ActionKillProcess `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
	Users     []*ActionLogoutUser  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *HostActionResponse) Reset() {
	*x = HostActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostActionResponse) ProtoMessage() {}

func (x *HostActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostActionResponse.ProtoReflect.Descriptor instead.
func (*HostActionResponse) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *HostActionResponse) GetProcesses() []*ActionKillProcess {
	if x != nil {
		return x.Processes
	}
	return nil
}

func (x *HostActionResponse) GetUsers() []*ActionLogoutUser {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_v1_service_proto protoreflect.FileDescriptor

var file_v1_service_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0c, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x18, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x3d, 0x0a, 0x11, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x10, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x75,
	0x0a, 0x12, 0x48, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x4f, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6c, 0x61, 0x6c, 0x6c, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x2f,
	0x70, 0x61, 0x6e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_service_proto_rawDescOnce sync.Once
	file_v1_service_proto_rawDescData = file_v1_service_proto_rawDesc
)

func file_v1_service_proto_rawDescGZIP() []byte {
	file_v1_service_proto_rawDescOnce.Do(func() {
		file_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_service_proto_rawDescData)
	})
	return file_v1_service_proto_rawDescData
}

var file_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_service_proto_goTypes = []interface{}{
	(*HostProcessReportRequest)(nil), // 0: v1.HostProcessReportRequest
	(*ActionKillProcess)(nil),        // 1: v1.ActionKillProcess
	(*ActionLogoutUser)(nil),         // 2: v1.ActionLogoutUser
	(*HostActionResponse)(nil),       // 3: v1.HostActionResponse
	(*ProcessReport)(nil),            // 4: v1.ProcessReport
}
var file_v1_service_proto_depIdxs = []int32{
	4, // 0: v1.HostProcessReportRequest.report:type_name -> v1.ProcessReport
	1, // 1: v1.HostActionResponse.processes:type_name -> v1.ActionKillProcess
	2, // 2: v1.HostActionResponse.users:type_name -> v1.ActionLogoutUser
	0, // 3: v1.SearchService.Report:input_type -> v1.HostProcessReportRequest
	3, // 4: v1.SearchService.Report:output_type -> v1.HostActionResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_service_proto_init() }
func file_v1_service_proto_init() {
	if File_v1_service_proto != nil {
		return
	}
	file_v1_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostProcessReportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionKillProcess); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionLogoutUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostActionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_service_proto_goTypes,
		DependencyIndexes: file_v1_service_proto_depIdxs,
		MessageInfos:      file_v1_service_proto_msgTypes,
	}.Build()
	File_v1_service_proto = out.File
	file_v1_service_proto_rawDesc = nil
	file_v1_service_proto_goTypes = nil
	file_v1_service_proto_depIdxs = nil
}