package consul

import (
	"log"

	api "github.com/hashicorp/consul/api"
)

func GetMuse() []string {
	// Get a new client
	client, err := api.NewClient(
		&api.Config{
			Address: "consul:8500",
			Scheme:  "http",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	catalog := client.Catalog()

	node, _, err := catalog.Node("muse", nil)
	if err != nil {
		// Add a new default service
		service := &api.AgentService{
			ID:      "calliope",
			Service: "calliope",
			Tags:    []string{"muse", "v1"},
			Port:    9090,
		}

		// Init the catalog
		reg := &api.CatalogRegistration{
			Datacenter: "dc1",
			Node:       "muse",
			Address:    "calliope",
			Service:    service,
		}
		if _, err := catalog.Register(reg, nil); err != nil {
			log.Fatal(err)
		}
	}

	if _, ok := node.Services["calliope"]; ok {
		log.Fatal("ServiceID:calliope is not registered")
	}

	// TODO: return all the services
	defSvc := node.Services["calliope"].Address + ":" + string(node.Services["calliope"].Port)
	return []string{defSvc}
}
