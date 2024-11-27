## Stock-ticker-service

### To run locally

- All Docker packaging and publish related code is done in the Kotlin repository. Please refer to section `To run locally` at https://github.com/shaz-sk/stock-ticker-service. It has instructions to pull or build the container and test.
  - Docker related files in that repository are Dockerfile and build_and_publish.sh

- To clone the Golang repository from github. 
  ```shell
  git clone https://github.com/shaz-sk/stock-ticker-go-service.git 
  cd stock-ticker-go-service
  go run main.go
  ``` 
- For API calls, in postman.
   ```  
   http://localhost:9090/api/v1/stock/report
   ```

### Information about code and features
- Go service with handlers and service layer pattern.
- stock SYMBOL, NDAYS and APIKEY as environment variables.
- Integration test using MockMvc and Wiremock. https://github.com/shaz-sk/stock-ticker-service/tree/master/src/test/kotlin/com/organization/stockTicker/integrationTests
- Open API spec is available ./openapi.yaml.
- Unit test for StockTickerService to show intent.

- Additional features implemented only in the kotlin repository are
  - Docker image that runs the microservice.
  - Docker build and publish script.
  - Circuit breaker to handle third party API errors by providing failure protection.
  - Actuator for health check endpoints.
  - Boilerplate code generation. API endpoints are generated during build using openapi code generator.
  - Reference : https://github.com/shaz-sk/stock-ticker-service


### Code details
- main.go --> handler.StockTickerHandler --> service.StockTickerService --> service.AlphaVantageClient --> service.StockTickerService --> mapper.StockDetailsMapper --> handler.StockTickerHandler
- StockTickerHandler - handles HTTP requests related to stock ticker information
- StockTickerServiceImpl - handles the business logic for retrieving and processing stock ticker data
- AlphaVantageClientImpl - calls Alpha Vantage API to retrieve stock data
- StockDetailsMapper - maps information received from Alpha Vantage to API response
- EnvConfig - holds all the configuration settings for environment variables
