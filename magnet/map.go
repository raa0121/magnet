package magnet

import (
	"encoding/json"
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
	collisionLeftUp, collisionRightDown Point
	positionX float64
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
				maps.Maps[i].Objects[j].collisionLeftUp = Point{20, 20}
				maps.Maps[i].Objects[j].collisionRightDown = Point{225, 245}
				maps.Maps[i].Objects[j].positionX = 0.0
			case 2:
				maps.Maps[i].Objects[j].Y = object2LeftUp
				maps.Maps[i].Objects[j].collisionLeftUp = Point{0, 100}
				maps.Maps[i].Objects[j].collisionRightDown = Point{256, 256}
				maps.Maps[i].Objects[j].positionX = 0.0
			}
		}
	}
}
