package lib

import (
	"context"
	"time"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	gormlogger "gorm.io/gorm/logger"
)

type Logger struct {
	*zap.SugaredLogger
}

type FxLogger struct {
	*Logger
}

type GormLogger struct {
	*Logger
	gormlogger.Config
}

var (
	globalLogger *Logger
	zapLogger    *zap.Logger
)

// GetLogger get the logger
func GetLogger() Logger {
	if globalLogger == nil {
		logger := newLogger(NewEnv())
		globalLogger = &logger
	}
	return *globalLogger
}

// sugared logger -> 加工的 logger
// 大部分的功能都是相同的；並且透過 `.Sugar()` 來獲取一個 `SugaredLogger`
// 然後使用 SugaredLogger 以 printf 格式紀錄語句
func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

// GetFxLogger gets logger for go-fx
// 依賴注入使用的 logger
func (l *Logger) GetFxLogger() fxevent.Logger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	// return &FxLogger
	return &FxLogger{Logger: newSugaredLogger(logger)}
}

// LogEvent log event for fx logger
// https://pkg.go.dev/go.uber.org/fx@v1.20.1/fxevent#Logger
// 提到了 type Logger interface { LogEvent(Event) }
// 因此我們需要實作 LogEvent
func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.Debug("OnStart hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.Debug("OnStart hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.Logger.Debug("OnStart hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.Logger.Debug("OnStop hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.Debug("OnStop hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.Logger.Debug("OnStop hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		l.Logger.Debug("supplied: ", zap.String("type", e.TypeName), zap.Error(e.Err))
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("provided: ", e.ConstructorName, " => ", rtype)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("decorated: ",
				zap.String("decorator", e.DecoratorName),
				zap.String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		l.Logger.Debug("invoking: ", e.FunctionName)
	case *fxevent.Started:
		if e.Err == nil {
			l.Logger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			l.Logger.Debug("initialized: custom fxevent.Logger -> ", e.ConstructorName)
		}
	}
}

// newLogger sets up logger
func newLogger(env Env) Logger {

	// 獲取 zap 設定檔(開發模式)
	config := zap.NewDevelopmentConfig()
	logOutput := env.LogOutput

	if env.Environment == "development" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if env.Environment == "production" && logOutput != "" {
		config.OutputPaths = []string{logOutput}
	}

	logLevel := env.LogLevel
	level := zap.PanicLevel
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zap.PanicLevel
	}
	config.Level.SetLevel(level)

	zapLogger, _ = config.Build()

	logger := newSugaredLogger(zapLogger)
	return *logger
}

// GetGormLogger gets the gorm framework logger
func (l Logger) GetGormLogger() *GormLogger {
	logger := zapLogger.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)

	return &GormLogger{
		Logger: newSugaredLogger(logger),
		Config: gormlogger.Config{
			LogLevel: gormlogger.Info,
		},
	}
}

// LogMode set log mode
func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info prints info
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.Debugf(str, args...)
	}
}

// Warn prints warn messages
func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		l.Warnf(str, args...)
	}

}

// Error prints error messages
func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		l.Errorf(str, args...)
	}
}

// Trace prints trace messages
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	if l.LogLevel >= gormlogger.Info {
		sql, rows := fc()
		l.Debug("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}

	if l.LogLevel >= gormlogger.Warn {
		sql, rows := fc()
		l.SugaredLogger.Warn("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}

	if l.LogLevel >= gormlogger.Error {
		sql, rows := fc()
		l.SugaredLogger.Error("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}
}
