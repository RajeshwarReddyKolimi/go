# Car Rental System

## Steps followed

- Created a folder `carrentalsystem` with folders:
  - models: To define structs like Reservation, User, etc.
  - usecases: Based on particular usecase, classified functionalities in different folders. Contains interfaces and their implementations.
    - car: Functionalities related to car alone like IsAvailable between given time.
    - crs: Functionalities which include users, cars, bookings and overall related to booking a car.
  - utils: Utility functions like parsing, calculating.
    and files:
  - main.go: The entry point of function where the required functionalities will be called from.
  - go.mod: Describes module's properties created using `go mod init`
