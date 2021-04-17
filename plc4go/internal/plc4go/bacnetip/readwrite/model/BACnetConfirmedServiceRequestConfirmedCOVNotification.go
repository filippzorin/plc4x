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
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const BACnetConfirmedServiceRequestConfirmedCOVNotification_SUBSCRIBERPROCESSIDENTIFIERHEADER uint8 = 0x09
const BACnetConfirmedServiceRequestConfirmedCOVNotification_MONITOREDOBJECTIDENTIFIERHEADER uint8 = 0x1C
const BACnetConfirmedServiceRequestConfirmedCOVNotification_ISSUECONFIRMEDNOTIFICATIONSHEADER uint8 = 0x2C
const BACnetConfirmedServiceRequestConfirmedCOVNotification_LIFETIMEHEADER uint8 = 0x07
const BACnetConfirmedServiceRequestConfirmedCOVNotification_LISTOFVALUESOPENINGTAG uint8 = 0x4E
const BACnetConfirmedServiceRequestConfirmedCOVNotification_LISTOFVALUESCLOSINGTAG uint8 = 0x4F

// The data-structure of this message
type BACnetConfirmedServiceRequestConfirmedCOVNotification struct {
	SubscriberProcessIdentifier               uint8
	MonitoredObjectType                       uint16
	MonitoredObjectInstanceNumber             uint32
	IssueConfirmedNotificationsType           uint16
	IssueConfirmedNotificationsInstanceNumber uint32
	LifetimeLength                            uint8
	LifetimeSeconds                           []int8
	Notifications                             []*BACnetTagWithContent
	Parent                                    *BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestConfirmedCOVNotification interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) ServiceChoice() uint8 {
	return 0x01
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func NewBACnetConfirmedServiceRequestConfirmedCOVNotification(subscriberProcessIdentifier uint8, monitoredObjectType uint16, monitoredObjectInstanceNumber uint32, issueConfirmedNotificationsType uint16, issueConfirmedNotificationsInstanceNumber uint32, lifetimeLength uint8, lifetimeSeconds []int8, notifications []*BACnetTagWithContent) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestConfirmedCOVNotification{
		SubscriberProcessIdentifier:               subscriberProcessIdentifier,
		MonitoredObjectType:                       monitoredObjectType,
		MonitoredObjectInstanceNumber:             monitoredObjectInstanceNumber,
		IssueConfirmedNotificationsType:           issueConfirmedNotificationsType,
		IssueConfirmedNotificationsInstanceNumber: issueConfirmedNotificationsInstanceNumber,
		LifetimeLength:                            lifetimeLength,
		LifetimeSeconds:                           lifetimeSeconds,
		Notifications:                             notifications,
		Parent:                                    NewBACnetConfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetConfirmedServiceRequestConfirmedCOVNotification(structType interface{}) *BACnetConfirmedServiceRequestConfirmedCOVNotification {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestConfirmedCOVNotification {
		if casted, ok := typ.(BACnetConfirmedServiceRequestConfirmedCOVNotification); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestConfirmedCOVNotification); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestConfirmedCOVNotification(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestConfirmedCOVNotification(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedCOVNotification"
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Const Field (subscriberProcessIdentifierHeader)
	lengthInBits += 8

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += 8

	// Const Field (monitoredObjectIdentifierHeader)
	lengthInBits += 8

	// Simple field (monitoredObjectType)
	lengthInBits += 10

	// Simple field (monitoredObjectInstanceNumber)
	lengthInBits += 22

	// Const Field (issueConfirmedNotificationsHeader)
	lengthInBits += 8

	// Simple field (issueConfirmedNotificationsType)
	lengthInBits += 10

	// Simple field (issueConfirmedNotificationsInstanceNumber)
	lengthInBits += 22

	// Const Field (lifetimeHeader)
	lengthInBits += 5

	// Simple field (lifetimeLength)
	lengthInBits += 3

	// Array field
	if len(m.LifetimeSeconds) > 0 {
		lengthInBits += 8 * uint16(len(m.LifetimeSeconds))
	}

	// Const Field (listOfValuesOpeningTag)
	lengthInBits += 8

	// Array field
	if len(m.Notifications) > 0 {
		for _, element := range m.Notifications {
			lengthInBits += element.LengthInBits()
		}
	}

	// Const Field (listOfValuesClosingTag)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedCOVNotificationParse(io *utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {

	// Const Field (subscriberProcessIdentifierHeader)
	subscriberProcessIdentifierHeader, _subscriberProcessIdentifierHeaderErr := io.ReadUint8(8)
	if _subscriberProcessIdentifierHeaderErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierHeaderErr, "Error parsing 'subscriberProcessIdentifierHeader' field")
	}
	if subscriberProcessIdentifierHeader != BACnetConfirmedServiceRequestConfirmedCOVNotification_SUBSCRIBERPROCESSIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestConfirmedCOVNotification_SUBSCRIBERPROCESSIDENTIFIERHEADER) + " but got " + fmt.Sprintf("%d", subscriberProcessIdentifierHeader))
	}

	// Simple Field (subscriberProcessIdentifier)
	subscriberProcessIdentifier, _subscriberProcessIdentifierErr := io.ReadUint8(8)
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field")
	}

	// Const Field (monitoredObjectIdentifierHeader)
	monitoredObjectIdentifierHeader, _monitoredObjectIdentifierHeaderErr := io.ReadUint8(8)
	if _monitoredObjectIdentifierHeaderErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierHeaderErr, "Error parsing 'monitoredObjectIdentifierHeader' field")
	}
	if monitoredObjectIdentifierHeader != BACnetConfirmedServiceRequestConfirmedCOVNotification_MONITOREDOBJECTIDENTIFIERHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestConfirmedCOVNotification_MONITOREDOBJECTIDENTIFIERHEADER) + " but got " + fmt.Sprintf("%d", monitoredObjectIdentifierHeader))
	}

	// Simple Field (monitoredObjectType)
	monitoredObjectType, _monitoredObjectTypeErr := io.ReadUint16(10)
	if _monitoredObjectTypeErr != nil {
		return nil, errors.Wrap(_monitoredObjectTypeErr, "Error parsing 'monitoredObjectType' field")
	}

	// Simple Field (monitoredObjectInstanceNumber)
	monitoredObjectInstanceNumber, _monitoredObjectInstanceNumberErr := io.ReadUint32(22)
	if _monitoredObjectInstanceNumberErr != nil {
		return nil, errors.Wrap(_monitoredObjectInstanceNumberErr, "Error parsing 'monitoredObjectInstanceNumber' field")
	}

	// Const Field (issueConfirmedNotificationsHeader)
	issueConfirmedNotificationsHeader, _issueConfirmedNotificationsHeaderErr := io.ReadUint8(8)
	if _issueConfirmedNotificationsHeaderErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsHeaderErr, "Error parsing 'issueConfirmedNotificationsHeader' field")
	}
	if issueConfirmedNotificationsHeader != BACnetConfirmedServiceRequestConfirmedCOVNotification_ISSUECONFIRMEDNOTIFICATIONSHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestConfirmedCOVNotification_ISSUECONFIRMEDNOTIFICATIONSHEADER) + " but got " + fmt.Sprintf("%d", issueConfirmedNotificationsHeader))
	}

	// Simple Field (issueConfirmedNotificationsType)
	issueConfirmedNotificationsType, _issueConfirmedNotificationsTypeErr := io.ReadUint16(10)
	if _issueConfirmedNotificationsTypeErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsTypeErr, "Error parsing 'issueConfirmedNotificationsType' field")
	}

	// Simple Field (issueConfirmedNotificationsInstanceNumber)
	issueConfirmedNotificationsInstanceNumber, _issueConfirmedNotificationsInstanceNumberErr := io.ReadUint32(22)
	if _issueConfirmedNotificationsInstanceNumberErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsInstanceNumberErr, "Error parsing 'issueConfirmedNotificationsInstanceNumber' field")
	}

	// Const Field (lifetimeHeader)
	lifetimeHeader, _lifetimeHeaderErr := io.ReadUint8(5)
	if _lifetimeHeaderErr != nil {
		return nil, errors.Wrap(_lifetimeHeaderErr, "Error parsing 'lifetimeHeader' field")
	}
	if lifetimeHeader != BACnetConfirmedServiceRequestConfirmedCOVNotification_LIFETIMEHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestConfirmedCOVNotification_LIFETIMEHEADER) + " but got " + fmt.Sprintf("%d", lifetimeHeader))
	}

	// Simple Field (lifetimeLength)
	lifetimeLength, _lifetimeLengthErr := io.ReadUint8(3)
	if _lifetimeLengthErr != nil {
		return nil, errors.Wrap(_lifetimeLengthErr, "Error parsing 'lifetimeLength' field")
	}

	// Array field (lifetimeSeconds)
	// Count array
	lifetimeSeconds := make([]int8, lifetimeLength)
	for curItem := uint16(0); curItem < uint16(lifetimeLength); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'lifetimeSeconds' field")
		}
		lifetimeSeconds[curItem] = _item
	}

	// Const Field (listOfValuesOpeningTag)
	listOfValuesOpeningTag, _listOfValuesOpeningTagErr := io.ReadUint8(8)
	if _listOfValuesOpeningTagErr != nil {
		return nil, errors.Wrap(_listOfValuesOpeningTagErr, "Error parsing 'listOfValuesOpeningTag' field")
	}
	if listOfValuesOpeningTag != BACnetConfirmedServiceRequestConfirmedCOVNotification_LISTOFVALUESOPENINGTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestConfirmedCOVNotification_LISTOFVALUESOPENINGTAG) + " but got " + fmt.Sprintf("%d", listOfValuesOpeningTag))
	}

	// Array field (notifications)
	// Length array
	notifications := make([]*BACnetTagWithContent, 0)
	_notificationsLength := uint16(len) - uint16(uint16(18))
	_notificationsEndPos := io.GetPos() + uint16(_notificationsLength)
	for io.GetPos() < _notificationsEndPos {
		_item, _err := BACnetTagWithContentParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'notifications' field")
		}
		notifications = append(notifications, _item)
	}

	// Const Field (listOfValuesClosingTag)
	listOfValuesClosingTag, _listOfValuesClosingTagErr := io.ReadUint8(8)
	if _listOfValuesClosingTagErr != nil {
		return nil, errors.Wrap(_listOfValuesClosingTagErr, "Error parsing 'listOfValuesClosingTag' field")
	}
	if listOfValuesClosingTag != BACnetConfirmedServiceRequestConfirmedCOVNotification_LISTOFVALUESCLOSINGTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetConfirmedServiceRequestConfirmedCOVNotification_LISTOFVALUESCLOSINGTAG) + " but got " + fmt.Sprintf("%d", listOfValuesClosingTag))
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestConfirmedCOVNotification{
		SubscriberProcessIdentifier:               subscriberProcessIdentifier,
		MonitoredObjectType:                       monitoredObjectType,
		MonitoredObjectInstanceNumber:             monitoredObjectInstanceNumber,
		IssueConfirmedNotificationsType:           issueConfirmedNotificationsType,
		IssueConfirmedNotificationsInstanceNumber: issueConfirmedNotificationsInstanceNumber,
		LifetimeLength:                            lifetimeLength,
		LifetimeSeconds:                           lifetimeSeconds,
		Notifications:                             notifications,
		Parent:                                    &BACnetConfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Const Field (subscriberProcessIdentifierHeader)
		_subscriberProcessIdentifierHeaderErr := io.WriteUint8(8, 0x09)
		if _subscriberProcessIdentifierHeaderErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierHeaderErr, "Error serializing 'subscriberProcessIdentifierHeader' field")
		}

		// Simple Field (subscriberProcessIdentifier)
		subscriberProcessIdentifier := uint8(m.SubscriberProcessIdentifier)
		_subscriberProcessIdentifierErr := io.WriteUint8(8, (subscriberProcessIdentifier))
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Const Field (monitoredObjectIdentifierHeader)
		_monitoredObjectIdentifierHeaderErr := io.WriteUint8(8, 0x1C)
		if _monitoredObjectIdentifierHeaderErr != nil {
			return errors.Wrap(_monitoredObjectIdentifierHeaderErr, "Error serializing 'monitoredObjectIdentifierHeader' field")
		}

		// Simple Field (monitoredObjectType)
		monitoredObjectType := uint16(m.MonitoredObjectType)
		_monitoredObjectTypeErr := io.WriteUint16(10, (monitoredObjectType))
		if _monitoredObjectTypeErr != nil {
			return errors.Wrap(_monitoredObjectTypeErr, "Error serializing 'monitoredObjectType' field")
		}

		// Simple Field (monitoredObjectInstanceNumber)
		monitoredObjectInstanceNumber := uint32(m.MonitoredObjectInstanceNumber)
		_monitoredObjectInstanceNumberErr := io.WriteUint32(22, (monitoredObjectInstanceNumber))
		if _monitoredObjectInstanceNumberErr != nil {
			return errors.Wrap(_monitoredObjectInstanceNumberErr, "Error serializing 'monitoredObjectInstanceNumber' field")
		}

		// Const Field (issueConfirmedNotificationsHeader)
		_issueConfirmedNotificationsHeaderErr := io.WriteUint8(8, 0x2C)
		if _issueConfirmedNotificationsHeaderErr != nil {
			return errors.Wrap(_issueConfirmedNotificationsHeaderErr, "Error serializing 'issueConfirmedNotificationsHeader' field")
		}

		// Simple Field (issueConfirmedNotificationsType)
		issueConfirmedNotificationsType := uint16(m.IssueConfirmedNotificationsType)
		_issueConfirmedNotificationsTypeErr := io.WriteUint16(10, (issueConfirmedNotificationsType))
		if _issueConfirmedNotificationsTypeErr != nil {
			return errors.Wrap(_issueConfirmedNotificationsTypeErr, "Error serializing 'issueConfirmedNotificationsType' field")
		}

		// Simple Field (issueConfirmedNotificationsInstanceNumber)
		issueConfirmedNotificationsInstanceNumber := uint32(m.IssueConfirmedNotificationsInstanceNumber)
		_issueConfirmedNotificationsInstanceNumberErr := io.WriteUint32(22, (issueConfirmedNotificationsInstanceNumber))
		if _issueConfirmedNotificationsInstanceNumberErr != nil {
			return errors.Wrap(_issueConfirmedNotificationsInstanceNumberErr, "Error serializing 'issueConfirmedNotificationsInstanceNumber' field")
		}

		// Const Field (lifetimeHeader)
		_lifetimeHeaderErr := io.WriteUint8(5, 0x07)
		if _lifetimeHeaderErr != nil {
			return errors.Wrap(_lifetimeHeaderErr, "Error serializing 'lifetimeHeader' field")
		}

		// Simple Field (lifetimeLength)
		lifetimeLength := uint8(m.LifetimeLength)
		_lifetimeLengthErr := io.WriteUint8(3, (lifetimeLength))
		if _lifetimeLengthErr != nil {
			return errors.Wrap(_lifetimeLengthErr, "Error serializing 'lifetimeLength' field")
		}

		// Array Field (lifetimeSeconds)
		if m.LifetimeSeconds != nil {
			for _, _element := range m.LifetimeSeconds {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'lifetimeSeconds' field")
				}
			}
		}

		// Const Field (listOfValuesOpeningTag)
		_listOfValuesOpeningTagErr := io.WriteUint8(8, 0x4E)
		if _listOfValuesOpeningTagErr != nil {
			return errors.Wrap(_listOfValuesOpeningTagErr, "Error serializing 'listOfValuesOpeningTag' field")
		}

		// Array Field (notifications)
		if m.Notifications != nil {
			for _, _element := range m.Notifications {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'notifications' field")
				}
			}
		}

		// Const Field (listOfValuesClosingTag)
		_listOfValuesClosingTagErr := io.WriteUint8(8, 0x4F)
		if _listOfValuesClosingTagErr != nil {
			return errors.Wrap(_listOfValuesClosingTagErr, "Error serializing 'listOfValuesClosingTag' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "subscriberProcessIdentifier":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SubscriberProcessIdentifier = data
			case "monitoredObjectType":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MonitoredObjectType = data
			case "monitoredObjectInstanceNumber":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MonitoredObjectInstanceNumber = data
			case "issueConfirmedNotificationsType":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.IssueConfirmedNotificationsType = data
			case "issueConfirmedNotificationsInstanceNumber":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.IssueConfirmedNotificationsInstanceNumber = data
			case "lifetimeLength":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.LifetimeLength = data
			case "lifetimeSeconds":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.LifetimeSeconds = utils.ByteArrayToInt8Array(_decoded[0:_len])
			case "notifications":
				var data []*BACnetTagWithContent
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Notifications = data
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

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.SubscriberProcessIdentifier, xml.StartElement{Name: xml.Name{Local: "subscriberProcessIdentifier"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MonitoredObjectType, xml.StartElement{Name: xml.Name{Local: "monitoredObjectType"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MonitoredObjectInstanceNumber, xml.StartElement{Name: xml.Name{Local: "monitoredObjectInstanceNumber"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.IssueConfirmedNotificationsType, xml.StartElement{Name: xml.Name{Local: "issueConfirmedNotificationsType"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.IssueConfirmedNotificationsInstanceNumber, xml.StartElement{Name: xml.Name{Local: "issueConfirmedNotificationsInstanceNumber"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.LifetimeLength, xml.StartElement{Name: xml.Name{Local: "lifetimeLength"}}); err != nil {
		return err
	}
	_encodedLifetimeSeconds := hex.EncodeToString(utils.Int8ArrayToByteArray(m.LifetimeSeconds))
	_encodedLifetimeSeconds = strings.ToUpper(_encodedLifetimeSeconds)
	if err := e.EncodeElement(_encodedLifetimeSeconds, xml.StartElement{Name: xml.Name{Local: "lifetimeSeconds"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "notifications"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.Notifications {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "notifications"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "notifications"}}); err != nil {
		return err
	}
	return nil
}

func (m BACnetConfirmedServiceRequestConfirmedCOVNotification) String() string {
	return string(m.Box("BACnetConfirmedServiceRequestConfirmedCOVNotification", utils.DefaultWidth*2))
}

func (m BACnetConfirmedServiceRequestConfirmedCOVNotification) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "BACnetConfirmedServiceRequestConfirmedCOVNotification"
	}
	boxChild := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		boxes = append(boxes, utils.BoxAnything("SubscriberProcessIdentifier", m.SubscriberProcessIdentifier, width-2))
		boxes = append(boxes, utils.BoxAnything("MonitoredObjectType", m.MonitoredObjectType, width-2))
		boxes = append(boxes, utils.BoxAnything("MonitoredObjectInstanceNumber", m.MonitoredObjectInstanceNumber, width-2))
		boxes = append(boxes, utils.BoxAnything("IssueConfirmedNotificationsType", m.IssueConfirmedNotificationsType, width-2))
		boxes = append(boxes, utils.BoxAnything("IssueConfirmedNotificationsInstanceNumber", m.IssueConfirmedNotificationsInstanceNumber, width-2))
		boxes = append(boxes, utils.BoxAnything("LifetimeLength", m.LifetimeLength, width-2))
		boxes = append(boxes, utils.BoxAnything("LifetimeSeconds", m.LifetimeSeconds, width-2))
		boxes = append(boxes, utils.BoxAnything("Notifications", m.Notifications, width-2))
		return boxes
	}
	return m.Parent.BoxParent(name, width, boxChild)
}
