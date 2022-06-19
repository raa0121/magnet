package magnet

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"

	resources "github.com/raa0121/magnet/magnet/internal"
)

const (
	object1LeftUp = 200
	object2LeftUp = 320
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
	Y float64
	ObjectType int `json:"object_type"`
	isHit bool
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
	for i, m := range maps.Maps {
		for j, o := range m.Objects{
			switch o.ObjectType {
			case 1:
				maps.Maps[i].Objects[j].Y = object1LeftUp
			case 2:
				maps.Maps[i].Objects[j].Y = object2LeftUp
			}
		}
	}
	fmt.Printf("%+v\n", maps)
}
