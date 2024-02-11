package api

var (
	RequestError        = "Could not parse the request"
	InternalServerError = "Internal server error"
	CredentialsError    = "Invalid Credentials"
	FetchingError       = "Could not fetch the data"
	QueryNotFoundError  = "Could not found the query parameter"
	EmailError          = "Please Provide the correct email"
	NameError           = "Please provide correct name"
	PasswordError       = "Password should be strong"
	UniqueError         = "ID must be unique"
	NoResourseFound     = "No Resourse Found"
)

var (
	Login  = "Login Successfully"
	Signup = "Signup successfully"

	Create      = "Created successfully!!!"
	Got         = "Got successfully!!!"
	Delete      = "Deleted successfully!!!"
	Update      = "Updated successfully!!!"
	CreateError = "Could not create.."
	GetError    = "Could not get the data"
	DeleteError = "Could not delete data"
	UpdateError = "Could not delete data"
)
