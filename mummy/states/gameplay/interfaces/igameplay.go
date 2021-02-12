package interfaces

//IGamePlayNotificable ...
type IGamePlayNotificable interface {
	OnCreateObject(t, x, y int)
	OnPrepreNewLevel()
	OnRequestPlayerPosition() (float64, float64)
}

//IStageNotificable ...
type IStageNotificable interface {
	GetTypeAt(x, y int) int
	SetTypeAt(x, y, t int)
}
