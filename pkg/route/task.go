package route

import (
	"github.com/pefish/go-core/api"
	"github.com/pefish/go-http/gorequest"
	"github.com/shadouzuo/backend/pkg/controller"
)

var TaskRoute = []*api.Api{
	api.NewApi(&api.NewApiParamsType{
		Description:    "",
		Path:           "/v1/add-task",
		Method:         gorequest.POST,
		Params:         controller.AddTaskParams{},
		ControllerFunc: controller.TaskController.AddTask,
	}),
	api.NewApi(&api.NewApiParamsType{
		Description:    "",
		Path:           "/v1/change-task",
		Method:         gorequest.POST,
		Params:         controller.ChangeTaskParams{},
		ControllerFunc: controller.TaskController.ChangeTask,
	}),
	api.NewApi(&api.NewApiParamsType{
		Description:    "",
		Path:           "/v1/list-tasks",
		Method:         gorequest.GET,
		ControllerFunc: controller.TaskController.ListTasks,
	}),
}
