package controller

import (
	"github.com/sintell/wsgame/connection"
	"github.com/sintell/wsgame/model"
	"github.com/sintell/wsgame/utils"
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
    world = model.
}

func Process(message *connection.Message) {

}
