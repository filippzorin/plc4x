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

// ApduDataExtDomainAddressSerialNumberWrite is the corresponding interface of ApduDataExtDomainAddressSerialNumberWrite
type ApduDataExtDomainAddressSerialNumberWrite interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtDomainAddressSerialNumberWriteExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtDomainAddressSerialNumberWrite.
// This is useful for switch cases.
type ApduDataExtDomainAddressSerialNumberWriteExactly interface {
	ApduDataExtDomainAddressSerialNumberWrite
	isApduDataExtDomainAddressSerialNumberWrite() bool
}

// _ApduDataExtDomainAddressSerialNumberWrite is the data-structure of this message
type _ApduDataExtDomainAddressSerialNumberWrite struct {
	*_ApduDataExt

	// Arguments.
	Length uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressSerialNumberWrite) GetExtApciType() uint8 {
	return 0x2E
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressSerialNumberWrite) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtDomainAddressSerialNumberWrite) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtDomainAddressSerialNumberWrite factory function for _ApduDataExtDomainAddressSerialNumberWrite
func NewApduDataExtDomainAddressSerialNumberWrite(length uint8) *_ApduDataExtDomainAddressSerialNumberWrite {
	_result := &_ApduDataExtDomainAddressSerialNumberWrite{
		Length:       length,
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressSerialNumberWrite(structType interface{}) ApduDataExtDomainAddressSerialNumberWrite {
	if casted, ok := structType.(ApduDataExtDomainAddressSerialNumberWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressSerialNumberWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressSerialNumberWrite) GetTypeName() string {
	return "ApduDataExtDomainAddressSerialNumberWrite"
}

func (m *_ApduDataExtDomainAddressSerialNumberWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtDomainAddressSerialNumberWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressSerialNumberWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtDomainAddressSerialNumberWriteParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtDomainAddressSerialNumberWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressSerialNumberWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressSerialNumberWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressSerialNumberWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressSerialNumberWrite")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtDomainAddressSerialNumberWrite{
		_ApduDataExt: &_ApduDataExt{},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtDomainAddressSerialNumberWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressSerialNumberWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressSerialNumberWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressSerialNumberWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressSerialNumberWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressSerialNumberWrite) isApduDataExtDomainAddressSerialNumberWrite() bool {
	return true
}

func (m *_ApduDataExtDomainAddressSerialNumberWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
