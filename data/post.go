package data

// Post is a structure used for deserializing data in Elasticsearch.
type Post struct {
	ID      string   `json:"id"`
	UserID  string   `json:"userId"`
	Plant   string   `json:"plant"`
	Info    string   `json:"info"`
	EnvData EnvData  `json:"envData"`
	Images  []string `json:"images"`
	Like    []string `json:"like"`
	Date    string   `json:"date"`
}

// EnvData is a structure used for serializing data in Elasticsearch.
type EnvData struct {
	Humidity    string `json:"humidity" validate:"required"`
	Temperature string `json:"temperature" validate:"required"`
	Light       string `json:"light" validate:"required"`
}

// PostMapping is a constant used for create index if it is not exist.
const PostMapping = `
{
	"posts": { 
	"properties": { 
			   "id": { 
				   "type":    "text" 
			   },
			   "userid": { 
				   "type":    "text" 
			   }, 
			   "plant": { 
				   "type":    "text" ,
				   "analyzer": "autocomplete"
			   }, 
			   "info": { 
				   "type":    "text" 
			   }, 
			   "content": { 
				   "type":    "text" 
			   },
			   "envdata": { 
				   "type":"nested",
				   "properties": {
					"humidity":{
					   "type": "text"
				   },
				   "temperature":{
					   "type":"text"
				   },
				   "light":{
					   "type":"text"
				   }
				   }
			   }, 
			   "date": { 
				   "type":    "text" 
			   } 
		   } 
	   } 
}
{"acknowledged":true}
`
