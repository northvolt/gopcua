// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package datatypes

import (
	"testing"

	"github.com/wmnsk/gopcua/utils/codectest"
	"github.com/google/go-cmp/cmp"
)

func TestAnonymousIdentityToken(t *testing.T) {
	cases := []codectest.Case{
		{
			Name:   "normal",
			Struct: NewAnonymousIdentityToken("anonymous"),
			Bytes: []byte{
				0x09, 0x00, 0x00, 0x00,
				0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
			},
		},
	}
	codectest.Run(t, cases, func(b []byte) (codectest.S, error) {
		return DecodeAnonymousIdentityToken(b)
	})
}

func TestUserNameIdentityToken(t *testing.T) {
		},
	},
}

func TestDecodeUserNameIdentityToken(t *testing.T) {
	for _, c := range userNameIdentityTokenCases {
		got, err := DecodeUserNameIdentityToken(c.serialized)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, c.structured); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestSerializeUserNameIdentityToken(t *testing.T) {
	for _, c := range userNameIdentityTokenCases {
		got, err := c.structured.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, c.serialized); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestUserNameIdentityTokenLen(t *testing.T) {
	for _, c := range userNameIdentityTokenCases {
		got := c.structured.Len()

		if diff := cmp.Diff(got, len(c.serialized)); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

var x509IdentityTokenCases = []struct {
	description string
	structured  *X509IdentityToken
	serialized  []byte
}{
	{
		"normal",
		NewX509IdentityToken("x509", "certificate"),
		[]byte{
			// PolicyID
			0x04, 0x00, 0x00, 0x00, 0x78, 0x35, 0x30, 0x39,
			// CertificateData
			0x0b, 0x00, 0x00, 0x00, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
		},
	},
}

func TestDecodeX509IdentityToken(t *testing.T) {
	for _, c := range x509IdentityTokenCases {
		got, err := DecodeX509IdentityToken(c.serialized)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, c.structured); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestSerializeX509IdentityToken(t *testing.T) {
	for _, c := range x509IdentityTokenCases {
		got, err := c.structured.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, c.serialized); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestX509IdentityTokenLen(t *testing.T) {
	for _, c := range x509IdentityTokenCases {
		got := c.structured.Len()

		if diff := cmp.Diff(got, len(c.serialized)); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

var issuedIdentityTokenCases = []struct {
	description string
	structured  *IssuedIdentityToken
	serialized  []byte
}{
	{
		"normal",
		NewIssuedIdentityToken("issued", []byte{0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64}, "plain"),
		[]byte{
			// PolicyID
			0x06, 0x00, 0x00, 0x00, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64,
			// TokenData
			0x08, 0x00, 0x00, 0x00, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
			// EncryptionAlgorithm
			0x05, 0x00, 0x00, 0x00, 0x70, 0x6c, 0x61, 0x69, 0x6e,
		},
	},
}

func TestDecodeIssuedIdentityToken(t *testing.T) {
	for _, c := range issuedIdentityTokenCases {
		got, err := DecodeIssuedIdentityToken(c.serialized)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, c.structured); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestSerializeIssuedIdentityToken(t *testing.T) {
	for _, c := range issuedIdentityTokenCases {
		got, err := c.structured.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, c.serialized); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}

func TestIssuedIdentityTokenLen(t *testing.T) {
	for _, c := range issuedIdentityTokenCases {
		got := c.structured.Len()

		if diff := cmp.Diff(got, len(c.serialized)); diff != "" {
			t.Errorf("%s failed\n%s", c.description, diff)
		}
	}
}
