DROP TABLE IF EXISTS listing;
CREATE TABLE listing(
	listingId BIGINT NOT NULL PRIMARY KEY,
	userId VARCHAR(36) NOT NULL,
	tweetId INT NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	listingname VARCHAR(255) NOT NULL,
	listingdescription VARCHAR(255) NOT NULL,
	`condition` VARCHAR(255) NOT NULL,
	listingprice INT NOT NULL,
	type VARCHAR(255) NOT NULL,
	stock INT NOT NULL,
	FOREIGN KEY (tweetId) REFERENCES tweets(tweetId),
	FOREIGN KEY (userId) REFERENCES users(userId)
);