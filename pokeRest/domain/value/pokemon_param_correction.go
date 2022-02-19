package value

var (
	NatureARise = ParamCorrection{param: Attack, rate: 1.1}
	NatureADown = ParamCorrection{param: Attack, rate: 0.9}
	NatureBRise = ParamCorrection{param: Deffence, rate: 1.1}
	NatureBDown = ParamCorrection{param: Deffence, rate: 0.9}
	NatureCRise = ParamCorrection{param: SpecialAttack, rate: 1.1}
	NatureCDown = ParamCorrection{param: SpecialAttack, rate: 0.9}
	NatureDRise = ParamCorrection{param: SpacialDefense, rate: 1.1}
	NatureDDown = ParamCorrection{param: SpacialDefense, rate: 0.9}
	NatureSRise = ParamCorrection{param: Speed, rate: 1.1}
	NatureSDown = ParamCorrection{param: Speed, rate: 0.9}
)

type ParamCorrection struct {
	param PokemonParam
	rate  float32
}
