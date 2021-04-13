## API methods

* List running environments
* Create environment
* Delete environment
* Get environment -> link to consul etc.
* List templates
* Get template
* Delete template
* Import environment(?)
* Update environment
* Update template

## To do

* Connect to consul
  * K / V
  * Connect
  * Service registration
* Vault for secrets + login
* User login
* Tracing
* Metrics
* Logs
* Api Docs
* Handle proxies

* IoC
* unit tests
* integration tests
* Http resiliency

* Data storage



## Backend

This will need its own data store. The data store will be async updated based on update
messages from other services

### Naming

* https://stackoverflow.com/questions/38842457/interface-naming-convention-golang

### Dependency injection

* dig
* inject
* wire

* Opinions: https://www.reddit.com/r/golang/comments/8jqjx8/dependency_injection_in_go/
* one way: https://www.elliotdwright.com/2018/02/27/how-i-structure-some-of-my-projects
* https://www.reddit.com/r/golang/comments/8jqjx8/dependency_injection_in_go/dz4hdi3/
* http://www.jerf.org/iri/post/2929
* https://scene-si.org/2016/06/16/dependency-injection-patterns-in-go/
* https://scene-si.org/2016/07/07/dependency-injection-continued/
* https://ieftimov.com/post/testing-in-go-dependency-injection/

### API documentation

* OpenApi and go
  * http-swagger and chi: https://github.com/swaggo/http-swagger/blob/master/example/go-chi/main.go#L1-L31
  * https://github.com/swaggo/swag
  * https://www.ribice.ba/swagger-golang/
  * https://webhookrelay.com/blog/2018/11/05/openapi-redoc-tutorial/
  * https://idratherbewriting.com/learnapidoc/docapis_introtoapis.html
  * https://mux.com/blog/an-adventure-in-openapi-v3-api-code-generation/
  * https://b2b-api.test.etraveli.com/docs/b2b/v1/#tag/Search
  * https://github.com/Redocly/redoc

## Front end

## Data storage and friends

Things we need to keep into account when deciding the data storage

* How are we going to keep the data consistent with the actual infrastructure. Infrastructure can change, either because somebody click-op-sed
  or because something failed
* Need to somehow ‘know’ what is in an environment, what is expected and what isn’t. Can keep track of those parts.
  * Storing
    * Environment
      * Resources
      * Tags
      * Name
      * Description
      * Entrypoint(s)
      * Status endpoints
    * Resource
      * Tags
      * Name
      * Description
      * Status endpoints
      * Dependencies -> Resource IDs / External resources
  * Do we need to store a dependency graph for quick retrieval?
* How are we going to deal with search?
* Database
  * Prefer faster reads over faster writes – We won’t be doing many writes compared to the number of reads we will do
  * How much data do we store – Not very much, definitely less than a Gb
  * How many users do we have – Not very many. Kinda 1 to start with
  * Data
    * Some is nested – Tags / Environment -> Resource links
    * Some of the data forms a graph – Resource dependencies
    * Some of the data needs to be searchable - All text fields like names and descriptions. Additionally also might want to
      search the dependency tree
    * Do we care about multi-master / geographic distribution? Probably not at the start
* Keep the data for a service together. No other services should touch this data directly
* Use events to communicate with other services
  * How are we going to report progress etc.

Also need

* Some way of keeping things up to date. Ideally we would get notified when things change, but we may have to poll. Can we
  link to Consul and keep track of the health status?