package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234BC",
		Username:  "alex",
	},
	"alexa": {
		AuthToken: "1234BC",
		Username:  "alexa",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    200,
		Username: "alex",
	},
	"alexa": {
		Coins:    400,
		Username: "alexa",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}

	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
