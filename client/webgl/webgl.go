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

package webgl

import (
	"errors"
	"fmt"
	"math"
	"syscall/js"
)

type Capability js.Value
type BufferBit int
type BufferTarget js.Value
type BufferUsage js.Value
type TextureTarget js.Value
type TextureFormat js.Value
type TextureUnit js.Value
type BlendFactor js.Value
type ShaderType js.Value
type ProgramParameterBool js.Value
type ShaderParameterBool js.Value
type DrawMode js.Value
type GlType js.Value

type Buffer js.Value
type Texture js.Value
type Program js.Value
type Shader js.Value
type UniformLocation js.Value
type AttribLocation js.Value

type WebGL struct {
	canvas, gl js.Value

	BLEND Capability

	COLOR_BUFFER_BIT BufferBit
	DEPTH_BUFFER_BIT BufferBit

	ARRAY_BUFFER BufferTarget
	STATIC_DRAW  BufferUsage
	DYNAMIC_DRAW BufferUsage

	TEXTURE_2D TextureTarget
	RGBA       TextureFormat
	TEXTURE0   TextureUnit

	ZERO                BlendFactor
	ONE                 BlendFactor
	SRC_ALPHA           BlendFactor
	ONE_MINUS_SRC_ALPHA BlendFactor

	VERTEX_SHADER   ShaderType
	FRAGMENT_SHADER ShaderType
	LINK_STATUS     ProgramParameterBool
	COMPILE_STATUS  ShaderParameterBool

	POINTS    DrawMode
	TRIANGLES DrawMode

	UNSIGNED_BYTE GlType
	FLOAT         GlType
	SHORT         GlType
}

func InitWebgl(canvas js.Value) (*WebGL, error) {
	w := WebGL{}

	w.canvas = canvas
	w.gl = canvas.Call("getContext", "webgl")

	// Post 1.14, this can be replaced with w.gl.IsNull()
	if w.gl.Type() == js.TypeNull {
		return nil, errors.New("Creating a webgl context is not supported.  This won't work.")
	}

	w.BLEND = Capability(w.gl.Get("BLEND"))

	w.COLOR_BUFFER_BIT = BufferBit(w.gl.Get("COLOR_BUFFER_BIT").Int())
	w.DEPTH_BUFFER_BIT = BufferBit(w.gl.Get("DEPTH_BUFFER_BIT").Int())

	w.ARRAY_BUFFER = BufferTarget(w.gl.Get("ARRAY_BUFFER"))

	w.STATIC_DRAW = BufferUsage(w.gl.Get("STATIC_DRAW"))
	w.DYNAMIC_DRAW = BufferUsage(w.gl.Get("DYNAMIC_DRAW"))

	w.TEXTURE_2D = TextureTarget(w.gl.Get("TEXTURE_2D"))
	w.RGBA = TextureFormat(w.gl.Get("RGBA"))
	w.TEXTURE0 = TextureUnit(w.gl.Get("TEXTURE0"))

	w.ZERO = BlendFactor(w.gl.Get("ZERO"))
	w.ONE = BlendFactor(w.gl.Get("ONE"))
	w.SRC_ALPHA = BlendFactor(w.gl.Get("SRC_ALPHA"))
	w.ONE_MINUS_SRC_ALPHA = BlendFactor(w.gl.Get("ONE_MINUS_SRC_ALPHA"))

	w.VERTEX_SHADER = ShaderType(w.gl.Get("VERTEX_SHADER"))
	w.FRAGMENT_SHADER = ShaderType(w.gl.Get("FRAGMENT_SHADER"))

	w.LINK_STATUS = ProgramParameterBool(w.gl.Get("LINK_STATUS"))
	w.COMPILE_STATUS = ShaderParameterBool(w.gl.Get("COMPILE_STATUS"))

	w.POINTS = DrawMode(w.gl.Get("POINTS"))
	w.TRIANGLES = DrawMode(w.gl.Get("TRIANGLES"))

	w.FLOAT = GlType(w.gl.Get("FLOAT"))
	w.SHORT = GlType(w.gl.Get("SHORT"))
	w.UNSIGNED_BYTE = GlType(w.gl.Get("UNSIGNED_BYTE"))

	return &w, nil
}

func (w *WebGL) ClearColor(r, g, b, a float64) {
	w.gl.Call("clearColor", r, g, b, a)
}

func (w *WebGL) Clear(colorBits BufferBit) {
	w.gl.Call("clear", int(colorBits))
}

func (w *WebGL) Enable(c Capability) {
	w.gl.Call("enable", js.Value(c))
}

/////////////////////////////////////////////
// Buffer
/////////////////////////////////////////////

func (w *WebGL) CreateBuffer() *Buffer {
	b := Buffer(w.gl.Call("createBuffer"))
	return &b
}

func (w *WebGL) BindBuffer(t BufferTarget, b *Buffer) {
	w.gl.Call("bindBuffer", js.Value(t), js.Value(*b))
}

func (w *WebGL) BufferDataF32(t BufferTarget, a []float32, u BufferUsage) {
	w.gl.Call("bufferData", js.Value(t), copyFloat32SliceToJS(a), js.Value(u))
}

func (w *WebGL) BufferDataSize(t BufferTarget, size int, u BufferUsage) {
	w.gl.Call("bufferData", js.Value(t), size, js.Value(u))
}

func (w *WebGL) BufferSubDataI16(t BufferTarget, offset int, a []int16) {
	w.gl.Call("bufferSubData", js.Value(t), offset, copyInt16SliceToJS(a))
}

func (w *WebGL) BufferSubDataF32(t BufferTarget, offset int, a []float32) {
	w.gl.Call("bufferSubData", js.Value(t), offset, copyFloat32SliceToJS(a))
}

/////////////////////////////////////////////
// Texture
/////////////////////////////////////////////

func (w *WebGL) CreateTexture() *Texture {
	t := Texture(w.gl.Call("createTexture"))
	return &t
}

func (w *WebGL) BindTexture(tt TextureTarget, t *Texture) {
	w.gl.Call("bindTexture", js.Value(tt), js.Value(*t))
}

func (w *WebGL) TexImage2D(target TextureTarget, level int, internalFormat TextureFormat, format TextureFormat, type_ GlType, pixels js.Value) {
	w.gl.Call("texImage2D", js.Value(target), level, js.Value(internalFormat), js.Value(format), js.Value(type_), pixels)
}

func (w *WebGL) GenerateMipmap(target TextureTarget) {
	w.gl.Call("generateMipmap", js.Value(target))
}

func (w *WebGL) ActiveTexture(u TextureUnit) {
	w.gl.Call("activeTexture", js.Value(u))
}

/////////////////////////////////////////////
// Program
/////////////////////////////////////////////

func (w *WebGL) CreateProgram() *Program {
	p := Program(w.gl.Call("createProgram"))
	return &p
}

func (w *WebGL) AttachShader(p *Program, s *Shader) {
	w.gl.Call("attachShader", js.Value(*p), js.Value(*s))
}

func (w *WebGL) LinkProgram(p *Program) {
	w.gl.Call("linkProgram", js.Value(*p))
}

func (w *WebGL) ValidateProgram(p *Program) {
	w.gl.Call("validateProgram", js.Value(*p))
}

func (w *WebGL) GetProgramParameterBool(p *Program, param ProgramParameterBool) bool {
	return w.gl.Call("getProgramParameter", js.Value(*p), js.Value(param)).Bool()
}

func (w *WebGL) GetProgramInfoLog(p *Program) string {
	return w.gl.Call("getProgramInfoLog", js.Value(*p)).String()
}

func (w *WebGL) UseProgram(p *Program) {
	w.gl.Call("useProgram", js.Value(*p))
}

func (w *WebGL) GetUniformLocation(p *Program, name string) UniformLocation {
	return UniformLocation(w.gl.Call("getUniformLocation", js.Value(*p), name))
}

func (w *WebGL) GetAttribLocation(p *Program, name string) AttribLocation {
	return AttribLocation(w.gl.Call("getAttribLocation", js.Value(*p), name))
}

/////////////////////////////////////////////
// Shader
/////////////////////////////////////////////

func (w *WebGL) CreateShader(t ShaderType) *Shader {
	s := Shader(w.gl.Call("createShader", js.Value(t)))
	return &s
}

func (w *WebGL) ShaderSource(s *Shader, code string) {
	w.gl.Call("shaderSource", js.Value(*s), code)
}

func (w *WebGL) CompileShader(s *Shader) {
	w.gl.Call("compileShader", js.Value(*s))
}

func (w *WebGL) GetShaderParameterBool(s *Shader, param ShaderParameterBool) bool {
	return w.gl.Call("getShaderParameter", js.Value(*s), js.Value(param)).Bool()
}

func (w *WebGL) GetShaderInfoLog(s *Shader) string {
	return w.gl.Call("getShaderInfoLog", js.Value(*s)).String()
}

/////////////////////////////////////////////
// Shader parameters and draw calls
/////////////////////////////////////////////

func (w *WebGL) Uniform1i(u UniformLocation, v int) {
	w.gl.Call("uniform1i", js.Value(u), v)
}

func (w *WebGL) Uniform2fv(u UniformLocation, v [2]float32) {
	w.gl.Call("uniform2fv", js.Value(u), copyFloat32SliceToJS(v[:]))
}

func (w *WebGL) Uniform4fv(u UniformLocation, v [4]float32) {
	w.gl.Call("uniform4fv", js.Value(u), copyFloat32SliceToJS(v[:]))
}

func (w *WebGL) EnableVertexAttribArray(a AttribLocation) {
	w.gl.Call("enableVertexAttribArray", js.Value(a))
}

func (w *WebGL) VertexAttribPointer(
	a AttribLocation, size int, arrayType GlType, normalized bool, stride, offset int64) {

	w.gl.Call("vertexAttribPointer", js.Value(a), size, js.Value(arrayType), normalized, stride, offset)
}

func (w *WebGL) DrawArrays(mode DrawMode, first, size int) {
	w.gl.Call("drawArrays", js.Value(mode), first, size)
}

func (w *WebGL) GetError() js.Value {
	return w.gl.Call("getError")
}

func (w *WebGL) BlendFunc(sfactor, dfactor BlendFactor) {
	w.gl.Call("blendFunc", js.Value(sfactor), js.Value(dfactor))
}

/////////////////////////////////////////////
// Utility functions (anything that wouldn't be on an OpenGL spec)
/////////////////////////////////////////////

func CreateProgram(w *WebGL, vertexShaderCode, fragmentShaderCode string) (*Program, error) {
	vertexShader, err := CompileShader(w, vertexShaderCode, w.VERTEX_SHADER)
	if err != nil {
		return nil, fmt.Errorf("Error compiling vertex shader: %s", err)
	}

	fragmentShader, err := CompileShader(w, fragmentShaderCode, w.FRAGMENT_SHADER)
	if err != nil {
		return nil, fmt.Errorf("Error compiling fragment shader: %s", err)
	}

	p := w.CreateProgram()
	w.AttachShader(p, vertexShader)
	w.AttachShader(p, fragmentShader)
	w.LinkProgram(p)
	w.ValidateProgram(p)

	if !w.GetProgramParameterBool(p, w.LINK_STATUS) {
		return nil, fmt.Errorf("Error linking shader: %s", err)
	}

	return p, nil
}

func CompileShader(w *WebGL, code string, t ShaderType) (*Shader, error) {
	s := w.CreateShader(t)

	w.ShaderSource(s, code)
	w.CompileShader(s)

	if !w.GetShaderParameterBool(s, w.COMPILE_STATUS) {
		return nil, fmt.Errorf("Error Compiling Shader: %s", w.GetShaderInfoLog(s))
	}
	return s, nil
}

func init() {
	if copyFloat32SliceToJS([]float32{5}).Index(0).Float() == 5 {
		return
	}
	swapEndianess = true
	if copyFloat32SliceToJS([]float32{5}).Index(0).Float() == 5 {
		return
	}
	panic("Could not determine endianess.")
}

var swapEndianess bool = false

const MaxArrayLength = 3 * 2 * 2 * 1000

const bytesPerFloatValue = 4

var preallocatedBuffer = js.Global().Get("ArrayBuffer").New(MaxArrayLength * bytesPerFloatValue)

func copyFloat32SliceToJS(a []float32) js.Value {
	if len(a) > MaxArrayLength {
		panic("Array too long to fit in pre-allocated buffer!")
	}

	// TODO: This is much faster, but could it be even better?  Would it be
	// possible to not do any adding to a buffer on the Go side, but from js copy
	// the float memory segment directly? (and then do any endian swapping.)
	// TODO: Pool a large buffer, do in chunks if not large enough?
	b := make([]byte, len(a)*bytesPerFloatValue)
	bi := 0
	for _, v := range a {
		vb := math.Float32bits(v)
		if swapEndianess {
			b[bi] = byte(vb)
			bi++
			b[bi] = byte(vb >> 8)
			bi++
			b[bi] = byte(vb >> 16)
			bi++
			b[bi] = byte(vb >> 24)
			bi++
		} else {
			b[bi] = byte(vb >> 24)
			bi++
			b[bi] = byte(vb >> 16)
			bi++
			b[bi] = byte(vb >> 8)
			bi++
			b[bi] = byte(vb)
			bi++
		}
	}
	js.CopyBytesToJS(js.Global().Get("Uint8Array").New(preallocatedBuffer), b)
	return js.Global().Get("Float32Array").New(preallocatedBuffer, 0, len(a))
}

func copyInt16SliceToJS(a []int16) js.Value {
	r := js.Global().Get("Int16Array").New(len(a))
	for i, v := range a {
		r.SetIndex(i, js.ValueOf(v))
	}
	return r
}

func LoadTexture(w *WebGL, image js.Value) *Texture {
	t := w.CreateTexture()
	w.BindTexture(w.TEXTURE_2D, t)
	w.TexImage2D(w.TEXTURE_2D, 0, w.RGBA, w.RGBA, w.UNSIGNED_BYTE, image)
	w.GenerateMipmap(w.TEXTURE_2D)
	return t
}
