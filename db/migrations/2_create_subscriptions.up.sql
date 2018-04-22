CREATE TABLE IF NOT EXISTS subscriptions (
    ID integer PRIMARY KEY,
    UserID integer,
    ChannelID varchar,
    ChannelName varchar,
    ChannelInfo text,
    UNIQUE (UserID, ChannelID)

    FOREIGN KEY (UserID) REFERENCES users (ID)
    ON DELETE CASCADE ON UPDATE NO ACTION
);

CREATE INDEX IF NOT EXISTS subscriptions_user ON subscriptions (UserID);