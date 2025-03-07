package main

import (
	"fmt"
	"sort"
)

func main() {
	r := []string{"EnRouteExtension",
		"CruisingLevelChangeExtension",
		"FlightRouteInformationExtension",
		"IataFlightDesignatorExtension",
		"ActualTimeOfArrivalExtension",
		"RouteTrajectoryGroupConatinerExtension",
		"FlightIdentificationExtension",
		"PlannedDelayExtension",
		"SpeedScheduleExtension",
		"FlightCapabilitiesExtension",
		"CruiseClimbStartExtension",
		"TrajectoryPoint4DExtension",
		"PersonOrOrganizationExtension",
		"SurvivalCapabilitiesExtension",
		"CruisingSpeedChangeExtension",
		"LevelConstraintExtension",
		"RouteTrajectoryConstraintExtension",
		"RouteTrajectoryElementExtension",
		"AircraftTypeExtension",
		"ActualTimeOfDepartureExtension",
		"AircraftExtension",
		"FlightExtension",
		"ArrivalExtension",
		"SpeedConstraintExtension",
		"RouteTrajectoryGroupExtension",
		"DepartureExtension",
		"TrajectoryPointPropertyExtension",
		"ProfilePointExtension",
		"TimeConstraintExtension",
		"RouteChangeExtension",
		"PerformanceProfileExtension"}

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	fmt.Println(r)

}
