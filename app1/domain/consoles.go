package domain

type Vendor struct {
	ID            string
	Name          string
	OriginCountry string
}

type Console struct {
	ID         string
	VendorID   string
	Name       string
	Vendor     *Vendor
	Generation int
}
