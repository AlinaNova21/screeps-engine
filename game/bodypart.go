package game

type BodyPartType string

const (
	Attack       BodyPartType = "attack"
	Carry        BodyPartType = "carry"
	Claim        BodyPartType = "claim"
	Heal         BodyPartType = "heal"
	Move         BodyPartType = "move"
	RangedAttack BodyPartType = "ranged_attack"
	Tough        BodyPartType = "tough"
	Work         BodyPartType = "work"
)

type BodyPart struct {
	Type    BodyPartType
	Hits    int8
	HitsMax int8
	Boost   string
}
