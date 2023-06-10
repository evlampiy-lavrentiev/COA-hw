package worker

import (
	"log"

	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
	"github.com/vmihailenco/msgpack"
)

type MPackWorkerCore struct {
}

func (nw MPackWorkerCore) Serialize(anek *types.Anek) []byte {
	res, err := msgpack.Marshal(anek)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (nw MPackWorkerCore) Deserialize(buf []byte) *types.Anek {
	res := types.Anek{}
	err := msgpack.Unmarshal(buf, &res)
	if err != nil {
		log.Panic(err)
	}
	return &res
}
