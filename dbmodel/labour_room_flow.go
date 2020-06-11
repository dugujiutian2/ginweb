package dbmodel

type LabourRoomFlow struct {
	RoomFlowId    int    `json:"room_flow_id" xorm:"not null pk autoincr comment('id') INT(10)"`
	LabourUnionId int    `json:"labour_union_id" xorm:"not null comment('公会id') index INT(10)"`
	Roomid        int    `json:"roomid" xorm:"not null comment('房间id') index INT(10)"`
	PresenterId   int    `json:"presenter_id" xorm:"not null comment('主持人id') index INT(10)"`
	GiveId        int    `json:"give_id" xorm:"not null comment('gift_send.give_id礼物id') unique INT(10)"`
	DefendOrderId string `json:"defend_order_id" xorm:"not null comment('守护订单id') index VARCHAR(10)"`
	Value         int    `json:"value" xorm:"not null comment('价值,也就是总收入') INT(10)"`
	Share         string `json:"share" xorm:"not null comment('公会长分成收入') DECIMAL(10,1)"`
	OtherShare    string `json:"other_share" xorm:"not null comment('相对人的分成收入(主播或者用户)') DECIMAL(10,1)"`
	Rate          int    `json:"rate" xorm:"not null comment('分成比例,单位是万') INT(10)"`
	Ptime         int    `json:"ptime" xorm:"not null comment('发生时间') index INT(10)"`
	Status        int    `json:"status" xorm:"not null comment('状态(0未核销,1已核销)') index SMALLINT(5)"`
}
