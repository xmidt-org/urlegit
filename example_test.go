// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit_test

import (
	"fmt"

	"github.com/xmidt-org/urlegit"
)

func Example() {
	c, err := urlegit.New(
		urlegit.OnlyAllowSchemes("https"),
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
