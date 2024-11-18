DROP TABLE IF EXISTS purchase;

CREATE TABLE purchase(
	purchaseId INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	userId VARCHAR(36) NOT NULL,
	listingId INT NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (userId) REFERENCES users(userId),
	FOREIGN KEY (listingId) REFERENCES listing(listingId)
);