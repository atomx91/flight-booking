# Flight booking

## Database
- Run the queries in _flight_booking.sql_
- Change database connection information in _config.yml_

## Generate *.pb.go file

`protoc -Iproto --go_out=paths=source_relative:./pb --go-grpc_out=paths=source_relative:./pb proto/*.proto`

## Server 
### gRPC Apis (port 4000)
`go run main.go`

### REST API (port 5000)
`go run main.go`


