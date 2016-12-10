package main

import (
	"flag"
	"fmt"
	"log"
	"net/smtp"
	"regexp"
	"strings"
	"time"

	"github.com/email"
)

type Sender struct {
	SmtpHost string
	SmtpPort string
	Username string
	Password string
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func validateSchedule(str string) ([]string, bool) {
	scheduleValidateLayout := regexp.MustCompile(`^(\d{4}|\*) (\d{2}|\*) (\d{2}|\*) (\d{2}|\*) (\d{2}|\*) (\d{2}|\*)$`)
	submatch := scheduleValidateLayout.FindStringSubmatch(str)
	return submatch, len(submatch) == 7
}

func transformScheduleTextToRegexp(submatch []string) *regexp.Regexp {
	regexpStr := "^"

	for index, value := range submatch {
		if index == 0 {
			continue
		}

		if index == 1 {
			regexpStr += strings.Replace(value, "*", `\d{4}`, 1)
			continue
		}

		regexpStr += " " + strings.Replace(value, "*", `\d{2}`, 1)
	}

	regexpStr += "$"

	return regexp.MustCompile(regexpStr)
}

func startGoCron(schedule string, sender *Sender, emailIns *email.Email) {
	if submatch, isMatch := validateSchedule(schedule); !isMatch {
		log.Fatal("invalid schedule format")
	} else {
		scheduleRuntimeLayout := transformScheduleTextToRegexp(submatch)
		layout := "2006 01 02 15 04 05"

		for {
			now := time.Now()

			if scheduleRuntimeLayout.Match([]byte(now.Format(layout))) {
				go sendMail(now, sender, emailIns)
			}

			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func sendMail(now time.Time, sender *Sender, emailIns *email.Email) {
	smtpAuth := smtp.PlainAuth("", sender.Username, sender.Password, sender.SmtpHost)
	smtpAddr := fmt.Sprintf("%s:%s", sender.SmtpHost, sender.SmtpPort)

	if err := emailIns.Send(smtpAddr, smtpAuth); err != nil {
		fmt.Println("email send failed. err: ", err)
	} else {
		fmt.Printf("email send success. time: %s, from: %s, to: %+v, subject: %s. \n", now.Format("2006-01-02 15:04:05"), emailIns.From, emailIns.To, emailIns.Subject)
	}
}

func main() {
	scheduleStr := ""
	schedule := flag.String("schedule", "", "2016 * 01 00 00 00")
	from := flag.String("from", "", "Someone <someone@example.com>")
	to := arrayFlags{}
	bcc := arrayFlags{}
	cc := arrayFlags{}
	flag.Var(&to, "to", "email address send to others")
	flag.Var(&bcc, "bcc", "bcc email address send to others")
	flag.Var(&cc, "cc", "cc email address send to others")
	subject := flag.String("subject", "", "email subject")
	content := flag.String("content", "", "email content")
	smtpHost := flag.String("smtpHost", "", "smtp host. like: smtp.gmail.com")
	smtpPort := flag.String("smtpPort", "", "smtp port. like: 587")
	username := flag.String("username", "", "username. like: someone@example.com")
	password := flag.String("password", "", "password. like: 1234")

	flag.Parse()

	emailIns := email.NewEmail()
	sender := Sender{}

	if *schedule == "" {
		log.Fatal("schedule cannot be empty")
	} else {
		scheduleStr = *schedule
	}

	if *from == "" {
		log.Fatal("from cannot be empty")
	} else {
		emailIns.From = *from
	}

	if len(to) == 0 {
		log.Fatal("to cannot be empty")
	} else {
		emailIns.To = []string(to)
	}

	if len(bcc) != 0 {
		emailIns.Bcc = []string(bcc)
	}

	if len(cc) != 0 {
		emailIns.Cc = []string(cc)
	}

	if *subject != "" {
		emailIns.Subject = *subject
	}

	if *content != "" {
		emailIns.HTML = []byte(*content)
	}

	if *smtpHost == "" {
		log.Fatal("smtp host cannot be empty")
	} else {
		sender.SmtpHost = *smtpHost
	}

	if *smtpPort == "" {
		log.Fatal("smtp port cannot be empty")
	} else {
		sender.SmtpPort = *smtpPort
	}

	if *username == "" {
		log.Fatal("username cannot be empty")
	} else {
		sender.Username = *username
	}

	if *password == "" {
		log.Fatal("password cannot be empty")
	} else {
		sender.Password = *password
	}

	startGoCron(scheduleStr, &sender, emailIns)
}
