package land

type Pool struct {
	StringData map[string]string
	HashData   map[string]HashData
	ListData   map[string][]string
	SetData    map[string][]string
	ZsetData   map[string]ZsetData
}

type HashData struct {
	Key string
	Val string
}

type ZsetData struct {
	Key   string
	Score string
}
