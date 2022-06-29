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

// BACnetEventType is an enum
type BACnetEventType uint16

type IBACnetEventType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetEventType_CHANGE_OF_BITSTRING       BACnetEventType = 0
	BACnetEventType_CHANGE_OF_STATE           BACnetEventType = 1
	BACnetEventType_CHANGE_OF_VALUE           BACnetEventType = 2
	BACnetEventType_COMMAND_FAILURE           BACnetEventType = 3
	BACnetEventType_FLOATING_LIMIT            BACnetEventType = 4
	BACnetEventType_OUT_OF_RANGE              BACnetEventType = 5
	BACnetEventType_CHANGE_OF_LIFE_SAFETY     BACnetEventType = 8
	BACnetEventType_EXTENDED                  BACnetEventType = 9
	BACnetEventType_BUFFER_READY              BACnetEventType = 10
	BACnetEventType_UNSIGNED_RANGE            BACnetEventType = 11
	BACnetEventType_ACCESS_EVENT              BACnetEventType = 13
	BACnetEventType_DOUBLE_OUT_OF_RANGE       BACnetEventType = 14
	BACnetEventType_SIGNED_OUT_OF_RANGE       BACnetEventType = 15
	BACnetEventType_UNSIGNED_OUT_OF_RANGE     BACnetEventType = 16
	BACnetEventType_CHANGE_OF_CHARACTERSTRING BACnetEventType = 17
	BACnetEventType_CHANGE_OF_STATUS_FLAGS    BACnetEventType = 18
	BACnetEventType_CHANGE_OF_RELIABILITY     BACnetEventType = 19
	BACnetEventType_NONE                      BACnetEventType = 20
	BACnetEventType_CHANGE_OF_DISCRETE_VALUE  BACnetEventType = 21
	BACnetEventType_CHANGE_OF_TIMER           BACnetEventType = 22
	BACnetEventType_VENDOR_PROPRIETARY_VALUE  BACnetEventType = 0xFFFF
)

var BACnetEventTypeValues []BACnetEventType

func init() {
	_ = errors.New
	BACnetEventTypeValues = []BACnetEventType{
		BACnetEventType_CHANGE_OF_BITSTRING,
		BACnetEventType_CHANGE_OF_STATE,
		BACnetEventType_CHANGE_OF_VALUE,
		BACnetEventType_COMMAND_FAILURE,
		BACnetEventType_FLOATING_LIMIT,
		BACnetEventType_OUT_OF_RANGE,
		BACnetEventType_CHANGE_OF_LIFE_SAFETY,
		BACnetEventType_EXTENDED,
		BACnetEventType_BUFFER_READY,
		BACnetEventType_UNSIGNED_RANGE,
		BACnetEventType_ACCESS_EVENT,
		BACnetEventType_DOUBLE_OUT_OF_RANGE,
		BACnetEventType_SIGNED_OUT_OF_RANGE,
		BACnetEventType_UNSIGNED_OUT_OF_RANGE,
		BACnetEventType_CHANGE_OF_CHARACTERSTRING,
		BACnetEventType_CHANGE_OF_STATUS_FLAGS,
		BACnetEventType_CHANGE_OF_RELIABILITY,
		BACnetEventType_NONE,
		BACnetEventType_CHANGE_OF_DISCRETE_VALUE,
		BACnetEventType_CHANGE_OF_TIMER,
		BACnetEventType_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetEventTypeByValue(value uint16) BACnetEventType {
	switch value {
	case 0:
		return BACnetEventType_CHANGE_OF_BITSTRING
	case 0xFFFF:
		return BACnetEventType_VENDOR_PROPRIETARY_VALUE
	case 1:
		return BACnetEventType_CHANGE_OF_STATE
	case 10:
		return BACnetEventType_BUFFER_READY
	case 11:
		return BACnetEventType_UNSIGNED_RANGE
	case 13:
		return BACnetEventType_ACCESS_EVENT
	case 14:
		return BACnetEventType_DOUBLE_OUT_OF_RANGE
	case 15:
		return BACnetEventType_SIGNED_OUT_OF_RANGE
	case 16:
		return BACnetEventType_UNSIGNED_OUT_OF_RANGE
	case 17:
		return BACnetEventType_CHANGE_OF_CHARACTERSTRING
	case 18:
		return BACnetEventType_CHANGE_OF_STATUS_FLAGS
	case 19:
		return BACnetEventType_CHANGE_OF_RELIABILITY
	case 2:
		return BACnetEventType_CHANGE_OF_VALUE
	case 20:
		return BACnetEventType_NONE
	case 21:
		return BACnetEventType_CHANGE_OF_DISCRETE_VALUE
	case 22:
		return BACnetEventType_CHANGE_OF_TIMER
	case 3:
		return BACnetEventType_COMMAND_FAILURE
	case 4:
		return BACnetEventType_FLOATING_LIMIT
	case 5:
		return BACnetEventType_OUT_OF_RANGE
	case 8:
		return BACnetEventType_CHANGE_OF_LIFE_SAFETY
	case 9:
		return BACnetEventType_EXTENDED
	}
	return 0
}

func BACnetEventTypeByName(value string) (enum BACnetEventType, ok bool) {
	ok = true
	switch value {
	case "CHANGE_OF_BITSTRING":
		enum = BACnetEventType_CHANGE_OF_BITSTRING
	case "VENDOR_PROPRIETARY_VALUE":
		enum = BACnetEventType_VENDOR_PROPRIETARY_VALUE
	case "CHANGE_OF_STATE":
		enum = BACnetEventType_CHANGE_OF_STATE
	case "BUFFER_READY":
		enum = BACnetEventType_BUFFER_READY
	case "UNSIGNED_RANGE":
		enum = BACnetEventType_UNSIGNED_RANGE
	case "ACCESS_EVENT":
		enum = BACnetEventType_ACCESS_EVENT
	case "DOUBLE_OUT_OF_RANGE":
		enum = BACnetEventType_DOUBLE_OUT_OF_RANGE
	case "SIGNED_OUT_OF_RANGE":
		enum = BACnetEventType_SIGNED_OUT_OF_RANGE
	case "UNSIGNED_OUT_OF_RANGE":
		enum = BACnetEventType_UNSIGNED_OUT_OF_RANGE
	case "CHANGE_OF_CHARACTERSTRING":
		enum = BACnetEventType_CHANGE_OF_CHARACTERSTRING
	case "CHANGE_OF_STATUS_FLAGS":
		enum = BACnetEventType_CHANGE_OF_STATUS_FLAGS
	case "CHANGE_OF_RELIABILITY":
		enum = BACnetEventType_CHANGE_OF_RELIABILITY
	case "CHANGE_OF_VALUE":
		enum = BACnetEventType_CHANGE_OF_VALUE
	case "NONE":
		enum = BACnetEventType_NONE
	case "CHANGE_OF_DISCRETE_VALUE":
		enum = BACnetEventType_CHANGE_OF_DISCRETE_VALUE
	case "CHANGE_OF_TIMER":
		enum = BACnetEventType_CHANGE_OF_TIMER
	case "COMMAND_FAILURE":
		enum = BACnetEventType_COMMAND_FAILURE
	case "FLOATING_LIMIT":
		enum = BACnetEventType_FLOATING_LIMIT
	case "OUT_OF_RANGE":
		enum = BACnetEventType_OUT_OF_RANGE
	case "CHANGE_OF_LIFE_SAFETY":
		enum = BACnetEventType_CHANGE_OF_LIFE_SAFETY
	case "EXTENDED":
		enum = BACnetEventType_EXTENDED
	default:
		enum = 0
		ok = false
	}
	return
}

func BACnetEventTypeKnows(value uint16) bool {
	for _, typeValue := range BACnetEventTypeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetEventType(structType interface{}) BACnetEventType {
	castFunc := func(typ interface{}) BACnetEventType {
		if sBACnetEventType, ok := typ.(BACnetEventType); ok {
			return sBACnetEventType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetEventType) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetEventType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventTypeParse(readBuffer utils.ReadBuffer) (BACnetEventType, error) {
	val, err := readBuffer.ReadUint16("BACnetEventType", 16)
	if err != nil {
		return 0, nil
	}
	return BACnetEventTypeByValue(val), nil
}

func (e BACnetEventType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetEventType", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetEventType) PLC4XEnumName() string {
	switch e {
	case BACnetEventType_CHANGE_OF_BITSTRING:
		return "CHANGE_OF_BITSTRING"
	case BACnetEventType_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetEventType_CHANGE_OF_STATE:
		return "CHANGE_OF_STATE"
	case BACnetEventType_BUFFER_READY:
		return "BUFFER_READY"
	case BACnetEventType_UNSIGNED_RANGE:
		return "UNSIGNED_RANGE"
	case BACnetEventType_ACCESS_EVENT:
		return "ACCESS_EVENT"
	case BACnetEventType_DOUBLE_OUT_OF_RANGE:
		return "DOUBLE_OUT_OF_RANGE"
	case BACnetEventType_SIGNED_OUT_OF_RANGE:
		return "SIGNED_OUT_OF_RANGE"
	case BACnetEventType_UNSIGNED_OUT_OF_RANGE:
		return "UNSIGNED_OUT_OF_RANGE"
	case BACnetEventType_CHANGE_OF_CHARACTERSTRING:
		return "CHANGE_OF_CHARACTERSTRING"
	case BACnetEventType_CHANGE_OF_STATUS_FLAGS:
		return "CHANGE_OF_STATUS_FLAGS"
	case BACnetEventType_CHANGE_OF_RELIABILITY:
		return "CHANGE_OF_RELIABILITY"
	case BACnetEventType_CHANGE_OF_VALUE:
		return "CHANGE_OF_VALUE"
	case BACnetEventType_NONE:
		return "NONE"
	case BACnetEventType_CHANGE_OF_DISCRETE_VALUE:
		return "CHANGE_OF_DISCRETE_VALUE"
	case BACnetEventType_CHANGE_OF_TIMER:
		return "CHANGE_OF_TIMER"
	case BACnetEventType_COMMAND_FAILURE:
		return "COMMAND_FAILURE"
	case BACnetEventType_FLOATING_LIMIT:
		return "FLOATING_LIMIT"
	case BACnetEventType_OUT_OF_RANGE:
		return "OUT_OF_RANGE"
	case BACnetEventType_CHANGE_OF_LIFE_SAFETY:
		return "CHANGE_OF_LIFE_SAFETY"
	case BACnetEventType_EXTENDED:
		return "EXTENDED"
	}
	return ""
}

func (e BACnetEventType) String() string {
	return e.PLC4XEnumName()
}
