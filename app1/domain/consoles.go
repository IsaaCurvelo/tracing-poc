package domain

type ExclusiveTitle struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Vendor struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	OriginCountry   string           `json:"origin_country"`
	ExclusiveTitles []ExclusiveTitle `json:"exclusive_titles"`
}

type Console struct {
	ID         string  `json:"id"`
	VendorID   string  `json:"vendor_id"`
	Name       string  `json:"name"`
	Vendor     *Vendor `json:"vendor"`
	Generation int     `json:"generation"`
}
