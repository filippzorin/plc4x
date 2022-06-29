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

// BACnetAccessZoneOccupancyState is an enum
type BACnetAccessZoneOccupancyState uint16

type IBACnetAccessZoneOccupancyState interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetAccessZoneOccupancyState_NORMAL                   BACnetAccessZoneOccupancyState = 0
	BACnetAccessZoneOccupancyState_BELOW_LOWER_LIMIT        BACnetAccessZoneOccupancyState = 1
	BACnetAccessZoneOccupancyState_AT_LOWER_LIMIT           BACnetAccessZoneOccupancyState = 2
	BACnetAccessZoneOccupancyState_AT_UPPER_LIMIT           BACnetAccessZoneOccupancyState = 3
	BACnetAccessZoneOccupancyState_ABOVE_UPPER_LIMIT        BACnetAccessZoneOccupancyState = 4
	BACnetAccessZoneOccupancyState_DISABLED                 BACnetAccessZoneOccupancyState = 5
	BACnetAccessZoneOccupancyState_NOT_SUPPORTED            BACnetAccessZoneOccupancyState = 6
	BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE BACnetAccessZoneOccupancyState = 0xFFFF
)

var BACnetAccessZoneOccupancyStateValues []BACnetAccessZoneOccupancyState

func init() {
	_ = errors.New
	BACnetAccessZoneOccupancyStateValues = []BACnetAccessZoneOccupancyState{
		BACnetAccessZoneOccupancyState_NORMAL,
		BACnetAccessZoneOccupancyState_BELOW_LOWER_LIMIT,
		BACnetAccessZoneOccupancyState_AT_LOWER_LIMIT,
		BACnetAccessZoneOccupancyState_AT_UPPER_LIMIT,
		BACnetAccessZoneOccupancyState_ABOVE_UPPER_LIMIT,
		BACnetAccessZoneOccupancyState_DISABLED,
		BACnetAccessZoneOccupancyState_NOT_SUPPORTED,
		BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAccessZoneOccupancyStateByValue(value uint16) BACnetAccessZoneOccupancyState {
	switch value {
	case 0:
		return BACnetAccessZoneOccupancyState_NORMAL
	case 0xFFFF:
		return BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetAccessZoneOccupancyState_BELOW_LOWER_LIMIT
	case 2:
		return BACnetAccessZoneOccupancyState_AT_LOWER_LIMIT
	case 3:
		return BACnetAccessZoneOccupancyState_AT_UPPER_LIMIT
	case 4:
		return BACnetAccessZoneOccupancyState_ABOVE_UPPER_LIMIT
	case 5:
		return BACnetAccessZoneOccupancyState_DISABLED
	case 6:
		return BACnetAccessZoneOccupancyState_NOT_SUPPORTED
	}
	return 0
}

func BACnetAccessZoneOccupancyStateByName(value string) (enum BACnetAccessZoneOccupancyState, ok bool) {
	ok = true
	switch value {
	case "NORMAL":
		enum = BACnetAccessZoneOccupancyState_NORMAL
	case "VENDOR_PROPRIETARY_VALUE":
		enum = BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE
	case "BELOW_LOWER_LIMIT":
		enum = BACnetAccessZoneOccupancyState_BELOW_LOWER_LIMIT
	case "AT_LOWER_LIMIT":
		enum = BACnetAccessZoneOccupancyState_AT_LOWER_LIMIT
	case "AT_UPPER_LIMIT":
		enum = BACnetAccessZoneOccupancyState_AT_UPPER_LIMIT
	case "ABOVE_UPPER_LIMIT":
		enum = BACnetAccessZoneOccupancyState_ABOVE_UPPER_LIMIT
	case "DISABLED":
		enum = BACnetAccessZoneOccupancyState_DISABLED
	case "NOT_SUPPORTED":
		enum = BACnetAccessZoneOccupancyState_NOT_SUPPORTED
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetAccessZoneOccupancyStateKnows(value uint16) bool {
	for _, typeValue := range BACnetAccessZoneOccupancyStateValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessZoneOccupancyState(structType interface{}) BACnetAccessZoneOccupancyState {
	castFunc := func(typ interface{}) BACnetAccessZoneOccupancyState {
		if sBACnetAccessZoneOccupancyState, ok := typ.(BACnetAccessZoneOccupancyState); ok {
			return sBACnetAccessZoneOccupancyState
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessZoneOccupancyState) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetAccessZoneOccupancyState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAccessZoneOccupancyStateParse(readBuffer utils.ReadBuffer) (BACnetAccessZoneOccupancyState, error) {
	val, err := readBuffer.ReadUint16("BACnetAccessZoneOccupancyState", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetAccessZoneOccupancyStateByValue(val), nil
}

func (e BACnetAccessZoneOccupancyState) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetAccessZoneOccupancyState", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccessZoneOccupancyState) PLC4XEnumName() string {
	switch e {
	case BACnetAccessZoneOccupancyState_NORMAL:
		return "NORMAL"
	case BACnetAccessZoneOccupancyState_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAccessZoneOccupancyState_BELOW_LOWER_LIMIT:
		return "BELOW_LOWER_LIMIT"
	case BACnetAccessZoneOccupancyState_AT_LOWER_LIMIT:
		return "AT_LOWER_LIMIT"
	case BACnetAccessZoneOccupancyState_AT_UPPER_LIMIT:
		return "AT_UPPER_LIMIT"
	case BACnetAccessZoneOccupancyState_ABOVE_UPPER_LIMIT:
		return "ABOVE_UPPER_LIMIT"
	case BACnetAccessZoneOccupancyState_DISABLED:
		return "DISABLED"
	case BACnetAccessZoneOccupancyState_NOT_SUPPORTED:
		return "NOT_SUPPORTED"
	}
	return ""
}

func (e BACnetAccessZoneOccupancyState) String() string {
	return e.PLC4XEnumName()
}
