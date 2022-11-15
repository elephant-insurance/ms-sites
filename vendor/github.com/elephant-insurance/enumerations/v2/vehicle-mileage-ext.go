package enumerations

// This file contains NON-AUTO-GENERATED code for the Vehicle Mileage enumeration
// This makes it possible to re-generate the enumeration without losing the features here

const (
	fourThousand    int = 4000
	sixThousand     int = 6000
	eightThousand   int = 8000
	tenThousand     int = 10000
	twelveThousand  int = 12000
	fifteenThousand int = 15000
	twentyThousand  int = 20000
)

// FromMileageInt returns a vehicle mileage ID for an integer value
// This makes it possible to turn an arbitrary number into a range value
func (e *EnumVehicleMileage) FromMileageInt(mileage *int) *VehicleMileageID {
	if mileage == nil {
		return nil
	}

	switch {
	case *mileage < fourThousand:
		return e.LessThan4000.ID.Clone()
	case *mileage < sixThousand:
		return e.From4000To5999.ID.Clone()
	case *mileage < eightThousand:
		return e.From6000To7999.ID.Clone()
	case *mileage < tenThousand:
		return e.From8000To9999.ID.Clone()
	case *mileage < twelveThousand:
		return e.From10000To11999.ID.Clone()
	case *mileage < fifteenThousand:
		return e.From12000To14999.ID.Clone()
	case *mileage < twentyThousand:
		return e.From15000To19999.ID.Clone()
	default:
		return e.MoreThan20000.ID.Clone()
	}
}
