package fcm

import "testing"

func TestValidate(t *testing.T) {
	t.Run("valid with token", func(t *testing.T) {
		msg := &Message{
			Token:      "test",
			TimeToLive: 3600,
			Data: map[string]interface{}{
				"message": "This is a Firebase Cloud Messaging Topic Message!",
			},
		}
		err := msg.Validate()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("invalid message", func(t *testing.T) {
		var msg *Message
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got <nil>", ErrInvalidMessage)
		}
	})

	t.Run("invalid target", func(t *testing.T) {
		msg := &Message{
			Data: map[string]interface{}{
				"message": "This is a Firebase Cloud Messaging Topic Message!",
			},
		}
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got nil", ErrInvalidTarget)
		}
	})

	t.Run("too many registration ids", func(t *testing.T) {
		msg := &Message{
			Token:           "test",
			RegistrationIDs: make([]string, 2000),
		}
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got <nil>", ErrToManyRegIDs)
		}
	})

	t.Run("invalid TTL", func(t *testing.T) {
		msg := &Message{
			Token:           "test",
			RegistrationIDs: []string{"reg_id"},
			TimeToLive:      2500000,
			Data: map[string]interface{}{
				"message": "This is a Firebase Cloud Messaging Topic Message!",
			},
		}
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got nil", ErrInvalidTimeToLive)
		}
	})

	t.Run("valid with registration ID", func(t *testing.T) {
		msg := &Message{
			RegistrationIDs: []string{"reg_id"},
			Data: map[string]interface{}{
				"message": "This is a Firebase Cloud Messaging Topic Message!",
			},
		}
		err := msg.Validate()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("valid with condition", func(t *testing.T) {
		msg := &Message{
			Condition: "'dogs' in topics || 'cats' in topics",
			Data: map[string]interface{}{
				"message": "This is a Firebase Cloud Messaging Topic Message!",
			},
		}
		err := msg.Validate()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("invalid condition", func(t *testing.T) {
		msg := &Message{
			Condition: "'TopicA' in topics && ('TopicB' in topics || 'TopicC' in topics || 'TopicD' in topics )",
			Data: map[string]interface{}{
				"message": "This is a Firebase Cloud Messaging Topic Message!",
			},
		}
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got nil", ErrInvalidTarget)
		}
	})
}
