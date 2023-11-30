package utils

type MessageSet map[string]string

func enLocaleMessages() MessageSet {
	return MessageSet{
		"BadRequest": "I expected something different from you",

		"InternalError": "Something's broken, try again later",

		"Start": "Hi, I will help you shorten your links\n" +
			"Also in the near future I will monitor their attendance and reveal secrets to you\n" +
			"\nBot does not store personal data",

		"Help": "Hi, I will help you shorten your links\n" +
			"I'm ready for these commands:\n\n" +
			"__/newlink__ \\- create a new short link\n" +
			"__/mylinks__ \\- edit and view your links",

		"Default": "I can't help you to do this right now \n" +
			"Send /help and find out which ones\\)",

		"commandNewLink": "Send me a link to shorten",

		"AddLinkNote": "Give it any note in order to not to lose link",

		"MyLinks": "Here is you links list\n\n" +
			"Choose any for further interaction",

		"menuNewLink": "create a new short link",

		"menuMyLinks": "edit and view your links",

		"ManageLink": "🔑 _note_: *%s*\n" +
			"🔗 _link_:\n`%s`\n",
	}
}

func ruLocaleMessages() MessageSet {
	return MessageSet{
		"BadRequest": "Я ожидал от тебя нечто иное",

		"InternalError": "У меня что-то сломалось, попробуй позже",

		"Start": "Привет, я помогу сокращать твои ссылки\n" +
			"Также скоро буду следить за их посещаемостью и открывать тебе секреты\n\n" +
			"Бот не хранит персональные даннные",

		"Help": "Привет, я помогу сокращать твои ссылки\n" +
			"Я готов к таким командам:\n\n" +
			"__/newlink__ - сократить ссылку\n" +
			"__/mylinks__ - список твоих сокращенных ссылок\n",

		"Default": "С этим я тебе пока не могу помочь\n" +
			"Отправь /help и узнаешь какую именно\\)",

		"commandNewLink": "Отправь мне ссылку для сокращения",

		"AddLinkNote": "Добавь какую-нибудь заметку, чтобы не потерять ссылку",

		"MyLinks": "Вот список твоих ссылок\n\n" +
			"Выбери любую ссылку для взаимодействия:",

		"menuNewLink": "сделать новую короткую ссылку",

		"menuMyLinks": "отредактировать и посмотреть ссылки",

		"ManageLink": "🔑 _заметка_: *%s*\n" +
			"🔗 _ссылка_:\n`%s`\n",
	}
}
