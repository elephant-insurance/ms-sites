package resp

// ResponseMessageDto model to hold Error Messages/Codes and only response format leaves micro service (other than crash)
type ResponseMessageDto struct {
	ValidationMessages []*MessageDetailDto `json:"validationMessages"`
	ErrorMessages      []*MessageDetailDto `json:"errorMessages"`
	HasError           bool                `json:"hasError"`
	HasValidation      bool                `json:"hasValidation"`
	Data               interface{}         `json:"data"`
}

// MessageDetailDto model to have error code and message
type MessageDetailDto struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// LoadDefaults => Load default values in case they are missing
func (responseMessage *ResponseMessageDto) LoadDefaults() {
	if responseMessage.ValidationMessages == nil {
		responseMessage.ValidationMessages = []*MessageDetailDto{}
	}
	if responseMessage.ErrorMessages == nil {
		responseMessage.ErrorMessages = []*MessageDetailDto{}
	}
	responseMessage.HasError = responseMessage.HasErrors()
	responseMessage.HasValidation = responseMessage.HasValidations()
}

// AddErrorDto => Adds the error to ResponseMessage Object
func (responseMessage *ResponseMessageDto) AddErrorDto(messageDto *MessageDetailDto) {
	if messageDto != nil {
		responseMessage.ErrorMessages = append(responseMessage.ErrorMessages, messageDto)
	}
}

// AddError => Adds the error to ResponseMessage Object
func (responseMessage *ResponseMessageDto) AddError(code, message string) {
	responseMessage.AddErrorDto(&MessageDetailDto{Code: code, Message: message})
}

// AddValidationMessageDto => Add on validation message to ResponseMessage Object
func (responseMessage *ResponseMessageDto) AddValidationMessageDto(messageDto *MessageDetailDto) {
	if messageDto != nil {
		responseMessage.ValidationMessages = append(responseMessage.ValidationMessages, messageDto)
	}
}

// AddValidationMessage => Add on validation message to ResponseMessage Object
func (responseMessage *ResponseMessageDto) AddValidationMessage(code, message string) {
	responseMessage.AddValidationMessageDto(&MessageDetailDto{Code: code, Message: message})
}

// SetPayLoad => sets the payLoad for the response object
func (responseMessage *ResponseMessageDto) SetPayLoad(payLoad interface{}) {
	responseMessage.Data = payLoad
}

// IsValid => tells that is there any validation or system error
func (responseMessage *ResponseMessageDto) IsValid() bool {
	return responseMessage == nil || (len(responseMessage.ValidationMessages) == 0 && len(responseMessage.ErrorMessages) == 0)
}

// HasValidations => tells that is there any validation or system error
func (responseMessage *ResponseMessageDto) HasValidations() bool {
	return responseMessage != nil && (len(responseMessage.ValidationMessages) > 0)
}

// HasErrors => tells that is there any system error
func (responseMessage *ResponseMessageDto) HasErrors() bool {
	return responseMessage != nil && len(responseMessage.ErrorMessages) > 0
}

// AddClassifiedResponseMessages => distinguish between type of Error [either System Error or Validation Error]
func (responseMessage *ResponseMessageDto) AddClassifiedResponseMessages(code string, message string) {
	switch code {
	default:
		if code == "" {
			code = "error"
		}
		responseMessage.AddError(code, message)
	}
}
