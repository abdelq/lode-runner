package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

// TODO Rename
func testMessage(t *testing.T, conn net.Conn, expectedMsg message) {
	t.Helper()

	receivedMsg := message{}
	if err := json.NewDecoder(conn).Decode(&receivedMsg); err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(expectedMsg, receivedMsg) {
		t.Errorf("expected: %s, received: %s", expectedMsg, receivedMsg)
	}
}
