// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

var (
	ErrSchemeNotAllowed     = fmt.Errorf("scheme not allowed")
	ErrHostnameEmpty        = fmt.Errorf("hostname is empty")
	ErrLoopback             = fmt.Errorf("loopback address")
	ErrDomainNotAllowed     = fmt.Errorf("domain not allowed")
	ErrRootDomainNotAllowed = fmt.Errorf("root domain not allowed")
	ErrInvalidInput         = fmt.Errorf("invalid input")
	ErrSubnetNotAllowed     = fmt.Errorf("subnet not allowed")
	ErrIPNotAllowed         = fmt.Errorf("IPs not allowed")
)

// Checker is a URL validator.
type Checker struct {
	schemeRules   []SchemeVador
	ipBeforeRules []IPVador
	resolver      Resolver
	hostRules     []HostVador
	ipRules       []IPVador
	err           error
	opts          []Option
}

// Option is an option for a Checker.
type Option interface {
	fmt.Stringer
	apply(*Checker)
}

// SchemeVador is a function that validates a scheme.
type SchemeVador func(string) error

// IPVador is a function that validates an IP.
type IPVador func(*net.IP) error

// HostVador is a function that validates a host.
type HostVador func(string) error

// New returns a new Checker with the provided options applied.
func New(opts ...Option) (*Checker, error) {
	c := Checker{
		opts: make([]Option, 0, len(opts)),
	}

	for _, opt := range opts {
		if opt != nil {
			opt.apply(&c)
			c.opts = append(c.opts, opt)
		}
	}

	if c.err != nil {
		return nil, c.err
	}
	return &c, nil
}

// Must returns a new Checker with the provided options applied. If an error
// occurs, it panics.
func Must(opts ...Option) *Checker {
	c, err := New(opts...)
	if err != nil {
		panic(err)
	}
	return c
}

// Legit returns true if the provided string is a valid URL based on the
// provided options.
func (c *Checker) Legit(s string) bool {
	return c.Text(s) == nil
}

// URLLegit returns true if the provided URL is valid based on the provided
// options.
func (c *Checker) URLegit(u *url.URL) bool {
	return c.URL(u) == nil
}

// Text returns an error if the provided string is not a valid URL based on
// the provided options.
func (c *Checker) Text(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	return c.URL(u)
}

// URLDetails returns an error if the provided URL is not valid based on
// the provided options.
func (c *Checker) URL(u *url.URL) error {
	if u == nil {
		return ErrInvalidInput
	}

	scheme := strings.ToLower(u.Scheme)
	for _, rule := range c.schemeRules {
		err := rule(scheme)
		if err != nil {
			return err
		}
	}

	host := strings.ToLower(u.Hostname())
	if host == "" {
		return ErrHostnameEmpty
	}

	var ips []net.IP
	ip := net.ParseIP(host)

	if ip != nil {
		ips = []net.IP{ip}
		for _, rule := range c.ipBeforeRules {
			err := rule(&ip)
			if err != nil {
				return err
			}
		}
	} else {
		for _, rule := range c.hostRules {
			err := rule(host)
			if err != nil {
				return err
			}
		}

		if c.resolver != nil {
			var err error
			// Replace the IPs with the newly resolved IPs.
			ips, err = c.resolver(host)
			if err != nil {
				return err
			}
		}
	}

	if len(ips) == 0 {
		return nil
	}

	for _, rule := range c.ipRules {
		for _, ip := range ips {
			err := rule(&ip)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Checker) String() string {
	inner := strings.Builder{}

	comma := ""
	for _, opt := range c.opts {
		inner.WriteString(comma)
		inner.WriteString(opt.String())
		comma = ", "
	}

	buf := strings.Builder{}
	buf.WriteString("urlegit.Checker{")
	innerStr := inner.String()
	if innerStr != "" {
		buf.WriteString(" ")
		buf.WriteString(innerStr)
		buf.WriteString(" ")
	}
	buf.WriteString("}")
	return buf.String()
}
