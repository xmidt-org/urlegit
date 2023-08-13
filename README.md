# urlegit

URLegit is a library for validating URLs.


[![Build Status](https://github.com/xmidt-org/urlegit/actions/workflows/ci.yml/badge.svg)](https://github.com/xmidt-org/urlegit/actions/workflows/ci.yml)
[![codecov.io](https://codecov.io/github/xmidt-org/urlegit/coverage.svg?branch=main)](http://codecov.io/github/xmidt-org/urlegit?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/xmidt-org/urlegit)](https://goreportcard.com/report/github.com/xmidt-org/urlegit)
[![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/xmidt-org/urlegit/blob/main/LICENSE)
[![GitHub Release](https://img.shields.io/github/release/xmidt-org/urlegit.svg)](CHANGELOG.md)
[![GoDoc](https://pkg.go.dev/badge/github.com/xmidt-org/urlegit)](https://pkg.go.dev/github.com/xmidt-org/urlegit)

## Summary

URL validation is complex and very specific to the task at hand.  This library's
goal is to provide a toolkit for making pre-flight checks against a URL easier.

## Usage

```go
package main

import (
	"fmt"

	"github.com/xmidt-org/urlegit"
)

func main() {
	c, err := urlegit.New(
		urlegit.OnlyScheme("https"),
		urlegit.ForbidSpecialUseDomains(),
	)

	if err != nil {
		panic(err)
	}

	url := "https://github.com"
	fmt.Printf("Is %q allowed? %t\n", url, c.Legit(url))
    
	// Output:
	// Is "https://github.com" allowed? true
}
```
[Go Playground](https://go.dev/play/p/QE93GEm6vrU)

## Resources

- https://www.w3.org/Addressing/URL/5_BNF.html
- https://datatracker.ietf.org/doc/html/rfc1738
- https://www.iana.org/assignments/special-use-domain-names/special-use-domain-names.xhtml
- https://www.iana.org/assignments/iana-ipv4-special-registry/iana-ipv4-special-registry.xhtml
- https://www.iana.org/assignments/iana-ipv6-special-registry/iana-ipv6-special-registry.xhtml

## Footnotes

This library originally started of in [xmidt-org/ancla](https://github.com/xmidt-org/ancla)
as webhook URL validation code.  [Here](https://github.com/xmidt-org/ancla/blob/09a683a1ca368cfc020eaef9345f3ebd7b79e825/webhookValidationConfig.go#L27) is a perma
link to that code.