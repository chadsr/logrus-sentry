package hook_test

import (
	"testing"

	sentryhook "github.com/chadsr/logrus-sentry"

	"github.com/sirupsen/logrus"
)

func TestHookFire(t *testing.T) {
	logger := logrus.New()

	levels := []logrus.Level{logrus.DebugLevel}
	h := sentryhook.New(levels)

	entry := logrus.NewEntry(logger)
	if err := h.Fire(entry); err != nil {
		t.Fatal(err)
	}
}
