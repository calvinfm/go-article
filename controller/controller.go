package controller

import (
	"context"
	"encoding/json"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"go-article/constant"
	"go-article/entity"
	"go-article/model"
	"go-article/utils"
	"net/http"
	"strings"
	"time"
	//"github.com/elastic/go-elasticsearch/v8"
)

type CommonController struct {
	ArticleModel model.ArticleModel
}

var Ring *redis.Ring
var Mycache *cache.Cache

var Ctx = context.TODO()
var Key = "fadhil"

func setRing() *redis.Ring {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})
	_ = ring.ForEachShard(Ctx, func(ctx context.Context, client *redis.Client) error {
		return client.FlushDB(ctx).Err()
	})

	return ring
}

func newCache(ring *redis.Ring) *cache.Cache {
	return cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
}

func (CommonController CommonController) AddArticle(w http.ResponseWriter, r *http.Request) {

	body := new(entity.AddArticle)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if utils.IsEmptyString(body.Author, body.Title, body.Body) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(utils.GetResNoData(constant.StatusError, constant.FailedAddArticle))
		return
	}

	isSuccess, err := CommonController.ArticleModel.AddArticle(body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(utils.GetResNoData(constant.StatusError, err.Error()))
		return
	}

	if !isSuccess {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(utils.GetResNoData(constant.StatusError, constant.FailedAddArticle))
		return
	}

	Mycache = nil

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(utils.GetResNoData(constant.StatusSuccess, constant.SuccessAddArticle))
	return
}

func (CommonController CommonController) GetArticle(w http.ResponseWriter, r *http.Request) {

	author := r.URL.Query().Get("author")
	query := r.URL.Query().Get("query")

	if Mycache == nil {
		article, err := CommonController.ArticleModel.GetArticle()

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(utils.GetResNoData(constant.StatusError, err.Error()))
			return
		}

		Ring = setRing()
		Mycache = newCache(Ring)

		obj := &article

		if err := Mycache.Set(&cache.Item{
			Ctx:   Ctx,
			Key:   Key,
			Value: obj,
			TTL:   time.Hour,
		}); err != nil {
			panic(err)
		}
	}

	articles := []entity.Article{}
	if err := Mycache.Get(Ctx, Key, &articles); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(utils.GetResNoData(constant.StatusError, err.Error()))
		return
	}

	if !utils.IsEmptyString(author) || !utils.IsEmptyString(query) {
		filteredArticles := []entity.Article{}
		for _, article := range articles {
			temp := false
			if !utils.IsEmptyString(author) && strings.Contains(article.Author, author) {
				temp = true
			}
			if !utils.IsEmptyString(query) &&
				(strings.Contains(article.Title, query) || strings.Contains(article.Body, query)) {
				temp = true
			}
			if temp {
				filteredArticles = append(filteredArticles, article)
			}
		}
		articles = filteredArticles
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(articles)
	return
}
