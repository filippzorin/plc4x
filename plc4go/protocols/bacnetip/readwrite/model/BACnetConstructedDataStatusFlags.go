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

// BACnetConstructedDataStatusFlags is the corresponding interface of BACnetConstructedDataStatusFlags
type BACnetConstructedDataStatusFlags interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetStatusFlagsTagged
}

// BACnetConstructedDataStatusFlagsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataStatusFlags.
// This is useful for switch cases.
type BACnetConstructedDataStatusFlagsExactly interface {
	BACnetConstructedDataStatusFlags
	isBACnetConstructedDataStatusFlags() bool
}

// _BACnetConstructedDataStatusFlags is the data-structure of this message
type _BACnetConstructedDataStatusFlags struct {
	*_BACnetConstructedData
	StatusFlags BACnetStatusFlagsTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStatusFlags) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStatusFlags) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STATUS_FLAGS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStatusFlags) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataStatusFlags) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStatusFlags) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStatusFlags) GetActualValue() BACnetStatusFlagsTagged {
	return CastBACnetStatusFlagsTagged(m.GetStatusFlags())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStatusFlags factory function for _BACnetConstructedDataStatusFlags
func NewBACnetConstructedDataStatusFlags(statusFlags BACnetStatusFlagsTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStatusFlags {
	_result := &_BACnetConstructedDataStatusFlags{
		StatusFlags:            statusFlags,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStatusFlags(structType interface{}) BACnetConstructedDataStatusFlags {
	if casted, ok := structType.(BACnetConstructedDataStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStatusFlags) GetTypeName() string {
	return "BACnetConstructedDataStatusFlags"
}

func (m *_BACnetConstructedDataStatusFlags) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataStatusFlags) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStatusFlags) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataStatusFlagsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataStatusFlags, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := _statusFlags.(BACnetStatusFlagsTagged)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Virtual field
	_actualValue := statusFlags
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStatusFlags")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataStatusFlags{
		StatusFlags:            statusFlags,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataStatusFlags) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStatusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStatusFlags")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusFlags")
		}
		_statusFlagsErr := writeBuffer.WriteSerializable(m.GetStatusFlags())
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusFlags")
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStatusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStatusFlags")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStatusFlags) isBACnetConstructedDataStatusFlags() bool {
	return true
}

func (m *_BACnetConstructedDataStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
