DROP TABLE IF EXISTS notifications;
CREATE TABLE notifications (
	notificationsId int AUTO_INCREMENT PRIMARY KEY,
    senderId varchar(36) NOT NULL,  -- 送信者ID
    replyId varchar(36) NOT NULL,    -- 受信者ID
    type VARCHAR(255) NOT NULL,        -- 通知の種類
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status ENUM('unread', 'read') DEFAULT 'unread',  -- ステータス
    contentId int,               -- リツイートやツイートID
    FOREIGN KEY (senderId) REFERENCES users(userId),
    FOREIGN KEY (replyId) REFERENCES users(userId)
);