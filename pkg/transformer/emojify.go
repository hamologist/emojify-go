package transformer

import (
	"github.com/enescakir/emoji"
	"hamologist.com/emojify/pkg/model"
	"math/rand"
	"strings"
	"time"
)

var (
	emojis []string
	r *rand.Rand
)

func EmojifyTransformer(payload model.EmojifyPayload) (model.EmojifyResponse, error) {
	message := payload.Message
	separated := strings.Split(message, " ")

	for i, word := range separated {
		word = word + emojis[r.Intn(len(emojis))]
		separated[i] = word
	}

	return model.EmojifyResponse{
		Message: strings.Join(separated, " "),
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