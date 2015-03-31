package model

import (
	"strconv"
)

const (
	DEFAULT_COMPANIES_NUMBER        = 3
	DEFAULT_BANKS_NUMBER            = 1
	DEFAULT_PRODUCTION_TYPES_NUMBER = 3
)

type WorldId int32

type World struct {
	Id              WorldId
	Market          *Market
	Companies       []*Company
	Banks           []*Bank
	ProductionTypes []*Production
}

func NewWorld() {
	world := World{}

	logger.Info("*Begin initialization*")
	logger.Info("Generate companies")
	for i := 0; i < DEFAULT_COMPANIES_NUMBER; i++ {
		world.Companies = append(world.Companies, NewCompany())
	}
	logger.Debug("%i", world.Companies)

	logger.Info("Generate market")
	world.Market = NewMarket()
	logger.Debug("%i", world.Market)

	logger.Info("Generate banks")
	for i := 0; i < DEFAULT_BANKS_NUMBER; i++ {
		world.Banks = append(world.Banks, NewBank())
	}
	logger.Debug("%i", world.Banks)

	logger.Info("Generate bank accounts for companies")
	for _, company := range world.Companies {
		world.Banks[0].OpenAccount(company)
	}
	logger.Info("Generate currency on bank accounts")
	logger.Debug("%i", world.Banks[0])

	logger.Info("Generate production types")
	for i := 0; i < DEFAULT_PRODUCTION_TYPES_NUMBER; i++ {
		world.ProductionTypes = append(world.ProductionTypes,
			NewProduction("Production "+strconv.Itoa(i), CurrencyAmount(i*0.75+10)))

		if i == DEFAULT_PRODUCTION_TYPES_NUMBER-1 {
			world.ProductionTypes[0].dependsOn = world.ProductionTypes[i]
		}
		if i > 0 {
			world.ProductionTypes[i].dependsOn = world.ProductionTypes[i-1]
		}
	}
	logger.Info("Set creation costs for production types")
	logger.Debug("%i", world.ProductionTypes)

	logger.Info("*End*")
}
