# iPA External Adapter

External Adapter for Chainlink to allow access to the Italian "Indice delle Pubbliche Amministrazioni".

Derived from [one](https://github.com/linkpoolio/alpha-vantage-cl-ea) of [LinkPool](https://linkpool.io/)'s adapters.

### Preconditions

Retrieve an API key from Alpha Vantage for free:

https://www.indicepa.gov.it/documentale/n-webservices.php

### Setup Instructions

#### Local Install
Make sure [Golang](https://golang.org/pkg/) is installed.

Build:
```
make build
```

Then run the adaptor:
```
./ipa-cl-ea -p <port> -apiKey <API_KEY>
```

##### Arguments

| Char   | Default  | Usage |
| ------ |:--------:| ----- |
| p      | 8080     | Port number to serve |
| apiKey | nil      | Your API Key for iPA |

#### Docker
To run the container:
```
docker run -it -p 8080:8080 reale/ipa-cl-ea -apiKey=yourkey
```

Container also supports passing in CLI arguments.
