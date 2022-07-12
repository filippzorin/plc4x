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

// SALCommandType is an enum
type SALCommandType uint8

type ISALCommandType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	SALCommandType_OFF            SALCommandType = 0x00
	SALCommandType_ON             SALCommandType = 0x01
	SALCommandType_RAMP_TO_LEVEL  SALCommandType = 0x02
	SALCommandType_TERMINATE_RAMP SALCommandType = 0x03
)

var SALCommandTypeValues []SALCommandType

func init() {
	_ = errors.New
	SALCommandTypeValues = []SALCommandType{
		SALCommandType_OFF,
		SALCommandType_ON,
		SALCommandType_RAMP_TO_LEVEL,
		SALCommandType_TERMINATE_RAMP,
	}
}

func SALCommandTypeByValue(value uint8) (enum SALCommandType, ok bool) {
	switch value {
	case 0x00:
		return SALCommandType_OFF, true
	case 0x01:
		return SALCommandType_ON, true
	case 0x02:
		return SALCommandType_RAMP_TO_LEVEL, true
	case 0x03:
		return SALCommandType_TERMINATE_RAMP, true
	}
	return 0, false
}

func SALCommandTypeByName(value string) (enum SALCommandType, ok bool) {
	switch value {
	case "OFF":
		return SALCommandType_OFF, true
	case "ON":
		return SALCommandType_ON, true
	case "RAMP_TO_LEVEL":
		return SALCommandType_RAMP_TO_LEVEL, true
	case "TERMINATE_RAMP":
		return SALCommandType_TERMINATE_RAMP, true
	}
	return 0, false
}

func SALCommandTypeKnows(value uint8) bool {
	for _, typeValue := range SALCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastSALCommandType(structType interface{}) SALCommandType {
	castFunc := func(typ interface{}) SALCommandType {
		if sSALCommandType, ok := typ.(SALCommandType); ok {
			return sSALCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m SALCommandType) GetLengthInBits() uint16 {
	return 4
}

func (m SALCommandType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALCommandTypeParse(readBuffer utils.ReadBuffer) (SALCommandType, error) {
	val, err := readBuffer.ReadUint8("SALCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading SALCommandType")
	}
	if enum, ok := SALCommandTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return SALCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e SALCommandType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("SALCommandType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e SALCommandType) PLC4XEnumName() string {
	switch e {
	case SALCommandType_OFF:
		return "OFF"
	case SALCommandType_ON:
		return "ON"
	case SALCommandType_RAMP_TO_LEVEL:
		return "RAMP_TO_LEVEL"
	case SALCommandType_TERMINATE_RAMP:
		return "TERMINATE_RAMP"
	}
	return ""
}

func (e SALCommandType) String() string {
	return e.PLC4XEnumName()
}
