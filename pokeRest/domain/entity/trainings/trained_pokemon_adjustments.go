package trainings

type TrainedPokemonAdjustments struct {
	items []*TrainedPokemonAdjustment
}

func NewTrainedPokemonAdjustments(items []*TrainedPokemonAdjustment) *TrainedPokemonAdjustments {
	return &TrainedPokemonAdjustments{items: items}
}

func (p TrainedPokemonAdjustments) Items() []*TrainedPokemonAdjustment {
	return p.items
}
