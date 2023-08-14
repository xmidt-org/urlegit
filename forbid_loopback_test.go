// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForbidLoopbackOption(t *testing.T) {
	loopback := []string{
		"http://127.0.0.1",
		"http://[::1]",
		"http://localhost",
	}

	tests := []sharedTest{
		{
			description: "forbid loopback, but host is not loopback",
			opt:         ForbidLoopback(),
			host:        "http://foo.com",
		}, {
			description: "forbid loopback, but ip is not loopback",
			opt:         ForbidLoopback(),
			host:        "http://192.168.1.1",
		}, {
			description: "forbid loopback",
			opt:         ForbidLoopback(),
			hosts:       loopback,
			expectedErr: ErrLoopback,
		}, {
			description: "forbid loopback with resolver",
			opts:        []Option{ForbidLoopback(), WithResolver(mockResolver)},
			hosts: []string{
				mockLoopbackURL,
				mockPrivateLoopbackURL,
			},
			expectedErr: ErrLoopback,
		}, {
			description: "forbid loopback with resolver, but host is not loopback",
			opts:        []Option{ForbidLoopback(), WithResolver(mockResolver)},
			host:        mockPrivateURL,
		}, {
			description: "forbid loopback with resolver, but resolver fails",
			opts:        []Option{ForbidLoopback(), WithResolver(mockResolver)},
			host:        mockUnsupportedURL,
			expectedErr: errAny,
		},
	}
	testCommon(t, tests)
}

func TestForbidLoopbackOptionString(t *testing.T) {
	opt := ForbidLoopback()
	assert.Equal(t, "ForbidLoopback()", opt.String())
}
