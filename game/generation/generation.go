// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate go run github.com/googleforgames/space-agon/game/generation
//go:generate go fmt github.com/googleforgames/space-agon/game
package main

import (
	"fmt"
	"os"
	"text/template"
)

var components = map[string]string{
	// "ExplosionDetails": "ExplosionDetails",
	"Lookup":         "Lookup",
	"MissileDetails": "MissileDetails",
	"Momentum":       "Vec2",
	"NetworkId":      "uint64",
	"Pos":            "Vec2",
	"Rot":            "float32",
	"ShipControl":    "ShipControl",
	// "SpawnEvent":       "SpawnType",
	"Spin":         "float32",
	"Sprite":       "Sprite",
	"TimedDestroy": "float32",
	"TimedExplode": "float32",

	"AffectedByGravity": "",
	"BoundLocation":     "",
	"CanExplode":        "",
	"FrameEndDelete":    "",
	"KeepInCamera":      "",
	"NetworkReceive":    "",
	"NetworkTransmit":   "",
	"ParticleSunDelete": "",
	"PointRender":       "",
}

var typeLiterals = map[string]string{
	// "ExplosionDetails": "ExplosionDetails{Initialized: false}",
	"float32":   "0",
	"uint64":    "0",
	"Lookup":    "<Lookup is special, this should be never invoked>",
	"SpawnType": "0",
	"Sprite":    "SpriteUnset",
}

func main() {
	fmt.Println("Starting generation")

	t := template.Must(template.ParseFiles("components.gotemplate"))
	f, err := os.Create("../components.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = t.Execute(f, getInfo())
	if err != nil {
		panic(err)
	}
	fmt.Println("Generation Complete")
}

type compType struct {
	Name                    string
	TrueType                string
	TypeLiteral             string
	GenerateTypeDeclaration bool
	ReturnPointer           bool
}

type Info struct {
	CompTypes map[string]*compType
	Comps     map[string]*compType
}

func getInfo() *Info {
	// Known types start with those defined in the file because they don't fit the
	// standard pattern.

	i := &Info{
		CompTypes: make(map[string]*compType),
		Comps:     make(map[string]*compType),
	}

	for k, v := range components {
		if v == "" {
			i.Comps[k] = nil
			continue
		}
		if _, ok := i.CompTypes[v]; !ok {
			c := &compType{
				Name:                    "comp_" + v,
				TrueType:                v,
				GenerateTypeDeclaration: true,
				ReturnPointer:           true,
			}

			if literal, ok := typeLiterals[v]; ok {
				c.TypeLiteral = literal
			} else {
				c.TypeLiteral = v + "{}"
			}

			i.CompTypes[v] = c
		}

		i.Comps[k] = i.CompTypes[v]
	}

	i.CompTypes["Lookup"].GenerateTypeDeclaration = false
	i.CompTypes["Lookup"].ReturnPointer = false

	return i
}
