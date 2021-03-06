package main

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/bwmarrin/discordgo"

	"github.com/rishabhsamb/HyBot/commandUtilities"
	"github.com/rishabhsamb/HyBot/outbursts"
	"github.com/rishabhsamb/HyBot/weatherCaller"
)

type commandHandler struct {
	obHandler outbursts.OutburstHandler
}

func (ch *commandHandler) init(ctx context.Context, client *firestore.Client) {
	ch.obHandler.Init(client, ctx)
	ch.obHandler.LoadOutbursts()
}

func (ch *commandHandler) driver(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "?") {
		fmt.Println("about to execute")
		ch.obHandler.Execute(s, m.ChannelID, strings.TrimSpace(m.Content))
		return
	}

	if strings.HasPrefix(m.Content, "?addOutburst ") {
		toParse := strings.TrimPrefix(m.Content, "?addOutburst ")
		commandSlice := commandUtilities.CommandSplit(toParse)
		key := ""
		foundKey := false
		var (
			messages       []string
			randomMessages []string
		)
		for _, val := range commandSlice {
			fmt.Println(val)
			if strings.HasPrefix(val, "key=") && !foundKey {
				keyIntermediate := strings.TrimPrefix(val, "key=\"") // REFACTOR THIS INTO A FUNCTION YOU REPEAT IT LIKE 4 TIMES
				key = strings.TrimSuffix(keyIntermediate, "\"")
				foundKey = true
			} else if strings.HasPrefix(val, "message=") {
				messageIntermediate := strings.TrimPrefix(val, "message=\"")
				messageFinal := strings.TrimSuffix(messageIntermediate, "\"")
				messages = append(messages, messageFinal)
			} else if strings.HasPrefix(val, "randomMessage=") {
				randomMessageIntermediate := strings.TrimPrefix(val, "randomMessage=\"")
				randomMessageFinal := strings.TrimSuffix(randomMessageIntermediate, "\"")
				randomMessages = append(randomMessages, randomMessageFinal)
			}
		}
		fmt.Println(key)
		fmt.Println(messages)
		fmt.Println(randomMessages)
		ch.obHandler.AddOutburst(key, messages, randomMessages)
		return
	}

	if strings.HasPrefix(m.Content, "?weather") {
		toParse := strings.TrimPrefix(m.Content, "?weather ")
		commandSlice := commandUtilities.CommandSplit(toParse)
		cityFinal := ""
		for _, val := range commandSlice {
			if strings.HasPrefix(val, "city=") {
				cityIntermediate := strings.TrimPrefix(val, "city=\"")
				cityFinal = strings.TrimSuffix(cityIntermediate, "\"")
				break
			}
		}
		cityFinal = strings.TrimSpace(cityFinal)
		if cityFinal != "" {
			stringMap := weatherCaller.WeatherCaller(cityFinal)
			s.ChannelMessageSend(m.ChannelID, "I hope you can read Kelvin b/c I'm not converting this for you.")
			for k, v := range stringMap {
				s.ChannelMessageSend(m.ChannelID, k+": "+v)
			}
		}
	}
}
