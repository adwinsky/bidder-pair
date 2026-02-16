package openrtb

type BidRequest struct {
	ID     string `json:"id"`
	Imp    []Imp  `json:"imp"`
	Site   *Site  `json:"site,omitempty"`
	App    *App   `json:"app,omitempty"`
	Device Device `json:"device"`
	User   User   `json:"user"`
}

type Imp struct {
	ID     string  `json:"id"`
	Banner *Banner `json:"banner,omitempty"`
	Video  *Video  `json:"video,omitempty"`
}

type Banner struct {
	W int `json:"w"`
	H int `json:"h"`
}

type Video struct {
	W int `json:"w"`
	H int `json:"h"`
}

type Site struct {
	Domain string `json:"domain"`
	Page   string `json:"page"`
}

type App struct {
	Name   string `json:"name"`
	Bundle string `json:"bundle"`
}

type Device struct {
	UA string `json:"ua"`
	IP string `json:"ip"`
}

type User struct {
	ID string `json:"id"`
}

func (b *BidRequest) ImpIDs() []string {
	out := make([]string, 0, len(b.Imp))
	for _, imp := range b.Imp {
		out = append(out, imp.ID)
	}
	return out
}
