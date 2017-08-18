package redir

import (
	"net/url"
	"github.com/satori/go.uuid"
	"encoding/base64"
)

type redirect struct {
	destination url.URL
	uniqueId string
}

type Redirect interface {
	GetDestination() string
	SetDestination(string) error

	ID() string
}
func (r *redirect) GetDestination() string {
	return r.destination.String()
}

func (r *redirect) SetDestination(newDestination string) error {
	dest, err := url.Parse(newDestination)
	if err != nil {
		return err
	}
	r.destination = *dest
	return nil
}

func (r *redirect) ID() string {
	return r.uniqueId
}

func New(url string) (Redirect, error) {
	redir := new(redirect)
	err := redir.SetDestination(url)

	/* Create UUID and encode it as base64 */
	genUUID := uuid.NewV4()
	redir.uniqueId = base64.StdEncoding.EncodeToString(genUUID.Bytes())

	return redir, err
}