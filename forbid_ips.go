// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"net"
)

// ForbidAnyIPs returns an Option that disallows any IP addresses using only
// rules that do not result in a network call.  Note that this does not
// apply to hostnames that are resolved to an IP address when a resolver is
// provided.
func ForbidAnyIPs() Option {
	return forbidAnyIPsOption{}
}

type forbidAnyIPsOption struct{}

func (forbidAnyIPsOption) String() string {
	return "ForbidAnyIPs()"
}

func (forbidAnyIPsOption) apply(c *Checker) {
	c.ipBeforeRules = append(c.ipBeforeRules, forbidIPs)
}

func forbidIPs(*net.IP) error {
	return ErrIPNotAllowed
}
