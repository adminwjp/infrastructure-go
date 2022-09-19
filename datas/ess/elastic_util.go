package ess

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type ElasticUtil struct {
	Host    string
	Account string
	Pwd     string
	// 作用域 容易 消失 指针 为 nil
	//panic: runtime error: invalid memory address or nil pointer dereference 无法使用
	Client *elastic.Client // 未使用 自带释放 资源?
	indexs []string
	used   int //使用 个数
}

//http://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Celasticsearch/%E6%93%8D%E4%BD%9Celasticsearch.html
//var client *elastic.Client // 未使用 自带释放 资源? 全局 变量 才有效 ? 最好 不使用时 不要使用该模块 即 改包 不然会 自动加载
func ElasticUtilByHost(host string) *ElasticUtil {
	return &ElasticUtil{Host: host}
}

func ElasticUtilInstance(host string, account string, pwd string) *ElasticUtil {
	return &ElasticUtil{Host: host, Account: account, Pwd: pwd, indexs: make([]string, 20)}
}

func NewElasticUtil() *ElasticUtil {
	return ElasticUtilByHost("http://127.0.0.1:9200")
}
func (elasticUtil *ElasticUtil) getClient() *elastic.Client {
	return elasticUtil.GetClient()
	//return  elasticUtil.client //panic: runtime error: invalid memory address or nil pointer dereference 无法使用
	//return  client

}
func GetClient(host ...string) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(host...))
	if err != nil {
		fmt.Println(client, err)
		//panic(err) // Handle error
		fmt.Println(" connection elastic fail")
		return nil
	} else {
		for _, v := range host {
			info, code, err := client.Ping(v).Do(context.Background())
			if err != nil {
				panic(err)
				fmt.Println("ping connection elastic fail")
				return nil
			}
			fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
			fmt.Println("connection elastic success")
		}
	}
	return client
}

func (elasticUtil *ElasticUtil) GetClient() *elastic.Client {
	if elasticUtil.Client!=nil{
		return  elasticUtil.Client
	}
	return GetClient(elasticUtil.Host)
}

func (elasticUtil *ElasticUtil) Connection() bool {
	//if client!=nil{
	if elasticUtil.Client != nil {
		return true
	}
	temp := elasticUtil.GetClient()
	//client=temp
	elasticUtil.Client = temp
	return temp != nil
}

func (elasticUtil *ElasticUtil) ReConnection() bool {
	temp := elasticUtil.getClient()
	if temp != nil {
		elasticUtil.Close()
	}
	temp = elasticUtil.GetClient()
	//client=temp
	elasticUtil.Client = temp
	return temp != nil
}

func (elasticUtil *ElasticUtil) Close() bool {
	temp := elasticUtil.getClient()
	if temp != nil {
		if elasticUtil.indexs != nil && len(elasticUtil.indexs) > 0 {
			for i := range elasticUtil.indexs {
				temp.CloseIndex(elasticUtil.indexs[i])
			}
			return true
		}
		return false
	}
	return false
}
func (elasticUtil *ElasticUtil) updateIndex(index string) {
	return
	if elasticUtil.indexs == nil {
		elasticUtil.indexs = make([]string, 20)
	}
	if elasticUtil.used == 0 {
		elasticUtil.indexs[elasticUtil.used] = index
		elasticUtil.used++
	} else {
		exists := false
		for i := 0; i < elasticUtil.used; i++ {
			if elasticUtil.indexs[i] == index {
				exists = true
				break
			}
		}
		if !exists {
			if elasticUtil.used == len(elasticUtil.indexs) {
				temp := make([]string, len(elasticUtil.indexs)+10)
				copy(temp, elasticUtil.indexs)
				//temp=append(temp,elasticUtil.indexs)
				//elasticUtil.indexs=temp
				elasticUtil.indexs = temp
			}
			elasticUtil.indexs[elasticUtil.used] = index
			elasticUtil.used++
		}
	}
}

func (elasticUtil *ElasticUtil) Insert(index string, data interface{}) bool {
	//elasticUtil.client.Index().Index(index).Do(context.Background())
	respnse, err := elasticUtil.getClient().Index().Index(index).BodyJson(data).Do(context.Background())
	elasticUtil.updateIndex(index)
	if err != nil {
		//panic(err) // Handle error
		return false
	}
	return respnse.Status > 0
}

func (elasticUtil *ElasticUtil) DeleteIndex(indexs ...string) bool {
	indices, err := elasticUtil.getClient().DeleteIndex(indexs...).Do(context.Background())
	if err != nil {
		return false
	}
	return indices.Acknowledged
}

func (elasticUtil ElasticUtil) CreateIndex(index string) bool {
	indices, err := elasticUtil.getClient().CreateIndex(index).Do(context.Background())
	if err != nil {
		return false
	}
	return indices.Acknowledged
}

func (elasticUtil *ElasticUtil) ExistsIndex(indexs ...string) bool {
	exists, err := elasticUtil.getClient().IndexExists(indexs...).Do(context.Background())
	if err != nil {
		return false
	}
	return exists
}

func (elasticUtil *ElasticUtil) GetIndex(indexs ...string) map[string]*elastic.IndicesGetResponse {
	indices, err := elasticUtil.getClient().IndexGet(indexs...).Do(context.Background())
	if err != nil {
		return nil
	}
	return indices
}

func (elasticUtil *ElasticUtil) CloseIndex(index string) bool {
	close, err := elasticUtil.getClient().CloseIndex(index).Do(context.Background())
	if err != nil {
		return false
	}
	return close.Acknowledged
}

func (elasticUtil *ElasticUtil) GetIndexTemplate(index string) *elastic.IndicesGetIndexTemplateResponse {
	response, err := elasticUtil.getClient().IndexGetIndexTemplate(index).Do(context.Background())
	if err != nil {
		return nil
	}
	return response
}

// create update delete
func (elasticUtil *ElasticUtil) Bulk(index string, requests ...elastic.BulkableRequest) bool {
	response, err := elasticUtil.getClient().Bulk().Index(index).Add(requests...).Do(context.Background())
	if err != nil {
		println("buk support create update delete,ex  ", err)
		return false
	}
	for i := range response.Succeeded() {
		return response.Succeeded()[i].Status > 0
	}
	return false
}

func (elasticUtil *ElasticUtil) PostSearch(query elastic.Query, from int, size int, sorts []string, sortDescs []string, indexs ...string) *elastic.SearchResult {
	return elasticUtil.search(true, query, from, size, sorts, sortDescs, indexs...)
}

func (elasticUtil *ElasticUtil) GetSearch(query elastic.Query, from int, size int, sorts []string, sortDescs []string, indexs ...string) *elastic.SearchResult {
	return elasticUtil.search(false, query, from, size, sorts, sortDescs, indexs...)
}

func (elasticUtil *ElasticUtil) search(post bool, query elastic.Query, from int, size int, sorts []string, sortDescs []string, indexs ...string) *elastic.SearchResult {
	search := elasticUtil.getClient().Search(indexs...)
	if sorts != nil {
		for i := range sorts {
			search = search.Sort(sorts[i], true)
		}
	}
	if sortDescs != nil {
		for i := range sortDescs {
			search = search.Sort(sorts[i], false)
		}
	}
	search = search.From(from).Size(size)
	if post {
		search = search.PostFilter(query)
	} else {
		search = search.Query(query)
	}
	result, err := search.Do(context.Background())
	if err != nil {
		return nil
	}
	return result
}

func (elasticUtil *ElasticUtil) InsertJson(index string, data string) bool {
	respnse, err := elasticUtil.getClient().Index().Index(index).BodyString(data).Do(context.Background())
	elasticUtil.updateIndex(index)
	if err != nil {
		return false
	}
	return respnse.Status > 0
}

func (elasticUtil *ElasticUtil) SaveOrUpdate(index string, id string, data interface{}) bool {
	respnse, err := elasticUtil.getClient().Update().Index(index).Id(id).Doc(data).Do(context.Background())
	elasticUtil.updateIndex(index)
	if err != nil {
		return false
	}
	return respnse.Status > 0
}

func (elasticUtil *ElasticUtil) Delete(index string, id string) bool {
	respnse, err := elasticUtil.getClient().Delete().Id(id).Index(index).Do(context.Background())
	elasticUtil.updateIndex(index)
	if err != nil {
		return false
	}
	return respnse.Status > 0
}

func (elasticUtil *ElasticUtil) Get(index string, id string) map[string]interface{} {
	result, err := elasticUtil.getClient().Get().Id(id).Index(index).Do(context.Background())
	elasticUtil.updateIndex(index)
	if err != nil {
		return nil
	}
	return result.Fields
	//return result.Found
}

func (elasticUtil *ElasticUtil) CreateIndexAlias(index string, alias string) bool {
	result, err := elasticUtil.getClient().Alias().Add(index, alias).Do(context.Background())
	if err != nil {
		return false
	}
	return result.Acknowledged
}

func (elasticUtil *ElasticUtil) DropIndexAlias(index string, alias string) bool {
	result, err := elasticUtil.getClient().Alias().Remove(index, alias).Do(context.Background())
	if err != nil {
		return false
	}
	return result.Acknowledged
}

func (elasticUtil *ElasticUtil) Count(indexs ...string) []int {
	response, err := elasticUtil.getClient().CatCount().Index(indexs...).Do(context.Background())
	if err != nil {
		return nil
	}
	var counts []int = make([]int, len(response))
	for i := range response {
		counts[i] = response[i].Count
	}
	return counts
}

func (elasticUtil *ElasticUtil) QueryCount(query elastic.Query, indexs ...string) int64 {
	count, err := elasticUtil.getClient().Count(indexs...).Query(query).Do(context.Background())
	if err != nil {
		return -1
	}
	return count
}

func (elasticUtil *ElasticUtil) Exists(index string, id string) bool {
	exists, err := elasticUtil.getClient().Exists().Index(index).Id(id).Do(context.Background())
	if err != nil {
		return false
	}
	return exists
}

func (elasticUtil *ElasticUtil) Disk() []elastic.CatAllocationResponseRow {
	disks, err := elasticUtil.getClient().CatAllocation().Do(context.Background())
	if err != nil {
		return nil
	}
	return disks
}

func (elasticUtil ElasticUtil) Index() []elastic.CatIndicesResponseRow {
	use, err := elasticUtil.getClient().CatIndices().Do(context.Background())
	if err != nil {
		return nil
	}
	return use
}

func (elasticUtil *ElasticUtil) DeleteByQuery(queyr elastic.Query, indexs ...string) *elastic.BulkIndexByScrollResponse {
	respnse, err := elasticUtil.getClient().DeleteByQuery(indexs...).Query(queyr).Do(context.Background())
	if err != nil {
		return nil
	}
	return respnse
}

func (elasticUtil ElasticUtil) DeleteScript(id string) bool {
	respnse, err := elasticUtil.getClient().DeleteScript().Id(id).Do(context.Background())
	if err != nil {
		return false
	}
	return respnse.Acknowledged
}

// create update delete
func (elasticUtil *ElasticUtil) BulkProcessor(index string, bulks []elastic.BulkableRequest) int64 {
	procssor, err := elasticUtil.getClient().BulkProcessor().Name(index).Do(context.Background())
	//procssor,err := elasticUtil.getClient().BulkProcessor().Name(index).BulkSize(bulkSize).BulkActions(len(bulks)).Do(context.Background())
	if err != nil {
		return -1
	}
	if bulks != nil && len(bulks) > 0 {
		for i := range bulks {
			procssor.Add(bulks[i])
		}
		err = procssor.Start(context.Background())
		if err != nil {
			return -1
		}
		err = procssor.Stop()
		if err != nil {
			return -1
		}
	}
	return procssor.Stats().Succeeded
}

func (elasticUtil *ElasticUtil) XPackWatchExecuteByString(body string) *elastic.XPackWatcherExecuteWatchResponse {
	execute, err := elasticUtil.getClient().XPackWatchExecute().BodyString(body).Do(context.Background())
	if err != nil {
		return nil
	}

	return execute
}

func (elasticUtil *ElasticUtil) MultiGet(items ...*elastic.MultiGetItem) *elastic.MgetResponse {
	response, err := elasticUtil.getClient().MultiGet().Add(items...).Do(context.Background())
	if err != nil {
		return nil
	}
	return response
}

type EsSearchHandler interface {
	Handler(data map[string]interface{}) bool
}

//应该 最多 只能 跑 1w ? 只能用分片
func (elasticUtil *ElasticUtil) MultiSearch(requests ...*elastic.SearchRequest) *elastic.MultiSearchResult {
	result, err := elasticUtil.getClient().MultiSearch().Add(requests...).Do(context.Background())
	if err != nil {
		return nil
	}
	return result
}

//应该 最多 只能 跑 1w ? 只能用分片
func (elasticUtil *ElasticUtil) MultiSearchByHandler(esSearchHandler EsSearchHandler, total int, request *elastic.SearchRequest) bool {

	var (
		from = 1
		rows = 1000
		size = 1000
	)
	client := elasticUtil.getClient()
	for {
		request.Size(size).From(from)
		result, err := client.MultiSearch().Add(request).Do(context.Background())
		if err != nil {
			return false
		}
		for i := range result.Responses {
			//panic: runtime error: invalid memory address or nil pointer dereference 跑一会儿就崩了
			//应该 最多 只能 跑 1w ? 只能用分片
			println(result.Responses[i].Hits)
			if result.Responses[i].Hits == nil || result.Responses[i].Hits.Hits == nil {
				break
			}
			for s := range result.Responses[i].Hits.Hits {
				var data map[string]interface{}
				buffer := result.Responses[i].Hits.Hits[s].Source
				err := json.Unmarshal(buffer, &data)
				if err != nil {
					continue
				}
				esSearchHandler.Handler(data)
			}
		}
		if total > size {
			size += rows
			from += rows
			if size >= total {
				size = total
			}
		} else {
			break
		}
	}
	return true
}

func (elasticUtil *ElasticUtil) HasPlugin(name string) bool {
	exists, err := elasticUtil.getClient().HasPlugin(name)
	if err != nil {
		return false
	}

	return exists
}

func (elasticUtil *ElasticUtil) Plugins() []string {
	plugins, err := elasticUtil.getClient().Plugins()
	if err != nil {
		return nil
	}

	return plugins
}
