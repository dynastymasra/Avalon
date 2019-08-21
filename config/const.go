package config

const (
	ServiceName = "Avalon"
	Version     = "0.1.0"
	RequestID   = "request_id"

	envServerAddress = "SERVER_ADDRESS"

	ISO8601Format = "2006-01-02T15:04:05.000Z"

	// Headers
	HeaderRequestID  = "X-Request-ID"
	HeaderAppVersion = "X-App-Version"

	// Database EnvVar
	envDatabaseHost         = "DATABASE_HOST"
	envDatabasePort         = "DATABASE_PORT"
	envDatabaseName         = "DATABASE_NAME"
	envDatabaseUsername     = "DATABASE_USERNAME"
	envDatabasePassword     = "DATABASE_PASSWORD"
	envDatabaseEnableLog    = "DATABASE_ENABLE_LOG"
	envDatabaseMaxOpenConns = "DATABASE_MAX_OPEN_CONNS"
	envDatabaseMaxIdleConns = "DATABASE_MAX_IDLE_CONNS"

	envLoggerFormat = "LOGGER_FORMAT"
	envLogLevel     = "LOG_LEVEL"
)