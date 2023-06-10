package worker

import (
	"log"

	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
	"github.com/linkedin/goavro"
)

const schemaJSON = `
{
	"type": "record",
	"name": "Struct",
	"fields": [
		{"name": "Str", "type": "string"},
        {"name": "Arr", "type": {"type": "array", "items": "int"}},
        {"name": "Dict", "type": {"type": "map", "values": "string"}},
        {"name": "Int", "type": "int"},
        {"name": "Float", "type": "double"}
	]
}`

type AvroWorkerCore struct {
}

func (nw AvroWorkerCore) Serialize(anek *types.Anek) []byte {
	schema, err := goavro.NewCodec(schemaJSON)
	if err != nil {
		log.Panic(err)
	}
	res, err := schema.BinaryFromNative(nil, anek)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (nw AvroWorkerCore) Deserialize(buf []byte) *types.Anek {
	schema, err := goavro.NewCodec(schemaJSON)
	if err != nil {
		log.Panic(err)
	}
	res, _, err := schema.NativeFromBinary(buf)
	if err != nil {
		log.Panic(err)
	}

	return res.(*types.Anek)
}
