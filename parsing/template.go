//
//  template.go
//  parsing
//
//  Created by d-exclaimation on 2:31 pm.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package parsing

import (
	"fmt"
	"strings"
	"time"
)

func (p *Parser) sourceContent() string {
	before := []string{
		"//",
		fmt.Sprintf("//  %s.c", p.module),
		"//",
		fmt.Sprintf("//  Created by d-exclaimation on %s.", time.Now().Format("01-02-2006 15:04:05")),
		"//",
		"",
		fmt.Sprintf("#include \"%s.h\"", p.module),
		"",
	}
	return strings.Join(append(before, p.sources...), "\n")
}

func (p *Parser) headerContent() string {
	before := []string{
		"//",
		fmt.Sprintf("//  %s.c", p.module),
		"//",
		fmt.Sprintf("//  Created by d-exclaimation on %s.", time.Now().Format("01-02-2006 15:04:05")),
		"//",
		"",
		fmt.Sprintf("#ifndef %s_H", strings.ToUpper(p.module)),
		fmt.Sprintf("#define %s_H", strings.ToUpper(p.module)),
		"",
	}
	curr := append(before, p.headers...)
	return strings.Join(append(curr, "", "#endif"), "\n")
}
