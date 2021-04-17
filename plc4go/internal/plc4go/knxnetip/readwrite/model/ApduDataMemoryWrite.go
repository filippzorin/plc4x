//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduDataMemoryWrite struct {
	Parent *ApduData
}

// The corresponding interface
type IApduDataMemoryWrite interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataMemoryWrite) ApciType() uint8 {
	return 0xA
}

func (m *ApduDataMemoryWrite) InitializeParent(parent *ApduData) {
}

func NewApduDataMemoryWrite() *ApduData {
	child := &ApduDataMemoryWrite{
		Parent: NewApduData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataMemoryWrite(structType interface{}) *ApduDataMemoryWrite {
	castFunc := func(typ interface{}) *ApduDataMemoryWrite {
		if casted, ok := typ.(ApduDataMemoryWrite); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataMemoryWrite); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataMemoryWrite(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataMemoryWrite(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataMemoryWrite) GetTypeName() string {
	return "ApduDataMemoryWrite"
}

func (m *ApduDataMemoryWrite) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataMemoryWrite) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataMemoryWrite) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataMemoryWriteParse(io *utils.ReadBuffer) (*ApduData, error) {

	// Create a partially initialized instance
	_child := &ApduDataMemoryWrite{
		Parent: &ApduData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataMemoryWrite) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataMemoryWrite) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *ApduDataMemoryWrite) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m ApduDataMemoryWrite) String() string {
	return string(m.Box("ApduDataMemoryWrite", utils.DefaultWidth*2))
}

func (m ApduDataMemoryWrite) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ApduDataMemoryWrite"
	}
	boxChild := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(name, width, boxChild)
}
