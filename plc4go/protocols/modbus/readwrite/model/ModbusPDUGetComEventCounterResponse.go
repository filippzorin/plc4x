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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUGetComEventCounterResponse is the corresponding interface of ModbusPDUGetComEventCounterResponse
type ModbusPDUGetComEventCounterResponse interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStatus returns Status (property field)
	GetStatus() uint16
	// GetEventCount returns EventCount (property field)
	GetEventCount() uint16
}

// ModbusPDUGetComEventCounterResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUGetComEventCounterResponse.
// This is useful for switch cases.
type ModbusPDUGetComEventCounterResponseExactly interface {
	ModbusPDUGetComEventCounterResponse
	isModbusPDUGetComEventCounterResponse() bool
}

// _ModbusPDUGetComEventCounterResponse is the data-structure of this message
type _ModbusPDUGetComEventCounterResponse struct {
	*_ModbusPDU
	Status     uint16
	EventCount uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUGetComEventCounterResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUGetComEventCounterResponse) GetFunctionFlag() uint8 {
	return 0x0B
}

func (m *_ModbusPDUGetComEventCounterResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUGetComEventCounterResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUGetComEventCounterResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUGetComEventCounterResponse) GetStatus() uint16 {
	return m.Status
}

func (m *_ModbusPDUGetComEventCounterResponse) GetEventCount() uint16 {
	return m.EventCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUGetComEventCounterResponse factory function for _ModbusPDUGetComEventCounterResponse
func NewModbusPDUGetComEventCounterResponse(status uint16, eventCount uint16) *_ModbusPDUGetComEventCounterResponse {
	_result := &_ModbusPDUGetComEventCounterResponse{
		Status:     status,
		EventCount: eventCount,
		_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUGetComEventCounterResponse(structType interface{}) ModbusPDUGetComEventCounterResponse {
	if casted, ok := structType.(ModbusPDUGetComEventCounterResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUGetComEventCounterResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUGetComEventCounterResponse) GetTypeName() string {
	return "ModbusPDUGetComEventCounterResponse"
}

func (m *_ModbusPDUGetComEventCounterResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUGetComEventCounterResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUGetComEventCounterResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUGetComEventCounterResponseParse(readBuffer utils.ReadBuffer, response bool) (ModbusPDUGetComEventCounterResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventCounterResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUGetComEventCounterResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint16("status", 16)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of ModbusPDUGetComEventCounterResponse")
	}
	status := _status

	// Simple Field (eventCount)
	_eventCount, _eventCountErr := readBuffer.ReadUint16("eventCount", 16)
	if _eventCountErr != nil {
		return nil, errors.Wrap(_eventCountErr, "Error parsing 'eventCount' field of ModbusPDUGetComEventCounterResponse")
	}
	eventCount := _eventCount

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventCounterResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUGetComEventCounterResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUGetComEventCounterResponse{
		_ModbusPDU: &_ModbusPDU{},
		Status:     status,
		EventCount: eventCount,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUGetComEventCounterResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventCounterResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUGetComEventCounterResponse")
		}

		// Simple Field (status)
		status := uint16(m.GetStatus())
		_statusErr := writeBuffer.WriteUint16("status", 16, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (eventCount)
		eventCount := uint16(m.GetEventCount())
		_eventCountErr := writeBuffer.WriteUint16("eventCount", 16, (eventCount))
		if _eventCountErr != nil {
			return errors.Wrap(_eventCountErr, "Error serializing 'eventCount' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventCounterResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUGetComEventCounterResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ModbusPDUGetComEventCounterResponse) isModbusPDUGetComEventCounterResponse() bool {
	return true
}

func (m *_ModbusPDUGetComEventCounterResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}