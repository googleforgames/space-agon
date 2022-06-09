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

//go:build js
// +build js

package main

import (
	"fmt"
	"log"
	"math"
	"syscall/js"

	"github.com/googleforgames/space-agon/client/webgl"
	"github.com/googleforgames/space-agon/game"
)

type graphics struct {
	w *webgl.WebGL

	width  int
	height int

	spritesheet  *webgl.Texture
	spriteShader *webgl.Program

	coords              []float32
	coordsBuffer        *webgl.Buffer
	textureCoords       []float32
	textureCoordsBuffer *webgl.Buffer

	uCenter [2]float32
	uScale  [2]float32

	written int

	pointShader   *webgl.Program
	pointCoords   []float32
	pointsWritten int
}

func NewGraphics() (*graphics, error) {
	g := &graphics{}

	// document.getElementById("game").parentElement.removeChild(document.getElementById("game"));

	canvas := js.Global().Get("document").Call("getElementById", "game")
	if canvas.Type() != js.TypeNull {
		canvas.Get("parentElement").Call("removeChild", canvas)
	}
	canvas = js.Global().Get("document").Call("createElement", "canvas")
	g.width = js.Global().Get("window").Get("innerWidth").Int()
	g.height = js.Global().Get("window").Get("innerHeight").Int()
	canvas.Set("width", g.width)
	canvas.Set("height", g.height)
	canvas.Set("id", "game")
	js.Global().Get("document").Call("getElementById", "container").Call("appendChild", canvas)

	var err error
	g.w, err = webgl.InitWebgl(canvas)
	if err != nil {
		return nil, err
	}

	g.w.Enable(g.w.BLEND)
	g.w.BlendFunc(g.w.SRC_ALPHA, g.w.ONE_MINUS_SRC_ALPHA)

	spritesheetElement := js.Global().Get("document").Call("getElementById", "spritesheet")
	// log.Println("WIDTH", spritesheetElement.Get("width"))
	// log.Println("HEIGHT", spritesheetElement.Get("height"))
	g.spritesheet = webgl.LoadTexture(g.w, spritesheetElement)

	g.spriteShader, err = webgl.CreateProgram(
		g.w,
		`
    uniform vec2 uCenter;
    uniform vec2 uScale; 

    attribute vec2 aVertexPosition;
    attribute vec2 aTextureCoord;

    varying highp vec2 vTextureCoord;

    void main() {
      gl_Position = vec4((aVertexPosition-uCenter)/uScale, 0.0, 1.0);
      vTextureCoord = aTextureCoord;
    }`,
		`
    precision highp float;

    varying highp vec2 vTextureCoord;

    uniform sampler2D uSampler;

    void main() {
      gl_FragColor = texture2D(uSampler, vTextureCoord);
    }
    `,
	)
	if err != nil {
		return nil, fmt.Errorf("Error building spriteShader: %w", err)
	}

	g.coords = make([]float32, webgl.MaxArrayLength/(3*2*2))
	g.textureCoords = make([]float32, len(g.coords))
	g.written = 0

	g.coordsBuffer = g.w.CreateBuffer()
	g.w.BindBuffer(g.w.ARRAY_BUFFER, g.coordsBuffer)

	// 4 comes from 4 bytes per float32
	g.w.BufferDataSize(g.w.ARRAY_BUFFER, 4*len(g.coords), g.w.DYNAMIC_DRAW)

	g.textureCoordsBuffer = g.w.CreateBuffer()
	g.w.BindBuffer(g.w.ARRAY_BUFFER, g.textureCoordsBuffer)
	g.w.BufferDataSize(g.w.ARRAY_BUFFER, 4*len(g.textureCoords), g.w.DYNAMIC_DRAW)

	g.pointShader, err = webgl.CreateProgram(
		g.w,
		`
    uniform vec2 uCenter;
    uniform vec2 uScale; 

    attribute vec2 aVertexPosition;

    void main() {
      gl_Position = vec4((aVertexPosition-uCenter)/uScale, 0.0, 1.0);
    }`,
		`
    precision highp float;

    void main() {
      gl_FragColor = vec4(1.0, 1.0, 1.0, 1.0);
    }
    `,
	)
	if err != nil {
		return nil, fmt.Errorf("Error building spriteShader: %w", err)
	}

	g.pointCoords = make([]float32, len(g.coords))

	return g, nil
}

func (g *graphics) Flush() {
	if g.pointsWritten > 0 {
		g.w.UseProgram(g.pointShader)

		uCenter := g.w.GetUniformLocation(g.pointShader, "uCenter")
		g.w.Uniform2fv(uCenter, g.uCenter)

		uScale := g.w.GetUniformLocation(g.pointShader, "uScale")
		g.w.Uniform2fv(uScale, g.uScale)

		g.w.BindBuffer(g.w.ARRAY_BUFFER, g.coordsBuffer)
		g.w.BufferSubDataF32(g.w.ARRAY_BUFFER, 0, g.pointCoords[:g.pointsWritten])
		aVertexPosition := g.w.GetAttribLocation(g.pointShader, "aVertexPosition")
		g.w.EnableVertexAttribArray(aVertexPosition)
		// Bind current array buffer to the given vertex attribute
		g.w.VertexAttribPointer(aVertexPosition, 2, g.w.FLOAT, false, 0, 0) // 2 = points per vertex

		g.w.DrawArrays(g.w.POINTS, 0, g.pointsWritten/2) // two floats per vertex

		g.pointsWritten = 0
	}

	if g.written > 0 {
		g.w.UseProgram(g.spriteShader)

		uCenter := g.w.GetUniformLocation(g.spriteShader, "uCenter")
		g.w.Uniform2fv(uCenter, g.uCenter)

		uScale := g.w.GetUniformLocation(g.spriteShader, "uScale")
		g.w.Uniform2fv(uScale, g.uScale)

		g.w.BindBuffer(g.w.ARRAY_BUFFER, g.coordsBuffer)
		g.w.BufferSubDataF32(g.w.ARRAY_BUFFER, 0, g.coords[:g.written])
		aVertexPosition := g.w.GetAttribLocation(g.spriteShader, "aVertexPosition")
		g.w.EnableVertexAttribArray(aVertexPosition)
		// Bind current array buffer to the given vertex attribute
		g.w.VertexAttribPointer(aVertexPosition, 2, g.w.FLOAT, false, 0, 0) // 2 = points per vertex

		g.w.BindBuffer(g.w.ARRAY_BUFFER, g.textureCoordsBuffer)
		g.w.BufferSubDataF32(g.w.ARRAY_BUFFER, 0, g.textureCoords[:g.written])
		aTextureCoord := g.w.GetAttribLocation(g.spriteShader, "aTextureCoord")
		g.w.EnableVertexAttribArray(aTextureCoord)
		// Bind current array buffer to the given vertex attribute
		g.w.VertexAttribPointer(aTextureCoord, 2, g.w.FLOAT, false, 0, 0) // 2 = points per vertex

		g.w.ActiveTexture(g.w.TEXTURE0)
		g.w.BindTexture(g.w.TEXTURE_2D, g.spritesheet)
		uSampler := g.w.GetUniformLocation(g.spriteShader, "uSampler")
		g.w.Uniform1i(uSampler, 0)

		g.w.DrawArrays(g.w.TRIANGLES, 0, g.written/2) // two floats per vertex

		g.written = 0
	}

	glError := g.w.GetError()
	if glError.Int() != 0 {
		log.Println("GL Error:", g.w.GetError())
	}
}

type Sprite struct {
	textureCoords []float32
	size          float32
}

var spritemap = map[game.Sprite]*Sprite{
	game.SpriteShip: &Sprite{
		textureCoords: genTexCoords(0, 0, 512, 512),
		size:          1,
	},
	game.SpriteEnemyShip: &Sprite{
		textureCoords: genTexCoords(512, 512, 1024, 1024),
		size:          1,
	},
	game.SpriteStar: &Sprite{
		textureCoords: genTexCoords(0, 512, 512, 1024),
		size:          10,
	},
	game.SpriteExplosionFlash: &Sprite{
		textureCoords: genTexCoords(0, 512, 512, 1024),
		// *2 because radius not diameter, *2 because the circle only takes up half
		// the sprite texture size.  Except that seems to big??
		size: game.ExplosionRadius * 2, // * 2 * 2,
	},
	game.SpriteMissile: &Sprite{
		textureCoords: genTexCoords(512, 0, 1024, 512),
		size:          1,
	},
}

func genTexCoords(xStart, yStart, xEnd, yEnd float32) []float32 {
	const width = 2048
	const height = width

	xStart /= width
	xEnd /= width
	yStart /= height
	yEnd /= height

	return []float32{
		xStart, yStart,
		xStart, yEnd,
		xEnd, yEnd,
		xStart, yStart,
		xEnd, yEnd,
		xEnd, yStart,
	}
}

//
//  0 1 : 6 7 ----------    : 10 11
//      |     \_            |
//      |       \_          |
//      |         \_        |
//      |           \_      |
//      |             \_    |
//  2 3 :     ----------4 5 : 8 9
//

func (g *graphics) Point(x, y float32) {
	g.pointCoords[g.pointsWritten] = x
	g.pointCoords[g.pointsWritten+1] = y
	g.pointsWritten += 2

	if g.pointsWritten >= len(g.pointCoords) {
		g.Flush()
	}
}

// TODO: spriteId -> size and texture location, rotation
func (g *graphics) Sprite(s *Sprite, centerx, centery, rotation float32) {
	coords := g.coords[g.written : g.written+12]
	textureCoords := g.textureCoords[g.written : g.written+12]

	cosSize := s.size * float32(math.Cos(float64(rotation)+(math.Pi/4))) * math.Sqrt2 / 2
	sinSize := s.size * float32(math.Sin(float64(rotation)+(math.Pi/4))) * math.Sqrt2 / 2

	coords[0] = centerx - sinSize
	coords[1] = centery + cosSize

	coords[2] = centerx - cosSize
	coords[3] = centery - sinSize

	coords[4] = centerx + sinSize
	coords[5] = centery - cosSize
	coords[6] = coords[0]
	coords[7] = coords[1]

	coords[8] = coords[4]
	coords[9] = coords[5]

	coords[10] = centerx + cosSize
	coords[11] = centery + sinSize
	for i := 0; i < 12; i++ {
		textureCoords[i] = s.textureCoords[i]
	}

	g.written += 12

	if g.written >= len(g.coords) {
		g.Flush()
	}

}

func (g *graphics) Clear() {
	g.w.ClearColor(0.039, 0.102, 0.247, 1)
	g.w.Clear(g.w.COLOR_BUFFER_BIT)
}

func (g *graphics) SetCamera(xmin, ymin, xmax, ymax float32) {
	g.uCenter[0] = (xmin + xmax) / 2
	g.uCenter[1] = (ymin + ymax) / 2
	g.uScale[0] = (xmax - xmin) / 2
	g.uScale[1] = (ymax - ymin) / 2

	aspectRatio := float32(g.width) / float32(g.height)

	if g.uScale[0]/aspectRatio > g.uScale[1] {
		g.uScale[1] = g.uScale[0] / aspectRatio
	} else {
		g.uScale[0] = g.uScale[1] * aspectRatio
	}
}
