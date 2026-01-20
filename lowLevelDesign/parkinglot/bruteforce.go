// LLD: Parking Lot - BRUTE FORCE Approach
// Time Complexity: O(levels * spots) for parking
// Simple and easy to understand - good starting point in interviews

package main

import (
	"errors"
	"fmt"
)

// ******** ENTITIES ********

type VehicleTypeBF int

const (
	CAR_BF VehicleTypeBF = iota
	BIKE_BF
)

type VehicleBF struct {
	Type VehicleTypeBF
	Reg  string
}

// Simple Spot - just knows its type and if occupied
type SpotBF struct {
	ID          int
	VehicleType VehicleTypeBF
	IsOccupied  bool
	Vehicle     *VehicleBF
}

// Level with simple slice of spots
type LevelBF struct {
	ID    int
	Spots []*SpotBF // Simple slice - no fancy data structures
}

// ParkingLot
type ParkingLotBF struct {
	Levels []*LevelBF
}

// ******** FACTORY FUNCTIONS ********

func NewSpotBF(id int, vType VehicleTypeBF) *SpotBF {
	return &SpotBF{ID: id, VehicleType: vType}
}

func NewLevelBF(id int, carSpots, bikeSpots int) *LevelBF {
	l := &LevelBF{
		ID:    id,
		Spots: make([]*SpotBF, 0, carSpots+bikeSpots),
	}

	// Add car spots
	for i := 1; i <= carSpots; i++ {
		l.Spots = append(l.Spots, NewSpotBF(i, CAR_BF))
	}

	// Add bike spots
	for i := carSpots + 1; i <= carSpots+bikeSpots; i++ {
		l.Spots = append(l.Spots, NewSpotBF(i, BIKE_BF))
	}

	return l
}

func NewParkingLotBF() *ParkingLotBF {
	return &ParkingLotBF{}
}

// ******** CORE METHODS ********

func (p *ParkingLotBF) AddLevel(level *LevelBF) {
	p.Levels = append(p.Levels, level)
}

// Park - BRUTE FORCE: O(levels * spots)
// Simply loop through all levels and all spots to find a free one
func (p *ParkingLotBF) Park(v *VehicleBF) (int, int, error) {
	// Loop through each level
	for _, level := range p.Levels {
		// Loop through each spot on this level
		for _, spot := range level.Spots {
			// Check if spot matches vehicle type AND is free
			if spot.VehicleType == v.Type && !spot.IsOccupied {
				// Found a free spot - park here
				spot.IsOccupied = true
				spot.Vehicle = v
				return level.ID, spot.ID, nil
			}
		}
	}
	return -1, -1, errors.New("no available spot for vehicle type")
}

// Unpark - O(levels * spots) to find the spot
func (p *ParkingLotBF) Unpark(levelID, spotID int) error {
	// Find the level
	for _, level := range p.Levels {
		if level.ID == levelID {
			// Find the spot
			for _, spot := range level.Spots {
				if spot.ID == spotID {
					if !spot.IsOccupied {
						return errors.New("spot already empty")
					}
					spot.IsOccupied = false
					spot.Vehicle = nil
					return nil
				}
			}
			return errors.New("invalid spot ID")
		}
	}
	return errors.New("invalid level ID")
}

// DisplayAvailability - O(levels * spots) to count free spots
func (p *ParkingLotBF) DisplayAvailability() {
	fmt.Println("----- Parking Availability (Brute Force) -----")
	for _, level := range p.Levels {
		carFree, bikeFree := 0, 0

		// Count free spots by looping through all
		for _, spot := range level.Spots {
			if !spot.IsOccupied {
				if spot.VehicleType == CAR_BF {
					carFree++
				} else {
					bikeFree++
				}
			}
		}

		fmt.Printf("Level %d: CAR=%d, BIKE=%d\n", level.ID, carFree, bikeFree)
	}
	fmt.Println("-----------------------------------------------")
}

// ******** DEMO ********

func mainBruteForce() {
	// Create parking lot
	lot := NewParkingLotBF()

	// Add levels
	lot.AddLevel(NewLevelBF(1, 3, 3)) // Level 1: 3 CAR, 3 BIKE spots
	lot.AddLevel(NewLevelBF(2, 2, 3)) // Level 2: 2 CAR, 3 BIKE spots

	lot.DisplayAvailability()

	// Park vehicles
	car := &VehicleBF{Type: CAR_BF, Reg: "CAR123"}
	bike := &VehicleBF{Type: BIKE_BF, Reg: "BIKE456"}

	vehicles := []*VehicleBF{car, car, car, car, bike}

	for _, v := range vehicles {
		levelID, spotID, err := lot.Park(v)
		if err != nil {
			fmt.Printf("Failed to park %s: %v\n", v.Reg, err)
		} else {
			fmt.Printf("Parked %s at Level %d, Spot %d\n", v.Reg, levelID, spotID)
		}
	}

	lot.DisplayAvailability()

	// Unpark
	err := lot.Unpark(1, 1)
	if err != nil {
		fmt.Printf("Failed to unpark: %v\n", err)
	} else {
		fmt.Println("Unparked vehicle from Level 1, Spot 1")
	}

	lot.DisplayAvailability()
}

/*
COMPARISON: Brute Force vs Optimized

┌─────────────────────┬─────────────────────┬─────────────────────┐
│ Operation           │ Brute Force         │ Optimized (Sets)    │
├─────────────────────┼─────────────────────┼─────────────────────┤
│ Park                │ O(L × S)            │ O(L) or O(1)*       │
│ Unpark              │ O(L × S)            │ O(L) or O(1)*       │
│ Check Availability  │ O(L × S)            │ O(1)                │
│ Memory              │ O(S)                │ O(S) + overhead     │
│ Code Complexity     │ Simple ✅           │ More complex        │
└─────────────────────┴─────────────────────┴─────────────────────┘

L = number of levels, S = spots per level
* O(1) if you also index levels by ID

WHEN TO USE BRUTE FORCE:
- Small parking lots (< 100 spots)
- Interview starting point (then optimize)
- Prototyping / MVP
- When simplicity > performance

WHEN TO OPTIMIZE:
- Large scale (1000+ spots)
- High throughput requirements
- Real-time systems
*/
