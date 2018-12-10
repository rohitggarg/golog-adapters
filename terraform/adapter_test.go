package terraform

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "bytes"
  "github.com/kataras/golog"
  "os"
)

func TestLogMessagePrinting(t *testing.T) {
  buffer := bytes.NewBufferString("")
  logger := golog.New()
  Adapt(logger)
  logger.SetLevel("error")
  logger.SetOutput(buffer)
  logger.Error("Some error occurred")
  assert.Contains(t, buffer.String(), "{\"@level\":\"error\",\"@message\":\"Some error occurred\",\"@timestamp\":\"")
}

func TestLevelGovernedByEnvVar(t *testing.T) {
  os.Setenv("TF_LOG", "info")
  logger := golog.New()
  Adapt(logger)
  assert.Equal(t, logger.Level, golog.InfoLevel)
}

func TestLevelGovernedByEnvVarCaseInsensitive(t *testing.T) {
  os.Setenv("TF_LOG", "INFO")
  logger := golog.New()
  Adapt(logger)
  assert.Equal(t, logger.Level, golog.InfoLevel)
}

func TestDefaultPrintToStdErr(t *testing.T) {
  logger := golog.New()
  Adapt(logger)
  assert.Equal(t, logger.Printer.Output, os.Stderr)
}
