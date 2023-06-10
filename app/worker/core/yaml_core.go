package worker

import (
	"log"

	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
	"gopkg.in/yaml.v3"
)

type YamlWorkerCore struct {
}

func (nw YamlWorkerCore) Serialize(anek *types.Anek) []byte {
	res, err := yaml.Marshal(anek)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (nw YamlWorkerCore) Deserialize(buf []byte) *types.Anek {
	res := types.Anek{}
	err := yaml.Unmarshal(buf, &res)
	if err != nil {
		log.Panic(err)
	}
	return &res
}
