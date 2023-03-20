package exchange_api

type Ace struct {
	Uid         string
	ApiKey      string
	SecurityKey string
}

func NewAce(uid string, apiKey string, securityKey string) *Ace {
	return &Ace{
		Uid:         uid,
		ApiKey:      apiKey,
		SecurityKey: securityKey,
	}
}
