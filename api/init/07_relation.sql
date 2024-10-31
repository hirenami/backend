DROP TABLE IF EXISTS relations;
CREATE TABLE relations (
    tweetId INT NOT NULL,  -- 関連するツイートのID
    replyId INT NOT NULL,  -- リプライされたツイートのID
    PRIMARY KEY (tweetId, replyId),  -- 複合主キー
    FOREIGN KEY (tweetId) REFERENCES tweets(tweetId),  -- 外部キー制約
    FOREIGN KEY (replyId) REFERENCES tweets(tweetId)   -- 外部キー制約
);