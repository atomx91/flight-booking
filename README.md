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


## APIs

### Client

- Located in folder `/client`
- Restful API served:

POST `/client` - register new client

PUT `/client/:username` - update client data by id

POST `/client/viewBookingHistory` - update client data

POST `/client/searchBooking` - Search Booking data

POST `/client/changePassword` - Change Password

- gRPC served:

Same with rest api

### Flight

- Located in folder `/flight`
- Restful API served:

POST `/flight` - Create Flight

GET `/flight/:id` - Get flight by id

PUT `/flight/:id` - Update Flight



- gRPC served:

Same with rest api

### Booking

- Located in folder `booking`
- Restful API served:

POST `/booking` - Create Booking

GET `/booking/guest` - Get list of user's reserved bookings

GET `/booking/cancel` - Cancel booking

- gRPC served:

Same with rest api