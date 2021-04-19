package observability

import (
	"os"

	"github.com/calvinverse/service.provisioning.ui.web/internal/meta"
	log "github.com/sirupsen/logrus"
)

var (
	defaultLogger *log.Entry
)

func getCustomLoggerWithDefaultFields(logger *log.Logger) *log.Entry {
	var entry log.FieldLogger
	entry = log.NewEntry(logger)

	return entry.WithFields(getDefaultLoggerFields())
}

func getDefaultLoggerFields() log.Fields {
	logFields := log.Fields{}

	logFields["application_name"] = meta.ApplicationName()
	logFields["application_version"] = meta.Version()

	return logFields
}

func getStandardLoggerWithDefaultFields() *log.Entry {
	return log.WithFields(getDefaultLoggerFields())
}

// InitializeLogger initializes a logger for the current application
func InitializeLogger() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)

	defaultLogger = getStandardLoggerWithDefaultFields()
}

// LogDebug logs a debug message with the given arguments
func LogDebug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// LogDebugWithFields logs a debug message with the given arguments and the given fields
func LogDebugWithFields(fields log.Fields, args ...interface{}) {
	defaultLogger.WithFields(fields).Debug(args...)
}

// LogError logs an error message with the given arguments
func LogError(args ...interface{}) {
	defaultLogger.Error(args...)
}

// LogError logs an error message with the arguments and the provided fields
func LogErrorWithFields(fields log.Fields, args ...interface{}) {
	defaultLogger.WithFields(fields).Error(args...)
}

// LogFatal logs a fatal message with the given arguments
func LogFatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// LogFatal logs a fatal message with the arguments and the provided fields
func LogFatalWithFields(fields log.Fields, args ...interface{}) {
	defaultLogger.WithFields(fields).Fatal(args...)
}

// LogInfo logs an info message with the given arguments
func LogInfo(args ...interface{}) {
	defaultLogger.Info(args...)
}

// LogInfo logs an info message with the arguments and the provided fields
func LogInfoWithFields(fields log.Fields, args ...interface{}) {
	defaultLogger.WithFields(fields).Info(args...)
}

// LogPanic logs a panic message with the given arguments
func LogPanic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// LogPanicWithFields logs a panic message with the arguments and the provided fields
func LogPanicWithFields(fields log.Fields, args ...interface{}) {
	defaultLogger.WithFields(fields).Panic(args...)
}

// LogWarn logs a warning message with the provided arguments
func LogWarn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// LogWarnWithFields logs a warning message with the arguments and the provided fields
func LogWarnWithFields(fields log.Fields, args ...interface{}) {
	defaultLogger.WithFields(fields).Warn(args...)
}

func NewLogger() *log.Entry {
	logger := log.New()

	logger.Formatter = &log.JSONFormatter{
		DisableTimestamp: true,
	}

	logger.Level = log.DebugLevel

	return getCustomLoggerWithDefaultFields(logger)
}

// SetLogLevel sets the minimum level of log messages that should be recorded
func SetLogLevel(level string) {
	switch level {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
