package logs

import (
	"encoding/json"
	"fmt"
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zaplogs *zap.Logger
var err error

func Launch() {
	zapconfig := zap.NewDevelopmentConfig()
	zapconfig.EncoderConfig.TimeKey = "timestamp"
	zapconfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapconfig.EncoderConfig.StacktraceKey = ""
	zaplogs, err = zapconfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}
func Info(message string, felids ...zap.Field) {
	zaplogs.Info(message, felids...)
}

func Dubug(message interface{}, felids ...zap.Field) {
	// ใช้ reflect ตรวจสอบว่าเป็น struct หรือไม่
	v := reflect.ValueOf(message)
	if v.Kind() == reflect.Struct {
		b, err := json.MarshalIndent(message, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling struct:", err)
			return
		}
		zaplogs.Debug("\n"+string(b), felids...)
	} else {
		zaplogs.Debug(fmt.Sprintf("%v", message), felids...)
	}

}

func Error(message interface{}, felids ...zap.Field) {
	switch err := message.(type) {
	case error:
		zaplogs.Error(err.Error(), felids...)
	case string:
		zaplogs.Error(err, felids...)
	}
}
func Warning(message string, felids ...zap.Field) {
	zaplogs.Warn(message, felids...)
}
func Panic(message string, felids ...zap.Field) {
	zaplogs.Panic(message, felids...)
}
