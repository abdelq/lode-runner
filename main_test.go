package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

// TODO Write a real dial helper function

// TODO Rename function
func testMessageReception(t *testing.T, conn io.Reader, expected message) {
	t.Helper()

	received := message{}
	if err := json.NewDecoder(conn).Decode(&received); err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("expected: %s, received: %s", expected, received)
	}
}
