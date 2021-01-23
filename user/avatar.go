package user

type Avatar struct {
	Id   uint32
	Name string
}

var AvatarValues = []Avatar{
	{0x1f43b, "bear"},
	{0x1f417, "boar"},
	{0x1f431, "cat"},
	{0x1f414, "chicken"},
	{0x1f42e, "cow"},
	{0x1f98c, "deer"},
	{0x1f436, "dog"},
	{0x1f432, "dragon"},
	{0x1f985, "eagle"},
	{0x1f98a, "fox"},
	{0x1f438, "frog"},
	{0x1f992, "giraffe"},
	{0x1f98d, "gorilla"},
	{0x1f439, "hamster"},
	{0x1f434, "horse"},
	{0x1f428, "koala"},
	{0x1f981, "lion"},
	{0x1f435, "monkey"},
	{0x1f42d, "mouse"},
	{0x1f43c, "panda"},
	{0x1f437, "pig"},
	{0x1f4a9, "poop"},
	{0x1f430, "rabbit"},
	{0x1f99d, "raccoon"},
	{0x1f98f, "rhinoceros"},
	{0x1f42f, "tiger"},
	{0x1f984, "unicorn"},
	{0x1f43a, "wolf"},
	{0x1f993, "zebra"},
}

var AvatarMap map[uint32]string

func AvatarName(id uint32) string {
	if AvatarMap == nil {
		AvatarMap = make(map[uint32]string)
	}

	if len(AvatarMap) == 0 {
		for _, value := range AvatarValues {
			AvatarMap[value.Id] = value.Name
		}
	}

	return AvatarMap[id]
}
