// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/gopcua/datatypes"
)

var readRequestCases = []struct {
	description string
	structured  *ReadRequest
	serialized  []byte
}{
	{
		"normal",
		NewReadRequest(
			NewRequestHeader(
				datatypes.NewOpaqueNodeID(0x00, []byte{
					0x08, 0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11,
					0xa6, 0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
				}),
				time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
				1, 0, 0, "", NewNullAdditionalHeader(), nil,
			),
			0, TimestampsToReturnBoth,
			datatypes.NewReadValueID(
				datatypes.NewFourByteNodeID(0, 2256),
				datatypes.IntegerIDValue,
				"", 0, "",
			),
		),
		[]byte{
			// TypeID
			0x01, 0x00, 0x77, 0x02,
			// AuthenticationToken
			0x05, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x08,
			0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11, 0xa6,
			0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
			// Timestamp
			0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
			// RequestHandle
			0x01, 0x00, 0x00, 0x00,
			// ReturnDiagnostics
			0x00, 0x00, 0x00, 0x00,
			// AuditEntryID
			0xff, 0xff, 0xff, 0xff,
			//TimeoutHint
			0x00, 0x00, 0x00, 0x00,
			// AdditionalHeader
			0x00, 0x00, 0x00,
			// MaxAge
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			// TimestampToReturn
			0x02, 0x00, 0x00, 0x00,
			// NodesToRead
			0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0xd0, 0x08,
			0x0d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff,
			0x00, 0x00, 0xff, 0xff, 0xff, 0xff,
		},
	},
}

func TestDecodeReadRequest(t *testing.T) {
	for _, c := range readRequestCases {
		got, err := DecodeReadRequest(c.serialized)
		if err != nil {
			t.Fatal(err)
		}

		// need to clear Payload here.
		got.Payload = nil

		if diff := cmp.Diff(got, c.structured, decodeCmpOpt); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestSerializeReadRequest(t *testing.T) {
	for _, c := range readRequestCases {
		got, err := c.structured.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, c.serialized); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestReadRequestLen(t *testing.T) {
	for _, c := range readRequestCases {
		got := c.structured.Len()

		if diff := cmp.Diff(got, len(c.serialized)); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestReadRequestServiceType(t *testing.T) {
	for _, c := range readRequestCases {
		if c.structured.ServiceType() != ServiceTypeReadRequest {
			t.Errorf(
				"ServiceType doesn't match. Want: %d, Got: %d",
				ServiceTypeReadRequest,
				c.structured.ServiceType(),
			)
		}
	}
}
