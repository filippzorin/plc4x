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

// BACnetConstructedDataLoopAction is the corresponding interface of BACnetConstructedDataLoopAction
type BACnetConstructedDataLoopAction interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAction returns Action (property field)
	GetAction() BACnetActionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetActionTagged
}

// BACnetConstructedDataLoopActionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLoopAction.
// This is useful for switch cases.
type BACnetConstructedDataLoopActionExactly interface {
	BACnetConstructedDataLoopAction
	isBACnetConstructedDataLoopAction() bool
}

// _BACnetConstructedDataLoopAction is the data-structure of this message
type _BACnetConstructedDataLoopAction struct {
	*_BACnetConstructedData
	Action BACnetActionTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLoopAction) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LOOP
}

func (m *_BACnetConstructedDataLoopAction) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLoopAction) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLoopAction) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLoopAction) GetAction() BACnetActionTagged {
	return m.Action
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLoopAction) GetActualValue() BACnetActionTagged {
	return CastBACnetActionTagged(m.GetAction())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLoopAction factory function for _BACnetConstructedDataLoopAction
func NewBACnetConstructedDataLoopAction(action BACnetActionTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLoopAction {
	_result := &_BACnetConstructedDataLoopAction{
		Action:                 action,
		TagNumber:              tagNumber,
		ArrayIndexArgument:     arrayIndexArgument,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLoopAction(structType interface{}) BACnetConstructedDataLoopAction {
	if casted, ok := structType.(BACnetConstructedDataLoopAction); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoopAction); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLoopAction) GetTypeName() string {
	return "BACnetConstructedDataLoopAction"
}

func (m *_BACnetConstructedDataLoopAction) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLoopAction) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (action)
	lengthInBits += m.Action.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLoopAction) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLoopActionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLoopAction, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoopAction"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLoopAction")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (action)
	if pullErr := readBuffer.PullContext("action"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for action")
	}
	_action, _actionErr := BACnetActionTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _actionErr != nil {
		return nil, errors.Wrap(_actionErr, "Error parsing 'action' field")
	}
	action := _action.(BACnetActionTagged)
	if closeErr := readBuffer.CloseContext("action"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for action")
	}

	// Virtual field
	_actualValue := action
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoopAction"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLoopAction")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLoopAction{
		Action:                 action,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLoopAction) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoopAction"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLoopAction")
		}

		// Simple Field (action)
		if pushErr := writeBuffer.PushContext("action"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for action")
		}
		_actionErr := writeBuffer.WriteSerializable(m.GetAction())
		if popErr := writeBuffer.PopContext("action"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for action")
		}
		if _actionErr != nil {
			return errors.Wrap(_actionErr, "Error serializing 'action' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoopAction"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLoopAction")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLoopAction) isBACnetConstructedDataLoopAction() bool {
	return true
}

func (m *_BACnetConstructedDataLoopAction) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
