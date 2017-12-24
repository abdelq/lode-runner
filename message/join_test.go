package message

import "testing"

// TODO Add level

func TestJoinParse(t *testing.T) {
	t.Run("Spectator", func(t *testing.T) {
		t.Parallel()
		msg := new(JoinMessage)

		validData := []byte(`{"name": "", "room": "test", "role": 255}`)
		if err := msg.Parse(validData); err != nil {
			t.Error(err)
			return
		}

		invalidData := []byte(`{"name": "", "room": "", "role": 255}`)
		if err := msg.Parse(invalidData); err.Error() != "invalid room" {
			t.Error(err)
		}
	})

	t.Run("Runner", func(t *testing.T) {
		t.Parallel()
		msg := new(JoinMessage)

		validData := []byte(`{"name": "runner", "room": "test"}`)
		if err := msg.Parse(validData); err != nil {
			t.Error(err)
			return
		}

		invalidRoomData := []byte(`{"name": "runner", "room": "", "role": 0}`)
		if err := msg.Parse(invalidRoomData); err.Error() != "invalid room" {
			t.Error(err)
		}

		invalidNameData := []byte(`{"name": "", "room": "test", "role": 38}`)
		if err := msg.Parse(invalidNameData); err.Error() != "invalid name" {
			t.Error(err)
		}
	})

	t.Run("Guard", func(t *testing.T) {
		t.Parallel()
		msg := new(JoinMessage)

		validData := []byte(`{"name": "guard", "room": "test", "role": 48}`)
		if err := msg.Parse(validData); err != nil {
			t.Error(err)
			return
		}

		invalidRoomData := []byte(`{"name": "guard", "room": "", "role": 48}`)
		if err := msg.Parse(invalidRoomData); err.Error() != "invalid room" {
			t.Error(err)
		}

		invalidNameData := []byte(`{"name": "", "room": "test", "role": 48}`)
		if err := msg.Parse(invalidNameData); err.Error() != "invalid name" {
			t.Error(err)
		}
	})
}
