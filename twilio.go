package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/loganballard/odds_briefing/config"
	logger "github.com/loganballard/odds_briefing/logger"
	odds "github.com/loganballard/odds_briefing/oddsapi"
)

func getTwilioInfo() (string, string, string, string) {
	var c config.Credentials
	c.LoadCredentials()
	return c.TwilioSid, c.TwilioAuthKey, c.TwilioNumberFrom, c.TwilioNumberTo
}

func setURLForMessaging(sid string) string {
	return "https://api.twilio.com/2010-04-01/Accounts/" + sid + "/Messages.json"
}

func setTwilioMessageBody(numberTo string, numberFrom string, msgBody string) *strings.Reader {
	msgData := url.Values{}

	msgData.Set("To", numberTo)
	msgData.Set("From", numberFrom)
	msgData.Set("Body", msgBody)
	msgDataReader := strings.NewReader(msgData.Encode())

	return msgDataReader
}

func buildTwilioRequest(msgBody string) http.Request {
	if len(msgBody) < 1 {
		logger.ErrorHelper(errors.New("msg to twilio was empty"))
	}

	sid, auth, numberFrom, numberTo := getTwilioInfo()
	urlStr := setURLForMessaging(sid)
	formattedMsgBody := setTwilioMessageBody(numberTo, numberFrom, msgBody)
	req, err := http.NewRequest("POST", urlStr, formattedMsgBody)
	if err != nil {
		logger.ErrorHelper(err)
	}

	req.SetBasicAuth(sid, auth)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return *req
}

func sendTwilioMsgFromGeneratedOddsData(gamblingMsg string) error {
	msgReq := buildTwilioRequest(gamblingMsg)
	client := &http.Client{}

	logger.Info(fmt.Sprintf("Sending Request to Twilio API: %v\n", msgReq))
	resp, err := client.Do(&msgReq)
	if err != nil {
		logger.ErrorHelper(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorHelper(err)
	}
	logger.Info(fmt.Sprintf("Twilio API Response: %v", string(respBody)))
	return nil
}

func sendFirstTotalsOddsAsMessage() error {
	nflTotalsForThisWeek := odds.GetNflTotalsOdds()
	msgToSend := nflTotalsForThisWeek[0]
	err := sendTwilioMsgFromGeneratedOddsData(msgToSend)
	return err
}

// PRETTY MUCH NEVER DO THIS
func sendFirstXTotalsOddsAsMessage(x int) error {
	var err error

	nflTotalsForThisWeek := odds.GetNflTotalsOdds()

	for i, msgToSend := range nflTotalsForThisWeek {
		if i >= (x - 1) {
			break
		}
		err = sendTwilioMsgFromGeneratedOddsData(msgToSend)
		if err != nil {
			logger.ErrorHelper(err)
		}
	}

	return err
}
