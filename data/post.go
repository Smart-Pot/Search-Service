package data

type Post struct {
	ID      string   `json:"id"`
	UserID  string   `json:"userId" validate:"required"`
	Plant   string   `json:"plant" validate:"required"`
	Info    string   `json:"info" validate:"required"`
	EnvData EnvData  `json:"envData" validate:"required"`
	Images  []string `json:"images" validate:"required"`
	Like    []string `json:"like"`
	Deleted bool     `json:"deleted"`
	Date    string   `json:"date"`
}

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
			   "like":{
				   "type":"text"
			   },
			   "deleted":{
				   "type":"boolean"
			   },
			   "date": { 
				   "type": 	"text" 
			   } 
		   } 
	   } 
}
{"acknowledged":true}
`
