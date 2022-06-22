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

// BACnetConstructedDataEventAlgorithmInhibitRef is the corresponding interface of BACnetConstructedDataEventAlgorithmInhibitRef
type BACnetConstructedDataEventAlgorithmInhibitRef interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEventAlgorithmInhibitRef returns EventAlgorithmInhibitRef (property field)
	GetEventAlgorithmInhibitRef() BACnetObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectPropertyReference
}

// BACnetConstructedDataEventAlgorithmInhibitRefExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventAlgorithmInhibitRef.
// This is useful for switch cases.
type BACnetConstructedDataEventAlgorithmInhibitRefExactly interface {
	BACnetConstructedDataEventAlgorithmInhibitRef
	isBACnetConstructedDataEventAlgorithmInhibitRef() bool
}

// _BACnetConstructedDataEventAlgorithmInhibitRef is the data-structure of this message
type _BACnetConstructedDataEventAlgorithmInhibitRef struct {
	*_BACnetConstructedData
	EventAlgorithmInhibitRef BACnetObjectPropertyReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_ALGORITHM_INHIBIT_REF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetEventAlgorithmInhibitRef() BACnetObjectPropertyReference {
	return m.EventAlgorithmInhibitRef
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetActualValue() BACnetObjectPropertyReference {
	return CastBACnetObjectPropertyReference(m.GetEventAlgorithmInhibitRef())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventAlgorithmInhibitRef factory function for _BACnetConstructedDataEventAlgorithmInhibitRef
func NewBACnetConstructedDataEventAlgorithmInhibitRef(eventAlgorithmInhibitRef BACnetObjectPropertyReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventAlgorithmInhibitRef {
	_result := &_BACnetConstructedDataEventAlgorithmInhibitRef{
		EventAlgorithmInhibitRef: eventAlgorithmInhibitRef,
		TagNumber:                tagNumber,
		ArrayIndexArgument:       arrayIndexArgument,
		_BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventAlgorithmInhibitRef(structType interface{}) BACnetConstructedDataEventAlgorithmInhibitRef {
	if casted, ok := structType.(BACnetConstructedDataEventAlgorithmInhibitRef); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventAlgorithmInhibitRef); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetTypeName() string {
	return "BACnetConstructedDataEventAlgorithmInhibitRef"
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventAlgorithmInhibitRef)
	lengthInBits += m.EventAlgorithmInhibitRef.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventAlgorithmInhibitRefParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventAlgorithmInhibitRef, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventAlgorithmInhibitRef"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventAlgorithmInhibitRef")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventAlgorithmInhibitRef)
	if pullErr := readBuffer.PullContext("eventAlgorithmInhibitRef"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventAlgorithmInhibitRef")
	}
	_eventAlgorithmInhibitRef, _eventAlgorithmInhibitRefErr := BACnetObjectPropertyReferenceParse(readBuffer)
	if _eventAlgorithmInhibitRefErr != nil {
		return nil, errors.Wrap(_eventAlgorithmInhibitRefErr, "Error parsing 'eventAlgorithmInhibitRef' field")
	}
	eventAlgorithmInhibitRef := _eventAlgorithmInhibitRef.(BACnetObjectPropertyReference)
	if closeErr := readBuffer.CloseContext("eventAlgorithmInhibitRef"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventAlgorithmInhibitRef")
	}

	// Virtual field
	_actualValue := eventAlgorithmInhibitRef
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventAlgorithmInhibitRef"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventAlgorithmInhibitRef")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEventAlgorithmInhibitRef{
		EventAlgorithmInhibitRef: eventAlgorithmInhibitRef,
		_BACnetConstructedData:   &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventAlgorithmInhibitRef"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventAlgorithmInhibitRef")
		}

		// Simple Field (eventAlgorithmInhibitRef)
		if pushErr := writeBuffer.PushContext("eventAlgorithmInhibitRef"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventAlgorithmInhibitRef")
		}
		_eventAlgorithmInhibitRefErr := writeBuffer.WriteSerializable(m.GetEventAlgorithmInhibitRef())
		if popErr := writeBuffer.PopContext("eventAlgorithmInhibitRef"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventAlgorithmInhibitRef")
		}
		if _eventAlgorithmInhibitRefErr != nil {
			return errors.Wrap(_eventAlgorithmInhibitRefErr, "Error serializing 'eventAlgorithmInhibitRef' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventAlgorithmInhibitRef"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventAlgorithmInhibitRef")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) isBACnetConstructedDataEventAlgorithmInhibitRef() bool {
	return true
}

func (m *_BACnetConstructedDataEventAlgorithmInhibitRef) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
