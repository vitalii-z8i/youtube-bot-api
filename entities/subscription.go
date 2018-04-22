package entities

// Subscription contains an info of subscribed channels
type Subscription struct {
	ID        int64 `db:"ID"`
	UserID    int64 `db:"UserID"`
	ChannelID int64 `db:"ChannelID"`
}
