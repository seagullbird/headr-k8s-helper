// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/apimachinery/pkg/runtime/generated.proto

/*
	Package runtime is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/apimachinery/pkg/runtime/generated.proto

	It has these top-level messages:
		RawExtension
		TypeMeta
		Unknown
*/
package runtime

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/ericchiang/k8s/util/intstr"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RawExtension is used to hold extensions in external versions.
//
// To use this, make a field which has RawExtension as its type in your external, versioned
// struct, and Object in your internal struct. You also need to register your
// various plugin types.
//
// // Internal package:
// type MyAPIObject struct {
// 	runtime.TypeMeta `json:",inline"`
// 	MyPlugin runtime.Object `json:"myPlugin"`
// }
// type PluginA struct {
// 	AOption string `json:"aOption"`
// }
//
// // External package:
// type MyAPIObject struct {
// 	runtime.TypeMeta `json:",inline"`
// 	MyPlugin runtime.RawExtension `json:"myPlugin"`
// }
// type PluginA struct {
// 	AOption string `json:"aOption"`
// }
//
// // On the wire, the JSON will look something like this:
// {
// 	"kind":"MyAPIObject",
// 	"apiVersion":"v1",
// 	"myPlugin": {
// 		"kind":"PluginA",
// 		"aOption":"foo",
// 	},
// }
//
// So what happens? Decode first uses json or yaml to unmarshal the serialized data into
// your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked.
// The next step is to copy (using pkg/conversion) into the internal struct. The runtime
// package's DefaultScheme has conversion functions installed which will unpack the
// JSON stored in RawExtension, turning it into the correct object type, and storing it
// in the Object. (TODO: In the case where the object is of an unknown type, a
// runtime.Unknown object will be created and stored.)
//
// +k8s:deepcopy-gen=true
// +protobuf=true
// +k8s:openapi-gen=true
type RawExtension struct {
	// Raw is the underlying serialization of this object.
	//
	// TODO: Determine how to detect ContentType and ContentEncoding of 'Raw' data.
	Raw              []byte `protobuf:"bytes,1,opt,name=raw" json:"raw,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RawExtension) Reset()                    { *m = RawExtension{} }
func (m *RawExtension) String() string            { return proto.CompactTextString(m) }
func (*RawExtension) ProtoMessage()               {}
func (*RawExtension) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *RawExtension) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

// TypeMeta is shared by all top level objects. The proper way to use it is to inline it in your type,
// like this:
// type MyAwesomeAPIObject struct {
//      runtime.TypeMeta    `json:",inline"`
//      ... // other fields
// }
// func (obj *MyAwesomeAPIObject) SetGroupVersionKind(gvk *metav1.GroupVersionKind) { metav1.UpdateTypeMeta(obj,gvk) }; GroupVersionKind() *GroupVersionKind
//
// TypeMeta is provided here for convenience. You may use it directly from this package or define
// your own with the same fields.
//
// +k8s:deepcopy-gen=false
// +protobuf=true
// +k8s:openapi-gen=true
type TypeMeta struct {
	// +optional
	ApiVersion *string `protobuf:"bytes,1,opt,name=apiVersion" json:"apiVersion,omitempty"`
	// +optional
	Kind             *string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TypeMeta) Reset()                    { *m = TypeMeta{} }
func (m *TypeMeta) String() string            { return proto.CompactTextString(m) }
func (*TypeMeta) ProtoMessage()               {}
func (*TypeMeta) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *TypeMeta) GetApiVersion() string {
	if m != nil && m.ApiVersion != nil {
		return *m.ApiVersion
	}
	return ""
}

func (m *TypeMeta) GetKind() string {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return ""
}

// Unknown allows api objects with unknown types to be passed-through. This can be used
// to deal with the API objects from a plug-in. Unknown objects still have functioning
// TypeMeta features-- kind, version, etc.
// TODO: Make this object have easy access to field based accessors and settors for
// metadata and field mutatation.
//
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +protobuf=true
// +k8s:openapi-gen=true
type Unknown struct {
	TypeMeta *TypeMeta `protobuf:"bytes,1,opt,name=typeMeta" json:"typeMeta,omitempty"`
	// Raw will hold the complete serialized object which couldn't be matched
	// with a registered type. Most likely, nothing should be done with this
	// except for passing it through the system.
	Raw []byte `protobuf:"bytes,2,opt,name=raw" json:"raw,omitempty"`
	// ContentEncoding is encoding used to encode 'Raw' data.
	// Unspecified means no encoding.
	ContentEncoding *string `protobuf:"bytes,3,opt,name=contentEncoding" json:"contentEncoding,omitempty"`
	// ContentType  is serialization method used to serialize 'Raw'.
	// Unspecified means ContentTypeJSON.
	ContentType      *string `protobuf:"bytes,4,opt,name=contentType" json:"contentType,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Unknown) Reset()                    { *m = Unknown{} }
func (m *Unknown) String() string            { return proto.CompactTextString(m) }
func (*Unknown) ProtoMessage()               {}
func (*Unknown) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *Unknown) GetTypeMeta() *TypeMeta {
	if m != nil {
		return m.TypeMeta
	}
	return nil
}

func (m *Unknown) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Unknown) GetContentEncoding() string {
	if m != nil && m.ContentEncoding != nil {
		return *m.ContentEncoding
	}
	return ""
}

func (m *Unknown) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func init() {
	proto.RegisterType((*RawExtension)(nil), "k8s.io.apimachinery.pkg.runtime.RawExtension")
	proto.RegisterType((*TypeMeta)(nil), "k8s.io.apimachinery.pkg.runtime.TypeMeta")
	proto.RegisterType((*Unknown)(nil), "k8s.io.apimachinery.pkg.runtime.Unknown")
}
func (m *RawExtension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RawExtension) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Raw != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Raw)))
		i += copy(dAtA[i:], m.Raw)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TypeMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TypeMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ApiVersion != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.ApiVersion)))
		i += copy(dAtA[i:], *m.ApiVersion)
	}
	if m.Kind != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.Kind)))
		i += copy(dAtA[i:], *m.Kind)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Unknown) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unknown) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TypeMeta != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.TypeMeta.Size()))
		n1, err := m.TypeMeta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Raw != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Raw)))
		i += copy(dAtA[i:], m.Raw)
	}
	if m.ContentEncoding != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.ContentEncoding)))
		i += copy(dAtA[i:], *m.ContentEncoding)
	}
	if m.ContentType != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.ContentType)))
		i += copy(dAtA[i:], *m.ContentType)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RawExtension) Size() (n int) {
	var l int
	_ = l
	if m.Raw != nil {
		l = len(m.Raw)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TypeMeta) Size() (n int) {
	var l int
	_ = l
	if m.ApiVersion != nil {
		l = len(*m.ApiVersion)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Kind != nil {
		l = len(*m.Kind)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Unknown) Size() (n int) {
	var l int
	_ = l
	if m.TypeMeta != nil {
		l = m.TypeMeta.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Raw != nil {
		l = len(m.Raw)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.ContentEncoding != nil {
		l = len(*m.ContentEncoding)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.ContentType != nil {
		l = len(*m.ContentType)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RawExtension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RawExtension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawExtension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Raw = append(m.Raw[:0], dAtA[iNdEx:postIndex]...)
			if m.Raw == nil {
				m.Raw = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TypeMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TypeMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TypeMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ApiVersion = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Kind = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Unknown) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Unknown: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unknown: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TypeMeta == nil {
				m.TypeMeta = &TypeMeta{}
			}
			if err := m.TypeMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Raw = append(m.Raw[:0], dAtA[iNdEx:postIndex]...)
			if m.Raw == nil {
				m.Raw = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentEncoding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ContentEncoding = &s
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ContentType = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/apimachinery/pkg/runtime/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x7f, 0xb7, 0x95, 0xda, 0xff, 0xb6, 0x12, 0xc8, 0x53, 0x60, 0x08, 0x55, 0xa6, 0xb2,
	0xd8, 0x12, 0x2c, 0x4c, 0x0c, 0x48, 0x19, 0x59, 0x22, 0x60, 0x60, 0xb3, 0x12, 0x2b, 0x58, 0xa1,
	0xd7, 0x96, 0x73, 0xab, 0xd0, 0x37, 0xe1, 0x29, 0x78, 0x0e, 0x46, 0x1e, 0x01, 0x85, 0x17, 0x41,
	0x31, 0x69, 0x15, 0x55, 0x42, 0x6c, 0x47, 0xc7, 0xe7, 0x3b, 0x3e, 0x17, 0x64, 0x75, 0x55, 0x0b,
	0x63, 0xa5, 0x72, 0x66, 0xad, 0xf2, 0x27, 0x83, 0xda, 0x6f, 0xa5, 0xab, 0x4a, 0xe9, 0x37, 0x48,
	0x66, 0xad, 0x65, 0xa9, 0x51, 0x7b, 0x45, 0xba, 0x10, 0xce, 0x5b, 0xb2, 0xfc, 0xec, 0x07, 0x10,
	0x43, 0x40, 0xb8, 0xaa, 0x14, 0x3d, 0x70, 0x7a, 0xf9, 0x5b, 0xe3, 0x86, 0xcc, 0xb3, 0x34, 0x48,
	0x35, 0xf9, 0xc3, 0xd6, 0x64, 0x09, 0x8b, 0x4c, 0x35, 0xe9, 0x0b, 0x69, 0xac, 0x8d, 0x45, 0x7e,
	0x0c, 0x63, 0xaf, 0x9a, 0x88, 0x2d, 0xd9, 0x6a, 0x91, 0x75, 0x32, 0xb9, 0x86, 0xd9, 0xdd, 0xd6,
	0xe9, 0x5b, 0x4d, 0x8a, 0xc7, 0x00, 0xca, 0x99, 0x07, 0xed, 0xbb, 0x6c, 0x08, 0xfd, 0xcf, 0x06,
	0x0e, 0xe7, 0x30, 0xa9, 0x0c, 0x16, 0xd1, 0x28, 0xbc, 0x04, 0x9d, 0xbc, 0x31, 0x98, 0xde, 0x63,
	0x85, 0xb6, 0x41, 0x9e, 0xc2, 0x8c, 0xfa, 0xae, 0x40, 0xcf, 0x2f, 0xce, 0xc5, 0x1f, 0x67, 0x89,
	0xdd, 0xe7, 0xd9, 0x1e, 0xdd, 0x8d, 0x1c, 0xed, 0x47, 0xf2, 0x15, 0x1c, 0xe5, 0x16, 0x49, 0x23,
	0xa5, 0x98, 0xdb, 0xc2, 0x60, 0x19, 0x8d, 0xc3, 0x86, 0x43, 0x9b, 0x2f, 0x61, 0xde, 0x5b, 0x5d,
	0x71, 0x34, 0x09, 0xa9, 0xa1, 0x75, 0x73, 0xf2, 0xde, 0xc6, 0xec, 0xa3, 0x8d, 0xd9, 0x67, 0x1b,
	0xb3, 0xd7, 0xaf, 0xf8, 0xdf, 0xe3, 0xb4, 0xdf, 0xf2, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x60, 0x1f,
	0x94, 0x77, 0xb5, 0x01, 0x00, 0x00,
}
