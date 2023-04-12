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

//go:generate ytags -x -f $GOFILE -t [json bson mapstructure form]

type Person struct {
	Person1
	NameAddr string `json:"name_addr" bson:"-" mapstructure:"name_addr" form:"name_addr"`
	Job      string `json:"job" bson:"-" mapstructure:"job" form:"job"`
	Age      int64  `json:"age" bson:"-" mapstructure:"age" form:"age"`
}
type Person1 struct {
	NNNNameAAAddr string `json:"nnn_name_aa_addr" bson:"-" mapstructure:"nnn_name_aa_addr" form:"nnn_name_aa_addr"`
	Job           string `json:"job" bson:"-" mapstructure:"job" form:"job"`
	Age           int64  `json:"age" bson:"-" mapstructure:"age" form:"age"`
}
