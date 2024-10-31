DROP TABLE IF EXISTS dms;
CREATE TABLE dms (
	dmsId INT AUTO_INCREMENT NOT NULL PRIMARY KEY ,
    senderId VARCHAR(36) NOT NULL,
    receiverId VARCHAR(36) NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    content VARCHAR(255) NOT NULL,
    media_url VARCHAR(255) DEFAULT NULL,
	status ENUM('unread', 'read') NOT NULL DEFAULT 'unread',
    FOREIGN KEY (senderId) REFERENCES users(userId),     -- 外部キー制約を追加
    FOREIGN KEY (receiverId) REFERENCES users(userId)     -- 外部キー制約を追加
);