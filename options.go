// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"net"
)

// Resolver is a function that returns a list of IP addresses for a given
// host.
//
// To use the default go resolver, use the net.LookupIP function.
type Resolver func(host string) ([]net.IP, error)

// WithResolver returns an Option that will use the given Resolver to resolve
// hostnames into IP addresses.
func WithResolver(r Resolver) Option {
	return resolverOption{r: r}
}

type resolverOption struct {
	r Resolver
}

func (r resolverOption) String() string {
	if r.r == nil {
		return "WithResolver(nil)"
	}
	return "WithResolver(resolver)"
}

func (r resolverOption) apply(c *Checker) {
	c.resolver = r.r
}

// Error returns an option that will cause the Checker to return the given
// error.
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

// CustomSchemeVador returns an Option that will use the given SchemeVador
// to validate schemes.
func CustomSchemeVador(s SchemeVador) Option {
	return customSchemeVadorOption{s: s}
}

type customSchemeVadorOption struct {
	s SchemeVador
}

func (o customSchemeVadorOption) String() string {
	return "CustomSchemeVador(vador)"
}

func (o customSchemeVadorOption) apply(c *Checker) {
	c.schemeRules = append(c.schemeRules, o.s)
}

// CustomHostVador returns an Option that will use the given HostVador
// to validate hosts.
func CustomHostVador(h HostVador) Option {
	return customHostVadorOption{h: h}
}

type customHostVadorOption struct {
	h HostVador
}

func (o customHostVadorOption) String() string {
	return "CustomHostVador(vador)"
}

func (o customHostVadorOption) apply(c *Checker) {
	c.hostRules = append(c.hostRules, o.h)
}

// CustomIPVador returns an Option that will use the given IPVador
// to validate IPs.
func CustomIPVador(i IPVador) Option {
	return customIPVadorOption{i: i}
}

type customIPVadorOption struct {
	i IPVador
}

func (o customIPVadorOption) String() string {
	return "CustomIPVador(vador)"
}

func (o customIPVadorOption) apply(c *Checker) {
	c.ipRules = append(c.ipRules, o.i)
}
