package worker

type WorkerCore interface {
	Serialize(anek *Anek) []byte
	Deserialize([]byte) *Anek
}

type NativeWorkerCore struct {
}

func (w NativeWorkerCore) Serialize(anek *Anek) []byte {

}

func (nw NativeWorkerCore) Deserialize([]byte) *Anek {
	return "Результат выполнения NativeWorkerCore"
}

type XmlWorkerCore struct {
}

func (nw XmlWorkerCore) Serialize(anek *Anek) []byte {

}

func (nw XmlWorkerCore) Deserialize([]byte) *Anek {
	return "Результат выполнения XmlWorkerCore"
}
