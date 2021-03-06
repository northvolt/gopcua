// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import "testing"

var testApplicationDescriptionBytes = [][]byte{
	{ // Single
		// ApplicationURI
		0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
		// ProductURI
		0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
		// ApplicationName
		0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d,
		0x6e, 0x61, 0x6d, 0x65,
		// ApplicationType
		0x00, 0x00, 0x00, 0x00,
		// GatewayServerURI
		0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
		// DiscoveryProfileURI
		0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x2d, 0x75, 0x72, 0x69,
		// DiscoveryURIs
		0x02, 0x00, 0x00, 0x00,
		0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x31,
		0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x32,
	},
	{ // Array
		// ArraySize
		0x02, 0x00, 0x00, 0x00,
		// ApplicationURI
		0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
		// ProductURI
		0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
		// ApplicationName
		0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d,
		0x6e, 0x61, 0x6d, 0x65,
		// ApplicationType
		0x00, 0x00, 0x00, 0x00,
		// GatewayServerURI
		0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
		// DiscoveryProfileURI
		0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x2d, 0x75, 0x72, 0x69,
		// DiscoveryURIs
		0x02, 0x00, 0x00, 0x00,
		0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x31,
		0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x32,
		// ApplicationURI
		0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
		// ProductURI
		0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
		// ApplicationName
		0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d,
		0x6e, 0x61, 0x6d, 0x65,
		// ApplicationType
		0x00, 0x00, 0x00, 0x00,
		// GatewayServerURI
		0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
		// DiscoveryProfileURI
		0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x2d, 0x75, 0x72, 0x69,
		// DiscoveryURIs
		0x02, 0x00, 0x00, 0x00,
		0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x31,
		0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x32,
	},
	{},
	{},
}

func TestDecodeApplicationDescription(t *testing.T) {
	a, err := DecodeApplicationDescription(testApplicationDescriptionBytes[0])
	if err != nil {
		t.Fatalf("Failed to decode ApplicationDescription: %s", err)
	}

	switch {
	case a.ApplicationURI.Get() != "app-uri":
		t.Errorf("ApplicationURI doesn't match. Want: %s, Got: %s", "app-uri", a.ApplicationURI.Get())
	}
	t.Log(a.String())
}

func TestDecodeApplicationDescriptionArray(t *testing.T) {
	a, err := DecodeApplicationDescriptionArray(testApplicationDescriptionBytes[1])
	if err != nil {
		t.Fatalf("Failed to decode ApplicationDescription: %s", err)
	}

	switch {
	case a.ArraySize != 2:
		t.Errorf("ArraySize doesn't match. Want: %d, Got: %d", 2, a.ArraySize)
	}
	t.Log(a)
}

func TestSerializeApplicationDescription(t *testing.T) {
	a := NewApplicationDescription(
		"app-uri",
		"prod-uri",
		"app-name",
		AppTypeServer,
		"gw-uri",
		"prof-uri",
		[]string{"discov-uri-1", "discov-uri-2"},
	)

	serialized, err := a.Serialize()
	if err != nil {
		t.Fatalf("Failed to serialize ApplicationDescription: %s", err)
	}

	for i, s := range serialized {
		x := testApplicationDescriptionBytes[0][i]
		if s != x {
			t.Errorf("Bytes doesn't match. Want: %#x, Got: %#x at %dth", x, s, i)
		}
	}
	t.Logf("%x", serialized)
}

func TestSerializeApplicationDescriptionArray(t *testing.T) {
	a := NewApplicationDescriptionArray(
		[]*ApplicationDescription{
			NewApplicationDescription(
				"app-uri",
				"prod-uri",
				"app-name",
				AppTypeServer,
				"gw-uri",
				"prof-uri",
				[]string{"discov-uri-1", "discov-uri-2"},
			),
			NewApplicationDescription(
				"app-uri",
				"prod-uri",
				"app-name",
				AppTypeServer,
				"gw-uri",
				"prof-uri",
				[]string{"discov-uri-1", "discov-uri-2"},
			),
		},
	)

	serialized, err := a.Serialize()
	if err != nil {
		t.Fatalf("Failed to serialize ApplicationDescription: %s", err)
	}

	for i, s := range serialized {
		x := testApplicationDescriptionBytes[1][i]
		if s != x {
			t.Errorf("Bytes doesn't match. Want: %#x, Got: %#x at %dth", x, s, i)
		}
	}
	t.Logf("%x", serialized)
}
