package data

import (
	"fmt"
	"log"
	"time"

	"github.com/elangovans/provsim-prototype/core"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

//ConfigMap ...
var ConfigMap = make(map[uuid.UUID]core.Provider)

/*
LoadSimulationConfigurations ...

*/
func LoadSimulationConfigurations() {
	var pr core.Provider
	b := time.Now()
	i := 0

	for _, config := range Configs {
		err := yaml.Unmarshal([]byte(config), &pr)
		if err != nil {
			log.Fatalf("cannot unmarshal data: %v", err)
		}

		pr.ID, err = uuid.NewUUID()
		if err != nil {
			log.Fatalf("cannot create a new UUID: %v", err)
		} else {
			/**
			Just print one provider config
			id for sameple
			**/
			if 0 == i {
				pr.Name = pr.Name + pr.ID.String()
				fmt.Println(pr.ID)
				i++
			}

			ConfigMap[pr.ID] = pr
		}
	}

	e := time.Now()
	elapsed := e.Sub(b)

	fmt.Printf("started at %s ended at %s, so the elapsed is %d to load number of %d configs\n",
		b, e, elapsed/1000000, len(Configs))
}

//Configs ...
var Configs = []string{`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`
name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
	`name: Chase-Simulator
id: efd958bc-2c65-4462-b774-b12d8ef97905
base-url: https://chase-sim.intuit.com 
login-url: https://chase-sim.intuit.com/login
provider-category: banking 
meta-data:
  base-provider: 
    name: Chase Bank
    id: 2c064f98-91a1-4d4b-9ab1-c5ea4d81bfdc
  cloned-from: ## if this was cloned from another provider and just for a reference 
    id: e7b92147-6fa8-4307-9d62-c0d9164774f8
capabilities: 
  - name: Savings Bank ## lookup capabilities map for entities supported if wants to inherit from the base template 
    type: Service
  - name: Checking Bank
    type: Service
    attributes:
      - ["pendig-tranaction", "include"]
      - ["something-else", "maybe"]
  - type: Security 
    name: oatuh2.0 
  - type: DataProvider 
    name: Chase-test-data-generator
    attributes: # name value pair, something like that? 
      - name: good-data
        weightage: 100% 
`,
}
