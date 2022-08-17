package controller

import "github.com/Symthy/PokeRest/internal/application/service/damages"

type DamageController struct {
	damageCalcService damages.DamageCalculationService
}
