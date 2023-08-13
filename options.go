// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"net"
)

// Resolver is a function that returns a list of IP addresses for a given
// host.
type Resolver func(host string) ([]net.IP, error)

func Error(err error) Option {
	return errorOption{err: err}
}

type errorOption struct {
	err error
}

func (o errorOption) String() string {
	if o.err == nil {
		return "Error()"
	}
	return "Error('" + o.err.Error() + "')"
}

func (o errorOption) apply(c *Checker) {
	c.err = o.err
}
