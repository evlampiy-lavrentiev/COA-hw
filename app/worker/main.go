package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	worker "github.com/evlampiy-lavrentiev/COA-hw/app/worker/core"
)

type Worker struct {
	Serializer worker.WorkerCore
	Conn       net.PacketConn
}

func MakeWorker(mode string, port int) *Worker {
	conn, err := net.ListenPacket("udp", mode+":"+strconv.Itoa(port))
	if err != nil {
		log.Panic(err)
	}
	return &Worker{Conn: conn}
}

func (w *Worker) CalcResponce() string {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		bytes, err := w.Serializer.Serialize(person)
	}
		totalTimeSerialize += time.Since(start).Microseconds()
		if err != nil {
			return "", fmt.Errorf("failed to serialize string: %v", err)
		}
		totalStructSize += len(bytes)

		start = time.Now()
		_, err = converter.Deserialize(bytes)
		totalTimeDeserialize += time.Since(start).Microseconds()
		if err != nil {
			return "", fmt.Errorf("failed to deserialize string: %v", err)
		}
	}
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
		bytesResponse := []byte(w.Serializer.FetchResult())
		go w.Conn.WriteTo(bytesResponse, remoteaddr)
	}
}
