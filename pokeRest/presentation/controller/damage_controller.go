package controller

import "github.com/Symthy/PokeRest/pokeRest/application/service/damages"

type DamageController struct {
	damageCalcService damages.DamageCalculationService
}
