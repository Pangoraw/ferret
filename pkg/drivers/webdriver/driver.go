package webdriver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MontFerret/ferret/pkg/drivers"
)

var UnimplementedError = errors.New("unimplemented")

type WebDriver struct {
	Client *Client
	SessionID string
	Connected bool
}

func New() *WebDriver {
	return &WebDriver{
		Client: NewClient("", ""),
		Connected: false,
		SessionID: "",
	}
}

func (wd *WebDriver) Name() string {
	return "webdriver.firefox"
}

func (wd *WebDriver) Open(ctx context.Context, params drivers.Params) (drivers.HTMLPage, error) {
	if wd.Connected || wd.SessionID != "" {
		return nil, errors.New("session already initialized")
	}

	payload := `
	{
		"capabilities": {"alwaysMatch": {"browserName": "firefox"}}
	}`
	type sessionResponseValue struct {
		sessionId string
	}
	type sessionResponse struct {
		value sessionResponseValue
	}

	resp, err := wd.Client.Post("session", []byte(payload))
	if err != nil {
		return nil, err
	}
	var session sessionResponse
	err = json.Unmarshal(resp, &session)
	if err != nil {
		fmt.Println("Parse error")
		return nil, err
	}

	wd.SessionID = session.value.sessionId

	fmt.Printf("resp: %s", string(resp))
	fmt.Printf("Connected to session %q\n", wd.SessionID)

	wd.Connected = true

	return nil, nil
}

func (wd *WebDriver) Parse(ctx context.Context, params drivers.ParseParams) (drivers.HTMLPage, error) {
	return nil, UnimplementedError
}

func (wd *WebDriver) Close() error {
	if !wd.Connected || wd.SessionID == "" {
		return nil
	}

	endpoint := fmt.Sprintf("session/%s", wd.SessionID)
	_, err := wd.Client.Delete(endpoint)
	if err != nil {
		return err
	}

	wd.Connected = false

	fmt.Println("Session closed")

	return nil
}



func (wd *WebDriver) Get(endpoint string) {}

func (wd *WebDriver) NewSession() (*Session, error) {
	return nil, UnimplementedError
}
