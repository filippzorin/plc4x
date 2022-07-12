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

// DestinationAddressType is an enum
type DestinationAddressType uint8

type IDestinationAddressType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	DestinationAddressType_PointToPointToMultiPoint DestinationAddressType = 0x03
	DestinationAddressType_PointToMultiPoint        DestinationAddressType = 0x05
	DestinationAddressType_PointToPoint             DestinationAddressType = 0x06
)

var DestinationAddressTypeValues []DestinationAddressType

func init() {
	_ = errors.New
	DestinationAddressTypeValues = []DestinationAddressType{
		DestinationAddressType_PointToPointToMultiPoint,
		DestinationAddressType_PointToMultiPoint,
		DestinationAddressType_PointToPoint,
	}
}

func DestinationAddressTypeByValue(value uint8) (enum DestinationAddressType, ok bool) {
	switch value {
	case 0x03:
		return DestinationAddressType_PointToPointToMultiPoint, true
	case 0x05:
		return DestinationAddressType_PointToMultiPoint, true
	case 0x06:
		return DestinationAddressType_PointToPoint, true
	}
	return 0, false
}

func DestinationAddressTypeByName(value string) (enum DestinationAddressType, ok bool) {
	switch value {
	case "PointToPointToMultiPoint":
		return DestinationAddressType_PointToPointToMultiPoint, true
	case "PointToMultiPoint":
		return DestinationAddressType_PointToMultiPoint, true
	case "PointToPoint":
		return DestinationAddressType_PointToPoint, true
	}
	return 0, false
}

func DestinationAddressTypeKnows(value uint8) bool {
	for _, typeValue := range DestinationAddressTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDestinationAddressType(structType interface{}) DestinationAddressType {
	castFunc := func(typ interface{}) DestinationAddressType {
		if sDestinationAddressType, ok := typ.(DestinationAddressType); ok {
			return sDestinationAddressType
		}
		return 0
	}
	return castFunc(structType)
}

func (m DestinationAddressType) GetLengthInBits() uint16 {
	return 3
}

func (m DestinationAddressType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DestinationAddressTypeParse(readBuffer utils.ReadBuffer) (DestinationAddressType, error) {
	val, err := readBuffer.ReadUint8("DestinationAddressType", 3)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DestinationAddressType")
	}
	if enum, ok := DestinationAddressTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return DestinationAddressType(val), nil
	} else {
		return enum, nil
	}
}

func (e DestinationAddressType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("DestinationAddressType", 3, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DestinationAddressType) PLC4XEnumName() string {
	switch e {
	case DestinationAddressType_PointToPointToMultiPoint:
		return "PointToPointToMultiPoint"
	case DestinationAddressType_PointToMultiPoint:
		return "PointToMultiPoint"
	case DestinationAddressType_PointToPoint:
		return "PointToPoint"
	}
	return ""
}

func (e DestinationAddressType) String() string {
	return e.PLC4XEnumName()
}
