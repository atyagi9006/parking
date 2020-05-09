# Parking

Parking is app which simulate a parking lot. It creates,allocates, release and provide status of paking. It also provides different type of lookups (i.e  regno for a colour, slotno for a color, slotno for a regno) 

## Table of Contents

  - [Installation](#install)
  - [Clean](#clean)
  - [Test](#test)
  - [Run](#run)
  - [Structure](#structure)

### Installation <a name="install"></a>

Clone this repo. And run:

```sh
make all
```

This command run everything in single shot (i.e clean, test, lint, run). To run each step seperatly execute the commands below one by one.

### Clean <a name="clean"></a>

To clean up all the generated artifacts, run:

```sh
make clean
```

### Test <a name="test"></a>

To test coverage, run:

```sh
make test
```

To see test-coverage details check test-coverage.html 

### Run <a name="test"></a>

To run app, run:

```sh
make run
```

This command runs app and prints out commentry and final result on commandline.

## Structure <a name="structure"></a>

```
.
├── Makefile
├── README.md
├── car
│   ├── car.go
│   └── car_test.go
├── cmd
│   ├── Command.go
│   ├── CommandCreateParkingLot.go
│   ├── CommandCreateParkingLot_test.go
│   ├── CommandLeave.go
│   ├── CommandPark.go
│   ├── CommandPark_test.go
│   ├── CommandRegistrationNumber.go
│   ├── CommandSlotNumberCarColor.go
│   ├── CommandSlotNumberCarNumber.go
│   ├── CommandStatus.go
│   └── Command_test.go
├── go.mod
├── input.txt
├── lot
│   ├── parking.go
│   └── parking_test.go
├── main.go
└── slot
    ├── slot.go
    └── slot_test.go
```
