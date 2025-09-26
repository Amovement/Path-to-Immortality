package model

type User struct {
	Username            string          `json:"username"`
	Attack              int64           `json:"attack"`              // 攻击
	Defense             int64           `json:"defense"`             // 防御
	Hp                  int64           `json:"hp"`                  // 生命
	HpLimit             int64           `json:"hpLimit"`             // 生命上限
	Speed               int64           `json:"speed"`               // 速度
	Exp                 int64           `json:"exp"`                 // 经验
	Level               int64           `json:"level"`               // 等级
	Gold                int64           `json:"gold"`                // 金币
	Potential           int64           `json:"potential"`           // 潜能
	NextCultivationTime int64           `json:"nextCultivationTime"` // 治疗和修炼占用同一个时间
	PassedChallengeId   []uint          `json:"passedChallengeId"`   // 通过的挑战 ID
	PassedChallengeTime map[uint]string `json:"passedChallengeTime"` // 通过的 Challenge 时间
	RestartCount        int64           `json:"restartCount"`        // 成功轮回转生的次数
}

const (
	UserInfoStorageKey = "_Path_2_Immortality_User_"
	UserOperatorLock   = "stat:lock"

	DefaultHp = 10
)

func NewUser() *User {
	return &User{
		Username:            "无名大侠",
		Hp:                  DefaultHp,
		HpLimit:             DefaultHp,
		Attack:              0,
		Defense:             0,
		Speed:               0,
		Exp:                 0,
		Level:               0,
		Gold:                0,
		Potential:           0,
		PassedChallengeTime: make(map[uint]string),
	}
}
