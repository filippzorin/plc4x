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

// TagClass is an enum
type TagClass uint8

type ITagClass interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	TagClass_APPLICATION_TAGS      TagClass = 0x0
	TagClass_CONTEXT_SPECIFIC_TAGS TagClass = 0x1
)

var TagClassValues []TagClass

func init() {
	_ = errors.New
	TagClassValues = []TagClass{
		TagClass_APPLICATION_TAGS,
		TagClass_CONTEXT_SPECIFIC_TAGS,
	}
}

func TagClassByValue(value uint8) (enum TagClass, ok bool) {
	switch value {
	case 0x0:
		return TagClass_APPLICATION_TAGS, true
	case 0x1:
		return TagClass_CONTEXT_SPECIFIC_TAGS, true
	}
	return 0, false
}

func TagClassByName(value string) (enum TagClass, ok bool) {
	switch value {
	case "APPLICATION_TAGS":
		return TagClass_APPLICATION_TAGS, true
	case "CONTEXT_SPECIFIC_TAGS":
		return TagClass_CONTEXT_SPECIFIC_TAGS, true
	}
	return 0, false
}

func TagClassKnows(value uint8) bool {
	for _, typeValue := range TagClassValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTagClass(structType interface{}) TagClass {
	castFunc := func(typ interface{}) TagClass {
		if sTagClass, ok := typ.(TagClass); ok {
			return sTagClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m TagClass) GetLengthInBits() uint16 {
	return 1
}

func (m TagClass) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TagClassParse(readBuffer utils.ReadBuffer) (TagClass, error) {
	val, err := readBuffer.ReadUint8("TagClass", 1)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TagClass")
	}
	if enum, ok := TagClassByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return TagClass(val), nil
	} else {
		return enum, nil
	}
}

func (e TagClass) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("TagClass", 1, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TagClass) PLC4XEnumName() string {
	switch e {
	case TagClass_APPLICATION_TAGS:
		return "APPLICATION_TAGS"
	case TagClass_CONTEXT_SPECIFIC_TAGS:
		return "CONTEXT_SPECIFIC_TAGS"
	}
	return ""
}

func (e TagClass) String() string {
	return e.PLC4XEnumName()
}
