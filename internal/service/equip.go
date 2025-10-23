package service

import (
	"encoding/json"
	"fmt"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/repo"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/utils"
	"math/rand"
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
// 然后根据用户已装备的装备ID列表，累加计算所有装备的属性值（攻击、防御、体魄、速度），
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
		if equipInfo == nil {
			continue
		}
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

// ForgeEquip 锻造
func (s *EquipService) ForgeEquip(uuid int64) string {
	var msg string
	// 检查身上是否存在装备 uuid
	var existed bool
	var forged bool // 是否锻造成功
	bag := getLocalBag()
	var ironCount int64
	for _, item := range bag.Items {
		if item.UUid == repo.DuanTieUUid {
			ironCount = item.Count
		}
	}
	if ironCount <= 0 {
		msg += " 背包内锻铁数量不足... "
		return msg
	}

	for ind, item := range bag.Items {
		if item.UUid == uuid && item.Type == model.ItemTypeEquip {
			existed = true
			if item.Status == model.ItemStatusEquip { // 必须卸下才能锻造
				msg += " 装备中的法器无法锻造 请先卸下... "
				return msg
			} else {
				equipSelected := bag.Items[ind].EquipInfo
				if equipSelected == nil {
					msg += " 获取装备失败... "
					return msg
				}
				// 检查身上锻铁数量
				if ironCount < equipSelected.Level {
					msg += " 背包内锻铁数量不足... 当前法器锻造需要 " + fmt.Sprint(equipSelected.Level) + " 块锻铁... "
					return msg
				}
				ironCount = ironCount - equipSelected.Level
				msg += "消耗锻铁 " + fmt.Sprint(equipSelected.Level) + " 块."
				forged = true

				successRate := 100 - (equipSelected.Level * 2)
				if rand.Int63n(100) < successRate {
					msg += " 锻造成功了... `"
					upgradedEquip := equipSelected.UpgradeEquip()
					bag.Items[ind].Description = upgradedEquip.Description
					bag.Items[ind].EquipInfo = &upgradedEquip
					msg += upgradedEquip.Name + "` 变强了! 当前 " + fmt.Sprintf("%d 级! ", upgradedEquip.Level)
				} else {
					msg += " 锻造失败... 锻铁破碎了.. `"
					msg += equipSelected.Name + "` 没有产生变化! "
				}
			}
		}
	}
	if forged {
		for ind, item := range bag.Items {
			if item.UUid == repo.DuanTieUUid {
				bag.Items[ind].Count = ironCount
				break
			}
		}
		updateLocalBag(bag)
	}

	if !existed {
		msg += " 不存在这样的装备... "
		return msg
	}
	return msg
}

// DestroyEquip 摧毁装备
func (s *EquipService) DestroyEquip(uuid int64) string {
	var msg string
	var existed bool
	// 检查身上是否存在装备 uuid
	bag := getLocalBag()

	for ind, item := range bag.Items {
		if item.UUid == uuid && item.Type == model.ItemTypeEquip {
			existed = true
			if item.Status == model.ItemStatusEquip { // 必须卸下才能摧毁
				msg += " 装备中的法器无法摧毁 请先卸下... "
				return msg
			} else {
				bag.Items[ind].Count = 0
				msg += " 摧毁了 `" + item.Name + "` "
				bag = addBagItemByUUid(bag, repo.DuanTieUUid, utils.Max(item.EquipInfo.Level/2, 1))
				msg += "获得了 " + fmt.Sprint(utils.Max(item.EquipInfo.Level/2, 1)) + " 块`锻铁`材料..."
				break
			}
		}
	}

	if !existed {
		msg += " 不存在这样的装备... "
		return msg
	} else {
		updateLocalBag(bag)
	}

	return msg
}
