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
package org.apache.plc4x.java.knxnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ApduControlContainer extends Apdu implements Message {

  // Accessors for discriminator values.
  public Byte getControl() {
    return (byte) 1;
  }

  // Properties.
  protected final ApduControl controlApdu;

  // Arguments.
  protected final Short dataLength;

  public ApduControlContainer(
      boolean numbered, byte counter, ApduControl controlApdu, Short dataLength) {
    super(numbered, counter, dataLength);
    this.controlApdu = controlApdu;
    this.dataLength = dataLength;
  }

  public ApduControl getControlApdu() {
    return controlApdu;
  }

  @Override
  protected void serializeApduChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ApduControlContainer");

    // Simple Field (controlApdu)
    writeSimpleField("controlApdu", controlApdu, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("ApduControlContainer");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ApduControlContainer _value = this;

    // Simple field (controlApdu)
    lengthInBits += controlApdu.getLengthInBits();

    return lengthInBits;
  }

  public static ApduControlContainerBuilder staticParseBuilder(
      ReadBuffer readBuffer, Short dataLength) throws ParseException {
    readBuffer.pullContext("ApduControlContainer");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    ApduControl controlApdu =
        readSimpleField(
            "controlApdu",
            new DataReaderComplexDefault<>(() -> ApduControl.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("ApduControlContainer");
    // Create the instance
    return new ApduControlContainerBuilder(controlApdu, dataLength);
  }

  public static class ApduControlContainerBuilder implements Apdu.ApduBuilder {
    private final ApduControl controlApdu;
    private final Short dataLength;

    public ApduControlContainerBuilder(ApduControl controlApdu, Short dataLength) {

      this.controlApdu = controlApdu;
      this.dataLength = dataLength;
    }

    public ApduControlContainer build(boolean numbered, byte counter, Short dataLength) {
      ApduControlContainer apduControlContainer =
          new ApduControlContainer(numbered, counter, controlApdu, dataLength);
      return apduControlContainer;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ApduControlContainer)) {
      return false;
    }
    ApduControlContainer that = (ApduControlContainer) o;
    return (getControlApdu() == that.getControlApdu()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getControlApdu());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}