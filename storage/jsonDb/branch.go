package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)


type branchRepo struct {
	fileName string
}

func NewBranchRepo(fileName string) *branchRepo {
	return &branchRepo{
		fileName: fileName,
	}
}

func (b *branchRepo) CreateBranch(req models.CreateBranch) (id string, err error){

	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return "", err
	}

	var branches []models.Branch
	err = json.Unmarshal(data, &branches)
	if err != nil {
		return "", err
	}

	id = uuid.New().String()

	branches = append(branches, models.Branch{Id: id, Name: req.Name})

	body, err := json.MarshalIndent(branches, "", " ")
	if err != nil{
		return "", err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil{
		return "", err
	}

	return id, nil
} 

func (b *branchRepo) UpdateBranch(req models.UpdateBranch) (err error){
	
	data, err := ioutil.ReadFile(b.fileName)
	if err != nil{
		return err
	}

	var branches []models.Branch
	
	err = json.Unmarshal(data, &branches)
	if err != nil{
		return err
	}

	for i, branch := range branches{
		if branch.Id == req.Id{
			branches[i].Name = req.Name

			data, err = json.MarshalIndent(branches, "", " ")
			if err != nil{
				return err
			}

			err = ioutil.WriteFile(b.fileName, data, os.ModePerm)
			if err != nil{
				return err
			}

			return nil
		}
	}

	return errors.New("not found user with such id")
}

func (b *branchRepo) DeleteBranch(req models.BranchPrimaryKey) (err error){

	data, err := ioutil.ReadFile(b.fileName)
	if err != nil{
		return err
	}

	var branches []models.Branch
	err = json.Unmarshal(data, &branches)
	if err != nil{
		return err
	}

	for i, branch := range branches{
		if branch.Id == req.Id{
			
			branches = append(branches[:i], branches[i+1:]...)
			
			data, err = json.MarshalIndent(branches, "", " ")
			if err != nil{
				return err
			}

			err = ioutil.WriteFile(b.fileName, data, os.ModePerm)
			if err != nil{
				return err
			}

			return nil
		}
	}

	return errors.New("not found user with input id")
}

func (b *branchRepo) Get_Branches() ([]models.Branch, error){

	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return []models.Branch{}, err
	}

	var branches []models.Branch
	err = json.Unmarshal(data, &branches)
	if err != nil {
		return []models.Branch{}, err
	}
	return branches, nil
}