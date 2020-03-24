package main

import (
	"fmt"
	"io/ioutil"
	"log"

	enumpb "./src/enum_example"
	simplepb "./src/simple"
	"github.com/golang/protobuf/proto"
)

// import "C:/Users/abhis/Desktop/goPrac/protbuff-go/src/simple"
func main() {
	readAndWriteDemo()
	doEnum()
}
func doEnum() {
	em := enumpb.EnumMessage{
		Id:           43,
		DayOfTheWeek: enumpb.DayOfTheweek_MONDAY,
	}
	fmt.Println(em)
}
func doSimple() *simplepb.SimpleMessage {
	ds := simplepb.SimpleMessage{
		Id:         123,
		IsSample:   true,
		Name:       "Parik",
		SampleList: []int32{1, 3, 4},
	}

	fmt.Println(ds)
	return &ds
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Unable to convert the data into byte for writing", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("unable to write file", err)
		return err
	}
	fmt.Println("Data saved")
	return nil
}

func readFromFile(fanme string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fanme)
	if err != nil {
		log.Fatalln("UNable to read file ", err)
		return err
	}
	err = proto.Unmarshal(in, pb) // or &simplepb.SimpleMessage{}
	if err != nil {
		log.Fatalln("Unable to unmarshal data ", err)
	}
	return nil
}

func readAndWriteDemo() {
	pb := doSimple()
	writeToFile("simple.bin", pb)
	pb2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", pb2)
	fmt.Println(pb2)
}
