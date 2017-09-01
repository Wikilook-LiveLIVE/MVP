package errors

var (
	//not available DB, method closed
	ERR_METHOD_UNAVAILABLE = &AppErr{0, "METHOD_UNAVAILABLE"}
	//invalid json, no mandatory fields
	ERR_INVALID_REQUEST = &AppErr{1, "INVALID_REQUEST"}
	//invalid or expired
	ERR_SECURITY_TOKEN_INVALID = &AppErr{2, "SECURITY_TOKEN_INVALID"}
	//customer not found
	ERR_CUSTOMER_NOT_FOUND = &AppErr{3, "CUSTOMER_NOT_FOUND"}

)

type AppErr struct {
	Id int
	Code  string
}

type Id interface {
	Id() int
}

type Code interface {
	Code() string
}

func (e *AppErr) GetId() int {
	return e.Id
}

func (e *AppErr) GetCode() string {
	return e.Code
}