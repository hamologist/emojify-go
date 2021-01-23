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
		checkOne := r.Intn(2)
		checkTwo := r.Intn(3)

		if checkTwo == 1 {
			for emojiCount := r.Intn(4); emojiCount > 0; emojiCount-- {
				word = word + emojis[r.Intn(len(emojis))]
			}
		} else if checkOne == 1 {
			word = word + emojis[r.Intn(len(emojis))]
		}
		separated[i] = word
	}

	return model.EmojifyResponse{
		Message: strings.Join(separated, " "),
	}, nil
}

func init() {
	emojiMap := emoji.Map()
	filter := make(map[string]bool)

	for key, val := range emojiMap {
		if !strings.Contains(key, ":flag") {
			emojis = append(emojis, val)
		} else {
			filter[val] = true
		}
	}

	for i, val := range emojis {
		if filter[val] == true {
			emojis[i] = emojis[len(emojis)-1]
			emojis[len(emojis)-1] = ""
			emojis = emojis[:len(emojis)-1]
		}
	}

	source := rand.NewSource(time.Now().Unix())
	r = rand.New(source)
}