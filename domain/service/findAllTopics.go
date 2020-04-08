package service

import (
	"log"
	"sort"

	"github.com/mimaken3/ShareIT-api/domain/model"
)

// 全トピックを取得
func (t *topicServiceStruct) FindAllTopics() (topics []model.Topic, err error) {
	topics, err = t.topicRepo.FindAllTopics()

	if err != nil {
		log.Println(err)
	}

	if len(topics) != 0 {
		// 「その他」を一番最後に移動する
		topic := model.Topic{}
		topic = topics[0]
		topics = unset(topics, 0)

		// クイックソートで並び替え
		sort.Slice(topics, func(i, j int) bool {
			strI := topics[i].TopicName
			strJ := topics[j].TopicName
			return strI < strJ
		})

		// 「その他」が最後にある、文字列順にソートしたトピック一覧を作成
		topics = append(topics, topic)
	}

	return
}

// スライスからn番目の要素を削除
func unset(s []model.Topic, i int) []model.Topic {
	// [:i] インデックi以前
	// [i+1:] インデックスi+1以降
	// ... s[:i]とs[i+1:]を組み合わせる
	s = append(s[:i], s[i+1:]...)

	//新しいスライスを用意
	n := make([]model.Topic, len(s))
	copy(n, s)

	return n
}
