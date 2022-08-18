package app

import (
	"github.com/adelmar92/GoApiClean/src/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
	"log"
)

const port = "8080"

func Init() {
	r := gin.New()
	handlers := dependencies.Start()
	ConfigureMappings(r, handlers)
	logEnvironment()
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("ERROR : Cannot Run Gin Framework " + err.Error())
	}
}

func logEnvironment() {
	log.Println("INFO : Starting Server on Port :" + port)
}
