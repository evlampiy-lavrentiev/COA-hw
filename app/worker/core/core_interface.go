package worker

import (
	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
)

type WorkerCore interface {
	Serialize(anek *types.Anek) []byte
	Deserialize([]byte) *types.Anek
}
