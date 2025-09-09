package models

type MdsListing struct {
	ID            int    `json:"id"`
	MdsName       string `json:"mdsName"`
	Comments      string `json:"comments"`
	EffectiveFrom string `json:"effectiveFrom"`
	EffectiveTo   string `json:"effectiveTo"`
	IsPpAgreed    bool   `json:"isPpAgreed"`
	ReferenceNo   string `json:"referenceNo"`
	FilePath      string `json:"filePath"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}
