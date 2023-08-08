# Logging

## Decision
The decision to use zerolog was based on
- Being the most performant out of the other loggers available - https://blog.logrocket.com/5-structured-logging-packages-for-go/
- Provides structured JSON logging for easier search, integration, readability and parsing

## Enchanments
- Check logs in the test to see if things are being logged correctly

[//]: # (Reference Links)
[structured-logging]: <https://www.loggly.com/use-cases/what-is-structured-logging-and-how-to-use-it/>