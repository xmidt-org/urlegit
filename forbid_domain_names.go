// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"fmt"
	"strings"
)

// ForbidDomainNames returns an Option that disallows the provided domain names.
// This is a case-insensitive match, and where the domain name is mached against
// the hostname starting from the end.  The character '*' matches everything in
// the single subdomain it is used.
//
// Example: hostname "www.example.com." will match any of the following domains:
//
//   - "*."
//   - "*.*."
//   - "*.*.*."
//   - "com."
//   - "*.com."
//   - "example.*."
//   - "example.com."
//   - "www.example.com."
//   - "*.example.com."
//   - etc...
//
// But will not match:
//
//   - "foo.www.example.com."
//   - "*.www.example.com."
//   - etc...
func ForbidDomainNames(domains ...string) Option {
	return forbidDomainNames("ForbidDomainNames", domains...)
}

func forbidDomainNames(name string, domains ...string) Option {
	n := forbidDomainNamesOption{
		optName: name,
		domains: make([]*domainName, len(domains)),
	}
	for _, domain := range domains {
		d, err := newDomainName(domain)
		if err != nil {
			return Error(fmt.Errorf("%w: invalid domain '%s'", ErrInvalidInput, domain))
		}
		n.domains = append(n.domains, d)
	}
	return &n
}

type forbidDomainNamesOption struct {
	optName string
	domains []*domainName
}

func (n forbidDomainNamesOption) String() string {
	b := strings.Builder{}

	b.WriteString(n.optName)
	b.WriteString("(")
	comma := ""
	for _, d := range n.domains {

		if d != nil {
			b.WriteString(comma)
			b.WriteString("'")
			b.WriteString(d.original)
			b.WriteString("'")
			comma = ", "
		}
	}
	b.WriteString(")")

	return b.String()
}

func (n forbidDomainNamesOption) apply(c *Checker) {
	c.hostRules = append(c.hostRules, forbidDomainNamesHostname(n.domains))
}

func forbidDomainNamesHostname(forbid []*domainName) hostVador {
	return func(host string) error {
		subs, err := hostnameProcess(host)
		if err != nil {
			return err
		}

		for _, forbidden := range forbid {
			if forbidden == nil {
				continue
			}
			if forbidden.Match(subs) {
				return ErrDomainNotAllowed
			}
		}

		return nil
	}
}

type domainName struct {
	original string
	subs     []string
}

func newDomainName(s string) (*domainName, error) {
	subs, err := hostnameProcess(s)
	if err != nil {
		return nil, err
	}

	return &domainName{
		original: s,
		subs:     subs,
	}, nil
}

func hostnameProcess(s string) ([]string, error) {
	s = strings.TrimSuffix(strings.ToLower(s), ".")

	subs := strings.Split(s, ".")
	rv := make([]string, 0, len(subs))

	for _, sub := range subs {
		if sub == "" {
			return nil, fmt.Errorf("%w: invalid domain '%s' zero length subdomain", ErrInvalidInput, s)
		}

		// Reverse the list.
		rv = append([]string{sub}, rv...)
	}

	return rv, nil
}

func (d *domainName) Match(target []string) bool {
	// Check for a match in the block list.
	for i, block := range d.subs {
		if i >= len(target) || (block != "*" && target[i] != block) {
			// No match, try the next block.
			return false
		}
	}
	return true
}
