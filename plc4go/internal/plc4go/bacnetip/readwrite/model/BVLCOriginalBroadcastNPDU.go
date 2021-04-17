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
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BVLCOriginalBroadcastNPDU struct {
	Npdu   *NPDU
	Parent *BVLC
}

// The corresponding interface
type IBVLCOriginalBroadcastNPDU interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCOriginalBroadcastNPDU) BvlcFunction() uint8 {
	return 0x0B
}

func (m *BVLCOriginalBroadcastNPDU) InitializeParent(parent *BVLC) {
}

func NewBVLCOriginalBroadcastNPDU(npdu *NPDU) *BVLC {
	child := &BVLCOriginalBroadcastNPDU{
		Npdu:   npdu,
		Parent: NewBVLC(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBVLCOriginalBroadcastNPDU(structType interface{}) *BVLCOriginalBroadcastNPDU {
	castFunc := func(typ interface{}) *BVLCOriginalBroadcastNPDU {
		if casted, ok := typ.(BVLCOriginalBroadcastNPDU); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCOriginalBroadcastNPDU); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCOriginalBroadcastNPDU(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCOriginalBroadcastNPDU(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCOriginalBroadcastNPDU) GetTypeName() string {
	return "BVLCOriginalBroadcastNPDU"
}

func (m *BVLCOriginalBroadcastNPDU) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCOriginalBroadcastNPDU) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (npdu)
	lengthInBits += m.Npdu.LengthInBits()

	return lengthInBits
}

func (m *BVLCOriginalBroadcastNPDU) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCOriginalBroadcastNPDUParse(io *utils.ReadBuffer, bvlcLength uint16) (*BVLC, error) {

	// Simple Field (npdu)
	npdu, _npduErr := NPDUParse(io, uint16(bvlcLength)-uint16(uint16(4)))
	if _npduErr != nil {
		return nil, errors.Wrap(_npduErr, "Error parsing 'npdu' field")
	}

	// Create a partially initialized instance
	_child := &BVLCOriginalBroadcastNPDU{
		Npdu:   npdu,
		Parent: &BVLC{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BVLCOriginalBroadcastNPDU) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (npdu)
		_npduErr := m.Npdu.Serialize(io)
		if _npduErr != nil {
			return errors.Wrap(_npduErr, "Error serializing 'npdu' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BVLCOriginalBroadcastNPDU) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "npdu":
				var data NPDU
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Npdu = &data
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

func (m *BVLCOriginalBroadcastNPDU) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Npdu, xml.StartElement{Name: xml.Name{Local: "npdu"}}); err != nil {
		return err
	}
	return nil
}

func (m BVLCOriginalBroadcastNPDU) String() string {
	return string(m.Box("BVLCOriginalBroadcastNPDU", utils.DefaultWidth*2))
}

func (m BVLCOriginalBroadcastNPDU) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "BVLCOriginalBroadcastNPDU"
	}
	boxChild := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		boxes = append(boxes, utils.BoxAnything("Npdu", m.Npdu, width-2))
		return boxes
	}
	return m.Parent.BoxParent(name, width, boxChild)
}
