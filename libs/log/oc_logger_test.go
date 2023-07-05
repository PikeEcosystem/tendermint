package log_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/PikeEcosystem/tendermint/libs/log"
)

func TestLoggerLogsItsErrors(t *testing.T) {
	var buf bytes.Buffer

	logger := log.NewOCLogger(&buf)
	logger.Info("foo", "baz baz", "bar")
	msg := strings.TrimSpace(buf.String())
	if !strings.Contains(msg, "foo") {
		t.Errorf("expected logger msg to contain ErrInvalidKey, got %s", msg)
	}
}

func TestInfo(t *testing.T) {
	var bufInfo bytes.Buffer

	l := log.NewOCLogger(&bufInfo)
	l.Info("Client initialized with old header (trusted is more recent)",
		"old", 42,
		"trustedHeight", "forty two",
		"trustedHash", []byte("test me"))

	msg := strings.TrimSpace(bufInfo.String())

	// Remove the timestamp information to allow
	// us to test against the expected message.
	receivedmsg := strings.Split(msg, "] ")[1]

	const expectedmsg = `Client initialized with old header
	(trusted is more recent) old=42 trustedHeight="forty two"
	trustedHash=74657374206D65`
	if strings.EqualFold(receivedmsg, expectedmsg) {
		t.Fatalf("received %s, expected %s", receivedmsg, expectedmsg)
	}
}

func TestDebug(t *testing.T) {
	var bufDebug bytes.Buffer

	ld := log.NewOCLogger(&bufDebug)
	ld.Debug("Client initialized with old header (trusted is more recent)",
		"old", 42,
		"trustedHeight", "forty two",
		"trustedHash", []byte("test me"))

	msg := strings.TrimSpace(bufDebug.String())

	// Remove the timestamp information to allow
	// us to test against the expected message.
	receivedmsg := strings.Split(msg, "] ")[1]

	const expectedmsg = `Client initialized with old header
	(trusted is more recent) old=42 trustedHeight="forty two"
	trustedHash=74657374206D65`
	if strings.EqualFold(receivedmsg, expectedmsg) {
		t.Fatalf("received %s, expected %s", receivedmsg, expectedmsg)
	}
}

func TestError(t *testing.T) {
	var bufErr bytes.Buffer

	le := log.NewOCLogger(&bufErr)
	le.Error("Client initialized with old header (trusted is more recent)",
		"old", 42,
		"trustedHeight", "forty two",
		"trustedHash", []byte("test me"))

	msg := strings.TrimSpace(bufErr.String())

	// Remove the timestamp information to allow
	// us to test against the expected message.
	receivedmsg := strings.Split(msg, "] ")[1]

	const expectedmsg = `Client initialized with old header
	(trusted is more recent) old=42 trustedHeight="forty two"
	trustedHash=74657374206D65`
	if strings.EqualFold(receivedmsg, expectedmsg) {
		t.Fatalf("received %s, expected %s", receivedmsg, expectedmsg)
	}
}

func BenchmarkOCLoggerSimple(b *testing.B) {
	benchmarkRunner(b, log.NewOCLogger(io.Discard), baseInfoMessage)
}

func BenchmarkOCLoggerContextual(b *testing.B) {
	benchmarkRunner(b, log.NewOCLogger(io.Discard), withInfoMessage)
}

func benchmarkRunner(b *testing.B, logger log.Logger, f func(log.Logger)) {
	lc := logger.With("common_key", "common_value")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(lc)
	}
}

var (
	baseInfoMessage = func(logger log.Logger) { logger.Info("foo_message", "foo_key", "foo_value") }
	withInfoMessage = func(logger log.Logger) { logger.With("a", "b").Info("c", "d", "f") }
)