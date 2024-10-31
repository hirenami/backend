DROP TABLE IF EXISTS tweets;
CREATE TABLE tweets(
	tweetId INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	userId VARCHAR(36) NOT NULL,
	retweetId INT DEFAULT NULL,
	isQuote BOOLEAN NOT NULL DEFAULT false,
	isReply BOOLEAN NOT NULL DEFAULT false,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	content VARCHAR(255) NOT NULL,
	media_url VARCHAR(255) DEFAULT NULL,
	likes INT NOT NULL DEFAULT 0,
	retweets INT NOT NULL DEFAULT 0,
	replies INT NOT NULL DEFAULT 0,
	impressions INT NOT NULL DEFAULT 0,
	isDeleted BOOLEAN NOT NULL DEFAULT false,
	FOREIGN KEY (userId) REFERENCES users(userId)
);