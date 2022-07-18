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

package readwrite

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Code generated by code-generation. DO NOT EDIT.

type CbusXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m CbusXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	switch typeName {
	case "MeteringData":
		return model.MeteringDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "EnableControlData":
		return model.EnableControlDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ApplicationAddress2":
		return model.ApplicationAddress2Parse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ApplicationAddress1":
		return model.ApplicationAddress1Parse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "RequestContext":
		return model.RequestContextParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TriggerControlData":
		return model.TriggerControlDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NetworkNumber":
		return model.NetworkNumberParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "RequestTermination":
		return model.RequestTerminationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ReplyOrConfirmation":
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 16)
		if err != nil {
			return nil, err
		}
		messageLength := uint16(parsedUint1)
		// TODO: find a way to parse the sub types
		var requestContext model.RequestContext
		return model.ReplyOrConfirmationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions, messageLength, requestContext)
	case "CBusMessage":
		isResponse := parserArguments[0] == "true"
		// TODO: find a way to parse the sub types
		var requestContext model.RequestContext
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		parsedUint3, err := strconv.ParseUint(parserArguments[3], 10, 16)
		if err != nil {
			return nil, err
		}
		messageLength := uint16(parsedUint3)
		return model.CBusMessageParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), isResponse, requestContext, cBusOptions, messageLength)
	case "CBusOptions":
		return model.CBusOptionsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CALDataOrSetParameter":
		return model.CALDataOrSetParameterParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TemperatureBroadcastData":
		return model.TemperatureBroadcastDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "PanicStatus":
		return model.PanicStatusParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "IdentifyReplyCommandUnitSummary":
		return model.IdentifyReplyCommandUnitSummaryParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BridgeCount":
		return model.BridgeCountParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "PowerUp":
		return model.PowerUpParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Reply":
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 16)
		if err != nil {
			return nil, err
		}
		replyLength := uint16(parsedUint1)
		// TODO: find a way to parse the sub types
		var requestContext model.RequestContext
		return model.ReplyParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions, replyLength, requestContext)
	case "InterfaceOptions1PowerUpSettings":
		return model.InterfaceOptions1PowerUpSettingsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "MonitoredSAL":
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		return model.MonitoredSALParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "ParameterChange":
		return model.ParameterChangeParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ReplyNetwork":
		return model.ReplyNetworkParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SerialNumber":
		return model.SerialNumberParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Confirmation":
		return model.ConfirmationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusPointToMultiPointCommand":
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		return model.CBusPointToMultiPointCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "StatusRequest":
		return model.StatusRequestParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "InterfaceOptions3":
		return model.InterfaceOptions3Parse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "InterfaceOptions1":
		return model.InterfaceOptions1Parse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "InterfaceOptions2":
		return model.InterfaceOptions2Parse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SecurityData":
		return model.SecurityDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NetworkProtocolControlInformation":
		return model.NetworkProtocolControlInformationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusHeader":
		return model.CBusHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Request":
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 16)
		if err != nil {
			return nil, err
		}
		messageLength := uint16(parsedUint1)
		return model.RequestParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions, messageLength)
	case "Alpha":
		return model.AlphaParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CALData":
		// TODO: find a way to parse the sub types
		var requestContext model.RequestContext
		return model.CALDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), requestContext)
	case "Checksum":
		return model.ChecksumParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CALReply":
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		// TODO: find a way to parse the sub types
		var requestContext model.RequestContext
		return model.CALReplyParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions, requestContext)
	case "CustomManufacturer":
		return model.CustomManufacturerParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "LightingData":
		return model.LightingDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NetworkRoute":
		return model.NetworkRouteParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "StandardFormatStatusReply":
		return model.StandardFormatStatusReplyParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ResponseTermination":
		return model.ResponseTerminationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SALData":
		applicationId, _ := model.ApplicationIdByName(parserArguments[0])
		return model.SALDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), applicationId)
	case "CBusCommand":
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		return model.CBusCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "TamperStatus":
		return model.TamperStatusParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "IdentifyReplyCommand":
		attribute, _ := model.AttributeByName(parserArguments[0])
		parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 5)
		if err != nil {
			return nil, err
		}
		numBytes := uint8(parsedUint1)
		return model.IdentifyReplyCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), attribute, numBytes)
	case "CBusConstants":
		return model.CBusConstantsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SerialInterfaceAddress":
		return model.SerialInterfaceAddressParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ZoneStatus":
		return model.ZoneStatusParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BridgeAddress":
		return model.BridgeAddressParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "StatusByte":
		return model.StatusByteParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "LightingLabelOptions":
		return model.LightingLabelOptionsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ExtendedStatusHeader":
		return model.ExtendedStatusHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CustomTypes":
		return model.CustomTypesParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TriggerControlLabelOptions":
		return model.TriggerControlLabelOptionsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "StatusHeader":
		return model.StatusHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "EncodedReply":
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		// TODO: find a way to parse the sub types
		var requestContext model.RequestContext
		return model.EncodedReplyParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions, requestContext)
	case "UnitAddress":
		return model.UnitAddressParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ExtendedFormatStatusReply":
		return model.ExtendedFormatStatusReplyParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SecurityArmCode":
		return model.SecurityArmCodeParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusPointToPointCommand":
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		return model.CBusPointToPointCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "CBusPointToPointToMultipointCommand":
		// TODO: find a way to parse the sub types
		var cBusOptions model.CBusOptions
		return model.CBusPointToPointToMultipointCommandParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "LogicAssignment":
		return model.LogicAssignmentParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
