package schema

type AdverseSymptomInfo struct {
	Symptom string `json:"symptom"`
	OAETerm string `json:"oaeTerm"`
}

type AdverseSymptomAdd struct {
	Symptom string `json:"symptom"`
	OAEId   *int64 `json:"oaeId"`
	OAETerm string `json:"oaeTerm"`
}