package worker

import (
	"encoding/xml"
	"log"

	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
)

type XmlWorkerCore struct {
}

func (nw XmlWorkerCore) Serialize(anek *types.Anek) []byte {
	res, err := xml.Marshal(anek)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (nw XmlWorkerCore) Deserialize(buf []byte) *types.Anek {
	res := types.Anek{}
	err := xml.Unmarshal(buf, &res)
	if err != nil {
		log.Panic(err)
	}
	return &res
}
