/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type IdentifyReplyCommandCurrentSenseLevels struct {
	*IdentifyReplyCommand
}

// The corresponding interface
type IIdentifyReplyCommandCurrentSenseLevels interface {
	IIdentifyReplyCommand
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *IdentifyReplyCommandCurrentSenseLevels) Attribute() Attribute {
	return Attribute_CurrentSenseLevels
}

func (m *IdentifyReplyCommandCurrentSenseLevels) GetAttribute() Attribute {
	return Attribute_CurrentSenseLevels
}

func (m *IdentifyReplyCommandCurrentSenseLevels) InitializeParent(parent *IdentifyReplyCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandCurrentSenseLevels factory function for IdentifyReplyCommandCurrentSenseLevels
func NewIdentifyReplyCommandCurrentSenseLevels() *IdentifyReplyCommand {
	child := &IdentifyReplyCommandCurrentSenseLevels{
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	child.Child = child
	return child.IdentifyReplyCommand
}

func CastIdentifyReplyCommandCurrentSenseLevels(structType interface{}) *IdentifyReplyCommandCurrentSenseLevels {
	castFunc := func(typ interface{}) *IdentifyReplyCommandCurrentSenseLevels {
		if casted, ok := typ.(IdentifyReplyCommandCurrentSenseLevels); ok {
			return &casted
		}
		if casted, ok := typ.(*IdentifyReplyCommandCurrentSenseLevels); ok {
			return casted
		}
		if casted, ok := typ.(IdentifyReplyCommand); ok {
			return CastIdentifyReplyCommandCurrentSenseLevels(casted.Child)
		}
		if casted, ok := typ.(*IdentifyReplyCommand); ok {
			return CastIdentifyReplyCommandCurrentSenseLevels(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *IdentifyReplyCommandCurrentSenseLevels) GetTypeName() string {
	return "IdentifyReplyCommandCurrentSenseLevels"
}

func (m *IdentifyReplyCommandCurrentSenseLevels) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandCurrentSenseLevels) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *IdentifyReplyCommandCurrentSenseLevels) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandCurrentSenseLevelsParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommand, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandCurrentSenseLevels"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandCurrentSenseLevels"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandCurrentSenseLevels{
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child.IdentifyReplyCommand, nil
}

func (m *IdentifyReplyCommandCurrentSenseLevels) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandCurrentSenseLevels"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandCurrentSenseLevels"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandCurrentSenseLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}