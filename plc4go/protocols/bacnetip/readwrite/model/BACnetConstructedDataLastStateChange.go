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

// BACnetConstructedDataLastStateChange is the corresponding interface of BACnetConstructedDataLastStateChange
type BACnetConstructedDataLastStateChange interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastStateChange returns LastStateChange (property field)
	GetLastStateChange() BACnetTimerTransitionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimerTransitionTagged
}

// BACnetConstructedDataLastStateChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastStateChange.
// This is useful for switch cases.
type BACnetConstructedDataLastStateChangeExactly interface {
	BACnetConstructedDataLastStateChange
	isBACnetConstructedDataLastStateChange() bool
}

// _BACnetConstructedDataLastStateChange is the data-structure of this message
type _BACnetConstructedDataLastStateChange struct {
	*_BACnetConstructedData
	LastStateChange BACnetTimerTransitionTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastStateChange) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_STATE_CHANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastStateChange) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLastStateChange) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetLastStateChange() BACnetTimerTransitionTagged {
	return m.LastStateChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastStateChange) GetActualValue() BACnetTimerTransitionTagged {
	return CastBACnetTimerTransitionTagged(m.GetLastStateChange())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastStateChange factory function for _BACnetConstructedDataLastStateChange
func NewBACnetConstructedDataLastStateChange(lastStateChange BACnetTimerTransitionTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastStateChange {
	_result := &_BACnetConstructedDataLastStateChange{
		LastStateChange:        lastStateChange,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastStateChange(structType interface{}) BACnetConstructedDataLastStateChange {
	if casted, ok := structType.(BACnetConstructedDataLastStateChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastStateChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastStateChange) GetTypeName() string {
	return "BACnetConstructedDataLastStateChange"
}

func (m *_BACnetConstructedDataLastStateChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLastStateChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lastStateChange)
	lengthInBits += m.LastStateChange.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastStateChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLastStateChangeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastStateChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastStateChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastStateChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastStateChange)
	if pullErr := readBuffer.PullContext("lastStateChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lastStateChange")
	}
	_lastStateChange, _lastStateChangeErr := BACnetTimerTransitionTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _lastStateChangeErr != nil {
		return nil, errors.Wrap(_lastStateChangeErr, "Error parsing 'lastStateChange' field")
	}
	lastStateChange := _lastStateChange.(BACnetTimerTransitionTagged)
	if closeErr := readBuffer.CloseContext("lastStateChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lastStateChange")
	}

	// Virtual field
	_actualValue := lastStateChange
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastStateChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastStateChange")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLastStateChange{
		LastStateChange:        lastStateChange,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLastStateChange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastStateChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastStateChange")
		}

		// Simple Field (lastStateChange)
		if pushErr := writeBuffer.PushContext("lastStateChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lastStateChange")
		}
		_lastStateChangeErr := writeBuffer.WriteSerializable(m.GetLastStateChange())
		if popErr := writeBuffer.PopContext("lastStateChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lastStateChange")
		}
		if _lastStateChangeErr != nil {
			return errors.Wrap(_lastStateChangeErr, "Error serializing 'lastStateChange' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastStateChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastStateChange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastStateChange) isBACnetConstructedDataLastStateChange() bool {
	return true
}

func (m *_BACnetConstructedDataLastStateChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
