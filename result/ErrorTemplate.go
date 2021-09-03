package result

const (
	EXIST_ERROR     = "要操作的对象已存在"
	NOT_EXIST_ERROR = "要操作的对象不存在"
	ADD_ERROR       = "添加 [%v] 失败"
	UPDATE_ERROR    = "更新 [%v] 失败"
	DEL_ERROR       = "删除 [%v] 失败"

	POOL_ADD_PRIZE_ERROR    = "奖池 [%v] 添加奖品 [%v] 失败"
	POOL_UPDATE_PRIZE_ERROR = "奖池 [%v] 更新奖品 [%v] 失败"
	POOL_DEL_PRIZE_ERROR    = "奖池 [%v] 删除奖品 [%v] 失败"
)
