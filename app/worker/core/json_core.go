package worker

import (
	"encoding/xml"
	"log"

	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
)

type JsonWorkerCore struct {
}

func (nw JsonWorkerCore) Serialize(anek *types.Anek) []byte {
	res, err := xml.Marshal(anek)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (nw JsonWorkerCore) Deserialize(buf []byte) *types.Anek {
	res := types.Anek{}
	err := xml.Unmarshal(buf, &res)
	if err != nil {
		log.Panic(err)
	}
	return &res
}
