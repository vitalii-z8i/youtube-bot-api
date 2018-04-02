--' example.sql

CREATE TABLE IF NOT EXISTS chats (
    ID integer PRIMARY KEY
);
CREATE TABLE IF NOT EXISTS users (
    ID integer PRIMARY KEY,
    FirstName varchar
);
CREATE TABLE IF NOT EXISTS messages (
    ID integer PRIMARY KEY,
    FromID integer,
    ChatID integer
    Date timestamp,
    Text text,
    PrevID integer,
    NextID integer,

    FOREIGN KEY (FromID) REFERENCES users (ID)
    ON DELETE CASCADE ON UPDATE NO ACTION,
    FOREIGN KEY (ChatID) REFERENCES chats (ID)
    ON DELETE CASCADE ON UPDATE NO ACTION,

    FOREIGN KEY (PrevID) REFERENCES messages (ID)
    FOREIGN KEY (NextID) REFERENCES messages (ID)
);
CREATE INDEX IF NOT EXISTS messages_user ON messages (FromID);
CREATE INDEX IF NOT EXISTS messages_chat ON messages (ChatID);
CREATE INDEX IF NOT EXISTS prev_messages ON messages (PrevID);
CREATE INDEX IF NOT EXISTS next_messages ON messages (NextID);