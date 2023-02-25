package controller

import "app/models"


func (c *Controller) CreateBranch(req models.CreateBranch) (id string, err error){

	id, err = c.store.Branch().CreateBranch(models.CreateBranch{Name: req.Name})
	if err != nil{
		return "", err
	}

	return id, nil
}

func (c *Controller) UpdateBranch(req models.UpdateBranch) (err error){

	err = c.store.Branch().UpdateBranch(req)
	if err != nil{
		return err
	}
	return nil

}

func (c *Controller) DeleteBranch(req models.BranchPrimaryKey) (err error){

	err = c.store.Branch().DeleteBranch(req)
	if err != nil{
		return err
	}

	return nil
}