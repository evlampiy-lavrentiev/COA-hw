package worker

import (
	"log"

	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
	"github.com/hamba/avro"
)

const schemaJSON = `
{
	"type": "record",
	"name": "Anek",
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
	schema, _ := avro.Parse(schemaJSON)
	res, err := avro.Marshal(schema, anek)

	if err != nil {
		log.Panic(err)
	}
	return res
}

func (nw AvroWorkerCore) Deserialize(buf []byte) *types.Anek {
	res := &types.Anek{}
	schema, _ := avro.Parse(schemaJSON)
	err := avro.Unmarshal(schema, buf, res)
	if err != nil {
		log.Panic(err)
	}
	return res
}
