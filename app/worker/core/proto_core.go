package worker

import (
	"log"

	pb "github.com/evlampiy-lavrentiev/COA-hw/anek.pb"
	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
	"google.golang.org/protobuf/proto"
)

type ProtoWorkerCore struct {
}

func (nw ProtoWorkerCore) Serialize(anek *types.Anek) []byte {
	protoAnek := anek.ConvertToProto()
	res, err := proto.Marshal(protoAnek)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (nw ProtoWorkerCore) Deserialize(buf []byte) *types.Anek {
	protoRes := pb.Anek{}
	err := proto.Unmarshal(buf, &protoRes)
	if err != nil {
		log.Panic(err)
	}
	return protoRes.ConvertToAnek()
}
