package types

type AllocateReq struct {
	Stat string `json:"stat"`
}

type AllocateResp struct {
	Msg string `json:"msg"`
}

type CultivationReq struct {
	Username string `json:"username,optional"`
}

type CultivationResp struct {
	NextTime string `json:"nextTime"` // 下一次可以修炼的时间
}

type DailyReq struct {
	Username string `json:"username,optional"`
}

type DailyResp struct {
	Msg string `json:"msg"`
}

type GetUserInfoReq struct {
	Username string `json:"username,optional"`
}

type GetUserInfoResp struct {
	Username            string `json:"username"`
	Attack              int64  `json:"attack"`              // 攻击
	Defense             int64  `json:"defense"`             // 防御
	Hp                  int64  `json:"hp"`                  // 生命
	HpLimit             int64  `json:"hpLimit"`             // 最大生命
	Exp                 int64  `json:"exp"`                 // 经验
	Cultivation         string `json:"level"`               // 修为, 与等级挂钩
	Gold                int64  `json:"gold"`                // 金币
	Potential           int64  `json:"potential"`           // 潜能
	NextCultivationTime string `json:"nextCultivationTime"` // 下次修炼和治疗时间
}

type HealReq struct {
	Username string `json:"username,optional"`
}

type HealResp struct {
	Msg string `json:"msg"`
}

type JoinChallengeReq struct {
	Username    string `json:"username,optional"`
	ChallengeId int64  `json:"challengeId"`
}

type JoinChallengeResp struct {
	Msg string `json:"msg"`
	Log string `json:"log"`
}

type ListChallengeItem struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	MonsterList []Monster `json:"monster"`
	Reward      string    `json:"reward"`
}

type ListChallengeReq struct {
	Username string `json:"username,optional"`
}

type ListChallengeResp struct {
	List []ListChallengeItem `json:"list"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Monster struct {
	Name        string `json:"name"`
	Hp          int64  `json:"hp"`
	HpLimit     int64  `json:"hpLimit"` // 最大生命
	Attack      int64  `json:"attack"`  // 攻击
	Defense     int64  `json:"defense"` // 防御
	Cultivation string `json:"level"`   // 修为, 与等级挂钩
	Speed       int64  `json:"speed"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Id    int64  `json:"id"`
	Token string `json:"token"`
}
