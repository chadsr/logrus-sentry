# Sentry hook for Logrus logger

[![Tests](https://github.com/chadsr/logrus-sentry/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/chadsr/logrus-sentry/actions/workflows/tests.yml)
[![codecov](https://codecov.io/gh/chadsr/logrus-sentry/branch/master/graph/badge.svg?token=C4GM92GQ3I)](https://codecov.io/gh/chadsr/logrus-sentry)

[Sentry](https://github.com/getsentry/sentry-go) hook for [Logrus](https://github.com/sirupsen/logrus) logger.

## Examples

Basic:

```go
import (
    "github.com/sirupsen/logrus"
    "github.com/getsentry/sentry-go"

    sentryhook "github.com/chadsr/logrus-sentry"
)

func main() {
    if err := sentry.Init(sentry.ClientOptions{Dsn: "aDSN"}); err != nil {
        log.Fatal(err)
    }
    
    logger := logrus.New()
    logger.AddHook(sentryhook.New([]logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel}))
    
    logger.Fatal("the error would be sent to sentry")
}
```

Customize Sentry event:

```go
import (
    "github.com/sirupsen/logrus"
    "github.com/getsentry/sentry-go"

    sentryhook "github.com/chadsr/logrus-sentry"
)

func main() {
    if err := sentry.Init(sentry.ClientOptions{Dsn: "aDSN"}); err != nil {
        log.Fatal(err)
    }
    
    logger := logrus.New()
    logger.AddHook(sentryhook.New(
        []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel},
        sentryhook.WithConverter(newConverter()),
    ))
    
    logger.
        WithField("corr_id", "aCorrId").
        WithField("user_id", "aUserId").
        Fatal("the error would be sent to sentry")
}

func newConverter() sentryhook.Converter {
    return func(entry *logrus.Entry, hub *sentry.Hub) *sentry.Event {
        event := sentryhook.DefaultConverter(entry, hub)

        if pkg, ok := entry.Data["pkg"].(string); ok {
            event.Logger = pkg
        }

        if corrID, ok := entry.Data["corr_id"].(string); ok {
            event.Tags["corr_id"] = corrID
        }

        if userID, ok := entry.Data["user_id"]; ok {
            event.Tags["user_id"] = fmt.Sprintf("%v", userID)
        }

        return event
    }
}
```
