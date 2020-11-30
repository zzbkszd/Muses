package sources

import (
	"fmt"
	"testing"
)

func TestTencentMS_SearchPageable(t *testing.T) {
	te := TencentMS{}
	res := te.Search("小小")
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestTencentMS_Album(t *testing.T) {
	te := TencentMS{}
	res := te.Album("0045sZAv2PGNvH")
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestTencentMS_Artist(t *testing.T) {
	te := TencentMS{}
	res := te.Artist("001uXFgt1kpLyI")
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestTencentMS_Lyric(t *testing.T) {
	te := TencentMS{}
	lyric := te.Lyric("003Inmwn25vHfD")
	fmt.Println(lyric)
}

func TestTencentMS_URL(t *testing.T) {
	te := TencentMS{}
	lyric := te.FetchUrl("003Inmwn25vHfD")
	fmt.Println(lyric)
}
