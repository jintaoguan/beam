// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	v1.proto

It has these top-level messages:
	Type
	FullType
	UserFn
	DynFn
	Fn
	CustomCoder
	MultiEdge
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Kind is mostly identical to reflect.TypeKind, expect we handle certain
// types specially, such as "error".
type Type_Kind int32

const (
	Type_INVALID Type_Kind = 0
	// Primitive.
	Type_BOOL   Type_Kind = 1
	Type_INT    Type_Kind = 2
	Type_INT8   Type_Kind = 3
	Type_INT16  Type_Kind = 4
	Type_INT32  Type_Kind = 5
	Type_INT64  Type_Kind = 6
	Type_UINT   Type_Kind = 7
	Type_UINT8  Type_Kind = 8
	Type_UINT16 Type_Kind = 9
	Type_UINT32 Type_Kind = 10
	Type_UINT64 Type_Kind = 11
	Type_STRING Type_Kind = 12
	// Non-primitive types.
	Type_SLICE    Type_Kind = 20
	Type_STRUCT   Type_Kind = 21
	Type_FUNC     Type_Kind = 22
	Type_CHAN     Type_Kind = 23
	Type_PTR      Type_Kind = 24
	Type_SPECIAL  Type_Kind = 25
	Type_EXTERNAL Type_Kind = 26
)

var Type_Kind_name = map[int32]string{
	0:  "INVALID",
	1:  "BOOL",
	2:  "INT",
	3:  "INT8",
	4:  "INT16",
	5:  "INT32",
	6:  "INT64",
	7:  "UINT",
	8:  "UINT8",
	9:  "UINT16",
	10: "UINT32",
	11: "UINT64",
	12: "STRING",
	20: "SLICE",
	21: "STRUCT",
	22: "FUNC",
	23: "CHAN",
	24: "PTR",
	25: "SPECIAL",
	26: "EXTERNAL",
}
var Type_Kind_value = map[string]int32{
	"INVALID":  0,
	"BOOL":     1,
	"INT":      2,
	"INT8":     3,
	"INT16":    4,
	"INT32":    5,
	"INT64":    6,
	"UINT":     7,
	"UINT8":    8,
	"UINT16":   9,
	"UINT32":   10,
	"UINT64":   11,
	"STRING":   12,
	"SLICE":    20,
	"STRUCT":   21,
	"FUNC":     22,
	"CHAN":     23,
	"PTR":      24,
	"SPECIAL":  25,
	"EXTERNAL": 26,
}

func (x Type_Kind) String() string {
	return proto.EnumName(Type_Kind_name, int32(x))
}
func (Type_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// ChanDir matches reflect.ChanDir.
type Type_ChanDir int32

const (
	Type_RECV Type_ChanDir = 0
	Type_SEND Type_ChanDir = 1
	Type_BOTH Type_ChanDir = 2
)

var Type_ChanDir_name = map[int32]string{
	0: "RECV",
	1: "SEND",
	2: "BOTH",
}
var Type_ChanDir_value = map[string]int32{
	"RECV": 0,
	"SEND": 1,
	"BOTH": 2,
}

func (x Type_ChanDir) String() string {
	return proto.EnumName(Type_ChanDir_name, int32(x))
}
func (Type_ChanDir) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Type_Special int32

const (
	Type_ILLEGAL Type_Special = 0
	// Go
	Type_ERROR   Type_Special = 1
	Type_CONTEXT Type_Special = 2
	Type_TYPE    Type_Special = 3
	// Beam
	Type_EVENTTIME     Type_Special = 10
	Type_KV            Type_Special = 11
	Type_GBK           Type_Special = 12
	Type_COGBK         Type_Special = 13
	Type_WINDOWEDVALUE Type_Special = 14
	Type_T             Type_Special = 15
	Type_U             Type_Special = 16
	Type_V             Type_Special = 17
	Type_W             Type_Special = 18
	Type_X             Type_Special = 19
	Type_Y             Type_Special = 20
	Type_Z             Type_Special = 21
)

var Type_Special_name = map[int32]string{
	0:  "ILLEGAL",
	1:  "ERROR",
	2:  "CONTEXT",
	3:  "TYPE",
	10: "EVENTTIME",
	11: "KV",
	12: "GBK",
	13: "COGBK",
	14: "WINDOWEDVALUE",
	15: "T",
	16: "U",
	17: "V",
	18: "W",
	19: "X",
	20: "Y",
	21: "Z",
}
var Type_Special_value = map[string]int32{
	"ILLEGAL":       0,
	"ERROR":         1,
	"CONTEXT":       2,
	"TYPE":          3,
	"EVENTTIME":     10,
	"KV":            11,
	"GBK":           12,
	"COGBK":         13,
	"WINDOWEDVALUE": 14,
	"T":             15,
	"U":             16,
	"V":             17,
	"W":             18,
	"X":             19,
	"Y":             20,
	"Z":             21,
}

func (x Type_Special) String() string {
	return proto.EnumName(Type_Special_name, int32(x))
}
func (Type_Special) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

type MultiEdge_Inbound_InputKind int32

const (
	MultiEdge_Inbound_INVALID   MultiEdge_Inbound_InputKind = 0
	MultiEdge_Inbound_MAIN      MultiEdge_Inbound_InputKind = 1
	MultiEdge_Inbound_SINGLETON MultiEdge_Inbound_InputKind = 2
	MultiEdge_Inbound_SLICE     MultiEdge_Inbound_InputKind = 3
	MultiEdge_Inbound_MAP       MultiEdge_Inbound_InputKind = 4
	MultiEdge_Inbound_MULTIMAP  MultiEdge_Inbound_InputKind = 5
	MultiEdge_Inbound_ITER      MultiEdge_Inbound_InputKind = 6
	MultiEdge_Inbound_REITER    MultiEdge_Inbound_InputKind = 7
)

var MultiEdge_Inbound_InputKind_name = map[int32]string{
	0: "INVALID",
	1: "MAIN",
	2: "SINGLETON",
	3: "SLICE",
	4: "MAP",
	5: "MULTIMAP",
	6: "ITER",
	7: "REITER",
}
var MultiEdge_Inbound_InputKind_value = map[string]int32{
	"INVALID":   0,
	"MAIN":      1,
	"SINGLETON": 2,
	"SLICE":     3,
	"MAP":       4,
	"MULTIMAP":  5,
	"ITER":      6,
	"REITER":    7,
}

func (x MultiEdge_Inbound_InputKind) String() string {
	return proto.EnumName(MultiEdge_Inbound_InputKind_name, int32(x))
}
func (MultiEdge_Inbound_InputKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0, 0}
}

// Type represents a serializable reflect.Type.
type Type struct {
	// (Required) Type kind.
	Kind Type_Kind `protobuf:"varint,1,opt,name=kind,enum=v1.Type_Kind" json:"kind,omitempty"`
	// (Optional) Element type (if SLICE, PTR or CHAN)
	Element *Type `protobuf:"bytes,2,opt,name=element" json:"element,omitempty"`
	// (Optional) Fields (if STRUCT).
	Fields []*Type_StructField `protobuf:"bytes,3,rep,name=fields" json:"fields,omitempty"`
	// (Optional) Parameter types (if FUNC).
	ParameterTypes []*Type `protobuf:"bytes,4,rep,name=parameter_types,json=parameterTypes" json:"parameter_types,omitempty"`
	// (Optional) Return types (if FUNC).
	ReturnTypes []*Type `protobuf:"bytes,5,rep,name=return_types,json=returnTypes" json:"return_types,omitempty"`
	// (Optional) Is variadic (if FUNC).
	IsVariadic bool `protobuf:"varint,6,opt,name=is_variadic,json=isVariadic" json:"is_variadic,omitempty"`
	// (Optional) Channel direction (if CHAN).
	ChanDir Type_ChanDir `protobuf:"varint,7,opt,name=chan_dir,json=chanDir,enum=v1.Type_ChanDir" json:"chan_dir,omitempty"`
	// (Optional) Special type (if SPECIAL)
	Special Type_Special `protobuf:"varint,8,opt,name=special,enum=v1.Type_Special" json:"special,omitempty"`
	// (Optional) Key for external types.
	// External types are types that are not directly serialized using
	// the types above, but rather indirectly serialized.  The wire format
	// holds a lookup key into a registry to reify the types in a worker from a
	// registry. The main usage of external serialization is to preserve
	// methods attached to types.
	ExternalKey string `protobuf:"bytes,9,opt,name=external_key,json=externalKey" json:"external_key,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Type) GetKind() Type_Kind {
	if m != nil {
		return m.Kind
	}
	return Type_INVALID
}

func (m *Type) GetElement() *Type {
	if m != nil {
		return m.Element
	}
	return nil
}

func (m *Type) GetFields() []*Type_StructField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetParameterTypes() []*Type {
	if m != nil {
		return m.ParameterTypes
	}
	return nil
}

func (m *Type) GetReturnTypes() []*Type {
	if m != nil {
		return m.ReturnTypes
	}
	return nil
}

func (m *Type) GetIsVariadic() bool {
	if m != nil {
		return m.IsVariadic
	}
	return false
}

func (m *Type) GetChanDir() Type_ChanDir {
	if m != nil {
		return m.ChanDir
	}
	return Type_RECV
}

func (m *Type) GetSpecial() Type_Special {
	if m != nil {
		return m.Special
	}
	return Type_ILLEGAL
}

func (m *Type) GetExternalKey() string {
	if m != nil {
		return m.ExternalKey
	}
	return ""
}

// StructField matches reflect.StructField.
type Type_StructField struct {
	Name      string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PkgPath   string  `protobuf:"bytes,2,opt,name=pkg_path,json=pkgPath" json:"pkg_path,omitempty"`
	Type      *Type   `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Tag       string  `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	Offset    int64   `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Index     []int32 `protobuf:"varint,6,rep,packed,name=index" json:"index,omitempty"`
	Anonymous bool    `protobuf:"varint,7,opt,name=anonymous" json:"anonymous,omitempty"`
}

func (m *Type_StructField) Reset()                    { *m = Type_StructField{} }
func (m *Type_StructField) String() string            { return proto.CompactTextString(m) }
func (*Type_StructField) ProtoMessage()               {}
func (*Type_StructField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Type_StructField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type_StructField) GetPkgPath() string {
	if m != nil {
		return m.PkgPath
	}
	return ""
}

func (m *Type_StructField) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Type_StructField) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Type_StructField) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Type_StructField) GetIndex() []int32 {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Type_StructField) GetAnonymous() bool {
	if m != nil {
		return m.Anonymous
	}
	return false
}

// FullType represents a serialized typex.FullType
type FullType struct {
	Type       *Type       `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Components []*FullType `protobuf:"bytes,2,rep,name=components" json:"components,omitempty"`
}

func (m *FullType) Reset()                    { *m = FullType{} }
func (m *FullType) String() string            { return proto.CompactTextString(m) }
func (*FullType) ProtoMessage()               {}
func (*FullType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FullType) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *FullType) GetComponents() []*FullType {
	if m != nil {
		return m.Components
	}
	return nil
}

// UserFn represents a serialized function reference. The
// implementation is notably not serialized and must be present (and
// somehow discoverable from the symbol name) on the decoding side.
type UserFn struct {
	// (Required) Symbol name of function.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (Required) Function type.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *UserFn) Reset()                    { *m = UserFn{} }
func (m *UserFn) String() string            { return proto.CompactTextString(m) }
func (*UserFn) ProtoMessage()               {}
func (*UserFn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserFn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserFn) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

// DynFn represents a serialized function generator.
type DynFn struct {
	// (Required) Name of the generated function.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (Required) Type of the generated function.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// (Required) Input to generator.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// (Required) Symbol name of generator (of type []byte ->
	// []reflect.Value -> []reflect.Value).
	Gen string `protobuf:"bytes,4,opt,name=gen" json:"gen,omitempty"`
}

func (m *DynFn) Reset()                    { *m = DynFn{} }
func (m *DynFn) String() string            { return proto.CompactTextString(m) }
func (*DynFn) ProtoMessage()               {}
func (*DynFn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DynFn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DynFn) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DynFn) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DynFn) GetGen() string {
	if m != nil {
		return m.Gen
	}
	return ""
}

// Fn represents a serialized function reference or struct.
type Fn struct {
	// (Optional) Function reference.
	Fn *UserFn `protobuf:"bytes,1,opt,name=fn" json:"fn,omitempty"`
	// (Optional) Struct type.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// (Optional) JSON-serialized value, if struct.
	Opt string `protobuf:"bytes,3,opt,name=opt" json:"opt,omitempty"`
	// (Optional) Function generator, if dynamic function.
	Dynfn *DynFn `protobuf:"bytes,4,opt,name=dynfn" json:"dynfn,omitempty"`
}

func (m *Fn) Reset()                    { *m = Fn{} }
func (m *Fn) String() string            { return proto.CompactTextString(m) }
func (*Fn) ProtoMessage()               {}
func (*Fn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Fn) GetFn() *UserFn {
	if m != nil {
		return m.Fn
	}
	return nil
}

func (m *Fn) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Fn) GetOpt() string {
	if m != nil {
		return m.Opt
	}
	return ""
}

func (m *Fn) GetDynfn() *DynFn {
	if m != nil {
		return m.Dynfn
	}
	return nil
}

// CustomCoder
type CustomCoder struct {
	// (Required) Name of the coder. For informational purposes only.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (Required) Concrete type being coded.
	Type *Type `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// (Required) Encoding function.
	Enc *UserFn `protobuf:"bytes,3,opt,name=enc" json:"enc,omitempty"`
	// (Required) Decoding function.
	Dec *UserFn `protobuf:"bytes,4,opt,name=dec" json:"dec,omitempty"`
}

func (m *CustomCoder) Reset()                    { *m = CustomCoder{} }
func (m *CustomCoder) String() string            { return proto.CompactTextString(m) }
func (*CustomCoder) ProtoMessage()               {}
func (*CustomCoder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CustomCoder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomCoder) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *CustomCoder) GetEnc() *UserFn {
	if m != nil {
		return m.Enc
	}
	return nil
}

func (m *CustomCoder) GetDec() *UserFn {
	if m != nil {
		return m.Dec
	}
	return nil
}

// MultiEdge represents a partly-serialized MultiEdge. It does not include
// node information, because runners manipulate the graph structure.
type MultiEdge struct {
	Fn       *Fn                   `protobuf:"bytes,1,opt,name=fn" json:"fn,omitempty"`
	Opcode   string                `protobuf:"bytes,4,opt,name=opcode" json:"opcode,omitempty"`
	Inbound  []*MultiEdge_Inbound  `protobuf:"bytes,2,rep,name=inbound" json:"inbound,omitempty"`
	Outbound []*MultiEdge_Outbound `protobuf:"bytes,3,rep,name=outbound" json:"outbound,omitempty"`
}

func (m *MultiEdge) Reset()                    { *m = MultiEdge{} }
func (m *MultiEdge) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge) ProtoMessage()               {}
func (*MultiEdge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MultiEdge) GetFn() *Fn {
	if m != nil {
		return m.Fn
	}
	return nil
}

func (m *MultiEdge) GetOpcode() string {
	if m != nil {
		return m.Opcode
	}
	return ""
}

func (m *MultiEdge) GetInbound() []*MultiEdge_Inbound {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *MultiEdge) GetOutbound() []*MultiEdge_Outbound {
	if m != nil {
		return m.Outbound
	}
	return nil
}

type MultiEdge_Inbound struct {
	Kind MultiEdge_Inbound_InputKind `protobuf:"varint,1,opt,name=kind,enum=v1.MultiEdge_Inbound_InputKind" json:"kind,omitempty"`
	Type *FullType                   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *MultiEdge_Inbound) Reset()                    { *m = MultiEdge_Inbound{} }
func (m *MultiEdge_Inbound) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge_Inbound) ProtoMessage()               {}
func (*MultiEdge_Inbound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *MultiEdge_Inbound) GetKind() MultiEdge_Inbound_InputKind {
	if m != nil {
		return m.Kind
	}
	return MultiEdge_Inbound_INVALID
}

func (m *MultiEdge_Inbound) GetType() *FullType {
	if m != nil {
		return m.Type
	}
	return nil
}

type MultiEdge_Outbound struct {
	Type *FullType `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *MultiEdge_Outbound) Reset()                    { *m = MultiEdge_Outbound{} }
func (m *MultiEdge_Outbound) String() string            { return proto.CompactTextString(m) }
func (*MultiEdge_Outbound) ProtoMessage()               {}
func (*MultiEdge_Outbound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 1} }

func (m *MultiEdge_Outbound) GetType() *FullType {
	if m != nil {
		return m.Type
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "v1.Type")
	proto.RegisterType((*Type_StructField)(nil), "v1.Type.StructField")
	proto.RegisterType((*FullType)(nil), "v1.FullType")
	proto.RegisterType((*UserFn)(nil), "v1.UserFn")
	proto.RegisterType((*DynFn)(nil), "v1.DynFn")
	proto.RegisterType((*Fn)(nil), "v1.Fn")
	proto.RegisterType((*CustomCoder)(nil), "v1.CustomCoder")
	proto.RegisterType((*MultiEdge)(nil), "v1.MultiEdge")
	proto.RegisterType((*MultiEdge_Inbound)(nil), "v1.MultiEdge.Inbound")
	proto.RegisterType((*MultiEdge_Outbound)(nil), "v1.MultiEdge.Outbound")
	proto.RegisterEnum("v1.Type_Kind", Type_Kind_name, Type_Kind_value)
	proto.RegisterEnum("v1.Type_ChanDir", Type_ChanDir_name, Type_ChanDir_value)
	proto.RegisterEnum("v1.Type_Special", Type_Special_name, Type_Special_value)
	proto.RegisterEnum("v1.MultiEdge_Inbound_InputKind", MultiEdge_Inbound_InputKind_name, MultiEdge_Inbound_InputKind_value)
}

func init() { proto.RegisterFile("v1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1014 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0x45, 0x89, 0x3f, 0x43, 0x39, 0xd9, 0x6c, 0x6d, 0x97, 0x31, 0x0c, 0x58, 0xe1, 0xa5,
	0x42, 0x63, 0xa8, 0xb0, 0x1c, 0x18, 0x45, 0x6f, 0x0a, 0x45, 0x3b, 0x84, 0x69, 0xca, 0x58, 0x51,
	0xb2, 0xd3, 0x8b, 0xc0, 0x88, 0x2b, 0x99, 0xb0, 0xb4, 0x64, 0x49, 0xca, 0x88, 0xd0, 0x57, 0xe9,
	0xbd, 0xaf, 0xd0, 0x77, 0xe8, 0x73, 0xf4, 0xda, 0x67, 0x28, 0x96, 0x3f, 0xf2, 0x4f, 0x5d, 0x14,
	0xc8, 0x69, 0xbf, 0x9d, 0x99, 0x6f, 0xe7, 0x67, 0x67, 0x67, 0x41, 0xb9, 0x3b, 0xea, 0xc4, 0x49,
	0x94, 0x45, 0xb8, 0x76, 0x77, 0x64, 0xfc, 0x2d, 0x43, 0xdd, 0x5b, 0xc7, 0x14, 0xbf, 0x85, 0xfa,
	0x6d, 0xc8, 0x02, 0x5d, 0x68, 0x09, 0xed, 0x97, 0xdd, 0xad, 0xce, 0xdd, 0x51, 0x87, 0xcb, 0x3b,
	0xe7, 0x21, 0x0b, 0x48, 0xae, 0xc2, 0x06, 0xc8, 0x74, 0x41, 0x97, 0x94, 0x65, 0x7a, 0xad, 0x25,
	0xb4, 0xb5, 0xae, 0x52, 0x59, 0x91, 0x4a, 0x81, 0x0f, 0x41, 0x9a, 0x85, 0x74, 0x11, 0xa4, 0xba,
	0xd8, 0x12, 0xdb, 0x5a, 0x77, 0x7b, 0x73, 0xd0, 0x30, 0x4b, 0x56, 0xd3, 0xec, 0x94, 0x2b, 0x49,
	0x69, 0x83, 0x8f, 0xe0, 0x55, 0xec, 0x27, 0xfe, 0x92, 0x66, 0x34, 0x99, 0x64, 0xeb, 0x98, 0xa6,
	0x7a, 0x3d, 0xa7, 0xdd, 0x9f, 0xfc, 0x72, 0x63, 0xc0, 0xb7, 0x29, 0x7e, 0x07, 0xcd, 0x84, 0x66,
	0xab, 0x84, 0x95, 0xf6, 0x8d, 0x27, 0xf6, 0x5a, 0xa1, 0x2d, 0x8c, 0x0f, 0x40, 0x0b, 0xd3, 0xc9,
	0x9d, 0x9f, 0x84, 0x7e, 0x10, 0x4e, 0x75, 0xa9, 0x25, 0xb4, 0x15, 0x02, 0x61, 0x3a, 0x2e, 0x25,
	0xf8, 0x1d, 0x28, 0xd3, 0x1b, 0x9f, 0x4d, 0x82, 0x30, 0xd1, 0xe5, 0x3c, 0x73, 0xb4, 0x09, 0xd8,
	0xbc, 0xf1, 0x59, 0x3f, 0x4c, 0x88, 0x3c, 0x2d, 0x00, 0xfe, 0x1e, 0xe4, 0x34, 0xa6, 0xd3, 0xd0,
	0x5f, 0xe8, 0xca, 0x13, 0xdb, 0x61, 0x21, 0x27, 0x95, 0x01, 0x7e, 0x0b, 0x4d, 0xfa, 0x25, 0xa3,
	0x09, 0xf3, 0x17, 0x93, 0x5b, 0xba, 0xd6, 0xd5, 0x96, 0xd0, 0x56, 0x89, 0x56, 0xc9, 0xce, 0xe9,
	0x7a, 0xef, 0x0f, 0x01, 0xb4, 0x07, 0x45, 0xc1, 0x18, 0xea, 0xcc, 0x5f, 0xd2, 0xfc, 0x06, 0x54,
	0x92, 0x63, 0xfc, 0x06, 0x94, 0xf8, 0x76, 0x3e, 0x89, 0xfd, 0xec, 0x26, 0xaf, 0xb9, 0x4a, 0xe4,
	0xf8, 0x76, 0x7e, 0xe9, 0x67, 0x37, 0x78, 0x1f, 0xea, 0xbc, 0x02, 0xba, 0xf8, 0xe4, 0x2a, 0x72,
	0x29, 0x46, 0x20, 0x66, 0xfe, 0x5c, 0xaf, 0xe7, 0x1c, 0x0e, 0xf1, 0x2e, 0x48, 0xd1, 0x6c, 0x96,
	0xd2, 0x4c, 0x6f, 0xb4, 0x84, 0xb6, 0x48, 0xca, 0x1d, 0xde, 0x86, 0x46, 0xc8, 0x02, 0xfa, 0x45,
	0x97, 0x5a, 0x62, 0xbb, 0x41, 0x8a, 0x0d, 0xde, 0x07, 0xd5, 0x67, 0x11, 0x5b, 0x2f, 0xa3, 0x55,
	0x9a, 0x57, 0x46, 0x21, 0xf7, 0x02, 0xe3, 0x2f, 0x01, 0xea, 0xbc, 0x31, 0xb0, 0x06, 0xb2, 0xed,
	0x8e, 0x7b, 0x8e, 0xdd, 0x47, 0x2f, 0xb0, 0x02, 0xf5, 0x0f, 0x83, 0x81, 0x83, 0x04, 0x2c, 0x83,
	0x68, 0xbb, 0x1e, 0xaa, 0x71, 0x91, 0xed, 0x7a, 0x3f, 0x22, 0x11, 0xab, 0xd0, 0xb0, 0x5d, 0xef,
	0xe8, 0x04, 0xd5, 0x4b, 0x78, 0xdc, 0x45, 0x8d, 0x12, 0x9e, 0xbc, 0x47, 0x12, 0x37, 0x1d, 0x71,
	0x92, 0xcc, 0x85, 0xa3, 0x9c, 0xa5, 0x60, 0x00, 0x69, 0x54, 0xd0, 0xd4, 0x0a, 0x1f, 0x77, 0x11,
	0x54, 0xf8, 0xe4, 0x3d, 0xd2, 0x38, 0x1e, 0x7a, 0xc4, 0x76, 0xcf, 0x50, 0x93, 0x53, 0x87, 0x8e,
	0x6d, 0x5a, 0x68, 0xbb, 0x14, 0x8f, 0x4c, 0x0f, 0xed, 0xf0, 0xb3, 0x4f, 0x47, 0xae, 0x89, 0x76,
	0x39, 0x32, 0x3f, 0xf6, 0x5c, 0xf4, 0x2d, 0x8f, 0xf1, 0xd2, 0x23, 0x48, 0xe7, 0x39, 0x0c, 0x2f,
	0x2d, 0xd3, 0xee, 0x39, 0xe8, 0x0d, 0x6e, 0x82, 0x62, 0x5d, 0x7b, 0x16, 0x71, 0x7b, 0x0e, 0xda,
	0x33, 0xbe, 0x03, 0xb9, 0xec, 0x02, 0x4e, 0x24, 0x96, 0x39, 0x2e, 0xd2, 0x1c, 0x5a, 0x6e, 0x1f,
	0x09, 0x45, 0xc2, 0xde, 0x47, 0x54, 0x33, 0x7e, 0x17, 0x40, 0x2e, 0x7b, 0x20, 0xaf, 0x89, 0xe3,
	0x58, 0x67, 0x3d, 0x07, 0xbd, 0xe0, 0x01, 0x59, 0x84, 0x0c, 0x08, 0x12, 0xb8, 0xdc, 0x1c, 0xb8,
	0x9e, 0x75, 0x5d, 0x16, 0xc6, 0xfb, 0x74, 0x69, 0x21, 0x11, 0x6f, 0x81, 0x6a, 0x8d, 0x2d, 0xd7,
	0xf3, 0xec, 0x0b, 0x0b, 0x01, 0x96, 0xa0, 0x76, 0x3e, 0x46, 0x1a, 0x0f, 0xef, 0xec, 0xc3, 0x79,
	0x91, 0x92, 0x39, 0xe0, 0x70, 0x0b, 0xbf, 0x86, 0xad, 0x2b, 0xdb, 0xed, 0x0f, 0xae, 0xac, 0xfe,
	0xb8, 0xe7, 0x8c, 0x2c, 0xf4, 0x12, 0x37, 0x40, 0xf0, 0xd0, 0x2b, 0xbe, 0x8c, 0x10, 0xe2, 0xcb,
	0x18, 0xbd, 0xe6, 0xcb, 0x15, 0xc2, 0x7c, 0xb9, 0x46, 0xdf, 0xf0, 0xe5, 0x13, 0xda, 0xe6, 0xcb,
	0xcf, 0x68, 0xc7, 0x18, 0x83, 0x72, 0xba, 0x5a, 0x2c, 0xf2, 0x37, 0x5f, 0xb5, 0x90, 0xf0, 0x6c,
	0x0b, 0x1d, 0x02, 0x4c, 0xa3, 0x65, 0x1c, 0x31, 0xca, 0xb2, 0x54, 0xaf, 0xe5, 0xef, 0xac, 0xc9,
	0x6d, 0x2a, 0x3e, 0x79, 0xa0, 0x37, 0x7e, 0x02, 0x69, 0x94, 0xd2, 0xe4, 0x94, 0x3d, 0xdb, 0xc7,
	0x95, 0xa7, 0xda, 0x73, 0x9e, 0x8c, 0x09, 0x34, 0xfa, 0x6b, 0xf6, 0x35, 0x54, 0xce, 0x08, 0xfc,
	0xcc, 0xcf, 0x5f, 0x41, 0x93, 0xe4, 0x98, 0xf7, 0xfe, 0x9c, 0xb2, 0xaa, 0xf7, 0xe7, 0x94, 0x19,
	0xbf, 0x40, 0xed, 0x94, 0xe1, 0x3d, 0xa8, 0xcd, 0x58, 0x99, 0x2c, 0xf0, 0x73, 0x8a, 0x80, 0x49,
	0x6d, 0xc6, 0xfe, 0xc7, 0x0b, 0x02, 0x31, 0x8a, 0xb3, 0xdc, 0x89, 0x4a, 0x38, 0xc4, 0x07, 0xd0,
	0x08, 0xd6, 0x6c, 0x56, 0x78, 0xd1, 0xba, 0x2a, 0x27, 0xe4, 0x39, 0x90, 0x42, 0x6e, 0xfc, 0x0a,
	0x9a, 0xb9, 0x4a, 0xb3, 0x68, 0x69, 0x46, 0x01, 0x4d, 0xbe, 0x22, 0xb3, 0x7d, 0x10, 0x29, 0x9b,
	0x96, 0xcf, 0xfb, 0x61, 0xb8, 0x5c, 0xcc, 0xb5, 0x01, 0x9d, 0x96, 0xde, 0x1f, 0x69, 0x03, 0x3a,
	0x35, 0x7e, 0x13, 0x41, 0xbd, 0x58, 0x2d, 0xb2, 0xd0, 0x0a, 0xe6, 0x14, 0xef, 0x3e, 0xc8, 0x5b,
	0xca, 0x2f, 0xb0, 0xc8, 0x99, 0x4f, 0x84, 0x78, 0x1a, 0x05, 0xb4, 0x2c, 0x55, 0xb9, 0xc3, 0x3f,
	0x80, 0x1c, 0xb2, 0xcf, 0xd1, 0x8a, 0x05, 0xe5, 0xad, 0xef, 0x70, 0xd2, 0xe6, 0xbc, 0x8e, 0x5d,
	0x28, 0x49, 0x65, 0x85, 0xbb, 0xa0, 0x44, 0xab, 0xac, 0x60, 0x14, 0x63, 0x7f, 0xf7, 0x31, 0x63,
	0x50, 0x6a, 0xc9, 0xc6, 0x6e, 0xef, 0x4f, 0x01, 0xe4, 0xf2, 0x20, 0x7c, 0xfc, 0xe8, 0xef, 0x39,
	0x78, 0xd6, 0x5b, 0xc7, 0x66, 0xf1, 0x2a, 0x7b, 0xf0, 0x1b, 0xb5, 0x1e, 0x55, 0xef, 0x71, 0x63,
	0x16, 0x6d, 0x15, 0x82, 0xba, 0x21, 0xfd, 0x6b, 0x52, 0x5d, 0xf4, 0x6c, 0x17, 0x09, 0xfc, 0xf5,
	0x0d, 0x6d, 0xf7, 0xcc, 0xb1, 0xbc, 0x81, 0x8b, 0x6a, 0xf7, 0xf3, 0x43, 0xe4, 0x0f, 0xf0, 0xa2,
	0x77, 0x89, 0xea, 0x7c, 0x24, 0x5c, 0x8c, 0x1c, 0xcf, 0xe6, 0xbb, 0x46, 0x3e, 0xd1, 0x3c, 0x8b,
	0x20, 0x89, 0x0f, 0x18, 0x62, 0xe5, 0x58, 0xde, 0x3b, 0x04, 0xa5, 0xca, 0x71, 0x13, 0x98, 0xf0,
	0x5f, 0x81, 0x7d, 0x96, 0xf2, 0xff, 0xf7, 0xf8, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0xfd,
	0x32, 0x8c, 0x8b, 0x07, 0x00, 0x00,
}