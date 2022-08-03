package domain

type ExclusiveTitle struct {
	ID       string
	Name     string
	VendorID string
}

type Vendor struct {
	ID            string
	Name          string
	OriginCountry string
	ExclusiveTitles []ExclusiveTitle
}
