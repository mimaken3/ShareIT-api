package service

import "github.com/mimaken3/ShareIT-api/domain/model"

// スライスからn番目の要素を削除
func UnsetTopics(s []model.Topic, i int) []model.Topic {
	// [:i] インデックi以前
	// [i+1:] インデックスi+1以降
	// ... s[:i]とs[i+1:]を組み合わせる
	s = append(s[:i], s[i+1:]...)

	//新しいスライスを用意
	n := make([]model.Topic, len(s))
	copy(n, s)

	return n
}

// スライスからn番目の要素を削除
func UnsetUsers(s []model.User, i int) []model.User {
	// [:i] インデックi以前
	// [i+1:] インデックスi+1以降
	// ... s[:i]とs[i+1:]を組み合わせる
	s = append(s[:i], s[i+1:]...)

	//新しいスライスを用意
	n := make([]model.User, len(s))
	copy(n, s)

	return n
}
