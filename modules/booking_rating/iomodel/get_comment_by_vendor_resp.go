package iomodel

import "paradise-booking/entities"

type GetCommentByVendorResp struct {
	ListRating []GetCommentUserByVendor
	// DataVendor *entities.Account `json:"vendor"`
}

type GetCommentUserByVendor struct {
	DataRating    entities.BookingRating
	DataPlace     *entities.Place     `json:"place"`
	DataPostGuide *entities.PostGuide `json:"post_guide"`
	DataUser      entities.Account    `json:"user"`
}
