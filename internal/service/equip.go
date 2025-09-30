package service

import (
	"encoding/json"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
)

type EquipService struct {
}

func NewEquipService() *EquipService {
	return &EquipService{}
}

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

// TakeOffEquip 卸下指定的装备
// 参数:
//
//	uuid: 要卸下的装备UUID
//
// 返回值:
//
//	string: 操作结果消息
func (s *EquipService) TakeOffEquip(uuid int) string {
	var msg string
	var tag bool
	user := getLocalUser()
	var newEquipArr []int64

	// 从用户装备列表中移除指定装备
	for _, equipId := range user.EquipArr {
		if equipId == int64(uuid) {
			tag = true
			continue
		} else {
			newEquipArr = append(newEquipArr, equipId)
		}
	}

	// 如果找到并移除了装备，则更新用户信息和背包状态
	if tag {
		user.EquipArr = newEquipArr
		updateUserInfo(user)

		bag := getLocalBag()
		// 更新背包中对应物品的状态为未装备
		for i, item := range bag.Items {
			if item.UUid == int64(uuid) {
				bag.Items[i].Status = 0
			}
		}
		updateLocalBag(bag)
		msg += "卸下法器成功"
		return msg
	} else {
		msg += "未穿着该法器"
		return msg
	}
}

// getUserEquipAttributes 计算并返回用户装备的总属性值
// 该函数会遍历用户背包中的装备物品，建立装备UUID到装备信息的映射，
// 然后根据用户已装备的装备ID列表，累加计算所有装备的属性值（攻击、防御、生命、速度），
// 并收集所有不重复的特殊属性
// 返回值：包含用户所有装备属性总和的Equip结构体
func getUserEquipAttributes() model.Equip {
	var (
		uuid2Equip      = make(map[int64]*model.Equip)
		equipAttributes = model.Equip{}
		specialsMap     = make(map[string]struct{})
	)

	user := getLocalUser()
	bag := getLocalBag()

	// 遍历用户背包，筛选出装备类型的物品并建立UUID到装备信息的映射
	for _, item := range bag.Items {
		if item.Type == model.ItemTypeEquip {
			equipInfo := item.EquipInfo
			uuid2Equip[item.UUid] = equipInfo
		}
	}

	// 遍历用户已装备的装备ID，累加计算装备属性总值
	for _, equipId := range user.EquipArr {
		equipInfo := uuid2Equip[equipId]
		equipAttributes.Attack += equipInfo.Attack
		equipAttributes.Defense += equipInfo.Defense
		equipAttributes.Hp += equipInfo.Hp
		equipAttributes.Speed += equipInfo.Speed
		// 收集不重复的特殊属性
		for _, special := range equipInfo.Special {
			if _, ok := specialsMap[special]; !ok {
				equipAttributes.Special = append(equipAttributes.Special, special)
			}
			specialsMap[special] = struct{}{}
		}
	}
	return equipAttributes
}

func (s *EquipService) GetUserEquipAttributes() string {
	e := getUserEquipAttributes()
	bytesE, _ := json.Marshal(e)
	return string(bytesE)
}
