DROP TABLE IF EXISTS follows;
CREATE TABLE follows (
    followerId VARCHAR(36) NOT NULL,  -- フォロワーのIDを指定
    followingId VARCHAR(36) NOT NULL,  -- フォローされているユーザーのIDを指定
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (followerId, followingId),  -- 複合主キーを設定
    FOREIGN KEY (followerId) REFERENCES users(userId),  -- 外部キー制約を追加
    FOREIGN KEY (followingId) REFERENCES users(userId)  -- 外部キー制約を追加
);