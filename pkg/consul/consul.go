package consul

import (
	"log"

	api "github.com/hashicorp/consul/api"
)

func GetMuse() []string {
	// Get a new client with custom config
	conf := api.DefaultConfig()
	conf.Address = "consul:8500"

	client, err := api.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	catalog := client.Catalog()

	meta := map[string]string{"muse": "muse"}
	services, _, err := catalog.Services(
		&api.QueryOptions{
			NodeMeta: meta,
		},
	)
	if err != nil {
		log.Fatal("consul error: " + err.Error())
	}

	// TODO: return all the services endpoints
	if _, ok := services["calliope"]; ok {
		log.Println("return calliope service address")
		s, _, err := catalog.Service("calliope", "", nil)
		if err != nil {
			log.Fatal("consul error: " + err.Error())
		}
		svcURL := s[0].Address + ":" + string(s[0].ServicePort)
		return []string{svcURL}
	}

	//return default value
	return []string{"http://calliope:9090"}
}
