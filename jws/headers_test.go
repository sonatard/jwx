package jws_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	publicKey := `{"kty":"RSA",
	             "n": "0vx7agoebGcQSuuPiLJXZptN9nndrQmbXEps2aiAFbWhM78LhWx4cbbfAAtVT86zwu1RK7aPFFxuhDR1L6tSoc_BJECPebWKRXjBZCiFV4n3oknjhMstn64tZ_2W-5JsGY4Hc5n9yBXArwl93lqt7_RN5w6Cf0h4QyQ5v-65YGjQR0_FDW2QvzqY368QQMicAtaSqzs8KJZgnYb9c7d0zgdAZHzu6qMQvRL5hajrn1n91CbOpbISD08qNLyrdkt-bFTWhAI4vMQFh6WeZu0fM4lFd2NcRwr3XPksINHaQ-G_xBniIqbw0Ls1jF44-csFCur-kEgU8awapJzKnqDKgw",
	             "e":"AQAB",
	             "alg":"RS256",
	             "kid":"2011-04-29"}`
	jwkPublicKeySet, err := jwk.ParseString(publicKey)
	if err != nil {
		t.Fatal("Failed to parse RSA public key")
	}
	certChain := []string{
		"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",
	}
	values := map[string]interface{}{
		jws.AlgorithmKey:          jwa.ES256,
		jws.ContentTypeKey:        "example",
		jws.CriticalKey:           []string{"exp"},
		jws.JWKKey:                jwkPublicKeySet.Keys[0],
		jws.JWKSetURLKey:          "https://www.jwk.com/key.json",
		jws.TypeKey:               "JWT",
		jws.KeyIDKey:              "e9bc097a-ce51-4036-9562-d2ade882db0d",
		jws.X509CertChainKey:      certChain,
		jws.X509CertThumbprintKey: "QzY0NjREMjkyQTI4RTU2RkE4MUJBRDExNzY1MUY1N0I4QjFCODlBOQ",
		jws.X509URLKey:            "https://www.x509.com/key.pem",
	}

	t.Run("Type", func(t *testing.T) {
		var h jws.Headers = jws.NewHeaders()
		_ = h
	})
	t.Run("Sanity", func(t *testing.T) {
		h := jws.NewHeaders()
		if !assert.NoError(t, json.Unmarshal([]byte(publicKey), h), "unmarshal public key should succeed") {
			return
		}
		if !assert.NotEmpty(t, h.KeyID()) { // これあったっけ…
			return
		}
	})
	t.Run("Private parameters", func(t *testing.T) {
		t.Run("Without standard headers", func(t *testing.T) {
			t.Parallel()
			const src = `{ "foo": 1, "bar": "two", "baz": true }`
			h := jws.NewHeaders()
			if !assert.NoError(t, json.Unmarshal([]byte(src), h), "unmarshal should succeed") {
				return
			}

			expected := map[string]interface{}{
				"foo": 1.0, // JSON has no such thing as integers here
				"bar": "two",
				"baz": true,
			}

			for key, value := range expected {
				v, ok := h.Get(key)
				if !assert.True(t, ok, `h.Get(%#v) should succeed`, key) {
					return
				}
				if !assert.Equal(t, v, value, `h.Get(%#v) should return %#v`, key, value) {
					return
				}
			}

			buf, err := json.Marshal(h)
			if !assert.NoError(t, err, `json.Marshal should succeed`) {
				return
			}
			if !assert.Equal(t, `{"bar":"two","baz":true,"foo":1}`, string(buf), `json.Marshal should succeed`) {
				return
			}
		})
		t.Run("With standard headers", func(t *testing.T) {
			t.Parallel()

			const src = `{ "alg": "ES256", "foo": 1, "bar": "two", "baz": true }`
			h := jws.NewHeaders()
			if !assert.NoError(t, json.Unmarshal([]byte(src), h), "unmarshal should succeed") {
				return
			}

			expected := map[string]interface{}{
				"alg": jwa.ES256,
				"foo": 1.0, // JSON has no such thing as integers here
				"bar": "two",
				"baz": true,
			}

			for key, value := range expected {
				v, ok := h.Get(key)
				if !assert.True(t, ok, `h.Get(%#v) should succeed`, key) {
					return
				}
				if !assert.Equal(t, v, value, `h.Get(%#v) should return %#v`, key, value) {
					return
				}
			}

			buf, err := json.Marshal(h)
			if !assert.NoError(t, err, `json.Marshal should succeed`) {
				return
			}
			if !assert.Equal(t, `{"alg":"ES256","bar":"two","baz":true,"foo":1}`, string(buf), `json.Marshal should succeed`) {
				return
			}
		})
	})
	t.Run("Roundtrip", func(t *testing.T) {
		h := jws.NewHeaders()
		for k, v := range values {
			if !assert.NoError(t, h.Set(k, v), "h.Set should succeed for %s", k) {
				return
			}
			got, ok := h.Get(k)
			if !assert.True(t, ok, "h.Get should succeed for %s", k) {
				return
			}
			//fmt.Println(reflect.TypeOf(got).String())
			//fmt.Println(reflect.TypeOf(v).String())
			if !reflect.DeepEqual(v, got) {
				t.Fatalf("Values do not match: (%v, %v)", v, got)
			}
		}
	})
	t.Run("JSON Roundtrip", func(t *testing.T) {
		h := jws.NewHeaders()
		for k, v := range values {
			err := h.Set(k, v)
			if err != nil {
				t.Fatalf("Set failed for %s", k)
			}
			got, ok := h.Get(k)
			if !ok {
				t.Fatalf("Set failed for %s", k)
			}
			if !reflect.DeepEqual(v, got) {
				t.Fatalf("Values do not match: (%v, %v)", v, got)
			}
		}
		hByte, err := json.Marshal(h)
		if err != nil {
			t.Fatal("Failed to JSON marshal")
		}
		hNew := jws.NewHeaders()

		if !assert.NoError(t, json.Unmarshal(hByte, &hNew), "json.Unmarshal should succeed for headers") {
			return
		}
	})
	t.Run("RoundtripError", func(t *testing.T) {
		type dummyStruct struct {
			dummy1 int
			dummy2 float64
		}
		dummy := &dummyStruct{1, 3.4}

		values := map[string]interface{}{
			jws.AlgorithmKey:              dummy,
			jws.ContentTypeKey:            dummy,
			jws.CriticalKey:               dummy,
			jws.JWKKey:                    dummy,
			jws.JWKSetURLKey:              dummy,
			jws.KeyIDKey:                  dummy,
			jws.TypeKey:                   dummy,
			jws.X509CertChainKey:          dummy,
			jws.X509CertThumbprintKey:     dummy,
			jws.X509CertThumbprintS256Key: dummy,
			jws.X509URLKey:                dummy,
		}

		h := jws.NewHeaders()
		for k, v := range values {
			err := h.Set(k, v)
			if err == nil {
				t.Fatalf("Setting %s value should have failed", k)
			}
		}
		err := h.Set("default", dummy) // private params
		if err != nil {
			t.Fatalf("Setting %s value failed", "default")
		}
		for k := range values {
			_, ok := h.Get(k)
			if ok {
				t.Fatalf("Getting %s value should have failed", k)
			}
		}
		_, ok := h.Get("default")
		if !ok {
			t.Fatal("Failed to get default value")
		}
	})
}
