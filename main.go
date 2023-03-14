package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	logpb "github.com/grcdeepak1/parse/proto/log"
	spb "github.com/openconfig/gribi/v1/proto/service"
)

func main() {
	gribi_data, err := ioutil.ReadFile("gribi_trace.pb")
	if err != nil {
		log.Fatal(err)
	}
	m := &logpb.Events{}
	err = proto.Unmarshal(gribi_data, m)
	if err != nil {
		fmt.Println("Fatal")
		log.Fatal(err)
	}

	fileName := "gribi_trace.txtpb"
	// Check if file exists
	if _, err := os.Stat(fileName); err == nil {
		// If file exists, delete it
		err := os.Remove(fileName)
		if err != nil {
			panic(err)
		}
	}
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, event := range m.GrpcEvents {
		m1 := &spb.ModifyRequest{}
		proto.Unmarshal(event.GetMessage().GetData(), m1)
		fmt.Println(m1)
		jsonString, err := json.Marshal(m1)
		if err != nil {
			panic(err)
		}
		_, err = f.WriteString(string(jsonString))
		if err != nil {
			log.Fatal(err)
		}
	}

}
