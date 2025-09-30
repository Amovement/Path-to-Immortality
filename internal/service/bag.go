package service

import (
	"encoding/json"
	"fmt"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/utils"
	"time"
)

type BagService struct {
}

func NewBagService() *BagService {
	return &BagService{}
}

func getLocalBag() *model.Bag {
	bag := model.NewBag()
	bagString, existed := utils.GetStorage(model.BagStorageKey)
	if existed {
		if IsProd() {
			bagString, _ = utils.Decrypt(bagString)
		}
		err := json.Unmarshal([]byte(bagString), &bag)
		if err != nil {
			fmt.Printf("[ERROR] %+v\n", err)
			bag = model.NewBag()
		}
	}
	return bag
}

func updateLocalBag(bag *model.Bag) {
	cleanBagCountZeroItem(bag)
	bagBytes, _ := json.Marshal(bag)
	bagString := string(bagBytes)
	if IsProd() {
		bagStringEncrypted, err := utils.Encrypt(bagString)
		if err != nil {
			fmt.Printf("[ERROR] %+v\n", err)
			return
		}
		bagString = bagStringEncrypted
	}
	utils.SetStorage(model.BagStorageKey, bagString)
}

// cleanBagCountZeroItem 清理背包中数量为0的物品项
//
//		参数:
//	  bag - 需要清理的背包对象指针
//
// 返回值:
//
//	*model.Bag - 清理后的背包对象指针
func cleanBagCountZeroItem(bag *model.Bag) *model.Bag {
	// 过滤出数量大于0的物品
	var retItems []*model.Item
	for _, item := range bag.Items {
		if item.Count > 0 {
			retItems = append(retItems, item)
		}
	}
	bag.Items = retItems
	return bag
}

// addBagItem 向本地背包中添加物品
//
//		参数:
//	  itemAdd: 要添加的物品指针
func addBagItem(bag *model.Bag, itemAdd *model.Item) *model.Bag {
	// 检查物品UUID是否有效，无效则直接返回
	if itemAdd.UUid == 0 { // 未知物品
		return bag
	}

	// 遍历背包中的物品，查找是否已存在相同UUID的物品
	for i, item := range bag.Items {
		// 如果找到相同UUID的物品，则增加数量并更新背包
		if item.UUid == itemAdd.UUid {
			bag.Items[i].Count += itemAdd.Count
			updateLocalBag(bag)
			return bag
		}
	}

	// 如果没有找到相同UUID的物品，则将新物品添加到背包中
	bag.Items = append(bag.Items, itemAdd)
	updateLocalBag(bag)
	return bag
}

func (s *BagService) GetBag() string {
	bag := getLocalBag()
	bagBytes, _ := json.Marshal(bag)
	return string(bagBytes)
}

// UseBagItem 使用背包中的物品
// id: 要使用的物品ID
// 返回值: 操作结果的提示信息
func (s *BagService) UseBagItem(id int64) string {
	key := model.UserOperatorLock // 锁
	if _, ok := CacheRedis.Load(key); ok {
		return "请求过于频繁"
	}
	CacheRedis.Store(key, struct{}{})
	defer CacheRedis.Delete(key)

	var msg string

	bag := getLocalBag()
	// 检查是否有该物品
	var existed bool
	for ind, item := range bag.Items {
		if item.UUid == id && item.Count > 0 {
			existed = true
			if item.Type == model.ItemTypeConsume { // 消耗品
				useLog, used := s.useConsumeItem(item.UUid)
				msg += useLog
				if used {
					// 使用成功则减掉数量 1
					bag.Items[ind].Count -= 1
					bag = cleanBagCountZeroItem(bag)
				}
			} else if item.Type == model.ItemTypeEquip { // 装备
				userLog := equipItem(item.UUid, bag)
				msg += userLog
			} else if item.Type == model.ItemTypeMaterial { // 材料
				useLog, _ := s.userMaterialItem(item.UUid, bag)
				msg += useLog
			}
			break
		}
	}

	if !existed {
		msg += "物品不存在"
	}
	updateLocalBag(bag)
	return msg
}

// useConsumeItem 使用指定ID的物品(消耗品类型)，根据物品ID对用户属性进行修改，并可能触发随机事件。
//
//		参数:
//	  - id: 物品ID，决定使用哪种物品及其效果
//
// 返回值:
//   - string: 操作结果的消息描述
//   - bool: 操作是否成功执行
func (s *BagService) useConsumeItem(id int64) (string, bool) {
	var msg string
	user := getLocalUser()

	// 检查当前用户是否已达到该物品的使用限制（耐药性检查）
	ok := s.checkGoodsLimit(user, id)
	if !ok {
		msg = msg + " 当前境界服用此物已经达到瓶颈.还是提升了境界再来吧..."
		return msg, false
	}

	// 随机生成一个数值用于判断是否发生负面效果
	randInt := utils.GetRandomInt64(1, user.Level+100)
	badThingHappened := false

	// 根据物品ID执行不同的逻辑处理
	switch id {
	case 1:
		msg = msg + "你的体魄上限小幅度增强了."
		user.HpLimit += 5
		if randInt <= 70 {
			badThingHappened = true
		}
	case 2:
		msg = msg + "你的攻击力小幅度提升了."
		user.Attack += 1
		if randInt <= 70 {
			badThingHappened = true
		}
	case 3:
		msg = msg + "你的防御力小幅度提升了."
		user.Defense += 1
		if randInt <= 70 {
			badThingHappened = true
		}
	case 4:
		msg = msg + "你的速度小幅度提升了."
		user.Speed += 1
		if randInt <= 70 {
			badThingHappened = true
		}
	case 5:
		msg = msg + " 你陷入了昏迷."
		if randInt == 50 {
			user.Exp = user.Level*10 - 5
			msg = msg + " 你好像明白了什么，突破只差临门一脚. "
		} else {
			msg = msg + " 什么都没有发生."
		}
	case 6:
		user.Exp += 10
		msg = msg + " 获得经验 10 点"
	case 7:
		user.Hp += 15
		if user.Hp > user.HpLimit {
			user.Hp = user.HpLimit
		}
		msg = msg + " 获得生命值 15 点"
	case 8:
		gold := utils.GetRandomInt64(1, utils.Max(user.Level, 150))
		user.Gold += gold
		msg = msg + "好运连连 获得金币 " + fmt.Sprint(gold) + "枚. "
	case 9:
		msg = msg + "气息更加绵长了."
		user.HpLimit += 10
	case 10:
		msg = msg + "攻击变得更为凌厉了."
		user.Attack += 2
	case 11:
		msg = msg + "叠甲, 过."
		user.Defense += 2
	case 12:
		msg = msg + "脚下生风!"
		user.Speed += 2
	case 13:
		msg = msg + "体内的力量涌出来了. 潜能 +1 ."
		user.Potential += 1
		if user.RestartCount > 0 {
			user.Potential = user.Potential + user.RestartCount
			msg += "体内的另外一个灵魂正在回应你, 你好像想起来了很多东西, 额外获得了 " + fmt.Sprint(user.RestartCount) + " 点潜能"
		}
	}

	// 如果触发了负面效果，则执行惩罚逻辑
	if badThingHappened {
		msg = msg + " 身体内的灵力暴走了!!!"
		user.Hp -= user.HpLimit / 10
		if user.Hp <= 0 {
			user.Hp = 1
		}
		msg = msg + " 损失 " + fmt.Sprint(user.HpLimit/10) + " 点生命值. "
		randInt = utils.GetRandomInt64(1, 5)
		if randInt == 1 {
			user.HpLimit -= 6
			if user.HpLimit <= 10 {
				user.HpLimit = 10
			}
			msg = msg + " 体魄上限扣除了. "
		} else if randInt == 2 {
			user.Attack -= 2
			if user.Attack < 0 {
				user.Attack = 0
			}
			msg = msg + " 攻击力下降了. "
		} else if randInt == 3 {
			user.Defense -= 2
			if user.Defense < 0 {
				user.Defense = 0
			}
			msg = msg + " 防御力下降了. "
		} else if randInt == 4 {
			user.Speed -= 2
			if user.Speed < 0 {
				user.Speed = 0
			}
			msg = msg + " 速度下降了. "
		}
	}

	// 更新用户信息到数据库或缓存
	updateUserInfo(user)
	return msg, true
}

func (s *BagService) checkGoodsLimit(user *model.User, goodsId int64) bool {
	switch goodsId {
	case 1:
		if user.HpLimit+5 > user.Level*3*5+10 {
			return false
		}
	case 2:
		if user.Attack+1 > user.Level*3 {
			return false
		}
	case 3:
		if user.Defense+1 > user.Level*3 {
			return false
		}
	case 4:
		if user.Speed+1 > user.Level*3 {
			return false
		}
	}

	if goodsId >= 9 && goodsId <= 12 {
		// 高级的药 需要检查总属性不超过等级的12倍
		totalStat := user.Attack + user.Defense + user.Speed + user.HpLimit/5 + user.Potential // 玩家的总属性
		if totalStat >= user.Level*12+user.RestartCount {
			return false
		}
	}

	return true
}

func (s *BagService) userMaterialItem(uuid int64, bag *model.Bag) (string, bool) {
	var (
		msg  string
		used bool
	)

	if uuid == 14 || uuid == 15 {
		if uuid == 14 {
			if !s.checkBagHasItem(bag, 15) {
				msg += "你好像没有`精魄`, 合成法器失败..."
				return msg, used
			}
		} else {
			if !s.checkBagHasItem(bag, 14) {
				msg += "你好像没有`玄晶`, 合成法器失败..."
				return msg, used
			}
		}
		bag = s.reduceItem(bag, 14)
		bag = s.reduceItem(bag, 15)
		// 合成装备
		equip := model.RandomEquip(0, 3, -1)
		if bag.RandomUUid == 0 {
			bag.RandomUUid = time.Now().Unix()
		}
		bag = addBagItem(bag, &model.Item{
			UUid:        bag.RandomUUid,
			Name:        equip.Name,
			Description: equip.GenerateDescription(),
			Count:       1,
			Type:        1,
			EquipInfo:   equip,
		})
		bag.RandomUUid = bag.RandomUUid + 1
		msg += fmt.Sprintf("法器打造成功: %s", equip.GenerateDescription())
		updateLocalBag(bag)

	} else {
		msg += " 这是仍未被发现的材料... 你还不知道怎么使用它"
	}

	return msg, used
}

// reduceItem 减少背包中指定物品的数量
// 该函数会查找背包中UUID匹配的物品，将其数量减1，如果物品数量变为0则清理该物品
//
// 参数:
//
//	bag - 背包对象指针，包含物品列表
//	uuid - 要减少的物品的唯一标识符
//
// 返回值:
//
//	返回更新后的背包对象指针
func (s *BagService) reduceItem(bag *model.Bag, uuid int64) *model.Bag {
	// 遍历背包中的所有物品，查找匹配UUID的物品
	for i, item := range bag.Items {
		if item.UUid == uuid {
			// 如果物品数量大于0，则减少1个
			if item.Count > 0 {
				bag.Items[i].Count--
				// 清理背包中数量为0的物品
				bag = cleanBagCountZeroItem(bag)
			}
			return bag
		}
	}
	return bag
}

// checkBagHasItem 检查背包中是否包含指定UUID且数量大于0的物品
//
//		参数:
//	  - bag: 要检查的背包对象
//	  - uuid: 要查找的物品UUID
//
// 返回值:
//   - bool: 如果找到UUID匹配且数量大于0的物品则返回true，否则返回false
func (s *BagService) checkBagHasItem(bag *model.Bag, uuid int64) bool {
	// 遍历背包中的所有物品
	for _, item := range bag.Items {
		// 检查物品UUID是否匹配且数量大于0
		if item.UUid == uuid && item.Count > 0 {
			return true
		}
	}
	return false
}
