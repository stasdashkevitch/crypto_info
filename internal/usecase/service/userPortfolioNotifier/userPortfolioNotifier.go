package userportfolionotifier

type UserPortfolioNotifier interface {
	Notify(subjecct, message string) string
}
