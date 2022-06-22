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

// BACnetConstructedDataDefaultStepIncrement is the corresponding interface of BACnetConstructedDataDefaultStepIncrement
type BACnetConstructedDataDefaultStepIncrement interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDefaultStepIncrement returns DefaultStepIncrement (property field)
	GetDefaultStepIncrement() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataDefaultStepIncrementExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDefaultStepIncrement.
// This is useful for switch cases.
type BACnetConstructedDataDefaultStepIncrementExactly interface {
	BACnetConstructedDataDefaultStepIncrement
	isBACnetConstructedDataDefaultStepIncrement() bool
}

// _BACnetConstructedDataDefaultStepIncrement is the data-structure of this message
type _BACnetConstructedDataDefaultStepIncrement struct {
	*_BACnetConstructedData
	DefaultStepIncrement BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDefaultStepIncrement) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDefaultStepIncrement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEFAULT_STEP_INCREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDefaultStepIncrement) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDefaultStepIncrement) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultStepIncrement) GetDefaultStepIncrement() BACnetApplicationTagReal {
	return m.DefaultStepIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultStepIncrement) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetDefaultStepIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDefaultStepIncrement factory function for _BACnetConstructedDataDefaultStepIncrement
func NewBACnetConstructedDataDefaultStepIncrement(defaultStepIncrement BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDefaultStepIncrement {
	_result := &_BACnetConstructedDataDefaultStepIncrement{
		DefaultStepIncrement:   defaultStepIncrement,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDefaultStepIncrement(structType interface{}) BACnetConstructedDataDefaultStepIncrement {
	if casted, ok := structType.(BACnetConstructedDataDefaultStepIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDefaultStepIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDefaultStepIncrement) GetTypeName() string {
	return "BACnetConstructedDataDefaultStepIncrement"
}

func (m *_BACnetConstructedDataDefaultStepIncrement) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDefaultStepIncrement) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (defaultStepIncrement)
	lengthInBits += m.DefaultStepIncrement.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDefaultStepIncrement) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDefaultStepIncrementParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDefaultStepIncrement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDefaultStepIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDefaultStepIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (defaultStepIncrement)
	if pullErr := readBuffer.PullContext("defaultStepIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for defaultStepIncrement")
	}
	_defaultStepIncrement, _defaultStepIncrementErr := BACnetApplicationTagParse(readBuffer)
	if _defaultStepIncrementErr != nil {
		return nil, errors.Wrap(_defaultStepIncrementErr, "Error parsing 'defaultStepIncrement' field")
	}
	defaultStepIncrement := _defaultStepIncrement.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("defaultStepIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for defaultStepIncrement")
	}

	// Virtual field
	_actualValue := defaultStepIncrement
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDefaultStepIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDefaultStepIncrement")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDefaultStepIncrement{
		DefaultStepIncrement:   defaultStepIncrement,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDefaultStepIncrement) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDefaultStepIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDefaultStepIncrement")
		}

		// Simple Field (defaultStepIncrement)
		if pushErr := writeBuffer.PushContext("defaultStepIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for defaultStepIncrement")
		}
		_defaultStepIncrementErr := writeBuffer.WriteSerializable(m.GetDefaultStepIncrement())
		if popErr := writeBuffer.PopContext("defaultStepIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for defaultStepIncrement")
		}
		if _defaultStepIncrementErr != nil {
			return errors.Wrap(_defaultStepIncrementErr, "Error serializing 'defaultStepIncrement' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDefaultStepIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDefaultStepIncrement")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDefaultStepIncrement) isBACnetConstructedDataDefaultStepIncrement() bool {
	return true
}

func (m *_BACnetConstructedDataDefaultStepIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
