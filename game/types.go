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

package game

import (
	"math"

	"github.com/googleforgames/space-agon/game/pb"
)

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
// Component types
////////////////////////////////////////////////////////////////////////////////

type Sprite uint16

const (
	SpriteUnset = Sprite(iota)
	SpriteShip
	SpriteEnemyShip
	SpriteMissile
	SpriteStar
	SpriteStarBit
	SpriteExplosionFlash
)

type Vec2 [2]float32

func (v *Vec2) ToProto() *pb.Vec2 {
	return &pb.Vec2{
		X: v[0],
		Y: v[1],
	}
}

func Vec2FromProto(p *pb.Vec2) Vec2 {
	return Vec2{p.X, p.Y}
}

func Vec2FromRadians(rad float32) Vec2 {
	sin, cos := math.Sincos(float64(rad))
	return Vec2{float32(cos), float32(sin)}
}

func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v[0] * s, v[1] * s}
}

func (v Vec2) Add(o Vec2) Vec2 {
	return Vec2{v[0] + o[0], v[1] + o[1]}
}

func (v Vec2) Sub(o Vec2) Vec2 {
	return Vec2{v[0] - o[0], v[1] - o[1]}
}

func (v *Vec2) AddEqual(o Vec2) {
	(*v)[0] += o[0]
	(*v)[1] += o[1]
}

func (v *Vec2) Length() float32 {
	x := (*v)[0]
	y := (*v)[1]
	return float32(math.Sqrt(float64(x*x + y*y)))
}

func (v Vec2) Normalize() Vec2 {
	return v.Scale(1 / v.Length())
}

func (v Vec2) Dot(o Vec2) float32 {
	return v[0]*o[0] + v[1]*o[1]
}

type Lookup [2]int

func (l *Lookup) Alive() bool {
	return l != nil && (*l)[0] >= 0
}

type ShipControl struct {
	Up           bool
	Down         bool
	Left         bool
	Right        bool
	Fire         bool
	FireCoolDown float32
}

// TODO: Use?
type PlayerConnectedEvent struct {
}

type ExplosionDetails struct {
	Initialized    bool
	MoreExplosions bool
}

type MissileDetails struct {
	Owner *Lookup
}
