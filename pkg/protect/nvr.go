package protect

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

type NVR struct {
	connected bool
	host      string
	port      int
	user      string
	password  string
	csrfToken string
	cookies   string
}

func NewNVR(host string, port int, user string, password string) *NVR {
	unifiProtectWebsocket := &NVR{
		connected: false,
		host:      host,
		port:      port,
		user:      user,
		password:  password,
	}

	return unifiProtectWebsocket
}

func (n *NVR) Authenticate() error {
	n.connected = false

	// Make a first call to retrive a csrf token
	if err := n.Call(http.MethodGet, "/", nil, nil); err != nil {
		return err
	}

	r := strings.NewReader(fmt.Sprintf(`{"password": "%s", "username": "%s"}`, n.password, n.user))

	if err := n.Call(http.MethodPost, "/api/auth/login", r, nil); err != nil {
		return err
	}

	n.connected = true

	return nil
}

func (n *NVR) GetSocketEvents() (*WebsocketEvent, error) {
	return NewWebsocketEvent(n)
}

func (n *NVR) GetLiveFeed(camera string, channel int) *LiveFeed {
	return NewLiveFeed(n, camera, channel)
}

func (n *NVR) Call(method string, url string, body io.Reader, out interface{}) error {
	var fullBody []byte
	if body != nil {
		fullBody, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}
		body = bytes.NewReader(fullBody)
	}

	if err := n.httpDo(method, url, body, out); err != nil {
		if err != ErrUnauthorized {
			return err
		}

		// Reconnect and retry
		if err := n.Authenticate(); err != nil {
			return err
		}
		// Re-create a body reader from the full body
		if body != nil {
			body = bytes.NewReader(fullBody)
		}
		return n.httpDo(method, url, body, out)
	}
	return nil
}

func (n *NVR) httpDo(method string, url string, body io.Reader, out interface{}) error {
	request, err := http.NewRequest(method, fmt.Sprintf("https://%s:%d%s", n.host, n.port, url), body)
	if err != nil {
		return err
	}

	request.Header.Set("Cookie", n.cookies)
	request.Header.Add("X-CSRF-Token", n.csrfToken)

	if body != nil {
		request.Header.Add("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if resp.StatusCode == 401 {
		return ErrUnauthorized
	}

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("invalid return code %d", resp.StatusCode)
	}

	n.csrfToken = resp.Header.Get("X-CSRF-Token")
	n.cookies = resp.Header.Get("Set-Cookie")

	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}

	return nil
}

func (n *NVR) GetBootstrap() (*Bootstrap, error) {
	bootstrap := &Bootstrap{}
	return bootstrap, n.Call(http.MethodGet, "/proxy/protect/api/bootstrap", nil, bootstrap)
}
