package handlers

type myHandler struct {
	hashedPasswords map[string]string
}

func NewMyHandler(psswds map[string]string) *myHandler {
	return &myHandler{hashedPasswords: psswds}
}