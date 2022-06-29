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

// UnitStatus is an enum
type UnitStatus uint8

type IUnitStatus interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	UnitStatus_OK          UnitStatus = 0
	UnitStatus_NACK        UnitStatus = 1
	UnitStatus_NO_RESPONSE UnitStatus = 2
)

var UnitStatusValues []UnitStatus

func init() {
	_ = errors.New
	UnitStatusValues = []UnitStatus{
		UnitStatus_OK,
		UnitStatus_NACK,
		UnitStatus_NO_RESPONSE,
	}
}

func UnitStatusByValue(value uint8) UnitStatus {
	switch value {
	case 0:
		return UnitStatus_OK
	case 1:
		return UnitStatus_NACK
	case 2:
		return UnitStatus_NO_RESPONSE
	}
	return 0
}

func UnitStatusByName(value string) (enum UnitStatus, ok bool) {
	ok = true
	switch value {
	case "OK":
		enum = UnitStatus_OK
	case "NACK":
		enum = UnitStatus_NACK
	case "NO_RESPONSE":
		enum = UnitStatus_NO_RESPONSE
	default:
		enum = 0
		ok = false
	}
	return
}

func UnitStatusKnows(value uint8) bool {
	for _, typeValue := range UnitStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastUnitStatus(structType interface{}) UnitStatus {
	castFunc := func(typ interface{}) UnitStatus {
		if sUnitStatus, ok := typ.(UnitStatus); ok {
			return sUnitStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m UnitStatus) GetLengthInBits() uint16 {
	return 8
}

func (m UnitStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func UnitStatusParse(readBuffer utils.ReadBuffer) (UnitStatus, error) {
	val, err := readBuffer.ReadUint8("UnitStatus", 8)
	if err != nil {
		return 0, nil
	}
	return UnitStatusByValue(val), nil
}

func (e UnitStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("UnitStatus", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e UnitStatus) PLC4XEnumName() string {
	switch e {
	case UnitStatus_OK:
		return "OK"
	case UnitStatus_NACK:
		return "NACK"
	case UnitStatus_NO_RESPONSE:
		return "NO_RESPONSE"
	}
	return ""
}

func (e UnitStatus) String() string {
	return e.PLC4XEnumName()
}
