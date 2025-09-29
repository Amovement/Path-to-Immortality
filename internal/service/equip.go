package service

import "github.com/Amovement/Path-to-Immortality-WASM/internal/model"

// equipItem 装备指定UUID的道具
// uuid: 要装备的道具唯一标识符
// bag: 用户背包信息
// 返回值: 装备操作的结果消息
func equipItem(uuid int64, bag *model.Bag) string {
	var msg string
	var existed bool
	wearingEquipType := make(map[uint]struct{}) // 身上已经装备的法器类型
	uuid2Equip := make(map[int64]*model.Equip)
	var willWearEquipType uint

	// 遍历背包物品，检查目标道具是否存在并收集装备信息
	for _, item := range bag.Items {
		if item.UUid == uuid {
			existed = true
			if item.Type != model.ItemTypeEquip {
				return "不是法器..."
			}
			willWearEquipType = item.EquipInfo.Type
		}
		if item.Type == model.ItemTypeEquip {
			uuid2Equip[item.UUid] = item.EquipInfo
		}
	}

	// 获取用户当前已装备的法器类型
	user := getLocalUser()
	for _, equipId := range user.EquipArr {
		if ty, ok := uuid2Equip[equipId]; ok {
			wearingEquipType[ty.Type] = struct{}{}
		}
	}

	// 执行装备逻辑：检查道具存在性和装备位置冲突
	if !existed {
		msg += "法器不存在"
	} else {
		if _, ok := wearingEquipType[willWearEquipType]; ok {
			msg += "身上已经装备相同部位的法器"
		} else {
			user.EquipArr = append(user.EquipArr, uuid)
			updateUserInfo(user)
			for i, _ := range bag.Items {
				if bag.Items[i].UUid == uuid {
					bag.Items[i].Status = 1
					break
				}
			}
			updateLocalBag(bag)
			msg += "装备`" + uuid2Equip[uuid].Name + "`成功"
		}
	}

	return msg
}
