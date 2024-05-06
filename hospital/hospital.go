package hospital

type HospitalSystem struct {
	Name       string
	Url        string
	PricingURL string
	FileNames  []string
}

type Hospital struct {
	Name         string
	FileName     string
	HospitalCode string
}
