package enemies

type tenemyState int

const (
	enemyShowing    tenemyState = tenemyState(0)
	enemyLeaving    tenemyState = tenemyState(1)
	enemyLookingfor tenemyState = tenemyState(2)
	enemyNextStep   tenemyState = tenemyState(3)
	enemyWalking    tenemyState = tenemyState(4)
)
