package data

type ErrorResponse struct {
	Error string
	Code  int
}

type StatusResponse struct {
	Result bool
}

type ResponseWithID struct {
	Id interface{}
}
