package worker

import (
	"bytes"
	"encoding/gob"

	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
)

type NativeWorkerCore struct {
}

func (w NativeWorkerCore) Serialize(anek *types.Anek) []byte {
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	encoder.Encode(anek)
	return buffer.Bytes()
}

func (nw NativeWorkerCore) Deserialize(buf []byte) *types.Anek {
	res := types.Anek{}
	buffer := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(buffer)
	decoder.Decode(&res)
	return &res
}
