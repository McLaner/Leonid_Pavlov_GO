package main

import (
	"awesomeProject/accounts/dto"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type AccountOperation struct {
	ServerPort     int
	ServerHost     string
	Operation      string
	AccountName    string
	AccountBalance int
	NewAccountName string
}

func (ao *AccountOperation) ExecuteOperation() error {
	switch strings.ToLower(ao.Operation) {
	case "create":
		return ao.createAccount()
	case "get":
		return ao.getAccount()
	case "delete":
		return ao.deleteAccount()
	case "patch":
		return ao.patchAccount()
	case "change":
		return ao.changeAccountName()
	default:
		return fmt.Errorf("unsupported operation: %s", ao.Operation)
	}
}

func (ao *AccountOperation) createAccount() error {
	reqBody := dto.CreateAccountRequest{
		Name:   ao.AccountName,
		Amount: ao.AccountBalance,
	}
	return ao.sendRequest("POST", "/account/create", reqBody)
}

func (ao *AccountOperation) getAccount() error {
	url := fmt.Sprintf("http://%s:%d/account?name=%s", ao.ServerHost, ao.ServerPort, ao.AccountName)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get account: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var accountInfo dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&accountInfo); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	fmt.Printf("Account: %s, Balance: %d\n", accountInfo.Name, accountInfo.Amount)
	return nil
}

func (ao *AccountOperation) deleteAccount() error {
	reqBody := dto.DeleteAccountRequest{Name: ao.AccountName}
	return ao.sendRequest("DELETE", "/account/delete", reqBody)
}

func (ao *AccountOperation) patchAccount() error {
	reqBody := dto.PatchAccountRequest{
		Name:   ao.AccountName,
		Amount: ao.AccountBalance,
	}
	return ao.sendRequest("PATCH", "/account/patch", reqBody)
}

func (ao *AccountOperation) changeAccountName() error {
	reqBody := dto.ChangeAccountRequest{
		Name:    ao.AccountName,
		NewName: ao.NewAccountName,
	}
	return ao.sendRequest("PUT", "/account/change", reqBody)
}

func (ao *AccountOperation) sendRequest(method, path string, body interface{}) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	url := fmt.Sprintf("http://%s:%d%s", ao.ServerHost, ao.ServerPort, path)
	req, err := http.NewRequest(method, url, strings.NewReader(string(jsonData)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

func main() {
	op := &AccountOperation{}

	flag.IntVar(&op.ServerPort, "port", 1323, "server port")
	flag.StringVar(&op.ServerHost, "host", "0.0.0.0", "server host")
	flag.StringVar(&op.Operation, "op", "", "operation to perform")
	flag.StringVar(&op.AccountName, "name", "", "account name")
	flag.IntVar(&op.AccountBalance, "balance", 0, "account balance")
	flag.StringVar(&op.NewAccountName, "newname", "", "new account name")

	flag.Parse()

	if err := op.ExecuteOperation(); err != nil {
		log.Fatalf("Operation failed: %v", err)
	}
}
