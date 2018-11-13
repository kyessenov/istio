// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ditto/config/config.proto

/*
	Package config is a generated protocol buffer package.

	It is generated from these files:
		ditto/config/config.proto

	It has these top-level messages:
		Requirement
		Endpoint
		APIRequirement
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Requirement struct {
	Actions  []string `protobuf:"bytes,1,rep,name=actions" json:"actions,omitempty"`
	Resource string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (m *Requirement) Reset()                    { *m = Requirement{} }
func (*Requirement) ProtoMessage()               {}
func (*Requirement) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Requirement) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Requirement) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type Endpoint struct {
	Endpoint     string         `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Method       string         `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Name         string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Requirements []*Requirement `protobuf:"bytes,4,rep,name=requirements" json:"requirements,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *Endpoint) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Endpoint) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Endpoint) GetRequirements() []*Requirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

type APIRequirement struct {
	Endpoints []*Endpoint `protobuf:"bytes,1,rep,name=endpoints" json:"endpoints,omitempty"`
	Version   string      `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *APIRequirement) Reset()                    { *m = APIRequirement{} }
func (*APIRequirement) ProtoMessage()               {}
func (*APIRequirement) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *APIRequirement) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *APIRequirement) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type Params struct {
	ApiRequirements []*APIRequirement `protobuf:"bytes,1,rep,name=api_requirements,json=apiRequirements" json:"api_requirements,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

func (m *Params) GetApiRequirements() []*APIRequirement {
	if m != nil {
		return m.ApiRequirements
	}
	return nil
}

func init() {
	proto.RegisterType((*Requirement)(nil), "config.Requirement")
	proto.RegisterType((*Endpoint)(nil), "config.Endpoint")
	proto.RegisterType((*APIRequirement)(nil), "config.APIRequirement")
	proto.RegisterType((*Params)(nil), "config.Params")
}
func (this *Requirement) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Requirement)
	if !ok {
		that2, ok := that.(Requirement)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Actions) != len(that1.Actions) {
		return false
	}
	for i := range this.Actions {
		if this.Actions[i] != that1.Actions[i] {
			return false
		}
	}
	if this.Resource != that1.Resource {
		return false
	}
	return true
}
func (this *Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint)
	if !ok {
		that2, ok := that.(Endpoint)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Endpoint != that1.Endpoint {
		return false
	}
	if this.Method != that1.Method {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if len(this.Requirements) != len(that1.Requirements) {
		return false
	}
	for i := range this.Requirements {
		if !this.Requirements[i].Equal(that1.Requirements[i]) {
			return false
		}
	}
	return true
}
func (this *APIRequirement) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*APIRequirement)
	if !ok {
		that2, ok := that.(APIRequirement)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Endpoints) != len(that1.Endpoints) {
		return false
	}
	for i := range this.Endpoints {
		if !this.Endpoints[i].Equal(that1.Endpoints[i]) {
			return false
		}
	}
	if this.Version != that1.Version {
		return false
	}
	return true
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ApiRequirements) != len(that1.ApiRequirements) {
		return false
	}
	for i := range this.ApiRequirements {
		if !this.ApiRequirements[i].Equal(that1.ApiRequirements[i]) {
			return false
		}
	}
	return true
}
func (this *Requirement) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&config.Requirement{")
	s = append(s, "Actions: "+fmt.Sprintf("%#v", this.Actions)+",\n")
	s = append(s, "Resource: "+fmt.Sprintf("%#v", this.Resource)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Endpoint) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&config.Endpoint{")
	s = append(s, "Endpoint: "+fmt.Sprintf("%#v", this.Endpoint)+",\n")
	s = append(s, "Method: "+fmt.Sprintf("%#v", this.Method)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	if this.Requirements != nil {
		s = append(s, "Requirements: "+fmt.Sprintf("%#v", this.Requirements)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *APIRequirement) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&config.APIRequirement{")
	if this.Endpoints != nil {
		s = append(s, "Endpoints: "+fmt.Sprintf("%#v", this.Endpoints)+",\n")
	}
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Params) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&config.Params{")
	if this.ApiRequirements != nil {
		s = append(s, "ApiRequirements: "+fmt.Sprintf("%#v", this.ApiRequirements)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringConfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Requirement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Requirement) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Actions) > 0 {
		for _, s := range m.Actions {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Resource) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Resource)))
		i += copy(dAtA[i:], m.Resource)
	}
	return i, nil
}

func (m *Endpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Endpoint) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Endpoint) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Endpoint)))
		i += copy(dAtA[i:], m.Endpoint)
	}
	if len(m.Method) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Method)))
		i += copy(dAtA[i:], m.Method)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Requirements) > 0 {
		for _, msg := range m.Requirements {
			dAtA[i] = 0x22
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *APIRequirement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *APIRequirement) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Endpoints) > 0 {
		for _, msg := range m.Endpoints {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Version) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	return i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ApiRequirements) > 0 {
		for _, msg := range m.ApiRequirements {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Requirement) Size() (n int) {
	var l int
	_ = l
	if len(m.Actions) > 0 {
		for _, s := range m.Actions {
			l = len(s)
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = len(m.Resource)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *Endpoint) Size() (n int) {
	var l int
	_ = l
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.Requirements) > 0 {
		for _, e := range m.Requirements {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func (m *APIRequirement) Size() (n int) {
	var l int
	_ = l
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	var l int
	_ = l
	if len(m.ApiRequirements) > 0 {
		for _, e := range m.ApiRequirements {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Requirement) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Requirement{`,
		`Actions:` + fmt.Sprintf("%v", this.Actions) + `,`,
		`Resource:` + fmt.Sprintf("%v", this.Resource) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Endpoint) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Endpoint{`,
		`Endpoint:` + fmt.Sprintf("%v", this.Endpoint) + `,`,
		`Method:` + fmt.Sprintf("%v", this.Method) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Requirements:` + strings.Replace(fmt.Sprintf("%v", this.Requirements), "Requirement", "Requirement", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *APIRequirement) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&APIRequirement{`,
		`Endpoints:` + strings.Replace(fmt.Sprintf("%v", this.Endpoints), "Endpoint", "Endpoint", 1) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`ApiRequirements:` + strings.Replace(fmt.Sprintf("%v", this.ApiRequirements), "APIRequirement", "APIRequirement", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Requirement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Requirement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Requirement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resource = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Endpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Endpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Endpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requirements = append(m.Requirements, &Requirement{})
			if err := m.Requirements[len(m.Requirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *APIRequirement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: APIRequirement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIRequirement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, &Endpoint{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiRequirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiRequirements = append(m.ApiRequirements, &APIRequirement{})
			if err := m.ApiRequirements[len(m.ApiRequirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ditto/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x31, 0x4f, 0x02, 0x31,
	0x18, 0x86, 0xaf, 0x42, 0x4e, 0xf8, 0x30, 0x4a, 0x6a, 0x42, 0xaa, 0x43, 0x43, 0x6e, 0x62, 0x30,
	0x98, 0xe0, 0xe0, 0x8c, 0xc6, 0xc1, 0xb8, 0x90, 0x8e, 0x2e, 0xa6, 0x42, 0xd5, 0x0e, 0xd7, 0x9e,
	0x6d, 0x71, 0x76, 0x77, 0xf1, 0x67, 0xf8, 0x53, 0x1c, 0x19, 0x1d, 0xa5, 0x2e, 0x8e, 0xfc, 0x04,
	0x03, 0xd7, 0xca, 0xdd, 0xd4, 0xef, 0x69, 0xde, 0xbc, 0x79, 0xbe, 0x16, 0x8e, 0x66, 0xd2, 0x39,
	0x7d, 0x3a, 0xd5, 0xea, 0x41, 0x3e, 0x86, 0x63, 0x58, 0x18, 0xed, 0x34, 0x4e, 0x4b, 0xca, 0x2e,
	0xa1, 0xc3, 0xc4, 0xf3, 0x5c, 0x1a, 0x91, 0x0b, 0xe5, 0x30, 0x81, 0x5d, 0x3e, 0x75, 0x52, 0x2b,
	0x4b, 0x50, 0xbf, 0x31, 0x68, 0xb3, 0x88, 0xf8, 0x18, 0x5a, 0x46, 0x58, 0x3d, 0x37, 0x53, 0x41,
	0x76, 0xfa, 0x68, 0xd0, 0x66, 0xff, 0x9c, 0xbd, 0x21, 0x68, 0x5d, 0xa9, 0x59, 0xa1, 0xa5, 0x72,
	0xeb, 0xa0, 0x08, 0x33, 0x41, 0x65, 0x30, 0x32, 0xee, 0x41, 0x9a, 0x0b, 0xf7, 0xa4, 0x67, 0xa1,
	0x22, 0x10, 0xc6, 0xd0, 0x54, 0x3c, 0x17, 0xa4, 0xb1, 0xb9, 0xdd, 0xcc, 0xf8, 0x1c, 0xf6, 0xcc,
	0xd6, 0xcc, 0x92, 0x66, 0xbf, 0x31, 0xe8, 0x8c, 0x0e, 0x87, 0x61, 0x8d, 0x8a, 0x35, 0xab, 0x05,
	0xb3, 0x5b, 0xd8, 0x1f, 0x4f, 0xae, 0xab, 0x5b, 0x0d, 0xa1, 0x1d, 0x15, 0xca, 0xbd, 0x3a, 0xa3,
	0x6e, 0xec, 0x89, 0xde, 0x6c, 0x1b, 0x59, 0xbf, 0xc2, 0x8b, 0x30, 0x56, 0x6a, 0x15, 0x3c, 0x23,
	0x66, 0x37, 0x90, 0x4e, 0xb8, 0xe1, 0xb9, 0xc5, 0x63, 0xe8, 0xf2, 0x42, 0xde, 0xd5, 0x14, 0xcb,
	0xea, 0x5e, 0xac, 0xae, 0x5b, 0xb0, 0x03, 0x5e, 0xc8, 0x0a, 0xdb, 0x8b, 0x93, 0xc5, 0x92, 0x26,
	0x5f, 0x4b, 0x9a, 0xac, 0x96, 0x14, 0xbd, 0x7a, 0x8a, 0x3e, 0x3c, 0x45, 0x9f, 0x9e, 0xa2, 0x85,
	0xa7, 0xe8, 0xdb, 0x53, 0xf4, 0xeb, 0x69, 0xb2, 0xf2, 0x14, 0xbd, 0xff, 0xd0, 0xe4, 0x3e, 0xdd,
	0x7c, 0xdc, 0xd9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xc5, 0x02, 0xb7, 0xd5, 0x01, 0x00,
	0x00,
}
