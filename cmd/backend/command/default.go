package command

import (
	"flag"

	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	"github.com/pefish/go-core/driver/logger"
	global_api_strategy "github.com/pefish/go-core/global-api-strategy"
	"github.com/pefish/go-core/service"
	go_logger "github.com/pefish/go-logger"
	go_mysql "github.com/pefish/go-mysql"
	task_driver "github.com/pefish/go-task-driver"
	"github.com/shadouzuo/backend/pkg/constant"
	"github.com/shadouzuo/backend/pkg/global"
	"github.com/shadouzuo/backend/pkg/route"
	"github.com/shadouzuo/backend/version"
)

type DefaultCommand struct {
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	flagSet.String("server-host", "0.0.0.0", "The host of web server.")
	flagSet.Int("server-port", 8000, "The port of web server.")
	return nil
}

func (dc *DefaultCommand) OnExited(command *commander.Commander) error {
	go_mysql.MysqlInstance.Close()
	return nil
}

func (dc *DefaultCommand) Init(command *commander.Commander) error {
	service.Service.SetName(version.AppName)
	logger.LoggerDriverInstance.Register(go_logger.Logger)

	err := go_config.ConfigManagerInstance.Unmarshal(&global.GlobalConfig)
	if err != nil {
		return err
	}

	go_mysql.MysqlInstance.SetLogger(go_logger.Logger)
	err = go_mysql.MysqlInstance.ConnectWithConfiguration(go_mysql.Configuration{
		Host:     global.GlobalConfig.Db.Host,
		Username: global.GlobalConfig.Db.User,
		Password: global.GlobalConfig.Db.Pass,
		Database: global.GlobalConfig.Db.Db,
	})
	if err != nil {
		return err
	}

	service.Service.SetHost(global.GlobalConfig.ServerHost)
	service.Service.SetPort(global.GlobalConfig.ServerPort)
	service.Service.SetPath(`/api`)
	global_api_strategy.ParamValidateStrategyInstance.SetErrorCode(constant.PARAM_ERROR)

	service.Service.SetRoutes(route.TaskRoute)

	return nil
}

func (dc *DefaultCommand) Start(command *commander.Commander) error {

	taskDriver := task_driver.NewTaskDriver()
	taskDriver.Register(service.Service)

	taskDriver.RunWait(command.Ctx)

	return nil
}
