package rabbitmq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	rabbitmq "github.com/messaging-go/rabbitmq"
)

func TestHandler_Process(t *testing.T) {
	t.Parallel()
	t.Run("panics with not implemented message", func(t *testing.T) {
		t.Parallel()

		handler := rabbitmq.New[int]()

		assert.Panics(t, func() {
			_ = handler.Process(t.Context(), 1, nil) //nolint:errcheck // this would not return
		})
	})
}
