package models

func UpdateLikes(postID int) error {

	tempDB := DBManager.Table("component_like")
	// 使用 db.Exec 方法更新點讚數
	result := tempDB.Exec("UPDATE posts SET likes = likes + 1 WHERE id = ?", postID)
	return result.Error
}
