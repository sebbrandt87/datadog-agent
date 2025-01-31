// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUTF16LEParserHandleMessages(t *testing.T) {
	parser := NewEncodedText(UTF16LE)
	testMsg := []byte{'F', 0x0, 'o', 0x0, 'o', 0x0}
	msg, _, _, _, err := parser.Parse(testMsg)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", string(msg))

	// We should support BOM
	testMsg = []byte{0xFF, 0xFE, 'F', 0x0, 'o', 0x0, 'o', 0x0}
	msg, _, _, _, err = parser.Parse(testMsg)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", string(msg))

	// BOM overrides endianness
	testMsg = []byte{0xFE, 0xFF, 0x0, 'F', 0x0, 'o', 0x0, 'o'}
	msg, _, _, _, err = parser.Parse(testMsg)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", string(msg))
}

func TestUTF16BEParserHandleMessages(t *testing.T) {
	parser := NewEncodedText(UTF16BE)
	testMsg := []byte{0x0, 'F', 0x0, 'o', 0x0, 'o'}
	msg, _, _, _, err := parser.Parse(testMsg)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", string(msg))

	// We should support BOM
	testMsg = []byte{0xFE, 0xFF, 0x0, 'F', 0x0, 'o', 0x0, 'o'}
	msg, _, _, _, err = parser.Parse(testMsg)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", string(msg))

	// BOM overrides endianness
	testMsg = []byte{0xFF, 0xFE, 'F', 0x0, 'o', 0x0, 'o', 0x0}
	msg, _, _, _, err = parser.Parse(testMsg)
	assert.Nil(t, err)
	assert.Equal(t, "Foo", string(msg))
}

func TestSHIFTJISParserHandleMessages(t *testing.T) {
	parser := NewEncodedText(SHIFTJIS)
	testMsg := []byte{0x93, 0xfa, 0x96, 0x7b}
	msg, _, _, _, err := parser.Parse(testMsg)
	assert.Nil(t, err)
	assert.Equal(t, "日本", string(msg))
}
