// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemeOption(t *testing.T) {
	tests := []sharedTest{
		{
			description: "just the scheme",
			opt:         OnlyAllowSchemes("http", "https"),
			noHttp:      true,
			hosts:       []string{"http://example.com", "https://example.com"},
		},
		{
			description: "ftp doesn't match http/s",
			opt:         OnlyAllowSchemes("ftp"),
			noHttp:      true,
			hosts:       []string{"http://example.com", "https://example.com"},
			expectedErr: ErrSchemeNotAllowed,
		},
		{
			description: "ftp doesn't match http/s",
			noHttp:      true,
			hosts:       []string{"ftp://example.com", "http://example.com", "https://example.com"},
		},
	}

	testCommon(t, tests)
}

func TestSchemeOptionString(t *testing.T) {
	opt := OnlyAllowSchemes("http")
	assert.Equal(t, "OnlyAllowSchemes('http')", opt.String())

	opt = OnlyAllowSchemes("http", "https")
	assert.Equal(t, "OnlyAllowSchemes('http', 'https')", opt.String())
}
