package logger

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// GormLogger 实现了 gorm logger.Interface，将 GORM 日志桥接到自定义 logger
type GormLogger struct {
	LogLevel                  gormLogger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

// NewGormLogger 创建一个新的 GORM logger 适配器
func NewGormLogger(logLevel gormLogger.LogLevel, slowThreshold time.Duration) *GormLogger {
	return &GormLogger{
		LogLevel:                  logLevel,
		SlowThreshold:             slowThreshold,
		IgnoreRecordNotFoundError: true, // 默认忽略 RecordNotFound 错误
	}
}

// LogMode 设置日志级别
func (l *GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info 记录信息日志
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		Info(msg, data...)
	}
}

// Warn 记录警告日志
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		Warning(msg, data...)
	}
}

// Error 记录错误日志
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		Error(msg, data...)
	}
}

// Trace 记录 SQL 查询日志
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= gormLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && l.LogLevel >= gormLogger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		// 记录 SQL 错误
		Error("SQL Error: %v | Elapsed: %.3fms | Rows: %d | SQL: %s",
			err,
			float64(elapsed.Nanoseconds())/1e6,
			rows,
			sql)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormLogger.Warn:
		// 记录慢查询
		Warning("Slow SQL (> %.0fms): Elapsed: %.3fms | Rows: %d | SQL: %s",
			float64(l.SlowThreshold.Milliseconds()),
			float64(elapsed.Nanoseconds())/1e6,
			rows,
			sql)
	case l.LogLevel == gormLogger.Info:
		// 记录所有 SQL 查询（仅在 Info 级别）
		Info("SQL Query: Elapsed: %.3fms | Rows: %d | SQL: %s",
			float64(elapsed.Nanoseconds())/1e6,
			rows,
			sql)
	}
}
