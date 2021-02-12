package interfaces

//IGamePlayNotificable ...
type IGamePlayNotificable interface {
	OnCreateObject(t, x, y int)
	OnPrepreNewLevel()
	OnRequestPlayerPosition() (float64, float64)
	OnScoreByStep()
}

//IStageNotificable ...
type IStageNotificable interface {
	GetTypeAt(x, y int) int
	SetTypeAt(x, y, t int)
}
