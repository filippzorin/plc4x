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

// BACnetConstructedDataOperationDirection is the corresponding interface of BACnetConstructedDataOperationDirection
type BACnetConstructedDataOperationDirection interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetOperationDirection returns OperationDirection (property field)
	GetOperationDirection() BACnetEscalatorOperationDirectionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEscalatorOperationDirectionTagged
}

// BACnetConstructedDataOperationDirectionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOperationDirection.
// This is useful for switch cases.
type BACnetConstructedDataOperationDirectionExactly interface {
	BACnetConstructedDataOperationDirection
	isBACnetConstructedDataOperationDirection() bool
}

// _BACnetConstructedDataOperationDirection is the data-structure of this message
type _BACnetConstructedDataOperationDirection struct {
	*_BACnetConstructedData
	OperationDirection BACnetEscalatorOperationDirectionTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOperationDirection) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OPERATION_DIRECTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOperationDirection) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataOperationDirection) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetOperationDirection() BACnetEscalatorOperationDirectionTagged {
	return m.OperationDirection
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetActualValue() BACnetEscalatorOperationDirectionTagged {
	return CastBACnetEscalatorOperationDirectionTagged(m.GetOperationDirection())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOperationDirection factory function for _BACnetConstructedDataOperationDirection
func NewBACnetConstructedDataOperationDirection(operationDirection BACnetEscalatorOperationDirectionTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOperationDirection {
	_result := &_BACnetConstructedDataOperationDirection{
		OperationDirection:     operationDirection,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOperationDirection(structType interface{}) BACnetConstructedDataOperationDirection {
	if casted, ok := structType.(BACnetConstructedDataOperationDirection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOperationDirection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOperationDirection) GetTypeName() string {
	return "BACnetConstructedDataOperationDirection"
}

func (m *_BACnetConstructedDataOperationDirection) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataOperationDirection) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (operationDirection)
	lengthInBits += m.OperationDirection.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOperationDirection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOperationDirectionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOperationDirection, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOperationDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOperationDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (operationDirection)
	if pullErr := readBuffer.PullContext("operationDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for operationDirection")
	}
	_operationDirection, _operationDirectionErr := BACnetEscalatorOperationDirectionTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _operationDirectionErr != nil {
		return nil, errors.Wrap(_operationDirectionErr, "Error parsing 'operationDirection' field")
	}
	operationDirection := _operationDirection.(BACnetEscalatorOperationDirectionTagged)
	if closeErr := readBuffer.CloseContext("operationDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for operationDirection")
	}

	// Virtual field
	_actualValue := operationDirection
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOperationDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOperationDirection")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataOperationDirection{
		OperationDirection:     operationDirection,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataOperationDirection) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOperationDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOperationDirection")
		}

		// Simple Field (operationDirection)
		if pushErr := writeBuffer.PushContext("operationDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for operationDirection")
		}
		_operationDirectionErr := writeBuffer.WriteSerializable(m.GetOperationDirection())
		if popErr := writeBuffer.PopContext("operationDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for operationDirection")
		}
		if _operationDirectionErr != nil {
			return errors.Wrap(_operationDirectionErr, "Error serializing 'operationDirection' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOperationDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOperationDirection")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOperationDirection) isBACnetConstructedDataOperationDirection() bool {
	return true
}

func (m *_BACnetConstructedDataOperationDirection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
