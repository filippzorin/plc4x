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

// BACnetConstructedDataAutoSlaveDiscovery is the corresponding interface of BACnetConstructedDataAutoSlaveDiscovery
type BACnetConstructedDataAutoSlaveDiscovery interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAutoSlaveDiscovery returns AutoSlaveDiscovery (property field)
	GetAutoSlaveDiscovery() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataAutoSlaveDiscoveryExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAutoSlaveDiscovery.
// This is useful for switch cases.
type BACnetConstructedDataAutoSlaveDiscoveryExactly interface {
	BACnetConstructedDataAutoSlaveDiscovery
	isBACnetConstructedDataAutoSlaveDiscovery() bool
}

// _BACnetConstructedDataAutoSlaveDiscovery is the data-structure of this message
type _BACnetConstructedDataAutoSlaveDiscovery struct {
	*_BACnetConstructedData
	AutoSlaveDiscovery BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAutoSlaveDiscovery) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAutoSlaveDiscovery) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTO_SLAVE_DISCOVERY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAutoSlaveDiscovery) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAutoSlaveDiscovery) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAutoSlaveDiscovery) GetAutoSlaveDiscovery() BACnetApplicationTagBoolean {
	return m.AutoSlaveDiscovery
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAutoSlaveDiscovery) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetAutoSlaveDiscovery())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAutoSlaveDiscovery factory function for _BACnetConstructedDataAutoSlaveDiscovery
func NewBACnetConstructedDataAutoSlaveDiscovery(autoSlaveDiscovery BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAutoSlaveDiscovery {
	_result := &_BACnetConstructedDataAutoSlaveDiscovery{
		AutoSlaveDiscovery:     autoSlaveDiscovery,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAutoSlaveDiscovery(structType interface{}) BACnetConstructedDataAutoSlaveDiscovery {
	if casted, ok := structType.(BACnetConstructedDataAutoSlaveDiscovery); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAutoSlaveDiscovery); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAutoSlaveDiscovery) GetTypeName() string {
	return "BACnetConstructedDataAutoSlaveDiscovery"
}

func (m *_BACnetConstructedDataAutoSlaveDiscovery) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAutoSlaveDiscovery) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (autoSlaveDiscovery)
	lengthInBits += m.AutoSlaveDiscovery.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAutoSlaveDiscovery) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAutoSlaveDiscoveryParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAutoSlaveDiscovery, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAutoSlaveDiscovery"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAutoSlaveDiscovery")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (autoSlaveDiscovery)
	if pullErr := readBuffer.PullContext("autoSlaveDiscovery"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for autoSlaveDiscovery")
	}
	_autoSlaveDiscovery, _autoSlaveDiscoveryErr := BACnetApplicationTagParse(readBuffer)
	if _autoSlaveDiscoveryErr != nil {
		return nil, errors.Wrap(_autoSlaveDiscoveryErr, "Error parsing 'autoSlaveDiscovery' field")
	}
	autoSlaveDiscovery := _autoSlaveDiscovery.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("autoSlaveDiscovery"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for autoSlaveDiscovery")
	}

	// Virtual field
	_actualValue := autoSlaveDiscovery
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAutoSlaveDiscovery"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAutoSlaveDiscovery")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAutoSlaveDiscovery{
		AutoSlaveDiscovery:     autoSlaveDiscovery,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAutoSlaveDiscovery) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAutoSlaveDiscovery"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAutoSlaveDiscovery")
		}

		// Simple Field (autoSlaveDiscovery)
		if pushErr := writeBuffer.PushContext("autoSlaveDiscovery"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for autoSlaveDiscovery")
		}
		_autoSlaveDiscoveryErr := writeBuffer.WriteSerializable(m.GetAutoSlaveDiscovery())
		if popErr := writeBuffer.PopContext("autoSlaveDiscovery"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for autoSlaveDiscovery")
		}
		if _autoSlaveDiscoveryErr != nil {
			return errors.Wrap(_autoSlaveDiscoveryErr, "Error serializing 'autoSlaveDiscovery' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAutoSlaveDiscovery"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAutoSlaveDiscovery")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAutoSlaveDiscovery) isBACnetConstructedDataAutoSlaveDiscovery() bool {
	return true
}

func (m *_BACnetConstructedDataAutoSlaveDiscovery) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
