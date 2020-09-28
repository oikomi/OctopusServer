package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommand(t *testing.T) {

	cmd := NewCommand()
	assert.NotNil(t, cmd)
	for _,cl  := range cmd.commandlines {
		if cl.Name == "start" {
			assert.Equal(t, "start a new gateway server", cl.Usage, "expected %s, got %s", "start a new gateway server", cl.Usage)
		}
	}

}
