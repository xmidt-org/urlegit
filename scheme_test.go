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
			opt:         Schemes("http", "https"),
			noHttp:      true,
			hosts:       []string{"http://foo.com", "https://foo.com"},
		},
		{
			description: "add schemes multiple times",
			opts: []Option{
				Schemes("https"),
				Schemes("http"),
			},
			noHttp: true,
			hosts:  []string{"http://foo.com", "https://foo.com"},
		},
		{
			description: "ftp doesn't match http/s",
			opt:         Schemes("ftp"),
			noHttp:      true,
			hosts:       []string{"http://foo.com", "https://foo.com"},
			expectedErr: ErrSchemeNotAllowed,
		},
		{
			description: "only ftp, but try other options",
			opts: []Option{
				OnlyScheme("ftp"),
				Schemes("http", "https"),
			},
			noHttp:      true,
			hosts:       []string{"http://foo.com", "https://foo.com"},
			expectedErr: ErrSchemeNotAllowed,
		},
		{
			description: "only ftp, but try other options, alt order",
			opts: []Option{
				Schemes("http", "https"),
				OnlyScheme("ftp"),
			},
			noHttp:      true,
			hosts:       []string{"http://foo.com", "https://foo.com"},
			expectedErr: ErrSchemeNotAllowed,
		},
		{
			description: "only ftp, but try other options",
			opts: []Option{
				OnlyScheme("ftp"),
				Schemes("http", "https"),
			},
			noHttp: true,
			host:   "ftp://foo.com",
		},
		{
			description: "only ftp, but try other options, alt order",
			opts: []Option{
				Schemes("http", "https"),
				OnlyScheme("ftp"),
			},
			noHttp: true,
			host:   "ftp://foo.com",
		},
	}

	testCommon(t, tests)
}

func TestSchemeOptionString(t *testing.T) {
	opt := Schemes("http")
	assert.Equal(t, "Schemes('http')", opt.String())

	opt = Schemes("http", "https")
	assert.Equal(t, "Schemes('http', 'https')", opt.String())

	opt = OnlyScheme("ftp")
	assert.Equal(t, "OnlyScheme('ftp')", opt.String())
}
