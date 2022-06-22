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

// BACnetConstructedDataBinaryOutputFeedbackValue is the corresponding interface of BACnetConstructedDataBinaryOutputFeedbackValue
type BACnetConstructedDataBinaryOutputFeedbackValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFeedbackValue returns FeedbackValue (property field)
	GetFeedbackValue() BACnetBinaryPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryPVTagged
}

// BACnetConstructedDataBinaryOutputFeedbackValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBinaryOutputFeedbackValue.
// This is useful for switch cases.
type BACnetConstructedDataBinaryOutputFeedbackValueExactly interface {
	BACnetConstructedDataBinaryOutputFeedbackValue
	isBACnetConstructedDataBinaryOutputFeedbackValue() bool
}

// _BACnetConstructedDataBinaryOutputFeedbackValue is the data-structure of this message
type _BACnetConstructedDataBinaryOutputFeedbackValue struct {
	*_BACnetConstructedData
	FeedbackValue BACnetBinaryPVTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_OUTPUT
}

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FEEDBACK_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) GetFeedbackValue() BACnetBinaryPVTagged {
	return m.FeedbackValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) GetActualValue() BACnetBinaryPVTagged {
	return CastBACnetBinaryPVTagged(m.GetFeedbackValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBinaryOutputFeedbackValue factory function for _BACnetConstructedDataBinaryOutputFeedbackValue
func NewBACnetConstructedDataBinaryOutputFeedbackValue(feedbackValue BACnetBinaryPVTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryOutputFeedbackValue {
	_result := &_BACnetConstructedDataBinaryOutputFeedbackValue{
		FeedbackValue:          feedbackValue,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryOutputFeedbackValue(structType interface{}) BACnetConstructedDataBinaryOutputFeedbackValue {
	if casted, ok := structType.(BACnetConstructedDataBinaryOutputFeedbackValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryOutputFeedbackValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryOutputFeedbackValue"
}

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (feedbackValue)
	lengthInBits += m.FeedbackValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBinaryOutputFeedbackValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBinaryOutputFeedbackValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryOutputFeedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryOutputFeedbackValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (feedbackValue)
	if pullErr := readBuffer.PullContext("feedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for feedbackValue")
	}
	_feedbackValue, _feedbackValueErr := BACnetBinaryPVTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _feedbackValueErr != nil {
		return nil, errors.Wrap(_feedbackValueErr, "Error parsing 'feedbackValue' field")
	}
	feedbackValue := _feedbackValue.(BACnetBinaryPVTagged)
	if closeErr := readBuffer.CloseContext("feedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for feedbackValue")
	}

	// Virtual field
	_actualValue := feedbackValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryOutputFeedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryOutputFeedbackValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBinaryOutputFeedbackValue{
		FeedbackValue:          feedbackValue,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryOutputFeedbackValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryOutputFeedbackValue")
		}

		// Simple Field (feedbackValue)
		if pushErr := writeBuffer.PushContext("feedbackValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for feedbackValue")
		}
		_feedbackValueErr := writeBuffer.WriteSerializable(m.GetFeedbackValue())
		if popErr := writeBuffer.PopContext("feedbackValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for feedbackValue")
		}
		if _feedbackValueErr != nil {
			return errors.Wrap(_feedbackValueErr, "Error serializing 'feedbackValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryOutputFeedbackValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryOutputFeedbackValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) isBACnetConstructedDataBinaryOutputFeedbackValue() bool {
	return true
}

func (m *_BACnetConstructedDataBinaryOutputFeedbackValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
