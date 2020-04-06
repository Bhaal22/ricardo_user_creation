package logic

import (
    "net/http"
    "bytes"
    "errors"
    "io/ioutil"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"

    "github.com/Bhaal22/ricardo_user_creation/utils/mocks"
)

func TestGoogleDNSInUS(t *testing.T) {
    country, error := Country("8.8.8.8")

    assert.NotNil(t, country)
    assert.Nil(t, error)
    assert.Equal(t, "US", country, "8.8.8.8 ip address should be located in US")
}

func TestSuccess(t *testing.T) {
    client := &mocks.MockHttpClient{}

    Client = client

    body := "FR"
    response := &http.Response{
        Status:        "200 OK",
        StatusCode:    200,
        Proto:         "HTTP/1.1",
        ProtoMajor:    1,
        ProtoMinor:    1,
        Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
        ContentLength: int64(len(body)),
        Request:       nil,
        Header:        make(http.Header, 0),
    }

    client.On("Do", mock.Anything).Return(response, nil)

    country, error := Country("1.1.1.1")

    assert.NotNil(t, country)
    assert.Nil(t, error)
    assert.Equal(t, "FR", country, "1.1.1.1 ip address should be located in FR")
}

func TestFailureOnCallingIpApiEndpoint(t *testing.T) {
    client := &mocks.MockHttpClient{}

    Client = client

    client.On("Do", mock.Anything).Return(&http.Response{}, errors.New("something bad happened !"))

    country, error := Country("1.1.1.1")

    assert.Empty(t, country)
    assert.NotNil(t, error)
}

func TestFailureOnResourceNotFound(t *testing.T) {
    client := &mocks.MockHttpClient{}

    Client = client

    body := "<h1>Not Found</h1><p>The requested resource was not found on this server.</p>"
    response := &http.Response{
        Status:        "404 Not Found",
        StatusCode:    404,
        Proto:         "HTTP/1.1",
        ProtoMajor:    1,
        ProtoMinor:    1,
        Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
        ContentLength: int64(len(body)),
        Request:       nil,
        Header:        make(http.Header, 0),
    }

    client.On("Do", mock.Anything).Return(response, errors.New("something bad happened !"))

    country, error := Country("1.1.1.1")

    assert.Empty(t, country)
    assert.NotNil(t, error)
}