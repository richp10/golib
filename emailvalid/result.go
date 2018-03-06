// From https://github.com/smancke/mailck
// Copyright (c) 2017 Sebastian Mancke - MIT License
// Further developed by Richard Phillips - MIT License

package emailvalid

type ResultState string

const (
	ValidState   ResultState = "valid"
	InvalidState ResultState = "invalid"
	ErrorState   ResultState = "error"
)

func (rs ResultState) String() string {
	return string(rs)
}

// Result contains the information about an email check.
type Result struct {
	Result       ResultState `json:"result"`
	ResultDetail string      `json:"resultDetail"`
	Message      string      `json:"message"`
}

var (
	Valid              = Result{ValidState, "mailboxChecked", "The email address is valid."}
	InvalidSyntax      = Result{InvalidState, "invalidSyntax", "The email format is invalid."}
	InvalidDomain      = Result{InvalidState, "invalidDomain", "The email domain does not exist."}
	MailboxUnavailable = Result{InvalidState, "mailboxUnavailable", "The email username does not exist."}
	Disposable         = Result{InvalidState, "disposable", "You cannot use disposable email addresses."}
	MailserverError    = Result{ErrorState, "mailserverError", "The target mailserver responded with an error."}
	TimeoutError       = Result{ErrorState, "timeoutError", "The connection with the mailserver timed out."}
	NetworkError       = Result{ErrorState, "networkError", "The connection to the mailserver could not be made."}
	ServiceError       = Result{ErrorState, "serviceError", "An internal error occured while checking."}
	ClientError        = Result{ErrorState, "clientError", "The request was was invalid."}
)

func (r Result) IsValid() bool {
	return r.Result == ValidState
}

func (r Result) IsInvalid() bool {
	return r.Result == InvalidState
}

func (r Result) IsError() bool {
	return r.Result == ErrorState
}
