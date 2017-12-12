package server

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/almostmoore/kadastr/api_server"
	"github.com/almostmoore/kadastr_tg/server/processors"
)

type Server struct {
	APIToken      string
	ApiClient     *api_server.Client
	commandRouter *processors.CommandRouter
}

func (s *Server) Run() {
	bot, err := tgbotapi.NewBotAPI(s.APIToken)
	if err != nil {
		log.Fatal("Не удалось запустить бота", err)
	}

	log.Printf("Авторизован под аккаунтом %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal("Не удалось подписаться на обновления", err)
	}

	s.BindCommands()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		s.ProcessMessage(bot, &update)
	}
}

func (s *Server) BindCommands() {
	s.commandRouter = &processors.CommandRouter{}

	featureProcessor := processors.NewFeatureProcessor(s.ApiClient)
	s.commandRouter.AddProcessor("search", featureProcessor)
	s.commandRouter.SetDefaultCommand("search")

	addParsingTaskProcessor := processors.NewAddParsingTaskProcessor(s.ApiClient)
	s.commandRouter.AddProcessor("doparsing", addParsingTaskProcessor)

	listParsingTaskProcessor := processors.NewListParsingTaskProcessor(s.ApiClient)
	s.commandRouter.AddProcessor("listparsing", listParsingTaskProcessor)

	helpProcessor := processors.HelpProcessor{}
	s.commandRouter.AddProcessor("start", helpProcessor)
	s.commandRouter.AddProcessor("help", helpProcessor)
}

func (s *Server) ProcessMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	answer, err := s.commandRouter.Run(update)

	if err != nil {
		answer = tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка. "+err.Error())
	}

	answer.ParseMode = tgbotapi.ModeMarkdown

	bot.Send(answer)
}
