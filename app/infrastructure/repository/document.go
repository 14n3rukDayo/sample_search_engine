package irepository

import (
	"context"
	"main/config"
	documentE "main/domain/document/entity"
	documentR "main/domain/document/repository"
	"main/util"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type document struct {
	rc *config.RedisClient
}

const DOCUMENT_ID_RESERVED = "doc-"

func NewDocumentRepository(rc *config.RedisClient) documentR.Document {
	return &document{rc: rc}
}

func prefixDocumentId(documentId int) string {
	id := DOCUMENT_ID_RESERVED + strconv.Itoa(documentId)
	return id
}

func (d *document) GetPrefixDocumentIdReserved() string {

	return DOCUMENT_ID_RESERVED
}
func (d *document) Add(document documentE.Document) error {
	err := d.rc.Client.SetNX(context.Background(), prefixDocumentId(document.Get().DocumentId), document.Get().Description, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (d *document) Get(documentId int) (document documentE.Document, err error) {
	result, err := d.rc.Client.Get(context.Background(), prefixDocumentId(documentId)).Result()
	if err == redis.Nil {
		return nil, util.NewInternalServerError()
	} else if err != nil {
		return nil, util.NewInternalServerError()
	}
	document = documentE.NewDocumentEntity(documentId, result)
	return document, nil
}

func (d *document) GenerateID() (documentId int, err error) {
	id, err := d.rc.Client.Incr(context.Background(), DOCUMENT_ID_KEY).Result()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (d *document) GetTotalNum() (total int, err error) {
	total, err = countDocPrefixedKeys(d.rc.Client, DOCUMENT_ID_RESERVED)
	if err != nil {
		return 0, util.NewInternalServerError()
	}
	return total, nil
}
