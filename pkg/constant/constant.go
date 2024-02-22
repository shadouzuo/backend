package constant

const (
	PARAM_ERROR uint64 = 1001 // 参数错误
)

type TaskStatusType uint64

const (
	TaskStatusType_WaitExec      TaskStatusType = 1
	TaskStatusType_Executing     TaskStatusType = 2
	TaskStatusType_Exited        TaskStatusType = 3
	TaskStatusType_ExitedWithErr TaskStatusType = 4
)
