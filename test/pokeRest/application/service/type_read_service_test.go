package service

import (
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/application/service/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TypeServiceTestSuite struct {
	suite.Suite
	serv types.TypeReadService
}

// Before
func (suite *TypeServiceTestSuite) SetupTest() {
	suite.serv = types.NewTypeReadService()
}

// After
func (suite *TypeServiceTestSuite) TearDownTest() {
}

func TestTypeControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TypeServiceTestSuite))
}

func (suite TypeServiceTestSuite) TestGetTypeCompatibility() {

	tests := []struct {
		expectedCompatibility [][]float32
		expectedTypeOrder     []string
		lang                  string
	}{
		{
			expectedCompatibility: typeTable,
			expectedTypeOrder:     jpTypeNames,
			lang:                  "ja-JP",
		},
		{
			expectedCompatibility: typeTable,
			expectedTypeOrder:     enTypeNames,
			lang:                  "en-US",
		},
	}

	for _, tt := range tests {
		types := suite.serv.GetTypeCompatibility()
		assert.EqualValues(suite.T(), tt.expectedCompatibility, types.GenerateTypeDamageRateTable())
		assert.EqualValues(suite.T(), tt.expectedCompatibility, types.GenerateTypeNames(tt.lang))
	}

}

var (
	typeTable = [][]float32{
		//ノ,  炎,  水,  電,  草,  氷,  闘,  毒,  地,  飛,  超,  虫,  岩,  霊,  竜,  悪,  鋼,  妖
		{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 0.5, 0.0, 1.0, 1.0, 0.5, 1.0}, /*ノ*/
		{1.0, 0.5, 0.5, 1.0, 2.0, 2.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 0.5, 1.0, 0.5, 1.0, 2.0, 1.0}, /*炎*/
		{1.0, 2.0, 0.5, 1.0, 0.5, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 1.0, 2.0, 1.0, 0.5, 1.0, 1.0, 1.0}, /*水*/
		{1.0, 1.0, 2.0, 0.5, 0.5, 1.0, 1.0, 1.0, 0.0, 2.0, 1.0, 1.0, 1.0, 1.0, 0.5, 1.0, 1.0, 1.0}, /*電*/
		{1.0, 0.5, 2.0, 1.0, 0.5, 1.0, 1.0, 0.5, 2.0, 0.5, 1.0, 0.5, 2.0, 1.0, 0.5, 1.0, 0.5, 1.0}, /*草*/
		{1.0, 0.5, 0.5, 1.0, 2.0, 0.5, 1.0, 1.0, 2.0, 2.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 0.5, 1.0}, /*氷*/
		{2.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 0.5, 1.0, 0.5, 0.5, 0.5, 2.0, 0.0, 1.0, 2.0, 2.0, 0.5}, /*闘*/
		{1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 0.5, 0.5, 1.0, 1.0, 1.0, 0.5, 0.5, 1.0, 1.0, 0.0, 2.0}, /*毒*/
		{1.0, 2.0, 1.0, 2.0, 0.5, 1.0, 1.0, 2.0, 1.0, 0.0, 1.0, 0.5, 2.0, 1.0, 1.0, 1.0, 2.0, 1.0}, /*地*/
		{1.0, 1.0, 1.0, 0.5, 2.0, 1.0, 2.0, 1.0, 1.0, 1.0, 1.0, 2.0, 0.5, 1.0, 1.0, 1.0, 0.5, 1.0}, /*飛*/
		{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 2.0, 1.0, 1.0, 0.5, 1.0, 1.0, 1.0, 1.0, 0.0, 0.5, 1.0}, /*超*/
		{1.0, 0.5, 1.0, 1.0, 2.0, 1.0, 0.5, 0.5, 1.0, 0.5, 2.0, 1.0, 1.0, 0.5, 1.0, 2.0, 0.5, 0.5}, /*虫*/
		{1.0, 2.0, 1.0, 1.0, 1.0, 2.0, 0.5, 1.0, 0.5, 2.0, 1.0, 2.0, 1.0, 1.0, 1.0, 1.0, 0.5, 1.0}, /*岩*/
		{0.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 2.0, 1.0, 0.5, 1.0, 1.0}, /*霊*/
		{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 0.5, 0.0}, /*竜*/
		{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 0.5, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 2.0, 1.0, 0.5, 1.0, 0.5}, /*悪*/
		{1.0, 0.5, 0.5, 0.5, 1.0, 2.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 1.0, 1.0, 1.0, 0.5, 2.0}, /*鋼*/
		{1.0, 0.5, 1.0, 1.0, 1.0, 1.0, 2.0, 0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 2.0, 2.0, 0.5, 1.0}, /*妖*/
	}
	jpTypeNames = []string{
		"ノーマル",
		"ほのお",
		"みず",
		"でんき",
		"くさ",
		"こおり",
		"かくとう",
		"どく",
		"じめん",
		"ひこう",
		"エスパー",
		"むし",
		"いわ",
		"ゴースト",
		"ドラゴン",
		"あく",
		"はがね",
		"フェアリー",
	}
	enTypeNames = []string{
		"Normal",
		"Fire",
		"Water",
		"Electric",
		"Grass",
		"Ice",
		"Fighting",
		"Poison",
		"Ground",
		"Flying",
		"Psychic",
		"Bug",
		"Rock",
		"Ghost",
		"Dragon",
		"Dark",
		"Steel",
		"Fairy",
	}
)
