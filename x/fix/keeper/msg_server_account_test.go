package keeper_test

import "time"

func (suite *KeeperTestSuite) TestRegisterAccount() {
	type args struct {
		address          string
		companyName      string
		website          string
		socialMediaLinks string
		CreatedAt        string
	}

	type errArgs struct {
		shouldPass bool
		contains   string
	}

	tests := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{
			name: "Valid Address",
			args: args{
				address:          suite.address[0].String(),
				companyName:      "CypherCore",
				website:          "CypherCore.io",
				socialMediaLinks: "@CypherCore",
				CreatedAt:        time.Now().Add(time.Minute * 10).Format("2006-01-02 15:04:05"),
			},
			errArgs: errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{
			name: "InValid Address",
			args: args{
				address:          "",
				companyName:      "CypherCore",
				website:          "CypherCore.io",
				socialMediaLinks: "@CypherCore",
				CreatedAt:        time.Now().Add(time.Minute * 10).Format("2006-01-02 15:04:05"),
			},
			errArgs: errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
	}

	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			// Test logic goes here
		})
	}

}
