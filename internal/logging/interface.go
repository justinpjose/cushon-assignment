package logging

type Logger interface {
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Infof(format string, v ...interface{})

	// Field adds a new field to the log which consists of a key and value
	Field(key string, v interface{})

	// CorrelationID creates a new correlation id, adds a correlation id field in the log and returns a new instance of the log which contains log information from the main() func
	CorrelationID() Logger
}
