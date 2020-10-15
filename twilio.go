package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getTwilioInfo() (string, string, string, string) {
	var c credentials
	c.loadCredentials()
	return c.TwilioSid, c.TwilioAuthKey, c.TwilioNumberFrom, c.TwilioNumberTo
}

func setUrlForMessaging(sid string) string {
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
		ErrorHelper(errors.New("msg to twilio was empty!"))
	}

	sid, auth, numberFrom, numberTo := getTwilioInfo()
	urlStr := setUrlForMessaging(sid)
	formattedMsgBody := setTwilioMessageBody(numberTo, numberFrom, msgBody)
	req, err := http.NewRequest("POST", urlStr, formattedMsgBody)
	if err != nil {
		ErrorHelper(err)
	}

	req.SetBasicAuth(sid, auth)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return *req
}

func sendTwilioMsgFromGeneratedOddsData(gamblingMsg string) error {
	msgReq := buildTwilioRequest(gamblingMsg)
	client := &http.Client{}

	InfoLogger.Printf("Sending Request to Twilio API: %v\n", msgReq)
	resp, err := client.Do(&msgReq)
	if err != nil {
		ErrorHelper(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ErrorHelper(err)
	}
	InfoLogger.Printf("Twilio API Response: %v", string(respBody))
	return nil
}

func SendFirstTotalsOddsAsMessage() error {
	nflTotalsForThisWeek := GetNflTotalsOdds()
	msgToSend := nflTotalsForThisWeek[0]
	err := sendTwilioMsgFromGeneratedOddsData(msgToSend)
	return err
}

// PRETTY MUCH NEVER DO THIS
func SendFirstXTotalsOddsAsMessage(x int) error {
	var err error

	nflTotalsForThisWeek := GetNflTotalsOdds()

	for i, msgToSend := range nflTotalsForThisWeek {
		if i >= (x - 1) {
			break
		}
		err = sendTwilioMsgFromGeneratedOddsData(msgToSend)
		if err != nil {
			ErrorHelper(err)
		}
	}

	return err
}
