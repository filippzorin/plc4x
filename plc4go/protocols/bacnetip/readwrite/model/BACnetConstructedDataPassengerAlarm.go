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

// BACnetConstructedDataPassengerAlarm is the corresponding interface of BACnetConstructedDataPassengerAlarm
type BACnetConstructedDataPassengerAlarm interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPassengerAlarm returns PassengerAlarm (property field)
	GetPassengerAlarm() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataPassengerAlarmExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPassengerAlarm.
// This is useful for switch cases.
type BACnetConstructedDataPassengerAlarmExactly interface {
	BACnetConstructedDataPassengerAlarm
	isBACnetConstructedDataPassengerAlarm() bool
}

// _BACnetConstructedDataPassengerAlarm is the data-structure of this message
type _BACnetConstructedDataPassengerAlarm struct {
	*_BACnetConstructedData
	PassengerAlarm BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPassengerAlarm) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPassengerAlarm) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PASSENGER_ALARM
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPassengerAlarm) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPassengerAlarm) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPassengerAlarm) GetPassengerAlarm() BACnetApplicationTagBoolean {
	return m.PassengerAlarm
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPassengerAlarm) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetPassengerAlarm())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPassengerAlarm factory function for _BACnetConstructedDataPassengerAlarm
func NewBACnetConstructedDataPassengerAlarm(passengerAlarm BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPassengerAlarm {
	_result := &_BACnetConstructedDataPassengerAlarm{
		PassengerAlarm:         passengerAlarm,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPassengerAlarm(structType interface{}) BACnetConstructedDataPassengerAlarm {
	if casted, ok := structType.(BACnetConstructedDataPassengerAlarm); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPassengerAlarm); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPassengerAlarm) GetTypeName() string {
	return "BACnetConstructedDataPassengerAlarm"
}

func (m *_BACnetConstructedDataPassengerAlarm) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataPassengerAlarm) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (passengerAlarm)
	lengthInBits += m.PassengerAlarm.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPassengerAlarm) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPassengerAlarmParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPassengerAlarm, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPassengerAlarm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPassengerAlarm")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (passengerAlarm)
	if pullErr := readBuffer.PullContext("passengerAlarm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for passengerAlarm")
	}
	_passengerAlarm, _passengerAlarmErr := BACnetApplicationTagParse(readBuffer)
	if _passengerAlarmErr != nil {
		return nil, errors.Wrap(_passengerAlarmErr, "Error parsing 'passengerAlarm' field")
	}
	passengerAlarm := _passengerAlarm.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("passengerAlarm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for passengerAlarm")
	}

	// Virtual field
	_actualValue := passengerAlarm
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPassengerAlarm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPassengerAlarm")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPassengerAlarm{
		PassengerAlarm:         passengerAlarm,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPassengerAlarm) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPassengerAlarm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPassengerAlarm")
		}

		// Simple Field (passengerAlarm)
		if pushErr := writeBuffer.PushContext("passengerAlarm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for passengerAlarm")
		}
		_passengerAlarmErr := writeBuffer.WriteSerializable(m.GetPassengerAlarm())
		if popErr := writeBuffer.PopContext("passengerAlarm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for passengerAlarm")
		}
		if _passengerAlarmErr != nil {
			return errors.Wrap(_passengerAlarmErr, "Error serializing 'passengerAlarm' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPassengerAlarm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPassengerAlarm")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPassengerAlarm) isBACnetConstructedDataPassengerAlarm() bool {
	return true
}

func (m *_BACnetConstructedDataPassengerAlarm) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
