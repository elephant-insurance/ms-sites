package mbuf

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
	"github.com/elephant-insurance/go-microservice-arch/v2/uf"
	"github.com/goccy/go-json"
)

func newAzureServiceBusMarshaler(resourceName, queueOrTopic, securityKeyName, securityKey string, ttl *time.Duration) (HTTPMarshaler, error) {
	var (
		td               = ttl
		sburl     string = fmt.Sprintf(AzureServiceBusURIPattern, resourceName, queueOrTopic)
		tf, tferr        = uf.NewAzureServiceBusTokenFactory(securityKey, securityKeyName, resourceName, queueOrTopic)
	)

	if tferr != nil {
		return nil, tferr
	}

	if td == nil {
		td = &defaultAzureServiceBusTTL
	}

	return &azureServiceBusMarshaler{
		failuresToMarshal: &clicker.Clicker{},
		queueOrTopic:      queueOrTopic,
		resourceName:      resourceName,
		securityKeyName:   securityKeyName,
		serviceBusURL:     sburl,
		tokenFactory:      tf,
		ttlDuration:       td,
	}, nil
}

type azureServiceBusMarshaler struct {
	failuresToMarshal *clicker.Clicker
	queueOrTopic      string
	resourceName      string
	securityKeyName   string
	serviceBusURL     string
	tokenFactory      uf.TokenFactory
	ttlDuration       *time.Duration
}

// Marshal sets up an http request containing messages to send to Azure Service Bus
// the actual structure expected is [][]byte, though [][]string will also work
// builds the array into a request body formatted for Azure
// returns an error if any item in the list cannot be cast to []byte
// Then we create the Azure SAS token auth header and return the entire request
func (asbm *azureServiceBusMarshaler) Marshal(entries []interface{}) (*http.Request, error) {
	body, err := buildAzureMessageBusArrayBody(entries)
	if err != nil {
		asbm.failuresToMarshal.Click(1)
		return nil, err
	}

	if len(body) == 0 {
		// no point sending an empty meessage
		return nil, nil
	}

	data := string(body)
	expiry := time.Now().Add(*asbm.ttlDuration)
	sasHeader, herr := asbm.tokenFactory.GetSASToken(&expiry)
	if herr != nil {
		return nil, herr
	}
	debugLog(sasHeader)

	rq, err := http.NewRequest(http.MethodPost, asbm.serviceBusURL, bytes.NewReader([]byte(data)))
	if err != nil {
		msg := errMsgFailedToCreateRequest
		return nil, errors.New(msg)
	}

	// rq.Header.Add automatically uppercases the key
	rq.Header.Add("Authorization", sasHeader)
	rq.Header.Add("Content-Type", azureServiceBusContentType)

	return rq, nil
}

func (asbm *azureServiceBusMarshaler) Diagnostics() map[string]interface{} {
	rtn := map[string]interface{}{
		diagnosticsFieldFailuresToMarshal: asbm.failuresToMarshal.Clicks,
		diagnosticsFieldAzureQueueOrTopic: asbm.queueOrTopic,
		diagnosticsFieldAzureResourceName: asbm.resourceName,
	}

	return rtn
}

// AzureServiceBusSettings is for use in configuration structs and their associated config.yml files
// Some required fields are optional here so that they may be specified globally for several different mbufs.
type AzureServiceBusSettings struct {
	MBUFSettings    *Settings `yaml:"MBUFSettings" config:"optional"`
	QueueOrTopic    string    `yaml:"QueueOrTopic" config:"public"`
	ResourceName    string    `yaml:"ResourceName" config:"public,optional"`
	SecurityKey     string    `yaml:"SecurityKey" config:"optional"`
	SecurityKeyName string    `yaml:"SecurityKeyName" config:"optional"`
	TTLSeconds      *int      `yaml:"TTLSeconds" config:"optional"`
}

type AzureServiceBusMessage struct {
	Body             string                 `json:"Body"`
	BrokerProperties map[string]interface{} `json:"BrokerProperties,omitempty"`
	UserProperties   map[string]interface{} `json:"UserProperties,omitempty"`
}

type BrokerProperties struct {
	Label *string
}

func buildAzureMessageBusArrayBody(entries []interface{}) ([]byte, error) {
	// This function writes the body of each message into a small struct so that it can be
	// properly separated on the other end
	// Currently it writes each body as a struct field string, then JSON marshals the struct
	// This produces an escaped-JSON string body field
	var msgAry = make([]AzureServiceBusMessage, len(entries))
	for i := 0; i < len(msgAry); i++ {
		thisEntry, ok := entries[i].([]byte)
		if !ok {
			return nil, errors.New(`unable to cast entry to byte array`)
		}

		thisEntryString := string(thisEntry)
		thisMsg := AzureServiceBusMessage{
			Body: thisEntryString,
		}
		msgAry[i] = thisMsg
	}

	rtn, merr := json.Marshal(msgAry)
	if merr != nil {
		return nil, merr
	}

	return rtn, nil
}

var (
	defaultAzureServiceBusTTL time.Duration = time.Hour * 24 * 7
)
