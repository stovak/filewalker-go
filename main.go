package main

import (
	"fmt"
	"log"

	"github.com/stovak/filewalker/pkg/util"
)

func main() {
	origMedia, err := util.GetListOfFilesFromDirectory(".mov", "/Volumes/One Touch/Dog Park Adjacent/original Media")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", origMedia)

	udids, err := idevice.UDIDList()
	if err != nil {
		panic(err)
	}

	for _, udid := range udids {
		fmt.Println(udid)
	}

}
