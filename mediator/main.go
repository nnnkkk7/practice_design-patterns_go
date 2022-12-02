package main

func main() {
	stationManeger := newStationManger()
	passengerTrain := &PassengerTrain{
		mediator: stationManeger,
	}
	freightTrain := &FreightTrain{
		mediator: stationManeger,
	}
	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.arrive()

}
