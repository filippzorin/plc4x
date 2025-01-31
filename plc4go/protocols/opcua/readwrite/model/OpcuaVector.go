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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaVector is the corresponding interface of OpcuaVector
type OpcuaVector interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// OpcuaVectorExactly can be used when we want exactly this type and not a type which fulfills OpcuaVector.
// This is useful for switch cases.
type OpcuaVectorExactly interface {
	OpcuaVector
	isOpcuaVector() bool
}

// _OpcuaVector is the data-structure of this message
type _OpcuaVector struct {
	*_ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaVector) GetIdentifier() string {
	return "18809"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaVector) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_OpcuaVector) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewOpcuaVector factory function for _OpcuaVector
func NewOpcuaVector() *_OpcuaVector {
	_result := &_OpcuaVector{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpcuaVector(structType any) OpcuaVector {
	if casted, ok := structType.(OpcuaVector); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaVector); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaVector) GetTypeName() string {
	return "OpcuaVector"
}

func (m *_OpcuaVector) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_OpcuaVector) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaVectorParse(ctx context.Context, theBytes []byte, identifier string) (OpcuaVector, error) {
	return OpcuaVectorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func OpcuaVectorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (OpcuaVector, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("OpcuaVector"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaVector")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("OpcuaVector"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaVector")
	}

	// Create a partially initialized instance
	_child := &_OpcuaVector{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_OpcuaVector) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaVector) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaVector"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaVector")
		}

		if popErr := writeBuffer.PopContext("OpcuaVector"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaVector")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpcuaVector) isOpcuaVector() bool {
	return true
}

func (m *_OpcuaVector) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
