package main

import (
	"flag"
	"github.com/janearc/sux/backend"
	"github.com/janearc/sux/sux"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	// parse command line args
	query := flag.String("query", "what is the funniest joke you've heard today", "query to pass to the backend")
	flag.Parse()

	// create the new service object
	S := sux.NewSux("config/config.yml", "config/version.yml", "config/secrets.yml")
	if S != nil {
		log.Infof("Sux [%s] instantiated", S.GetVersionBuild())
	} else {
		log.Fatalf("Failed to instantiate Sux")
	}

	openai := backend.NewOpenAITransport(S.GetConfig())
	if openai != nil {
		log.Infof("OpenAI transport instantiated")
	} else {
		log.Fatalf("Failed to instantiate OpenAI transport")
	}

	S.AddBackend("OpenAI", openai)

	log.Info("Sux instantiated, backend created, backend added to service.")
	log.Infof("Query: [%s]", *query)

	request, err := openai.OpenAIRequest(*query)
	if err != nil {
		log.Fatalf("OpenAI request failed with error: %v", err)
	}
	log.Infof("OpenAI response: [%s]", request)
}
