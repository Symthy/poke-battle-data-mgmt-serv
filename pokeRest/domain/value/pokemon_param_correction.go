package value

var (
	NatureARise = ParamCorrection{param: A, rate: 1.1}
	NatureADown = ParamCorrection{param: A, rate: 0.9}
	NatureBRise = ParamCorrection{param: B, rate: 1.1}
	NatureBDown = ParamCorrection{param: B, rate: 0.9}
	NatureCRise = ParamCorrection{param: C, rate: 1.1}
	NatureCDown = ParamCorrection{param: C, rate: 0.9}
	NatureDRise = ParamCorrection{param: D, rate: 1.1}
	NatureDDown = ParamCorrection{param: D, rate: 0.9}
	NatureSRise = ParamCorrection{param: S, rate: 1.1}
	NatureSDown = ParamCorrection{param: S, rate: 0.9}
)

type ParamCorrection struct {
	param PokemonParam
	rate  float32
}
