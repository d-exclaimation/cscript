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
	prefixes := []string{
		"#include",
		"#define",
		"#typedef",
		"/**",
		" *",
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return false
}

func IsFunctionDeclaration(line string, nextLine string) bool {
	clean := strings.TrimSpace(line)
	return strings.HasSuffix(clean, ")") &&
		strings.Contains(clean, "(") &&
		strings.HasSuffix(strings.TrimSpace(nextLine), "{")
}

func IsStaticDeclaration(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "static")
}
