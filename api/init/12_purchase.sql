DROP TABLE IF EXISTS purchase;

CREATE TABLE purchase(
	purchaseId INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	userId VARCHAR(36) NOT NULL,
	listingId BIGINT NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	status ENUM('注文確定', '配送中', '出荷完了','キャンセル') NOT NULL DEFAULT '注文確定',
	FOREIGN KEY (userId) REFERENCES users(userId),
	FOREIGN KEY (listingId) REFERENCES listing(listingId)
);