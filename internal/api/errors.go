package api

var (
	RequestError        = "could not parse the request"
	InternalServerError = "internal server error"
	CredentialsError    = "invalid credentials"
	FetchingError       = "could not fetch the data"
	QueryNotFoundError  = "could not found the query parameter"
	EmailError          = "please provide the correct email"
	NameError           = "please provide correct name"
	PasswordError       = "password should be strong"
	UniqueError         = "id must be unique"
	UniqueTransaction   = "transaction id must be unique"
	NoResourseFound     = "no resourse found"
	NotAuthorized       = "not quthorized"
)

var (
	Login  = "login successfully"
	Signup = "signup successfully"

	Create      = "created successfully!!!"
	Got         = "got successfully!!!"
	Delete      = "ddeleted successfully!!!"
	Update      = "updated successfully!!!"
	CreateError = "could not create.."
	GetError    = "could not get the data"
	DeleteError = "could not delete data"
	UpdateError = "could not update data"
)
