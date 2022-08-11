package rss

const reqRssMD = `### 摸鱼机器人提醒你

#### %s
%s
`

const reqRssContent = `
##### [%s](%s)
> %s
`

const emailBody = `
<div
    style="
      display: flex;
      flex-direction: column;
      justify-content: flex-start;
      align-items: center;
      align-content: flex-start;
      background-color: #f1f5f9;
      padding: 20px;
    "
  >
    <div style="align-content: center; width: 600px; background-color: #fff;">
       %s
       %s
    </div>
</div>`

const emailBodyTitle = `
<h2 style="margin: 25px;">%s</h2>
`

const emailBodyContent = `
 <div style="margin: 25px;">
	<p style="
		color: #999999;
		font-size: 12px;
		font-weight: 400;
		margin: 0;
		margin-bottom: 3px;
	">
		%s
	</p>
	<h3 style="font-weight: 400; margin: 0">
          <a
            href="%s"
            rel="noopener"
            target="_blank"
            style="text-decoration: none">
            %s
          </a>
 	</h3>
	<div>%s</div>
</div>
`

var receiver = []string{
	//"chenzuo@hotmail.com",
	"1120873075@qq.com",
}

var Configs = []RequestConfig{
	{Url: "https://sspai.com", IsDingTalk: true, IsEmail: true, Receiver: receiver},
	{Url: "https://www.appinn.com", IsDingTalk: true, IsEmail: true, Receiver: receiver},
	//{Url: "https://855.fun",IsDingTalk: false,IsEmail: false},
}
