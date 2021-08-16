package handy

import "errors"

var ErrAssertion = errors.New("assertion")

func Assert(value bool, WhatWillPanic interface{}) {
	if !value {
		if WhatWillPanic == nil {
			panic(ErrAssertion)
		}
		panic(WhatWillPanic)
	}
}
