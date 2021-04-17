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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const BVLC_BACNETTYPE uint8 = 0x81

// The data-structure of this message
type BVLC struct {
	Child IBVLCChild
}

// The corresponding interface
type IBVLC interface {
	BvlcFunction() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IBVLCParent interface {
	SerializeParent(io utils.WriteBuffer, child IBVLC, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBVLCChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *BVLC)
	GetTypeName() string
	IBVLC
	utils.AsciiBoxer
}

func NewBVLC() *BVLC {
	return &BVLC{}
}

func CastBVLC(structType interface{}) *BVLC {
	castFunc := func(typ interface{}) *BVLC {
		if casted, ok := typ.(BVLC); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLC); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLC) GetTypeName() string {
	return "BVLC"
}

func (m *BVLC) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLC) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BVLC) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (bacnetType)
	lengthInBits += 8
	// Discriminator Field (bvlcFunction)
	lengthInBits += 8

	// Implicit Field (bvlcLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *BVLC) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCParse(io *utils.ReadBuffer) (*BVLC, error) {

	// Const Field (bacnetType)
	bacnetType, _bacnetTypeErr := io.ReadUint8(8)
	if _bacnetTypeErr != nil {
		return nil, errors.Wrap(_bacnetTypeErr, "Error parsing 'bacnetType' field")
	}
	if bacnetType != BVLC_BACNETTYPE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BVLC_BACNETTYPE) + " but got " + fmt.Sprintf("%d", bacnetType))
	}

	// Discriminator Field (bvlcFunction) (Used as input to a switch field)
	bvlcFunction, _bvlcFunctionErr := io.ReadUint8(8)
	if _bvlcFunctionErr != nil {
		return nil, errors.Wrap(_bvlcFunctionErr, "Error parsing 'bvlcFunction' field")
	}

	// Implicit Field (bvlcLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	bvlcLength, _bvlcLengthErr := io.ReadUint16(16)
	_ = bvlcLength
	if _bvlcLengthErr != nil {
		return nil, errors.Wrap(_bvlcLengthErr, "Error parsing 'bvlcLength' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BVLC
	var typeSwitchError error
	switch {
	case bvlcFunction == 0x00: // BVLCResult
		_parent, typeSwitchError = BVLCResultParse(io)
	case bvlcFunction == 0x01: // BVLCWideBroadcastDistributionTable
		_parent, typeSwitchError = BVLCWideBroadcastDistributionTableParse(io)
	case bvlcFunction == 0x02: // BVLCReadBroadcastDistributionTable
		_parent, typeSwitchError = BVLCReadBroadcastDistributionTableParse(io)
	case bvlcFunction == 0x03: // BVLCReadBroadcastDistributionTableAck
		_parent, typeSwitchError = BVLCReadBroadcastDistributionTableAckParse(io)
	case bvlcFunction == 0x04: // BVLCForwardedNPDU
		_parent, typeSwitchError = BVLCForwardedNPDUParse(io, bvlcLength)
	case bvlcFunction == 0x05: // BVLCRegisterForeignDevice
		_parent, typeSwitchError = BVLCRegisterForeignDeviceParse(io)
	case bvlcFunction == 0x06: // BVLCReadForeignDeviceTable
		_parent, typeSwitchError = BVLCReadForeignDeviceTableParse(io)
	case bvlcFunction == 0x07: // BVLCReadForeignDeviceTableAck
		_parent, typeSwitchError = BVLCReadForeignDeviceTableAckParse(io)
	case bvlcFunction == 0x08: // BVLCDeleteForeignDeviceTableEntry
		_parent, typeSwitchError = BVLCDeleteForeignDeviceTableEntryParse(io)
	case bvlcFunction == 0x09: // BVLCDistributeBroadcastToNetwork
		_parent, typeSwitchError = BVLCDistributeBroadcastToNetworkParse(io)
	case bvlcFunction == 0x0A: // BVLCOriginalUnicastNPDU
		_parent, typeSwitchError = BVLCOriginalUnicastNPDUParse(io, bvlcLength)
	case bvlcFunction == 0x0B: // BVLCOriginalBroadcastNPDU
		_parent, typeSwitchError = BVLCOriginalBroadcastNPDUParse(io, bvlcLength)
	case bvlcFunction == 0x0C: // BVLCSecureBVLL
		_parent, typeSwitchError = BVLCSecureBVLLParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *BVLC) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *BVLC) SerializeParent(io utils.WriteBuffer, child IBVLC, serializeChildFunction func() error) error {

	// Const Field (bacnetType)
	_bacnetTypeErr := io.WriteUint8(8, 0x81)
	if _bacnetTypeErr != nil {
		return errors.Wrap(_bacnetTypeErr, "Error serializing 'bacnetType' field")
	}

	// Discriminator Field (bvlcFunction) (Used as input to a switch field)
	bvlcFunction := uint8(child.BvlcFunction())
	_bvlcFunctionErr := io.WriteUint8(8, (bvlcFunction))

	if _bvlcFunctionErr != nil {
		return errors.Wrap(_bvlcFunctionErr, "Error serializing 'bvlcFunction' field")
	}

	// Implicit Field (bvlcLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	bvlcLength := uint16(uint16(m.LengthInBytes()))
	_bvlcLengthErr := io.WriteUint16(16, (bvlcLength))
	if _bvlcLengthErr != nil {
		return errors.Wrap(_bvlcLengthErr, "Error serializing 'bvlcLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	return nil
}

func (m *BVLC) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// BVLCResult needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BVLCResult":
			if m.Child == nil {
				m.Child = &BVLCResult{
					Parent: m,
				}
			}
		// BVLCWideBroadcastDistributionTable needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BVLCWideBroadcastDistributionTable":
			if m.Child == nil {
				m.Child = &BVLCWideBroadcastDistributionTable{
					Parent: m,
				}
			}
		// BVLCReadBroadcastDistributionTable needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadBroadcastDistributionTable":
			if m.Child == nil {
				m.Child = &BVLCReadBroadcastDistributionTable{
					Parent: m,
				}
			}
		// BVLCReadBroadcastDistributionTableAck needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadBroadcastDistributionTableAck":
			if m.Child == nil {
				m.Child = &BVLCReadBroadcastDistributionTableAck{
					Parent: m,
				}
			}
		// BVLCRegisterForeignDevice needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BVLCRegisterForeignDevice":
			if m.Child == nil {
				m.Child = &BVLCRegisterForeignDevice{
					Parent: m,
				}
			}
		// BVLCReadForeignDeviceTable needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadForeignDeviceTable":
			if m.Child == nil {
				m.Child = &BVLCReadForeignDeviceTable{
					Parent: m,
				}
			}
		// BVLCReadForeignDeviceTableAck needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadForeignDeviceTableAck":
			if m.Child == nil {
				m.Child = &BVLCReadForeignDeviceTableAck{
					Parent: m,
				}
			}
		// BVLCDeleteForeignDeviceTableEntry needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BVLCDeleteForeignDeviceTableEntry":
			if m.Child == nil {
				m.Child = &BVLCDeleteForeignDeviceTableEntry{
					Parent: m,
				}
			}
		// BVLCDistributeBroadcastToNetwork needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BVLCDistributeBroadcastToNetwork":
			if m.Child == nil {
				m.Child = &BVLCDistributeBroadcastToNetwork{
					Parent: m,
				}
			}
		// BVLCSecureBVLL needs special treatment as it has no fields
		case "org.apache.plc4x.java.bacnetip.readwrite.BVLCSecureBVLL":
			if m.Child == nil {
				m.Child = &BVLCSecureBVLL{
					Parent: m,
				}
			}
		}
	}
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of BVLC")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCResult":
					var dt *BVLCResult
					if m.Child != nil {
						dt = m.Child.(*BVLCResult)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCWideBroadcastDistributionTable":
					var dt *BVLCWideBroadcastDistributionTable
					if m.Child != nil {
						dt = m.Child.(*BVLCWideBroadcastDistributionTable)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadBroadcastDistributionTable":
					var dt *BVLCReadBroadcastDistributionTable
					if m.Child != nil {
						dt = m.Child.(*BVLCReadBroadcastDistributionTable)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadBroadcastDistributionTableAck":
					var dt *BVLCReadBroadcastDistributionTableAck
					if m.Child != nil {
						dt = m.Child.(*BVLCReadBroadcastDistributionTableAck)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCForwardedNPDU":
					var dt *BVLCForwardedNPDU
					if m.Child != nil {
						dt = m.Child.(*BVLCForwardedNPDU)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCRegisterForeignDevice":
					var dt *BVLCRegisterForeignDevice
					if m.Child != nil {
						dt = m.Child.(*BVLCRegisterForeignDevice)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadForeignDeviceTable":
					var dt *BVLCReadForeignDeviceTable
					if m.Child != nil {
						dt = m.Child.(*BVLCReadForeignDeviceTable)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadForeignDeviceTableAck":
					var dt *BVLCReadForeignDeviceTableAck
					if m.Child != nil {
						dt = m.Child.(*BVLCReadForeignDeviceTableAck)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCDeleteForeignDeviceTableEntry":
					var dt *BVLCDeleteForeignDeviceTableEntry
					if m.Child != nil {
						dt = m.Child.(*BVLCDeleteForeignDeviceTableEntry)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCDistributeBroadcastToNetwork":
					var dt *BVLCDistributeBroadcastToNetwork
					if m.Child != nil {
						dt = m.Child.(*BVLCDistributeBroadcastToNetwork)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCOriginalUnicastNPDU":
					var dt *BVLCOriginalUnicastNPDU
					if m.Child != nil {
						dt = m.Child.(*BVLCOriginalUnicastNPDU)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCOriginalBroadcastNPDU":
					var dt *BVLCOriginalBroadcastNPDU
					if m.Child != nil {
						dt = m.Child.(*BVLCOriginalBroadcastNPDU)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.bacnetip.readwrite.BVLCSecureBVLL":
					var dt *BVLCSecureBVLL
					if m.Child != nil {
						dt = m.Child.(*BVLCSecureBVLL)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				}
			}
		}
	}
}

func (m *BVLC) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.bacnetip.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m BVLC) String() string {
	return string(m.Box("BVLC", utils.DefaultWidth*2))
}

func (m *BVLC) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

func (m *BVLC) BoxParent(name string, width int, boxChild func() []utils.AsciiBox) utils.AsciiBox {
	if name == "" {
		name = "BVLC"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, boxChild()...)
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
