// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"net"
	"net/url"
)

const (
	mockLoopbackURL        = "http://mock-loopback.com"
	mockPrivateLoopbackURL = "http://mock-private-loopback.com"
	mockLoopbackPrivateURL = "http://mock-loopback-private.com"
	mockPrivateURL         = "http://mock-private.com"
	mockUnsupportedURL     = "http://mock-unsupported.com"
)

func mockResolver(s string) ([]net.IP, error) {
	loopback := net.ParseIP("127.0.0.1")
	local := net.ParseIP("192.168.1.1")

	switch s {
	case getFQDN(mockLoopbackURL):
		return []net.IP{loopback}, nil
	case getFQDN(mockPrivateLoopbackURL):
		return []net.IP{local, loopback}, nil
	case getFQDN(mockLoopbackPrivateURL):
		return []net.IP{loopback, local}, nil
	case getFQDN(mockPrivateURL):
		return []net.IP{local}, nil
	case getFQDN(mockUnsupportedURL):
	}

	return nil, errAny
}

func getFQDN(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u.Hostname()
}
