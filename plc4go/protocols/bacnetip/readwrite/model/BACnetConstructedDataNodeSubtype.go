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

// BACnetConstructedDataNodeSubtype is the corresponding interface of BACnetConstructedDataNodeSubtype
type BACnetConstructedDataNodeSubtype interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNodeSubType returns NodeSubType (property field)
	GetNodeSubType() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataNodeSubtypeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNodeSubtype.
// This is useful for switch cases.
type BACnetConstructedDataNodeSubtypeExactly interface {
	BACnetConstructedDataNodeSubtype
	isBACnetConstructedDataNodeSubtype() bool
}

// _BACnetConstructedDataNodeSubtype is the data-structure of this message
type _BACnetConstructedDataNodeSubtype struct {
	*_BACnetConstructedData
	NodeSubType BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNodeSubtype) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNodeSubtype) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NODE_SUBTYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNodeSubtype) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataNodeSubtype) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNodeSubtype) GetNodeSubType() BACnetApplicationTagCharacterString {
	return m.NodeSubType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNodeSubtype) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetNodeSubType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNodeSubtype factory function for _BACnetConstructedDataNodeSubtype
func NewBACnetConstructedDataNodeSubtype(nodeSubType BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNodeSubtype {
	_result := &_BACnetConstructedDataNodeSubtype{
		NodeSubType:            nodeSubType,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNodeSubtype(structType interface{}) BACnetConstructedDataNodeSubtype {
	if casted, ok := structType.(BACnetConstructedDataNodeSubtype); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNodeSubtype); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNodeSubtype) GetTypeName() string {
	return "BACnetConstructedDataNodeSubtype"
}

func (m *_BACnetConstructedDataNodeSubtype) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataNodeSubtype) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nodeSubType)
	lengthInBits += m.NodeSubType.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNodeSubtype) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNodeSubtypeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNodeSubtype, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNodeSubtype"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNodeSubtype")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeSubType)
	if pullErr := readBuffer.PullContext("nodeSubType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeSubType")
	}
	_nodeSubType, _nodeSubTypeErr := BACnetApplicationTagParse(readBuffer)
	if _nodeSubTypeErr != nil {
		return nil, errors.Wrap(_nodeSubTypeErr, "Error parsing 'nodeSubType' field")
	}
	nodeSubType := _nodeSubType.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("nodeSubType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeSubType")
	}

	// Virtual field
	_actualValue := nodeSubType
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNodeSubtype"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNodeSubtype")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataNodeSubtype{
		NodeSubType:            nodeSubType,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataNodeSubtype) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNodeSubtype"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNodeSubtype")
		}

		// Simple Field (nodeSubType)
		if pushErr := writeBuffer.PushContext("nodeSubType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodeSubType")
		}
		_nodeSubTypeErr := writeBuffer.WriteSerializable(m.GetNodeSubType())
		if popErr := writeBuffer.PopContext("nodeSubType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodeSubType")
		}
		if _nodeSubTypeErr != nil {
			return errors.Wrap(_nodeSubTypeErr, "Error serializing 'nodeSubType' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNodeSubtype"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNodeSubtype")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNodeSubtype) isBACnetConstructedDataNodeSubtype() bool {
	return true
}

func (m *_BACnetConstructedDataNodeSubtype) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
