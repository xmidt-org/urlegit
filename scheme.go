// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import "strings"

// Schemes returns an Option that allows any of the provided schemes.
// Multiple Schemes() options are allowed and the schemes are combined.
// If an OnlyScheme() option is provided, Schemes() options are ignored.
//
// Scheme validation is case-insensitive.
func Schemes(s ...string) Option {
	for i, v := range s {
		s[i] = strings.ToLower(v)
	}
	return schemesOption{schemes: s}
}

type schemesOption struct {
	schemes []string
}

func (s schemesOption) String() string {
	return "Schemes('" + strings.Join(s.schemes, "', '") + "')"
}

func (s schemesOption) apply(c *Checker) {
	if !c.onlyOneScheme {
		c.allowedSchemes = append(c.allowedSchemes, s.schemes...)
	}
}

//------------------------------------------------------------------------------

// OnlyScheme returns an Option that allows only the provided scheme.  The last
// OnlyScheme option wins.  Schemes() options are ignored if an OnlyScheme
// option is provided.
//
// Scheme validation is case-insensitive.
func OnlyScheme(s string) Option {
	return onlySchemeOption(strings.ToLower(s))
}

type onlySchemeOption string

func (o onlySchemeOption) String() string {
	return "OnlyScheme('" + string(o) + "')"
}

func (o onlySchemeOption) apply(c *Checker) {
	c.onlyOneScheme = true
	c.allowedSchemes = []string{string(o)}
}
