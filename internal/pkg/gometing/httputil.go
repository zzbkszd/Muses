package gometing

import (
	"fmt"
)

func CreateFormQuery(params map[string]string) string {
	ps := ""
	for k, v := range params {
		ps += fmt.Sprintf("&%s=%s", k, v)
	}
	return ps[1:]
}
