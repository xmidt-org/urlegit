// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForbidSubnetOption(t *testing.T) {
	tests := []sharedTest{
		{
			description: "forbid subnet",
			opt:         ForbidSubnet("10.0.0.0/8"),
			host:        "http://192.168.1.1",
		}, {
			description: "forbid subnet, matched",
			opt:         ForbidSubnet("10.0.0.0/8"),
			host:        "http://10.0.0.1",
			expectedErr: ErrSubnetNotAllowed,
		}, {
			description: "invalid subnet",
			opt:         ForbidSubnet("10.0.0.0/"),
			failOnNew:   true,
			expectedErr: ErrInvalidInput,
		}, {
			description: "too many resolvers",
			opt:         ForbidSubnet("10.0.0.0/8", mockResolver, mockResolver),
			failOnNew:   true,
			expectedErr: ErrInvalidInput,
		}, {
			description: "forbid subnet with resolver, no match",
			opt:         ForbidSubnet("10.0.0.0/8", mockResolver),
			host:        mockPrivateURL,
		}, {
			description: "forbid subnet with resolver, resolver error",
			opt:         ForbidSubnet("10.0.0.0/8", mockResolver),
			host:        mockUnsupportedURL,
			expectedErr: errAny,
		}, {
			description: "forbid subnet with resolver, disallowed subnet",
			opt:         ForbidSubnet("192.168.1.0/8", mockResolver),
			host:        mockPrivateURL,
			expectedErr: ErrSubnetNotAllowed,
		},
	}
	testCommon(t, tests)
}

func TestForbidSubnetsOption(t *testing.T) {
	tests := []sharedTest{
		{
			description: "forbid subnet",
			opt:         ForbidSubnets([]string{"10.0.0.0/8", "10.0.0.0/24"}),
			host:        "http://192.168.1.1",
		},
	}
	testCommon(t, tests)
}

func TestForbidSubnetOptionString(t *testing.T) {
	opt := ForbidSubnet("10.0.0.0/8")
	assert.Equal(t, "ForbidSubnet('10.0.0.0/8')", opt.String())

	opt = ForbidSubnet("10.0.0.0/8", mockResolver)
	assert.Equal(t, "ForbidSubnet('10.0.0.0/8', resolver)", opt.String())

	opt = ForbidSubnets([]string{"10.0.0.0/8", "10.0.0.0/24"})
	assert.Equal(t, "ForbidSubnet('10.0.0.0/8', '10.0.0.0/24')", opt.String())

	opt = ForbidSubnets([]string{"10.0.0.0/8", "10.0.0.0/24"}, mockResolver)
	assert.Equal(t, "ForbidSubnet('10.0.0.0/8', '10.0.0.0/24', resolver)", opt.String())
}
