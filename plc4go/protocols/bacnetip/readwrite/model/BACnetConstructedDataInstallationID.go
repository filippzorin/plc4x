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

// BACnetConstructedDataInstallationID is the corresponding interface of BACnetConstructedDataInstallationID
type BACnetConstructedDataInstallationID interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInstallationId returns InstallationId (property field)
	GetInstallationId() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataInstallationIDExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataInstallationID.
// This is useful for switch cases.
type BACnetConstructedDataInstallationIDExactly interface {
	BACnetConstructedDataInstallationID
	isBACnetConstructedDataInstallationID() bool
}

// _BACnetConstructedDataInstallationID is the data-structure of this message
type _BACnetConstructedDataInstallationID struct {
	*_BACnetConstructedData
	InstallationId BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInstallationID) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInstallationID) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INSTALLATION_ID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInstallationID) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataInstallationID) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInstallationID) GetInstallationId() BACnetApplicationTagUnsignedInteger {
	return m.InstallationId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInstallationID) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetInstallationId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInstallationID factory function for _BACnetConstructedDataInstallationID
func NewBACnetConstructedDataInstallationID(installationId BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInstallationID {
	_result := &_BACnetConstructedDataInstallationID{
		InstallationId:         installationId,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInstallationID(structType interface{}) BACnetConstructedDataInstallationID {
	if casted, ok := structType.(BACnetConstructedDataInstallationID); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInstallationID); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInstallationID) GetTypeName() string {
	return "BACnetConstructedDataInstallationID"
}

func (m *_BACnetConstructedDataInstallationID) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataInstallationID) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (installationId)
	lengthInBits += m.InstallationId.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInstallationID) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInstallationIDParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataInstallationID, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInstallationID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInstallationID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (installationId)
	if pullErr := readBuffer.PullContext("installationId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for installationId")
	}
	_installationId, _installationIdErr := BACnetApplicationTagParse(readBuffer)
	if _installationIdErr != nil {
		return nil, errors.Wrap(_installationIdErr, "Error parsing 'installationId' field")
	}
	installationId := _installationId.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("installationId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for installationId")
	}

	// Virtual field
	_actualValue := installationId
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInstallationID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInstallationID")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataInstallationID{
		InstallationId:         installationId,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataInstallationID) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInstallationID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInstallationID")
		}

		// Simple Field (installationId)
		if pushErr := writeBuffer.PushContext("installationId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for installationId")
		}
		_installationIdErr := writeBuffer.WriteSerializable(m.GetInstallationId())
		if popErr := writeBuffer.PopContext("installationId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for installationId")
		}
		if _installationIdErr != nil {
			return errors.Wrap(_installationIdErr, "Error serializing 'installationId' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInstallationID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInstallationID")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInstallationID) isBACnetConstructedDataInstallationID() bool {
	return true
}

func (m *_BACnetConstructedDataInstallationID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
