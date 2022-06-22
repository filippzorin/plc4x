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

// LRawCon is the corresponding interface of LRawCon
type LRawCon interface {
	utils.LengthAware
	utils.Serializable
	CEMI
}

// LRawConExactly can be used when we want exactly this type and not a type which fulfills LRawCon.
// This is useful for switch cases.
type LRawConExactly interface {
	LRawCon
	isLRawCon() bool
}

// _LRawCon is the data-structure of this message
type _LRawCon struct {
	*_CEMI

	// Arguments.
	Size uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LRawCon) GetMessageCode() uint8 {
	return 0x2F
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LRawCon) InitializeParent(parent CEMI) {}

func (m *_LRawCon) GetParent() CEMI {
	return m._CEMI
}

// NewLRawCon factory function for _LRawCon
func NewLRawCon(size uint16) *_LRawCon {
	_result := &_LRawCon{
		Size:  size,
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLRawCon(structType interface{}) LRawCon {
	if casted, ok := structType.(LRawCon); ok {
		return casted
	}
	if casted, ok := structType.(*LRawCon); ok {
		return *casted
	}
	return nil
}

func (m *_LRawCon) GetTypeName() string {
	return "LRawCon"
}

func (m *_LRawCon) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_LRawCon) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_LRawCon) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LRawConParse(readBuffer utils.ReadBuffer, size uint16) (LRawCon, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LRawCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LRawCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LRawCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LRawCon")
	}

	// Create a partially initialized instance
	_child := &_LRawCon{
		_CEMI: &_CEMI{},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_LRawCon) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LRawCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LRawCon")
		}

		if popErr := writeBuffer.PopContext("LRawCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LRawCon")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_LRawCon) isLRawCon() bool {
	return true
}

func (m *_LRawCon) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
