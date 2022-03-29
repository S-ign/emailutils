package email

import (
	"fmt"
	"os"
	"testing"

	"github.com/s-ign/emailutils/config"
)

func TestCreateAuth(t *testing.T) {
	auth1, err := config.CreateAuth(
		"avery.carty@mojodomo.com",
		"MOJO",
	)
	if err != nil {
		return
	}
	_, err2 := config.CreateAuth(
		"avery.carty@mojodomo.com",
		"MAJO",
	)

	_, err3 := config.CreateAuth(
		"avery.carty@mojodomo.com",
		"MIJO",
	)

	cases := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "works if MOJO envirnoment variable set",
			got:  fmt.Sprintf("%v", auth1),
			want: fmt.Sprintf("%v", &config.Auth{
				From: "avery.carty@mojodomo.com", Password: os.Getenv("MOJO")}),
		},
		{
			name: "return error if environment variable is unset",
			got:  fmt.Sprintf("%v", err2),
			want: fmt.Sprintf("CreateAuth: password environment variable is unset"),
		},
		{
			name: "return error if environment variable is set but blank",
			got:  fmt.Sprintf("%v", err3),
			want: fmt.Sprintf("CreateAuth: password environment variable blank"),
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("name: %v -> got: %v, want: %v", c.name, c.got, c.want),
			func(t *testing.T) {
				fmt.Println(c)
				if c.got != c.want {
					t.Errorf("name: %v -> got: %v, want: %v", c.name, c.got, c.want)
				}
			})
	}
}
