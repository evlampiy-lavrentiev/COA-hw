package types

import (
	"encoding/xml"
	"io"

	pb "github.com/evlampiy-lavrentiev/COA-hw/anek.pb"
)

type StringMap map[string]string

type Anek struct {
	Str   string
	Int   int
	Arr   []int
	Dict  StringMap
	Float float64
}

func MakeAnek() *Anek {
	return &Anek{
		Str: `Знаете почему меня называют на работе 007?
0 - желаний работать
0 - мотивации
7 - перекуров за час`,
		Int: 228,
		Arr: []int{1, 3, 3, 7},
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

func (anek *Anek) ConvertToProto() *pb.Anek {
	return &pb.Anek{
		Str:   anek.Str,
		Int:   int64(anek.Int),
		Arr:   anek.Arr,
		Dict:  anek.Dict,
		Float: anek.Float,
	}
}

func (protoAnek *pb.Anek) ConvertToAnek(pb *Anek) *Anek {
	return &Anek{
		Str:   protoAnek.Str,
		Int:   protoAnek.Int,
		Arr:   protoAnek.Arr,
		Dict:  protoAnek.Dict,
		Float: protoAnek.Float,
	}
}
