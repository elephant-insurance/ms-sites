package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// StateID uniquely identifies a particular State
type StateID string

// Clone creates a safe, independent copy of a StateID
func (i *StateID) Clone() *StateID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two StateIds are equivalent
func (i *StateID) Equals(j *StateID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *StateID that is either valid or nil
func (i *StateID) ID() *StateID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *StateID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the StateID corresponds to a recognized State
func (i *StateID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return State.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *StateID) ValidatedID() *ValidatedStateID {
	if i != nil {
		return &ValidatedStateID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *StateID) MarshalJSON() ([]byte, error) {
	if i == nil || *i == "" {
		return []byte("null"), nil
	}

	if !i.Valid() {
		err := errors.New(errorMarshalInvalidID)
		return nil, err
	}

	istring := string(*i)

	return []byte(`"` + istring + `"`), nil
}

func (i *StateID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := StateID(dataString)
	item := State.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	stateAlaskaID             StateID = "AK"
	stateAlabamaID            StateID = "AL"
	stateArkansasID           StateID = "AR"
	stateArizonaID            StateID = "AZ"
	stateCaliforniaID         StateID = "CA"
	stateColoradoID           StateID = "CO"
	stateConnecticutID        StateID = "CT"
	stateDistrictOfColumbiaID StateID = "DC"
	stateDelawareID           StateID = "DE"
	stateFloridaID            StateID = "FL"
	stateGeorgiaID            StateID = "GA"
	stateHawaiiID             StateID = "HI"
	stateIowaID               StateID = "IA"
	stateIdahoID              StateID = "ID"
	stateIllinoisID           StateID = "IL"
	stateIndianaID            StateID = "IN"
	stateKansasID             StateID = "KS"
	stateKentuckyID           StateID = "KY"
	stateLouisianaID          StateID = "LA"
	stateMassachusettsID      StateID = "MA"
	stateMarylandID           StateID = "MD"
	stateMaineID              StateID = "ME"
	stateMichiganID           StateID = "MI"
	stateMinnesotaID          StateID = "MN"
	stateMissouriID           StateID = "MO"
	stateMississippiID        StateID = "MS"
	stateMontanaID            StateID = "MT"
	stateNorthCarolinaID      StateID = "NC"
	stateNorthDakotaID        StateID = "ND"
	stateNebraskaID           StateID = "NE"
	stateNewHampshireID       StateID = "NH"
	stateNewJerseyID          StateID = "NJ"
	stateNewMexicoID          StateID = "NM"
	stateNevadaID             StateID = "NV"
	stateNewYorkID            StateID = "NY"
	stateOhioID               StateID = "OH"
	stateOklahomaID           StateID = "OK"
	stateOregonID             StateID = "OR"
	statePennsylvaniaID       StateID = "PA"
	stateRhodeIslandID        StateID = "RI"
	stateSouthCarolinaID      StateID = "SC"
	stateSouthDakotaID        StateID = "SD"
	stateTennesseeID          StateID = "TN"
	stateTexasID              StateID = "TX"
	stateUtahID               StateID = "UT"
	stateVirginiaID           StateID = "VA"
	stateVermontID            StateID = "VT"
	stateWashingtonID         StateID = "WA"
	stateWestVirginiaID       StateID = "WV"
	stateWisconsinID          StateID = "WI"
	stateWyomingID            StateID = "WY"
	stateNonUSID              StateID = "ZZ"
)

// EnumStateItem describes an entry in an enumeration of State
type EnumStateItem struct {
	ID        StateID           `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	FullName    string
	DisplayName string
}

var (
	stateAlaska             = EnumStateItem{stateAlaskaID, "AK", map[string]string{"FullName": "Alaska", "DisplayName": "Alaska"}, "Alaska", 1, "Alaska", "Alaska"}
	stateAlabama            = EnumStateItem{stateAlabamaID, "AL", map[string]string{"FullName": "Alabama", "DisplayName": "Alabama"}, "Alabama", 2, "Alabama", "Alabama"}
	stateArkansas           = EnumStateItem{stateArkansasID, "AR", map[string]string{"FullName": "Arkansas", "DisplayName": "Arkansas"}, "Arkansas", 3, "Arkansas", "Arkansas"}
	stateArizona            = EnumStateItem{stateArizonaID, "AZ", map[string]string{"FullName": "Arizona", "DisplayName": "Arizona"}, "Arizona", 4, "Arizona", "Arizona"}
	stateCalifornia         = EnumStateItem{stateCaliforniaID, "CA", map[string]string{"FullName": "California", "DisplayName": "California"}, "California", 5, "California", "California"}
	stateColorado           = EnumStateItem{stateColoradoID, "CO", map[string]string{"FullName": "Colorado", "DisplayName": "Colorado"}, "Colorado", 6, "Colorado", "Colorado"}
	stateConnecticut        = EnumStateItem{stateConnecticutID, "CT", map[string]string{"FullName": "Connecticut", "DisplayName": "Connecticut"}, "Connecticut", 7, "Connecticut", "Connecticut"}
	stateDistrictOfColumbia = EnumStateItem{stateDistrictOfColumbiaID, "DC", map[string]string{"FullName": "DistrictOfColumbia", "DisplayName": "District Of Columbia"}, "DistrictOfColumbia", 8, "DistrictOfColumbia", "District Of Columbia"}
	stateDelaware           = EnumStateItem{stateDelawareID, "DE", map[string]string{"FullName": "Delaware", "DisplayName": "Delaware"}, "Delaware", 9, "Delaware", "Delaware"}
	stateFlorida            = EnumStateItem{stateFloridaID, "FL", map[string]string{"FullName": "Florida", "DisplayName": "Florida"}, "Florida", 10, "Florida", "Florida"}
	stateGeorgia            = EnumStateItem{stateGeorgiaID, "GA", map[string]string{"FullName": "Georgia", "DisplayName": "Georgia"}, "Georgia", 11, "Georgia", "Georgia"}
	stateHawaii             = EnumStateItem{stateHawaiiID, "HI", map[string]string{"FullName": "Hawaii", "DisplayName": "Hawaii"}, "Hawaii", 12, "Hawaii", "Hawaii"}
	stateIowa               = EnumStateItem{stateIowaID, "IA", map[string]string{"FullName": "Iowa", "DisplayName": "Iowa"}, "Iowa", 13, "Iowa", "Iowa"}
	stateIdaho              = EnumStateItem{stateIdahoID, "ID", map[string]string{"FullName": "Idaho", "DisplayName": "Idaho"}, "Idaho", 14, "Idaho", "Idaho"}
	stateIllinois           = EnumStateItem{stateIllinoisID, "IL", map[string]string{"FullName": "Illinois", "DisplayName": "Illinois"}, "Illinois", 15, "Illinois", "Illinois"}
	stateIndiana            = EnumStateItem{stateIndianaID, "IN", map[string]string{"FullName": "Indiana", "DisplayName": "Indiana"}, "Indiana", 16, "Indiana", "Indiana"}
	stateKansas             = EnumStateItem{stateKansasID, "KS", map[string]string{"FullName": "Kansas", "DisplayName": "Kansas"}, "Kansas", 17, "Kansas", "Kansas"}
	stateKentucky           = EnumStateItem{stateKentuckyID, "KY", map[string]string{"FullName": "Kentucky", "DisplayName": "Kentucky"}, "Kentucky", 18, "Kentucky", "Kentucky"}
	stateLouisiana          = EnumStateItem{stateLouisianaID, "LA", map[string]string{"FullName": "Louisiana", "DisplayName": "Louisiana"}, "Louisiana", 19, "Louisiana", "Louisiana"}
	stateMassachusetts      = EnumStateItem{stateMassachusettsID, "MA", map[string]string{"FullName": "Massachusetts", "DisplayName": "Massachusetts"}, "Massachusetts", 20, "Massachusetts", "Massachusetts"}
	stateMaryland           = EnumStateItem{stateMarylandID, "MD", map[string]string{"FullName": "Maryland", "DisplayName": "Maryland"}, "Maryland", 21, "Maryland", "Maryland"}
	stateMaine              = EnumStateItem{stateMaineID, "ME", map[string]string{"FullName": "Maine", "DisplayName": "Maine"}, "Maine", 22, "Maine", "Maine"}
	stateMichigan           = EnumStateItem{stateMichiganID, "MI", map[string]string{"FullName": "Michigan", "DisplayName": "Michigan"}, "Michigan", 23, "Michigan", "Michigan"}
	stateMinnesota          = EnumStateItem{stateMinnesotaID, "MN", map[string]string{"FullName": "Minnesota", "DisplayName": "Minnesota"}, "Minnesota", 24, "Minnesota", "Minnesota"}
	stateMissouri           = EnumStateItem{stateMissouriID, "MO", map[string]string{"FullName": "Missouri", "DisplayName": "Missouri"}, "Missouri", 25, "Missouri", "Missouri"}
	stateMississippi        = EnumStateItem{stateMississippiID, "MS", map[string]string{"FullName": "Mississippi", "DisplayName": "Mississippi"}, "Mississippi", 26, "Mississippi", "Mississippi"}
	stateMontana            = EnumStateItem{stateMontanaID, "MT", map[string]string{"FullName": "Montana", "DisplayName": "Montana"}, "Montana", 27, "Montana", "Montana"}
	stateNorthCarolina      = EnumStateItem{stateNorthCarolinaID, "NC", map[string]string{"FullName": "NorthCarolina", "DisplayName": "North Carolina"}, "NorthCarolina", 28, "NorthCarolina", "North Carolina"}
	stateNorthDakota        = EnumStateItem{stateNorthDakotaID, "ND", map[string]string{"FullName": "NorthDakota", "DisplayName": "North Dakota"}, "NorthDakota", 29, "NorthDakota", "North Dakota"}
	stateNebraska           = EnumStateItem{stateNebraskaID, "NE", map[string]string{"FullName": "Nebraska", "DisplayName": "Nebraska"}, "Nebraska", 30, "Nebraska", "Nebraska"}
	stateNewHampshire       = EnumStateItem{stateNewHampshireID, "NH", map[string]string{"FullName": "NewHampshire", "DisplayName": "New Hampshire"}, "NewHampshire", 31, "NewHampshire", "New Hampshire"}
	stateNewJersey          = EnumStateItem{stateNewJerseyID, "NJ", map[string]string{"FullName": "NewJersey", "DisplayName": "New Jersey"}, "NewJersey", 32, "NewJersey", "New Jersey"}
	stateNewMexico          = EnumStateItem{stateNewMexicoID, "NM", map[string]string{"FullName": "NewMexico", "DisplayName": "New Mexico"}, "NewMexico", 33, "NewMexico", "New Mexico"}
	stateNevada             = EnumStateItem{stateNevadaID, "NV", map[string]string{"FullName": "Nevada", "DisplayName": "Nevada"}, "Nevada", 34, "Nevada", "Nevada"}
	stateNewYork            = EnumStateItem{stateNewYorkID, "NY", map[string]string{"FullName": "NewYork", "DisplayName": "NewYork"}, "NewYork", 35, "NewYork", "NewYork"}
	stateOhio               = EnumStateItem{stateOhioID, "OH", map[string]string{"FullName": "Ohio", "DisplayName": "Ohio"}, "Ohio", 36, "Ohio", "Ohio"}
	stateOklahoma           = EnumStateItem{stateOklahomaID, "OK", map[string]string{"FullName": "Oklahoma", "DisplayName": "Oklahoma"}, "Oklahoma", 37, "Oklahoma", "Oklahoma"}
	stateOregon             = EnumStateItem{stateOregonID, "OR", map[string]string{"FullName": "Oregon", "DisplayName": "Oregon"}, "Oregon", 38, "Oregon", "Oregon"}
	statePennsylvania       = EnumStateItem{statePennsylvaniaID, "PA", map[string]string{"FullName": "Pennsylvania", "DisplayName": "Pennsylvania"}, "Pennsylvania", 39, "Pennsylvania", "Pennsylvania"}
	stateRhodeIsland        = EnumStateItem{stateRhodeIslandID, "RI", map[string]string{"FullName": "RhodeIsland", "DisplayName": "Rhode Island"}, "RhodeIsland", 40, "RhodeIsland", "Rhode Island"}
	stateSouthCarolina      = EnumStateItem{stateSouthCarolinaID, "SC", map[string]string{"FullName": "SouthCarolina", "DisplayName": "South Carolina"}, "SouthCarolina", 41, "SouthCarolina", "South Carolina"}
	stateSouthDakota        = EnumStateItem{stateSouthDakotaID, "SD", map[string]string{"FullName": "SouthDakota", "DisplayName": "South Dakota"}, "SouthDakota", 42, "SouthDakota", "South Dakota"}
	stateTennessee          = EnumStateItem{stateTennesseeID, "TN", map[string]string{"FullName": "Tennessee", "DisplayName": "Tennessee"}, "Tennessee", 43, "Tennessee", "Tennessee"}
	stateTexas              = EnumStateItem{stateTexasID, "TX", map[string]string{"FullName": "Texas", "DisplayName": "Texas"}, "Texas", 44, "Texas", "Texas"}
	stateUtah               = EnumStateItem{stateUtahID, "UT", map[string]string{"FullName": "Utah", "DisplayName": "Utah"}, "Utah", 45, "Utah", "Utah"}
	stateVirginia           = EnumStateItem{stateVirginiaID, "VA", map[string]string{"FullName": "Virginia", "DisplayName": "Virginia"}, "Virginia", 46, "Virginia", "Virginia"}
	stateVermont            = EnumStateItem{stateVermontID, "VT", map[string]string{"FullName": "Vermont", "DisplayName": "Vermont"}, "Vermont", 47, "Vermont", "Vermont"}
	stateWashington         = EnumStateItem{stateWashingtonID, "WA", map[string]string{"FullName": "Washington", "DisplayName": "Washington"}, "Washington", 48, "Washington", "Washington"}
	stateWestVirginia       = EnumStateItem{stateWestVirginiaID, "WV", map[string]string{"FullName": "WestVirginia", "DisplayName": "West Virginia"}, "WestVirginia", 49, "WestVirginia", "West Virginia"}
	stateWisconsin          = EnumStateItem{stateWisconsinID, "WI", map[string]string{"FullName": "Wisconsin", "DisplayName": "Wisconsin"}, "Wisconsin", 50, "Wisconsin", "Wisconsin"}
	stateWyoming            = EnumStateItem{stateWyomingID, "WY", map[string]string{"FullName": "Wyoming", "DisplayName": "Wyoming"}, "Wyoming", 51, "Wyoming", "Wyoming"}
	stateNonUS              = EnumStateItem{stateNonUSID, "ZZ", map[string]string{"FullName": "Non-US", "DisplayName": "Non-US State"}, "NonUS", 52, "Non-US", "Non-US State"}
)

// EnumState is a collection of State items
type EnumState struct {
	Description string
	Items       []*EnumStateItem
	Name        string

	Alaska             *EnumStateItem
	Alabama            *EnumStateItem
	Arkansas           *EnumStateItem
	Arizona            *EnumStateItem
	California         *EnumStateItem
	Colorado           *EnumStateItem
	Connecticut        *EnumStateItem
	DistrictOfColumbia *EnumStateItem
	Delaware           *EnumStateItem
	Florida            *EnumStateItem
	Georgia            *EnumStateItem
	Hawaii             *EnumStateItem
	Iowa               *EnumStateItem
	Idaho              *EnumStateItem
	Illinois           *EnumStateItem
	Indiana            *EnumStateItem
	Kansas             *EnumStateItem
	Kentucky           *EnumStateItem
	Louisiana          *EnumStateItem
	Massachusetts      *EnumStateItem
	Maryland           *EnumStateItem
	Maine              *EnumStateItem
	Michigan           *EnumStateItem
	Minnesota          *EnumStateItem
	Missouri           *EnumStateItem
	Mississippi        *EnumStateItem
	Montana            *EnumStateItem
	NorthCarolina      *EnumStateItem
	NorthDakota        *EnumStateItem
	Nebraska           *EnumStateItem
	NewHampshire       *EnumStateItem
	NewJersey          *EnumStateItem
	NewMexico          *EnumStateItem
	Nevada             *EnumStateItem
	NewYork            *EnumStateItem
	Ohio               *EnumStateItem
	Oklahoma           *EnumStateItem
	Oregon             *EnumStateItem
	Pennsylvania       *EnumStateItem
	RhodeIsland        *EnumStateItem
	SouthCarolina      *EnumStateItem
	SouthDakota        *EnumStateItem
	Tennessee          *EnumStateItem
	Texas              *EnumStateItem
	Utah               *EnumStateItem
	Virginia           *EnumStateItem
	Vermont            *EnumStateItem
	Washington         *EnumStateItem
	WestVirginia       *EnumStateItem
	Wisconsin          *EnumStateItem
	Wyoming            *EnumStateItem
	NonUS              *EnumStateItem

	itemDict map[string]*EnumStateItem
}

// State is a public singleton instance of EnumState
// representing States of the USA
var State = &EnumState{
	Description: "States of the USA",
	Items: []*EnumStateItem{
		&stateAlaska,
		&stateAlabama,
		&stateArkansas,
		&stateArizona,
		&stateCalifornia,
		&stateColorado,
		&stateConnecticut,
		&stateDistrictOfColumbia,
		&stateDelaware,
		&stateFlorida,
		&stateGeorgia,
		&stateHawaii,
		&stateIowa,
		&stateIdaho,
		&stateIllinois,
		&stateIndiana,
		&stateKansas,
		&stateKentucky,
		&stateLouisiana,
		&stateMassachusetts,
		&stateMaryland,
		&stateMaine,
		&stateMichigan,
		&stateMinnesota,
		&stateMissouri,
		&stateMississippi,
		&stateMontana,
		&stateNorthCarolina,
		&stateNorthDakota,
		&stateNebraska,
		&stateNewHampshire,
		&stateNewJersey,
		&stateNewMexico,
		&stateNevada,
		&stateNewYork,
		&stateOhio,
		&stateOklahoma,
		&stateOregon,
		&statePennsylvania,
		&stateRhodeIsland,
		&stateSouthCarolina,
		&stateSouthDakota,
		&stateTennessee,
		&stateTexas,
		&stateUtah,
		&stateVirginia,
		&stateVermont,
		&stateWashington,
		&stateWestVirginia,
		&stateWisconsin,
		&stateWyoming,
		&stateNonUS,
	},
	Name:               "EnumState",
	Alaska:             &stateAlaska,
	Alabama:            &stateAlabama,
	Arkansas:           &stateArkansas,
	Arizona:            &stateArizona,
	California:         &stateCalifornia,
	Colorado:           &stateColorado,
	Connecticut:        &stateConnecticut,
	DistrictOfColumbia: &stateDistrictOfColumbia,
	Delaware:           &stateDelaware,
	Florida:            &stateFlorida,
	Georgia:            &stateGeorgia,
	Hawaii:             &stateHawaii,
	Iowa:               &stateIowa,
	Idaho:              &stateIdaho,
	Illinois:           &stateIllinois,
	Indiana:            &stateIndiana,
	Kansas:             &stateKansas,
	Kentucky:           &stateKentucky,
	Louisiana:          &stateLouisiana,
	Massachusetts:      &stateMassachusetts,
	Maryland:           &stateMaryland,
	Maine:              &stateMaine,
	Michigan:           &stateMichigan,
	Minnesota:          &stateMinnesota,
	Missouri:           &stateMissouri,
	Mississippi:        &stateMississippi,
	Montana:            &stateMontana,
	NorthCarolina:      &stateNorthCarolina,
	NorthDakota:        &stateNorthDakota,
	Nebraska:           &stateNebraska,
	NewHampshire:       &stateNewHampshire,
	NewJersey:          &stateNewJersey,
	NewMexico:          &stateNewMexico,
	Nevada:             &stateNevada,
	NewYork:            &stateNewYork,
	Ohio:               &stateOhio,
	Oklahoma:           &stateOklahoma,
	Oregon:             &stateOregon,
	Pennsylvania:       &statePennsylvania,
	RhodeIsland:        &stateRhodeIsland,
	SouthCarolina:      &stateSouthCarolina,
	SouthDakota:        &stateSouthDakota,
	Tennessee:          &stateTennessee,
	Texas:              &stateTexas,
	Utah:               &stateUtah,
	Virginia:           &stateVirginia,
	Vermont:            &stateVermont,
	Washington:         &stateWashington,
	WestVirginia:       &stateWestVirginia,
	Wisconsin:          &stateWisconsin,
	Wyoming:            &stateWyoming,
	NonUS:              &stateNonUS,

	itemDict: map[string]*EnumStateItem{
		strings.ToLower(string(stateAlaskaID)):             &stateAlaska,
		strings.ToLower(string(stateAlabamaID)):            &stateAlabama,
		strings.ToLower(string(stateArkansasID)):           &stateArkansas,
		strings.ToLower(string(stateArizonaID)):            &stateArizona,
		strings.ToLower(string(stateCaliforniaID)):         &stateCalifornia,
		strings.ToLower(string(stateColoradoID)):           &stateColorado,
		strings.ToLower(string(stateConnecticutID)):        &stateConnecticut,
		strings.ToLower(string(stateDistrictOfColumbiaID)): &stateDistrictOfColumbia,
		strings.ToLower(string(stateDelawareID)):           &stateDelaware,
		strings.ToLower(string(stateFloridaID)):            &stateFlorida,
		strings.ToLower(string(stateGeorgiaID)):            &stateGeorgia,
		strings.ToLower(string(stateHawaiiID)):             &stateHawaii,
		strings.ToLower(string(stateIowaID)):               &stateIowa,
		strings.ToLower(string(stateIdahoID)):              &stateIdaho,
		strings.ToLower(string(stateIllinoisID)):           &stateIllinois,
		strings.ToLower(string(stateIndianaID)):            &stateIndiana,
		strings.ToLower(string(stateKansasID)):             &stateKansas,
		strings.ToLower(string(stateKentuckyID)):           &stateKentucky,
		strings.ToLower(string(stateLouisianaID)):          &stateLouisiana,
		strings.ToLower(string(stateMassachusettsID)):      &stateMassachusetts,
		strings.ToLower(string(stateMarylandID)):           &stateMaryland,
		strings.ToLower(string(stateMaineID)):              &stateMaine,
		strings.ToLower(string(stateMichiganID)):           &stateMichigan,
		strings.ToLower(string(stateMinnesotaID)):          &stateMinnesota,
		strings.ToLower(string(stateMissouriID)):           &stateMissouri,
		strings.ToLower(string(stateMississippiID)):        &stateMississippi,
		strings.ToLower(string(stateMontanaID)):            &stateMontana,
		strings.ToLower(string(stateNorthCarolinaID)):      &stateNorthCarolina,
		strings.ToLower(string(stateNorthDakotaID)):        &stateNorthDakota,
		strings.ToLower(string(stateNebraskaID)):           &stateNebraska,
		strings.ToLower(string(stateNewHampshireID)):       &stateNewHampshire,
		strings.ToLower(string(stateNewJerseyID)):          &stateNewJersey,
		strings.ToLower(string(stateNewMexicoID)):          &stateNewMexico,
		strings.ToLower(string(stateNevadaID)):             &stateNevada,
		strings.ToLower(string(stateNewYorkID)):            &stateNewYork,
		strings.ToLower(string(stateOhioID)):               &stateOhio,
		strings.ToLower(string(stateOklahomaID)):           &stateOklahoma,
		strings.ToLower(string(stateOregonID)):             &stateOregon,
		strings.ToLower(string(statePennsylvaniaID)):       &statePennsylvania,
		strings.ToLower(string(stateRhodeIslandID)):        &stateRhodeIsland,
		strings.ToLower(string(stateSouthCarolinaID)):      &stateSouthCarolina,
		strings.ToLower(string(stateSouthDakotaID)):        &stateSouthDakota,
		strings.ToLower(string(stateTennesseeID)):          &stateTennessee,
		strings.ToLower(string(stateTexasID)):              &stateTexas,
		strings.ToLower(string(stateUtahID)):               &stateUtah,
		strings.ToLower(string(stateVirginiaID)):           &stateVirginia,
		strings.ToLower(string(stateVermontID)):            &stateVermont,
		strings.ToLower(string(stateWashingtonID)):         &stateWashington,
		strings.ToLower(string(stateWestVirginiaID)):       &stateWestVirginia,
		strings.ToLower(string(stateWisconsinID)):          &stateWisconsin,
		strings.ToLower(string(stateWyomingID)):            &stateWyoming,
		strings.ToLower(string(stateNonUSID)):              &stateNonUS,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumState) ByID(id StateIdentifier) *EnumStateItem {
	if e == nil || id == nil {
		return nil
	}

	if idx := id.ID(); idx != nil {
		idxString := strings.ToLower(string(*idx))

		if rtn, ok := e.itemDict[idxString]; ok {
			return rtn
		}
	}

	return nil
}

// ByIDString retrieves an entry by a string representation of its ID
func (e *EnumState) ByIDString(idx string) *EnumStateItem {
	if e == nil || len(e.itemDict) == 0 || idx == "" {
		return nil
	}

	normIdx := strings.ToLower(idx)
	if rtn, ok := e.itemDict[normIdx]; ok {
		return rtn
	}

	return nil
}

// ByIndex retrieves an entry based on its index (NOT NECESSARILY sort order) value
func (e *EnumState) ByIndex(idx int) *EnumStateItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedStateID is a struct that is designed to replace a *StateID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *StateID it contains while being a better JSON citizen.
type ValidatedStateID struct {
	// id will point to a valid StateID, if possible
	// If id is nil, then ValidatedStateID.Valid() will return false.
	id *StateID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedStateID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedStateID
func (vi *ValidatedStateID) Clone() *ValidatedStateID {
	if vi == nil {
		return nil
	}

	var cid *StateID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedStateID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedStateIds represent the same State
func (vi *ValidatedStateID) Equals(vj *ValidatedStateID) bool {
	if vi == nil && vj == nil {
		return true
	}

	if vi == nil || vj == nil {
		return false
	}

	if vi.id == nil && vj.id == nil {
		return true
	}

	if vi.id == nil || vj.id == nil {
		return false
	}

	return vi.id.Equals(vj.id)
}

// Valid returns true if and only if the ValidatedStateID corresponds to a recognized State
func (vi *ValidatedStateID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedStateID) ID() *StateID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedStateID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedStateID) ValidatedID() *ValidatedStateID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedStateID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedStateID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedStateID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedStateID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedStateID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := StateID(capString)
	item := State.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		if vi.Errors == nil {
			vi.Errors = []error{}
		}
		vi.Errors = append(vi.Errors, err)
		return nil
	}

	vi.id = item.ID.Clone()

	return nil
}

func (vi ValidatedStateID) String() string {
	return vi.ToIDString()
}

type StateIdentifier interface {
	ID() *StateID
	Valid() bool
}
