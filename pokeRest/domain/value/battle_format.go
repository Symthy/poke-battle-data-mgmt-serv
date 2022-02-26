package value

type BattleFormat string

const (
	Single BattleFormat = "Single"
	Double BattleFormat = "Double"
)

func resolveBattleFormat(value string) BattleFormat {
	if value == "Double" {
		return Double
	}
	return Single
}
