package route

import (
	"strings"

	"github.com/tidwall/gjson"

	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/errx"
	"github.com/starudream/go-lib/httpx"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/router"

	"github.com/starudream/mirai-go"
	"github.com/starudream/openai-go"
)

const modelGPT35 = "gpt-3.5-turbo"

var (
	botQQ int64

	targetGroups = map[int64]bool{}

	openaiClient *openai.Client
)

func init() {
	botQQ = config.GetInt64("mirai.bot_qq")

	_targetGroup := config.GetInt("mirai.target_group")
	targetGroups[int64(_targetGroup)] = true
	_targetGroups := config.GetIntSlice("mirai.target_groups")
	for _, tg := range _targetGroups {
		targetGroups[int64(tg)] = true
	}

	apiKey := config.GetString("openai.api_key")
	if apiKey == "" {
		log.Fatal().Msgf("openai api key is empty")
	}
	var err error
	openaiClient, err = openai.NewClient(apiKey, openai.WithClient(httpx.Client()))
	if err != nil {
		log.Fatal().Msgf("openai client init failed: %v", err)
	}
}

func index(c *router.Context) {
	req := &mirai.Message{}
	if c.BindJSON(req) != nil {
		return
	}

	if req.Type != mirai.MessageTypeGroup {
		return
	}

	bs, _ := req.MarshalJSON()
	body := string(bs)

	if gjson.Get(body, "messageChain.#(type==At).target").Int() != botQQ {
		return
	}

	target := gjson.Get(body, "sender.group.id").Int()

	if !targetGroups[target] {
		return
	}

	var (
		targetQQ  = gjson.Get(body, "sender.id").Int()
		messageId = gjson.Get(body, "messageChain.#(type==Source).id").Int()
		message   = strings.TrimPrefix(gjson.Get(body, "messageChain.#(type==Plain).text").String(), " ")
	)

	res, _, err := openaiClient.ChatCompletions(openai.ChatCompletionsReq{
		Model: modelGPT35,
		Messages: []openai.CompletionsMessage{
			openai.NewCompletionsMessage("user", message),
		},
	})
	if err != nil {
		log.Error().Msgf("openai chat failed: %v", err)
		c.AbortWithError(errx.ErrInternal)
		return
	}

	messages := res.GetMessage()
	if len(messages) == 0 {
		c.AbortWithError(errx.ErrInternal)
		return
	}

	text := messages[0].Content

	log.Ctx(c).Info().Msgf("openai chat: %s", text)

	resp := &Resp{
		Command: "sendGroupMessage",
		Content: &mirai.SendMessageReq{
			Target: target,
			Quote:  messageId,
			MessageChain: []mirai.MessageInfoInterface{
				&mirai.MessageInfoAt{
					Type:   mirai.MessageInfoTypeAt,
					Target: targetQQ,
				},
				&mirai.MessageInfoPlain{
					Type: mirai.MessageInfoTypePlain,
					Text: " " + text,
				},
			},
		},
	}

	c.OK(resp)
}

type Resp struct {
	Command string `json:"command,omitempty"`
	Content any    `json:"content,omitempty"`
}
