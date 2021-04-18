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
type AdsDeleteDeviceNotificationResponse struct {
	Result ReturnCode
	Parent *AdsData
}

// The corresponding interface
type IAdsDeleteDeviceNotificationResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsDeleteDeviceNotificationResponse) CommandId() CommandId {
	return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
}

func (m *AdsDeleteDeviceNotificationResponse) Response() bool {
	return true
}

func (m *AdsDeleteDeviceNotificationResponse) InitializeParent(parent *AdsData) {
}

func NewAdsDeleteDeviceNotificationResponse(result ReturnCode) *AdsData {
	child := &AdsDeleteDeviceNotificationResponse{
		Result: result,
		Parent: NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsDeleteDeviceNotificationResponse(structType interface{}) *AdsDeleteDeviceNotificationResponse {
	castFunc := func(typ interface{}) *AdsDeleteDeviceNotificationResponse {
		if casted, ok := typ.(AdsDeleteDeviceNotificationResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsDeleteDeviceNotificationResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsDeleteDeviceNotificationResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsDeleteDeviceNotificationResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsDeleteDeviceNotificationResponse) GetTypeName() string {
	return "AdsDeleteDeviceNotificationResponse"
}

func (m *AdsDeleteDeviceNotificationResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsDeleteDeviceNotificationResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsDeleteDeviceNotificationResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsDeleteDeviceNotificationResponseParse(io utils.ReadBuffer) (*AdsData, error) {

	// Simple Field (result)
	result, _resultErr := ReturnCodeParse(io)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}

	// Create a partially initialized instance
	_child := &AdsDeleteDeviceNotificationResponse{
		Result: result,
		Parent: &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsDeleteDeviceNotificationResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("AdsDeleteDeviceNotificationResponse")

		// Simple Field (result)
		_resultErr := m.Result.Serialize(io)
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		io.PopContext("AdsDeleteDeviceNotificationResponse")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsDeleteDeviceNotificationResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "result":
				var data ReturnCode
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Result = data
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

func (m *AdsDeleteDeviceNotificationResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Result, xml.StartElement{Name: xml.Name{Local: "result"}}); err != nil {
		return err
	}
	return nil
}

func (m AdsDeleteDeviceNotificationResponse) String() string {
	return string(m.Box("", 120))
}

func (m AdsDeleteDeviceNotificationResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "AdsDeleteDeviceNotificationResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Complex field (case complex)
		boxes = append(boxes, m.Result.Box("result", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
