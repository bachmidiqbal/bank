# bank

# How to run

> go run main.go

# How to use this service

## APIs

## create account

request

> curl -X POST http://localhost:8000/create-account \
   -H 'Content-Type: application/json' \
   -d '{"name":"Iqbal","address":"Indonesia","phone":"+6211223344"}'

response

```
{
    "name": "Iqbal",
    "address": "Indonesia",
    "phone": "+6211223344",
    "Number": 1000,
    "Balance": 0
}
```

## deposit

request

> curl -X POST http://localhost:8000/deposit \
   -H 'Content-Type: application/json' \
   -d '{"number":"1000","amount":"50"}'

response

```
{
    "name": "Iqbal",
    "address": "Indonesia",
    "phone": "+6211223344",
    "Number": 1000,
    "Balance": 50
}
```

## withdraw

request

> curl -X POST http://localhost:8000/withdraw \
   -H 'Content-Type: application/json' \
   -d '{"number":"1000","amount":"10"}'

response

```
{
    "name": "Iqbal",
    "address": "Indonesia",
    "phone": "+6211223344",
    "Number": 1000,
    "Balance": 40
}
```

## statement

request

> curl "http://localhost:8000/statement?number=1000"

response

```
{
    "name": "Iqbal",
    "address": "Indonesia",
    "phone": "+6211223344",
    "Number": 1000,
    "Balance": 40
}
```

## transfer

request

> curl -X POST http://localhost:8000/transfer \
   -H 'Content-Type: application/json' \
   -d '{"from":"1000","to":"1001","amount":"20"}'

response

```
[
    {
        "name": "Iqbal",
        "address": "Indonesia",
        "phone": "+6211223344",
        "Number": 1000,
        "Balance": 20
    },
    {
        "name": "Nabila",
        "address": "Indonesia",
        "phone": "+6233445566",
        "Number": 1001,
        "Balance": 20
    }
]
```
