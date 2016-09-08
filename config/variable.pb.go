// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/variable.proto
// DO NOT EDIT!

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing a Variable.
type VariableDef struct {
	// Name of the variable tensor.
	VariableName string `protobuf:"bytes,1,opt,name=variable_name,json=variableName" json:"variable_name,omitempty"`
	// Name of the initializer op.
	InitializerName string `protobuf:"bytes,2,opt,name=initializer_name,json=initializerName" json:"initializer_name,omitempty"`
	// Name of the snapshot tensor.
	SnapshotName string `protobuf:"bytes,3,opt,name=snapshot_name,json=snapshotName" json:"snapshot_name,omitempty"`
	// Support for saving variables as slices of a larger variable.
	SaveSliceInfoDef *SaveSliceInfoDef `protobuf:"bytes,4,opt,name=save_slice_info_def,json=saveSliceInfoDef" json:"save_slice_info_def,omitempty"`
}

func (m *VariableDef) Reset()                    { *m = VariableDef{} }
func (m *VariableDef) String() string            { return proto.CompactTextString(m) }
func (*VariableDef) ProtoMessage()               {}
func (*VariableDef) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *VariableDef) GetSaveSliceInfoDef() *SaveSliceInfoDef {
	if m != nil {
		return m.SaveSliceInfoDef
	}
	return nil
}

type SaveSliceInfoDef struct {
	// Name of the full variable of which this is a slice.
	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	// Shape of the full variable.
	FullShape []int32 `protobuf:"varint,2,rep,packed,name=full_shape,json=fullShape" json:"full_shape,omitempty"`
	// Offset of this variable into the full variable.
	VarOffset []int32 `protobuf:"varint,3,rep,packed,name=var_offset,json=varOffset" json:"var_offset,omitempty"`
	// Shape of this variable.
	VarShape []int32 `protobuf:"varint,4,rep,packed,name=var_shape,json=varShape" json:"var_shape,omitempty"`
}

func (m *SaveSliceInfoDef) Reset()                    { *m = SaveSliceInfoDef{} }
func (m *SaveSliceInfoDef) String() string            { return proto.CompactTextString(m) }
func (*SaveSliceInfoDef) ProtoMessage()               {}
func (*SaveSliceInfoDef) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{1} }

func init() {
	proto.RegisterType((*VariableDef)(nil), "tensorflow.VariableDef")
	proto.RegisterType((*SaveSliceInfoDef)(nil), "tensorflow.SaveSliceInfoDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/variable.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xd9, 0xaf, 0xfd, 0xa4, 0x9d, 0x5a, 0x2d, 0xf1, 0x12, 0x50, 0xa1, 0xb4, 0x97, 0x78,
	0x69, 0x40, 0xdf, 0xa0, 0xf4, 0x22, 0x82, 0x96, 0x04, 0xbc, 0x86, 0x69, 0x9d, 0xb5, 0x8b, 0x9b,
	0xdd, 0xb0, 0x1b, 0xb7, 0xe0, 0x23, 0xf8, 0x72, 0xbe, 0x8e, 0x47, 0xd9, 0x8d, 0xb1, 0xa1, 0xd7,
	0xdf, 0xff, 0x37, 0x93, 0xcc, 0xfe, 0x21, 0xa9, 0x49, 0x59, 0x6d, 0xb8, 0xd4, 0xfb, 0x74, 0xab,
	0x0d, 0xa5, 0xdc, 0x60, 0x49, 0x7b, 0x6d, 0xde, 0x52, 0x87, 0x46, 0xe0, 0x46, 0xd2, 0xa2, 0x32,
	0xba, 0xd6, 0x11, 0x1c, 0xcc, 0xd9, 0x17, 0x83, 0xd1, 0xf3, 0x6f, 0xbc, 0x22, 0x1e, 0xcd, 0x61,
	0xdc, 0xda, 0x85, 0xc2, 0x92, 0x62, 0x36, 0x65, 0xc9, 0x30, 0x3b, 0x6d, 0xe1, 0x23, 0x96, 0x14,
	0xdd, 0xc0, 0x44, 0x28, 0x51, 0x0b, 0x94, 0xe2, 0x83, 0x4c, 0xe3, 0xfd, 0x0b, 0xde, 0x79, 0x87,
	0x07, 0x75, 0x0e, 0x63, 0xab, 0xb0, 0xb2, 0x3b, 0x5d, 0x37, 0x5e, 0xaf, 0xd9, 0xd7, 0xc2, 0x20,
	0x3d, 0xc0, 0x85, 0x45, 0x47, 0x85, 0x95, 0x62, 0x4b, 0x85, 0x50, 0x5c, 0x17, 0x2f, 0xc4, 0xe3,
	0xfe, 0x94, 0x25, 0xa3, 0xdb, 0xab, 0xc5, 0xe1, 0x77, 0x17, 0x39, 0x3a, 0xca, 0xbd, 0x75, 0xaf,
	0xb8, 0x5e, 0x11, 0xcf, 0x26, 0xf6, 0x88, 0xcc, 0x3e, 0x19, 0x4c, 0x8e, 0xb5, 0xe8, 0x12, 0x86,
	0xfc, 0x5d, 0xca, 0xee, 0x49, 0x03, 0x0f, 0xc2, 0xe7, 0xaf, 0x01, 0x42, 0x68, 0x77, 0x58, 0xf9,
	0x43, 0x7a, 0xc9, 0xff, 0x2c, 0xe8, 0xb9, 0x07, 0x3e, 0x76, 0x68, 0x0a, 0xcd, 0xb9, 0xa5, 0x3a,
	0xee, 0x35, 0xb1, 0x43, 0xf3, 0x14, 0x80, 0x5f, 0xed, 0xe3, 0x66, 0xb8, 0x1f, 0xd2, 0x81, 0x43,
	0x13, 0x66, 0x97, 0x29, 0xc4, 0xda, 0xbc, 0x76, 0x2f, 0xf8, 0x6b, 0x65, 0x79, 0xd6, 0xbe, 0xfb,
	0xda, 0xb7, 0x62, 0xd7, 0xec, 0x9b, 0xb1, 0xcd, 0x49, 0xa8, 0xe8, 0xee, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0x57, 0x9f, 0xbb, 0xce, 0x01, 0x00, 0x00,
}