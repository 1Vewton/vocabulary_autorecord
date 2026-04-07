package error_interface

type Error struct {
	Reason  string
	Err     error
	IsError bool
}
