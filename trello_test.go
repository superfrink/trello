package trello

import (
	//	"github.com/dchest/uniuri"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTrelloHasContext(t *testing.T) {
	trello := NewTrello(TrelloParams{})
	assert.NotNil(t, trello.Context)
}

func TestNewTrelloHasClient(t *testing.T) {
	trello := NewTrello(TrelloParams{})
	assert.NotNil(t, trello.Context.client)
}
