DROP TABLE IF EXISTS purchase;

CREATE TABLE purchase(
	purchaseId INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	userId VARCHAR(36) NOT NULL,
	listingId BIGINT NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	status ENUM('listing', 'completed', 'cancelled','purchased') NOT NULL DEFAULT 'purchased',
	FOREIGN KEY (userId) REFERENCES users(userId),
	FOREIGN KEY (listingId) REFERENCES listing(listingId)
);