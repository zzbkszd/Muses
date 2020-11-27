package gometing

type RemoteMusicInfo struct {
	Id         string // 音乐id
	Name       string // 名称
	Artist     string // 艺术家
	ArtistId   string // 艺术家ID
	Album      string // 专辑名称
	AlbumId    string // 专辑ID
	SourceName string // 数据源名称
}

/**
加载一个请求session，返回一个header列表
*/
func LoadMockSession(key string) map[string]string {
	switch key {
	case "tencent":
		return map[string]string{
			"Referer":         "http://y.qq.com",
			"Cookie":          "pgv_pvi=22038528; pgv_si=s3156287488; pgv_pvid=5535248600; yplayer_open=1; ts_last=y.qq.com/portal/player.html; ts_uid=4847550686; yq_index=0; qqmusic_fromtag=66; player_exist=1",
			"User-Agent":      "QQ%E9%9F%B3%E4%B9%90/54409 CFNetwork/901.1 Darwin/17.6.0 (x86_64)",
			"Accept":          "*/*",
			"Accept-Language": "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4",
			"Connection":      "keep-alive",
			"Content-Type":    "application/x-www-form-urlencoded",
		}
	}
	return nil
}
