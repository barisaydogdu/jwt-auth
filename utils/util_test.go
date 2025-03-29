package utils

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func TestCreateToken(t *testing.T) {
	token, err := CreateToken("barisaydogdu@mail.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestVerifyToken(t *testing.T) {
	token, err := CreateToken("barisaydogdu@mail.com")
	if err != nil {
		t.Error(err)
	}

	verifiedToken, err := VerifyToken(token)
	if err != nil {
		t.Error(err)
	}

	t.Log(verifiedToken)
}
