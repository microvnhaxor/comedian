package chat

import (
	"testing"

	"github.com/maddevsio/comedian/config"
	"github.com/stretchr/testify/assert"
)

func TestCleanMessage(t *testing.T) {
	c, err := config.Get()
	assert.NoError(t, err)
	s, err := NewSlack(c)
	assert.NoError(t, err)
	s.myUsername = "comedian"
	assert.NoError(t, err)
	text, ok := s.cleanMessage("<@comedian> hey there")
	assert.Equal(t, "hey there", text)
	assert.True(t, ok)
	text, ok = s.cleanMessage("What's up?")
	assert.Equal(t, "What's up?", text)
	assert.False(t, ok)
}

func TestSendMessage(t *testing.T) {
	c, err := config.Get()
	assert.NoError(t, err)
	s, err := NewSlack(c)
	assert.NoError(t, err)
	channel := "AnychannelID"
	err = s.SendMessage(channel, "yoooo!")
	assert.NoError(t, err)
}
