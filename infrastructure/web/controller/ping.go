package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dynastymasra/avalon/infrastructure/provider"

	"github.com/dynastymasra/avalon/config"

	log "github.com/sirupsen/logrus"

	"github.com/dynastymasra/avalon/infrastructure/web/formatter"
)

func Ping(provider *provider.Instance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		log := log.WithField(config.RequestID, r.Context().Value(config.HeaderRequestID))

		if err := provider.Postgres.Ping(); err != nil {
			log.WithError(err).Infoln("Failed ping postgres")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		q, _ := BuildMappingToString(SearchSetting(), SearchMappings("consumerID"))
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formatter.ObjectResponse(q).Stringify())
	}
}

// Mapper for mapping elastic mapping
type Mapper map[string]interface{}

// FieldContents struct
type FieldContents struct {
	Type     string `json:"type,omitempty"`
	Index    string `json:"index,omitempty"`
	Analyzer string `json:"analyzer,omitempty"`
}

// Fields struct
type Fields struct {
	Original   *FieldContents `json:"original,omitempty"`
	English    *FieldContents `json:"english,omitempty"`
	Indonesian *FieldContents `json:"indonesian,omitempty"`
	Ngrammer   *FieldContents `json:"ngrammer,omitempty"`
}

// Mapping struct
type Mapping struct {
	Type   string `json:"type"`
	Fields Fields `json:"fields"`
}

// TypeMatcher struct
type TypeMatcher struct {
	MatchMappingType string  `json:"match_mapping_type"`
	PathMatch        string  `json:"path_match"`
	Mapping          Mapping `json:"mapping"`
}

// MapperID struct
type MapperID struct {
	ConsumerID    string
	StringMatcher Mapper
}

// SearchSetting configuration elastic
func SearchSetting() Mapper {
	var ngramTokenizer, avalonNgramTokenizer, ngramAnalyzer, avalonNgramAnalyzer, settingGroups, analysis Mapper

	ngramTokenizer = make(Mapper)
	avalonNgramTokenizer = make(Mapper)
	ngramAnalyzer = make(Mapper)
	avalonNgramAnalyzer = make(Mapper)
	settingGroups = make(Mapper)
	analysis = make(Mapper)

	ngramAnalyzer["tokenizer"] = "avalon_ngram_tokenizer"
	ngramAnalyzer["filter"] = []string{"lowercase"}
	avalonNgramAnalyzer["avalon_ngram_analyzer"] = ngramAnalyzer

	ngramTokenizer["type"] = "ngram"
	ngramTokenizer["min_gram"] = 3
	ngramTokenizer["max_gram"] = 10
	ngramTokenizer["token_chars"] = []string{"letter", "digit", "whitespace"}
	avalonNgramTokenizer["avalon_ngram_tokenizer"] = ngramTokenizer

	settingGroups["tokenizer"] = avalonNgramTokenizer
	settingGroups["analyzer"] = avalonNgramAnalyzer

	analysis["analysis"] = settingGroups

	return analysis
}

// SearchMappings function
func SearchMappings(consumerID string) Mapper {
	mappings := NewMapperID(consumerID)
	return mappings.ToMap()
}

// BootstrapStringMapper fucntion
func BootstrapStringMapper() *TypeMatcher {
	ngrammer := FieldContents{}
	ngrammer.Type = "text"
	ngrammer.Analyzer = "avalon_ngram_analyzer"

	fields := Fields{}
	fields.Ngrammer = &ngrammer

	mapping := Mapping{}
	mapping.Type = "keyword"
	mapping.Fields = fields

	typematcher := TypeMatcher{}
	typematcher.MatchMappingType = "string"
	typematcher.PathMatch = "*"
	typematcher.Mapping = mapping

	return &typematcher
}

// NewMapperID function
func NewMapperID(consumerID string) *MapperID {
	nm := MapperID{ConsumerID: consumerID}

	stringTemplate, err := convertStructToMapViaJSON(BootstrapStringMapper())
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "mapping.go",
			"package":  "model.es",
			"function": "NewMapperID.convertStructToMapViaJSON",
		}).Error(err)
	}

	nm.StringMatcher = stringTemplate

	return &nm
}

func convertStructToMapViaJSON(i interface{}) (Mapper, error) {

	var toMap map[string]interface{}

	toJSON, err := json.Marshal(i)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "mapping.go",
			"package":  "model.es",
			"function": "convertStructToMapViaJSON.Marshal",
		}).Error(err)
		return nil, err
	}

	err = json.Unmarshal(toJSON, &toMap)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "mapping.go",
			"package":  "model.es",
			"function": "convertStructToMapViaJSON.Unmarshal",
		}).Error(err)
		return nil, err
	}

	return toMap, nil

}

// BuildMappingToString parse mapping interface to string
func BuildMappingToString(settings, mappings Mapper) (string, error) {
	var mapper Mapper

	mapper = make(Mapper)

	mapper["settings"] = settings
	mapper["mappings"] = mappings["mappings"]

	byteValue, err := json.Marshal(mapper)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "mapping.go",
			"package":  "model.es",
			"function": "BuildMappingToString.Marshal",
		}).Error(err)

		return "", err
	}

	return string(byteValue), nil
}

// ToMap function
func (mi *MapperID) ToMap() Mapper {

	var mappings, esType, dt, strMatcher Mapper

	dt = make(Mapper)
	strMatcher = make(Mapper)
	esType = make(Mapper)
	mappings = make(Mapper)

	strMatcher["strings"] = mi.StringMatcher
	dt["dynamic_templates"] = []Mapper{strMatcher}
	esType[mi.ConsumerID] = dt
	mappings["mappings"] = esType

	return mappings
}
