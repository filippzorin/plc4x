/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetContextTagEnumerated is the corresponding interface of BACnetContextTagEnumerated
type BACnetContextTagEnumerated interface {
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadEnumerated
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint32
}

// BACnetContextTagEnumeratedExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagEnumerated.
// This is useful for switch cases.
type BACnetContextTagEnumeratedExactly interface {
	BACnetContextTagEnumerated
	isBACnetContextTagEnumerated() bool
}

// _BACnetContextTagEnumerated is the data-structure of this message
type _BACnetContextTagEnumerated struct {
	*_BACnetContextTag
	Payload BACnetTagPayloadEnumerated

	// Arguments.
	TagNumberArgument uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagEnumerated) GetDataType() BACnetDataType {
	return BACnetDataType_ENUMERATED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagEnumerated) InitializeParent(parent BACnetContextTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetContextTagEnumerated) GetParent() BACnetContextTag {
	return m._BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagEnumerated) GetPayload() BACnetTagPayloadEnumerated {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTagEnumerated) GetActualValue() uint32 {
	return uint32(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagEnumerated factory function for _BACnetContextTagEnumerated
func NewBACnetContextTagEnumerated(payload BACnetTagPayloadEnumerated, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagEnumerated {
	_result := &_BACnetContextTagEnumerated{
		Payload:           payload,
		TagNumberArgument: tagNumberArgument,
		_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagEnumerated(structType interface{}) BACnetContextTagEnumerated {
	if casted, ok := structType.(BACnetContextTagEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagEnumerated) GetTypeName() string {
	return "BACnetContextTagEnumerated"
}

func (m *_BACnetContextTagEnumerated) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetContextTagEnumerated) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagEnumerated) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagEnumeratedParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, header BACnetTagHeader) (BACnetContextTagEnumerated, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadEnumeratedParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := _payload.(BACnetTagPayloadEnumerated)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_actualValue := payload.GetActualValue()
	actualValue := uint32(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagEnumerated")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagEnumerated{
		Payload:           payload,
		_BACnetContextTag: &_BACnetContextTag{},
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagEnumerated) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagEnumerated")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagEnumerated")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetContextTagEnumerated) isBACnetContextTagEnumerated() bool {
	return true
}

func (m *_BACnetContextTagEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
