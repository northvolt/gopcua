// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/gopcua/datatypes"
	"github.com/wmnsk/gopcua/utils/codectest"
)

func TestCloseSessionRequest(t *testing.T) {
	cases := []codectest.Case{
		{
			Name: "normal",
			Struct: NewCloseSessionRequest(
				NewRequestHeader(
					datatypes.NewOpaqueNodeID(0x00, []byte{
						0x08, 0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11,
						0xa6, 0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
					}),
					time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
					1, 0, 0, "", NewNullAdditionalHeader(), nil,
				),
				true,
			),
			Bytes: []byte{ // CloseSessionRequest
				// TypeID
				0x01, 0x00, 0xd9, 0x01,
				// RequestHeader
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
				// DeleteSubscription
				0x01,
			},
		},
	}
	codectest.Run(t, cases, func(b []byte) (codectest.S, error) {
		v, err := DecodeCloseSessionRequest(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})

	t.Run("service-id", func(t *testing.T) {
		id := new(CloseSessionRequest).ServiceType()
		if got, want := id, uint16(ServiceTypeCloseSessionRequest); got != want {
			t.Fatalf("got %d want %d", got, want)
		}
	})
}

// option to regard []T{} and []T{nil} as equal
// https://godoc.org/github.com/google/go-cmp/cmp#example-Option--EqualEmpty
var decodeCmpOpt = cmp.FilterValues(func(x, y interface{}) bool {
	vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
	return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
		(vx.Kind() == reflect.Slice) && (vx.Len() == 0 && vy.Len() == 0)
}, cmp.Comparer(func(_, _ interface{}) bool { return true }))
