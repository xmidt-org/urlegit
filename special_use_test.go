// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"testing"
)

func TestForbidSpecialUseDomainsOption(t *testing.T) {
	tests := []sharedTest{
		{
			description: "invalid",
			opt:         ForbidSpecialUseDomains(),
			hosts: []string{
				"http://example.com",
				"http://example.org",
				"http://foo.alt",
				"http://foo.example",
				"http://foo.invalid",
				"http://foo.local",
				"http://foo.localhost",
				"http://foo.test",
			},
			expectedErr: ErrDomainNotAllowed,
		},
	}

	testCommon(t, tests)
}
