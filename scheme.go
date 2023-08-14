// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import "strings"

// OnlyAllowSchemes returns an Option that allows any of the provided schemes.
// Multiple Schemes() options are allowed, but they are applied seprately.
//
// Scheme validation is case-insensitive.
func OnlyAllowSchemes(s ...string) Option {
	for i, v := range s {
		s[i] = strings.ToLower(v)
	}
	return schemesOption{schemes: s}
}

type schemesOption struct {
	schemes []string
}

func (s schemesOption) String() string {
	return "OnlyAllowSchemes('" + strings.Join(s.schemes, "', '") + "')"
}

func (s schemesOption) apply(c *Checker) {
	c.schemeRules = append(c.schemeRules, schemeChecker(s.schemes...))
}

func schemeChecker(schemes ...string) SchemeVador {
	return func(s string) error {
		for _, v := range schemes {
			if s == v {
				return nil
			}
		}
		return ErrSchemeNotAllowed
	}
}
