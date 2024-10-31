DROP TABLE IF EXISTS likes;
CREATE TABLE likes (
    userId VARCHAR(36) NOT NULL,  -- 'userId related to users.userId' の修正。ユーザーIDを指定
    tweetId INT NOT NULL,  -- 'tweetId related to tweets.tweetId' の修正。ツイートIDを指定
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (userId, tweetId),  -- 複合主キーを設定
    FOREIGN KEY (userId) REFERENCES users(userId),  -- 外部キー制約を追加
    FOREIGN KEY (tweetId) REFERENCES tweets(tweetId)  -- 外部キー制約を追加
);