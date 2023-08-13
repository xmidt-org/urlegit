// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type sharedTest struct {
	description string
	opt         Option
	opts        []Option
	hosts       []string
	host        string
	noHttp      bool
	failOnNew   bool
	expectedErr error
}

func testCommon(t *testing.T, tests []sharedTest) {
	t.Helper()

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			opts := make([]Option, 0, len(tc.opts)+1)
			if !tc.noHttp {
				opts = append(opts, Schemes("http"))
			}
			opts = append(opts, tc.opt)
			opts = append(opts, tc.opts...)
			c, err := New(opts...)
			if tc.failOnNew {
				assert.ErrorIs(err, tc.expectedErr)
				return
			}
			require.NoError(err)

			hosts := tc.hosts
			if tc.host != "" {
				hosts = []string{tc.host}
			}

			for _, host := range hosts {
				err = c.Text(host)

				if errors.Is(tc.expectedErr, errAny) {
					assert.Error(err, "host: %s", host)
				} else {
					assert.ErrorIs(err, tc.expectedErr, "host: %s", host)
				}
			}
		})
	}
}
