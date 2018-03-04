package workflow

type Notification struct {
	Api			string	`json:"api"`
	PlatId		string 	`json:"platId"`
	BusinessId 	int		`json:"businessId"`
	OpenId		string	`json:"openId"`
}
