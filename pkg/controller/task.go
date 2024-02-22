package controller

import (
	_type "github.com/pefish/go-core-type/api-session"
	go_error "github.com/pefish/go-error"
	go_logger "github.com/pefish/go-logger"
	go_mysql "github.com/pefish/go-mysql"
	"github.com/shadouzuo/backend/pkg/constant"
	"github.com/shadouzuo/backend/pkg/db"
)

type TaskControllerType struct {
}

var TaskController = TaskControllerType{}

type AddTaskParams struct {
	Name     string                 `json:"name" validate:"required" desc:""`
	Desc     string                 `json:"desc" desc:""`
	Interval uint64                 `json:"interval" desc:""`
	Data     map[string]interface{} `json:"data" desc:""`
}

func (t *TaskControllerType) AddTask(apiSession _type.IApiSession) (interface{}, *go_error.ErrorInfo) {
	var params AddTaskParams
	err := apiSession.ScanParams(&params)
	if err != nil {
		go_logger.Logger.ErrorF("Read params error. %+v", err)
		return nil, go_error.INTERNAL_ERROR
	}

	_, rowsAffected, err := go_mysql.MysqlInstance.Insert(
		"task",
		db.Task{
			Name:     params.Name,
			Desc:     params.Desc,
			Interval: params.Interval,
			Data:     params.Data,
			Status:   constant.TaskStatusType_Exited,
		},
	)
	if err != nil {
		go_logger.Logger.Error(err)
		return nil, go_error.INTERNAL_ERROR
	}
	if rowsAffected == 0 {
		return nil, go_error.WrapWithStr("Insert failed.")
	}

	return true, nil
}

type ChangeTaskParams struct {
	Id       uint64                  `json:"id" validate:"required"`
	Name     string                  `json:"name" validate:"required"`
	Desc     string                  `json:"desc"`
	Interval uint64                  `json:"interval"`
	Data     map[string]interface{}  `json:"data" validate:"required"`
	Status   constant.TaskStatusType `json:"status" desc:""`
}

func (t *TaskControllerType) ChangeTask(apiSession _type.IApiSession) (interface{}, *go_error.ErrorInfo) {
	var params ChangeTaskParams
	err := apiSession.ScanParams(&params)
	if err != nil {
		go_logger.Logger.ErrorF("Read params error. %+v", err)
		return nil, go_error.INTERNAL_ERROR
	}

	_, rowsAffected, err := go_mysql.MysqlInstance.Update(
		&go_mysql.UpdateParams{
			TableName: "task",
			Update: db.Task{
				Name:     params.Name,
				Desc:     params.Desc,
				Interval: params.Interval,
				Data:     params.Data,
				Status:   params.Status,
			},
			Where: map[string]interface{}{
				"id": params.Id,
			},
		},
	)
	if err != nil {
		go_logger.Logger.Error(err)
		return nil, go_error.INTERNAL_ERROR
	}
	if rowsAffected == 0 {
		return nil, go_error.WrapWithStr("Update failed.")
	}

	return true, nil
}

func (t *TaskControllerType) ListTasks(apiSession _type.IApiSession) (interface{}, *go_error.ErrorInfo) {
	tasks := make([]db.Task, 0)

	err := go_mysql.MysqlInstance.Select(
		&tasks,
		&go_mysql.SelectParams{
			TableName: "task",
			Select:    "*",
		},
	)
	if err != nil {
		go_logger.Logger.Error(err)
		return nil, go_error.INTERNAL_ERROR
	}

	return tasks, nil
}
