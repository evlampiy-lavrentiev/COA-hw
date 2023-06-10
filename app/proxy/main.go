package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/evlampiy-lavrentiev/COA-hw/util"
)

type Proxy struct {
	ResponseChan chan string
	Connections  map[string]*net.Conn
}

func MakeProxy(configPath string) *Proxy {

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Panicf("Error in config(%s) reading: %s", configPath, err)
	}
	var config map[string]int

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Panicf("Error in JSON Unmarshal: %v", err)
	}

	proxy := Proxy{
		Connections:  map[string]*net.Conn{},
		ResponseChan: make(chan string, 100)}
	for format, port := range config {
		log.Printf("make connector %s %v", format, port)
		proxy.Connections[format] = util.MakeDialConnector(format, port)
	}

	return &proxy
}

func (s Proxy) WorkersListen() {
	for mode, conn := range s.Connections {
		go func(mode string, conn *net.Conn) {
			buf := make([]byte, 65507)
			for {
				size, err := (*conn).Read(buf)
				if err != nil {
					log.Panicf("Error while reading worker response: %s", err)
				}
				log.Printf("get response from %s", mode)
				s.ResponseChan <- string(buf[:size])
			}
		}(mode, conn)
	}
}

func (s *Proxy) handleGetResult(w http.ResponseWriter, r *http.Request) {
	format := r.URL.Query().Get("format")

	workers_cnt := 1
	switch format {
	case "all":
		workers_cnt = len(s.Connections)
		log.Fatalf("not impl YET")
	default:
		conn, ok := s.Connections[format]
		if ok {
			fmt.Fprintf(*conn, "test your jokes pls")
		} else {
			log.Printf("Got not implemented format: %s", format)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	var answers []string
	for i := 0; i < workers_cnt; i++ {
		answers = append(answers, <-s.ResponseChan)
	}
	fmt.Fprintf(w, "Results:\n%s\n", strings.Join(answers, "\n"))
}

func main() {
	s := MakeProxy("config.json")

	http.HandleFunc("/get_result/", s.handleGetResult)

	go s.WorkersListen()

	log.Print("Proxy started!")
	log.Fatal(http.ListenAndServe("0.0.0.0:2000", nil))
}
