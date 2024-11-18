DROP TABLE IF EXISTS listing;
CREATE TABLE listing(
	listingId INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	userId VARCHAR(36) NOT NULL,
	tweetId INT NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	listingname VARCHAR(255) NOT NULL,
	listingdescription VARCHAR(255) NOT NULL,
	listingprice DECIMAL(10,2) NOT NULL,
	stock INT NOT NULL,
	FOREIGN KEY (tweetId) REFERENCES tweets(tweetId),
	FOREIGN KEY (userId) REFERENCES users(userId)
);