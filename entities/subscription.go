package entities

// Subscription contains an info of subscribed channels
type Subscription struct {
	ID          int64  `db:"ID"`
	UserID      int64  `db:"UserID"`
	ChannelID   string `db:"ChannelID"`
	ChannelName string `db:"ChannelName"`
	ChannelInfo string `db:"ChannelInfo"`
	Channel     YTChannel
}

// YTChannel contains basic info of YT channel itself
type YTChannel struct {
	ChannelID   string `json:"callback_data" db:"YTChannelID"`
	ChannelName string `json:"text" db:"ChannelName"`
	ChannelInfo string `db:"ChannelInfo" json:"-"`
}
