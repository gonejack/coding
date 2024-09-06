package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// Foreground colors.
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Color represents a text color.
type Color uint8

// Add adds the coloring to the given string.
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}

var (
	_levelToColor = map[zapcore.Level]Color{
		zapcore.DebugLevel:  Magenta,
		zapcore.InfoLevel:   Blue,
		zapcore.WarnLevel:   Yellow,
		zapcore.ErrorLevel:  Red,
		zapcore.DPanicLevel: Red,
		zapcore.PanicLevel:  Red,
		zapcore.FatalLevel:  Red,
	}
	_unknownLevelColor         = Red
	_levelToCapitalColorString = make(map[zapcore.Level]string, len(_levelToColor))
)

func init() {
	for level, color := range _levelToColor {
		var text string
		switch level {
		case zapcore.DebugLevel:
			text = "DBUG"
		case zapcore.InfoLevel:
			text = "INFO"
		case zapcore.WarnLevel:
			text = "WARN"
		case zapcore.ErrorLevel:
			text = "ERRO"
		case zapcore.DPanicLevel:
			text = "DPANIC"
		case zapcore.PanicLevel:
			text = "PANIC"
		case zapcore.FatalLevel:
			text = "FATA"
		default:
			text = fmt.Sprintf("LEVEL(%d)", level)
		}
		_levelToCapitalColorString[level] = fmt.Sprintf("[%s]", color.Add(text))
	}
}

func CapitalColorLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	s, ok := _levelToCapitalColorString[l]
	if !ok {
		s = _unknownLevelColor.Add(l.CapitalString())
	}
	enc.AppendString(s)
}

func NewLogger() *zap.Logger {
	cfg := zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "ts",
		NameKey:          "logger",
		ConsoleSeparator: " ",
		EncodeName:       zapcore.FullNameEncoder,
		EncodeLevel:      CapitalColorLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout(fmt.Sprintf("[%s]", time.DateTime)),
		EncodeDuration:   zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zap.DebugLevel)
	return zap.New(core).WithOptions()
}

func main() {
	x := NewLogger()
	xx := x.With(
		zap.String("package", "main"),
		zap.String("action", "action"),
	).Named("thislogger").Sugar()
	xx.Infof("message_%d", 1)
	xx = xx.With("hello", "world")
	xx.Debugf("abc_%d", 122)
}
