package middleware

import (
	"errors"
)



var adminRoleID=1;
//Admin ...
func Admin(ID int) error {


	if(adminRoleID==ID){
		return nil
	}
	return errors.New("user not admin")
	
}
