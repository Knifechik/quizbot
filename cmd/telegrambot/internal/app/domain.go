package app

type UserInfo struct {
	ChatID           int
	QuestNumber      int
	LastMessageID    int
	CountRightAnswer int
	Answer           string
	Finished         bool
	Quests           []Questions
}
