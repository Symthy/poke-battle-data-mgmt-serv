package value

type CorrectionValues struct {
	values []CorrectionValue
}

func NewCorrectionValues(values []CorrectionValue) CorrectionValues {
	return CorrectionValues{values}
}
