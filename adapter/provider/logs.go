package provider

import "go.uber.org/zap"

const (
	INFO  = "Info"
	WARN  = "Warn"
	ERROR = "Error"
	DEBUG = "Debug"
)

type SystemLogger struct {
	zapLogger     *zap.Logger
	logger        *zap.SugaredLogger
	msg           string
	typeLog       string
	keysAndValues []interface{}
}

func NewLogger() *SystemLogger {
	cfg := zap.NewProductionConfig()
	//	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := cfg.Build(zap.AddCallerSkip(2))

	zapLogger := logger.Named("SmartChat-Logger")

	return &SystemLogger{
		zapLogger:     zapLogger,
		logger:        zapLogger.Sugar(),
		msg:           "None",
		typeLog:       INFO,
		keysAndValues: nil,
	}
}

func (l *SystemLogger) ZapSync() {
	l.zapLogger.Sync()
}

func (l *SystemLogger) NewLog(msg string, requestUUID string, keysAndValues ...interface{}) *SystemLogger {
	l.toClean()

	l.msg = msg

	l.keysAndValues = append(l.keysAndValues, "requestID", requestUUID)

	if len(keysAndValues) > 0 {
		l.keysAndValues = append(l.keysAndValues, keysAndValues...)
	}

	return l
}

func (l *SystemLogger) Info() *SystemLogger {
	l.typeLog = INFO
	return l
}

func (l *SystemLogger) Warn() *SystemLogger {
	l.typeLog = WARN
	return l
}

func (l *SystemLogger) Error() *SystemLogger {
	l.typeLog = ERROR
	return l
}

func (l *SystemLogger) Debug() *SystemLogger {
	l.typeLog = DEBUG
	return l
}

func (l *SystemLogger) Description(description string) *SystemLogger {
	l.keysAndValues = append(l.keysAndValues, "description", description)
	return l
}

func (l *SystemLogger) Phase(phase string) *SystemLogger {
	l.keysAndValues = append(l.keysAndValues, "phase", phase)
	return l
}

func (l *SystemLogger) Request() {
	l.msg = l.msg + "_REQUEST"
	l.Exe()
}

func (l *SystemLogger) Response() {
	l.msg = l.msg + "_RESPONSE"
	l.Exe()
}

func (l *SystemLogger) toClean() {
	l.msg = ""
	l.typeLog = INFO
	l.keysAndValues = nil
}

func (l *SystemLogger) Exe() {
	//errMsg := fmt.Sprintf("%s: %+v", l.msg, l.keysAndValues)
	//err := errors.New("\n\nMensage error: " + errMsg)
	//err = errors.WithStack(err)

	switch l.typeLog {
	case INFO:
		l.logger.Infow(l.msg, l.keysAndValues...)
	case WARN:
		l.logger.Warnw(l.msg, l.keysAndValues...)
	case ERROR:
		l.logger.Errorw(l.msg, l.keysAndValues...)
	case DEBUG:
		l.logger.Debugw(l.msg, l.keysAndValues...)
	}

	l.toClean()
}
