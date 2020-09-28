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

* IoC
* unit tests
* integration tests
* Http resiliency



## Backend

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