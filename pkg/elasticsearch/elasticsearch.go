package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"kumiko/pkg/logger"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/spf13/viper"
)

var Client *elasticsearch.Client

func InitElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			viper.GetString("elasticsearch.addr"), // 例如 "http://localhost:9200"
		},
		Username: viper.GetString("elasticsearch.username"),
		Password: viper.GetString("elasticsearch.password"),
	}

	var err error
	Client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		logger.StdError("Elasticsearch连接失败: %v", err)
	}

	// 测试连接
	res, err := Client.Info()
	if err != nil {
		logger.StdError("Elasticsearch Info失败: %v", err)
	}
	defer res.Body.Close()
	logger.StdInfo("Elasticsearch连接成功")
}

// 新增文档
func IndexDoc(index string, docID string, body interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		logger.StdError("Elasticsearch文档序列化失败: %v", err)
		return err
	}
	res, err := Client.Index(
		index,
		bytes.NewReader(data),
		Client.Index.WithDocumentID(docID),
		Client.Index.WithContext(context.Background()),
	)
	if err != nil {
		logger.StdError("Elasticsearch新增文档失败: %v", err)
		return err
	}
	defer res.Body.Close()
	return nil
}

// 删除文档
func DeleteDoc(index string, docID string) error {
	res, err := Client.Delete(
		index,
		docID,
		Client.Delete.WithContext(context.Background()),
	)
	if err != nil {
		logger.StdError("Elasticsearch删除文档失败: %v", err)
		return err
	}
	defer res.Body.Close()
	return nil
}

// 查询文档
func SearchDoc(index string, query map[string]interface{}) ([]map[string]interface{}, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		logger.StdError("Elasticsearch查询序列化失败: %v", err)
		return nil, err
	}
	res, err := Client.Search(
		Client.Search.WithIndex(index),
		Client.Search.WithBody(&buf),
		Client.Search.WithContext(context.Background()),
	)
	if err != nil {
		logger.StdError("Elasticsearch查询失败: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		logger.StdError("Elasticsearch结果解析失败: %v", err)
		return nil, err
	}

	// 解析 hits
	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	var results []map[string]interface{}
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		results = append(results, source)
	}
	return results, nil
}
