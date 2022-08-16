package service

import "time"

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

const reqHolidayMD = `### 摸鱼机器人提醒你

> 最后苦逼摸鱼的一天，明天就是: **%s** 的假期了！！！

`

const reqWorkMD = `### 摸鱼机器人提醒你

> 明天就要苦逼的摸鱼，今天是最后的疯狂！！！

`

var timezone, _ = time.LoadLocation("Asia/Shanghai")

const reqImageMD = `### 摸鱼机器人提醒你

一言

%s

![每日美图](%s)

`

const reqImageHtml = `
<body>
  <div
    style="
      display: flex;
      flex-direction: column;
      justify-content: flex-start;
      align-items: center;
      align-content: flex-start;
      background-color: #f1f5f9;
      padding: 20px;
      height: 100%s
    "
  >
    <div style="align-content: center; width: 45%s background-color: #fff">
      <h2 style="margin: 25px">一言</h2>
      <div style="margin: 25px">
		<p style="
			color: #999999;
			font-size: 16px;
			font-weight: 400;
		">
			%s
		</p>
        <image
          src="%s"
          alt="Bing Image"
          width="100%s
        />
      </div>
    </div>
  </div>
</body>
`
