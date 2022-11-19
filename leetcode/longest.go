package main

import (
	"fmt"
	"strings"
)

func longgesCommon(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[0 : len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func loggest() {
	flower := []string{"flower", "flow", "flight"}
	fmt.Println(longgesCommon(flower))
}
