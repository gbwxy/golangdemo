package my

type MyError interface {
	error
	Message() string
}
