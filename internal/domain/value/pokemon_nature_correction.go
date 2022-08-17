package value

var (
	NatureARise = NatureCorrection{param: A, rate: 1.1}
	NatureADown = NatureCorrection{param: A, rate: 0.9}
	NatureBRise = NatureCorrection{param: B, rate: 1.1}
	NatureBDown = NatureCorrection{param: B, rate: 0.9}
	NatureCRise = NatureCorrection{param: C, rate: 1.1}
	NatureCDown = NatureCorrection{param: C, rate: 0.9}
	NatureDRise = NatureCorrection{param: D, rate: 1.1}
	NatureDDown = NatureCorrection{param: D, rate: 0.9}
	NatureSRise = NatureCorrection{param: S, rate: 1.1}
	NatureSDown = NatureCorrection{param: S, rate: 0.9}
)

type NatureCorrection struct {
	param PokemonParam
	rate  float32
}
