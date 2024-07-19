package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Khel": {
		Username:   "Khel",
		Auth_token: "109GHI",
	},
	"Emz": {
		Username:   "Emz",
		Auth_token: "KE103",
	},
	"Azrael": {
		Username:   "Azrael",
		Auth_token: "314KEH",
	},
}

var mockMoneyDetails = map[string]MoneyDetails{
	"Khel": {
		Username: "Khel",
		Money:    31,
	},
	"Emz": {
		Username: "Emz",
		Money:    41,
	},
	"Azrael": {
		Username: "Azrael",
		Money:    91,
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserMoney(username string) *MoneyDetails {
	time.Sleep(1 * time.Second)

	var clientData = MoneyDetails{}

	clientData, ok := mockMoneyDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
