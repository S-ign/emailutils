package email

import (
	"fmt"
	"testing"

	"github.com/S-ign/emailutils/config"
)

func TestCreateAuth(t *testing.T) {
	auth1, err := config.CreateAuth(
		"test@test.com",
		"testtesttest",
	)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cases := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "normal",
			got:  fmt.Sprintf("%v", auth1),
			want: fmt.Sprintf("%v", &config.Auth{
				From: "test@test.com", Password: "testtesttest"}),
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
