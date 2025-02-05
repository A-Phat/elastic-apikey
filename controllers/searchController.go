package controllers

import (
	"bytes"
	"context"
	"elastic-apikey/config"
	"encoding/json"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gin-gonic/gin"
)

func InsertDocument(c *gin.Context) {
	var doc map[string]interface{}
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := json.Marshal(doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode JSON"})
		return
	}

	req := esapi.IndexRequest{
		Index:      "api_key_index",
		DocumentID: doc["id"].(string),
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), config.ESClient)
	if err != nil || res.IsError() {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert document"})
		return
	}
	defer res.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Document inserted successfully"})
}

func SearchDocument(c *gin.Context) {
	query := c.Query("q")

	searchBody := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": query,
			},
		},
	}

	data, _ := json.Marshal(searchBody)
	req := esapi.SearchRequest{
		Index: []string{"api_key_index"},
		Body:  bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), config.ESClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search"})
		return
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	c.JSON(http.StatusOK, result)
}

