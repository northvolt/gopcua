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

var openSecureChannelRequestCases = []struct {
	description string
	structured  *OpenSecureChannelRequest
	serialized  []byte
}{
	{
		"normal",
		NewOpenSecureChannelRequest(
			NewRequestHeader(
				datatypes.NewOpaqueNodeID(0x00, []byte{
					0x08, 0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11,
					0xa6, 0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
				}),
				time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
				1, 0, 0, "", NewNullAdditionalHeader(), nil,
			),
			0,
			ReqTypeIssue,
			SecModeNone,
			6000000,
			nil,
		),
		[]byte{
			// TypeID
			0x01, 0x00, 0xbe, 0x01,
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
			// TimeoutHint
			0x00, 0x00, 0x00, 0x00,
			// AdditionalHeader
			0x00, 0x00, 0x00,
			// ClientProtocolVersion
			0x00, 0x00, 0x00, 0x00,
			// SecurityTokenRequestType
			0x00, 0x00, 0x00, 0x00,
			// MessageSecurityMode
			0x01, 0x00, 0x00, 0x00,
			// ClientNonce
			0xff, 0xff, 0xff, 0xff,
			// RequestedLifetime
			0x80, 0x8d, 0x5b, 0x00,
		},
	},
}

func TestDecodeOpenSecureChannelRequest(t *testing.T) {
	for _, c := range openSecureChannelRequestCases {
		got, err := DecodeOpenSecureChannelRequest(c.serialized)
		if err != nil {
			t.Fatal(err)
		}

		// need to clear Payload here.
		got.Payload = nil

		if diff := cmp.Diff(got, c.structured); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestSerializeOpenSecureChannelRequest(t *testing.T) {
	for _, c := range openSecureChannelRequestCases {
		got, err := c.structured.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, c.serialized); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestOpenSecureChannelRequestLen(t *testing.T) {
	for _, c := range openSecureChannelRequestCases {
		got := c.structured.Len()

		if diff := cmp.Diff(got, len(c.serialized)); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestOpenSecureChannelRequestServiceType(t *testing.T) {
	for _, c := range openSecureChannelRequestCases {
		if c.structured.ServiceType() != ServiceTypeOpenSecureChannelRequest {
			t.Errorf(
				"ServiceType doesn't match. Want: %d, Got: %d",
				ServiceTypeOpenSecureChannelRequest,
				c.structured.ServiceType(),
			)
		}
	}
}
