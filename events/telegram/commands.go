package telegram

import (
	"log"
	"math/rand"
	"strings"
)

const (
	changeTargetCmd = "/target"
	HumorCmd        = "/anecdote"
	HelpCmd         = "/help"
	StartCmd        = "/start"
)

const admin = "Vanv1k"

var TargetPerson = ""

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if targetSenderCmd(username) {
		return p.sendReply(chatID)
	}
	commandParts := strings.Fields(text)
	if len(commandParts) == 0 {
		return nil
	}
	switch commandParts[0] {
	case HumorCmd:
		return p.sendRandomAnecdote(chatID)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	case changeTargetCmd:
		if username == admin && len(commandParts) > 1 {
			TargetPerson = commandParts[1]
			message := "New target is " + TargetPerson
			return p.tg.SendMessage(chatID, message)
		}
	}
	return nil
}

func (p *Processor) sendReply(chatID int) error {
	return p.tg.SendSticker(chatID, msgReply[rand.Intn(len(msgReply))])
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func (p *Processor) sendRandomAnecdote(chatID int) error {
	return p.tg.SendMessage(chatID, `Тут должен был быть анекдот. Потом`)
}

func targetSenderCmd(username string) bool {
	return username == TargetPerson
}
