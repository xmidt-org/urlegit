// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForbidAnyIPsOption(t *testing.T) {
	tests := []sharedTest{
		{
			description: "forbid any IPs, no match",
			opt:         ForbidAnyIPs(),
			host:        "http://example.com",
		}, {
			description: "forbid any IPs, match",
			opt:         ForbidAnyIPs(),
			host:        "http://192.168.1.1",
			expectedErr: ErrIPNotAllowed,
		},
	}
	testCommon(t, tests)
}

func TestForbidAnyIPsOptionString(t *testing.T) {
	opt := ForbidAnyIPs()
	assert.Equal(t, "ForbidAnyIPs()", opt.String())
}
