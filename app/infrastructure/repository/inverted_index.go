package irepository

import (
	"context"
	"encoding/json"
	"main/config"
	invertedIndexE "main/domain/invertedindex/entity"
	invertedIndexR "main/domain/invertedindex/repository"
	documentScoreVO "main/domain/invertedindex/valueobject"
	"main/util"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type invertedIndex struct {
	rc *config.RedisClient
}

func NewInvertedIndexRepository(rc *config.RedisClient) invertedIndexR.InvertedIndex {
	return &invertedIndex{rc: rc}
}

const ALL_DOCUMENT_TOKENS_BUM = "all_document_tokens_num"

func (ii *invertedIndex) MultiUpsert(invertedIndexes []invertedIndexE.InvertedIndex) error {
	// TODO: エラーハンドリングとトランザクションの追加

	for _, invertedIndex := range invertedIndexes {
		token := invertedIndex.Get().Token
		documentScore := invertedIndex.Get().DocumentScore

		interfaceSlice := make([]interface{}, len(documentScore))
		for i, v := range documentScore {
			jsonData, err := json.Marshal(v)
			if err != nil {
				return util.NewAddOperationError([]string{err.Error()})
			}
			interfaceSlice[i] = jsonData
		}

		// Redisに追加
		err := ii.rc.Client.RPush(context.Background(), token, interfaceSlice...).Err()
		if err != nil {
			return util.NewAddOperationError([]string{err.Error()})
		}
	}
	return nil
}
func (ii *invertedIndex) Get(word string) (invertedIndexE.InvertedIndex, error) {
	values, err := ii.rc.Client.LRange(context.Background(), word, 0, -1).Result()
	if err != nil {
		return nil, util.NewInternalServerError()
	}
	var documentIds []documentScoreVO.DocumentScore
	for _, value := range values {
		var docId documentScoreVO.DocumentScore
		err := json.Unmarshal([]byte(value), &docId)
		if err != nil {
			return nil, util.NewInternalServerError()
		}
		documentIds = append(documentIds, docId)
	}

	invertedIndex := invertedIndexE.NewInvertedIndex(word, documentIds)
	return invertedIndex, nil
}

func (ii *invertedIndex) AddAllDL(invertedIndexes []invertedIndexE.InvertedIndex) error {

	err := ii.rc.Client.Set(context.Background(), ALL_DOCUMENT_TOKENS_BUM, len(invertedIndexes), 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (ii *invertedIndex) GetAllDLNum() (allDLNum int, err error) {
	result, err := ii.rc.Client.Get(context.Background(), ALL_DOCUMENT_TOKENS_BUM).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, util.NewInternalServerError()
	}
	allDLNum, err = strconv.Atoi(result)
	if err != nil {
		return 0, util.NewInternalServerError()
	}
	return allDLNum, nil
}
