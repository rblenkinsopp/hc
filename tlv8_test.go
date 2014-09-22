package gohap

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "encoding/hex"
    "bytes"
)

func TestTLV8SetByte(t *testing.T) {
    container := TLV8Container{}
    container.SetByte(1, 0xF)
    assert.Equal(t, container.GetByte(1), byte(0xF))
}

func TestTLV8GetBytes(t *testing.T) {
    data := "0102AFFA"
    rawMessage, _ := hex.DecodeString(data)    
    message := bytes.NewBuffer(rawMessage)
    container, err := ReadTLV8(message)
    assert.Nil(t, err)
    assert.Equal(t, container.GetBytes(1), []byte{0xAF, 0xFA})
}

func TestTLV8GetBytesFromMultipleSource(t *testing.T) {
    data := "0102AFFA0103BFFBAA"
    rawMessage, _ := hex.DecodeString(data)
    
    message := bytes.NewBuffer(rawMessage)
    container, err := ReadTLV8(message)
    assert.Nil(t, err)
    assert.Equal(t, container.GetBytes(1), []byte{0xAF, 0xFA, 0xBF, 0xFB, 0xAA})
}

func TestTLV8SetMoreThanMaxBytes(t *testing.T) {
    container := &TLV8Container{}
    data := "00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF" // 384 bytes
    bytes, _ := hex.DecodeString(data)
    assert.Equal(t, len(bytes), 384)
    
    container.SetBytes(1, bytes)
    
    // split up in 255 chunks
    // 01(type)FF(length=255)bytes...01(type)81(length=129)bytes...
    expected_data := "01FF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEE0181FF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF" // 384 bytes
    expected_bytes, _ := hex.DecodeString(expected_data)
    assert.Equal(t, container.BytesBuffer().Bytes(), expected_bytes)
}

func TestTLV8SetBytes(t *testing.T) {
    container := &TLV8Container{}
    container.SetBytes(1, []byte{0xAF, 0xFA})
    assert.Equal(t, container.GetBytes(1), []byte{0xAF, 0xFA})
}

func TestTLV8BytesBuffer(t *testing.T) {
    container := &TLV8Container{}
    container.SetBytes(1, []byte{0xAF, 0xFA})
    
    assert.Equal(t, container.BytesBuffer().Bytes(), []byte{0x01, 0x02, 0xAF, 0xFA})
}

func TestTLV8String(t *testing.T) {
    container := &TLV8Container{}
    container.SetString(1, "Hello World")
    
    assert.Equal(t, container.GetString(1), "Hello World")
}