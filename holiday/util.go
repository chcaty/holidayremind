package holiday

import (
	"fmt"
	"holidayRemind/common"
	"holidayRemind/dingtalk"
)

const ReqHolidayMD = `### 摸鱼机器人提醒你

> 最后苦逼谜语的一天，明天就是: **%s** 的假期了！！！

`

const ReqWorkMD = `### 摸鱼机器人提醒你

> 明天就要苦逼的摸鱼，今天是最后的疯狂！！！

`

func CommonSendMessage(msg common.Message) (int, error) {
	err := dingtalk.SendMdMessage(msg)
	if err != nil {
		fmt.Printf("SendMdMessage to DingTalk error: %s", err.Error())
		return 500, err
	}
	fmt.Printf("SendMdMessage to DingTalk success: %s", msg.Text)
	return 200, nil
}
