# Elsa Coding Challenge

## Architecture Diagram

View on [PlantUML Web Editor](https://www.plantuml.com/plantuml/uml/bLDFSzew4B_xK-pZ4fXnyEMTtYa4CadU8mala6PwOcHh0JMov4hhCBJJxrwjOWmEqtPoO93N-ktlpwulF0b7TQcJlvGfT2qHDaIL_oVBdDWEruesTLvxT8KrX8Q6XIspIWl3c7DFKgiy6_zzFkFIqzqqAuKdTA6q7DkoieOn5VS_Wze4S1aAG1k4bTNQRfLPWrO6GHW9dTVA-N2o4CX1sq9elGizMfMd1CbqzFdXQR5ylbtSBATNrrVta_uWIKYHHcZlWbHYxKGP-1qyej3d2rKYV0eCUQs5sy7_jVe6ewhIgX2ah4cISV0kWaZNBxGAFLBeZUC1-i-OGyQt1xrGvT_CsLSbqSCg-6Myq8fsxEIBLIO54YyO7Nvb7hRzgd0RAxwm3bFGA2Iwt0edq-5ms1iS14polim3oKIGo8L79hopPcqdbtoSatM1cUVbqqZWGPcLBQEHD9A42xKc3trX_3RUYxG3iAit82VJ2lJls1rAvTc_A3QH_P4BU-u6qx4G1OGkXpRBIriR6QBkYqe5v6XsUvntRm3V4m0upAmlQjhmnV07Do4oZOufhvKvVXpnGZKUo89NQnF8-TZ6xNUUiEH8kmJ7aKDlUdZw4vq_yAsHm0dpmZDjHid4dIoa877Apj92GrwyZ_6rnbBmTLugWdQeXT15hGLX19pvDqLHMrr9hfCzbiQgVYJ98-fsoPlKvqsE38sSzWKD2_ft_d2VtImMitd07WCRFzT8pJvqrVqMsva4jpZPfNJlg1lsILTksqIxtzgHbDquU4si8mmzaDfrXldalz3W1Czoid3nwClwDKRcxo8W9cJs_P02DdwxrcEZjcC-aduqWwwQ2pIoBlLF)

## Prerequisites

- Having `redis` installed with version >= 7.2.0, or having `docker` installed
- If you have `docker` installed, there's a `docker-compose.yml` file already, run this command at repository level to start redis container:

```
docker-compose up -d
```

## Run the project

- Run the `main.go` file in `cmd/script` to prepare and add data to redis
- Run the `main.go` file in `cmd/http` to start http server
- There're two API that you can test with the http server

### Get the leaderboard

```
curl --location 'http://localhost:9000/v1/leaderboard/test-key'
```

### Get rank and score of an user

```
curl --location 'http://localhost:9000/v1/leaderboard/test-key/user-rank?userId=user-100'
```
