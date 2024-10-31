DROP TABLE IF EXISTS blocks;
CREATE TABLE blocks (
    blockerId VARCHAR(36) NOT NULL,  -- 'related to users.userId' を修正。ブロッカーのIDを指定
    blockedId VARCHAR(36) NOT NULL,   -- 'related to users.userId' を修正。ブロックされたユーザーのIDを指定
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (blockerId, blockedId),  -- 複合主キーを設定
    FOREIGN KEY (blockerId) REFERENCES users(userId),  -- 外部キー制約を追加
    FOREIGN KEY (blockedId) REFERENCES users(userId)   -- 外部キー制約を追加
);