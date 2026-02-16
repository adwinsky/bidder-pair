package decision

type DecisionRequest struct {
	RequestID string   `json:"request_id"`
	UserID    string   `json:"user_id"`
	ImpIDs    []string `json:"imp_ids"`
}

type DecisionResponse struct {
	Allow      bool    `json:"allow"`
	Price      float64 `json:"price"`
	CreativeID string  `json:"creative_id"`
}
