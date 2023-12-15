package tools

import(
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{

	"konstantinos": {
		AuthToken: "kon-token",
		Username: "konstantinos",
	},
	"john": {
		AuthToken: "john-token",
		Username: "john",
	},
	"peter": {
		AuthToken: "peter-token",
		Username: "peter",
	},
}


var mockBalanceDetails = map[string]BalanceDetails{

	"konstantinos": {
		Balance: 100,
		Username: "konstantinos",
	},
	"john": {
		Balance: 300,
		Username: "john",
	},
	"peter": {
		Balance: 900,
		Username: "peter",
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

func (d *mockDB) GetUserBalance (username string) *BalanceDetails {
	time.Sleep(time.Second * 1)

	var clientData = BalanceDetails{}
	clientData, ok := mockBalanceDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}