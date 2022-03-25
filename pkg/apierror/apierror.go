package apierror

import "k8s.io/apimachinery/pkg/util/json"

type ApiError struct {
	Msg      string  `json:"message"`
	ErrorMsg *string `json:"error"`
	Code     int     `json:"code"`
	Success  bool    `json:"success"`
}

func (a ApiError) Error() string {
	return a.toJson()
}

func (a ApiError) toJson() string {
	j, _ := json.Marshal(a)
	return string(j)
}

func GetError(code int, message error) error {
	err := ErrorDictionary[code]
	if message != nil {
		e := message.Error()
		err.ErrorMsg = &e
	}
	return err
}

const (
	K8sInitError = 1000 + iota
	BootstrapError
	K8sClientError
	OauthClientError
	PersistenceInitError
	LabelNotSupportedError
)

const (
	K8sInitErrorMsg         = "Error reading kubernetes configuration"
	K8sClientErrorMsg       = "Error Creating a kubernetes clientset"
	OauthClientErrorMsg     = "Error Creating a oauth client"
	PersistenceInitErrorMsg = "Error creating persistence client"
	BootstraperrorMsg       = "Error in bootstrapping service"
	LabelNotSupportedMsg    = "Label not supported"
)

var ErrorDictionary = map[int]ApiError{
	K8sInitError: {
		Msg:     K8sInitErrorMsg,
		Code:    K8sInitError,
		Success: false,
	},
	K8sClientError: {
		Msg:     K8sClientErrorMsg,
		Code:    K8sClientError,
		Success: false,
	},
	OauthClientError: {
		Msg:     OauthClientErrorMsg,
		Code:    OauthClientError,
		Success: false,
	},
	PersistenceInitError: {
		Msg:     PersistenceInitErrorMsg,
		Code:    PersistenceInitError,
		Success: false,
	},
	BootstrapError: {
		Msg:     BootstraperrorMsg,
		Code:    BootstrapError,
		Success: false,
	},
	LabelNotSupportedError: {
		Msg:     LabelNotSupportedMsg,
		Code:    LabelNotSupportedError,
		Success: false,
	},
}
