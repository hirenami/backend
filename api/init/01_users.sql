DROP TABLE IF EXISTS users;
CREATE TABLE users (
  firebaseUid VARCHAR(128) PRIMARY KEY, 	-- FIREBASE UID
  userId VARCHAR(36) NOT NULL UNIQUE, 	-- USER Userid
  username VARCHAR(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  header_image VARCHAR(255) NOT NULL DEFAULT 'https://firebasestorage.googleapis.com/v0/b/term6-namito-hirezaki.appspot.com/o/grey.png?alt=media&token=3cfa3b15-5419-4807-932d-d19e10c52ff3',
  icon_image VARCHAR(255) NOT NULL DEFAULT 'https://firebasestorage.googleapis.com/v0/b/term6-namito-hirezaki.appspot.com/o/default_profile_400x400.png?alt=media&token=44ace5f1-ef11-481f-9618-ba7d07e96b5d',
  biography VARCHAR(255) DEFAULT '' NOT NULL,
  isPrivate BOOLEAN NOT NULL DEFAULT false,
  isFrozen BOOLEAN  NOT NULL DEFAULT false,
  isDeleted BOOLEAN NOT NULL DEFAULT false,
  isAdmin BOOLEAN NOT NULL DEFAULT false
);
