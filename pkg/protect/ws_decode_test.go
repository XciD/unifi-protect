package protect

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	messageHex := "010101000000006f789c1dcb4b0a02311045d1bdd4d8c06baac9c71d8863175049bd4183a6455a45c4bd1b9d5d2e9cb758db96b5cb5eee57b78db293cee7e9df071fbbe98c422da172ca616eca505d6b488a48aa47f334d065759e8f7c0dd11fb731961f8e00992c436bb4099a13006591cf17ca3523020201010000000034789cab562a2d28c9cc4d55b2323430b13431343030d051ca492c2e094e4dcd030a9a1959181b5b185a9a9b9919d602002dd00c64"

	message, err := hex.DecodeString(messageHex)

	if err != nil {
		panic(err)
	}

	decodedMessage, err := DecodeWsMessage(message)

	action, err := decodedMessage.GetAction()

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "update", action.Action)
	assert.Equal(t, "600ee7a803b6a103870003e9", action.ID)
	assert.Equal(t, "nvr", action.ModelKey)
	assert.Equal(t, "c3409e39-be18-4c3e-bd3b-7306ee3d6ad7", action.NewUpdateID)

	payload := map[string]interface{}{}

	if err := decodedMessage.Payload.GetJSON(&payload); err != nil {
		panic(err)
	}
	assert.Equal(t, float64(104941000), payload["uptime"])
	assert.Equal(t, float64(1628338197661), payload["lastSeen"])
}
