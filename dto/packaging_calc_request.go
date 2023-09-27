package dto

// PackagingCalcRequest - represents packaging calculate request data
type PackagingCalcRequest struct {
	ItemsCount int `json:"itemsCount"`
}

type PackagingCalcRespItems struct {
	PackageSize  int `json:"packageSize"`
	PackageCount int `json:"packageCount"`
}

type PackagingCalcResponse struct {
	Data []PackagingCalcRespItems `json:"data"`
}

type AddPackageSizeRequest struct {
	PackageSize int `json:"packageSize"`
}

type RemovePackageSizeRequest struct {
	PackageSize int `json:"packageSize"`
}
