package types

import (
	"encoding/xml"
	"io"

	"github.com/golang/protobuf/proto"
)

type StringMap map[string]string

type Anek struct {
	Str   string    `protobuf:"bytes,1,opt,name=str,proto3"`
	Int   int32     `protobuf:"varint,2,opt,name=int,proto3"`
	Arr   []int32   `protobuf:"varint,3,rep,name=arr,proto3"`
	Dict  StringMap `protobuf:"bytes,4,rep,name=dict,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Float float64   `protobuf:"fixed64,5,opt,name=float,proto3"`
}

func MakeAnek() *Anek {
	return &Anek{
		Str: `Знаете почему меня называют на работе 007?
0 - желаний работать
0 - мотивации
7 - перекуров за час`,
		Int: 228,
		Arr: []int32{1, 3, 3, 7},
		Dict: map[string]string{
			"Rzaka": "9-10",
			"Smysl": "5-6",
			"Ziza":  "10000",
		},
		Float: 3.141592653589793238462643383279}
}

// https://stackoverflow.com/questions/30928770/marshall-map-to-xml-in-go
type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = StringMap{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

func (anek *Anek) Reset() {
	*anek = Anek{}
}
func (anek *Anek) String() string {
	return proto.CompactTextString(anek)
}
func (_ *Anek) ProtoMessage() {
}
