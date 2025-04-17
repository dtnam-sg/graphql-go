package resolver

import (
	"demo-graphql/graph/model"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func findRoles(input model.CreateUserInput, roleList []*model.Role) ([]*model.Role, error) {
	var roles []*model.Role
	for _, roleID := range input.RoleIds {
		if roleID == nil {
			continue
		}
		for _, role := range roleList {
			if role.ID == *roleID {
				roles = append(roles, role)
				break
			}
		}
	}
	return roles, nil

}

func findFunctions(input model.CreateUserInput, functionList []*model.Function) ([]*model.Function, error) {
	var functions []*model.Function
	if input.FunctionIds != nil {
		for _, funcID := range input.FunctionIds {
			if funcID == nil {
				continue
			}
			for _, f := range functionList {
				if f.ID == *funcID {
					functions = append(functions, f)
					break
				}
			}
		}
	}
	return functions, nil
}

func NewBaseInfoWithDefault() *model.BaseInfo {
	return &model.BaseInfo{
		Status:      "ACTIVE",
		CreatedBy:   "System",
		UpdatedBy:   "System",
		CreatedTime: timestamppb.New(time.Now()).AsTime(),
		UpdatedTime: timestamppb.New(time.Now()).AsTime(),
	}
}
