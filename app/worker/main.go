package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	worker "github.com/evlampiy-lavrentiev/COA-hw/app/worker/core"
	types "github.com/evlampiy-lavrentiev/COA-hw/app/worker/types"
	"github.com/evlampiy-lavrentiev/COA-hw/util"
)

type Worker struct {
	Mode          string
	Serializer    worker.WorkerCore
	Conn          net.PacketConn
	MulticastConn *net.UDPConn
}

func MakeWorker(mode string, port int) *Worker {
	conn, err := net.ListenPacket("udp", mode+":"+strconv.Itoa(port))
	if err != nil {
		log.Panic(err)
	}
	return &Worker{
		Mode:          mode,
		Conn:          conn,
		MulticastConn: util.MakeMulticastUDPConnector(os.Getenv("MULTICAST_ADDRESS")),
	}
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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
	case "proto":
		w.Serializer = worker.ProtoWorkerCore{}
	case "avro":
		w.Serializer = worker.AvroWorkerCore{}
	case "yaml":
		w.Serializer = worker.YamlWorkerCore{}
	case "mpack":
		w.Serializer = worker.MPackWorkerCore{}
	default:
		log.Panic("Not implemented")
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		p := make([]byte, 2048)
		defer wg.Done()
		for {
			n, remoteaddr, err := w.Conn.ReadFrom(p)
			log.Printf("Read message from %v %s", remoteaddr, p[:n])
			if err != nil {
				log.Printf("Error reading from %v", err)
				continue
			}
			bytesResponse := []byte(w.CalcResponce())
			go w.Conn.WriteTo(bytesResponse, remoteaddr)
		}
	}()

	go func() {
		p2 := make([]byte, 2048)
		defer wg.Done()
		for {
			_, remoteaddr, err := w.MulticastConn.ReadFromUDP(p2)

			log.Printf("Get multicast msg from %v", remoteaddr)
			if err != nil {
				log.Printf("Error reading from %v", err)
				continue
			}
			bytesResponse := []byte(w.CalcResponce())
			_, err = w.MulticastConn.WriteToUDP(bytesResponse, remoteaddr)
			if err != nil {
				log.Printf("error sending response to Proxy. Error: %s", err)
			}
		}
	}()

	log.Printf("Started %s!", mode)
	wg.Wait()
}
