package service

import (
	"encoding/json"
	"fmt"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/model"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/repo"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/types"
	"github.com/Amovement/Path-to-Immortality-WASM/internal/utils"
	"log"
	"time"
)

type MythicPlusService struct {
}

func NewMythicPlusService() *MythicPlusService {
	return &MythicPlusService{}
}

func (s *MythicPlusService) LowerTheMythicPlus() string {
	key := fmt.Sprintf("stat:lock") // 角色属性锁
	if _, ok := CacheRedis.Load(key); ok {
		return "请求过于频繁"
	}
	CacheRedis.Store(key, struct{}{})
	defer CacheRedis.Delete(key)

	var msg string
	user := getLocalUser()
	nowTime := time.Now().Unix()
	if nowTime < user.Mythic.NextOpenTime {
		msg = "请稍后再来...秘境下一次可降级时间为 " + time.Unix(user.Mythic.NextOpenTime, 0).Format("2006-01-02 15:04:05")
		return msg
	}

	nowLevel := utils.Max(user.Mythic.Level-1, 0)
	user.Mythic.Level = nowLevel
	nextOpenTime := time.Now().Add(time.Minute * 15).Unix() // 15分钟
	user.Mythic.NextOpenTime = nextOpenTime

	monsterCount := utils.Max(nowLevel/10, 1)
	var monsters []*model.Monster
	for i := 0; i < int(monsterCount); i++ {
		monsters = append(monsters, model.GenerateRandomMythicMonster(nowLevel))
	}
	user.Mythic.Monsters = monsters

	updateUserInfo(user)
	msg += "秘境已刷新... 当前秘境等级: " + fmt.Sprintf("%d", nowLevel)
	return msg
}

func (s *MythicPlusService) GetMythicInfo() []byte {
	user := getLocalUser()
	if len(user.Mythic.Monsters) == 0 {
		monsterCount := utils.Max(user.Mythic.Level/10, 1)
		var monsters []*model.Monster
		for i := 0; i < int(monsterCount); i++ {
			monsters = append(monsters, model.GenerateRandomMythicMonster(user.Mythic.Level))
		}
		user.Mythic.Monsters = monsters
	}
	updateUserInfo(user)
	user.Mythic.Description = s.generateDescription(user.Mythic.Monsters)
	bytes, _ := json.Marshal(user.Mythic)
	return bytes
}

// generateDescription 获取词缀说明
func (s *MythicPlusService) generateDescription(monsters []*model.Monster) string {
	special := make(map[string]struct{})
	var description string
	for _, monster := range monsters {
		for _, specialName := range monster.Special {
			if _, ok := special[specialName]; ok {
				continue
			}
			special[specialName] = struct{}{}

			description += "<li>" + fmt.Sprintf("%s - %s", specialName, model.SpecialsDescription[specialName]) + "</li>"
		}
	}
	log.Println(description)
	if description == "" {
		description += "<li>本层无特殊词缀</li>"
	}
	return description
}

func (s *MythicPlusService) JoinMythic() (string, string) {
	key := fmt.Sprintf("stat:lock") // 角色属性锁
	if _, ok := CacheRedis.Load(key); ok {
		return "请求过于频繁", ""
	}
	CacheRedis.Store(key, struct{}{})
	defer CacheRedis.Delete(key)

	var msg string
	user := getLocalUser()
	if len(user.Mythic.Monsters) == 0 {
		msg += "请先刷新秘境..."
		return msg, msg
	}

	var monsterList []types.Monster
	for _, monster := range user.Mythic.Monsters {
		monsterList = append(monsterList, types.Monster{
			Name:        monster.Name,
			Hp:          monster.Hp,
			HpLimit:     monster.HpLimit,
			Attack:      monster.Attack,
			Defense:     monster.Defense,
			Speed:       monster.Speed,
			Special:     monster.Special,
			Cultivation: monster.Cultivation,
		})
	}

	msg, fightLog := FightCore(user, monsterList)
	if msg == fightWin {
		equipAttr := getUserEquipAttributes()
		if model.CheckHasSpecial(equipAttr.Special, model.SpecialsBleed) {
			user.Hp -= (user.HpLimit * 5) / 100
			if user.Hp < 0 {
				user.Hp = 1
			}
			msg = msg + " `流血诅咒`扣除了" + fmt.Sprint((user.HpLimit*5)/100) + "点体魄. "
		}
		msg = msg + " 第 " + fmt.Sprint(user.Mythic.Level) + " 层秘境战斗胜利. "

		user.Gold += user.Mythic.Level * 100
		bag := getLocalBag()
		addBagItemByUUid(bag, repo.DuanTieUUid, (user.Mythic.Level/3)+1)
		updateLocalBag(bag)
		msg = msg + " 获得 " + fmt.Sprint(user.Mythic.Level*100) + " 金币. " + " 获得 " + fmt.Sprint((user.Mythic.Level/3)+1) + " 锻铁"

		// 更新秘境信息
		user.Mythic.Level++
		msg += " 秘境层数上升了, 当前秘境层数等级: " + fmt.Sprint(user.Mythic.Level) + " 层"
		monsterCount := utils.Max(user.Mythic.Level/10, 1)
		var monsters []*model.Monster
		for i := 0; i < int(monsterCount); i++ {
			monsters = append(monsters, model.GenerateRandomMythicMonster(user.Mythic.Level))
		}
		user.Mythic.Monsters = monsters

	} else if msg == fightLose {
		// 掉 10% 经验惩罚 + 损失身上 10% 的金币
		// 如果身上经验不足 10% 则直接掉级，所有属性减 1
		if user.Exp <= user.Level { // 升级经验是等级的 10 倍
			user.Level = utils.Max(user.Level-1, 0)
			user.Attack = utils.Max(user.Attack-1, 0)
			user.Defense = utils.Max(user.Defense-1, 0)
			user.Speed = utils.Max(user.Speed-1, 0)
			user.HpLimit = utils.Max(user.HpLimit-1, model.DefaultHp)
			msg = msg + " 经验不足, 您的境界跌落了(各项属性下降)！"
		} else {
			msg = msg + " 修为损失 " + fmt.Sprint(user.Exp/10) + ". "
			user.Exp = user.Exp - user.Exp/10
		}
		msg = msg + " 金币减少 " + fmt.Sprint(user.Gold/10) + ". "
		user.Gold = user.Gold - user.Gold/10

		user.Potential = utils.Max(user.Potential-3, 0) // 潜能也扣
		user.Hp = 1
	}

	nextOpenTime := time.Now().Add(time.Minute * 15).Unix() // 15分钟
	user.Mythic.NextOpenTime = nextOpenTime
	updateUserInfo(user)

	return msg, fightLog + "\n" + msg
}
