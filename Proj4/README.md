## To start app
`go run server.go`
### Example endpoint tests:
#### List all endpoints:
```curl localhost:1323/ | jq```
#### Simple object:
```curl -X POST localhost:1323/worker  -H 'Content-Type: application/json' -d '{"Firstname": "Janusz", "Surname": "Nowak", "Position": "Boss"}'```

```curl localhost:1323/worker | jq```
#### Composite object:
```curl -X POST localhost:1323/shift -H 'Content-Type: application/json' -d '{"EnterTime": "2022-05-23 09:00",  "ExitTime": "2022-05-22 17:30", "WorkerID": 1}'```

```curl localhost:1323/shift | jq```

```curl -X POST localhost:1323/purchase  -H 'Content-Type: application/json' -d '{"CustomerID": 1, "ProductID":1}'```