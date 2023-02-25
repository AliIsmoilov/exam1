package models

type BranchPrimaryKey struct{
	Id		string		`json:"id"`
}

type Branch struct{
	Id		string		`json:"id"`
	Name	string		`json:"name"`
}

type CreateBranch struct{
	Name	string		`json:"name"`
}

type UpdateBranch struct{
	Id		string		`json:"id"`
	Name	string		`json:"name"`
}