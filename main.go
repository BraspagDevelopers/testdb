package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "testdb <server> <database>",
	Short: "Simple tool to check SQL database connectivity",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		cobra.CheckErr(err)
		password, err := cmd.Flags().GetString("password")
		cobra.CheckErr(err)

		dbOptions := DBOptions{
			Server:   args[0],
			Database: args[1],
			Username: coalesce(username, DefaultUsername),
			Password: coalesce(password, DefaultPassword),
		}

		db, err := dbOptions.connect()
		cobra.CheckErr(err)
		defer db.Close()

		tests := []Test{
			{"SELECT 1", `SELECT 1;`},
		}

		namelen := testNameMaxLen(tests)

		for _, test := range tests {
			result, err := RunTest(db, test)
			printResult(test, result, err, namelen)
		}
	},
}

func testNameMaxLen(tests []Test) int {
	var max int
	for _, t := range tests {
		length := len(t.Name)
		if length > max {
			max = length
		}
	}
	return max
}

func main() {
	cmd.Flags().StringP("username", "U", "", "Username")
	cmd.Flags().StringP("password", "P", "", "Password")
	cobra.CheckErr(cmd.Execute())
}
