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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetDoorStatus is an enum
type BACnetDoorStatus uint16

type IBACnetDoorStatus interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetDoorStatus_CLOSED                   BACnetDoorStatus = 0
	BACnetDoorStatus_OPENED                   BACnetDoorStatus = 1
	BACnetDoorStatus_UNKNOWN                  BACnetDoorStatus = 2
	BACnetDoorStatus_DOOR_FAULT               BACnetDoorStatus = 3
	BACnetDoorStatus_UNUSED                   BACnetDoorStatus = 4
	BACnetDoorStatus_NONE                     BACnetDoorStatus = 5
	BACnetDoorStatus_CLOSING                  BACnetDoorStatus = 6
	BACnetDoorStatus_OPENING                  BACnetDoorStatus = 7
	BACnetDoorStatus_SAFETY_LOCKED            BACnetDoorStatus = 8
	BACnetDoorStatus_LIMITED_OPENED           BACnetDoorStatus = 9
	BACnetDoorStatus_VENDOR_PROPRIETARY_VALUE BACnetDoorStatus = 0xFFFF
)

var BACnetDoorStatusValues []BACnetDoorStatus

func init() {
	_ = errors.New
	BACnetDoorStatusValues = []BACnetDoorStatus{
		BACnetDoorStatus_CLOSED,
		BACnetDoorStatus_OPENED,
		BACnetDoorStatus_UNKNOWN,
		BACnetDoorStatus_DOOR_FAULT,
		BACnetDoorStatus_UNUSED,
		BACnetDoorStatus_NONE,
		BACnetDoorStatus_CLOSING,
		BACnetDoorStatus_OPENING,
		BACnetDoorStatus_SAFETY_LOCKED,
		BACnetDoorStatus_LIMITED_OPENED,
		BACnetDoorStatus_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetDoorStatusByValue(value uint16) (enum BACnetDoorStatus, ok bool) {
	switch value {
	case 0:
		return BACnetDoorStatus_CLOSED, true
	case 0xFFFF:
		return BACnetDoorStatus_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetDoorStatus_OPENED, true
	case 2:
		return BACnetDoorStatus_UNKNOWN, true
	case 3:
		return BACnetDoorStatus_DOOR_FAULT, true
	case 4:
		return BACnetDoorStatus_UNUSED, true
	case 5:
		return BACnetDoorStatus_NONE, true
	case 6:
		return BACnetDoorStatus_CLOSING, true
	case 7:
		return BACnetDoorStatus_OPENING, true
	case 8:
		return BACnetDoorStatus_SAFETY_LOCKED, true
	case 9:
		return BACnetDoorStatus_LIMITED_OPENED, true
	}
	return 0, false
}

func BACnetDoorStatusByName(value string) (enum BACnetDoorStatus, ok bool) {
	switch value {
	case "CLOSED":
		return BACnetDoorStatus_CLOSED, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetDoorStatus_VENDOR_PROPRIETARY_VALUE, true
	case "OPENED":
		return BACnetDoorStatus_OPENED, true
	case "UNKNOWN":
		return BACnetDoorStatus_UNKNOWN, true
	case "DOOR_FAULT":
		return BACnetDoorStatus_DOOR_FAULT, true
	case "UNUSED":
		return BACnetDoorStatus_UNUSED, true
	case "NONE":
		return BACnetDoorStatus_NONE, true
	case "CLOSING":
		return BACnetDoorStatus_CLOSING, true
	case "OPENING":
		return BACnetDoorStatus_OPENING, true
	case "SAFETY_LOCKED":
		return BACnetDoorStatus_SAFETY_LOCKED, true
	case "LIMITED_OPENED":
		return BACnetDoorStatus_LIMITED_OPENED, true
	}
	return 0, false
}

func BACnetDoorStatusKnows(value uint16) bool {
	for _, typeValue := range BACnetDoorStatusValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetDoorStatus(structType interface{}) BACnetDoorStatus {
	castFunc := func(typ interface{}) BACnetDoorStatus {
		if sBACnetDoorStatus, ok := typ.(BACnetDoorStatus); ok {
			return sBACnetDoorStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetDoorStatus) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetDoorStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDoorStatusParse(readBuffer utils.ReadBuffer) (BACnetDoorStatus, error) {
	val, err := readBuffer.ReadUint16("BACnetDoorStatus", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetDoorStatus")
	}
	if enum, ok := BACnetDoorStatusByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetDoorStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetDoorStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetDoorStatus", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetDoorStatus) PLC4XEnumName() string {
	switch e {
	case BACnetDoorStatus_CLOSED:
		return "CLOSED"
	case BACnetDoorStatus_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetDoorStatus_OPENED:
		return "OPENED"
	case BACnetDoorStatus_UNKNOWN:
		return "UNKNOWN"
	case BACnetDoorStatus_DOOR_FAULT:
		return "DOOR_FAULT"
	case BACnetDoorStatus_UNUSED:
		return "UNUSED"
	case BACnetDoorStatus_NONE:
		return "NONE"
	case BACnetDoorStatus_CLOSING:
		return "CLOSING"
	case BACnetDoorStatus_OPENING:
		return "OPENING"
	case BACnetDoorStatus_SAFETY_LOCKED:
		return "SAFETY_LOCKED"
	case BACnetDoorStatus_LIMITED_OPENED:
		return "LIMITED_OPENED"
	}
	return ""
}

func (e BACnetDoorStatus) String() string {
	return e.PLC4XEnumName()
}
