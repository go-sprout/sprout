package pesticide

import (
	"testing"
	"time"
)

// ForceTimeLocal temporarily sets [time.Local] for test purpose.
func ForceTimeLocal(t *testing.T, local *time.Location) {
	t.Helper()

	originalLocal := time.Local
	time.Local = local
	t.Cleanup(func() { time.Local = originalLocal })
}
