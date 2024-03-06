package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"testuser": {
		AuthToken: "testtoken",
		Username:  "testuser",
	},
	"testuser2": {
		AuthToken: "testtoken2",
		Username:  "testuser2",
	},
}
var mockCoinDetails = map[string]CoinDetails{
	"testuser": {
		Coins:    100,
		Username: "testuser",
	},
	"testuser2": {

		Coins:    200,
		Username: "testuser2",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate a database call
	time.Sleep(1 * time.Second)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (db *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate a database call
	time.Sleep(1 * time.Second)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}
