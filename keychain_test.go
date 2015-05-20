package keychain

import (
	"testing"
)


var app = "github.com/lunixbochs/go-keychain/test"

func TestError(t *testing.T){

    _, err := Find(app, "IDONTEXIST")
    if err == nil {
        t.Fatal("Found password IDONTEXIST?")
    }

    if err.Error() == "Unknown error" {
        t.Fatal("Error should not be 'Unknown error'")
    }

    t.Logf("error was '%q'\n", err)
}

func TestFlow(t *testing.T) {
	err := Add(app, "testuser", "password")
	if err != nil {
		t.Fatal(err)
	}
	pass, err := Find(app, "testuser")
	if err != nil {
		t.Fatal(err)
	}
	if pass != "password" {
		t.Fatalf("password did not match: %s", pass)
	}
	err = Remove(app, "testuser")
	if err != nil {
		t.Fatal(err)
	}
	pass, err = Find(app, "testuser")
	if err == nil || pass != "" {
		t.Fatal("password remove failed")
	}
}
