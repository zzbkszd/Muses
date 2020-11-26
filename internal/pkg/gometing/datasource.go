package gometing

type MusicSource interface {
	// 根据关键字搜索音乐
	Search(keyword string) []*RemoteMusicInfo
	SearchPageable(keyword string, page, limit int) []*RemoteMusicInfo
	// 根据ID获取单个音乐
	Song(id string) *RemoteMusicInfo
	// 获取专辑信息
	Album(albumId string) []*RemoteMusicInfo
	// 获取艺术家信息
	Artist(artistId string) []*RemoteMusicInfo

	// 歌词
	Lyric(music *RemoteMusicInfo)
	// 专辑图片
	Picture(music *RemoteMusicInfo)
	// 音乐下载连接
	FetchUrl(music *RemoteMusicInfo)
}

// QQ音乐
type TencentMS struct {
}

func (TencentMS) SearchPageable(keyword string, page, limit int) {
	url := "https://c.y.qq.com/soso/fcgi-bin/client_search_cp"
	params := map[string]string{
		"format": "json",
	}

}
