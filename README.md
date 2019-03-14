# iPA External Adapter

![iPA logo](logo_ipa.png)

External Adapter for Chainlink to allow access to the Italian "Indice delle Pubbliche Amministrazioni" (iPA).

Derived from [one](https://github.com/linkpoolio/alpha-vantage-cl-ea) of [LinkPool](https://linkpool.io/)'s adapters.

### Preconditions

Retrieve an API key from the iPA website for free:

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

### Usage

To call the API, you need to send a POST request to `http://localhost:<port>/query` with the request body being of the ChainLink `RunResult` type.

The `data` passed in should match the parameter options from the iPA docs.

For example,
```
curl -X POST -H 'Content-Type: application/json' -d '{ "jobRunId": "1234", "data": { "CF": "97735020584" }}' http://localhost:8080/query
```
should return:
```json
{
  "jobRunId": "",
  "status": "",
  "error": null,
  "pending": false,
  "data": {
    "data": [
      {
        "cod_amm": "agid",
        "des_amm": "Agenzia per L'Italia Digitale",
        "OU": [
          {
            "des_ou": "Uff_eFatturaPA",
            "stato_canale": "A",
            "cod_uni_ou": "UF4NU9"
          },
          {
            "des_ou": "Organizzazione e Gestione del Personale",
            "stato_canale": "A",
            "cod_uni_ou": "ITLER2"
          },
          {
            "des_ou": "Contabilita' Finanza e Funzionamento",
            "stato_canale": "A",
            "cod_uni_ou": "F7VRDL"
          }
        ]
      }
    ],
    "result": {
      "cod_err": 0,
      "desc_err": "",
      "num_items": 3
    }
  }
}
```
