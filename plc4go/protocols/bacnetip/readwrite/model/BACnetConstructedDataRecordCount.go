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

// BACnetConstructedDataRecordCount is the corresponding interface of BACnetConstructedDataRecordCount
type BACnetConstructedDataRecordCount interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRecordCount returns RecordCount (property field)
	GetRecordCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataRecordCountExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRecordCount.
// This is useful for switch cases.
type BACnetConstructedDataRecordCountExactly interface {
	BACnetConstructedDataRecordCount
	isBACnetConstructedDataRecordCount() bool
}

// _BACnetConstructedDataRecordCount is the data-structure of this message
type _BACnetConstructedDataRecordCount struct {
	*_BACnetConstructedData
	RecordCount BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRecordCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRecordCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RECORD_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRecordCount) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataRecordCount) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRecordCount) GetRecordCount() BACnetApplicationTagUnsignedInteger {
	return m.RecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRecordCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetRecordCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRecordCount factory function for _BACnetConstructedDataRecordCount
func NewBACnetConstructedDataRecordCount(recordCount BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRecordCount {
	_result := &_BACnetConstructedDataRecordCount{
		RecordCount:            recordCount,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRecordCount(structType interface{}) BACnetConstructedDataRecordCount {
	if casted, ok := structType.(BACnetConstructedDataRecordCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRecordCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRecordCount) GetTypeName() string {
	return "BACnetConstructedDataRecordCount"
}

func (m *_BACnetConstructedDataRecordCount) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataRecordCount) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (recordCount)
	lengthInBits += m.RecordCount.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRecordCount) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRecordCountParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRecordCount, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRecordCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (recordCount)
	if pullErr := readBuffer.PullContext("recordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for recordCount")
	}
	_recordCount, _recordCountErr := BACnetApplicationTagParse(readBuffer)
	if _recordCountErr != nil {
		return nil, errors.Wrap(_recordCountErr, "Error parsing 'recordCount' field")
	}
	recordCount := _recordCount.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("recordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for recordCount")
	}

	// Virtual field
	_actualValue := recordCount
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRecordCount")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataRecordCount{
		RecordCount:            recordCount,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataRecordCount) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRecordCount")
		}

		// Simple Field (recordCount)
		if pushErr := writeBuffer.PushContext("recordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for recordCount")
		}
		_recordCountErr := writeBuffer.WriteSerializable(m.GetRecordCount())
		if popErr := writeBuffer.PopContext("recordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for recordCount")
		}
		if _recordCountErr != nil {
			return errors.Wrap(_recordCountErr, "Error serializing 'recordCount' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRecordCount")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRecordCount) isBACnetConstructedDataRecordCount() bool {
	return true
}

func (m *_BACnetConstructedDataRecordCount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
