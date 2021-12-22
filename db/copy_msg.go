package db

import (
	"context"
	"fmt"
	"os"
	"runtime"
)

var name, _ = os.Hostname()
var fromSystem = fmt.Sprintf("%s-%s-%s", name, runtime.GOOS, runtime.GOARCH)

func SaveCopyMsg(ctx context.Context, text string) error {
	_, err := client.CopyMsg.Create().SetFrom(fromSystem).SetBody(text).Save(ctx)
	return err
}
