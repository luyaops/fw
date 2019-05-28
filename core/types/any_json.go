package types

import (
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

type DynamicMessage struct {
	RawJson []byte `protobuf:"bytes,1,opt,name=rawJson"`
}

func (m *DynamicMessage) Reset() {
	m.RawJson = []byte("{}")
}

func (m *DynamicMessage) String() string {
	return string(m.RawJson)
}

func (m *DynamicMessage) ProtoMessage() {
}

func (m *DynamicMessage) MarshalJSONPB(jm *jsonpb.Marshaler) ([]byte, error) {
	return m.RawJson, nil
}

func (m *DynamicMessage) UnmarshalJSONPB(jum *jsonpb.Unmarshaler, js []byte) error {
	m.RawJson = js
	return nil
}

func init() {
	proto.RegisterType((*DynamicMessage)(nil), "shcmp.protobuf.dynamicMessage")
}
