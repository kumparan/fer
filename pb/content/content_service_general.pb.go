// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/content/content_service_general.proto

package content

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Empty used when an RPC doesn't need to return any message
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// FindByIDRequest :nodoc:
type FindByIDRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindByIDRequest) Reset()         { *m = FindByIDRequest{} }
func (m *FindByIDRequest) String() string { return proto.CompactTextString(m) }
func (*FindByIDRequest) ProtoMessage()    {}
func (*FindByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{1}
}

func (m *FindByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIDRequest.Unmarshal(m, b)
}
func (m *FindByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIDRequest.Marshal(b, m, deterministic)
}
func (m *FindByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIDRequest.Merge(m, src)
}
func (m *FindByIDRequest) XXX_Size() int {
	return xxx_messageInfo_FindByIDRequest.Size(m)
}
func (m *FindByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIDRequest proto.InternalMessageInfo

func (m *FindByIDRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// FindBySlugRequest :nodoc:
type FindBySlugRequest struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindBySlugRequest) Reset()         { *m = FindBySlugRequest{} }
func (m *FindBySlugRequest) String() string { return proto.CompactTextString(m) }
func (*FindBySlugRequest) ProtoMessage()    {}
func (*FindBySlugRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{2}
}

func (m *FindBySlugRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindBySlugRequest.Unmarshal(m, b)
}
func (m *FindBySlugRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindBySlugRequest.Marshal(b, m, deterministic)
}
func (m *FindBySlugRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindBySlugRequest.Merge(m, src)
}
func (m *FindBySlugRequest) XXX_Size() int {
	return xxx_messageInfo_FindBySlugRequest.Size(m)
}
func (m *FindBySlugRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindBySlugRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindBySlugRequest proto.InternalMessageInfo

func (m *FindBySlugRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

// FindByIDsRequest :nodoc:
type FindByIDsRequest struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindByIDsRequest) Reset()         { *m = FindByIDsRequest{} }
func (m *FindByIDsRequest) String() string { return proto.CompactTextString(m) }
func (*FindByIDsRequest) ProtoMessage()    {}
func (*FindByIDsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{3}
}

func (m *FindByIDsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIDsRequest.Unmarshal(m, b)
}
func (m *FindByIDsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIDsRequest.Marshal(b, m, deterministic)
}
func (m *FindByIDsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIDsRequest.Merge(m, src)
}
func (m *FindByIDsRequest) XXX_Size() int {
	return xxx_messageInfo_FindByIDsRequest.Size(m)
}
func (m *FindByIDsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIDsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIDsRequest proto.InternalMessageInfo

func (m *FindByIDsRequest) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

// FindMultiRequest :nodoc:
type FindMultiRequest struct {
	//
	//USAGE
	//1. by pagination: should have page and size. Example: ?page=2&size=10.
	//2. before-date: should only have before and size. Example: ?before=2017-09-10HH:mm:ssZ&size=25
	//3. after-date: should only have after and size. Example: ?after=2017-09-10HH:mm:ssZ&size=25
	//
	//Using 1, page will be set to 1 if page is not given
	//Using 2 or 3, page will be omitted.
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size"`
	Before               string   `protobuf:"bytes,3,opt,name=before,proto3" json:"before"`
	After                string   `protobuf:"bytes,4,opt,name=after,proto3" json:"after"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMultiRequest) Reset()         { *m = FindMultiRequest{} }
func (m *FindMultiRequest) String() string { return proto.CompactTextString(m) }
func (*FindMultiRequest) ProtoMessage()    {}
func (*FindMultiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{4}
}

func (m *FindMultiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMultiRequest.Unmarshal(m, b)
}
func (m *FindMultiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMultiRequest.Marshal(b, m, deterministic)
}
func (m *FindMultiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMultiRequest.Merge(m, src)
}
func (m *FindMultiRequest) XXX_Size() int {
	return xxx_messageInfo_FindMultiRequest.Size(m)
}
func (m *FindMultiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMultiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMultiRequest proto.InternalMessageInfo

func (m *FindMultiRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *FindMultiRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FindMultiRequest) GetBefore() string {
	if m != nil {
		return m.Before
	}
	return ""
}

func (m *FindMultiRequest) GetAfter() string {
	if m != nil {
		return m.After
	}
	return ""
}

// FindMultiByParentIDRequest :nodoc:
type FindMultiByParentIDRequest struct {
	AssociateId          int64             `protobuf:"varint,1,opt,name=associate_id,json=associateId,proto3" json:"associate_id"`
	Filter               *FindMultiRequest `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FindMultiByParentIDRequest) Reset()         { *m = FindMultiByParentIDRequest{} }
func (m *FindMultiByParentIDRequest) String() string { return proto.CompactTextString(m) }
func (*FindMultiByParentIDRequest) ProtoMessage()    {}
func (*FindMultiByParentIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{5}
}

func (m *FindMultiByParentIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMultiByParentIDRequest.Unmarshal(m, b)
}
func (m *FindMultiByParentIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMultiByParentIDRequest.Marshal(b, m, deterministic)
}
func (m *FindMultiByParentIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMultiByParentIDRequest.Merge(m, src)
}
func (m *FindMultiByParentIDRequest) XXX_Size() int {
	return xxx_messageInfo_FindMultiByParentIDRequest.Size(m)
}
func (m *FindMultiByParentIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMultiByParentIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMultiByParentIDRequest proto.InternalMessageInfo

func (m *FindMultiByParentIDRequest) GetAssociateId() int64 {
	if m != nil {
		return m.AssociateId
	}
	return 0
}

func (m *FindMultiByParentIDRequest) GetFilter() *FindMultiRequest {
	if m != nil {
		return m.Filter
	}
	return nil
}

// FindMultiByParentIDsRequest :nodoc:
type FindMultiByParentIDsRequest struct {
	AssociateIds         []int64           `protobuf:"varint,1,rep,packed,name=associate_ids,json=associateIds,proto3" json:"associate_ids"`
	Filter               *FindMultiRequest `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FindMultiByParentIDsRequest) Reset()         { *m = FindMultiByParentIDsRequest{} }
func (m *FindMultiByParentIDsRequest) String() string { return proto.CompactTextString(m) }
func (*FindMultiByParentIDsRequest) ProtoMessage()    {}
func (*FindMultiByParentIDsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{6}
}

func (m *FindMultiByParentIDsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMultiByParentIDsRequest.Unmarshal(m, b)
}
func (m *FindMultiByParentIDsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMultiByParentIDsRequest.Marshal(b, m, deterministic)
}
func (m *FindMultiByParentIDsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMultiByParentIDsRequest.Merge(m, src)
}
func (m *FindMultiByParentIDsRequest) XXX_Size() int {
	return xxx_messageInfo_FindMultiByParentIDsRequest.Size(m)
}
func (m *FindMultiByParentIDsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMultiByParentIDsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMultiByParentIDsRequest proto.InternalMessageInfo

func (m *FindMultiByParentIDsRequest) GetAssociateIds() []int64 {
	if m != nil {
		return m.AssociateIds
	}
	return nil
}

func (m *FindMultiByParentIDsRequest) GetFilter() *FindMultiRequest {
	if m != nil {
		return m.Filter
	}
	return nil
}

// FindMultiResponse :nodoc:
type FindMultiResponse struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMultiResponse) Reset()         { *m = FindMultiResponse{} }
func (m *FindMultiResponse) String() string { return proto.CompactTextString(m) }
func (*FindMultiResponse) ProtoMessage()    {}
func (*FindMultiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{7}
}

func (m *FindMultiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMultiResponse.Unmarshal(m, b)
}
func (m *FindMultiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMultiResponse.Marshal(b, m, deterministic)
}
func (m *FindMultiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMultiResponse.Merge(m, src)
}
func (m *FindMultiResponse) XXX_Size() int {
	return xxx_messageInfo_FindMultiResponse.Size(m)
}
func (m *FindMultiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMultiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMultiResponse proto.InternalMessageInfo

func (m *FindMultiResponse) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FindMultiResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// BooleanResponse :nodoc:
type BooleanResponse struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BooleanResponse) Reset()         { *m = BooleanResponse{} }
func (m *BooleanResponse) String() string { return proto.CompactTextString(m) }
func (*BooleanResponse) ProtoMessage()    {}
func (*BooleanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{8}
}

func (m *BooleanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BooleanResponse.Unmarshal(m, b)
}
func (m *BooleanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BooleanResponse.Marshal(b, m, deterministic)
}
func (m *BooleanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BooleanResponse.Merge(m, src)
}
func (m *BooleanResponse) XXX_Size() int {
	return xxx_messageInfo_BooleanResponse.Size(m)
}
func (m *BooleanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BooleanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BooleanResponse proto.InternalMessageInfo

func (m *BooleanResponse) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

// DeleteByIDRequest can be used to many kind of objects
type DeleteByIDRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	ObjectId             int64    `protobuf:"varint,2,opt,name=object_id,json=objectId,proto3" json:"object_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteByIDRequest) Reset()         { *m = DeleteByIDRequest{} }
func (m *DeleteByIDRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteByIDRequest) ProtoMessage()    {}
func (*DeleteByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{9}
}

func (m *DeleteByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteByIDRequest.Unmarshal(m, b)
}
func (m *DeleteByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteByIDRequest.Marshal(b, m, deterministic)
}
func (m *DeleteByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteByIDRequest.Merge(m, src)
}
func (m *DeleteByIDRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteByIDRequest.Size(m)
}
func (m *DeleteByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteByIDRequest proto.InternalMessageInfo

func (m *DeleteByIDRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DeleteByIDRequest) GetObjectId() int64 {
	if m != nil {
		return m.ObjectId
	}
	return 0
}

// MutateByIDRequest can be used to many kind of objects
type MutateByIDRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	ObjectId             int64    `protobuf:"varint,2,opt,name=object_id,json=objectId,proto3" json:"object_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateByIDRequest) Reset()         { *m = MutateByIDRequest{} }
func (m *MutateByIDRequest) String() string { return proto.CompactTextString(m) }
func (*MutateByIDRequest) ProtoMessage()    {}
func (*MutateByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_114fbaaa806d4cf1, []int{10}
}

func (m *MutateByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateByIDRequest.Unmarshal(m, b)
}
func (m *MutateByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateByIDRequest.Marshal(b, m, deterministic)
}
func (m *MutateByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateByIDRequest.Merge(m, src)
}
func (m *MutateByIDRequest) XXX_Size() int {
	return xxx_messageInfo_MutateByIDRequest.Size(m)
}
func (m *MutateByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateByIDRequest proto.InternalMessageInfo

func (m *MutateByIDRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MutateByIDRequest) GetObjectId() int64 {
	if m != nil {
		return m.ObjectId
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.content.Empty")
	proto.RegisterType((*FindByIDRequest)(nil), "pb.content.FindByIDRequest")
	proto.RegisterType((*FindBySlugRequest)(nil), "pb.content.FindBySlugRequest")
	proto.RegisterType((*FindByIDsRequest)(nil), "pb.content.FindByIDsRequest")
	proto.RegisterType((*FindMultiRequest)(nil), "pb.content.FindMultiRequest")
	proto.RegisterType((*FindMultiByParentIDRequest)(nil), "pb.content.FindMultiByParentIDRequest")
	proto.RegisterType((*FindMultiByParentIDsRequest)(nil), "pb.content.FindMultiByParentIDsRequest")
	proto.RegisterType((*FindMultiResponse)(nil), "pb.content.FindMultiResponse")
	proto.RegisterType((*BooleanResponse)(nil), "pb.content.BooleanResponse")
	proto.RegisterType((*DeleteByIDRequest)(nil), "pb.content.DeleteByIDRequest")
	proto.RegisterType((*MutateByIDRequest)(nil), "pb.content.MutateByIDRequest")
}

func init() {
	proto.RegisterFile("pb/content/content_service_general.proto", fileDescriptor_114fbaaa806d4cf1)
}

var fileDescriptor_114fbaaa806d4cf1 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0x9b, 0x6d, 0xba, 0x9d, 0x5d, 0xd8, 0x36, 0xaa, 0x20, 0xa2, 0x1c, 0x5a, 0x83, 0xd4,
	0x9e, 0x82, 0x04, 0xdc, 0xb8, 0x45, 0x05, 0x29, 0x87, 0x4a, 0x28, 0xdc, 0xb8, 0x54, 0x4e, 0x32,
	0x0d, 0x46, 0xc6, 0x0e, 0xb1, 0x5d, 0x51, 0xbe, 0x1e, 0x39, 0x4e, 0xd2, 0x08, 0x38, 0x21, 0x4e,
	0x99, 0x37, 0x7e, 0x79, 0xcf, 0xf3, 0xc6, 0xb0, 0xab, 0xb2, 0x57, 0xb9, 0x14, 0x1a, 0x85, 0xee,
	0xbe, 0x47, 0x85, 0xf5, 0x99, 0xe5, 0x78, 0x2c, 0x51, 0x60, 0x4d, 0x79, 0x54, 0xd5, 0x52, 0xcb,
	0x00, 0xaa, 0x2c, 0x6a, 0x19, 0x64, 0x0a, 0x93, 0xf7, 0xdf, 0x2a, 0x7d, 0x21, 0x1b, 0x78, 0xf8,
	0xc0, 0x44, 0x11, 0x5f, 0x92, 0x7d, 0x8a, 0xdf, 0x0d, 0x2a, 0x1d, 0x3c, 0x86, 0x31, 0x2b, 0xc2,
	0xd1, 0x7a, 0xb4, 0xf3, 0xd2, 0x31, 0x2b, 0xc8, 0x16, 0x16, 0x8e, 0xf2, 0x89, 0x9b, 0xb2, 0x23,
	0x05, 0x70, 0xa3, 0xb8, 0x29, 0x1b, 0xda, 0x2c, 0x6d, 0x6a, 0xf2, 0x12, 0xe6, 0x9d, 0x96, 0xea,
	0x78, 0x73, 0xf0, 0x58, 0xa1, 0xc2, 0xd1, 0xda, 0xdb, 0x79, 0xa9, 0x2d, 0xc9, 0x17, 0xc7, 0x3a,
	0x18, 0xae, 0xd9, 0x40, 0xad, 0xa2, 0x25, 0xb6, 0xa6, 0x4d, 0xdd, 0x38, 0xb0, 0x9f, 0x18, 0x8e,
	0x5d, 0xcf, 0xd6, 0xc1, 0x13, 0xf0, 0x33, 0x3c, 0xc9, 0x1a, 0x43, 0xaf, 0xf1, 0x6d, 0x51, 0xb0,
	0x84, 0x09, 0x3d, 0x69, 0xac, 0xc3, 0x9b, 0xa6, 0xed, 0x00, 0x31, 0xf0, 0xac, 0x77, 0x8a, 0x2f,
	0x1f, 0x69, 0x8d, 0x42, 0x5f, 0xc7, 0xdc, 0xc0, 0x3d, 0x55, 0x4a, 0xe6, 0x8c, 0x6a, 0x3c, 0xf6,
	0x03, 0xdf, 0xf5, 0xbd, 0xa4, 0x08, 0xde, 0x82, 0x7f, 0x62, 0xdc, 0xea, 0xda, 0x4b, 0xdc, 0xbd,
	0x7e, 0x1e, 0x5d, 0x23, 0x8c, 0x7e, 0x1f, 0x22, 0x6d, 0xb9, 0xe4, 0x07, 0xac, 0xfe, 0x62, 0xdb,
	0x27, 0xf2, 0x02, 0x1e, 0x0d, 0x7d, 0xbb, 0x6c, 0xee, 0x07, 0xc6, 0xea, 0x1f, 0x9d, 0xdf, 0xb9,
	0x4d, 0xb5, 0x67, 0xaa, 0x92, 0x42, 0xe1, 0x9f, 0x1b, 0xb0, 0x69, 0xe5, 0xd2, 0x08, 0xdd, 0x46,
	0xeb, 0x00, 0xd9, 0xc2, 0x43, 0x2c, 0x25, 0x47, 0x2a, 0xfa, 0x5f, 0x97, 0x30, 0x39, 0x53, 0x6e,
	0xdc, 0x5e, 0x6e, 0x53, 0x07, 0x48, 0x02, 0x8b, 0x3d, 0x72, 0xd4, 0x38, 0x7c, 0x34, 0x4f, 0x61,
	0x6a, 0x14, 0xd6, 0xd7, 0x20, 0x7d, 0x0b, 0x93, 0x22, 0x58, 0xc1, 0x4c, 0x66, 0x5f, 0x31, 0xd7,
	0xf6, 0xc8, 0x19, 0xde, 0xba, 0x46, 0x52, 0x58, 0xa9, 0x83, 0xd1, 0xf4, 0x3f, 0x48, 0xc5, 0xb3,
	0xcf, 0xd3, 0x36, 0x9f, 0xcc, 0x6f, 0xde, 0xfb, 0x9b, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb,
	0x93, 0x80, 0xd6, 0x1b, 0x03, 0x00, 0x00,
}
