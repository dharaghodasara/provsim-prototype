# Provider Simulator - design specificaiton 
- Provider 
    - has basic set of fields 
    - Provider has set of capabilities 
- Capabilities 
    - can be added or removed thru simple configuration files 
    - are considered as "addons" to a provider. 
    - exposes `DataGeneerator` interface 
        - `DataGenerator` takes one paramenter `scenario` and returns `JSON node`
        - Quesions: 
            1. How do we handle the user associations with the data?  
            2. Should the original data be cloned and stored against the users? 
            3. If we stored against the users, when a request comes for a existing_user, pick the data from the user and manupulate else pick from the template and persist for the users. 
    - listens to events 
        - `onEvent` they react and change state of the data 
- CapabilityDecorator



