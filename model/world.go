package model

import (
	"math/rand"
)

const (
	DEFAULT_COMPANIES_NUMBER = 3
	DEFAULT_BANKS_NUMBER     = 1
)

type WorldId int32
type WorldTick int64
type WorldStatus int8

type WorldMetaInfo struct {
	Id     WorldId
	Tick   WorldTick
	Status WorldStatus
}

type World struct {
	MetaInfo WorldMetaInfo

	Market    *Market
	Companies []*Company
	Banks     []*Bank
}

func NewWorld() *World {
	world := World{MetaInfo: WorldMetaInfo{Id: WorldId(rand.Int31()), Tick: 0, Status: 0}}

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
	logger.Debug("%i", world.Banks[0])

	logger.Info("Generate currency on bank accounts")
	logger.Info("Generate production types")
	logger.Info("Set creation costs for production types")
	logger.Info("*End*")

	return &world
}
