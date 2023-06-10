package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	worker "github.com/evlampiy-lavrentiev/COA-hw/app/worker/core"
	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
)

type Worker struct {
	Mode       string
	Serializer worker.WorkerCore
	Conn       net.PacketConn
}

func MakeWorker(mode string, port int) *Worker {
	conn, err := net.ListenPacket("udp", mode+":"+strconv.Itoa(port))
	if err != nil {
		log.Panic(err)
	}
	return &Worker{
		Mode: mode,
		Conn: conn}
}

func (w *Worker) CalcResponce() string {
	anek := types.MakeAnek()
	bytes := make([]byte, 0)

	start := time.Now()
	for i := 0; i < 1000; i++ {
		bytes = w.Serializer.Serialize(anek)
	}
	ser_ms := float64(time.Since(start).Nanoseconds()) / 1e6

	start = time.Now()
	for i := 0; i < 1000; i++ {
		_ = w.Serializer.Deserialize(bytes)
	}
	deser_ms := float64(time.Since(start).Nanoseconds()) / 1e6

	return fmt.Sprintf("%s - %v - %vus - %vus", w.Mode, len(bytes), ser_ms, deser_ms)
}

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	mode := os.Args[2]

	w := MakeWorker(mode, port)

	switch mode {
	case "native":
		w.Serializer = worker.NativeWorkerCore{}
	case "xml":
		w.Serializer = worker.XmlWorkerCore{}
	case "json":
		w.Serializer = worker.JsonWorkerCore{}
	default:
		log.Panic("Not implemented")
	}

	p := make([]byte, 2048)
	log.Printf("Start %s worker loop at port=%v", mode, port)

	for {
		n, remoteaddr, err := w.Conn.ReadFrom(p)
		log.Printf("Read a message from %v %s", remoteaddr, p[:n])
		if err != nil {
			log.Printf("Error reading from %v", err)
			continue
		}
		bytesResponse := []byte(w.CalcResponce())
		go w.Conn.WriteTo(bytesResponse, remoteaddr)
	}
}
