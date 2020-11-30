package sources

import "github.com/zzbkszd/Muses/internal/pkg/gometing"

type MusicSource interface {
	// 根据关键字搜索音乐
	Search(keyword string) []*gometing.RemoteMusicInfo
	SearchPageable(keyword string, page, limit int) []*gometing.RemoteMusicInfo
	// 根据ID获取单个音乐
	Song(id string) *gometing.RemoteMusicInfo
	// 获取专辑信息
	Album(albumId string) []*gometing.RemoteMusicInfo
	// 获取艺术家信息
	Artist(artistId string) []*gometing.RemoteMusicInfo

	// 歌词
	Lyric(musicId string) string
	// 音乐下载连接
	FetchUrl(musicId string) string
}
