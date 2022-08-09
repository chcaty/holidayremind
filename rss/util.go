package rss

const ReqRssMD = `### 摸鱼机器人提醒你

#### %s
%s
`

const ReqRssContent = `
##### [%s](%s)
> %s
`

type Channel int

const (
	Sspai Channel = iota
	Zztt
)
