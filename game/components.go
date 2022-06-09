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

// THIS FILE IS GENERATED, DO NOT EDIT
// THIS FILE IS GENERATED, DO NOT EDIT
// THIS FILE IS GENERATED, DO NOT EDIT
// THIS FILE IS GENERATED, DO NOT EDIT
// THIS FILE IS GENERATED, DO NOT EDIT
// THIS FILE IS GENERATED, DO NOT EDIT
// THIS FILE IS GENERATED, DO NOT EDIT
// THIS FILE IS GENERATED, DO NOT EDIT
// THIS FILE IS GENERATED, DO NOT EDIT
// THIS FILE IS GENERATED, DO NOT EDIT
// THIS FILE IS GENERATED, DO NOT EDIT
// If you want to define new components, edit game/generation/generation.go
// Then run: go generate github.com/googleforgames/space-agon/game/generation

package game

// Lookup has a special implementation which keeps track of how it is moved,
// just include it manually here.
type comp_Lookup []*Lookup

func (c *comp_Lookup) Swap(j1, j2 int) {
	(*c)[j1], (*c)[j2] = (*c)[j2], (*c)[j1]
	(*c)[j1][1] = j1
	(*c)[j2][1] = j2
}

func (c *comp_Lookup) Extend(i int) {
	j := len(*c)
	*c = append(*c, &Lookup{i, j})
}

func (c *comp_Lookup) RemoveLast() {
	j := len(*c) - 1
	(*c)[j][0] = -2
	(*c)[j][1] = -3
	*c = (*c)[:j]
}

type comp_MissileDetails []MissileDetails

func (c *comp_MissileDetails) Swap(j1, j2 int) {
	(*c)[j1], (*c)[j2] = (*c)[j2], (*c)[j1]
}

func (c *comp_MissileDetails) Extend(i int) {
	*c = append(*c, MissileDetails{})
}

func (c *comp_MissileDetails) RemoveLast() {
	*c = (*c)[:len(*c)-1]
}

type comp_ShipControl []ShipControl

func (c *comp_ShipControl) Swap(j1, j2 int) {
	(*c)[j1], (*c)[j2] = (*c)[j2], (*c)[j1]
}

func (c *comp_ShipControl) Extend(i int) {
	*c = append(*c, ShipControl{})
}

func (c *comp_ShipControl) RemoveLast() {
	*c = (*c)[:len(*c)-1]
}

type comp_Sprite []Sprite

func (c *comp_Sprite) Swap(j1, j2 int) {
	(*c)[j1], (*c)[j2] = (*c)[j2], (*c)[j1]
}

func (c *comp_Sprite) Extend(i int) {
	*c = append(*c, SpriteUnset)
}

func (c *comp_Sprite) RemoveLast() {
	*c = (*c)[:len(*c)-1]
}

type comp_Vec2 []Vec2

func (c *comp_Vec2) Swap(j1, j2 int) {
	(*c)[j1], (*c)[j2] = (*c)[j2], (*c)[j1]
}

func (c *comp_Vec2) Extend(i int) {
	*c = append(*c, Vec2{})
}

func (c *comp_Vec2) RemoveLast() {
	*c = (*c)[:len(*c)-1]
}

type comp_float32 []float32

func (c *comp_float32) Swap(j1, j2 int) {
	(*c)[j1], (*c)[j2] = (*c)[j2], (*c)[j1]
}

func (c *comp_float32) Extend(i int) {
	*c = append(*c, 0)
}

func (c *comp_float32) RemoveLast() {
	*c = (*c)[:len(*c)-1]
}

type comp_uint64 []uint64

func (c *comp_uint64) Swap(j1, j2 int) {
	(*c)[j1], (*c)[j2] = (*c)[j2], (*c)[j1]
}

func (c *comp_uint64) Extend(i int) {
	*c = append(*c, 0)
}

func (c *comp_uint64) RemoveLast() {
	*c = (*c)[:len(*c)-1]
}

const (
	AffectedByGravityKey = CompKey(iota)
	BoundLocationKey     = CompKey(iota)
	CanExplodeKey        = CompKey(iota)
	FrameEndDeleteKey    = CompKey(iota)
	KeepInCameraKey      = CompKey(iota)
	LookupKey            = CompKey(iota)
	MissileDetailsKey    = CompKey(iota)
	MomentumKey          = CompKey(iota)
	NetworkIdKey         = CompKey(iota)
	NetworkReceiveKey    = CompKey(iota)
	NetworkTransmitKey   = CompKey(iota)
	ParticleSunDeleteKey = CompKey(iota)
	PointRenderKey       = CompKey(iota)
	PosKey               = CompKey(iota)
	RotKey               = CompKey(iota)
	ShipControlKey       = CompKey(iota)
	SpinKey              = CompKey(iota)
	SpriteKey            = CompKey(iota)
	TimedDestroyKey      = CompKey(iota)
	TimedExplodeKey      = CompKey(iota)

	doNotMoveOrUseLastKeyForNumberOfKeys
)

type EntityBag struct {
	count    int
	comps    []Comp
	compsKey compsKey

	Lookup         *comp_Lookup
	MissileDetails *comp_MissileDetails
	Momentum       *comp_Vec2
	NetworkId      *comp_uint64
	Pos            *comp_Vec2
	Rot            *comp_float32
	ShipControl    *comp_ShipControl
	Spin           *comp_float32
	Sprite         *comp_Sprite
	TimedDestroy   *comp_float32
	TimedExplode   *comp_float32
}

func newEntityBag(compsKey *compsKey) *EntityBag {
	bag := &EntityBag{
		count:    0,
		comps:    nil,
		compsKey: *compsKey,
	}

	if inRequirement(compsKey, LookupKey) {
		bag.Lookup = &comp_Lookup{}
		bag.comps = append(bag.comps, bag.Lookup)
	}

	if inRequirement(compsKey, MissileDetailsKey) {
		bag.MissileDetails = &comp_MissileDetails{}
		bag.comps = append(bag.comps, bag.MissileDetails)
	}

	if inRequirement(compsKey, MomentumKey) {
		bag.Momentum = &comp_Vec2{}
		bag.comps = append(bag.comps, bag.Momentum)
	}

	if inRequirement(compsKey, NetworkIdKey) {
		bag.NetworkId = &comp_uint64{}
		bag.comps = append(bag.comps, bag.NetworkId)
	}

	if inRequirement(compsKey, PosKey) {
		bag.Pos = &comp_Vec2{}
		bag.comps = append(bag.comps, bag.Pos)
	}

	if inRequirement(compsKey, RotKey) {
		bag.Rot = &comp_float32{}
		bag.comps = append(bag.comps, bag.Rot)
	}

	if inRequirement(compsKey, ShipControlKey) {
		bag.ShipControl = &comp_ShipControl{}
		bag.comps = append(bag.comps, bag.ShipControl)
	}

	if inRequirement(compsKey, SpinKey) {
		bag.Spin = &comp_float32{}
		bag.comps = append(bag.comps, bag.Spin)
	}

	if inRequirement(compsKey, SpriteKey) {
		bag.Sprite = &comp_Sprite{}
		bag.comps = append(bag.comps, bag.Sprite)
	}

	if inRequirement(compsKey, TimedDestroyKey) {
		bag.TimedDestroy = &comp_float32{}
		bag.comps = append(bag.comps, bag.TimedDestroy)
	}

	if inRequirement(compsKey, TimedExplodeKey) {
		bag.TimedExplode = &comp_float32{}
		bag.comps = append(bag.comps, bag.TimedExplode)
	}

	return bag
}

func (iter *Iter) Lookup() *Lookup {
	comp := iter.e.bags[iter.i].Lookup
	if comp == nil {
		return nil
	}
	return (*comp)[iter.j]
}

func (iter *Iter) MissileDetails() *MissileDetails {
	comp := iter.e.bags[iter.i].MissileDetails
	if comp == nil {
		return nil
	}
	return &(*comp)[iter.j]
}

func (iter *Iter) Momentum() *Vec2 {
	comp := iter.e.bags[iter.i].Momentum
	if comp == nil {
		return nil
	}
	return &(*comp)[iter.j]
}

func (iter *Iter) NetworkId() *uint64 {
	comp := iter.e.bags[iter.i].NetworkId
	if comp == nil {
		return nil
	}
	return &(*comp)[iter.j]
}

func (iter *Iter) Pos() *Vec2 {
	comp := iter.e.bags[iter.i].Pos
	if comp == nil {
		return nil
	}
	return &(*comp)[iter.j]
}

func (iter *Iter) Rot() *float32 {
	comp := iter.e.bags[iter.i].Rot
	if comp == nil {
		return nil
	}
	return &(*comp)[iter.j]
}

func (iter *Iter) ShipControl() *ShipControl {
	comp := iter.e.bags[iter.i].ShipControl
	if comp == nil {
		return nil
	}
	return &(*comp)[iter.j]
}

func (iter *Iter) Spin() *float32 {
	comp := iter.e.bags[iter.i].Spin
	if comp == nil {
		return nil
	}
	return &(*comp)[iter.j]
}

func (iter *Iter) Sprite() *Sprite {
	comp := iter.e.bags[iter.i].Sprite
	if comp == nil {
		return nil
	}
	return &(*comp)[iter.j]
}

func (iter *Iter) TimedDestroy() *float32 {
	comp := iter.e.bags[iter.i].TimedDestroy
	if comp == nil {
		return nil
	}
	return &(*comp)[iter.j]
}

func (iter *Iter) TimedExplode() *float32 {
	comp := iter.e.bags[iter.i].TimedExplode
	if comp == nil {
		return nil
	}
	return &(*comp)[iter.j]
}

func inRequirement(compsKey *compsKey, compKey CompKey) bool {
	return 0 < (*compsKey)[compKey/compsKeyUnitSize]&(1<<(compKey%compsKeyUnitSize))
}

func (e *EntityBag) Add(i int) int {
	j := e.count
	e.count++
	for _, c := range e.comps {
		c.Extend(i)
	}
	return j
}

func (e *EntityBag) Remove(i int) {
	e.count--
	for _, c := range e.comps {
		c.Swap(e.count, i)
		c.RemoveLast()
	}
}

type Iter struct {
	e            *Entities
	i            int
	j            int
	requirements compsKey
}

func (iter *Iter) Require(k CompKey) {
	iter.requirements[k/compsKeyUnitSize] |= 1 << (k % compsKeyUnitSize)
}

func (iter *Iter) Next() bool {
	iter.j++
	for iter.i == -1 || iter.j >= iter.e.bags[iter.i].count {
		for {
			iter.i++
			if iter.i >= len(iter.e.bags) {
				return false
			}
			if iter.meetsRequirements(iter.e.bags[iter.i]) {
				break
			}
		}
		iter.j = 0
	}
	return true
}

func (iter *Iter) meetsRequirements(bag *EntityBag) bool {
	for i := 0; i < len(iter.requirements); i++ {
		if iter.requirements[i] != (iter.requirements[i] & bag.compsKey[i]) {
			return false
		}
	}
	return true
}

func (iter *Iter) New() {
	var ok bool
	iter.i, ok = iter.e.bagsByKey[iter.requirements]
	if !ok {
		iter.e.bagsByKey[iter.requirements] = len(iter.e.bags)
		iter.i = len(iter.e.bags)
		iter.e.bags = append(iter.e.bags, newEntityBag(&iter.requirements))
	}

	iter.j = iter.e.bags[iter.i].Add(iter.i)
}

func (iter *Iter) Get(indices *Lookup) {
	iter.i = (*indices)[0]
	iter.j = (*indices)[1]
}

func (iter *Iter) Remove() {
	iter.e.bags[iter.i].Remove(iter.j)
	// So that a call to next will arrive at this index, which now contains  a
	// different entity.
	iter.j--
}

type CompKey uint16
type compsKey [doNotMoveOrUseLastKeyForNumberOfKeys/compsKeyUnitSize + 1]uint8

const compsKeyUnitSize = 8

type Entities struct {
	bags      []*EntityBag
	bagsByKey map[compsKey]int
}

func newEntities() *Entities {
	return &Entities{
		bagsByKey: make(map[compsKey]int),
	}
}

func (e *Entities) NewIter() *Iter {
	return &Iter{
		e: e,
		i: -1,
		j: -1,
	}
}

type Comp interface {
	Swap(j1, j2 int)
	Extend(i int)
	RemoveLast()
}
