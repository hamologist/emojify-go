package transformer

import (
	"github.com/enescakir/emoji"
	"hamologist.com/emojify/pkg/model"
	"math/rand"
	"time"
)

var (
	emojis []string
	r *rand.Rand
)

func EmojifyTransformer(payload model.EmojifyPayload) (model.EmojifyResponse, error) {
	return model.EmojifyResponse{
		Message: emojis[r.Intn(len(emojis))],
	}, nil
}

func init() {
	emojiMap := emoji.Map()

	i := 0
	emojis = make([]string, len(emojiMap))
	for _, val := range emojiMap {
		emojis[i] = val
		i += 1
	}

	source := rand.NewSource(time.Now().Unix())
	r = rand.New(source)
}