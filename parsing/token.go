//
//  token.go
//  parsing
//
//  Created by d-exclaimation on 11:03 am.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package parsing

import "strings"

func IsDefinition(line string) bool {
	defSuffix := []string{
		"#include",
		"#define",
		"#typedef",
		"struct",
		"enum",
	}
	for _, suffix := range defSuffix {
		if strings.HasPrefix(line, suffix) {
			return true
		}
	}
	return false
}

func IsFunctionDeclaration(line string, nextLine string) bool {
	return strings.HasSuffix(line, ")") &&
		strings.Contains(line, "(") &&
		strings.HasSuffix(nextLine, "{")
}
