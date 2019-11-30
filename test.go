package main

// //go:generate easytags $GOFILE json,bson,mapstructure,form
// type Person struct {
// 	NameAddr string `json:"name_addr" bson:"name_addr" mapstructure:"name_addr" form:"name_addr"`
// 	Job      string `json:"job" bson:"job" mapstructure:"job" form:"job"`
// 	Age      int64  `json:"age" bson:"age" mapstructure:"age" form:"age"`
// }
// type Person1 struct {
// 	NNNNameAAAddr string `json:"nnn_name_aa_addr" bson:"nnn_name_aa_addr" mapstructure:"nnn_name_aa_addr" form:"nnn_name_aa_addr"`
// 	Job           string `json:"job" bson:"job" mapstructure:"job" form:"job"`
// 	Age           int64  `json:"age" bson:"age" mapstructure:"age" form:"age"`
// }

//go:generate ytags -f $GOFILE -t json,bson,mapstructure,form

type Person struct {
	Person1
	NameAddr string `json:"NameAddr" bson:"NameAddr" mapstructure:"NameAddr" form:"NameAddr"`
	Job      string `json:"Job" bson:"Job" mapstructure:"Job" form:"Job"`
	Age      int64  `json:"Age" bson:"Age" mapstructure:"Age" form:"Age"`
}
type Person1 struct {
	NNNNameAAAddr string `json:"NNNNameAAAddr" bson:"NNNNameAAAddr" mapstructure:"NNNNameAAAddr" form:"NNNNameAAAddr"`
	Job           string `json:"Job" bson:"Job" mapstructure:"Job" form:"Job"`
	Age           int64  `json:"Age" bson:"Age" mapstructure:"Age" form:"Age"`
}
