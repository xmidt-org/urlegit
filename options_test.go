// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func customSchemeVador(s string) error {
	if s == "wss" {
		return nil
	}
	return ErrSchemeNotAllowed
}

func customHostVador(h string) error {
	if h == "example.com" {
		return ErrDomainNotAllowed
	}
	return nil
}

func customIPVador(ip *net.IP) error {
	exclude := net.ParseIP("192.168.1.1")
	if ip == nil {
		return nil
	}
	if ip.Equal(exclude) {
		return ErrIPNotAllowed
	}
	return nil
}

func TestResolverOptionString(t *testing.T) {
	opt := WithResolver(nil)
	assert.Equal(t, "WithResolver(nil)", opt.String())

	opt = WithResolver(mockResolver)
	assert.Equal(t, "WithResolver(resolver)", opt.String())
}

func TestCustomSchemeVador(t *testing.T) {
	tests := []sharedTest{
		{
			description: "use custom scheme, match",
			opt:         CustomSchemeVador(customSchemeVador),
			host:        "wss://example.com",
			noHttp:      true,
		}, {
			description: "use custom scheme, no match",
			opt:         CustomSchemeVador(customSchemeVador),
			host:        "ftp://example.com",
			expectedErr: ErrSchemeNotAllowed,
			noHttp:      true,
		},
	}
	testCommon(t, tests)
}

func TestCustomSchemeVadorString(t *testing.T) {
	opt := CustomSchemeVador(customSchemeVador)
	assert.Equal(t, "CustomSchemeVador(vador)", opt.String())
}

func TestCustomHostVador(t *testing.T) {
	tests := []sharedTest{
		{
			description: "use custom host vador, no match",
			opt:         CustomHostVador(customHostVador),
			host:        "http://testing.test",
		}, {
			description: "use custom host vador, match",
			opt:         CustomHostVador(customHostVador),
			host:        "http://example.com",
			expectedErr: ErrDomainNotAllowed,
		},
	}
	testCommon(t, tests)
}

func TestCustomHostVadorString(t *testing.T) {
	opt := CustomHostVador(customHostVador)
	assert.Equal(t, "CustomHostVador(vador)", opt.String())
}

func TestCustomIPVador(t *testing.T) {
	tests := []sharedTest{
		{
			description: "use custom ip vador, no match",
			opt:         CustomIPVador(customIPVador),
			host:        "http://192.168.1.2",
		}, {
			description: "use custom ip vador, match",
			opt:         CustomIPVador(customIPVador),
			host:        "http://192.168.1.1",
			expectedErr: ErrIPNotAllowed,
		},
	}
	testCommon(t, tests)
}

func TestCustomIPVadorString(t *testing.T) {
	opt := CustomIPVador(customIPVador)
	assert.Equal(t, "CustomIPVador(vador)", opt.String())
}
