package mbuf

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/elephant-insurance/go-microservice-arch/v2/clicker"
)

type azureLogMarshaler struct {
	authBase          string
	failuresToMarshal *clicker.Clicker
	logType           string
	relayURL          string
	sharedKey         string
	timeStampField    string
	workspaceID       string
}

func newAzureLogMarshaler(workspaceID, sharedKey, logType, timeStampField string) HTTPMarshaler {
	return &azureLogMarshaler{
		failuresToMarshal: &clicker.Clicker{},
		logType:           logType,
		relayURL:          fmt.Sprintf(azureURLPattern, workspaceID),
		sharedKey:         sharedKey,
		timeStampField:    timeStampField,
		workspaceID:       workspaceID,
	}
}

// Marshal sets up an http request containing log messages to send to Azure
// the actual structure expected is [][]byte, though [][]string will also work
// builds the array into a request body formatted for Azure
// returns an error if any item in the list cannot be cast to []byte
// Then we create the baroque Azure security header and return the entire request
func (af *azureLogMarshaler) Marshal(entries []interface{}) (*http.Request, error) {
	body, err := buildJSONArrayBody(entries)
	if err != nil {
		af.failuresToMarshal.Click(1)
		return nil, err
	}

	if len(body) == 0 {
		// no point sending an empty meessage
		return nil, nil
	}

	data := string(body)

	metaDataString, dateString := buildAzureLogAnalyticsMetaDataString(data)

	hashedString, err := af.buildAzureLogAnalyticsSignature(metaDataString)
	debugLog("Hashed string: " + hashedString)

	if err != nil {
		// we've already logged it and the calling thread is long gone
		return nil, errors.New("error building Azure signature")
	}

	if af.authBase == "" {
		af.authBase = fmt.Sprintf(azureAuthBasePattern, af.workspaceID)
	}
	sig := af.authBase + hashedString
	debugLog("Signature: " + sig)

	rq, err := http.NewRequest(http.MethodPost, af.relayURL, bytes.NewReader([]byte(data)))
	if err != nil {
		msg := errMsgFailedToCreateRequest
		return nil, errors.New(msg)
	}

	// rq.Header.Add automatically uppercases the key
	rq.Header.Add("Log-Type", af.logType)
	debugLog("Azure Log Type (table name): " + af.logType)
	rq.Header.Add("Authorization", sig)
	rq.Header.Add("Content-Type", "application/json")
	// assign directly to the Header map to retain lowercase key name
	rq.Header["x-ms-date"] = []string{dateString}
	rq.Header["time-generated-field"] = []string{af.timeStampField}
	rq.Header.Add("Accept", "application/json")

	return rq, nil
}

func (af *azureLogMarshaler) buildAzureLogAnalyticsSignature(msg string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(af.sharedKey)
	if err != nil {
		ermsg := "failed to decode RelayEndpointAzureSharedKey: "
		debugLog(ermsg + err.Error())
		return "", err
	}

	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(msg))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

func buildAzureLogAnalyticsMetaDataString(body string) (metadata, date string) {
	// build the Azure header, which is complicated
	bodyLength := strconv.Itoa(utf8.RuneCountInString(body))
	debugLog(fmt.Sprintf("DATA (%v bytes): \n%v", bodyLength, body))
	dateString := time.Now().UTC().Format(time.RFC1123)

	// Azure doesn't like "UTC"
	dateString = strings.Replace(dateString, "UTC", "GMT", -1)

	metaDataString := "POST\n" + bodyLength + "\napplication/json\n" + "x-ms-date:" + dateString + "\n/api/logs"
	debugLog("Azure metadata string: " + metaDataString)

	return metaDataString, dateString
}

func (af *azureLogMarshaler) Diagnostics() map[string]interface{} {
	rtn := map[string]interface{}{
		diagnosticsFieldFailuresToMarshal: af.failuresToMarshal.Clicks,
		diagnosticsFieldAzureWorkspaceID:  af.workspaceID,
		diagnosticsFieldLogType:           af.logType,
	}

	return rtn
}
