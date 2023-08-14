// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	errAny = errors.New("any error")
)

func TestNew(t *testing.T) {
	tests := []struct {
		description string
		opts        []Option
		expectedErr error
	}{
		{
			description: "simple case",
		},
		{
			description: "error case",
			opts:        []Option{Error(ErrSchemeNotAllowed)},
			expectedErr: ErrSchemeNotAllowed,
		},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			assert := assert.New(t)

			got, err := New(tc.opts...)

			assert.ErrorIs(err, tc.expectedErr)
			if tc.expectedErr == nil {
				assert.NotNil(got)
			}
		})
	}
}

func TestMust(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		Must(Error(ErrSchemeNotAllowed))
	})

	c := Must()
	assert.NotNil(c)
}

func TestText(t *testing.T) {
	tests := []sharedTest{
		{
			description: "no hostname",
			host:        "http://",
			expectedErr: ErrHostnameEmpty,
		}, {
			description: "url.Parse error",
			host:        ":invalid",
			expectedErr: errAny,
		},
	}
	testCommon(t, tests)
}

func TestURL(t *testing.T) {
	assert := assert.New(t)

	c := Must()
	err := c.URL(nil)
	assert.ErrorIs(err, ErrInvalidInput)
}

func TestLegit(t *testing.T) {
	c := Must(OnlyAllowSchemes("http"))
	assert.Equal(t, true, c.Legit("http://example.com"))
	assert.Equal(t, false, c.Legit("https://example.com"))
}

func TestURLegit(t *testing.T) {
	good, _ := url.Parse("http://example.com")
	bad, _ := url.Parse("https://example.com")

	c := Must(OnlyAllowSchemes("http"))

	assert.Equal(t, true, c.URLegit(good))
	assert.Equal(t, false, c.URLegit(bad))
}

func TestCheckerString(t *testing.T) {
	c := Must()
	assert.Equal(t, "urlegit.Checker{}", c.String())

	c = Must(OnlyAllowSchemes("http"))
	assert.Equal(t, "urlegit.Checker{ OnlyAllowSchemes('http') }", c.String())

	c = Must(OnlyAllowSchemes("http", "https"))
	assert.Equal(t, "urlegit.Checker{ OnlyAllowSchemes('http', 'https') }", c.String())
}

func TestErrorString(t *testing.T) {
	assert.Equal(t, "Error()", Error(nil).String())
	assert.Equal(t, "Error('any error')", Error(errAny).String())
}
