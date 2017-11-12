package message

import "testing"

// TODO Test the unmarshal
// TODO Use t.Error instead of t.Fail
func TestParse(t *testing.T) {
	t.Parallel() // TODO Make sure it's parallel

	t.Run("Spectator", func(t *testing.T) {
		validData := []byte(`{"name": "", "room": "test", "role": 255}`)
		if new(JoinMessage).Parse(validData) != nil {
			t.Fail()
			return
		}

		invalidData := []byte(`{"name": "", "room": " ", "role": 255}`)
		if new(JoinMessage).Parse(invalidData).Error() != "invalid room" {
			t.Fail()
		}
	})

	// TODO Test with 0 as role
	t.Run("Runner", func(t *testing.T) {
		validData := []byte(`{"name": "runner", "room": "test", "role": 38}`)
		if new(JoinMessage).Parse(validData) != nil {
			t.Fail()
			return
		}

		invalidRoomData := []byte(`{"name": "runner", "room": " ", "role": 38}`)
		if new(JoinMessage).Parse(invalidRoomData).Error() != "invalid room" {
			t.Fail()
		}

		invalidNameData := []byte(`{"name": "", "room": "test", "role": 38}`)
		if new(JoinMessage).Parse(invalidNameData).Error() != "invalid name" {
			t.Fail()
		}
	})

	t.Run("Guard", func(t *testing.T) {
		validData := []byte(`{"name": "guard", "room": "test", "role": 48}`)
		if new(JoinMessage).Parse(validData) != nil {
			t.Fail()
			return
		}

		invalidRoomData := []byte(`{"name": "guard", "room": " ", "role": 48}`)
		if new(JoinMessage).Parse(invalidRoomData).Error() != "invalid room" {
			t.Fail()
		}

		invalidNameData := []byte(`{"name": "", "room": "test", "role": 48}`)
		if new(JoinMessage).Parse(invalidNameData).Error() != "invalid name" {
			t.Fail()
		}
	})
}
