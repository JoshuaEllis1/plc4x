//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "encoding/xml"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetErrorConfirmedPrivateTransfer struct {
    Parent *BACnetError
    IBACnetErrorConfirmedPrivateTransfer
}

// The corresponding interface
type IBACnetErrorConfirmedPrivateTransfer interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetErrorConfirmedPrivateTransfer) ServiceChoice() uint8 {
    return 0x12
}


func (m *BACnetErrorConfirmedPrivateTransfer) InitializeParent(parent *BACnetError) {
}

func NewBACnetErrorConfirmedPrivateTransfer() *BACnetError {
    child := &BACnetErrorConfirmedPrivateTransfer{
        Parent: NewBACnetError(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastBACnetErrorConfirmedPrivateTransfer(structType interface{}) BACnetErrorConfirmedPrivateTransfer {
    castFunc := func(typ interface{}) BACnetErrorConfirmedPrivateTransfer {
        if casted, ok := typ.(BACnetErrorConfirmedPrivateTransfer); ok {
            return casted
        }
        if casted, ok := typ.(*BACnetErrorConfirmedPrivateTransfer); ok {
            return *casted
        }
        if casted, ok := typ.(BACnetError); ok {
            return CastBACnetErrorConfirmedPrivateTransfer(casted.Child)
        }
        if casted, ok := typ.(*BACnetError); ok {
            return CastBACnetErrorConfirmedPrivateTransfer(casted.Child)
        }
        return BACnetErrorConfirmedPrivateTransfer{}
    }
    return castFunc(structType)
}

func (m *BACnetErrorConfirmedPrivateTransfer) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *BACnetErrorConfirmedPrivateTransfer) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetErrorConfirmedPrivateTransferParse(io *utils.ReadBuffer) (*BACnetError, error) {

    // Create a partially initialized instance
    _child := &BACnetErrorConfirmedPrivateTransfer{
        Parent: &BACnetError{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *BACnetErrorConfirmedPrivateTransfer) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetErrorConfirmedPrivateTransfer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            }
        }
    }
}

func (m *BACnetErrorConfirmedPrivateTransfer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorConfirmedPrivateTransfer"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}
