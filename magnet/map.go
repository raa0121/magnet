package magnet

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"

	resources "github.com/raa0121/magnet/magnet/internal"
)

var maps Maps

type Maps struct {
	Maps []Map `json:"maps"`
}

type Map struct {
	Objects []Object `json:"objects"`
}

type Object struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	ObjectType int `json:"object_type"`
}

func init() {
	mapInit()
}
func mapInit() {
	b, err := fs.ReadFile(resources.Embed, "maps.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &maps)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", maps)
}
