package utils

import (
	"fmt"
	"strings"
)

func ParseUrl(host, url string, env map[string]string) string {
	for k, v := range env {
		url = strings.ReplaceAll(url, fmt.Sprintf("{{%v}}", k), v)
	}
	return fmt.Sprintf("%v%v", host, url)
}
