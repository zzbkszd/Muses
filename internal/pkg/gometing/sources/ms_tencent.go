package sources

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/zzbkszd/Muses/internal/pkg/common"
	"github.com/zzbkszd/Muses/internal/pkg/gometing"
	"math/rand"
	"strconv"
)

// QQ音乐
type TencentMS struct {
}

func (ms TencentMS) Search(keyword string) []*gometing.RemoteMusicInfo {
	return ms.SearchPageable(keyword, 1, 10)
}
func (TencentMS) SearchPageable(keyword string, page, limit int) []*gometing.RemoteMusicInfo {
	url := "https://c.y.qq.com/soso/fcgi-bin/client_search_cp"
	params := map[string]string{
		"format":   "json",
		"p":        strconv.Itoa(page),
		"n":        strconv.Itoa(limit),
		"w":        keyword,
		"aggr":     "1",
		"lossless": "1",
		"cr":       "1",
		"new_json": "1",
	}

	response, err := common.DoGet(url, params, gometing.LoadMockSession("tencent"))
	if err != nil {
		return nil
	}

	return parseSongs(gjson.Get(response, "data.song.list"))
}

func (TencentMS) Song(id string) *gometing.RemoteMusicInfo {
	url := "https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg"
	params := map[string]string{
		"format":   "json",
		"platform": "yqq",
		"songmid":  id,
	}

	response, err := common.DoGet(url, params, gometing.LoadMockSession("tencent"))
	if err != nil {
		return nil
	}

	return parseSong(gjson.Get(response, "data.0"))
}

func (TencentMS) Album(albumId string) []*gometing.RemoteMusicInfo {
	url := "https://c.y.qq.com/v8/fcg-bin/fcg_v8_album_detail_cp.fcg"
	params := map[string]string{
		"format":   "json",
		"platform": "mac",
		"albummid": albumId,
		"newsong":  "1",
	}

	response, err := common.DoGet(url, params, gometing.LoadMockSession("tencent"))
	if err != nil {
		return nil
	}
	return parseSongs(gjson.Get(response, "data.getSongInfo"))
}

func (TencentMS) Artist(artistId string) []*gometing.RemoteMusicInfo {
	url := "https://c.y.qq.com/v8/fcg-bin/fcg_v8_singer_track_cp.fcg"
	params := map[string]string{
		"format":    "json",
		"platform":  "mac",
		"singermid": artistId,
		"newsong":   "1",
		"begin":     "0",
		"num":       "30",
		"order":     "listen",
	}

	response, err := common.DoGet(url, params, gometing.LoadMockSession("tencent"))
	if err != nil {
		return nil
	}

	return parseSongs(gjson.Get(response, "data.list.#.musicData"))
}

func (TencentMS) Lyric(musicId string) string {
	url := "https://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric_new.fcg"
	params := map[string]string{
		"songmid": musicId,
		"g_tk":    "5381",
	}

	response, err := common.DoGet(url, params, gometing.LoadMockSession("tencent"))
	if err != nil {
		return ""
	}
	//
	//return parseSongs(gjson.Get(response, "data.list.#.musicData"))
	fmt.Println(response)
	json := response[len("MusicJsonCallback(") : len(response)-1]
	lyric := gjson.Get(json, "lyric").String()
	decode, err := base64.StdEncoding.DecodeString(lyric)
	if err != nil {
		return ""
	}
	return string(decode)
}

func (TencentMS) FetchUrl(musicId string) string {
	url := "https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg"
	infoParam := map[string]string{
		"format":   "json",
		"platform": "yqq",
		"songmid":  musicId,
	}

	info, err := common.DoGet(url, infoParam, gometing.LoadMockSession("tencent"))
	if err != nil {
		return ""
	}

	types := [][]string{
		{"size_320mp3", "320", "M800", "mp3"},
		{"size_128mp3", "128", "M500", "mp3"},
		{"size_192aac", "192", "C600", "m4a"},
	}

	songmids := make([]string, 0)
	filename := make([]string, 0)
	songtype := make([]int64, 0)
	for _, vo := range types {
		songmids = append(songmids, gjson.Get(info, "data.0.mid").String())
		filename = append(filename, vo[2]+gjson.Get(info, "data.0.file.media_mid").String()+"."+vo[3])
		songtype = append(songtype, gjson.Get(info, "data.0.type").Int())
	}
	params := []map[string]interface{}{
		{
			"guid":      strconv.Itoa(int(rand.Int63n(10000000000))),
			"songmid":   songmids,
			"filename":  filename,
			"songtype":  songtype,
			"uin":       0,
			"loginflag": 1,
			"platform":  20,
		},
	}

	payload, err := json.Marshal([]map[string]interface{}{{"req_0": params}})
	if err != nil {
		return ""
	}
	fetchurl := "https://u.y.qq.com/cgi-bin/musicu.fcg"
	fetchParam := map[string]string{
		"format":      "json",
		"platform":    "yqq.json",
		"needNewCode": "0",
		"data":        string(payload),
	}

	fmt.Println(string(payload))
	fetchResponse, err := common.DoGet(fetchurl, fetchParam, gometing.LoadMockSession("tencent"))
	fmt.Println(fetchResponse)
	return ""

}

func parseSongs(songs gjson.Result) []*gometing.RemoteMusicInfo {
	if songs.IsArray() {
		res := make([]*gometing.RemoteMusicInfo, 0)
		for _, song := range songs.Array() {
			res = append(res, parseSong(song))
		}
		return res
	} else {
		return nil
	}
}

func parseSong(song gjson.Result) *gometing.RemoteMusicInfo {
	format := "https://y.gtimg.cn/music/photo_new/T002R300x300M000%s.jpg?max_age=2592000"
	return &gometing.RemoteMusicInfo{
		Id:         song.Get("mid").String(),
		Name:       song.Get("name").String(),
		Artist:     song.Get("singer.0.name").String(),
		ArtistId:   song.Get("singer.0.mid").String(),
		Album:      song.Get("album.name").String(),
		AlbumId:    song.Get("album.mid").String(),
		Picture:    fmt.Sprintf(format, song.Get("mid").String()),
		SourceName: "tencent",
	}
}
