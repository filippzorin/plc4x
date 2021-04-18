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
type ModbusPDUWriteFileRecordRequest struct {
	Items  []*ModbusPDUWriteFileRecordRequestItem
	Parent *ModbusPDU
}

// The corresponding interface
type IModbusPDUWriteFileRecordRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUWriteFileRecordRequest) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUWriteFileRecordRequest) FunctionFlag() uint8 {
	return 0x15
}

func (m *ModbusPDUWriteFileRecordRequest) Response() bool {
	return false
}

func (m *ModbusPDUWriteFileRecordRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUWriteFileRecordRequest(items []*ModbusPDUWriteFileRecordRequestItem) *ModbusPDU {
	child := &ModbusPDUWriteFileRecordRequest{
		Items:  items,
		Parent: NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUWriteFileRecordRequest(structType interface{}) *ModbusPDUWriteFileRecordRequest {
	castFunc := func(typ interface{}) *ModbusPDUWriteFileRecordRequest {
		if casted, ok := typ.(ModbusPDUWriteFileRecordRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUWriteFileRecordRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUWriteFileRecordRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUWriteFileRecordRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUWriteFileRecordRequest) GetTypeName() string {
	return "ModbusPDUWriteFileRecordRequest"
}

func (m *ModbusPDUWriteFileRecordRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUWriteFileRecordRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.LengthInBits()
		}
	}

	return lengthInBits
}

func (m *ModbusPDUWriteFileRecordRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUWriteFileRecordRequestParse(io utils.ReadBuffer) (*ModbusPDU, error) {

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := io.ReadUint8(8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Array field (items)
	// Length array
	items := make([]*ModbusPDUWriteFileRecordRequestItem, 0)
	_itemsLength := byteCount
	_itemsEndPos := io.GetPos() + uint16(_itemsLength)
	for io.GetPos() < _itemsEndPos {
		_item, _err := ModbusPDUWriteFileRecordRequestItemParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items = append(items, _item)
	}

	// Create a partially initialized instance
	_child := &ModbusPDUWriteFileRecordRequest{
		Items:  items,
		Parent: &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUWriteFileRecordRequest) Serialize(io utils.WriteBuffer) error {
	itemsArraySizeInBytes := func(items []*ModbusPDUWriteFileRecordRequestItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.LengthInBytes())
		}
		return sizeInBytes
	}
	ser := func() error {
		io.PushContext("ModbusPDUWriteFileRecordRequest")

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.Items)))
		_byteCountErr := io.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (items)
		if m.Items != nil {
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
		}

		io.PopContext("ModbusPDUWriteFileRecordRequest")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUWriteFileRecordRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "items":
				var data []*ModbusPDUWriteFileRecordRequestItem
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Items = data
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

func (m *ModbusPDUWriteFileRecordRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.Items {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "items"}}); err != nil {
		return err
	}
	return nil
}

func (m ModbusPDUWriteFileRecordRequest) String() string {
	return string(m.Box("", 120))
}

func (m ModbusPDUWriteFileRecordRequest) Box(name string, width int) utils.AsciiBox {
	boxName := "ModbusPDUWriteFileRecordRequest"
	if name != "" {
		boxName += "/" + name
	}
	itemsArraySizeInBytes := func(items []*ModbusPDUWriteFileRecordRequestItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.LengthInBytes())
		}
		return sizeInBytes
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Implicit Field (byteCount)
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.Items)))
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("ByteCount", byteCount, -1))
		// Array Field (items)
		if m.Items != nil {
			// Complex array base type
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Items {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Items", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
