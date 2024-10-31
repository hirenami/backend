package dao

import (
	"api/sqlc"
	"context"
	"database/sql"
)

func (d *Dao) Test(ctx context.Context, tx *sql.Tx, userId, content, media_url string) error {
	// トランザクション用のクエリを生成
	txQueries := d.WithTx(tx)

	var mediaUrl sql.NullString
	if media_url == "" {
		mediaUrl = sql.NullString{
			String: "",
			Valid:  false, // NULLを示す
		}
	} else {
		mediaUrl = sql.NullString{
			String: media_url,
			Valid:  true, // 有効なURL
		}
	}

	arg := sqlc.CreateTweetParams{
		Userid:   userId,
		Content:  content,
		MediaUrl: mediaUrl,
	}

	// トランザクション内でクエリを実行
	return txQueries.CreateTweet(ctx, arg)
}
