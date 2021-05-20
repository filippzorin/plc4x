//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type SysexCommandReportFirmwareResponse struct {
	MajorVersion uint8
	MinorVersion uint8
	FileName     []int8
	Parent       *SysexCommand
}

// The corresponding interface
type ISysexCommandReportFirmwareResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SysexCommandReportFirmwareResponse) CommandType() uint8 {
	return 0x79
}

func (m *SysexCommandReportFirmwareResponse) Response() bool {
	return true
}

func (m *SysexCommandReportFirmwareResponse) InitializeParent(parent *SysexCommand) {
}

func NewSysexCommandReportFirmwareResponse(majorVersion uint8, minorVersion uint8, fileName []int8) *SysexCommand {
	child := &SysexCommandReportFirmwareResponse{
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		FileName:     fileName,
		Parent:       NewSysexCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastSysexCommandReportFirmwareResponse(structType interface{}) *SysexCommandReportFirmwareResponse {
	castFunc := func(typ interface{}) *SysexCommandReportFirmwareResponse {
		if casted, ok := typ.(SysexCommandReportFirmwareResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*SysexCommandReportFirmwareResponse); ok {
			return casted
		}
		if casted, ok := typ.(SysexCommand); ok {
			return CastSysexCommandReportFirmwareResponse(casted.Child)
		}
		if casted, ok := typ.(*SysexCommand); ok {
			return CastSysexCommandReportFirmwareResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SysexCommandReportFirmwareResponse) GetTypeName() string {
	return "SysexCommandReportFirmwareResponse"
}

func (m *SysexCommandReportFirmwareResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SysexCommandReportFirmwareResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	// Manual Array Field (fileName)
	fileName := m.FileName
	lengthInBits += FirmataUtilsLengthSysexString(fileName) * 8

	return lengthInBits
}

func (m *SysexCommandReportFirmwareResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SysexCommandReportFirmwareResponseParse(readBuffer utils.ReadBuffer) (*SysexCommand, error) {
	if pullErr := readBuffer.PullContext("SysexCommandReportFirmwareResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (majorVersion)
	majorVersion, _majorVersionErr := readBuffer.ReadUint8("majorVersion", 8)
	if _majorVersionErr != nil {
		return nil, errors.Wrap(_majorVersionErr, "Error parsing 'majorVersion' field")
	}

	// Simple Field (minorVersion)
	minorVersion, _minorVersionErr := readBuffer.ReadUint8("minorVersion", 8)
	if _minorVersionErr != nil {
		return nil, errors.Wrap(_minorVersionErr, "Error parsing 'minorVersion' field")
	}
	if pullErr := readBuffer.PullContext("fileName", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Manual Array Field (fileName)
	// Terminated array
	_fileNameList := make([]int8, 0)
	for !((bool)(FirmataUtilsIsSysexEnd(readBuffer))) {
		_fileNameList = append(_fileNameList, ((int8)(FirmataUtilsParseSysexString(readBuffer))))

	}
	fileName := make([]int8, len(_fileNameList))
	for i := 0; i < len(_fileNameList); i++ {
		fileName[i] = int8(_fileNameList[i])
	}
	if closeErr := readBuffer.CloseContext("fileName", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("SysexCommandReportFirmwareResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandReportFirmwareResponse{
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		FileName:     fileName,
		Parent:       &SysexCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *SysexCommandReportFirmwareResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandReportFirmwareResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (majorVersion)
		majorVersion := uint8(m.MajorVersion)
		_majorVersionErr := writeBuffer.WriteUint8("majorVersion", 8, (majorVersion))
		if _majorVersionErr != nil {
			return errors.Wrap(_majorVersionErr, "Error serializing 'majorVersion' field")
		}

		// Simple Field (minorVersion)
		minorVersion := uint8(m.MinorVersion)
		_minorVersionErr := writeBuffer.WriteUint8("minorVersion", 8, (minorVersion))
		if _minorVersionErr != nil {
			return errors.Wrap(_minorVersionErr, "Error serializing 'minorVersion' field")
		}

		// Manual Array Field (fileName)
		if m.FileName != nil {
			if pushErr := writeBuffer.PushContext("fileName", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, Element := range m.FileName {
				FirmataUtilsSerializeSysexString(writeBuffer, Element)
			}
			if popErr := writeBuffer.PopContext("fileName", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("SysexCommandReportFirmwareResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *SysexCommandReportFirmwareResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}