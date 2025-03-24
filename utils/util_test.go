package utils

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func TestCreateToken(t *testing.T) {
	token, err := createToken("barisaydogdu@mail.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
