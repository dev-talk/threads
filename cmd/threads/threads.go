package main

import (
	"context"
	"fmt"
	"runtime"

	flag "github.com/bborbe/flagenv"
	"github.com/kolide/kit/version"
	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	appName = "DevTalk Threads"
	appKey  = "threads"
)

var (
	dbg         = flag.Bool("debug", false, "enable debug mode")
	sentryDsn   = flag.String("sentryDsn", "", "sentry dsn key")
	versionInfo = flag.Bool("version", true, "show version info")
	port        = flag.Int("port", 8080, "http server port (default: 8080)")
)

func main() {
	flag.Parse()

	if *versionInfo {
		v := version.Version()
		fmt.Printf("-- %s --\n", appName)
		fmt.Printf(" - version: %s\n", v.Version)
		fmt.Printf("   branch: \t%s\n", v.Branch)
		fmt.Printf("   revision: \t%s\n", v.Revision)
		fmt.Printf("   build date: \t%s\n", v.BuildDate)
		fmt.Printf("   build user: \t%s\n", v.BuildUser)
		fmt.Printf("   go version: \t%s\n", v.GoVersion)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())

	var zapFields []zapcore.Field
	if !*dbg {
		zapFields = []zapcore.Field{
			zap.String("app", appKey),
			zap.String("version", version.Version().Version),
		}
	}

	logger := log.New(*sentryDsn, *dbg).WithFields(zapFields...)
	defer logger.Sync()
	logger.Info("preparing")

	ctx := log.WithLogger(context.Background(), logger)

	log.From(ctx).Info("finished")
}
