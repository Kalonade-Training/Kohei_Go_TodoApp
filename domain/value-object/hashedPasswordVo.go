package value_object

type HashedPassword struct {
	value string
}

func NewHashedPassword(value string) HashedPassword {
	return HashedPassword{value: value}
}

func (hp HashedPassword) Value() string {
	return hp.value
}