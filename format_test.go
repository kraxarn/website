package main

import (
	"github.com/kraxarn/website/format"
	"github.com/kraxarn/website/tests"
	"testing"
)

func TestFormatFileSize(t *testing.T) {
	tests.Eq(t, "1", format.FileSize(1))
	tests.Eq(t, "1k", format.FileSize(1_000))
	tests.Eq(t, "1M", format.FileSize(1_000_000))
	tests.Eq(t, "1G", format.FileSize(1_000_000_000))
}
