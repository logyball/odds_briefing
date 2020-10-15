### Sports Odds

Getting the Odds API information from [here](https://the-odds-api.com/).  The API Key is free for 500/requests a month.  

[Api Documentation](https://the-odds-api.com/liveapi/guides/v3/).

### Twilio API

Twilio is used for sending text messages.  [Check out their API Documentation](https://www.twilio.com/docs/usage/api).  This one isn't free.  It's $1/mo to register a phone number and $0.0075 per text.

### Setup

Copy [credentials.yml.example](./credentials/credentials.yml.example) and call it `credentials.yml`. Obtain API keys and put them in the required places in the credentials.yml file.

### Testing

To test:

`$> go test -v`

To test that the api calls are working:

`$> go test -v -tags=api_tests`

If you really want to run up your Twilio bill:

`$> TEST_HARMFUL_STUFF=true go test -v -tags=api_tests`
