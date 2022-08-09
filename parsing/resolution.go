//
//  resolution.go
//  parsing
//
//  Created by d-exclaimation on 2:58 pm.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package parsing

import (
	"errors"
	"strings"
)

func module(filename string) (string, error) {
	module, _, found := strings.Cut(filename, ".c")
	if found {
		return module, nil
	}
	module, _, found = strings.Cut(filename, ".h")
	if found {
		return module, nil
	}
	module, _, found = strings.Cut(filename, ".cscript")
	if found {
		return module, nil
	}
	return module, errors.New("not a detectable C source file")
}
