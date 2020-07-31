package muelle

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	sr "github.com/TheSmallBoat/carlo/streaming_rpc"
	"github.com/lithdew/kademlia"
)

type twinProvider struct {
	prd    *sr.Provider
	header map[string]string
}

func newTwinProvider(prd *sr.Provider) *twinProvider {
	return &twinProvider{
		prd:    prd,
		header: map[string]string{ActionHeader: ActionTwin, BodyTitleHeader: PayLoad},
	}
}

func (t *twinProvider) KadID() *kademlia.ID {
	return t.prd.KadID()
}

func (t *twinProvider) Push(data []byte) error {
	stream, errP := t.prd.Push([]string{ServiceTwin}, t.header, ioutil.NopCloser(bytes.NewReader(data)))
	if errP != nil {
		return errors.New(fmt.Sprintf("Unable to push data to the twin %s: %s", t.prd.Addr(), errP))
	}
	res, err := ioutil.ReadAll(stream.Reader)
	if !errors.Is(err, io.EOF) {
		return err
	}
	if stream.Header.Headers[ResponseTwinHeader] == ServiceTwin {
		if stream.Header.Headers[ResponseStatusHeader] == ResponseFailure {
			return errors.New(fmt.Sprintf("The wrong response header for %s : %s", stream.Header.Headers[BodyTitleHeader], string(res)))
		}

		if (stream.Header.Headers[ResponseStatusHeader] == ResponseSuccess) && (stream.Header.Headers[BodyTitleHeader] == ProcessTime) {
			// process time: string(res)
			return nil
		}
		return errors.New(fmt.Sprint("Found the unknown response Header"))
	} else {
		return errors.New(fmt.Sprintf("The wrong response header for %s : %s", ResponseTwinHeader, stream.Header.Headers[ResponseTwinHeader]))
	}
}
