package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type PatchAccountRequest struct {
	Name    string `json:"name"`
	Amount  int    `json:"amount"`
	NewName string `json:"newname"`
}

type ChangeAccountRequest struct {
	Name    string `json:"name"`
	NewName string `json:"newname"`
	Amount  int    `json:"amount"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
