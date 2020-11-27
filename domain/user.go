package domain

type UserSimpleInfo struct {
	UserID      int     `gorm:"column:id" json:"UserID"`
	UserName    string  `gorm:"column:name" json:"UserName"`
	UserHN      *string `gorm:"column:handle_name" json:"UserHN"`
	UserImg     *string `gorm:"column:img" json:"UserImg"`
	GoaledCount int64   `gorm:"column:goaled_count" json:"GoaledCount"`
}

type LoginInfo struct {
	UserID   int    `gorm:"column:id"`
	UserName string `gorm:"column:name"`
	Password []byte `gorm:"column:password"`
}

type UserDetailInfo struct {
	ID          int     `gorm:"column:id; type:autoIncrement"`
	Name        string  `gorm:"column:name;unique"`
	HN          *string `gorm:"column:handle_name"`
	Img         *string `gorm:"column:img"`
	GoaledCount int64   `gorm:"column:goaled_count; default:0"`
	FinalGoal   *string `gorm:"column:final_goal"`
	Profile     *string `gorm:"column:profile"`
	Twitter     *string `gorm:"column:twitter"`
	Instagram   *string `gorm:"column:instagram"`
	Facebook    *string `gorm:"column:facebook"`
	Github      *string `gorm:"column:github"`
	URL         *string `gorm:"column:url"`
}
