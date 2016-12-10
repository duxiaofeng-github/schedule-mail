# schedule-mail
schedule email sender

## Dependence
- Go v1.5 or above
- [github.com/jordan-wright/email](https://github.com/jordan-wright/email)

## How To Start
- `git clone git@github.com:duxiaofeng-github/schedule-mail.git`
- `cd schedule-mail`
- `git submodule update`
- `source env.sh`
- `make`
- for development `go run src/main.go --from 'someone <someone@gmail.com>' --to someone@gmail.com --cc someone@gmail.com --bcc someone@gmail.com --subject 'test' --content '<h1>Fancy Html is supported, too!</h1>' --smtpHost smtp.gmail.com --smtpPort 587 --username someone@gmail.com --password yourpwd --schedule "2016 12 10 17 * 00"`
- `./schedule-mail-macos --from 'someone <someone@gmail.com>' --to someone@gmail.com --cc someone@gmail.com --bcc someone@gmail.com --subject 'test' --content '<h1>Fancy Html is supported, too!</h1>' --smtpHost smtp.gmail.com --smtpPort 587 --username someone@gmail.com --password yourpwd --schedule "2016 12 10 17 * 00"`
- or `./schedule-mail-linux --from 'someone <someone@gmail.com>' --to someone@gmail.com --cc someone@gmail.com --bcc someone@gmail.com --subject 'test' --content '<h1>Fancy Html is supported, too!</h1>' --smtpHost smtp.gmail.com --smtpPort 587 --username someone@gmail.com --password yourpwd --schedule "2016 12 10 17 * 00"`
- have fun