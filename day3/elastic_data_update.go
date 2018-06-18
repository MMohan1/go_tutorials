package main

import ("context"
	   "encoding/json"
	   "fmt"
	   // "reflect"
	   "time"
       "strings"
       "gopkg.in/olivere/elastic.v5")


// elastichosts=['edge-prod-ha-ta-poc01-es56-app.hirealchemy.com']
// elasticport=443
// query_radius=150
// query_timeout=60000s
// username=edge-prod-ha-ta-poc01-es56-app
// password=edge-prod-ha-ta-poc01-e$56-app


const uname string = "edge-prod-ha-ta-poc01-es56-engg"
const pwd  string = "edge-prod-ha-ta-poc01-e$56-Engg"
const connection_str string = "http://139.162.4.158:4507"

func main() {
     client, err := elastic.NewClient(
             elastic.SetURL(connection_str),
             elastic.SetSniff(false),
             elastic.SetHealthcheckInterval(10*time.Second),
             // elastic.SetRetrier(NewCustomRetrier()),
             elastic.SetBasicAuth(uname, pwd))
	 if err != nil {
        fmt.Println("not able to connect to elastic client")
	    	panic(err)
	        }
      index_name := "candidates"
      exists, err := client.IndexExists(index_name).Do(context.Background())
      if err != nil {
             fmt.Println("some error come to get the index existens")
      }
      if !exists {
         fmt.Println("the index doesn't exists")
	     // Index does not exist yet.
      }
      var group_id string = "5b0eab6702538874de300e56"
      query := elastic.NewBoolQuery()
      query = query.Must(elastic.NewTermQuery("group_ids", group_id))
      // src, err := query.Source()
     //  if err != nil {
     //     panic(err)
     // }
     // data, err := json.MarshalIndent(src, "", "  ")
     // if err != nil {   
     // panic(err)
     // }    
     // fmt.Println(string(data))
     // var ttyp data
    search, err := client.Search().
            Index(index_name).    
            Type("resume").
            Query(query).
            From(0).Size(1495).Do(context.Background())
    if err != nil {
             fmt.Println("some error come to fire the elastic query")
      }
   fmt.Println("Total documents foudn is => ", search.Hits.TotalHits)
   for _, hit := range search.Hits.Hits {
       fmt.Println("I am here")
       var dat map[string] interface{}
        err := json.Unmarshal(*hit.Source, &dat)
        if err != nil {
           fmt.Println(err)    
            // Deserialization 
        // Work with tweet
        }
        fmt.Println("-------------")
        file_data := dat["company_level_config"].(map[string]interface{})
        default_data := file_data["default"].(map[string]interface{})
        file_name := default_data["file_name"]
        fmt.Println(file_name.(string))
        candidate_name_arry := strings.Split(file_name.(string), ".")
        new_name := candidate_name_arry[0]
        fmt.Println(new_name)
        dat["name"] = new_name
        doc_id := hit.Id
        client.Update().Index(index_name).Type("resume").Id(doc_id).Doc(map[string]interface{}{"name": dat["name"]}).Do(context.Background())
        // fmt.Println("updated id: ", update.Id)
        fmt.Println(dat["name"], doc_id)
    
    }
}


