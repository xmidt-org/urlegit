// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestForbidDomainNamesOption(t *testing.T) {
	tests := []sharedTest{
		{
			description: "invalid",
			opt:         ForbidDomainNames("foo.com"),
			host:        "http://foo.com",
			expectedErr: ErrDomainNotAllowed,
		}, {
			description: "invalid with trailing dot",
			opt:         ForbidDomainNames("foo.com."),
			host:        "http://foo.com",
			expectedErr: ErrDomainNotAllowed,
		}, {
			description: "invalid, missing subdomain",
			opt:         ForbidDomainNames("foo.com"),
			host:        "http://cars..foo.com",
			expectedErr: ErrInvalidInput,
		}, {
			description: "invalid option argument, empty subdomain",
			opt:         ForbidDomainNames("foo..com"),
			failOnNew:   true,
			expectedErr: ErrInvalidInput,
		}, {
			description: "happy path",
			opt:         ForbidDomainNames(),
			host:        "http://example.com",
		}, {
			description: "happy path",
			opt:         ForbidDomainNames("foo.com."),
			host:        "http://example.com",
		},
	}
	testCommon(t, tests)
}

func Test_hostnameProcess(t *testing.T) {
	tests := []struct {
		s           string
		want        []string
		expectedErr error
	}{
		{
			s:    "example.com",
			want: []string{"com", "example"},
		}, {
			s:    "Example.Com.",
			want: []string{"com", "example"},
		}, {
			s:           "example..com",
			expectedErr: ErrInvalidInput,
		},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			assert := assert.New(t)

			got, err := hostnameProcess(tc.s)

			assert.ErrorIs(err, tc.expectedErr)
			if tc.expectedErr == nil {
				assert.Equal(tc.want, got)
			}
		})
	}
}

func Test_newDomainName(t *testing.T) {
	tests := []struct {
		s           string
		want        domainName
		expectedErr error
	}{
		{
			s:    "example.com",
			want: domainName{original: "example.com", subs: []string{"com", "example"}},
		}, {
			s:           "example..com",
			expectedErr: ErrInvalidInput,
		},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			assert := assert.New(t)
			//require := require.New(t)

			got, err := newDomainName(tc.s)

			assert.ErrorIs(err, tc.expectedErr)
			if tc.expectedErr == nil {
				assert.Equal(&tc.want, got)
			}
		})
	}
}

func Test_domainName_Match(t *testing.T) {
	tests := []struct {
		input  string
		target string
		match  bool
	}{
		{
			input:  "example.com",
			target: "example.com",
			match:  true,
		}, {
			input:  "example.*",
			target: "example.com",
			match:  true,
		}, {
			input:  "*.*",
			target: "example.com",
			match:  true,
		}, {
			input:  "*",
			target: "example.com",
			match:  true,
		}, {
			input:  "specific.example.com",
			target: "example.com",
			match:  false,
		}, {
			input:  "*.example.com",
			target: "example.com",
			match:  false,
		}, {
			input:  "*.*.*",
			target: "example.com",
			match:  false,
		},
	}
	for _, tc := range tests {
		desc := tc.input + " does not match " + tc.target
		if tc.match {
			desc = tc.input + " matches " + tc.target
		}
		t.Run(desc, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			d, err := newDomainName(tc.input)
			require.NoError(err)

			target, err := hostnameProcess(tc.target)
			require.NoError(err)

			assert.Equal(tc.match, d.Match(target))
		})
	}
}

func TestForbidDomainNamesOptionString(t *testing.T) {
	opt := ForbidDomainNames()
	assert.Equal(t, "ForbidDomainNames()", opt.String())

	opt = ForbidDomainNames("foo.bar")
	assert.Equal(t, "ForbidDomainNames('foo.bar')", opt.String())
}
