// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

// ForbidSpecialUseDomains returns an Option that disallows the use of special
// use domains.  See https://www.iana.org/assignments/special-use-domain-names/special-use-domain-names.xhtml
func ForbidSpecialUseDomains() Option {
	return forbidDomainNames("ForbidSpecialUseDomains",
		"*.alt",
		"*.example",
		"*.invalid",
		"*.local",
		"*.localhost",
		"*.test",
		"example.*",
	)
}

// ForbidSpecialUseSubnets returns an Option that disallows the use of special
// use domains that do not result in a network call, unless a resolver is
// provided. If a resolver is provided, it will be used to resolve the hostname
// and check the returned IP addresses for loopback addresses.
// See https://www.iana.org/assignments/iana-ipv4-special-registry/iana-ipv4-special-registry.xhtml
// and https://www.iana.org/assignments/iana-ipv6-special-registry/iana-ipv6-special-registry.xhtml
/*
func ForbidSpecialUseSubnets(resolver ...Resolver) Option {
	return forbidSubnetsOption("ForbidSpecialUseSubnets",
		[]string{
			"0.0.0.0/8",          //local ipv4
			"fe80::/10",          //local ipv6
			"255.255.255.255/32", //broadcast to neighbors
			"2001::/32",          //ipv6 TEREDO prefix
			"2001:5::/32",        //EID space for lisp
			"2002::/16",          //ipv6 6to4
			"fc00::/7",           //ipv6 unique local
			"192.0.0.0/24",       //ipv4 IANA
			"2001:0000::/23",     //ipv6 IANA
			"224.0.0.1/32",       //ipv4 multicast
		},
		resolver...,
	)
}
*/
