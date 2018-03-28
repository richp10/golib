package env

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/armon/consul-api"
	"github.com/asaskevich/govalidator"
)

// Watches a Consul *Node*
func ConsulWatch(nodename string, key string, schema string) {
	// Get a new consul client
	if key == "" {
		key = nodename
	}

	consul, _ := consulapi.NewClient(consulapi.DefaultConfig())
	health, _, err := consul.Health().Node(nodename, nil)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	// && health[0].Status == "passing"
	if len(health) > 0 {
		// open a goroutine to watch remote changes forever
		go func() {
			ip := ""
			for {
				//println("Check: " + nodename)
				node, _, err := consul.Catalog().Node(nodename, nil)
				if err != nil {
					time.Sleep(time.Second * 5) // delay after fail..
					continue
				}
				newip := node.Node.Address

				if ip != newip {
					// First set the requested environment setting using the ip and optional schema
					log.Info("SET KEY: "+key+" TO: " + schema+newip)
					viper.Set(key, schema+newip)
					ip = newip
					// Create a flag that we can use to check whether
					// the node IP address has been changed
					viper.Set(nodename+"-ipchanged", true)
				}
				time.Sleep(time.Second * 10) // delay after each request
			}
		}()

	} else {
		log.Fatalf("No docker container: " + nodename)
	}

}

// Watches using a consul KV Key
func ConsulWatchIpKey(nodename string, key string, schema string) {
	// Get a new consul client
	if key == "" {
		key = nodename
	}

	consul, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	// open a goroutine to watch remote changes forever
	go func() {
		ip := ""
		for {
			kvpair, _, err := consul.KV().Get("ip-store/"+nodename, nil)
			if err != nil {
				log.Info("Consul Get KV Failed: ", err)
				time.Sleep(time.Second * 10) // delay after fail..
				continue
			}

			if !govalidator.IsIP(string(kvpair.Value)) {
				log.Info("Consul Get KV IP Invalid")
				time.Sleep(time.Second * 10) // delay after fail..
				continue
			}

			newip := string(kvpair.Value)
			if ip != newip {
				// First set the requested environment setting using the ip and optional schema
				log.Info("SET KEY: " + key + " TO: " + schema + newip)
				viper.Set(key, schema+newip)
				ip = newip
				// Create a flag that we can use to check whether
				// the node IP address has been changed
				viper.Set(nodename+"-ipchanged", true)
			}
			time.Sleep(time.Second * 10) // delay after each request
		}
	}()

}
