package main

import (
	"github.com/kraxarn/website/tests"
	"testing"
)

func TestFormatFileSize(t *testing.T) {
	tests.Eq(t, "1", FormatFileSize(1))
	tests.Eq(t, "1k", FormatFileSize(1_000))
	tests.Eq(t, "1M", FormatFileSize(1_000_000))
	tests.Eq(t, "1G", FormatFileSize(1_000_000_000))
}
