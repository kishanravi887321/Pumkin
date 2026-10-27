package models

type User struct {
	UserId        string `json:"userId" bson:"userId"`
	UserName      string `json:"userName" bson:"userName"`
	Email         string `json:"email" bson:"email"`
	Password      string `json:"-" bson:"password"`
	AvtarURL      string `json:"avatarUrl" bson:"avatarUrl"`
	TelePhone     string `json:"telePhone" bson:"telePhone"`
	RecoveryEmail string `json:"recoveryEmail" bson:"recoveryEmail"`
	UserMasterKey string `json:"userMasterKey" bson:"userMasterKey"`
	CreatedAt     int64  `json:"createdAt" bson:"createdAt" description:"Timestamp of user creation"`
}
