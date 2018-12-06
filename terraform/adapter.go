package terraform

import (
  "github.com/kataras/golog"
  "github.com/kataras/pio"
  "encoding/json"
  "os"
)

type terraformLog struct {
  Level string `json:"@level"`
  Message string `json:"@message"`
  Timestamp string `json:"@timestamp"`
}

func Adapt(logger *golog.Logger) {
  terraformLogLevel := os.Getenv("TF_LOG")
  logger.SetLevel(terraformLogLevel)
  logger.SetTimeFormat("2006-01-02T15:04:05.000000Z07:00")
  logger.Printer = pio.NewTextPrinter("terraform", os.Stderr).EnableDirectOutput().Hijack(logHijack)
}

var logHijack = func(ctx *pio.Ctx) {
  l, ok := ctx.Value.(*golog.Log)
  if !ok {
    ctx.Next()
    return
  }

  logLine := &terraformLog{
    Level: golog.Levels[l.Level].Name,
    Timestamp: l.FormatTime(),
    Message: string(l.Logger.Prefix) + l.Message,
  }

  ctx.Store(json.Marshal(logLine))
  ctx.Next()
}
