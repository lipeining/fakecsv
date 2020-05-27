package model

// Column 输入列属性 max min 根据输入另外转化为 datetime, int64, len(string)
type Column struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	// Fake     string `json:"fake"`
	Max      string `json:"max"`
	Min      string `json:"min"`
	Default  string `json:"default"`
	NotNull  bool   `json:"notNull"`
	Unique   bool   `json:"unique"`
	Autoincr bool   `json:"autoincr"`
}
