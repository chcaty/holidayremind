package smtp

type SmtpConfig struct {
	Host      string
	Port      string
	UserName  string
	Password  string
	MaxClient int
}

var Smtp = SmtpConfig{
	Host:      "smtp.163.com",
	Port:      "25",
	UserName:  "chcaty@163.com",
	Password:  "LVAULLJARBXIKAAC",
	MaxClient: 5,
}

var Receiver = []string{
	//"chenzuo@hotmail.com",
	"1120873075@qq.com",
}

const EmailBody = `
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

const EmailBodyTitle = `
<h2 style="margin: 25px;">%s</h2>
`

const EmailBodyContent = `
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
