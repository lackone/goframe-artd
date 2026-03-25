package logic

const (
	STATUS_ENABLE  = 1
	STATUS_DISABLE = -1
	AUTH_TYPE_MENU = 1
	AUTH_TYPE_BTN  = 2
	IS_SUPER_YES   = 1
	IS_SUPER_NO    = -1
)

var ImageFields = []string{
    "website_logo",
    "website_favicon",
}

type CommonLogic struct{
	
}

func NewCommonLogic() *CommonLogic {
	return &CommonLogic{}
}