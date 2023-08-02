package main

import (
	"context"

	"github.com/edaniels/golog"
	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"

	"thegreatco/viam-appliedmotion/stf10_ip"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("appliedmotion"))
}

func mainWithArgs(ctx context.Context, args []string, logger golog.Logger) (err error) {
	custom_module, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	err = custom_module.AddModelFromRegistry(ctx, board.API, stf10_ip.Model)
	if err != nil {
		return err
	}

	err = custom_module.Start(ctx)
	defer custom_module.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
