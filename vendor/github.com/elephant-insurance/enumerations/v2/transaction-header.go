package enumerations

	import (
		"encoding/json"
		"encoding/xml"
		"errors"
		"strings"
	)

// TXHeaderID uniquely identifies a particular TXHeader
type TXHeaderID string

// Clone creates a safe, independent copy of a TXHeaderID
func (i *TXHeaderID) Clone() *TXHeaderID {
	if i == nil {
		return nil
	}

	rtn := *i

	return &rtn
}

// Equals returns true if and only if two TXHeaderIds are equivalent
func (i *TXHeaderID) Equals(j *TXHeaderID) bool {
	if i == nil && j == nil {
		return true
	}

	if i == nil || j == nil {
		return false
	}

	return *i == *j
}

// ID returns this identifier as a *TXHeaderID that is either valid or nil 
func (i *TXHeaderID) ID() *TXHeaderID {
	if i != nil && i.Valid() {
		return i.Clone()
	}

	return nil
}

// ToIDString returns a string representation of this identifier, or empty string if it is invalid
func (i *TXHeaderID) ToIDString() string {
	if i != nil && *i != "" && i.Valid() {
		return string(*i)
	}

	return ""
}

// Valid returns true if and only if the TXHeaderID corresponds to a recognized TXHeader
func (i *TXHeaderID) Valid() bool {
	if i == nil || *i == "" {
		return false
	}

	return TXHeader.ByIDString(string(*i)) != nil
}

// ToValidated returns this identifier as a Validated type that is safe to unmarshal
func (i *TXHeaderID) ValidatedID() *ValidatedTXHeaderID {
	if i != nil {
		return &ValidatedTXHeaderID{
			id: i.ID(),
		}
	}

	return nil
}

func (i *TXHeaderID) MarshalJSON() ([]byte, error) {
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

func (i *TXHeaderID) UnmarshalJSON(data []byte) error {
	var dataString string
	err := json.Unmarshal(data, &dataString)
	if err != nil {
		return err
	}

	if dataString == "" {
		i = nil
		return nil
	}

	id := TXHeaderID(dataString)
	item := TXHeader.ByID(&id)
	if item == nil {
		err := errors.New(errorUnmarshalInvalidID)
		return err
	}

	*i = item.ID
	return nil
}

const (
	tXHeaderBrandID TXHeaderID = "txbrand"
	tXHeaderDomainID TXHeaderID = "txdomain"
	tXHeaderIDID TXHeaderID = "txid"
	tXHeaderIntegratorID TXHeaderID = "txintegrator"
	tXHeaderSourceID TXHeaderID = "txsource"
	tXHeaderTypeID TXHeaderID = "txtype"
	tXHeaderIPAddressID TXHeaderID = "txip"
	tXHeaderInstanceID TXHeaderID = "txinstance"
)

// EnumTXHeaderItem describes an entry in an enumeration of TXHeader
type EnumTXHeaderItem struct {
	ID        TXHeaderID `json:"Value"`
	Desc      string `json:"Description,omitempty"`
	Meta      map[string]string `json:",omitempty"`
	Name	  string `json:"Name"`
	SortOrder int

	// Meta Properties
	HeaderKey	 string
}

var (
	tXHeaderBrand = EnumTXHeaderItem{tXHeaderBrandID, "Transaction Brand (*BrandID)", map[string]string{"HeaderKey": "Admiral-Txbrand"}, "Brand", 1, "Admiral-Txbrand"}
	tXHeaderDomain = EnumTXHeaderItem{tXHeaderDomainID, "Transaction Domain (*AccountDomainID)", map[string]string{"HeaderKey": "Admiral-Txdomain"}, "Domain", 2, "Admiral-Txdomain"}
	tXHeaderID = EnumTXHeaderItem{tXHeaderIDID, "Transaction ID (string)", map[string]string{"HeaderKey": "Admiral-Txid"}, "ID", 3, "Admiral-Txid"}
	tXHeaderIntegrator = EnumTXHeaderItem{tXHeaderIntegratorID, "Transaction Integrator (*IntegrationPartnerID)", map[string]string{"HeaderKey": "Admiral-Txintegrator"}, "Integrator", 4, "Admiral-Txintegrator"}
	tXHeaderSource = EnumTXHeaderItem{tXHeaderSourceID, "Transaction Source (*SourceOfBusinessID)", map[string]string{"HeaderKey": "Admiral-Txsource"}, "Source", 5, "Admiral-Txsource"}
	tXHeaderType = EnumTXHeaderItem{tXHeaderTypeID, "Transaction Type (string)", map[string]string{"HeaderKey": "Admiral-Txtype"}, "Type", 6, "Admiral-Txtype"}
	tXHeaderIPAddress = EnumTXHeaderItem{tXHeaderIPAddressID, "Transaction IP Address (string)", map[string]string{"HeaderKey": "Admiral-Txip"}, "IPAddress", 7, "Admiral-Txip"}
	tXHeaderInstance = EnumTXHeaderItem{tXHeaderInstanceID, "Transaction Server Instance (string)", map[string]string{"HeaderKey": "Admiral-Txinstance"}, "Instance", 8, "Admiral-Txinstance"}
)

// EnumTXHeader is a collection of TXHeader items
type EnumTXHeader struct {
	Description string
	Items []*EnumTXHeaderItem
	Name  string
	
	Brand  *EnumTXHeaderItem
	Domain  *EnumTXHeaderItem
	ID  *EnumTXHeaderItem
	Integrator  *EnumTXHeaderItem
	Source  *EnumTXHeaderItem
	Type  *EnumTXHeaderItem
	IPAddress  *EnumTXHeaderItem
	Instance  *EnumTXHeaderItem

	itemDict map[string]*EnumTXHeaderItem
}

// TXHeader is a public singleton instance of EnumTXHeader
// representing Elephant microservice transaction headers
var TXHeader = &EnumTXHeader{
	Description: "Elephant microservice transaction headers",
	Items:  []*EnumTXHeaderItem{
		&tXHeaderBrand,
		&tXHeaderDomain,
		&tXHeaderID,
		&tXHeaderIntegrator,
		&tXHeaderSource,
		&tXHeaderType,
		&tXHeaderIPAddress,
		&tXHeaderInstance,
	},
	Name:   "EnumTXHeader",
	Brand: &tXHeaderBrand,
	Domain: &tXHeaderDomain,
	ID: &tXHeaderID,
	Integrator: &tXHeaderIntegrator,
	Source: &tXHeaderSource,
	Type: &tXHeaderType,
	IPAddress: &tXHeaderIPAddress,
	Instance: &tXHeaderInstance,

	itemDict: map[string]*EnumTXHeaderItem{
		strings.ToLower(string(tXHeaderBrandID)): &tXHeaderBrand,
		strings.ToLower(string(tXHeaderDomainID)): &tXHeaderDomain,
		strings.ToLower(string(tXHeaderIDID)): &tXHeaderID,
		strings.ToLower(string(tXHeaderIntegratorID)): &tXHeaderIntegrator,
		strings.ToLower(string(tXHeaderSourceID)): &tXHeaderSource,
		strings.ToLower(string(tXHeaderTypeID)): &tXHeaderType,
		strings.ToLower(string(tXHeaderIPAddressID)): &tXHeaderIPAddress,
		strings.ToLower(string(tXHeaderInstanceID)): &tXHeaderInstance,
	},
}

// ByID retrieves an entry by its native ID
func (e *EnumTXHeader) ByID(id TXHeaderIdentifier) *EnumTXHeaderItem {
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
func (e *EnumTXHeader) ByIDString(idx string) *EnumTXHeaderItem {
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
func (e *EnumTXHeader) ByIndex(idx int) *EnumTXHeaderItem {
	if e == nil || len(e.Items) < (idx+1) {
		return nil
	}

	return e.Items[idx]
}

// ValidatedTXHeaderID is a struct that is designed to replace a *TXHeaderID property in a JSON document
// where we need more control over marshaling and unmarshaling, including the ability to capture an
// invalid value during unmarshaling without returning an error.
// It is designed to behave exactly like the *TXHeaderID it contains while being a better JSON citizen.
type ValidatedTXHeaderID struct {
	// id will point to a valid TXHeaderID, if possible
	// If id is nil, then ValidatedTXHeaderID.Valid() will return false.
	id *TXHeaderID
	// capturedValue is the string representation of what we unmarshaled, if applicable
	capturedValue *string
	// Errors collects any errors we receved or generated while unmarshaling this
	Errors []error
}

// CapturedValue returns the raw string value of this field when it was unmarshaled from JSON, if applicable
func (vi *ValidatedTXHeaderID) CapturedValue() *string {
	if vi != nil {
		return vi.capturedValue
	}

	return nil
}

// Clone creates a safe, independent copy of a ValidatedTXHeaderID
func (vi *ValidatedTXHeaderID) Clone() *ValidatedTXHeaderID {
	if vi == nil {
		return nil
	}

	var cid *TXHeaderID
	if vi.id != nil {
		cid = vi.id.Clone()
	}

	rtn := ValidatedTXHeaderID{
		id: cid,
	}

	return &rtn
}

// Equals returns true if and only if two ValidatedTXHeaderIds represent the same TXHeader
func (vi *ValidatedTXHeaderID) Equals(vj *ValidatedTXHeaderID) bool {
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

// Valid returns true if and only if the ValidatedTXHeaderID corresponds to a recognized TXHeader
func (vi *ValidatedTXHeaderID) Valid() bool {
	if vi == nil || vi.id == nil {
		return false
	}

	return vi.id.Valid()
}

func (vi *ValidatedTXHeaderID) ID() *TXHeaderID {
	if vi != nil && vi.id != nil {
		return vi.id.ID()
	}

	return nil
}

func (vi *ValidatedTXHeaderID) ToIDString() string {
	if vi != nil && vi.id != nil {
		return string(*vi.id)
	}

	return ""
}

func (vi *ValidatedTXHeaderID) ValidatedID() *ValidatedTXHeaderID {
	if vi != nil {
		return vi.Clone()
	}

	return nil
}

func (vi *ValidatedTXHeaderID) MarshalJSON() ([]byte, error) {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return []byte(nullString), nil
	}

	return vi.id.MarshalJSON()
}

func (vi *ValidatedTXHeaderID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const nullString string = "null"

	if vi == nil || vi.id == nil {
		return e.EncodeElement("", start)
	}

	return e.EncodeElement(string(*vi.id), start)
}


// UnmarshalXML is guaranteed not to return a error!
func (vi *ValidatedTXHeaderID) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var xmlData string
	if err := decoder.DecodeElement(&xmlData, &start); err != nil {
		return err
	}

	return vi.unmarshalValue(xmlData)
}

// UnmarshalJSON is guaranteed not to return a error!
func (vi *ValidatedTXHeaderID) UnmarshalJSON(data []byte) error {
	// shouldn't be possible but a promise is a promise and we promise not to error out:
	if vi == nil {
		return nil
	}

	// capture the raw string in case it doesn't parse nicely
	capString := string(data)
	return vi.unmarshalValue(capString)
}

func (vi *ValidatedTXHeaderID) unmarshalValue(capString string) error {
	capString = strings.Replace(capString, "\"", "", -1)
	vi.capturedValue = &capString

	// empty string is invalid, but not an error
	if capString == "" {
		return nil
	}

	id := TXHeaderID(capString)
	item := TXHeader.ByID(&id)
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


func (vi ValidatedTXHeaderID) String() string {
	return vi.ToIDString()
}

type TXHeaderIdentifier interface {
	ID() *TXHeaderID
	Valid() bool
}
