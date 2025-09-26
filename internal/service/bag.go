package service

import (
	"encoding/json"
	"fmt"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/utils"
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

func cleanBagCountZeroItem(bag *model.Bag) *model.Bag {
	var retItems []*model.Item
	for _, item := range bag.Items {
		if item.Count > 0 {
			retItems = append(retItems, item)
		}
	}
	bag.Items = retItems
	return bag
}

func addBagItem(itemAdd *model.Item) {
	if itemAdd.UUid == 0 { // 未知物品
		return
	}
	bag := getLocalBag()
	for i, item := range bag.Items {
		if item.UUid == itemAdd.UUid {
			bag.Items[i].Count += itemAdd.Count
			updateLocalBag(bag)
			return
		}
	}

	bag.Items = append(bag.Items, itemAdd)
	updateLocalBag(bag)
}

func (s *BagService) GetBag() string {
	bag := getLocalBag()
	bagBytes, _ := json.Marshal(bag)
	return string(bagBytes)
}

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
				useLog, used := s.useItem(item.UUid)
				msg += useLog
				if used {
					// 使用成功则减掉数量 1
					bag.Items[ind].Count -= 1
					cleanBagCountZeroItem(bag)
				}
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

// useItem 使用指定ID的物品，根据物品ID对用户属性进行修改，并可能触发随机事件。
// 参数:
//   - id: 物品ID，决定使用哪种物品及其效果
//
// 返回值:
//   - string: 操作结果的消息描述
//   - bool: 操作是否成功执行
func (s *BagService) useItem(id int64) (string, bool) {
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
