package gometing

import (
	"testing"
)

func TestTencentMS_SearchPageable(t *testing.T) {
	te := TencentMS{}
	te.Search("小小")
}
