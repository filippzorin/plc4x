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
	"github.com/rs/zerolog/log"
	"io"
	"strconv"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type APDUComplexAck struct {
	SegmentedMessage   bool
	MoreFollows        bool
	OriginalInvokeId   uint8
	SequenceNumber     *uint8
	ProposedWindowSize *uint8
	ServiceAck         *BACnetServiceAck
	Parent             *APDU
}

// The corresponding interface
type IAPDUComplexAck interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUComplexAck) ApduType() uint8 {
	return 0x3
}

func (m *APDUComplexAck) InitializeParent(parent *APDU) {
}

func NewAPDUComplexAck(segmentedMessage bool, moreFollows bool, originalInvokeId uint8, sequenceNumber *uint8, proposedWindowSize *uint8, serviceAck *BACnetServiceAck) *APDU {
	child := &APDUComplexAck{
		SegmentedMessage:   segmentedMessage,
		MoreFollows:        moreFollows,
		OriginalInvokeId:   originalInvokeId,
		SequenceNumber:     sequenceNumber,
		ProposedWindowSize: proposedWindowSize,
		ServiceAck:         serviceAck,
		Parent:             NewAPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAPDUComplexAck(structType interface{}) *APDUComplexAck {
	castFunc := func(typ interface{}) *APDUComplexAck {
		if casted, ok := typ.(APDUComplexAck); ok {
			return &casted
		}
		if casted, ok := typ.(*APDUComplexAck); ok {
			return casted
		}
		if casted, ok := typ.(APDU); ok {
			return CastAPDUComplexAck(casted.Child)
		}
		if casted, ok := typ.(*APDU); ok {
			return CastAPDUComplexAck(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *APDUComplexAck) GetTypeName() string {
	return "APDUComplexAck"
}

func (m *APDUComplexAck) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *APDUComplexAck) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (segmentedMessage)
	lengthInBits += 1

	// Simple field (moreFollows)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Optional Field (sequenceNumber)
	if m.SequenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (proposedWindowSize)
	if m.ProposedWindowSize != nil {
		lengthInBits += 8
	}

	// Simple field (serviceAck)
	lengthInBits += m.ServiceAck.LengthInBits()

	return lengthInBits
}

func (m *APDUComplexAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func APDUComplexAckParse(io utils.ReadBuffer) (*APDU, error) {

	// Simple Field (segmentedMessage)
	segmentedMessage, _segmentedMessageErr := io.ReadBit()
	if _segmentedMessageErr != nil {
		return nil, errors.Wrap(_segmentedMessageErr, "Error parsing 'segmentedMessage' field")
	}

	// Simple Field (moreFollows)
	moreFollows, _moreFollowsErr := io.ReadBit()
	if _moreFollowsErr != nil {
		return nil, errors.Wrap(_moreFollowsErr, "Error parsing 'moreFollows' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8(2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (originalInvokeId)
	originalInvokeId, _originalInvokeIdErr := io.ReadUint8(8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}

	// Optional Field (sequenceNumber) (Can be skipped, if a given expression evaluates to false)
	var sequenceNumber *uint8 = nil
	if segmentedMessage {
		_val, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'sequenceNumber' field")
		}
		sequenceNumber = &_val
	}

	// Optional Field (proposedWindowSize) (Can be skipped, if a given expression evaluates to false)
	var proposedWindowSize *uint8 = nil
	if segmentedMessage {
		_val, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'proposedWindowSize' field")
		}
		proposedWindowSize = &_val
	}

	// Simple Field (serviceAck)
	serviceAck, _serviceAckErr := BACnetServiceAckParse(io)
	if _serviceAckErr != nil {
		return nil, errors.Wrap(_serviceAckErr, "Error parsing 'serviceAck' field")
	}

	// Create a partially initialized instance
	_child := &APDUComplexAck{
		SegmentedMessage:   segmentedMessage,
		MoreFollows:        moreFollows,
		OriginalInvokeId:   originalInvokeId,
		SequenceNumber:     sequenceNumber,
		ProposedWindowSize: proposedWindowSize,
		ServiceAck:         serviceAck,
		Parent:             &APDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *APDUComplexAck) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("APDUComplexAck")

		// Simple Field (segmentedMessage)
		segmentedMessage := bool(m.SegmentedMessage)
		_segmentedMessageErr := io.WriteBit("segmentedMessage", (segmentedMessage))
		if _segmentedMessageErr != nil {
			return errors.Wrap(_segmentedMessageErr, "Error serializing 'segmentedMessage' field")
		}

		// Simple Field (moreFollows)
		moreFollows := bool(m.MoreFollows)
		_moreFollowsErr := io.WriteBit("moreFollows", (moreFollows))
		if _moreFollowsErr != nil {
			return errors.Wrap(_moreFollowsErr, "Error serializing 'moreFollows' field")
		}

		// Reserved Field (reserved)
		{
			_err := io.WriteUint8("reserved", 2, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := io.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Optional Field (sequenceNumber) (Can be skipped, if the value is null)
		var sequenceNumber *uint8 = nil
		if m.SequenceNumber != nil {
			sequenceNumber = m.SequenceNumber
			_sequenceNumberErr := io.WriteUint8("sequenceNumber", 8, *(sequenceNumber))
			if _sequenceNumberErr != nil {
				return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
			}
		}

		// Optional Field (proposedWindowSize) (Can be skipped, if the value is null)
		var proposedWindowSize *uint8 = nil
		if m.ProposedWindowSize != nil {
			proposedWindowSize = m.ProposedWindowSize
			_proposedWindowSizeErr := io.WriteUint8("proposedWindowSize", 8, *(proposedWindowSize))
			if _proposedWindowSizeErr != nil {
				return errors.Wrap(_proposedWindowSizeErr, "Error serializing 'proposedWindowSize' field")
			}
		}

		// Simple Field (serviceAck)
		_serviceAckErr := m.ServiceAck.Serialize(io)
		if _serviceAckErr != nil {
			return errors.Wrap(_serviceAckErr, "Error serializing 'serviceAck' field")
		}

		io.PopContext("APDUComplexAck")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *APDUComplexAck) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "segmentedMessage":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SegmentedMessage = data
			case "moreFollows":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MoreFollows = data
			case "originalInvokeId":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.OriginalInvokeId = data
			case "sequenceNumber":
				// When working with pointers we need to check for an empty element
				var dataString string
				if err := d.DecodeElement(&dataString, &tok); err != nil {
					return err
				}
				if dataString != "" {
					atoi, err := strconv.Atoi(dataString)
					if err != nil {
						return err
					}
					data := uint8(atoi)
					m.SequenceNumber = &data
				}
			case "proposedWindowSize":
				// When working with pointers we need to check for an empty element
				var dataString string
				if err := d.DecodeElement(&dataString, &tok); err != nil {
					return err
				}
				if dataString != "" {
					atoi, err := strconv.Atoi(dataString)
					if err != nil {
						return err
					}
					data := uint8(atoi)
					m.ProposedWindowSize = &data
				}
			case "serviceAck":
				var dt *BACnetServiceAck
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.ServiceAck = dt
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

func (m *APDUComplexAck) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.SegmentedMessage, xml.StartElement{Name: xml.Name{Local: "segmentedMessage"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MoreFollows, xml.StartElement{Name: xml.Name{Local: "moreFollows"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.OriginalInvokeId, xml.StartElement{Name: xml.Name{Local: "originalInvokeId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SequenceNumber, xml.StartElement{Name: xml.Name{Local: "sequenceNumber"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ProposedWindowSize, xml.StartElement{Name: xml.Name{Local: "proposedWindowSize"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ServiceAck, xml.StartElement{Name: xml.Name{Local: "serviceAck"}}); err != nil {
		return err
	}
	return nil
}

func (m APDUComplexAck) String() string {
	return string(m.Box("", 120))
}

func (m APDUComplexAck) Box(name string, width int) utils.AsciiBox {
	boxName := "APDUComplexAck"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// bool can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("SegmentedMessage", m.SegmentedMessage, -1))
		// Simple field (case simple)
		// bool can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("MoreFollows", m.MoreFollows, -1))
		// Reserved Field (reserved)
		// reserved field can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("reserved", uint8(0), -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("OriginalInvokeId", m.OriginalInvokeId, -1))
		// Optional Field (sequenceNumber) (Can be skipped, if the value is null)
		var sequenceNumber *uint8 = nil
		if m.SequenceNumber != nil {
			sequenceNumber = m.SequenceNumber
			// uint8 can be boxed as anything with the least amount of space
			boxes = append(boxes, utils.BoxAnything("SequenceNumber", *(sequenceNumber), -1))
		}
		// Optional Field (proposedWindowSize) (Can be skipped, if the value is null)
		var proposedWindowSize *uint8 = nil
		if m.ProposedWindowSize != nil {
			proposedWindowSize = m.ProposedWindowSize
			// uint8 can be boxed as anything with the least amount of space
			boxes = append(boxes, utils.BoxAnything("ProposedWindowSize", *(proposedWindowSize), -1))
		}
		// Complex field (case complex)
		boxes = append(boxes, m.ServiceAck.Box("serviceAck", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
