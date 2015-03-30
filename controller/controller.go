package controller

import (
	"github.com/sintell/i-Economics/connection"
	"github.com/sintell/i-Economics/model"
	"github.com/sintell/i-Economics/utils"
	"log"
	"os"
)

var worlds map[model.WorldId]*model.World
var logger *utils.Logger

func init() {
	logger = utils.NewLogger(utils.DEBUG, os.Stdout, log.Ldate|log.Ltime|log.Lmicroseconds)
	worlds = make(map[model.WorldId]*model.World)
}

func GenerateWorld() model.WorldId {
	world := model.NewWorld()
	worldId := world.MetaInfo.Id
	logger.Info("Creating new world [%s]", worldId)

	worlds[worldId] = world
	return worldId
}

func Process(message *connection.Message) {

}
