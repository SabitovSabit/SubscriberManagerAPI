package models


type Subscriber struct {
	Id      int
	Name    string
	IsFree  bool
	AddDate string
}

type BaseResult struct{
	Result   bool
	Message  string
}