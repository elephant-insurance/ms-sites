package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// MilitaryServiceID uniquely identifies a particular MilitaryService
type MilitaryServiceID string

// Clone creates a safe, independent copy of a MilitaryServiceID
func (i *MilitaryServiceID) Clone() *MilitaryServiceID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two MilitaryServiceIds are equivalent
func (i *MilitaryServiceID) Equals(j *MilitaryServiceID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *MilitaryServiceID that is either valid or nil
func (i *MilitaryServiceID) ID() *MilitaryServiceID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *MilitaryServiceID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the MilitaryServiceID corresponds to a recognized MilitaryService
func (i *MilitaryServiceID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return MilitaryService.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *MilitaryServiceID) ValidatedID() *ValidatedMilitaryServiceID {
	if i != nil {
		return &ValidatedMilitaryServiceID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *MilitaryServiceID) MarshalJSON() ([]byte, error) {
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

func (i *MilitaryServiceID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := MilitaryServiceID(dataString)
	item := MilitaryService.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	militaryServiceAirForceE1AirmanBasicID                           MilitaryServiceID = "1"
	militaryServiceAirForceE2AirmanID                                MilitaryServiceID = "2"
	militaryServiceAirForceE3AirmanFirstClassID                      MilitaryServiceID = "3"
	militaryServiceAirForceE4SeniorAirmanID                          MilitaryServiceID = "4"
	militaryServiceAirForceE4SergeantID                              MilitaryServiceID = "5"
	militaryServiceAirForceE5StaffSergeantID                         MilitaryServiceID = "6"
	militaryServiceAirForceE6TechnicalSergeantID                     MilitaryServiceID = "7"
	militaryServiceAirForceE7FirstSergeantID                         MilitaryServiceID = "8"
	militaryServiceAirForceE7MasterSergeantID                        MilitaryServiceID = "9"
	militaryServiceAirForceE8FirstSergeantID                         MilitaryServiceID = "10"
	militaryServiceAirForceE8SeniorMasterSergeantID                  MilitaryServiceID = "11"
	militaryServiceAirForceE9ChiefMasterSergeantID                   MilitaryServiceID = "12"
	militaryServiceAirForceE9ChiefMasterSergeantOfTheAirForceID      MilitaryServiceID = "13"
	militaryServiceAirForceE9FirstSergeantID                         MilitaryServiceID = "14"
	militaryServiceAirForceO1SecondLieutenantID                      MilitaryServiceID = "15"
	militaryServiceAirForceO2FirstLieutenantID                       MilitaryServiceID = "17"
	militaryServiceAirForceO3CaptainID                               MilitaryServiceID = "18"
	militaryServiceAirForceO4MajorID                                 MilitaryServiceID = "19"
	militaryServiceAirForceO5LieutenantColonelID                     MilitaryServiceID = "20"
	militaryServiceAirForceO6ColonelID                               MilitaryServiceID = "21"
	militaryServiceAirForceO7BrigadierGeneralID                      MilitaryServiceID = "22"
	militaryServiceAirForceO8MajorGeneralID                          MilitaryServiceID = "23"
	militaryServiceAirForceO9LieutenantGeneralID                     MilitaryServiceID = "24"
	militaryServiceAirForceO10GeneralID                              MilitaryServiceID = "16"
	militaryServiceArmyE1PrivateID                                   MilitaryServiceID = "25"
	militaryServiceArmyE2PrivateID                                   MilitaryServiceID = "26"
	militaryServiceArmyE3PrivateFirstClassID                         MilitaryServiceID = "27"
	militaryServiceArmyE4CorporalID                                  MilitaryServiceID = "28"
	militaryServiceArmyE4SpecialistID                                MilitaryServiceID = "29"
	militaryServiceArmyE4PCorporalID                                 MilitaryServiceID = "30"
	militaryServiceArmyE4PSpecialistID                               MilitaryServiceID = "31"
	militaryServiceArmyE5SergeantID                                  MilitaryServiceID = "32"
	militaryServiceArmyE5PSergeantID                                 MilitaryServiceID = "33"
	militaryServiceArmyE6StaffSergeantID                             MilitaryServiceID = "34"
	militaryServiceArmyE6PStaffSergeantID                            MilitaryServiceID = "35"
	militaryServiceArmyE7SergeantFirstClassID                        MilitaryServiceID = "36"
	militaryServiceArmyE8FirstSergeantID                             MilitaryServiceID = "37"
	militaryServiceArmyE8MasterSergeantID                            MilitaryServiceID = "38"
	militaryServiceArmyE9CommandSergeantMajorID                      MilitaryServiceID = "39"
	militaryServiceArmyE9SergeantMajorID                             MilitaryServiceID = "40"
	militaryServiceArmyE9SergeantMajorOfTheArmyID                    MilitaryServiceID = "41"
	militaryServiceArmyO1SecondLieutenantID                          MilitaryServiceID = "42"
	militaryServiceArmyO2FirstLieutenantID                           MilitaryServiceID = "44"
	militaryServiceArmyO3CaptainID                                   MilitaryServiceID = "45"
	militaryServiceArmyO4MajorID                                     MilitaryServiceID = "46"
	militaryServiceArmyO5LieutenantColonelID                         MilitaryServiceID = "47"
	militaryServiceArmyO6ColonelID                                   MilitaryServiceID = "48"
	militaryServiceArmyO7BrigadierGeneralID                          MilitaryServiceID = "49"
	militaryServiceArmyO8MajorGeneralID                              MilitaryServiceID = "50"
	militaryServiceArmyO9LieutenantGeneralID                         MilitaryServiceID = "51"
	militaryServiceArmyO10GeneralID                                  MilitaryServiceID = "43"
	militaryServiceArmyW1WarrantOfficerID                            MilitaryServiceID = "52"
	militaryServiceArmyW2ChiefWarrantOfficerID                       MilitaryServiceID = "53"
	militaryServiceArmyW3ChiefWarrantOfficerID                       MilitaryServiceID = "54"
	militaryServiceArmyW4ChiefWarrantOfficerID                       MilitaryServiceID = "55"
	militaryServiceArmyW5ChiefWarrantOfficerID                       MilitaryServiceID = "56"
	militaryServiceCoastGuardE1SeamanRecruitID                       MilitaryServiceID = "57"
	militaryServiceCoastGuardE2SeamanApprenticeID                    MilitaryServiceID = "58"
	militaryServiceCoastGuardE3SeamanID                              MilitaryServiceID = "59"
	militaryServiceCoastGuardE4PettyOfficerThirdClassID              MilitaryServiceID = "60"
	militaryServiceCoastGuardE5PettyOfficerSecondClassID             MilitaryServiceID = "61"
	militaryServiceCoastGuardE6PettyOfficerFirstClassID              MilitaryServiceID = "62"
	militaryServiceCoastGuardE7ChiefPettyOfficerID                   MilitaryServiceID = "63"
	militaryServiceCoastGuardE8SeniorChiefPettyOfficerID             MilitaryServiceID = "64"
	militaryServiceCoastGuardE9FleetCommandMasterChiefPettyOfficerID MilitaryServiceID = "65"
	militaryServiceCoastGuardE9MasterChiefPettyOfficerID             MilitaryServiceID = "66"
	militaryServiceCoastGuardO1EnsignID                              MilitaryServiceID = "67"
	militaryServiceCoastGuardO2LieutenantJuniorGradeID               MilitaryServiceID = "70"
	militaryServiceCoastGuardO3LieutenantID                          MilitaryServiceID = "71"
	militaryServiceCoastGuardO4LieutenantCommanderID                 MilitaryServiceID = "72"
	militaryServiceCoastGuardO5CommanderID                           MilitaryServiceID = "73"
	militaryServiceCoastGuardO6CaptainID                             MilitaryServiceID = "74"
	militaryServiceCoastGuardO7RearAdmiralID                         MilitaryServiceID = "75"
	militaryServiceCoastGuardO8RearAdmiralID                         MilitaryServiceID = "76"
	militaryServiceCoastGuardO9ViceAdmiralID                         MilitaryServiceID = "77"
	militaryServiceCoastGuardO10AdmiralID                            MilitaryServiceID = "68"
	militaryServiceCoastGuardO10FleetAdmiralID                       MilitaryServiceID = "69"
	militaryServiceMarinesE1PrivateID                                MilitaryServiceID = "78"
	militaryServiceMarinesE2PrivateFirstClassID                      MilitaryServiceID = "79"
	militaryServiceMarinesE3LanceCorporalID                          MilitaryServiceID = "80"
	militaryServiceMarinesE4CorporalID                               MilitaryServiceID = "81"
	militaryServiceMarinesE5SergeantID                               MilitaryServiceID = "82"
	militaryServiceMarinesE6StaffSergeantID                          MilitaryServiceID = "83"
	militaryServiceMarinesE7GunnerySergeantID                        MilitaryServiceID = "84"
	militaryServiceMarinesE8FirstSergeantID                          MilitaryServiceID = "85"
	militaryServiceMarinesE8MasterSergeantID                         MilitaryServiceID = "86"
	militaryServiceMarinesE9MasterGunnerySergeantID                  MilitaryServiceID = "87"
	militaryServiceMarinesE9SergeantMajorID                          MilitaryServiceID = "88"
	militaryServiceMarinesE9SergeantMajorOfTheMarineCorpsID          MilitaryServiceID = "89"
	militaryServiceMarinesO1SecondLieutenantID                       MilitaryServiceID = "90"
	militaryServiceMarinesO2FirstLieutenantID                        MilitaryServiceID = "92"
	militaryServiceMarinesO3CaptainID                                MilitaryServiceID = "93"
	militaryServiceMarinesO4MajorID                                  MilitaryServiceID = "94"
	militaryServiceMarinesO5LieutenantColonelID                      MilitaryServiceID = "95"
	militaryServiceMarinesO6ColonelID                                MilitaryServiceID = "96"
	militaryServiceMarinesO7BrigadierGeneralID                       MilitaryServiceID = "97"
	militaryServiceMarinesO8MajorGeneralID                           MilitaryServiceID = "98"
	militaryServiceMarinesO9LieutenantGeneralID                      MilitaryServiceID = "99"
	militaryServiceMarinesO10GeneralID                               MilitaryServiceID = "91"
	militaryServiceMarinesW1WarrantOfficerID                         MilitaryServiceID = "100"
	militaryServiceMarinesW2ChiefWarrantOfficerID                    MilitaryServiceID = "101"
	militaryServiceMarinesW3ChiefWarrantOfficerID                    MilitaryServiceID = "102"
	militaryServiceMarinesW4ChiefWarrantOfficerID                    MilitaryServiceID = "103"
	militaryServiceMarinesW5ChiefWarrantOfficerID                    MilitaryServiceID = "104"
	militaryServiceNavyE1SeamanRecruitID                             MilitaryServiceID = "105"
	militaryServiceNavyE2SeamanApprenticeID                          MilitaryServiceID = "106"
	militaryServiceNavyE3SeamanID                                    MilitaryServiceID = "107"
	militaryServiceNavyE4PettyOfficerThirdClassID                    MilitaryServiceID = "108"
	militaryServiceNavyE5PettyOfficerSecondClassID                   MilitaryServiceID = "109"
	militaryServiceNavyE6PettyOfficerFirstClassID                    MilitaryServiceID = "110"
	militaryServiceNavyE7ChiefPettyOfficerID                         MilitaryServiceID = "111"
	militaryServiceNavyE8SeniorChiefPettyOfficerID                   MilitaryServiceID = "112"
	militaryServiceNavyE9MasterChiefPettyOfficerID                   MilitaryServiceID = "113"
	militaryServiceNavyE9MasterChiefPettyOfficerOfTheNavyID          MilitaryServiceID = "114"
	militaryServiceNavyO1EnsignID                                    MilitaryServiceID = "115"
	militaryServiceNavyO2LieutenantJuniorGradeID                     MilitaryServiceID = "117"
	militaryServiceNavyO3LieutenantID                                MilitaryServiceID = "118"
	militaryServiceNavyO4LieutenantCommanderID                       MilitaryServiceID = "119"
	militaryServiceNavyO5CommanderID                                 MilitaryServiceID = "120"
	militaryServiceNavyO6CaptainID                                   MilitaryServiceID = "121"
	militaryServiceNavyO7RearAdmiralID                               MilitaryServiceID = "122"
	militaryServiceNavyO8RearAdmiralID                               MilitaryServiceID = "123"
	militaryServiceNavyO9ViceAdmiralID                               MilitaryServiceID = "124"
	militaryServiceNavyO10AdmiralID                                  MilitaryServiceID = "116"
	militaryServiceNavyW1WarrantOfficerID                            MilitaryServiceID = "125"
	militaryServiceNavyW2ChiefWarrantOfficerID                       MilitaryServiceID = "126"
	militaryServiceNavyW3ChiefWarrantOfficerID                       MilitaryServiceID = "127"
	militaryServiceNavyW4ChiefWarrantOfficerID                       MilitaryServiceID = "128"
)

// EnumMilitaryServiceItem describes an entry in an enumeration of MilitaryService
type EnumMilitaryServiceItem struct {
	ID        MilitaryServiceID `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int

	// Meta Properties
	Branch *EnumMilitaryBranchItem
	Rank   *EnumMilitaryRankItem
}

var (
	militaryServiceAirForceE1AirmanBasic                           = EnumMilitaryServiceItem{militaryServiceAirForceE1AirmanBasicID, "Air Force E1 Airman Basic", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE1ID)}, "AirForceE1AirmanBasic", 1, &militaryBranchAirForce, &militaryRankEnlistedE1}
	militaryServiceAirForceE2Airman                                = EnumMilitaryServiceItem{militaryServiceAirForceE2AirmanID, "Air Force E2 Airman", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE2ID)}, "AirForceE2Airman", 2, &militaryBranchAirForce, &militaryRankEnlistedE2}
	militaryServiceAirForceE3AirmanFirstClass                      = EnumMilitaryServiceItem{militaryServiceAirForceE3AirmanFirstClassID, "Air Force E3 Airman First Class", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE3ID)}, "AirForceE3AirmanFirstClass", 3, &militaryBranchAirForce, &militaryRankEnlistedE3}
	militaryServiceAirForceE4SeniorAirman                          = EnumMilitaryServiceItem{militaryServiceAirForceE4SeniorAirmanID, "Air Force E4 Senior Airman", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE4ID)}, "AirForceE4SeniorAirman", 4, &militaryBranchAirForce, &militaryRankEnlistedE4}
	militaryServiceAirForceE4Sergeant                              = EnumMilitaryServiceItem{militaryServiceAirForceE4SergeantID, "Air Force E4 Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE4ID)}, "AirForceE4Sergeant", 5, &militaryBranchAirForce, &militaryRankEnlistedE4}
	militaryServiceAirForceE5StaffSergeant                         = EnumMilitaryServiceItem{militaryServiceAirForceE5StaffSergeantID, "Air Force E5 Staff Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE5ID)}, "AirForceE5StaffSergeant", 6, &militaryBranchAirForce, &militaryRankNonCommissionedE5}
	militaryServiceAirForceE6TechnicalSergeant                     = EnumMilitaryServiceItem{militaryServiceAirForceE6TechnicalSergeantID, "Air Force E6 Technical Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE6ID)}, "AirForceE6TechnicalSergeant", 7, &militaryBranchAirForce, &militaryRankNonCommissionedE6}
	militaryServiceAirForceE7FirstSergeant                         = EnumMilitaryServiceItem{militaryServiceAirForceE7FirstSergeantID, "Air Force E7 First Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE7ID)}, "AirForceE7FirstSergeant", 8, &militaryBranchAirForce, &militaryRankNonCommissionedE7}
	militaryServiceAirForceE7MasterSergeant                        = EnumMilitaryServiceItem{militaryServiceAirForceE7MasterSergeantID, "Air Force E7 Master Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE7ID)}, "AirForceE7MasterSergeant", 9, &militaryBranchAirForce, &militaryRankNonCommissionedE7}
	militaryServiceAirForceE8FirstSergeant                         = EnumMilitaryServiceItem{militaryServiceAirForceE8FirstSergeantID, "Air Force E8 First Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE8ID)}, "AirForceE8FirstSergeant", 10, &militaryBranchAirForce, &militaryRankNonCommissionedE8}
	militaryServiceAirForceE8SeniorMasterSergeant                  = EnumMilitaryServiceItem{militaryServiceAirForceE8SeniorMasterSergeantID, "Air Force E8 Senior Master Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE8ID)}, "AirForceE8SeniorMasterSergeant", 11, &militaryBranchAirForce, &militaryRankNonCommissionedE8}
	militaryServiceAirForceE9ChiefMasterSergeant                   = EnumMilitaryServiceItem{militaryServiceAirForceE9ChiefMasterSergeantID, "Air Force E9 Chief Master Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "AirForceE9ChiefMasterSergeant", 12, &militaryBranchAirForce, &militaryRankNonCommissionedE9}
	militaryServiceAirForceE9ChiefMasterSergeantOfTheAirForce      = EnumMilitaryServiceItem{militaryServiceAirForceE9ChiefMasterSergeantOfTheAirForceID, "Air Force E9 Chief Master Sergeant Of The AirForce", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "AirForceE9ChiefMasterSergeantOfTheAirForce", 13, &militaryBranchAirForce, &militaryRankNonCommissionedE9}
	militaryServiceAirForceE9FirstSergeant                         = EnumMilitaryServiceItem{militaryServiceAirForceE9FirstSergeantID, "Air Force E9 First Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "AirForceE9FirstSergeant", 14, &militaryBranchAirForce, &militaryRankNonCommissionedE9}
	militaryServiceAirForceO1SecondLieutenant                      = EnumMilitaryServiceItem{militaryServiceAirForceO1SecondLieutenantID, "Air Force O1 Second Lieutenant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO1ID)}, "AirForceO1SecondLieutenant", 15, &militaryBranchAirForce, &militaryRankCommissionedOfficerO1}
	militaryServiceAirForceO2FirstLieutenant                       = EnumMilitaryServiceItem{militaryServiceAirForceO2FirstLieutenantID, "Air Force O2 First Lieutenant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO2ID)}, "AirForceO2FirstLieutenant", 16, &militaryBranchAirForce, &militaryRankCommissionedOfficerO2}
	militaryServiceAirForceO3Captain                               = EnumMilitaryServiceItem{militaryServiceAirForceO3CaptainID, "Air Force O3 Captain", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO3ID)}, "AirForceO3Captain", 17, &militaryBranchAirForce, &militaryRankCommissionedOfficerO3}
	militaryServiceAirForceO4Major                                 = EnumMilitaryServiceItem{militaryServiceAirForceO4MajorID, "Air Force O4 Major", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO4ID)}, "AirForceO4Major", 18, &militaryBranchAirForce, &militaryRankCommissionedOfficerO4}
	militaryServiceAirForceO5LieutenantColonel                     = EnumMilitaryServiceItem{militaryServiceAirForceO5LieutenantColonelID, "Air Force O5 Lieutenant Colonel", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO5ID)}, "AirForceO5LieutenantColonel", 19, &militaryBranchAirForce, &militaryRankCommissionedOfficerO5}
	militaryServiceAirForceO6Colonel                               = EnumMilitaryServiceItem{militaryServiceAirForceO6ColonelID, "Air Force O6 Colonel", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO6ID)}, "AirForceO6Colonel", 20, &militaryBranchAirForce, &militaryRankCommissionedOfficerO6}
	militaryServiceAirForceO7BrigadierGeneral                      = EnumMilitaryServiceItem{militaryServiceAirForceO7BrigadierGeneralID, "Air Force O7 Brigadier General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO7ID)}, "AirForceO7BrigadierGeneral", 21, &militaryBranchAirForce, &militaryRankCommissionedOfficerO7}
	militaryServiceAirForceO8MajorGeneral                          = EnumMilitaryServiceItem{militaryServiceAirForceO8MajorGeneralID, "Air Force O8 Major General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO8ID)}, "AirForceO8MajorGeneral", 22, &militaryBranchAirForce, &militaryRankCommissionedOfficerO8}
	militaryServiceAirForceO9LieutenantGeneral                     = EnumMilitaryServiceItem{militaryServiceAirForceO9LieutenantGeneralID, "Air Force O9 Lieutenant General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO9ID)}, "AirForceO9LieutenantGeneral", 23, &militaryBranchAirForce, &militaryRankCommissionedOfficerO9}
	militaryServiceAirForceO10General                              = EnumMilitaryServiceItem{militaryServiceAirForceO10GeneralID, "Air Force O10 General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchAirForceID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO10ID)}, "AirForceO10General", 24, &militaryBranchAirForce, &militaryRankCommissionedOfficerO10}
	militaryServiceArmyE1Private                                   = EnumMilitaryServiceItem{militaryServiceArmyE1PrivateID, "Army E1 Private", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE1ID)}, "ArmyE1Private", 25, &militaryBranchArmy, &militaryRankEnlistedE1}
	militaryServiceArmyE2Private                                   = EnumMilitaryServiceItem{militaryServiceArmyE2PrivateID, "Army E2 Private", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE2ID)}, "ArmyE2Private", 26, &militaryBranchArmy, &militaryRankEnlistedE2}
	militaryServiceArmyE3PrivateFirstClass                         = EnumMilitaryServiceItem{militaryServiceArmyE3PrivateFirstClassID, "Army E3 Private First Class", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE3ID)}, "ArmyE3PrivateFirstClass", 27, &militaryBranchArmy, &militaryRankEnlistedE3}
	militaryServiceArmyE4Corporal                                  = EnumMilitaryServiceItem{militaryServiceArmyE4CorporalID, "Army E4 Corporal", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE4ID)}, "ArmyE4Corporal", 28, &militaryBranchArmy, &militaryRankEnlistedE4}
	militaryServiceArmyE4Specialist                                = EnumMilitaryServiceItem{militaryServiceArmyE4SpecialistID, "Army E4 Specialist", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE4ID)}, "ArmyE4Specialist", 29, &militaryBranchArmy, &militaryRankEnlistedE4}
	militaryServiceArmyE4PCorporal                                 = EnumMilitaryServiceItem{militaryServiceArmyE4PCorporalID, "Army E4P Corporal", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE4ID)}, "ArmyE4PCorporal", 30, &militaryBranchArmy, &militaryRankEnlistedE4}
	militaryServiceArmyE4PSpecialist                               = EnumMilitaryServiceItem{militaryServiceArmyE4PSpecialistID, "Army E4P Specialist", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE4ID)}, "ArmyE4PSpecialist", 31, &militaryBranchArmy, &militaryRankEnlistedE4}
	militaryServiceArmyE5Sergeant                                  = EnumMilitaryServiceItem{militaryServiceArmyE5SergeantID, "Army E5 Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE5ID)}, "ArmyE5Sergeant", 32, &militaryBranchArmy, &militaryRankNonCommissionedE5}
	militaryServiceArmyE5PSergeant                                 = EnumMilitaryServiceItem{militaryServiceArmyE5PSergeantID, "Army E5P Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE5ID)}, "ArmyE5PSergeant", 33, &militaryBranchArmy, &militaryRankNonCommissionedE5}
	militaryServiceArmyE6StaffSergeant                             = EnumMilitaryServiceItem{militaryServiceArmyE6StaffSergeantID, "Army E6 Staff Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE6ID)}, "ArmyE6StaffSergeant", 34, &militaryBranchArmy, &militaryRankNonCommissionedE6}
	militaryServiceArmyE6PStaffSergeant                            = EnumMilitaryServiceItem{militaryServiceArmyE6PStaffSergeantID, "Army E6P Staff Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE6ID)}, "ArmyE6PStaffSergeant", 35, &militaryBranchArmy, &militaryRankNonCommissionedE6}
	militaryServiceArmyE7SergeantFirstClass                        = EnumMilitaryServiceItem{militaryServiceArmyE7SergeantFirstClassID, "Army E7 Sergeant First Class", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE7ID)}, "ArmyE7SergeantFirstClass", 36, &militaryBranchArmy, &militaryRankNonCommissionedE7}
	militaryServiceArmyE8FirstSergeant                             = EnumMilitaryServiceItem{militaryServiceArmyE8FirstSergeantID, "Army E8 First Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE8ID)}, "ArmyE8FirstSergeant", 37, &militaryBranchArmy, &militaryRankNonCommissionedE8}
	militaryServiceArmyE8MasterSergeant                            = EnumMilitaryServiceItem{militaryServiceArmyE8MasterSergeantID, "Army E8 Master Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE8ID)}, "ArmyE8MasterSergeant", 38, &militaryBranchArmy, &militaryRankNonCommissionedE8}
	militaryServiceArmyE9CommandSergeantMajor                      = EnumMilitaryServiceItem{militaryServiceArmyE9CommandSergeantMajorID, "Army E9 Command Sergeant Major", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "ArmyE9CommandSergeantMajor", 39, &militaryBranchArmy, &militaryRankNonCommissionedE9}
	militaryServiceArmyE9SergeantMajor                             = EnumMilitaryServiceItem{militaryServiceArmyE9SergeantMajorID, "Army E9 Sergeant Major", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "ArmyE9SergeantMajor", 40, &militaryBranchArmy, &militaryRankNonCommissionedE9}
	militaryServiceArmyE9SergeantMajorOfTheArmy                    = EnumMilitaryServiceItem{militaryServiceArmyE9SergeantMajorOfTheArmyID, "Army E9 Sergeant Major Of The Army", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "ArmyE9SergeantMajorOfTheArmy", 41, &militaryBranchArmy, &militaryRankNonCommissionedE9}
	militaryServiceArmyO1SecondLieutenant                          = EnumMilitaryServiceItem{militaryServiceArmyO1SecondLieutenantID, "Army O1 Second Lieutenant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO1ID)}, "ArmyO1SecondLieutenant", 42, &militaryBranchArmy, &militaryRankCommissionedOfficerO1}
	militaryServiceArmyO2FirstLieutenant                           = EnumMilitaryServiceItem{militaryServiceArmyO2FirstLieutenantID, "Army O2 First Lieutenant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO2ID)}, "ArmyO2FirstLieutenant", 43, &militaryBranchArmy, &militaryRankCommissionedOfficerO2}
	militaryServiceArmyO3Captain                                   = EnumMilitaryServiceItem{militaryServiceArmyO3CaptainID, "Army O3 Captain", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO3ID)}, "ArmyO3Captain", 44, &militaryBranchArmy, &militaryRankCommissionedOfficerO3}
	militaryServiceArmyO4Major                                     = EnumMilitaryServiceItem{militaryServiceArmyO4MajorID, "Army O4 Major", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO4ID)}, "ArmyO4Major", 45, &militaryBranchArmy, &militaryRankCommissionedOfficerO4}
	militaryServiceArmyO5LieutenantColonel                         = EnumMilitaryServiceItem{militaryServiceArmyO5LieutenantColonelID, "Army O5 Lieutenant Colonel", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO5ID)}, "ArmyO5LieutenantColonel", 46, &militaryBranchArmy, &militaryRankCommissionedOfficerO5}
	militaryServiceArmyO6Colonel                                   = EnumMilitaryServiceItem{militaryServiceArmyO6ColonelID, "Army O6 Colonel", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO6ID)}, "ArmyO6Colonel", 47, &militaryBranchArmy, &militaryRankCommissionedOfficerO6}
	militaryServiceArmyO7BrigadierGeneral                          = EnumMilitaryServiceItem{militaryServiceArmyO7BrigadierGeneralID, "Army O7 Brigadier General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO7ID)}, "ArmyO7BrigadierGeneral", 48, &militaryBranchArmy, &militaryRankCommissionedOfficerO7}
	militaryServiceArmyO8MajorGeneral                              = EnumMilitaryServiceItem{militaryServiceArmyO8MajorGeneralID, "Army O8 Major General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO8ID)}, "ArmyO8MajorGeneral", 49, &militaryBranchArmy, &militaryRankCommissionedOfficerO8}
	militaryServiceArmyO9LieutenantGeneral                         = EnumMilitaryServiceItem{militaryServiceArmyO9LieutenantGeneralID, "Army O9 Lieutenant General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO9ID)}, "ArmyO9LieutenantGeneral", 50, &militaryBranchArmy, &militaryRankCommissionedOfficerO9}
	militaryServiceArmyO10General                                  = EnumMilitaryServiceItem{militaryServiceArmyO10GeneralID, "Army O10 General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO10ID)}, "ArmyO10General", 51, &militaryBranchArmy, &militaryRankCommissionedOfficerO10}
	militaryServiceArmyW1WarrantOfficer                            = EnumMilitaryServiceItem{militaryServiceArmyW1WarrantOfficerID, "Army W1 Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW1ID)}, "ArmyW1WarrantOfficer", 52, &militaryBranchArmy, &militaryRankWarrantOfficerOW1}
	militaryServiceArmyW2ChiefWarrantOfficer                       = EnumMilitaryServiceItem{militaryServiceArmyW2ChiefWarrantOfficerID, "Army W2 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW2ID)}, "ArmyW2ChiefWarrantOfficer", 53, &militaryBranchArmy, &militaryRankWarrantOfficerOW2}
	militaryServiceArmyW3ChiefWarrantOfficer                       = EnumMilitaryServiceItem{militaryServiceArmyW3ChiefWarrantOfficerID, "Army W3 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW3ID)}, "ArmyW3ChiefWarrantOfficer", 54, &militaryBranchArmy, &militaryRankWarrantOfficerOW3}
	militaryServiceArmyW4ChiefWarrantOfficer                       = EnumMilitaryServiceItem{militaryServiceArmyW4ChiefWarrantOfficerID, "Army W4 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW4ID)}, "ArmyW4ChiefWarrantOfficer", 55, &militaryBranchArmy, &militaryRankWarrantOfficerOW4}
	militaryServiceArmyW5ChiefWarrantOfficer                       = EnumMilitaryServiceItem{militaryServiceArmyW5ChiefWarrantOfficerID, "Army W5 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchArmyID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW5ID)}, "ArmyW5ChiefWarrantOfficer", 56, &militaryBranchArmy, &militaryRankWarrantOfficerOW5}
	militaryServiceCoastGuardE1SeamanRecruit                       = EnumMilitaryServiceItem{militaryServiceCoastGuardE1SeamanRecruitID, "Coast Guard E1 Seaman Recruit", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE1ID)}, "CoastGuardE1SeamanRecruit", 57, &militaryBranchCoastGuard, &militaryRankEnlistedE1}
	militaryServiceCoastGuardE2SeamanApprentice                    = EnumMilitaryServiceItem{militaryServiceCoastGuardE2SeamanApprenticeID, "Coast Guard E2 Seaman Apprentice", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE2ID)}, "CoastGuardE2SeamanApprentice", 58, &militaryBranchCoastGuard, &militaryRankEnlistedE2}
	militaryServiceCoastGuardE3Seaman                              = EnumMilitaryServiceItem{militaryServiceCoastGuardE3SeamanID, "Coast Guard E3 Seaman", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE3ID)}, "CoastGuardE3Seaman", 59, &militaryBranchCoastGuard, &militaryRankEnlistedE3}
	militaryServiceCoastGuardE4PettyOfficerThirdClass              = EnumMilitaryServiceItem{militaryServiceCoastGuardE4PettyOfficerThirdClassID, "Coast GuardE4 Petty Officer Third Class", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE4ID)}, "CoastGuardE4PettyOfficerThirdClass", 60, &militaryBranchCoastGuard, &militaryRankEnlistedE4}
	militaryServiceCoastGuardE5PettyOfficerSecondClass             = EnumMilitaryServiceItem{militaryServiceCoastGuardE5PettyOfficerSecondClassID, "Coast Guard E5 Petty Officer Second Class", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE5ID)}, "CoastGuardE5PettyOfficerSecondClass", 61, &militaryBranchCoastGuard, &militaryRankNonCommissionedE5}
	militaryServiceCoastGuardE6PettyOfficerFirstClass              = EnumMilitaryServiceItem{militaryServiceCoastGuardE6PettyOfficerFirstClassID, "Coast Guard E6 Petty Officer First Class", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE6ID)}, "CoastGuardE6PettyOfficerFirstClass", 62, &militaryBranchCoastGuard, &militaryRankNonCommissionedE6}
	militaryServiceCoastGuardE7ChiefPettyOfficer                   = EnumMilitaryServiceItem{militaryServiceCoastGuardE7ChiefPettyOfficerID, "Coast Guard E7 Chief Petty Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE7ID)}, "CoastGuardE7ChiefPettyOfficer", 63, &militaryBranchCoastGuard, &militaryRankNonCommissionedE7}
	militaryServiceCoastGuardE8SeniorChiefPettyOfficer             = EnumMilitaryServiceItem{militaryServiceCoastGuardE8SeniorChiefPettyOfficerID, "Coast Guard E8 Senior Chief Petty Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE8ID)}, "CoastGuardE8SeniorChiefPettyOfficer", 64, &militaryBranchCoastGuard, &militaryRankNonCommissionedE8}
	militaryServiceCoastGuardE9FleetCommandMasterChiefPettyOfficer = EnumMilitaryServiceItem{militaryServiceCoastGuardE9FleetCommandMasterChiefPettyOfficerID, "Coast Guard E9 Fleet/Command Master Chief Petty Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "CoastGuardE9FleetCommandMasterChiefPettyOfficer", 65, &militaryBranchCoastGuard, &militaryRankNonCommissionedE9}
	militaryServiceCoastGuardE9MasterChiefPettyOfficer             = EnumMilitaryServiceItem{militaryServiceCoastGuardE9MasterChiefPettyOfficerID, "Coast Guard E9 Master Chief Petty Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "CoastGuardE9MasterChiefPettyOfficer", 66, &militaryBranchCoastGuard, &militaryRankNonCommissionedE9}
	militaryServiceCoastGuardO1Ensign                              = EnumMilitaryServiceItem{militaryServiceCoastGuardO1EnsignID, "Coast Guard O1 Ensign", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO1ID)}, "CoastGuardO1Ensign", 67, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO1}
	militaryServiceCoastGuardO2LieutenantJuniorGrade               = EnumMilitaryServiceItem{militaryServiceCoastGuardO2LieutenantJuniorGradeID, "Coast Guard O2 Lieutenant Junior Grade", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO2ID)}, "CoastGuardO2LieutenantJuniorGrade", 68, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO2}
	militaryServiceCoastGuardO3Lieutenant                          = EnumMilitaryServiceItem{militaryServiceCoastGuardO3LieutenantID, "Coast Guard O3 Lieutenant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO3ID)}, "CoastGuardO3Lieutenant", 69, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO3}
	militaryServiceCoastGuardO4LieutenantCommander                 = EnumMilitaryServiceItem{militaryServiceCoastGuardO4LieutenantCommanderID, "Coast Guard O4 Lieutenant Commander", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO4ID)}, "CoastGuardO4LieutenantCommander", 70, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO4}
	militaryServiceCoastGuardO5Commander                           = EnumMilitaryServiceItem{militaryServiceCoastGuardO5CommanderID, "Coast Guard O5 Commander", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO5ID)}, "CoastGuardO5Commander", 71, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO5}
	militaryServiceCoastGuardO6Captain                             = EnumMilitaryServiceItem{militaryServiceCoastGuardO6CaptainID, "Coast Guard O6 Captain", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO6ID)}, "CoastGuardO6Captain", 72, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO6}
	militaryServiceCoastGuardO7RearAdmiral                         = EnumMilitaryServiceItem{militaryServiceCoastGuardO7RearAdmiralID, "Coast Guard O7 Rear Admiral", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO7ID)}, "CoastGuardO7RearAdmiral", 73, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO7}
	militaryServiceCoastGuardO8RearAdmiral                         = EnumMilitaryServiceItem{militaryServiceCoastGuardO8RearAdmiralID, "Coast Guard O8 Rear Admiral", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO8ID)}, "CoastGuardO8RearAdmiral", 74, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO8}
	militaryServiceCoastGuardO9ViceAdmiral                         = EnumMilitaryServiceItem{militaryServiceCoastGuardO9ViceAdmiralID, "Coast Guard O9 Vice Admiral", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO9ID)}, "CoastGuardO9ViceAdmiral", 75, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO9}
	militaryServiceCoastGuardO10Admiral                            = EnumMilitaryServiceItem{militaryServiceCoastGuardO10AdmiralID, "Coast Guard O10 Admiral", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO10ID)}, "CoastGuardO10Admiral", 76, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO10}
	militaryServiceCoastGuardO10FleetAdmiral                       = EnumMilitaryServiceItem{militaryServiceCoastGuardO10FleetAdmiralID, "Coast Guard O10 Fleet Admiral", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchCoastGuardID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO10ID)}, "CoastGuardO10FleetAdmiral", 77, &militaryBranchCoastGuard, &militaryRankCommissionedOfficerO10}
	militaryServiceMarinesE1Private                                = EnumMilitaryServiceItem{militaryServiceMarinesE1PrivateID, "Marines E1 Private", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE1ID)}, "MarinesE1Private", 78, &militaryBranchMarines, &militaryRankEnlistedE1}
	militaryServiceMarinesE2PrivateFirstClass                      = EnumMilitaryServiceItem{militaryServiceMarinesE2PrivateFirstClassID, "Marines E2 Private First Class", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE2ID)}, "MarinesE2PrivateFirstClass", 79, &militaryBranchMarines, &militaryRankEnlistedE2}
	militaryServiceMarinesE3LanceCorporal                          = EnumMilitaryServiceItem{militaryServiceMarinesE3LanceCorporalID, "Marines E3 Lance Corporal", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE3ID)}, "MarinesE3LanceCorporal", 80, &militaryBranchMarines, &militaryRankEnlistedE3}
	militaryServiceMarinesE4Corporal                               = EnumMilitaryServiceItem{militaryServiceMarinesE4CorporalID, "Marines E4 Corporal", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE4ID)}, "MarinesE4Corporal", 81, &militaryBranchMarines, &militaryRankEnlistedE4}
	militaryServiceMarinesE5Sergeant                               = EnumMilitaryServiceItem{militaryServiceMarinesE5SergeantID, "Marines E5 Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE5ID)}, "MarinesE5Sergeant", 82, &militaryBranchMarines, &militaryRankNonCommissionedE5}
	militaryServiceMarinesE6StaffSergeant                          = EnumMilitaryServiceItem{militaryServiceMarinesE6StaffSergeantID, "Marines E6 Staff Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE6ID)}, "MarinesE6StaffSergeant", 83, &militaryBranchMarines, &militaryRankNonCommissionedE6}
	militaryServiceMarinesE7GunnerySergeant                        = EnumMilitaryServiceItem{militaryServiceMarinesE7GunnerySergeantID, "Marines E7 Gunnery Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE7ID)}, "MarinesE7GunnerySergeant", 84, &militaryBranchMarines, &militaryRankNonCommissionedE7}
	militaryServiceMarinesE8FirstSergeant                          = EnumMilitaryServiceItem{militaryServiceMarinesE8FirstSergeantID, "Marines E8 First Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE8ID)}, "MarinesE8FirstSergeant", 85, &militaryBranchMarines, &militaryRankNonCommissionedE8}
	militaryServiceMarinesE8MasterSergeant                         = EnumMilitaryServiceItem{militaryServiceMarinesE8MasterSergeantID, "Marines E8 Master Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE8ID)}, "MarinesE8MasterSergeant", 86, &militaryBranchMarines, &militaryRankNonCommissionedE8}
	militaryServiceMarinesE9MasterGunnerySergeant                  = EnumMilitaryServiceItem{militaryServiceMarinesE9MasterGunnerySergeantID, "Marines E9 Master Gunnery Sergeant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "MarinesE9MasterGunnerySergeant", 87, &militaryBranchMarines, &militaryRankNonCommissionedE9}
	militaryServiceMarinesE9SergeantMajor                          = EnumMilitaryServiceItem{militaryServiceMarinesE9SergeantMajorID, "Marines E9 Sergeant Major", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "MarinesE9SergeantMajor", 88, &militaryBranchMarines, &militaryRankNonCommissionedE9}
	militaryServiceMarinesE9SergeantMajorOfTheMarineCorps          = EnumMilitaryServiceItem{militaryServiceMarinesE9SergeantMajorOfTheMarineCorpsID, "Marines E9 Sergeant Major Of The Marine Corps", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "MarinesE9SergeantMajorOfTheMarineCorps", 89, &militaryBranchMarines, &militaryRankNonCommissionedE9}
	militaryServiceMarinesO1SecondLieutenant                       = EnumMilitaryServiceItem{militaryServiceMarinesO1SecondLieutenantID, "Marines O1 Second Lieutenant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO1ID)}, "MarinesO1SecondLieutenant", 90, &militaryBranchMarines, &militaryRankCommissionedOfficerO1}
	militaryServiceMarinesO2FirstLieutenant                        = EnumMilitaryServiceItem{militaryServiceMarinesO2FirstLieutenantID, "Marines O2 First Lieutenant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO2ID)}, "MarinesO2FirstLieutenant", 91, &militaryBranchMarines, &militaryRankCommissionedOfficerO2}
	militaryServiceMarinesO3Captain                                = EnumMilitaryServiceItem{militaryServiceMarinesO3CaptainID, "Marines O3 Captain", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO3ID)}, "MarinesO3Captain", 92, &militaryBranchMarines, &militaryRankCommissionedOfficerO3}
	militaryServiceMarinesO4Major                                  = EnumMilitaryServiceItem{militaryServiceMarinesO4MajorID, "Marines O4 Major", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO4ID)}, "MarinesO4Major", 93, &militaryBranchMarines, &militaryRankCommissionedOfficerO4}
	militaryServiceMarinesO5LieutenantColonel                      = EnumMilitaryServiceItem{militaryServiceMarinesO5LieutenantColonelID, "Marines O5 Lieutenant Colonel", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO5ID)}, "MarinesO5LieutenantColonel", 94, &militaryBranchMarines, &militaryRankCommissionedOfficerO5}
	militaryServiceMarinesO6Colonel                                = EnumMilitaryServiceItem{militaryServiceMarinesO6ColonelID, "Marines O6 Colonel", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO6ID)}, "MarinesO6Colonel", 95, &militaryBranchMarines, &militaryRankCommissionedOfficerO6}
	militaryServiceMarinesO7BrigadierGeneral                       = EnumMilitaryServiceItem{militaryServiceMarinesO7BrigadierGeneralID, "Marines O7 Brigadier General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO7ID)}, "MarinesO7BrigadierGeneral", 96, &militaryBranchMarines, &militaryRankCommissionedOfficerO7}
	militaryServiceMarinesO8MajorGeneral                           = EnumMilitaryServiceItem{militaryServiceMarinesO8MajorGeneralID, "Marines O8 Major General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO8ID)}, "MarinesO8MajorGeneral", 97, &militaryBranchMarines, &militaryRankCommissionedOfficerO8}
	militaryServiceMarinesO9LieutenantGeneral                      = EnumMilitaryServiceItem{militaryServiceMarinesO9LieutenantGeneralID, "Marines O9 Lieutenant General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO9ID)}, "MarinesO9LieutenantGeneral", 98, &militaryBranchMarines, &militaryRankCommissionedOfficerO9}
	militaryServiceMarinesO10General                               = EnumMilitaryServiceItem{militaryServiceMarinesO10GeneralID, "Marines O10 General", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO10ID)}, "MarinesO10General", 99, &militaryBranchMarines, &militaryRankCommissionedOfficerO10}
	militaryServiceMarinesW1WarrantOfficer                         = EnumMilitaryServiceItem{militaryServiceMarinesW1WarrantOfficerID, "Marines W1 Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW1ID)}, "MarinesW1WarrantOfficer", 100, &militaryBranchMarines, &militaryRankWarrantOfficerOW1}
	militaryServiceMarinesW2ChiefWarrantOfficer                    = EnumMilitaryServiceItem{militaryServiceMarinesW2ChiefWarrantOfficerID, "Marines W2 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW2ID)}, "MarinesW2ChiefWarrantOfficer", 101, &militaryBranchMarines, &militaryRankWarrantOfficerOW2}
	militaryServiceMarinesW3ChiefWarrantOfficer                    = EnumMilitaryServiceItem{militaryServiceMarinesW3ChiefWarrantOfficerID, "Marines W3 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW3ID)}, "MarinesW3ChiefWarrantOfficer", 102, &militaryBranchMarines, &militaryRankWarrantOfficerOW3}
	militaryServiceMarinesW4ChiefWarrantOfficer                    = EnumMilitaryServiceItem{militaryServiceMarinesW4ChiefWarrantOfficerID, "Marines W4 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW4ID)}, "MarinesW4ChiefWarrantOfficer", 103, &militaryBranchMarines, &militaryRankWarrantOfficerOW4}
	militaryServiceMarinesW5ChiefWarrantOfficer                    = EnumMilitaryServiceItem{militaryServiceMarinesW5ChiefWarrantOfficerID, "Marines W5 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchMarinesID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW5ID)}, "MarinesW5ChiefWarrantOfficer", 104, &militaryBranchMarines, &militaryRankWarrantOfficerOW5}
	militaryServiceNavyE1SeamanRecruit                             = EnumMilitaryServiceItem{militaryServiceNavyE1SeamanRecruitID, "Navy E1 Seaman Recruit", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE1ID)}, "NavyE1SeamanRecruit", 105, &militaryBranchMarines, &militaryRankEnlistedE1}
	militaryServiceNavyE2SeamanApprentice                          = EnumMilitaryServiceItem{militaryServiceNavyE2SeamanApprenticeID, "Navy E2 Seaman Apprentice", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE2ID)}, "NavyE2SeamanApprentice", 106, &militaryBranchMarines, &militaryRankEnlistedE2}
	militaryServiceNavyE3Seaman                                    = EnumMilitaryServiceItem{militaryServiceNavyE3SeamanID, "Navy E3 Seaman", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE3ID)}, "NavyE3Seaman", 107, &militaryBranchMarines, &militaryRankEnlistedE3}
	militaryServiceNavyE4PettyOfficerThirdClass                    = EnumMilitaryServiceItem{militaryServiceNavyE4PettyOfficerThirdClassID, "Navy E4 Petty Officer Third Class", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankEnlistedE4ID)}, "NavyE4PettyOfficerThirdClass", 108, &militaryBranchMarines, &militaryRankEnlistedE4}
	militaryServiceNavyE5PettyOfficerSecondClass                   = EnumMilitaryServiceItem{militaryServiceNavyE5PettyOfficerSecondClassID, "Navy E5 Petty Officer Second Class", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE5ID)}, "NavyE5PettyOfficerSecondClass", 109, &militaryBranchMarines, &militaryRankNonCommissionedE5}
	militaryServiceNavyE6PettyOfficerFirstClass                    = EnumMilitaryServiceItem{militaryServiceNavyE6PettyOfficerFirstClassID, "Navy E6 Petty Officer First Class", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE6ID)}, "NavyE6PettyOfficerFirstClass", 110, &militaryBranchMarines, &militaryRankNonCommissionedE6}
	militaryServiceNavyE7ChiefPettyOfficer                         = EnumMilitaryServiceItem{militaryServiceNavyE7ChiefPettyOfficerID, "Navy E7 Chief Petty Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE7ID)}, "NavyE7ChiefPettyOfficer", 111, &militaryBranchMarines, &militaryRankNonCommissionedE7}
	militaryServiceNavyE8SeniorChiefPettyOfficer                   = EnumMilitaryServiceItem{militaryServiceNavyE8SeniorChiefPettyOfficerID, "Navy E8 Senior Chief Petty Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE8ID)}, "NavyE8SeniorChiefPettyOfficer", 112, &militaryBranchMarines, &militaryRankNonCommissionedE8}
	militaryServiceNavyE9MasterChiefPettyOfficer                   = EnumMilitaryServiceItem{militaryServiceNavyE9MasterChiefPettyOfficerID, "Navy E9 Master Chief Petty Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "NavyE9MasterChiefPettyOfficer", 113, &militaryBranchMarines, &militaryRankNonCommissionedE9}
	militaryServiceNavyE9MasterChiefPettyOfficerOfTheNavy          = EnumMilitaryServiceItem{militaryServiceNavyE9MasterChiefPettyOfficerOfTheNavyID, "Navy E9 Master Chief Petty Officer Of The Navy", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankNonCommissionedE9ID)}, "NavyE9MasterChiefPettyOfficerOfTheNavy", 114, &militaryBranchMarines, &militaryRankNonCommissionedE9}
	militaryServiceNavyO1Ensign                                    = EnumMilitaryServiceItem{militaryServiceNavyO1EnsignID, "Navy O1 Ensign", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO1ID)}, "NavyO1Ensign", 115, &militaryBranchMarines, &militaryRankCommissionedOfficerO1}
	militaryServiceNavyO2LieutenantJuniorGrade                     = EnumMilitaryServiceItem{militaryServiceNavyO2LieutenantJuniorGradeID, "Navy O2 Lieutenant Junior Grade", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO2ID)}, "NavyO2LieutenantJuniorGrade", 116, &militaryBranchMarines, &militaryRankCommissionedOfficerO2}
	militaryServiceNavyO3Lieutenant                                = EnumMilitaryServiceItem{militaryServiceNavyO3LieutenantID, "Navy O3 Lieutenant", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO3ID)}, "NavyO3Lieutenant", 117, &militaryBranchMarines, &militaryRankCommissionedOfficerO3}
	militaryServiceNavyO4LieutenantCommander                       = EnumMilitaryServiceItem{militaryServiceNavyO4LieutenantCommanderID, "Navy O4 Lieutenant Commander", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO4ID)}, "NavyO4LieutenantCommander", 118, &militaryBranchMarines, &militaryRankCommissionedOfficerO4}
	militaryServiceNavyO5Commander                                 = EnumMilitaryServiceItem{militaryServiceNavyO5CommanderID, "Navy O5 Commander", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO5ID)}, "NavyO5Commander", 119, &militaryBranchMarines, &militaryRankCommissionedOfficerO5}
	militaryServiceNavyO6Captain                                   = EnumMilitaryServiceItem{militaryServiceNavyO6CaptainID, "Navy O6 Captain", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO6ID)}, "NavyO6Captain", 120, &militaryBranchMarines, &militaryRankCommissionedOfficerO6}
	militaryServiceNavyO7RearAdmiral                               = EnumMilitaryServiceItem{militaryServiceNavyO7RearAdmiralID, "Navy O7 Rear Admiral", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO7ID)}, "NavyO7RearAdmiral", 121, &militaryBranchMarines, &militaryRankCommissionedOfficerO7}
	militaryServiceNavyO8RearAdmiral                               = EnumMilitaryServiceItem{militaryServiceNavyO8RearAdmiralID, "Navy O8 Rear Admiral", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO8ID)}, "NavyO8RearAdmiral", 122, &militaryBranchMarines, &militaryRankCommissionedOfficerO8}
	militaryServiceNavyO9ViceAdmiral                               = EnumMilitaryServiceItem{militaryServiceNavyO9ViceAdmiralID, "Navy O9 Vice Admiral", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO9ID)}, "NavyO9ViceAdmiral", 123, &militaryBranchMarines, &militaryRankCommissionedOfficerO9}
	militaryServiceNavyO10Admiral                                  = EnumMilitaryServiceItem{militaryServiceNavyO10AdmiralID, "Navy O10 Admiral", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankCommissionedOfficerO10ID)}, "NavyO10Admiral", 124, &militaryBranchMarines, &militaryRankCommissionedOfficerO10}
	militaryServiceNavyW1WarrantOfficer                            = EnumMilitaryServiceItem{militaryServiceNavyW1WarrantOfficerID, "Navy W1 Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW1ID)}, "NavyW1WarrantOfficer", 125, &militaryBranchMarines, &militaryRankWarrantOfficerOW1}
	militaryServiceNavyW2ChiefWarrantOfficer                       = EnumMilitaryServiceItem{militaryServiceNavyW2ChiefWarrantOfficerID, "Navy W2 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW2ID)}, "NavyW2ChiefWarrantOfficer", 126, &militaryBranchMarines, &militaryRankWarrantOfficerOW2}
	militaryServiceNavyW3ChiefWarrantOfficer                       = EnumMilitaryServiceItem{militaryServiceNavyW3ChiefWarrantOfficerID, "Navy W3 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW3ID)}, "NavyW3ChiefWarrantOfficer", 127, &militaryBranchMarines, &militaryRankWarrantOfficerOW3}
	militaryServiceNavyW4ChiefWarrantOfficer                       = EnumMilitaryServiceItem{militaryServiceNavyW4ChiefWarrantOfficerID, "Navy W4 Chief Warrant Officer", map[string]string{MilitaryServiceMetaBranchKey: string(militaryBranchNavyID), MilitaryServiceMetaRankKey: string(militaryRankWarrantOfficerOW4ID)}, "NavyW4ChiefWarrantOfficer", 128, &militaryBranchMarines, &militaryRankWarrantOfficerOW4}
)

// EnumMilitaryService is a collection of MilitaryService items
type EnumMilitaryService struct {
	Description string
	Items       []*EnumMilitaryServiceItem
	Name        string

	AirForceE1AirmanBasic                           *EnumMilitaryServiceItem
	AirForceE2Airman                                *EnumMilitaryServiceItem
	AirForceE3AirmanFirstClass                      *EnumMilitaryServiceItem
	AirForceE4SeniorAirman                          *EnumMilitaryServiceItem
	AirForceE4Sergeant                              *EnumMilitaryServiceItem
	AirForceE5StaffSergeant                         *EnumMilitaryServiceItem
	AirForceE6TechnicalSergeant                     *EnumMilitaryServiceItem
	AirForceE7FirstSergeant                         *EnumMilitaryServiceItem
	AirForceE7MasterSergeant                        *EnumMilitaryServiceItem
	AirForceE8FirstSergeant                         *EnumMilitaryServiceItem
	AirForceE8SeniorMasterSergeant                  *EnumMilitaryServiceItem
	AirForceE9ChiefMasterSergeant                   *EnumMilitaryServiceItem
	AirForceE9ChiefMasterSergeantOfTheAirForce      *EnumMilitaryServiceItem
	AirForceE9FirstSergeant                         *EnumMilitaryServiceItem
	AirForceO1SecondLieutenant                      *EnumMilitaryServiceItem
	AirForceO2FirstLieutenant                       *EnumMilitaryServiceItem
	AirForceO3Captain                               *EnumMilitaryServiceItem
	AirForceO4Major                                 *EnumMilitaryServiceItem
	AirForceO5LieutenantColonel                     *EnumMilitaryServiceItem
	AirForceO6Colonel                               *EnumMilitaryServiceItem
	AirForceO7BrigadierGeneral                      *EnumMilitaryServiceItem
	AirForceO8MajorGeneral                          *EnumMilitaryServiceItem
	AirForceO9LieutenantGeneral                     *EnumMilitaryServiceItem
	AirForceO10General                              *EnumMilitaryServiceItem
	ArmyE1Private                                   *EnumMilitaryServiceItem
	ArmyE2Private                                   *EnumMilitaryServiceItem
	ArmyE3PrivateFirstClass                         *EnumMilitaryServiceItem
	ArmyE4Corporal                                  *EnumMilitaryServiceItem
	ArmyE4Specialist                                *EnumMilitaryServiceItem
	ArmyE4PCorporal                                 *EnumMilitaryServiceItem
	ArmyE4PSpecialist                               *EnumMilitaryServiceItem
	ArmyE5Sergeant                                  *EnumMilitaryServiceItem
	ArmyE5PSergeant                                 *EnumMilitaryServiceItem
	ArmyE6StaffSergeant                             *EnumMilitaryServiceItem
	ArmyE6PStaffSergeant                            *EnumMilitaryServiceItem
	ArmyE7SergeantFirstClass                        *EnumMilitaryServiceItem
	ArmyE8FirstSergeant                             *EnumMilitaryServiceItem
	ArmyE8MasterSergeant                            *EnumMilitaryServiceItem
	ArmyE9CommandSergeantMajor                      *EnumMilitaryServiceItem
	ArmyE9SergeantMajor                             *EnumMilitaryServiceItem
	ArmyE9SergeantMajorOfTheArmy                    *EnumMilitaryServiceItem
	ArmyO1SecondLieutenant                          *EnumMilitaryServiceItem
	ArmyO2FirstLieutenant                           *EnumMilitaryServiceItem
	ArmyO3Captain                                   *EnumMilitaryServiceItem
	ArmyO4Major                                     *EnumMilitaryServiceItem
	ArmyO5LieutenantColonel                         *EnumMilitaryServiceItem
	ArmyO6Colonel                                   *EnumMilitaryServiceItem
	ArmyO7BrigadierGeneral                          *EnumMilitaryServiceItem
	ArmyO8MajorGeneral                              *EnumMilitaryServiceItem
	ArmyO9LieutenantGeneral                         *EnumMilitaryServiceItem
	ArmyO10General                                  *EnumMilitaryServiceItem
	ArmyW1WarrantOfficer                            *EnumMilitaryServiceItem
	ArmyW2ChiefWarrantOfficer                       *EnumMilitaryServiceItem
	ArmyW3ChiefWarrantOfficer                       *EnumMilitaryServiceItem
	ArmyW4ChiefWarrantOfficer                       *EnumMilitaryServiceItem
	ArmyW5ChiefWarrantOfficer                       *EnumMilitaryServiceItem
	CoastGuardE1SeamanRecruit                       *EnumMilitaryServiceItem
	CoastGuardE2SeamanApprentice                    *EnumMilitaryServiceItem
	CoastGuardE3Seaman                              *EnumMilitaryServiceItem
	CoastGuardE4PettyOfficerThirdClass              *EnumMilitaryServiceItem
	CoastGuardE5PettyOfficerSecondClass             *EnumMilitaryServiceItem
	CoastGuardE6PettyOfficerFirstClass              *EnumMilitaryServiceItem
	CoastGuardE7ChiefPettyOfficer                   *EnumMilitaryServiceItem
	CoastGuardE8SeniorChiefPettyOfficer             *EnumMilitaryServiceItem
	CoastGuardE9FleetCommandMasterChiefPettyOfficer *EnumMilitaryServiceItem
	CoastGuardE9MasterChiefPettyOfficer             *EnumMilitaryServiceItem
	CoastGuardO1Ensign                              *EnumMilitaryServiceItem
	CoastGuardO2LieutenantJuniorGrade               *EnumMilitaryServiceItem
	CoastGuardO3Lieutenant                          *EnumMilitaryServiceItem
	CoastGuardO4LieutenantCommander                 *EnumMilitaryServiceItem
	CoastGuardO5Commander                           *EnumMilitaryServiceItem
	CoastGuardO6Captain                             *EnumMilitaryServiceItem
	CoastGuardO7RearAdmiral                         *EnumMilitaryServiceItem
	CoastGuardO8RearAdmiral                         *EnumMilitaryServiceItem
	CoastGuardO9ViceAdmiral                         *EnumMilitaryServiceItem
	CoastGuardO10Admiral                            *EnumMilitaryServiceItem
	CoastGuardO10FleetAdmiral                       *EnumMilitaryServiceItem
	MarinesE1Private                                *EnumMilitaryServiceItem
	MarinesE2PrivateFirstClass                      *EnumMilitaryServiceItem
	MarinesE3LanceCorporal                          *EnumMilitaryServiceItem
	MarinesE4Corporal                               *EnumMilitaryServiceItem
	MarinesE5Sergeant                               *EnumMilitaryServiceItem
	MarinesE6StaffSergeant                          *EnumMilitaryServiceItem
	MarinesE7GunnerySergeant                        *EnumMilitaryServiceItem
	MarinesE8FirstSergeant                          *EnumMilitaryServiceItem
	MarinesE8MasterSergeant                         *EnumMilitaryServiceItem
	MarinesE9MasterGunnerySergeant                  *EnumMilitaryServiceItem
	MarinesE9SergeantMajor                          *EnumMilitaryServiceItem
	MarinesE9SergeantMajorOfTheMarineCorps          *EnumMilitaryServiceItem
	MarinesO1SecondLieutenant                       *EnumMilitaryServiceItem
	MarinesO2FirstLieutenant                        *EnumMilitaryServiceItem
	MarinesO3Captain                                *EnumMilitaryServiceItem
	MarinesO4Major                                  *EnumMilitaryServiceItem
	MarinesO5LieutenantColonel                      *EnumMilitaryServiceItem
	MarinesO6Colonel                                *EnumMilitaryServiceItem
	MarinesO7BrigadierGeneral                       *EnumMilitaryServiceItem
	MarinesO8MajorGeneral                           *EnumMilitaryServiceItem
	MarinesO9LieutenantGeneral                      *EnumMilitaryServiceItem
	MarinesO10General                               *EnumMilitaryServiceItem
	MarinesW1WarrantOfficer                         *EnumMilitaryServiceItem
	MarinesW2ChiefWarrantOfficer                    *EnumMilitaryServiceItem
	MarinesW3ChiefWarrantOfficer                    *EnumMilitaryServiceItem
	MarinesW4ChiefWarrantOfficer                    *EnumMilitaryServiceItem
	MarinesW5ChiefWarrantOfficer                    *EnumMilitaryServiceItem
	NavyE1SeamanRecruit                             *EnumMilitaryServiceItem
	NavyE2SeamanApprentice                          *EnumMilitaryServiceItem
	NavyE3Seaman                                    *EnumMilitaryServiceItem
	NavyE4PettyOfficerThirdClass                    *EnumMilitaryServiceItem
	NavyE5PettyOfficerSecondClass                   *EnumMilitaryServiceItem
	NavyE6PettyOfficerFirstClass                    *EnumMilitaryServiceItem
	NavyE7ChiefPettyOfficer                         *EnumMilitaryServiceItem
	NavyE8SeniorChiefPettyOfficer                   *EnumMilitaryServiceItem
	NavyE9MasterChiefPettyOfficer                   *EnumMilitaryServiceItem
	NavyE9MasterChiefPettyOfficerOfTheNavy          *EnumMilitaryServiceItem
	NavyO1Ensign                                    *EnumMilitaryServiceItem
	NavyO2LieutenantJuniorGrade                     *EnumMilitaryServiceItem
	NavyO3Lieutenant                                *EnumMilitaryServiceItem
	NavyO4LieutenantCommander                       *EnumMilitaryServiceItem
	NavyO5Commander                                 *EnumMilitaryServiceItem
	NavyO6Captain                                   *EnumMilitaryServiceItem
	NavyO7RearAdmiral                               *EnumMilitaryServiceItem
	NavyO8RearAdmiral                               *EnumMilitaryServiceItem
	NavyO9ViceAdmiral                               *EnumMilitaryServiceItem
	NavyO10Admiral                                  *EnumMilitaryServiceItem
	NavyW1WarrantOfficer                            *EnumMilitaryServiceItem
	NavyW2ChiefWarrantOfficer                       *EnumMilitaryServiceItem
	NavyW3ChiefWarrantOfficer                       *EnumMilitaryServiceItem
	NavyW4ChiefWarrantOfficer                       *EnumMilitaryServiceItem

	itemDict map[string]*EnumMilitaryServiceItem
}

// MilitaryService is a public singleton instance of EnumMilitaryService
// representing military ranks and branches
var MilitaryService = &EnumMilitaryService{
	Description: "military ranks and branches",
	Items: []*EnumMilitaryServiceItem{
		&militaryServiceAirForceE1AirmanBasic,
		&militaryServiceAirForceE2Airman,
		&militaryServiceAirForceE3AirmanFirstClass,
		&militaryServiceAirForceE4SeniorAirman,
		&militaryServiceAirForceE4Sergeant,
		&militaryServiceAirForceE5StaffSergeant,
		&militaryServiceAirForceE6TechnicalSergeant,
		&militaryServiceAirForceE7FirstSergeant,
		&militaryServiceAirForceE7MasterSergeant,
		&militaryServiceAirForceE8FirstSergeant,
		&militaryServiceAirForceE8SeniorMasterSergeant,
		&militaryServiceAirForceE9ChiefMasterSergeant,
		&militaryServiceAirForceE9ChiefMasterSergeantOfTheAirForce,
		&militaryServiceAirForceE9FirstSergeant,
		&militaryServiceAirForceO1SecondLieutenant,
		&militaryServiceAirForceO2FirstLieutenant,
		&militaryServiceAirForceO3Captain,
		&militaryServiceAirForceO4Major,
		&militaryServiceAirForceO5LieutenantColonel,
		&militaryServiceAirForceO6Colonel,
		&militaryServiceAirForceO7BrigadierGeneral,
		&militaryServiceAirForceO8MajorGeneral,
		&militaryServiceAirForceO9LieutenantGeneral,
		&militaryServiceAirForceO10General,
		&militaryServiceArmyE1Private,
		&militaryServiceArmyE2Private,
		&militaryServiceArmyE3PrivateFirstClass,
		&militaryServiceArmyE4Corporal,
		&militaryServiceArmyE4Specialist,
		&militaryServiceArmyE4PCorporal,
		&militaryServiceArmyE4PSpecialist,
		&militaryServiceArmyE5Sergeant,
		&militaryServiceArmyE5PSergeant,
		&militaryServiceArmyE6StaffSergeant,
		&militaryServiceArmyE6PStaffSergeant,
		&militaryServiceArmyE7SergeantFirstClass,
		&militaryServiceArmyE8FirstSergeant,
		&militaryServiceArmyE8MasterSergeant,
		&militaryServiceArmyE9CommandSergeantMajor,
		&militaryServiceArmyE9SergeantMajor,
		&militaryServiceArmyE9SergeantMajorOfTheArmy,
		&militaryServiceArmyO1SecondLieutenant,
		&militaryServiceArmyO2FirstLieutenant,
		&militaryServiceArmyO3Captain,
		&militaryServiceArmyO4Major,
		&militaryServiceArmyO5LieutenantColonel,
		&militaryServiceArmyO6Colonel,
		&militaryServiceArmyO7BrigadierGeneral,
		&militaryServiceArmyO8MajorGeneral,
		&militaryServiceArmyO9LieutenantGeneral,
		&militaryServiceArmyO10General,
		&militaryServiceArmyW1WarrantOfficer,
		&militaryServiceArmyW2ChiefWarrantOfficer,
		&militaryServiceArmyW3ChiefWarrantOfficer,
		&militaryServiceArmyW4ChiefWarrantOfficer,
		&militaryServiceArmyW5ChiefWarrantOfficer,
		&militaryServiceCoastGuardE1SeamanRecruit,
		&militaryServiceCoastGuardE2SeamanApprentice,
		&militaryServiceCoastGuardE3Seaman,
		&militaryServiceCoastGuardE4PettyOfficerThirdClass,
		&militaryServiceCoastGuardE5PettyOfficerSecondClass,
		&militaryServiceCoastGuardE6PettyOfficerFirstClass,
		&militaryServiceCoastGuardE7ChiefPettyOfficer,
		&militaryServiceCoastGuardE8SeniorChiefPettyOfficer,
		&militaryServiceCoastGuardE9FleetCommandMasterChiefPettyOfficer,
		&militaryServiceCoastGuardE9MasterChiefPettyOfficer,
		&militaryServiceCoastGuardO1Ensign,
		&militaryServiceCoastGuardO2LieutenantJuniorGrade,
		&militaryServiceCoastGuardO3Lieutenant,
		&militaryServiceCoastGuardO4LieutenantCommander,
		&militaryServiceCoastGuardO5Commander,
		&militaryServiceCoastGuardO6Captain,
		&militaryServiceCoastGuardO7RearAdmiral,
		&militaryServiceCoastGuardO8RearAdmiral,
		&militaryServiceCoastGuardO9ViceAdmiral,
		&militaryServiceCoastGuardO10Admiral,
		&militaryServiceCoastGuardO10FleetAdmiral,
		&militaryServiceMarinesE1Private,
		&militaryServiceMarinesE2PrivateFirstClass,
		&militaryServiceMarinesE3LanceCorporal,
		&militaryServiceMarinesE4Corporal,
		&militaryServiceMarinesE5Sergeant,
		&militaryServiceMarinesE6StaffSergeant,
		&militaryServiceMarinesE7GunnerySergeant,
		&militaryServiceMarinesE8FirstSergeant,
		&militaryServiceMarinesE8MasterSergeant,
		&militaryServiceMarinesE9MasterGunnerySergeant,
		&militaryServiceMarinesE9SergeantMajor,
		&militaryServiceMarinesE9SergeantMajorOfTheMarineCorps,
		&militaryServiceMarinesO1SecondLieutenant,
		&militaryServiceMarinesO2FirstLieutenant,
		&militaryServiceMarinesO3Captain,
		&militaryServiceMarinesO4Major,
		&militaryServiceMarinesO5LieutenantColonel,
		&militaryServiceMarinesO6Colonel,
		&militaryServiceMarinesO7BrigadierGeneral,
		&militaryServiceMarinesO8MajorGeneral,
		&militaryServiceMarinesO9LieutenantGeneral,
		&militaryServiceMarinesO10General,
		&militaryServiceMarinesW1WarrantOfficer,
		&militaryServiceMarinesW2ChiefWarrantOfficer,
		&militaryServiceMarinesW3ChiefWarrantOfficer,
		&militaryServiceMarinesW4ChiefWarrantOfficer,
		&militaryServiceMarinesW5ChiefWarrantOfficer,
		&militaryServiceNavyE1SeamanRecruit,
		&militaryServiceNavyE2SeamanApprentice,
		&militaryServiceNavyE3Seaman,
		&militaryServiceNavyE4PettyOfficerThirdClass,
		&militaryServiceNavyE5PettyOfficerSecondClass,
		&militaryServiceNavyE6PettyOfficerFirstClass,
		&militaryServiceNavyE7ChiefPettyOfficer,
		&militaryServiceNavyE8SeniorChiefPettyOfficer,
		&militaryServiceNavyE9MasterChiefPettyOfficer,
		&militaryServiceNavyE9MasterChiefPettyOfficerOfTheNavy,
		&militaryServiceNavyO1Ensign,
		&militaryServiceNavyO2LieutenantJuniorGrade,
		&militaryServiceNavyO3Lieutenant,
		&militaryServiceNavyO4LieutenantCommander,
		&militaryServiceNavyO5Commander,
		&militaryServiceNavyO6Captain,
		&militaryServiceNavyO7RearAdmiral,
		&militaryServiceNavyO8RearAdmiral,
		&militaryServiceNavyO9ViceAdmiral,
		&militaryServiceNavyO10Admiral,
		&militaryServiceNavyW1WarrantOfficer,
		&militaryServiceNavyW2ChiefWarrantOfficer,
		&militaryServiceNavyW3ChiefWarrantOfficer,
		&militaryServiceNavyW4ChiefWarrantOfficer,
	},
	Name:                                            "EnumMilitaryService",
	AirForceE1AirmanBasic:                           &militaryServiceAirForceE1AirmanBasic,
	AirForceE2Airman:                                &militaryServiceAirForceE2Airman,
	AirForceE3AirmanFirstClass:                      &militaryServiceAirForceE3AirmanFirstClass,
	AirForceE4SeniorAirman:                          &militaryServiceAirForceE4SeniorAirman,
	AirForceE4Sergeant:                              &militaryServiceAirForceE4Sergeant,
	AirForceE5StaffSergeant:                         &militaryServiceAirForceE5StaffSergeant,
	AirForceE6TechnicalSergeant:                     &militaryServiceAirForceE6TechnicalSergeant,
	AirForceE7FirstSergeant:                         &militaryServiceAirForceE7FirstSergeant,
	AirForceE7MasterSergeant:                        &militaryServiceAirForceE7MasterSergeant,
	AirForceE8FirstSergeant:                         &militaryServiceAirForceE8FirstSergeant,
	AirForceE8SeniorMasterSergeant:                  &militaryServiceAirForceE8SeniorMasterSergeant,
	AirForceE9ChiefMasterSergeant:                   &militaryServiceAirForceE9ChiefMasterSergeant,
	AirForceE9ChiefMasterSergeantOfTheAirForce:      &militaryServiceAirForceE9ChiefMasterSergeantOfTheAirForce,
	AirForceE9FirstSergeant:                         &militaryServiceAirForceE9FirstSergeant,
	AirForceO1SecondLieutenant:                      &militaryServiceAirForceO1SecondLieutenant,
	AirForceO2FirstLieutenant:                       &militaryServiceAirForceO2FirstLieutenant,
	AirForceO3Captain:                               &militaryServiceAirForceO3Captain,
	AirForceO4Major:                                 &militaryServiceAirForceO4Major,
	AirForceO5LieutenantColonel:                     &militaryServiceAirForceO5LieutenantColonel,
	AirForceO6Colonel:                               &militaryServiceAirForceO6Colonel,
	AirForceO7BrigadierGeneral:                      &militaryServiceAirForceO7BrigadierGeneral,
	AirForceO8MajorGeneral:                          &militaryServiceAirForceO8MajorGeneral,
	AirForceO9LieutenantGeneral:                     &militaryServiceAirForceO9LieutenantGeneral,
	AirForceO10General:                              &militaryServiceAirForceO10General,
	ArmyE1Private:                                   &militaryServiceArmyE1Private,
	ArmyE2Private:                                   &militaryServiceArmyE2Private,
	ArmyE3PrivateFirstClass:                         &militaryServiceArmyE3PrivateFirstClass,
	ArmyE4Corporal:                                  &militaryServiceArmyE4Corporal,
	ArmyE4Specialist:                                &militaryServiceArmyE4Specialist,
	ArmyE4PCorporal:                                 &militaryServiceArmyE4PCorporal,
	ArmyE4PSpecialist:                               &militaryServiceArmyE4PSpecialist,
	ArmyE5Sergeant:                                  &militaryServiceArmyE5Sergeant,
	ArmyE5PSergeant:                                 &militaryServiceArmyE5PSergeant,
	ArmyE6StaffSergeant:                             &militaryServiceArmyE6StaffSergeant,
	ArmyE6PStaffSergeant:                            &militaryServiceArmyE6PStaffSergeant,
	ArmyE7SergeantFirstClass:                        &militaryServiceArmyE7SergeantFirstClass,
	ArmyE8FirstSergeant:                             &militaryServiceArmyE8FirstSergeant,
	ArmyE8MasterSergeant:                            &militaryServiceArmyE8MasterSergeant,
	ArmyE9CommandSergeantMajor:                      &militaryServiceArmyE9CommandSergeantMajor,
	ArmyE9SergeantMajor:                             &militaryServiceArmyE9SergeantMajor,
	ArmyE9SergeantMajorOfTheArmy:                    &militaryServiceArmyE9SergeantMajorOfTheArmy,
	ArmyO1SecondLieutenant:                          &militaryServiceArmyO1SecondLieutenant,
	ArmyO2FirstLieutenant:                           &militaryServiceArmyO2FirstLieutenant,
	ArmyO3Captain:                                   &militaryServiceArmyO3Captain,
	ArmyO4Major:                                     &militaryServiceArmyO4Major,
	ArmyO5LieutenantColonel:                         &militaryServiceArmyO5LieutenantColonel,
	ArmyO6Colonel:                                   &militaryServiceArmyO6Colonel,
	ArmyO7BrigadierGeneral:                          &militaryServiceArmyO7BrigadierGeneral,
	ArmyO8MajorGeneral:                              &militaryServiceArmyO8MajorGeneral,
	ArmyO9LieutenantGeneral:                         &militaryServiceArmyO9LieutenantGeneral,
	ArmyO10General:                                  &militaryServiceArmyO10General,
	ArmyW1WarrantOfficer:                            &militaryServiceArmyW1WarrantOfficer,
	ArmyW2ChiefWarrantOfficer:                       &militaryServiceArmyW2ChiefWarrantOfficer,
	ArmyW3ChiefWarrantOfficer:                       &militaryServiceArmyW3ChiefWarrantOfficer,
	ArmyW4ChiefWarrantOfficer:                       &militaryServiceArmyW4ChiefWarrantOfficer,
	ArmyW5ChiefWarrantOfficer:                       &militaryServiceArmyW5ChiefWarrantOfficer,
	CoastGuardE1SeamanRecruit:                       &militaryServiceCoastGuardE1SeamanRecruit,
	CoastGuardE2SeamanApprentice:                    &militaryServiceCoastGuardE2SeamanApprentice,
	CoastGuardE3Seaman:                              &militaryServiceCoastGuardE3Seaman,
	CoastGuardE4PettyOfficerThirdClass:              &militaryServiceCoastGuardE4PettyOfficerThirdClass,
	CoastGuardE5PettyOfficerSecondClass:             &militaryServiceCoastGuardE5PettyOfficerSecondClass,
	CoastGuardE6PettyOfficerFirstClass:              &militaryServiceCoastGuardE6PettyOfficerFirstClass,
	CoastGuardE7ChiefPettyOfficer:                   &militaryServiceCoastGuardE7ChiefPettyOfficer,
	CoastGuardE8SeniorChiefPettyOfficer:             &militaryServiceCoastGuardE8SeniorChiefPettyOfficer,
	CoastGuardE9FleetCommandMasterChiefPettyOfficer: &militaryServiceCoastGuardE9FleetCommandMasterChiefPettyOfficer,
	CoastGuardE9MasterChiefPettyOfficer:             &militaryServiceCoastGuardE9MasterChiefPettyOfficer,
	CoastGuardO1Ensign:                              &militaryServiceCoastGuardO1Ensign,
	CoastGuardO2LieutenantJuniorGrade:               &militaryServiceCoastGuardO2LieutenantJuniorGrade,
	CoastGuardO3Lieutenant:                          &militaryServiceCoastGuardO3Lieutenant,
	CoastGuardO4LieutenantCommander:                 &militaryServiceCoastGuardO4LieutenantCommander,
	CoastGuardO5Commander:                           &militaryServiceCoastGuardO5Commander,
	CoastGuardO6Captain:                             &militaryServiceCoastGuardO6Captain,
	CoastGuardO7RearAdmiral:                         &militaryServiceCoastGuardO7RearAdmiral,
	CoastGuardO8RearAdmiral:                         &militaryServiceCoastGuardO8RearAdmiral,
	CoastGuardO9ViceAdmiral:                         &militaryServiceCoastGuardO9ViceAdmiral,
	CoastGuardO10Admiral:                            &militaryServiceCoastGuardO10Admiral,
	CoastGuardO10FleetAdmiral:                       &militaryServiceCoastGuardO10FleetAdmiral,
	MarinesE1Private:                                &militaryServiceMarinesE1Private,
	MarinesE2PrivateFirstClass:                      &militaryServiceMarinesE2PrivateFirstClass,
	MarinesE3LanceCorporal:                          &militaryServiceMarinesE3LanceCorporal,
	MarinesE4Corporal:                               &militaryServiceMarinesE4Corporal,
	MarinesE5Sergeant:                               &militaryServiceMarinesE5Sergeant,
	MarinesE6StaffSergeant:                          &militaryServiceMarinesE6StaffSergeant,
	MarinesE7GunnerySergeant:                        &militaryServiceMarinesE7GunnerySergeant,
	MarinesE8FirstSergeant:                          &militaryServiceMarinesE8FirstSergeant,
	MarinesE8MasterSergeant:                         &militaryServiceMarinesE8MasterSergeant,
	MarinesE9MasterGunnerySergeant:                  &militaryServiceMarinesE9MasterGunnerySergeant,
	MarinesE9SergeantMajor:                          &militaryServiceMarinesE9SergeantMajor,
	MarinesE9SergeantMajorOfTheMarineCorps:          &militaryServiceMarinesE9SergeantMajorOfTheMarineCorps,
	MarinesO1SecondLieutenant:                       &militaryServiceMarinesO1SecondLieutenant,
	MarinesO2FirstLieutenant:                        &militaryServiceMarinesO2FirstLieutenant,
	MarinesO3Captain:                                &militaryServiceMarinesO3Captain,
	MarinesO4Major:                                  &militaryServiceMarinesO4Major,
	MarinesO5LieutenantColonel:                      &militaryServiceMarinesO5LieutenantColonel,
	MarinesO6Colonel:                                &militaryServiceMarinesO6Colonel,
	MarinesO7BrigadierGeneral:                       &militaryServiceMarinesO7BrigadierGeneral,
	MarinesO8MajorGeneral:                           &militaryServiceMarinesO8MajorGeneral,
	MarinesO9LieutenantGeneral:                      &militaryServiceMarinesO9LieutenantGeneral,
	MarinesO10General:                               &militaryServiceMarinesO10General,
	MarinesW1WarrantOfficer:                         &militaryServiceMarinesW1WarrantOfficer,
	MarinesW2ChiefWarrantOfficer:                    &militaryServiceMarinesW2ChiefWarrantOfficer,
	MarinesW3ChiefWarrantOfficer:                    &militaryServiceMarinesW3ChiefWarrantOfficer,
	MarinesW4ChiefWarrantOfficer:                    &militaryServiceMarinesW4ChiefWarrantOfficer,
	MarinesW5ChiefWarrantOfficer:                    &militaryServiceMarinesW5ChiefWarrantOfficer,
	NavyE1SeamanRecruit:                             &militaryServiceNavyE1SeamanRecruit,
	NavyE2SeamanApprentice:                          &militaryServiceNavyE2SeamanApprentice,
	NavyE3Seaman:                                    &militaryServiceNavyE3Seaman,
	NavyE4PettyOfficerThirdClass:                    &militaryServiceNavyE4PettyOfficerThirdClass,
	NavyE5PettyOfficerSecondClass:                   &militaryServiceNavyE5PettyOfficerSecondClass,
	NavyE6PettyOfficerFirstClass:                    &militaryServiceNavyE6PettyOfficerFirstClass,
	NavyE7ChiefPettyOfficer:                         &militaryServiceNavyE7ChiefPettyOfficer,
	NavyE8SeniorChiefPettyOfficer:                   &militaryServiceNavyE8SeniorChiefPettyOfficer,
	NavyE9MasterChiefPettyOfficer:                   &militaryServiceNavyE9MasterChiefPettyOfficer,
	NavyE9MasterChiefPettyOfficerOfTheNavy:          &militaryServiceNavyE9MasterChiefPettyOfficerOfTheNavy,
	NavyO1Ensign:                                    &militaryServiceNavyO1Ensign,
	NavyO2LieutenantJuniorGrade:                     &militaryServiceNavyO2LieutenantJuniorGrade,
	NavyO3Lieutenant:                                &militaryServiceNavyO3Lieutenant,
	NavyO4LieutenantCommander:                       &militaryServiceNavyO4LieutenantCommander,
	NavyO5Commander:                                 &militaryServiceNavyO5Commander,
	NavyO6Captain:                                   &militaryServiceNavyO6Captain,
	NavyO7RearAdmiral:                               &militaryServiceNavyO7RearAdmiral,
	NavyO8RearAdmiral:                               &militaryServiceNavyO8RearAdmiral,
	NavyO9ViceAdmiral:                               &militaryServiceNavyO9ViceAdmiral,
	NavyO10Admiral:                                  &militaryServiceNavyO10Admiral,
	NavyW1WarrantOfficer:                            &militaryServiceNavyW1WarrantOfficer,
	NavyW2ChiefWarrantOfficer:                       &militaryServiceNavyW2ChiefWarrantOfficer,
	NavyW3ChiefWarrantOfficer:                       &militaryServiceNavyW3ChiefWarrantOfficer,
	NavyW4ChiefWarrantOfficer:                       &militaryServiceNavyW4ChiefWarrantOfficer,

	itemDict: map[string]*EnumMilitaryServiceItem{
		strings.ToLower(string(militaryServiceAirForceE1AirmanBasicID)):                           &militaryServiceAirForceE1AirmanBasic,
		strings.ToLower(string(militaryServiceAirForceE2AirmanID)):                                &militaryServiceAirForceE2Airman,
		strings.ToLower(string(militaryServiceAirForceE3AirmanFirstClassID)):                      &militaryServiceAirForceE3AirmanFirstClass,
		strings.ToLower(string(militaryServiceAirForceE4SeniorAirmanID)):                          &militaryServiceAirForceE4SeniorAirman,
		strings.ToLower(string(militaryServiceAirForceE4SergeantID)):                              &militaryServiceAirForceE4Sergeant,
		strings.ToLower(string(militaryServiceAirForceE5StaffSergeantID)):                         &militaryServiceAirForceE5StaffSergeant,
		strings.ToLower(string(militaryServiceAirForceE6TechnicalSergeantID)):                     &militaryServiceAirForceE6TechnicalSergeant,
		strings.ToLower(string(militaryServiceAirForceE7FirstSergeantID)):                         &militaryServiceAirForceE7FirstSergeant,
		strings.ToLower(string(militaryServiceAirForceE7MasterSergeantID)):                        &militaryServiceAirForceE7MasterSergeant,
		strings.ToLower(string(militaryServiceAirForceE8FirstSergeantID)):                         &militaryServiceAirForceE8FirstSergeant,
		strings.ToLower(string(militaryServiceAirForceE8SeniorMasterSergeantID)):                  &militaryServiceAirForceE8SeniorMasterSergeant,
		strings.ToLower(string(militaryServiceAirForceE9ChiefMasterSergeantID)):                   &militaryServiceAirForceE9ChiefMasterSergeant,
		strings.ToLower(string(militaryServiceAirForceE9ChiefMasterSergeantOfTheAirForceID)):      &militaryServiceAirForceE9ChiefMasterSergeantOfTheAirForce,
		strings.ToLower(string(militaryServiceAirForceE9FirstSergeantID)):                         &militaryServiceAirForceE9FirstSergeant,
		strings.ToLower(string(militaryServiceAirForceO1SecondLieutenantID)):                      &militaryServiceAirForceO1SecondLieutenant,
		strings.ToLower(string(militaryServiceAirForceO2FirstLieutenantID)):                       &militaryServiceAirForceO2FirstLieutenant,
		strings.ToLower(string(militaryServiceAirForceO3CaptainID)):                               &militaryServiceAirForceO3Captain,
		strings.ToLower(string(militaryServiceAirForceO4MajorID)):                                 &militaryServiceAirForceO4Major,
		strings.ToLower(string(militaryServiceAirForceO5LieutenantColonelID)):                     &militaryServiceAirForceO5LieutenantColonel,
		strings.ToLower(string(militaryServiceAirForceO6ColonelID)):                               &militaryServiceAirForceO6Colonel,
		strings.ToLower(string(militaryServiceAirForceO7BrigadierGeneralID)):                      &militaryServiceAirForceO7BrigadierGeneral,
		strings.ToLower(string(militaryServiceAirForceO8MajorGeneralID)):                          &militaryServiceAirForceO8MajorGeneral,
		strings.ToLower(string(militaryServiceAirForceO9LieutenantGeneralID)):                     &militaryServiceAirForceO9LieutenantGeneral,
		strings.ToLower(string(militaryServiceAirForceO10GeneralID)):                              &militaryServiceAirForceO10General,
		strings.ToLower(string(militaryServiceArmyE1PrivateID)):                                   &militaryServiceArmyE1Private,
		strings.ToLower(string(militaryServiceArmyE2PrivateID)):                                   &militaryServiceArmyE2Private,
		strings.ToLower(string(militaryServiceArmyE3PrivateFirstClassID)):                         &militaryServiceArmyE3PrivateFirstClass,
		strings.ToLower(string(militaryServiceArmyE4CorporalID)):                                  &militaryServiceArmyE4Corporal,
		strings.ToLower(string(militaryServiceArmyE4SpecialistID)):                                &militaryServiceArmyE4Specialist,
		strings.ToLower(string(militaryServiceArmyE4PCorporalID)):                                 &militaryServiceArmyE4PCorporal,
		strings.ToLower(string(militaryServiceArmyE4PSpecialistID)):                               &militaryServiceArmyE4PSpecialist,
		strings.ToLower(string(militaryServiceArmyE5SergeantID)):                                  &militaryServiceArmyE5Sergeant,
		strings.ToLower(string(militaryServiceArmyE5PSergeantID)):                                 &militaryServiceArmyE5PSergeant,
		strings.ToLower(string(militaryServiceArmyE6StaffSergeantID)):                             &militaryServiceArmyE6StaffSergeant,
		strings.ToLower(string(militaryServiceArmyE6PStaffSergeantID)):                            &militaryServiceArmyE6PStaffSergeant,
		strings.ToLower(string(militaryServiceArmyE7SergeantFirstClassID)):                        &militaryServiceArmyE7SergeantFirstClass,
		strings.ToLower(string(militaryServiceArmyE8FirstSergeantID)):                             &militaryServiceArmyE8FirstSergeant,
		strings.ToLower(string(militaryServiceArmyE8MasterSergeantID)):                            &militaryServiceArmyE8MasterSergeant,
		strings.ToLower(string(militaryServiceArmyE9CommandSergeantMajorID)):                      &militaryServiceArmyE9CommandSergeantMajor,
		strings.ToLower(string(militaryServiceArmyE9SergeantMajorID)):                             &militaryServiceArmyE9SergeantMajor,
		strings.ToLower(string(militaryServiceArmyE9SergeantMajorOfTheArmyID)):                    &militaryServiceArmyE9SergeantMajorOfTheArmy,
		strings.ToLower(string(militaryServiceArmyO1SecondLieutenantID)):                          &militaryServiceArmyO1SecondLieutenant,
		strings.ToLower(string(militaryServiceArmyO2FirstLieutenantID)):                           &militaryServiceArmyO2FirstLieutenant,
		strings.ToLower(string(militaryServiceArmyO3CaptainID)):                                   &militaryServiceArmyO3Captain,
		strings.ToLower(string(militaryServiceArmyO4MajorID)):                                     &militaryServiceArmyO4Major,
		strings.ToLower(string(militaryServiceArmyO5LieutenantColonelID)):                         &militaryServiceArmyO5LieutenantColonel,
		strings.ToLower(string(militaryServiceArmyO6ColonelID)):                                   &militaryServiceArmyO6Colonel,
		strings.ToLower(string(militaryServiceArmyO7BrigadierGeneralID)):                          &militaryServiceArmyO7BrigadierGeneral,
		strings.ToLower(string(militaryServiceArmyO8MajorGeneralID)):                              &militaryServiceArmyO8MajorGeneral,
		strings.ToLower(string(militaryServiceArmyO9LieutenantGeneralID)):                         &militaryServiceArmyO9LieutenantGeneral,
		strings.ToLower(string(militaryServiceArmyO10GeneralID)):                                  &militaryServiceArmyO10General,
		strings.ToLower(string(militaryServiceArmyW1WarrantOfficerID)):                            &militaryServiceArmyW1WarrantOfficer,
		strings.ToLower(string(militaryServiceArmyW2ChiefWarrantOfficerID)):                       &militaryServiceArmyW2ChiefWarrantOfficer,
		strings.ToLower(string(militaryServiceArmyW3ChiefWarrantOfficerID)):                       &militaryServiceArmyW3ChiefWarrantOfficer,
		strings.ToLower(string(militaryServiceArmyW4ChiefWarrantOfficerID)):                       &militaryServiceArmyW4ChiefWarrantOfficer,
		strings.ToLower(string(militaryServiceArmyW5ChiefWarrantOfficerID)):                       &militaryServiceArmyW5ChiefWarrantOfficer,
		strings.ToLower(string(militaryServiceCoastGuardE1SeamanRecruitID)):                       &militaryServiceCoastGuardE1SeamanRecruit,
		strings.ToLower(string(militaryServiceCoastGuardE2SeamanApprenticeID)):                    &militaryServiceCoastGuardE2SeamanApprentice,
		strings.ToLower(string(militaryServiceCoastGuardE3SeamanID)):                              &militaryServiceCoastGuardE3Seaman,
		strings.ToLower(string(militaryServiceCoastGuardE4PettyOfficerThirdClassID)):              &militaryServiceCoastGuardE4PettyOfficerThirdClass,
		strings.ToLower(string(militaryServiceCoastGuardE5PettyOfficerSecondClassID)):             &militaryServiceCoastGuardE5PettyOfficerSecondClass,
		strings.ToLower(string(militaryServiceCoastGuardE6PettyOfficerFirstClassID)):              &militaryServiceCoastGuardE6PettyOfficerFirstClass,
		strings.ToLower(string(militaryServiceCoastGuardE7ChiefPettyOfficerID)):                   &militaryServiceCoastGuardE7ChiefPettyOfficer,
		strings.ToLower(string(militaryServiceCoastGuardE8SeniorChiefPettyOfficerID)):             &militaryServiceCoastGuardE8SeniorChiefPettyOfficer,
		strings.ToLower(string(militaryServiceCoastGuardE9FleetCommandMasterChiefPettyOfficerID)): &militaryServiceCoastGuardE9FleetCommandMasterChiefPettyOfficer,
		strings.ToLower(string(militaryServiceCoastGuardE9MasterChiefPettyOfficerID)):             &militaryServiceCoastGuardE9MasterChiefPettyOfficer,
		strings.ToLower(string(militaryServiceCoastGuardO1EnsignID)):                              &militaryServiceCoastGuardO1Ensign,
		strings.ToLower(string(militaryServiceCoastGuardO2LieutenantJuniorGradeID)):               &militaryServiceCoastGuardO2LieutenantJuniorGrade,
		strings.ToLower(string(militaryServiceCoastGuardO3LieutenantID)):                          &militaryServiceCoastGuardO3Lieutenant,
		strings.ToLower(string(militaryServiceCoastGuardO4LieutenantCommanderID)):                 &militaryServiceCoastGuardO4LieutenantCommander,
		strings.ToLower(string(militaryServiceCoastGuardO5CommanderID)):                           &militaryServiceCoastGuardO5Commander,
		strings.ToLower(string(militaryServiceCoastGuardO6CaptainID)):                             &militaryServiceCoastGuardO6Captain,
		strings.ToLower(string(militaryServiceCoastGuardO7RearAdmiralID)):                         &militaryServiceCoastGuardO7RearAdmiral,
		strings.ToLower(string(militaryServiceCoastGuardO8RearAdmiralID)):                         &militaryServiceCoastGuardO8RearAdmiral,
		strings.ToLower(string(militaryServiceCoastGuardO9ViceAdmiralID)):                         &militaryServiceCoastGuardO9ViceAdmiral,
		strings.ToLower(string(militaryServiceCoastGuardO10AdmiralID)):                            &militaryServiceCoastGuardO10Admiral,
		strings.ToLower(string(militaryServiceCoastGuardO10FleetAdmiralID)):                       &militaryServiceCoastGuardO10FleetAdmiral,
		strings.ToLower(string(militaryServiceMarinesE1PrivateID)):                                &militaryServiceMarinesE1Private,
		strings.ToLower(string(militaryServiceMarinesE2PrivateFirstClassID)):                      &militaryServiceMarinesE2PrivateFirstClass,
		strings.ToLower(string(militaryServiceMarinesE3LanceCorporalID)):                          &militaryServiceMarinesE3LanceCorporal,
		strings.ToLower(string(militaryServiceMarinesE4CorporalID)):                               &militaryServiceMarinesE4Corporal,
		strings.ToLower(string(militaryServiceMarinesE5SergeantID)):                               &militaryServiceMarinesE5Sergeant,
		strings.ToLower(string(militaryServiceMarinesE6StaffSergeantID)):                          &militaryServiceMarinesE6StaffSergeant,
		strings.ToLower(string(militaryServiceMarinesE7GunnerySergeantID)):                        &militaryServiceMarinesE7GunnerySergeant,
		strings.ToLower(string(militaryServiceMarinesE8FirstSergeantID)):                          &militaryServiceMarinesE8FirstSergeant,
		strings.ToLower(string(militaryServiceMarinesE8MasterSergeantID)):                         &militaryServiceMarinesE8MasterSergeant,
		strings.ToLower(string(militaryServiceMarinesE9MasterGunnerySergeantID)):                  &militaryServiceMarinesE9MasterGunnerySergeant,
		strings.ToLower(string(militaryServiceMarinesE9SergeantMajorID)):                          &militaryServiceMarinesE9SergeantMajor,
		strings.ToLower(string(militaryServiceMarinesE9SergeantMajorOfTheMarineCorpsID)):          &militaryServiceMarinesE9SergeantMajorOfTheMarineCorps,
		strings.ToLower(string(militaryServiceMarinesO1SecondLieutenantID)):                       &militaryServiceMarinesO1SecondLieutenant,
		strings.ToLower(string(militaryServiceMarinesO2FirstLieutenantID)):                        &militaryServiceMarinesO2FirstLieutenant,
		strings.ToLower(string(militaryServiceMarinesO3CaptainID)):                                &militaryServiceMarinesO3Captain,
		strings.ToLower(string(militaryServiceMarinesO4MajorID)):                                  &militaryServiceMarinesO4Major,
		strings.ToLower(string(militaryServiceMarinesO5LieutenantColonelID)):                      &militaryServiceMarinesO5LieutenantColonel,
		strings.ToLower(string(militaryServiceMarinesO6ColonelID)):                                &militaryServiceMarinesO6Colonel,
		strings.ToLower(string(militaryServiceMarinesO7BrigadierGeneralID)):                       &militaryServiceMarinesO7BrigadierGeneral,
		strings.ToLower(string(militaryServiceMarinesO8MajorGeneralID)):                           &militaryServiceMarinesO8MajorGeneral,
		strings.ToLower(string(militaryServiceMarinesO9LieutenantGeneralID)):                      &militaryServiceMarinesO9LieutenantGeneral,
		strings.ToLower(string(militaryServiceMarinesO10GeneralID)):                               &militaryServiceMarinesO10General,
		strings.ToLower(string(militaryServiceMarinesW1WarrantOfficerID)):                         &militaryServiceMarinesW1WarrantOfficer,
		strings.ToLower(string(militaryServiceMarinesW2ChiefWarrantOfficerID)):                    &militaryServiceMarinesW2ChiefWarrantOfficer,
		strings.ToLower(string(militaryServiceMarinesW3ChiefWarrantOfficerID)):                    &militaryServiceMarinesW3ChiefWarrantOfficer,
		strings.ToLower(string(militaryServiceMarinesW4ChiefWarrantOfficerID)):                    &militaryServiceMarinesW4ChiefWarrantOfficer,
		strings.ToLower(string(militaryServiceMarinesW5ChiefWarrantOfficerID)):                    &militaryServiceMarinesW5ChiefWarrantOfficer,
		strings.ToLower(string(militaryServiceNavyE1SeamanRecruitID)):                             &militaryServiceNavyE1SeamanRecruit,
		strings.ToLower(string(militaryServiceNavyE2SeamanApprenticeID)):                          &militaryServiceNavyE2SeamanApprentice,
		strings.ToLower(string(militaryServiceNavyE3SeamanID)):                                    &militaryServiceNavyE3Seaman,
		strings.ToLower(string(militaryServiceNavyE4PettyOfficerThirdClassID)):                    &militaryServiceNavyE4PettyOfficerThirdClass,
		strings.ToLower(string(militaryServiceNavyE5PettyOfficerSecondClassID)):                   &militaryServiceNavyE5PettyOfficerSecondClass,
		strings.ToLower(string(militaryServiceNavyE6PettyOfficerFirstClassID)):                    &militaryServiceNavyE6PettyOfficerFirstClass,
		strings.ToLower(string(militaryServiceNavyE7ChiefPettyOfficerID)):                         &militaryServiceNavyE7ChiefPettyOfficer,
		strings.ToLower(string(militaryServiceNavyE8SeniorChiefPettyOfficerID)):                   &militaryServiceNavyE8SeniorChiefPettyOfficer,
		strings.ToLower(string(militaryServiceNavyE9MasterChiefPettyOfficerID)):                   &militaryServiceNavyE9MasterChiefPettyOfficer,
		strings.ToLower(string(militaryServiceNavyE9MasterChiefPettyOfficerOfTheNavyID)):          &militaryServiceNavyE9MasterChiefPettyOfficerOfTheNavy,
		strings.ToLower(string(militaryServiceNavyO1EnsignID)):                                    &militaryServiceNavyO1Ensign,
		strings.ToLower(string(militaryServiceNavyO2LieutenantJuniorGradeID)):                     &militaryServiceNavyO2LieutenantJuniorGrade,
		strings.ToLower(string(militaryServiceNavyO3LieutenantID)):                                &militaryServiceNavyO3Lieutenant,
		strings.ToLower(string(militaryServiceNavyO4LieutenantCommanderID)):                       &militaryServiceNavyO4LieutenantCommander,
		strings.ToLower(string(militaryServiceNavyO5CommanderID)):                                 &militaryServiceNavyO5Commander,
		strings.ToLower(string(militaryServiceNavyO6CaptainID)):                                   &militaryServiceNavyO6Captain,
		strings.ToLower(string(militaryServiceNavyO7RearAdmiralID)):                               &militaryServiceNavyO7RearAdmiral,
		strings.ToLower(string(militaryServiceNavyO8RearAdmiralID)):                               &militaryServiceNavyO8RearAdmiral,
		strings.ToLower(string(militaryServiceNavyO9ViceAdmiralID)):                               &militaryServiceNavyO9ViceAdmiral,
		strings.ToLower(string(militaryServiceNavyO10AdmiralID)):                                  &militaryServiceNavyO10Admiral,
		strings.ToLower(string(militaryServiceNavyW1WarrantOfficerID)):                            &militaryServiceNavyW1WarrantOfficer,
		strings.ToLower(string(militaryServiceNavyW2ChiefWarrantOfficerID)):                       &militaryServiceNavyW2ChiefWarrantOfficer,
		strings.ToLower(string(militaryServiceNavyW3ChiefWarrantOfficerID)):                       &militaryServiceNavyW3ChiefWarrantOfficer,
		strings.ToLower(string(militaryServiceNavyW4ChiefWarrantOfficerID)):                       &militaryServiceNavyW4ChiefWarrantOfficer,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumMilitaryService) ByID(id MilitaryServiceIdentifier) *EnumMilitaryServiceItem {
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
func (e *EnumMilitaryService) ByIDString(idx string) *EnumMilitaryServiceItem {
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
func (e *EnumMilitaryService) ByIndex(idx int) *EnumMilitaryServiceItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedMilitaryServiceID is a struct that is designed to replace a *MilitaryServiceID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *MilitaryServiceID it contains while being a better JSON citizen.
type ValidatedMilitaryServiceID struct {
	// id will point to a valid MilitaryServiceID, if possible
	// If id is nil, then ValidatedMilitaryServiceID.Valid() will return false.
	id *MilitaryServiceID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedMilitaryServiceID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedMilitaryServiceID
func (vi *ValidatedMilitaryServiceID) Clone() *ValidatedMilitaryServiceID {
	if vi == nil {
		return nil
	}

	var cid *MilitaryServiceID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedMilitaryServiceID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedMilitaryServiceIds represent the same MilitaryService
func (vi *ValidatedMilitaryServiceID) Equals(vj *ValidatedMilitaryServiceID) bool {
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

// Valid returns true if and only if the ValidatedMilitaryServiceID corresponds to a recognized MilitaryService
func (vi *ValidatedMilitaryServiceID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedMilitaryServiceID) ID() *MilitaryServiceID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedMilitaryServiceID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedMilitaryServiceID) ValidatedID() *ValidatedMilitaryServiceID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedMilitaryServiceID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedMilitaryServiceID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedMilitaryServiceID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedMilitaryServiceID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedMilitaryServiceID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := MilitaryServiceID(capString)
	item := MilitaryService.ByID(&id)
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

func (vi ValidatedMilitaryServiceID) String() string {
	return vi.ToIDString()
}

type MilitaryServiceIdentifier interface {
	ID() *MilitaryServiceID
	Valid() bool
}
