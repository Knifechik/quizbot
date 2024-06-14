package app

type Questions struct {
	Quest      string
	Answers    []string
	GoodAnswer string
}

var Quiz = []map[int]Questions{
	QuizEasy,
	QuizMedium,
	QuizHard,
}

var QuizEasy = map[int]Questions{
	0: {
		Quest: "Стоит ли Эдгару играть в ВОВ?",
		Answers: []string{
			"Да", "Может быть", "Незнаю", "Нет",
		},
		GoodAnswer: "button_4",
	},
	1: {
		Quest: "Работает ли это у Антона на говнокоде?",
		Answers: []string{
			"Конечно нет", "Конечно да", "Это не он писал", "Кто его знает?",
		},
		GoodAnswer: "button_2",
	},
	2: {
		Quest: "Хорошо ли обучает Эдгар?",
		Answers: []string{
			"Не", "Думаю нет", "Да, просто прекрасно", "Он обучает?",
		},
		GoodAnswer: "button_3",
	},
	3: {
		Quest: "Зачем Антон писал это, при том что он сам не знает, как это работает?",
		Answers: []string{
			"Потому что так надо", "Для развития", "Всё равно переписывать всё", "Хз",
		},
		GoodAnswer: "button_1",
	},
	4: {
		Quest: "Вот зачем ты попросил меня добавить пятый вопрос, всё же так хорош было?",
		Answers: []string{
			"Чтоб не расслаблялся", "Так надо", "Это практика", "Я чё, попросил добавить 5 вопрос?",
		},
		GoodAnswer: "button_1",
	},
}

var QuizMedium = map[int]Questions{
	0: {
		Quest: "Хочется ли чтоб это работало и было красиво?",
		Answers: []string{
			"Да, очень", "Нет, зачем оно надо", "Иди в пень", "Воздержусь",
		},
		GoodAnswer: "button_1",
	},
	1: {
		Quest: "Пьёт ли Эдгар пиво?",
		Answers: []string{
			"Он и есть пиво", "Хлещет, как не в себя", "Не, он за ЗОЖ", "Пока никто не видит",
		},
		GoodAnswer: "button_4",
	},
	2: {
		Quest: "Любимая игра Эдгара?",
		Answers: []string{
			"WoW", "Minecraft", "Крестки-нолики", "Играть вменяемого человека",
		},
		GoodAnswer: "button_4",
	},
	3: {
		Quest: "Я пишу уже 9 вопрос и фантазия кончается, но была ли она?",
		Answers: []string{
			"Нет", "Лучше б попросил чатГПТ", "Хватит писать, просто поставь номера", "Вопросы - бомба",
		},
		GoodAnswer: "button_1",
	},
	4: {
		Quest: "Давай просто любимое число моё угадаем",
		Answers: []string{
			"3", "5", "13", "7",
		},
		GoodAnswer: "button_4",
	},
}

var QuizHard = map[int]Questions{
	0: {
		Quest: "Имя актёра Доктора Стрэнджа",
		Answers: []string{
			"Бандерлог Кукумбер", "Бубалех Кандибобер", "Баклажан Киберскотч", "Бенедикт Камбербэтч",
		},
		GoodAnswer: "button_4",
	},
	1: {
		Quest: "Лучший актёр?",
		Answers: []string{
			"Райан Гослинг", "Райан Рейнольдс", "Барбарис Курувпечь", "Джонни Дэпп",
		},
		GoodAnswer: "button_3",
	},
	2: {
		Quest: "Сколько пальцев у человека?",
		Answers: []string{
			"5", "48", "32", "20",
		},
		GoodAnswer: "button_4",
	},
	3: {
		Quest: "Лучший крафт?",
		Answers: []string{
			"Варкрафт", "Старкрафт", "Майнкрафт", "Крафтовое пиво",
		},
		GoodAnswer: "button_1",
	},
	4: {
		Quest: "Да?",
		Answers: []string{
			"Да", "Нет", "Не знаю", "Воздержусь",
		},
		GoodAnswer: "button_1",
	},
}
