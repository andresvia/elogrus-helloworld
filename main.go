package main

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/olivere/elastic.v3"
	"gopkg.in/sohlich/elogrus.v1"
	"os"
	"time"
)

func main() {

	t := time.Tick(2 * time.Second)

	println(os.Getenv("ELASTICSEARCH_URL"))
	client, err := elastic.NewClient(elastic.SetURL(os.Getenv("ELASTICSEARCH_URL")))
	if err != nil {
		panic("error elasticsearch " + err.Error())
	}
	println("cliente elastic creado")

	hook, err := elogrus.NewElasticHook(client, "loghost", logrus.DebugLevel, "logstash")
	if err != nil {
		panic("error elogrus" + err.Error())
	}
	println("hook elogrus creado")

	logrus.AddHook(hook)

	for range t {
		now := time.Now().String()
		logrus.WithFields(logrus.Fields{
			"now": now,
		}).Info("Now!")
	}
}
