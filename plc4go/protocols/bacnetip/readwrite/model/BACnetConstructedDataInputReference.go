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

// BACnetConstructedDataInputReference is the corresponding interface of BACnetConstructedDataInputReference
type BACnetConstructedDataInputReference interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInputReference returns InputReference (property field)
	GetInputReference() BACnetObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectPropertyReference
}

// BACnetConstructedDataInputReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataInputReference.
// This is useful for switch cases.
type BACnetConstructedDataInputReferenceExactly interface {
	BACnetConstructedDataInputReference
	isBACnetConstructedDataInputReference() bool
}

// _BACnetConstructedDataInputReference is the data-structure of this message
type _BACnetConstructedDataInputReference struct {
	*_BACnetConstructedData
	InputReference BACnetObjectPropertyReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInputReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInputReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INPUT_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInputReference) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataInputReference) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInputReference) GetInputReference() BACnetObjectPropertyReference {
	return m.InputReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInputReference) GetActualValue() BACnetObjectPropertyReference {
	return CastBACnetObjectPropertyReference(m.GetInputReference())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInputReference factory function for _BACnetConstructedDataInputReference
func NewBACnetConstructedDataInputReference(inputReference BACnetObjectPropertyReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInputReference {
	_result := &_BACnetConstructedDataInputReference{
		InputReference:         inputReference,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInputReference(structType interface{}) BACnetConstructedDataInputReference {
	if casted, ok := structType.(BACnetConstructedDataInputReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInputReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInputReference) GetTypeName() string {
	return "BACnetConstructedDataInputReference"
}

func (m *_BACnetConstructedDataInputReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataInputReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (inputReference)
	lengthInBits += m.InputReference.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInputReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInputReferenceParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataInputReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInputReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInputReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (inputReference)
	if pullErr := readBuffer.PullContext("inputReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for inputReference")
	}
	_inputReference, _inputReferenceErr := BACnetObjectPropertyReferenceParse(readBuffer)
	if _inputReferenceErr != nil {
		return nil, errors.Wrap(_inputReferenceErr, "Error parsing 'inputReference' field")
	}
	inputReference := _inputReference.(BACnetObjectPropertyReference)
	if closeErr := readBuffer.CloseContext("inputReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for inputReference")
	}

	// Virtual field
	_actualValue := inputReference
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInputReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInputReference")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataInputReference{
		InputReference:         inputReference,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataInputReference) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInputReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInputReference")
		}

		// Simple Field (inputReference)
		if pushErr := writeBuffer.PushContext("inputReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for inputReference")
		}
		_inputReferenceErr := writeBuffer.WriteSerializable(m.GetInputReference())
		if popErr := writeBuffer.PopContext("inputReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for inputReference")
		}
		if _inputReferenceErr != nil {
			return errors.Wrap(_inputReferenceErr, "Error serializing 'inputReference' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInputReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInputReference")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInputReference) isBACnetConstructedDataInputReference() bool {
	return true
}

func (m *_BACnetConstructedDataInputReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
