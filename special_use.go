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
