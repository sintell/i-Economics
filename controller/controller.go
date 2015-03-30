package controller

import (
<<<<<<< HEAD
	"github.com/sintell/i-Economics/connection"
	"github.com/sintell/i-Economics/model"
	"github.com/sintell/i-Economics/utils"
=======
	"github.com/sintell/wsgame/connection"
	"github.com/sintell/wsgame/model"
	"github.com/sintell/wsgame/utils"
>>>>>>> 11847977dfa6b15a112f4ed4dc378f42a5e36670
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
<<<<<<< HEAD
	world := model.NewWorld()
	worldId := world.MetaInfo.Id
	logger.Info("Creating new world [%s]", worldId)

	worlds[worldId] = world
	return worldId
=======
    world = model.
>>>>>>> 11847977dfa6b15a112f4ed4dc378f42a5e36670
}

func Process(message *connection.Message) {

}
