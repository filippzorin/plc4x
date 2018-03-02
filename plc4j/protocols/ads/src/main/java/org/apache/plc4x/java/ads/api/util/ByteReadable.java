/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 */
package org.apache.plc4x.java.ads.api.util;

import io.netty.buffer.ByteBuf;
import org.apache.commons.io.HexDump;

import java.io.ByteArrayOutputStream;
import java.io.IOException;

@FunctionalInterface
public interface ByteReadable extends LengthSupplier {

    default byte[] getBytes() {
        ByteBuf byteBuf = getByteBuf();
        byte[] result = new byte[byteBuf.writerIndex()];
        byteBuf.readBytes(result);
        return result;
    }

    ByteBuf getByteBuf();

    @Override
    default long getCalculatedLength() {
        return getByteBuf().readableBytes();
    }

    default String dump() throws IOException {
        try (ByteArrayOutputStream byteArrayOutputStream = new ByteArrayOutputStream()) {
            HexDump.dump(getBytes(), 0, byteArrayOutputStream, 0);
            return toString() + HexDump.EOL + byteArrayOutputStream.toString();
        }
    }
}
