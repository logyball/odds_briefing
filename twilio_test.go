package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTwilioKeys(t *testing.T) {
	sid, auth, from, to := getTwilioInfo()
	assert.NotEmpty(t, sid)
	assert.NotEmpty(t, auth)
	assert.NotEmpty(t, from)
	assert.NotEmpty(t, to)
}

func TestBuildTwilioReq(t *testing.T) {
	msgReq := buildTwilioRequest("message")
	_, _, numberFrom, numberTo := getTwilioInfo()
	_, _, basicAuthIsUsed := msgReq.BasicAuth()
	msgBodyArr, err := ioutil.ReadAll(msgReq.Body)
	if err != nil {
		assert.Fail(t, "getting message body failed")
	}
	msgBody := string(msgBodyArr)

	assert.NotEmpty(t, msgReq)
	assert.Contains(t, msgReq.Header.Values("Content-Type"), "application/x-www-form-urlencoded")
	assert.Contains(t, msgReq.Header.Values("Accept"), "application/json")
	assert.True(t, basicAuthIsUsed)
	assert.Contains(t, msgBody, "message")
	assert.Contains(t, msgBody, numberTo[1:])   // strip URL-encoded "+"
	assert.Contains(t, msgBody, numberFrom[1:]) // strip URL-encoded "+"
}
