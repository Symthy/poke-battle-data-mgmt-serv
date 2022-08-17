package battles

type SeasonPeriods struct {
	items []SeasonPeriod
}

func NewSeasonPeriods(items []SeasonPeriod) SeasonPeriods {
	return SeasonPeriods{items: items}
}

func (s SeasonPeriods) Items() []SeasonPeriod {
	return s.items
}
