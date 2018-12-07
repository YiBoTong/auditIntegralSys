package fun

import (
	"auditIntegralSys/SystemSetup/db/rbac"
	rbacEntity "auditIntegralSys/SystemSetup/entity"
	"auditIntegralSys/Worker/db/menu"
	db_rbac2 "auditIntegralSys/Worker/db/rbac"
	"auditIntegralSys/Worker/entity"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
)

// 根据parentId查询所有菜单及子菜单
func GetAllMenu(parentId int, queryIsUse bool) ([]entity.Menus, error) {
	var allMenu []entity.Menus
	where := g.Map{}
	if queryIsUse {
		where["is_use"] = 1
	}

	menuList, err := db_menu.Menus(parentId, where)
	for _, v := range menuList {
		var childMenu []entity.Menu
		var childList []map[string]interface{}
		id := gconv.Int(v["id"])
		if gconv.Bool(v["has_child"]) {
			childList, err = db_menu.Menus(id, where)
			for _, cv := range childList {
				childMenu = append(childMenu, entity.Menu{
					Id:       gconv.Int(cv["id"]),
					Path:     gconv.String(cv["path"]),
					Order:    gconv.Int(cv["order"]),
					ParentId: gconv.Int(cv["parentId"]),
					HasChild: gconv.Bool(cv["has_child"]),
					Time:     gconv.String(cv["time"]),
					IsUse:    gconv.Bool(cv["is_use"]),
					Meta: entity.Meta{
						Title:   gconv.String(cv["title"]),
						Icon:    gconv.String(cv["icon"]),
						NoCache: gconv.Bool(cv["no_cache"]),
					},
				})
			}
		}
		if err != nil {
			break
		}
		allMenu = append(allMenu, entity.Menus{
			Menu: entity.Menu{
				Id:       id,
				Path:     gconv.String(v["path"]),
				Order:    gconv.Int(v["order"]),
				ParentId: gconv.Int(v["parentId"]),
				HasChild: gconv.Bool(v["has_child"]),
				Time:     gconv.String(v["time"]),
				IsUse:    gconv.Bool(v["is_use"]),
				Meta: entity.Meta{
					Title:   gconv.String(v["title"]),
					Icon:    gconv.String(v["icon"]),
					NoCache: gconv.Bool(v["no_cache"]),
				},
			},
			Children: childMenu,
		})
	}
	return allMenu, err
}

// 根据parentId和角色key查询所有菜单及子菜单
func GetAllMenuRbac(parentId int, key string) ([]rbacEntity.Rbacs, error) {
	var allMenu []rbacEntity.Rbacs

	menuList, err := db_rbac.Get(key, parentId)
	for _, v := range menuList {
		var childMenu []rbacEntity.Rbac
		var childList []map[string]interface{}
		menuId := gconv.Int(v["id"])
		if gconv.Bool(v["has_child"]) {
			childList, err = db_rbac.Get(key, menuId)
			for _, cv := range childList {
				childMenu = append(childMenu, rbacEntity.Rbac{
					Id:      gconv.Int(cv["rid"]),
					Key:     gconv.String(cv["key"]),
					Title:   gconv.String(cv["title"]),
					MenuId:  gconv.Int(cv["id"]),
					IsRead:  gconv.Bool(cv["is_read"]),
					IsWrite: gconv.Bool(cv["is_write"]),
				})
			}
		}
		if err != nil {
			break
		}
		allMenu = append(allMenu, rbacEntity.Rbacs{
			Rbac: rbacEntity.Rbac{
				Id:      gconv.Int(v["rid"]),
				Key:     gconv.String(v["key"]),
				Title:   gconv.String(v["title"]),
				MenuId:  menuId,
				IsRead:  gconv.Bool(v["is_read"]),
				IsWrite: gconv.Bool(v["is_write"]),
			},
			Children: childMenu,
		})
	}
	return allMenu, err
}

// 根据parentId和角色key查询有权限的菜单及子菜单
func GetRbacMenu(parentId int, key string) ([]entity.RbacMenus, error) {
	var allMenu []entity.RbacMenus

	menuList, err := db_rbac2.GetRbacMenu(key, parentId)
	for _, v := range menuList {
		var childMenu []entity.RbacMenu
		var childList []map[string]interface{}
		menuId := gconv.Int(v["id"])
		if gconv.Bool(v["has_child"]) {
			childList, err = db_rbac2.GetRbacMenu(key, menuId)
			for _, cv := range childList {
				childMenu = append(childMenu, entity.RbacMenu{
					Menu: entity.Menu{
						Id:       gconv.Int(cv["id"]),
						Path:     gconv.String(cv["path"]),
						Order:    gconv.Int(cv["order"]),
						ParentId: gconv.Int(cv["parentId"]),
						HasChild: gconv.Bool(cv["has_child"]),
						Time:     gconv.String(cv["time"]),
						IsUse:    gconv.Bool(cv["is_use"]),
						Meta: entity.Meta{
							Title:   gconv.String(cv["title"]),
							Icon:    gconv.String(cv["icon"]),
							NoCache: gconv.Bool(cv["no_cache"]),
						},
					},
					Rbac: entity.Rbac{
						IsRead:  gconv.Bool(cv["is_read"]),
						IsWrite: gconv.Bool(cv["is_write"]),
					},
				})
			}
		}
		if err != nil {
			break
		}
		allMenu = append(allMenu, entity.RbacMenus{
			RbacMenu: entity.RbacMenu{
				Menu: entity.Menu{
					Id:       gconv.Int(v["id"]),
					Path:     gconv.String(v["path"]),
					Order:    gconv.Int(v["order"]),
					ParentId: gconv.Int(v["parentId"]),
					HasChild: gconv.Bool(v["has_child"]),
					Time:     gconv.String(v["time"]),
					IsUse:    gconv.Bool(v["is_use"]),
					Meta: entity.Meta{
						Title:   gconv.String(v["title"]),
						Icon:    gconv.String(v["icon"]),
						NoCache: gconv.Bool(v["no_cache"]),
					},
				},
				Rbac: entity.Rbac{
					IsRead:  gconv.Bool(v["is_read"]),
					IsWrite: gconv.Bool(v["is_write"]),
				},
			},
			Children: childMenu,
		})
	}
	return allMenu, err
}
