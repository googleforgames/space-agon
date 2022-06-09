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

package protostream

import (
	"github.com/golang/protobuf/proto"
	"io"
)

type ReaderWriter interface {
	io.Reader
	io.Writer
}

type ProtoStream struct {
	rw   ReaderWriter
	b    []byte
	read int
}

func NewProtoStream(rw ReaderWriter) *ProtoStream {
	return &ProtoStream{
		rw: rw,
		b:  make([]byte, 10), // 10 is enough to read the largest varint
	}
}

func (p *ProtoStream) Send(m proto.Message) error {
	b, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	b = append(proto.EncodeVarint(uint64(len(b))), b...)

	_, err = p.rw.Write(b)
	return err
}

func (p *ProtoStream) Recv(m proto.Message) error {
	vLength := 0
	mLength := uint64(0)
	for {
		mLength, vLength = proto.DecodeVarint(p.b[:p.read])
		if vLength != 0 {
			break
		}
		n, err := p.rw.Read(p.b[p.read:])
		if err != nil {
			return err
		}
		p.read += n
	}

	total := vLength + int(mLength)

	if len(p.b) < total {
		old := p.b[:p.read]
		p.b = make([]byte, total)
		copy(p.b, old)
	}
	for p.read < total {
		n, err := p.rw.Read(p.b[p.read:])
		if err != nil {
			return err
		}
		p.read += n
	}

	err := proto.Unmarshal(p.b[vLength:total], m)
	if err != nil {
		return err
	}

	// Move unused to start
	copy(p.b, p.b[total:p.read])
	p.read -= total

	return nil
}
