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
	"log"
	"math"
	"math/rand"

	"github.com/golang/protobuf/proto"
	"github.com/mbychkowski/space-agon/space-agon/game/pb"
)

type Game struct {
	E             *Entities
	initialized   bool
	NextNetworkId uint64

	ControlledShip *Lookup
	timeDead       float32
	NetworkIds     map[uint64]*Lookup
}

func NewGame() *Game {
	g := &Game{
		E: newEntities(),
		// Oh man, this is such a bad hack.
		NextNetworkId: uint64(rand.Int63()),
		// NewClientUpdate: NewNetworkUpdate(),

		timeDead:   100,
		NetworkIds: make(map[uint64]*Lookup),
	}

	return g
}

func (g *Game) NextNid() uint64 {
	// TODO: ensure only being called by host somehow?
	nid := g.NextNetworkId
	g.NextNetworkId++
	return nid
}

type Keystate struct {
	Press   bool
	Hold    bool
	Release bool
}

func (k *Keystate) FrameEndReset() {
	k.Press = false
	k.Release = false
}

func (k *Keystate) Down() {
	if !k.Hold {
		k.Press = true
		k.Hold = true
	}
}

func (k *Keystate) Up() {
	if k.Hold {
		k.Release = true
		k.Hold = false
	}
}

func getNid(g *Game, i *Iter, nid uint64) bool {
	lookup, ok := g.NetworkIds[nid]
	if !ok || !lookup.Alive() {
		return false
	}
	i.Get(lookup)
	return true
}

type Input struct {
	Up    Keystate
	Down  Keystate
	Left  Keystate
	Right Keystate
	Fire  Keystate
	Dt    float32
	// Whether entities which only exist for render should be created.
	IsRendered bool
	// Whether code which runs only if the input is going to control
	// a player ship should run.
	IsPlayer    bool
	IsConnected bool
	// Whether the this instance is the host, if not it is a client.
	IsHost   bool
	Cid      int64
	Memos    []*pb.Memo
	MemosOut []*pb.Memo
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) SendTo(to int64, actual proto.Message) {
	i.SendMemo(&pb.Memo{
		Recipient: &pb.Memo_To{
			To: to,
		},
	}, actual)
}

func (i *Input) BroadcastOthers(actual proto.Message) {
	i.SendMemo(&pb.Memo{
		Recipient: &pb.Memo_EveryoneBut{
			EveryoneBut: i.Cid,
		},
	}, actual)
}

func (i *Input) BroadcastAll(actual proto.Message) {
	i.SendMemo(&pb.Memo{
		Recipient: &pb.Memo_Everyone{},
	}, actual)
}

func (i *Input) SendMemo(partial *pb.Memo, actual proto.Message) {
	switch a := actual.(type) {
	case *pb.PosTracks:
		partial.Actual = &pb.Memo_PosTracks{PosTracks: a}
	case *pb.MomentumTracks:
		partial.Actual = &pb.Memo_MomentumTracks{MomentumTracks: a}
	case *pb.RotTracks:
		partial.Actual = &pb.Memo_RotTracks{RotTracks: a}
	case *pb.SpinTracks:
		partial.Actual = &pb.Memo_SpinTracks{SpinTracks: a}
	case *pb.ShipControlTrack:
		partial.Actual = &pb.Memo_ShipControlTrack{ShipControlTrack: a}
	// case *pb.SpawnEvent:
	// 	partial.Actual = &pb.Memo_SpawnEvent{SpawnEvent: a}
	case *pb.DestroyEvent:
		partial.Actual = &pb.Memo_DestroyEvent{DestroyEvent: a}
	case *pb.ShootMissile:
		partial.Actual = &pb.Memo_ShootMissile{ShootMissile: a}
	case *pb.SpawnMissile:
		partial.Actual = &pb.Memo_SpawnMissile{SpawnMissile: a}
	case *pb.SpawnExplosion:
		partial.Actual = &pb.Memo_SpawnExplosion{SpawnExplosion: a}
	case *pb.SpawnShip:
		partial.Actual = &pb.Memo_SpawnShip{SpawnShip: a}
	case *pb.RegisterPlayer:
		partial.Actual = &pb.Memo_RegisterPlayer{RegisterPlayer: a}
	default:
		panic("Unknown memo actual type")
	}

	i.MemosOut = append(i.MemosOut, partial)
}

func (inp *Input) FrameEndReset() {
	inp.Up.FrameEndReset()
	inp.Down.FrameEndReset()
	inp.Left.FrameEndReset()
	inp.Right.FrameEndReset()
	inp.Fire.FrameEndReset()
}

const ExplosionRadius = 2

func (g *Game) Step(input *Input) {
	if g.ControlledShip.Alive() {
		i := g.E.NewIter()
		i.Get(g.ControlledShip)

		shipControl := i.ShipControl()
		shipControl.Up = input.Up.Hold
		shipControl.Down = input.Down.Hold
		shipControl.Left = input.Left.Hold
		shipControl.Right = input.Right.Hold
		shipControl.Fire = input.Fire.Hold
	}

	for _, memo := range input.Memos {
		switch actual := memo.Actual.(type) {
		case *pb.Memo_DestroyEvent:
			destroyEvent := actual.DestroyEvent

			i := g.E.NewIter()
			if getNid(g, i, destroyEvent.Nid) {
				i.Remove()
			}

		// case *pb.Memo_SpawnEvent:
		// 	spawnEvent := actual.SpawnEvent

		// 	i := g.E.NewIter()
		// 	i.Require(LookupKey)
		// 	i.Require(NetworkReceiveKey)

		// 	switch spawnEvent.SpawnType {
		// 	case pb.SpawnEvent_SHIP:
		// 		spawnSpaceship(i)
		// 	// case pb.SpawnEvent_MISSILE:
		// 	// 	spawnMissile(i)
		// 	// case pb.SpawnEvent_EXPLOSION:
		// 	// 	spawnExplosion(i)
		// 	default:
		// 		log.Println("Asked to spawn", spawnEvent.SpawnType)
		// 		panic("Spawn what now?")
		// 	}

		// 	*i.NetworkId() = spawnEvent.Nid
		// 	g.NetworkIds[spawnEvent.Nid] = i.Lookup()

		case *pb.Memo_PosTracks:
			posTracks := actual.PosTracks
			i := g.E.NewIter()

			for index, nid := range posTracks.Nid {
				if getNid(g, i, nid) {
					*i.Pos() = Vec2{posTracks.X[index], posTracks.Y[index]}
				}
			}

		case *pb.Memo_RotTracks:
			rotTracks := actual.RotTracks
			i := g.E.NewIter()

			for index, nid := range rotTracks.Nid {
				if getNid(g, i, nid) {
					*i.Rot() = rotTracks.R[index]
				}
			}

		case *pb.Memo_MomentumTracks:
			momentumTracks := actual.MomentumTracks
			i := g.E.NewIter()

			for index, nid := range momentumTracks.Nid {
				if getNid(g, i, nid) {
					*i.Momentum() = Vec2{momentumTracks.X[index], momentumTracks.Y[index]}
				}
			}

		case *pb.Memo_SpinTracks:
			spinTracks := actual.SpinTracks
			i := g.E.NewIter()

			for index, nid := range spinTracks.Nid {
				if getNid(g, i, nid) {
					*i.Spin() = spinTracks.S[index]
				}
			}

		case *pb.Memo_ShipControlTrack:
			shipControlTrack := actual.ShipControlTrack
			i := g.E.NewIter()

			if getNid(g, i, shipControlTrack.Nid) {
				sc := i.ShipControl()
				sc.Up = shipControlTrack.Up
				sc.Left = shipControlTrack.Left
				sc.Right = shipControlTrack.Right
			}

		case *pb.Memo_ShootMissile:
			shootMissile := actual.ShootMissile

			i := g.E.NewIter()
			if getNid(g, i, shootMissile.Owner) {

				const MissileSpeed = 13
				momentum := *i.Momentum()
				momentum.AddEqual(Vec2FromRadians(*i.Rot()).Scale(MissileSpeed))

				input.BroadcastAll(&pb.SpawnMissile{
					Nid:      g.NextNid(),
					Owner:    shootMissile.Owner,
					Pos:      i.Pos().ToProto(),
					Momentum: momentum.ToProto(),
					Rot:      *i.Rot(),
					Spin:     *i.Spin(),
				})
			}

		case *pb.Memo_SpawnMissile:
			spawnMissile := actual.SpawnMissile

			i := g.E.NewIter()
			if input.IsHost {
				i.Require(NetworkTransmitKey)
				i.Require(CanExplodeKey)
				i.Require(TimedExplodeKey)
			} else {
				i.Require(NetworkReceiveKey)
			}

			i.Require(NetworkIdKey)
			i.Require(PosKey)
			i.Require(MomentumKey)
			i.Require(RotKey)
			i.Require(SpinKey)
			i.Require(SpriteKey)
			i.Require(AffectedByGravityKey)
			i.Require(LookupKey)
			i.Require(CanExplodeKey)
			i.Require(MissileDetailsKey)

			i.New()

			if input.IsHost {
				*i.TimedExplode() = 2
			}

			*i.NetworkId() = spawnMissile.Nid
			g.NetworkIds[spawnMissile.Nid] = i.Lookup()
			*i.Pos() = Vec2FromProto(spawnMissile.Pos)
			*i.Momentum() = Vec2FromProto(spawnMissile.Momentum)
			*i.Rot() = spawnMissile.Rot
			*i.Spin() = spawnMissile.Spin
			*i.Sprite() = SpriteMissile
			i.MissileDetails().Owner = g.NetworkIds[spawnMissile.Owner]

		case *pb.Memo_SpawnExplosion:
			spawnExplosion := actual.SpawnExplosion

			pos := Vec2FromProto(spawnExplosion.Pos)
			momentum := Vec2FromProto(spawnExplosion.Momentum)

			if input.IsHost {
				i := g.E.NewIter()
				i.Require(PosKey)
				i.Require(CanExplodeKey)
				i.Require(NetworkIdKey)

				for i.Next() {
					diff := pos.Sub(*i.Pos())
					if diff.Length() < ExplosionRadius {
						iMomentum := Vec2{}
						if i.Momentum() != nil {
							iMomentum = *i.Momentum()
						}
						input.BroadcastOthers(&pb.DestroyEvent{
							Nid: *i.NetworkId(),
						})
						input.BroadcastAll(&pb.SpawnExplosion{
							Pos:      i.Pos().ToProto(),
							Momentum: iMomentum.ToProto(),
						})
						i.Remove()
					}
				}
			}

			if input.IsRendered { // explosion fun :D
				i := g.E.NewIter() // particle iter
				i.Require(PosKey)
				i.Require(MomentumKey)
				i.Require(TimedDestroyKey)
				i.Require(PointRenderKey)
				i.Require(ParticleSunDeleteKey)

				for j := 0; j < 1000; j++ {
					speed := rand.Float32()*10 + 0.01
					dir := rand.Float32() * math.Pi * 2
					ttl := 1.0/speed + rand.Float32()
					if ttl > 3 {
						ttl = 3
					}

					i.New()
					*i.Pos() = pos.Add(Vec2FromRadians(rand.Float32() * math.Pi * 2).Scale(rand.Float32() * ExplosionRadius))
					*i.Momentum() = momentum.Add(Vec2FromRadians(dir).Scale(speed))
					*i.TimedDestroy() = ttl
				}
			}

			if input.IsRendered {
				i := g.E.NewIter()
				i.Require(PosKey)
				i.Require(MomentumKey)
				i.Require(TimedDestroyKey)
				i.Require(SpriteKey)

				i.New()
				*i.Pos() = pos
				*i.Momentum() = momentum
				*i.TimedDestroy() = 0.07
				*i.Sprite() = SpriteExplosionFlash
			}

		case *pb.Memo_SpawnShip:
			spawnShip := actual.SpawnShip

			var pos Vec2
			var r float32
			{
				type possibility struct {
					pos     Vec2
					r       float32
					closest float32
				}
				possibilities := []possibility{}
				for r := float32(0); r < math.Pi*2; r += math.Pi / 6 {
					possibilities = append(possibilities, possibility{
						pos:     Vec2FromRadians(r).Scale(14),
						r:       r,
						closest: float32(math.Inf(1)),
					})
				}
				i := g.E.NewIter()
				i.Require(PosKey)
				i.Require(CanExplodeKey)

				for i.Next() {
					for j := range possibilities {
						diff := i.Pos().Sub(possibilities[j].pos)
						dist := diff.Length()
						if dist < possibilities[j].closest {
							possibilities[j].closest = dist
						}
					}
				}

				best := possibilities[0]
				for _, j := range possibilities[1:] {
					if j.closest > best.closest {
						best = j
					}
				}

				pos = best.pos
				r = best.r
			}

			i := g.E.NewIter()
			if spawnShip.Authority == input.Cid {
				i.Require(NetworkTransmitKey)
			} else {
				i.Require(NetworkReceiveKey)
			}

			i.Require(PosKey)
			i.Require(RotKey)
			i.Require(SpriteKey)
			i.Require(KeepInCameraKey)
			i.Require(SpinKey)
			i.Require(MomentumKey)
			i.Require(ShipControlKey)
			i.Require(LookupKey)
			i.Require(AffectedByGravityKey)
			i.Require(NetworkIdKey)
			i.Require(BoundLocationKey)
			i.Require(CanExplodeKey)
			i.New()

			// *i.Pos() = Vec2FromProto(spawnShip.Pos)
			// *i.Momentum() = Vec2FromProto(spawnShip.Momentum)
			*i.Pos() = pos
			*i.Momentum() = Vec2FromRadians(r + math.Pi/2).Scale(3.5)
			*i.Rot() = r + math.Pi/2
			*i.Spin() = spawnShip.Spin
			// pos := i.Pos()
			// (*pos)[0] = 7
			// (*pos)[1] = 0
			// (*i.Momentum())[1] = 5
			// *i.Rot() = 0
			// *i.NetworkId() = g.NextNetworkId
			// g.NextNetworkId++

			*i.NetworkId() = spawnShip.Nid
			g.NetworkIds[spawnShip.Nid] = i.Lookup()

			if spawnShip.Authority == input.Cid {
				g.ControlledShip = i.Lookup()
				*i.Sprite() = SpriteShip
			} else {
				*i.Sprite() = SpriteEnemyShip
			}

		case *pb.Memo_RegisterPlayer:
			registerPlayer := actual.RegisterPlayer

			input.BroadcastAll(&pb.SpawnShip{
				Nid:       g.NextNid(),
				Authority: registerPlayer.Cid,
				// Pos:       (&Vec2{14, 0}).ToProto(),
				// Momentum:  (&Vec2{0, 3.5}).ToProto(),
				// Rot:  0,
				// Spin: 0,
			})

		default:
			log.Fatal("Unknown message type:", actual)
		}
	}

	if !g.initialized {
		if input.IsRendered { // spawn stars
			{ // Big star
				i := g.E.NewIter()
				i.Require(PosKey)
				i.Require(SpriteKey)
				i.New()

				*i.Sprite() = SpriteStar
			}

			{
				i := g.E.NewIter()
				i.Require(PosKey)
				i.Require(PointRenderKey)

				// spawn small stars
				const density = 0.05
				const starBoxRadius = 200
				for j := 0; j < int(density*starBoxRadius*starBoxRadius); j++ {
					i.New()
					// *i.Sprite() = SpriteStarBit
					*i.Pos() = Vec2{
						rand.Float32()*starBoxRadius*2 - starBoxRadius,
						rand.Float32()*starBoxRadius*2 - starBoxRadius,
					}
				}
			}
		}

		g.initialized = true
	}

	if input.IsPlayer && input.IsConnected { // spawn/respawn
		if !g.ControlledShip.Alive() {
			g.timeDead += input.Dt

			if g.timeDead > 4 {
				g.timeDead = 0

				input.SendTo(0, &pb.RegisterPlayer{
					Cid: input.Cid,
				})
			}
		}
	}

	// Explosions cause more explosions
	// for {
	// 	newExplosions := [][2]Vec2(nil)
	// 	{
	// 		i := g.E.NewIter()
	// 		i.Require(PosKey)
	// 		i.Require(ExplosionDetailsKey)

	// 		for i.Next() {
	// 			if !i.ExplosionDetails().MoreExplosions {
	// 				i.ExplosionDetails().MoreExplosions = true
	// 				other := g.E.NewIter()
	// 				other.Require(PosKey)
	// 				other.Require(LookupKey)
	// 				other.Require(CanExplodeKey)
	// 				other.Require(NetworkTransmitKey)
	// 				for other.Next() {
	// 					if i.Lookup() == other.Lookup() {
	// 						continue
	// 					}
	// 					diff := i.Pos().Sub(*other.Pos())
	// 					if diff.Length() < 1 {
	// 						newExplosions = append(newExplosions, [2]Vec2{*other.Pos(), *other.Momentum()})
	// 						if other.NetworkId() != nil {

	// 							input.BroadcastOthers(&pb.DestroyEvent{
	// 								Nid: uint64(*other.NetworkId()),
	// 							})

	// 							other.Remove()
	// 							break
	// 						}
	// 					}

	// 				}
	// 			}
	// 		}

	// 		if len(newExplosions) == 0 {
	// 			break
	// 		}

	// 		for _, posMomentum := range newExplosions {
	// 			ie := g.E.NewIter()
	// 			ie.Require(NetworkTransmitKey)
	// 			ie.Require(TimedDestroyKey)
	// 			spawnExplosion(ie)
	// 			*ie.Pos() = posMomentum[0]
	// 			*ie.Momentum() = posMomentum[1]
	// 			*ie.NetworkId() = g.NextNetworkId
	// 			*ie.TimedDestroy() = 0.5
	// 			g.NextNetworkId++

	// 			input.BroadcastOthers(&pb.SpawnEvent{
	// 				Nid:       uint64(*ie.NetworkId()),
	// 				SpawnType: pb.SpawnEvent_EXPLOSION,
	// 			})
	// 		}
	// 	}
	// }

	// if input.IsRendered { // explosion fun :D
	// 	i := g.E.NewIter()
	// 	i.Require(PosKey)
	// 	i.Require(MomentumKey)
	// 	i.Require(ExplosionDetailsKey)

	// 	pi := g.E.NewIter() // particle iter
	// 	pi.Require(PosKey)
	// 	pi.Require(MomentumKey)
	// 	pi.Require(TimedDestroyKey)
	// 	pi.Require(PointRenderKey)
	// 	pi.Require(ParticleSunDeleteKey)

	// 	for i.Next() {
	// 		if !i.ExplosionDetails().Initialized {
	// 			i.ExplosionDetails().Initialized = true
	// 			for j := 0; j < 1000; j++ {
	// 				speed := rand.Float32()*10 + 0.01
	// 				dir := rand.Float32() * math.Pi * 2
	// 				ttl := 1.0/speed + rand.Float32()
	// 				if ttl > 3 {
	// 					ttl = 3
	// 				}

	// 				pi.New()
	// 				*pi.Pos() = *i.Pos()
	// 				*pi.Momentum() = i.Momentum().Add(Vec2FromRadians(dir).Scale(speed))
	// 				*pi.TimedDestroy() = ttl
	// 			}
	// 		}
	// 	}
	// }

	if input.IsRendered { // Spawn sun particles
		i := g.E.NewIter()
		i.Require(PosKey)
		i.Require(PointRenderKey)
		i.Require(MomentumKey)
		i.Require(TimedDestroyKey)

		for j := 0; j < 10; j++ {
			i.New()
			rad := rand.Float32() * 2 * math.Pi
			*i.Pos() = Vec2FromRadians(rad)
			rad += rand.Float32()*2 - 1
			*i.Momentum() = Vec2FromRadians(rad).Scale(rand.Float32()*5 + 1)
			*i.TimedDestroy() = rand.Float32()*2 + 1
		}
	}

	{
		i := g.E.NewIter()
		i.Require(TimedDestroyKey)
		for i.Next() {
			*i.TimedDestroy() -= input.Dt
			if *i.TimedDestroy() <= 0 {
				if i.NetworkId() != nil {
					input.BroadcastOthers(&pb.DestroyEvent{
						Nid: uint64(*i.NetworkId()),
					})
				}
				i.Remove()
			}
		}
	}

	{
		i := g.E.NewIter()
		i.Require(TimedExplodeKey)
		i.Require(PosKey)
		i.Require(MomentumKey)
		i.Require(NetworkIdKey)
		for i.Next() {
			*i.TimedExplode() -= input.Dt
			if *i.TimedExplode() <= 0 {
				input.BroadcastOthers(&pb.DestroyEvent{
					Nid: *i.NetworkId(),
				})
				input.BroadcastAll(&pb.SpawnExplosion{
					Pos:      i.Pos().ToProto(),
					Momentum: i.Momentum().ToProto(),
				})
				i.Remove()
			}
		}
	}

	{ // Explode in the sun
		i := g.E.NewIter()
		i.Require(CanExplodeKey)
		i.Require(PosKey)
		i.Require(NetworkTransmitKey)
		for i.Next() {
			if i.Pos().Length() < 3 {
				input.BroadcastOthers(&pb.DestroyEvent{
					Nid: *i.NetworkId(),
				})
				input.BroadcastAll(&pb.SpawnExplosion{
					Pos:      i.Pos().ToProto(),
					Momentum: i.Momentum().ToProto(),
				})
				i.Remove()
			}
		}
	}

	{ // Explode When colliding
		i := g.E.NewIter()
		i.Require(MissileDetailsKey)
		i.Require(PosKey)
		i.Require(LookupKey)
		i.Require(NetworkTransmitKey)
		for i.Next() {
			other := g.E.NewIter()
			other.Require(PosKey)
			other.Require(LookupKey)
			other.Require(CanExplodeKey)
			for other.Next() {
				if i.Lookup() == other.Lookup() || i.MissileDetails().Owner == other.Lookup() {
					continue
				}
				diff := i.Pos().Sub(*other.Pos())
				if diff.Length() < ExplosionRadius*0.8 {
					input.BroadcastOthers(&pb.DestroyEvent{
						Nid: *i.NetworkId(),
					})
					input.BroadcastAll(&pb.SpawnExplosion{
						Pos:      i.Pos().ToProto(),
						Momentum: i.Momentum().ToProto(),
					})
					i.Remove()
					break
				}
			}
		}
	}

	{
		i := g.E.NewIter()
		i.Require(PosKey)
		i.Require(ParticleSunDeleteKey)
		for i.Next() {
			if i.Pos().Length() < 2.3 {
				i.Remove()
			}
		}
	}

	{
		i := g.E.NewIter()
		i.Require(PosKey)
		i.Require(RotKey)
		i.Require(ShipControlKey)
		i.Require(SpinKey)
		i.Require(MomentumKey)
		i.Require(LookupKey)
		i.Require(NetworkIdKey)
		for i.Next() {
			///////////////////////////
			// Ship Movement Controls
			///////////////////////////
			const rotationForSpeed = 5
			const rotationAgainstSpeed = 10
			const forwardSpeed = 4

			spinDesire := float32(0)
			if i.ShipControl().Left {
				spinDesire++
			}
			if i.ShipControl().Right {
				spinDesire--
			}
			if !i.ShipControl().Left && !i.ShipControl().Right {
				s := *i.Spin()
				if s < -0.5 {
					spinDesire += 0.1
				} else if s > 0.5 {
					spinDesire -= 0.1
				}
			}

			// Game feel: Stopping spin is easier than starting it.
			if (spinDesire < 0) == (*i.Spin() < 0) {
				spinDesire *= rotationForSpeed
			} else {
				spinDesire *= rotationAgainstSpeed
			}

			*i.Spin() += spinDesire * input.Dt

			if i.ShipControl().Up {
				dx := float32(math.Cos(float64(*i.Rot()))) * forwardSpeed * input.Dt
				dy := float32(math.Sin(float64(*i.Rot()))) * forwardSpeed * input.Dt

				(*i.Momentum())[0] += dx
				(*i.Momentum())[1] += dy
			}

			///////////////////////////
			// Ship Weapons
			///////////////////////////
			i.ShipControl().FireCoolDown -= input.Dt

			if i.ShipControl().FireCoolDown <= 0 && i.ShipControl().Fire {
				input.SendTo(0, &pb.ShootMissile{
					Owner: *i.NetworkId(),
				})

				i.ShipControl().FireCoolDown = 0.5
				// i.ShipControl().FireCoolDown = 5
			}
		}
	}

	if input.IsRendered { // Spawn ship movement particles
		i := g.E.NewIter()
		i.Require(PosKey)
		i.Require(RotKey)
		i.Require(ShipControlKey)
		i.Require(MomentumKey)

		ip := g.E.NewIter()
		ip.Require(PosKey)
		ip.Require(PointRenderKey)
		ip.Require(MomentumKey)
		ip.Require(TimedDestroyKey)
		ip.Require(ParticleSunDeleteKey)

		for i.Next() {
			const pushFactor = 5
			emitPoint := i.Pos().Sub(Vec2FromRadians(*i.Rot()).Scale(0.4))

			if i.ShipControl().Up {
				ip.New()
				*ip.Pos() = emitPoint

				angleOut := *i.Rot() + math.Pi + (rand.Float32()-0.5)/3
				*ip.Momentum() = i.Momentum().Add(Vec2FromRadians(angleOut).Scale(pushFactor))
				*ip.TimedDestroy() = rand.Float32()*2 + 1
			}
			if i.ShipControl().Left {
				ip.New()
				*ip.Pos() = emitPoint

				angleOut := *i.Rot() + math.Pi/2 + (rand.Float32()-0.5)/3
				*ip.Momentum() = i.Momentum().Add(Vec2FromRadians(angleOut).Scale(pushFactor))
				*ip.TimedDestroy() = rand.Float32()*2 + 1
			}
			if i.ShipControl().Right {
				ip.New()
				*ip.Pos() = emitPoint

				angleOut := *i.Rot() + math.Pi*3/2 + (rand.Float32()-0.5)/3
				*ip.Momentum() = i.Momentum().Add(Vec2FromRadians(angleOut).Scale(pushFactor))
				*ip.TimedDestroy() = rand.Float32()*2 + 1
			}
		}
	}

	{
		i := g.E.NewIter()
		i.Require(RotKey)
		i.Require(MomentumKey)
		i.Require(MissileDetailsKey)

		for i.Next() {
			const pushFactor = 10
			i.Momentum().AddEqual(Vec2FromRadians(*i.Rot()).Scale(pushFactor * input.Dt))
		}
	}

	if input.IsRendered { // Spawn missile trail particles
		i := g.E.NewIter()
		i.Require(RotKey)
		i.Require(MissileDetailsKey)
		i.Require(PosKey)

		ip := g.E.NewIter()
		ip.Require(PosKey)
		ip.Require(PointRenderKey)
		ip.Require(MomentumKey)
		ip.Require(TimedDestroyKey)
		ip.Require(ParticleSunDeleteKey)

		for i.Next() {
			const pushFactor = 5

			for j := 0; j < 4; j++ {
				ip.New()
				*ip.Pos() = *i.Pos()

				angleOut := *i.Rot() + math.Pi + (rand.Float32()-0.5)/2
				*ip.Momentum() = i.Momentum().Add(Vec2FromRadians(angleOut).Scale(pushFactor))
				ip.Pos().AddEqual(ip.Momentum().Scale(float32(j) * input.Dt / 4))
				*ip.TimedDestroy() = rand.Float32()*2 + 1
			}
		}
	}

	{
		i := g.E.NewIter()
		i.Require(RotKey)
		i.Require(SpinKey)

		for i.Next() {
			*i.Rot() += *i.Spin() * input.Dt
		}
	}

	{
		i := g.E.NewIter()
		i.Require(BoundLocationKey)
		i.Require(PosKey)
		i.Require(MomentumKey)

		for i.Next() {
			l := i.Pos().Length()
			if l > 50 {
				scale := 50 / l
				*i.Pos() = i.Pos().Scale(scale)
				// Calculate the momentum in the direction of the invisible wall, and
				// cancel it out.
				i.Momentum().AddEqual(i.Pos().Normalize().Scale((*i.Pos()).Normalize().Scale(-1).Dot(*i.Momentum())))
			}
		}
	}

	{
		// Force of gravity = gravconst * mass1 * mass2 / (distance)^2

		// Update value = Dt * const * normalized direction vector / (distance)^2

		// = Dt * const * (-1 * Pos / Pos.Lenght) / (Pos.Length) ^ 2
		// = Dt * const * -1 * Pos / Pos.Length ^ 3
		// = Pos.Scale(Dt * const * -1 / Pos.Length ^ 3)

		// Pos.Length = (x*x + y*y) ^ 1/2
		// sqrt then cube will probably be faster than taking to the power of 1.5?

		i := g.E.NewIter()
		i.Require(PosKey)
		i.Require(AffectedByGravityKey)
		i.Require(MomentumKey)

		const gravityStrength = 200

		for i.Next() {
			length := i.Pos().Length()
			lengthCubed := length * length * length
			i.Momentum().AddEqual(i.Pos().Scale(-1 * gravityStrength * input.Dt / lengthCubed))
		}
	}

	{
		i := g.E.NewIter()
		i.Require(PosKey)
		i.Require(MomentumKey)

		for i.Next() {
			i.Pos().AddEqual(i.Momentum().Scale(input.Dt))
		}
	}

	{
		i := g.E.NewIter()
		i.Require(FrameEndDeleteKey)
		for i.Next() {
			i.Remove()
		}
	}

	{
		posTracks := &pb.PosTracks{}

		i := g.E.NewIter()
		i.Require(PosKey)
		i.Require(NetworkTransmitKey)
		i.Require(NetworkIdKey)

		for i.Next() {
			posTracks.Nid = append(posTracks.Nid, *i.NetworkId())
			pos := *i.Pos()
			posTracks.X = append(posTracks.X, pos[0])
			posTracks.Y = append(posTracks.Y, pos[1])
		}

		input.BroadcastOthers(posTracks)
	}

	{
		rotTracks := &pb.RotTracks{}

		i := g.E.NewIter()
		i.Require(RotKey)
		i.Require(NetworkTransmitKey)
		i.Require(NetworkIdKey)

		for i.Next() {
			rotTracks.Nid = append(rotTracks.Nid, *i.NetworkId())
			rotTracks.R = append(rotTracks.R, *i.Rot())
		}

		input.BroadcastOthers(rotTracks)
	}

	{
		momentumTracks := &pb.MomentumTracks{}

		i := g.E.NewIter()
		i.Require(MomentumKey)
		i.Require(NetworkTransmitKey)
		i.Require(NetworkIdKey)

		for i.Next() {
			momentumTracks.Nid = append(momentumTracks.Nid, *i.NetworkId())
			momentum := *i.Momentum()
			momentumTracks.X = append(momentumTracks.X, momentum[0])
			momentumTracks.Y = append(momentumTracks.Y, momentum[1])
		}

		input.BroadcastOthers(momentumTracks)
	}

	{
		spinTracks := &pb.SpinTracks{}

		i := g.E.NewIter()
		i.Require(SpinKey)
		i.Require(NetworkTransmitKey)
		i.Require(NetworkIdKey)

		for i.Next() {
			spinTracks.Nid = append(spinTracks.Nid, *i.NetworkId())
			spinTracks.S = append(spinTracks.S, *i.Spin())
		}

		input.BroadcastOthers(spinTracks)
	}

	{
		i := g.E.NewIter()
		i.Require(ShipControlKey)
		i.Require(NetworkTransmitKey)
		i.Require(NetworkIdKey)

		for i.Next() {
			shipControlTrack := &pb.ShipControlTrack{}

			shipControlTrack.Nid = *i.NetworkId()
			sc := i.ShipControl()
			shipControlTrack.Up = sc.Up
			shipControlTrack.Left = sc.Left
			shipControlTrack.Right = sc.Right

			input.BroadcastOthers(shipControlTrack)
		}
	}
}

// func spawnSpaceship(i *Iter) {
// 	i.Require(PosKey)
// 	i.Require(RotKey)
// 	i.Require(SpriteKey)
// 	i.Require(KeepInCameraKey)
// 	i.Require(SpinKey)
// 	i.Require(MomentumKey)
// 	i.Require(ShipControlKey)
// 	i.Require(LookupKey)
// 	i.Require(AffectedByGravityKey)
// 	i.Require(NetworkIdKey)
// 	i.Require(BoundLocationKey)
// 	i.Require(CanExplodeKey)
// 	i.New()

// 	*i.Sprite() = SpriteShip
// }

// func spawnExplosion(i *Iter) {
// 	i.Require(PosKey)
// 	i.Require(MomentumKey)
// 	i.Require(ExplosionDetailsKey)
// 	i.Require(NetworkIdKey)
// 	i.New()
// 	// REMOVE TO TRIGGER WASM BUG, IS ERRONOUSLY TRUE
// 	i.ExplosionDetails().Initialized = false
// 	// log.Println("Is initialized", i.ExplosionDetails().Initialized)
// }
