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

// BACnetConstructedDataInstanceOf is the corresponding interface of BACnetConstructedDataInstanceOf
type BACnetConstructedDataInstanceOf interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInstanceOf returns InstanceOf (property field)
	GetInstanceOf() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataInstanceOfExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataInstanceOf.
// This is useful for switch cases.
type BACnetConstructedDataInstanceOfExactly interface {
	BACnetConstructedDataInstanceOf
	isBACnetConstructedDataInstanceOf() bool
}

// _BACnetConstructedDataInstanceOf is the data-structure of this message
type _BACnetConstructedDataInstanceOf struct {
	*_BACnetConstructedData
	InstanceOf BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInstanceOf) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInstanceOf) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INSTANCE_OF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInstanceOf) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataInstanceOf) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInstanceOf) GetInstanceOf() BACnetApplicationTagCharacterString {
	return m.InstanceOf
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInstanceOf) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetInstanceOf())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInstanceOf factory function for _BACnetConstructedDataInstanceOf
func NewBACnetConstructedDataInstanceOf(instanceOf BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInstanceOf {
	_result := &_BACnetConstructedDataInstanceOf{
		InstanceOf:             instanceOf,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInstanceOf(structType interface{}) BACnetConstructedDataInstanceOf {
	if casted, ok := structType.(BACnetConstructedDataInstanceOf); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInstanceOf); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInstanceOf) GetTypeName() string {
	return "BACnetConstructedDataInstanceOf"
}

func (m *_BACnetConstructedDataInstanceOf) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataInstanceOf) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (instanceOf)
	lengthInBits += m.InstanceOf.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInstanceOf) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInstanceOfParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataInstanceOf, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInstanceOf"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInstanceOf")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (instanceOf)
	if pullErr := readBuffer.PullContext("instanceOf"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for instanceOf")
	}
	_instanceOf, _instanceOfErr := BACnetApplicationTagParse(readBuffer)
	if _instanceOfErr != nil {
		return nil, errors.Wrap(_instanceOfErr, "Error parsing 'instanceOf' field")
	}
	instanceOf := _instanceOf.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("instanceOf"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for instanceOf")
	}

	// Virtual field
	_actualValue := instanceOf
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInstanceOf"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInstanceOf")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataInstanceOf{
		InstanceOf:             instanceOf,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataInstanceOf) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInstanceOf"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInstanceOf")
		}

		// Simple Field (instanceOf)
		if pushErr := writeBuffer.PushContext("instanceOf"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for instanceOf")
		}
		_instanceOfErr := writeBuffer.WriteSerializable(m.GetInstanceOf())
		if popErr := writeBuffer.PopContext("instanceOf"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for instanceOf")
		}
		if _instanceOfErr != nil {
			return errors.Wrap(_instanceOfErr, "Error serializing 'instanceOf' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInstanceOf"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInstanceOf")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInstanceOf) isBACnetConstructedDataInstanceOf() bool {
	return true
}

func (m *_BACnetConstructedDataInstanceOf) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
