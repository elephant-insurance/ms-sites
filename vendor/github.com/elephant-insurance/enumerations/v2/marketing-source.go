package enumerations

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

// MarketingSourceID uniquely identifies a particular MarketingSource
type MarketingSourceID string

// Clone creates a safe, independent copy of a MarketingSourceID
func (i *MarketingSourceID) Clone() *MarketingSourceID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two MarketingSourceIds are equivalent
func (i *MarketingSourceID) Equals(j *MarketingSourceID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *MarketingSourceID that is either valid or nil
func (i *MarketingSourceID) ID() *MarketingSourceID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *MarketingSourceID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the MarketingSourceID corresponds to a recognized MarketingSource
func (i *MarketingSourceID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return MarketingSource.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *MarketingSourceID) ValidatedID() *ValidatedMarketingSourceID {
	if i != nil {
		return &ValidatedMarketingSourceID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *MarketingSourceID) MarshalJSON() ([]byte, error) {
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

func (i *MarketingSourceID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := MarketingSourceID(dataString)
	item := MarketingSource.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	marketingSourceBillboardID      MarketingSourceID = "billboard"
	marketingSourceEmailID          MarketingSourceID = "email"
	marketingSourceInternetAdID     MarketingSourceID = "internetad"
	marketingSourceRadioID          MarketingSourceID = "radio"
	marketingSourceReferralID       MarketingSourceID = "referral"
	marketingSourceRepeatCustomerID MarketingSourceID = "repeatcustomer"
	marketingSourceSearchEngineID   MarketingSourceID = "searchengine"
	marketingSourceSocialMediaID    MarketingSourceID = "socialmedia"
	marketingSourceSportingEventID  MarketingSourceID = "sportingevent"
	marketingSourceSpotifyID        MarketingSourceID = "spotify"
	marketingSourceTVID             MarketingSourceID = "tv"
	marketingSourceYouTubeID        MarketingSourceID = "youtube"
	marketingSourceOtherID          MarketingSourceID = "other"
)

// EnumMarketingSourceItem describes an entry in an enumeration of MarketingSource
type EnumMarketingSourceItem struct {
	ID        MarketingSourceID `json:"Value"`
	Desc      string            `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name      string            `json:"Name"`
	SortOrder int
}

var (
	marketingSourceBillboard      = EnumMarketingSourceItem{marketingSourceBillboardID, "Billboard", nil, "Billboard", 1}
	marketingSourceEmail          = EnumMarketingSourceItem{marketingSourceEmailID, "Email", nil, "Email", 2}
	marketingSourceInternetAd     = EnumMarketingSourceItem{marketingSourceInternetAdID, "Internet Ad", nil, "InternetAd", 3}
	marketingSourceRadio          = EnumMarketingSourceItem{marketingSourceRadioID, "Radio", nil, "Radio", 4}
	marketingSourceReferral       = EnumMarketingSourceItem{marketingSourceReferralID, "Referral", nil, "Referral", 5}
	marketingSourceRepeatCustomer = EnumMarketingSourceItem{marketingSourceRepeatCustomerID, "Repeat Customer", nil, "RepeatCustomer", 6}
	marketingSourceSearchEngine   = EnumMarketingSourceItem{marketingSourceSearchEngineID, "Search Engine", nil, "SearchEngine", 7}
	marketingSourceSocialMedia    = EnumMarketingSourceItem{marketingSourceSocialMediaID, "Social Media", nil, "SocialMedia", 8}
	marketingSourceSportingEvent  = EnumMarketingSourceItem{marketingSourceSportingEventID, "Sporting Event", nil, "SportingEvent", 9}
	marketingSourceSpotify        = EnumMarketingSourceItem{marketingSourceSpotifyID, "Spotify", nil, "Spotify", 10}
	marketingSourceTV             = EnumMarketingSourceItem{marketingSourceTVID, "TV", nil, "TV", 11}
	marketingSourceYouTube        = EnumMarketingSourceItem{marketingSourceYouTubeID, "YouTube", nil, "YouTube", 12}
	marketingSourceOther          = EnumMarketingSourceItem{marketingSourceOtherID, "Other", nil, "Other", 13}
)

// EnumMarketingSource is a collection of MarketingSource items
type EnumMarketingSource struct {
	Description string
	Items       []*EnumMarketingSourceItem
	Name        string

	Billboard      *EnumMarketingSourceItem
	Email          *EnumMarketingSourceItem
	InternetAd     *EnumMarketingSourceItem
	Radio          *EnumMarketingSourceItem
	Referral       *EnumMarketingSourceItem
	RepeatCustomer *EnumMarketingSourceItem
	SearchEngine   *EnumMarketingSourceItem
	SocialMedia    *EnumMarketingSourceItem
	SportingEvent  *EnumMarketingSourceItem
	Spotify        *EnumMarketingSourceItem
	TV             *EnumMarketingSourceItem
	YouTube        *EnumMarketingSourceItem
	Other          *EnumMarketingSourceItem

	itemDict map[string]*EnumMarketingSourceItem
}

// MarketingSource is a public singleton instance of EnumMarketingSource
// representing places where customers learned about Elephant
var MarketingSource = &EnumMarketingSource{
	Description: "places where customers learned about Elephant",
	Items: []*EnumMarketingSourceItem{
		&marketingSourceBillboard,
		&marketingSourceEmail,
		&marketingSourceInternetAd,
		&marketingSourceRadio,
		&marketingSourceReferral,
		&marketingSourceRepeatCustomer,
		&marketingSourceSearchEngine,
		&marketingSourceSocialMedia,
		&marketingSourceSportingEvent,
		&marketingSourceSpotify,
		&marketingSourceTV,
		&marketingSourceYouTube,
		&marketingSourceOther,
	},
	Name:           "EnumMarketingSource",
	Billboard:      &marketingSourceBillboard,
	Email:          &marketingSourceEmail,
	InternetAd:     &marketingSourceInternetAd,
	Radio:          &marketingSourceRadio,
	Referral:       &marketingSourceReferral,
	RepeatCustomer: &marketingSourceRepeatCustomer,
	SearchEngine:   &marketingSourceSearchEngine,
	SocialMedia:    &marketingSourceSocialMedia,
	SportingEvent:  &marketingSourceSportingEvent,
	Spotify:        &marketingSourceSpotify,
	TV:             &marketingSourceTV,
	YouTube:        &marketingSourceYouTube,
	Other:          &marketingSourceOther,

	itemDict: map[string]*EnumMarketingSourceItem{
		strings.ToLower(string(marketingSourceBillboardID)):      &marketingSourceBillboard,
		strings.ToLower(string(marketingSourceEmailID)):          &marketingSourceEmail,
		strings.ToLower(string(marketingSourceInternetAdID)):     &marketingSourceInternetAd,
		strings.ToLower(string(marketingSourceRadioID)):          &marketingSourceRadio,
		strings.ToLower(string(marketingSourceReferralID)):       &marketingSourceReferral,
		strings.ToLower(string(marketingSourceRepeatCustomerID)): &marketingSourceRepeatCustomer,
		strings.ToLower(string(marketingSourceSearchEngineID)):   &marketingSourceSearchEngine,
		strings.ToLower(string(marketingSourceSocialMediaID)):    &marketingSourceSocialMedia,
		strings.ToLower(string(marketingSourceSportingEventID)):  &marketingSourceSportingEvent,
		strings.ToLower(string(marketingSourceSpotifyID)):        &marketingSourceSpotify,
		strings.ToLower(string(marketingSourceTVID)):             &marketingSourceTV,
		strings.ToLower(string(marketingSourceYouTubeID)):        &marketingSourceYouTube,
		strings.ToLower(string(marketingSourceOtherID)):          &marketingSourceOther,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumMarketingSource) ByID(id MarketingSourceIdentifier) *EnumMarketingSourceItem {
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
func (e *EnumMarketingSource) ByIDString(idx string) *EnumMarketingSourceItem {
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
func (e *EnumMarketingSource) ByIndex(idx int) *EnumMarketingSourceItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedMarketingSourceID is a struct that is designed to replace a *MarketingSourceID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *MarketingSourceID it contains while being a better JSON citizen.
type ValidatedMarketingSourceID struct {
	// id will point to a valid MarketingSourceID, if possible
	// If id is nil, then ValidatedMarketingSourceID.Valid() will return false.
	id *MarketingSourceID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedMarketingSourceID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedMarketingSourceID
func (vi *ValidatedMarketingSourceID) Clone() *ValidatedMarketingSourceID {
	if vi == nil {
		return nil
	}

	var cid *MarketingSourceID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedMarketingSourceID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedMarketingSourceIds represent the same MarketingSource
func (vi *ValidatedMarketingSourceID) Equals(vj *ValidatedMarketingSourceID) bool {
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

// Valid returns true if and only if the ValidatedMarketingSourceID corresponds to a recognized MarketingSource
func (vi *ValidatedMarketingSourceID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedMarketingSourceID) ID() *MarketingSourceID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedMarketingSourceID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedMarketingSourceID) ValidatedID() *ValidatedMarketingSourceID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedMarketingSourceID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedMarketingSourceID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}

// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedMarketingSourceID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedMarketingSourceID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedMarketingSourceID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := MarketingSourceID(capString)
	item := MarketingSource.ByID(&id)
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

func (vi ValidatedMarketingSourceID) String() string {
	return vi.ToIDString()
}

type MarketingSourceIdentifier interface {
	ID() *MarketingSourceID
	Valid() bool
}
