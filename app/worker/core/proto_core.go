package worker

import (
	"log"

	// pb "github.com/evlampiy-lavrentiev/COA-hw/anek.pb"
	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
	"github.com/golang/protobuf/proto"
	// "google.golang.org/protobuf/proto"
)

type ProtoWorkerCore struct {
}

func (nw ProtoWorkerCore) Serialize(anek *types.Anek) []byte {
	// protoAnek := anek.ConvertToProto()
	// res, err := proto.Marshal(protoAnek)
	res, err := proto.Marshal(anek)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (nw ProtoWorkerCore) Deserialize(buf []byte) *types.Anek {
	// protoRes := pb.Anek{}
	// err := proto.Unmarshal(buf, &protoRes)
	res := &types.Anek{}
	err := proto.Unmarshal(buf, res)
	if err != nil {
		log.Panic(err)
	}
	return res
}
