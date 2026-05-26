package deepseek

const (
	BaseURL          = "https://api.deepseek.com"
	DefaultTestModel = "deepseek-v4-flash"
)

type Model struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	Created     int64  `json:"created"`
	OwnedBy     string `json:"owned_by"`
	Type        string `json:"type"`
	DisplayName string `json:"display_name"`
}

var DefaultModels = []Model{
	{ID: "deepseek-chat", Object: "model", Created: 1704067200, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek Chat"},
	{ID: "deepseek-reasoner", Object: "model", Created: 1704067200, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek Reasoner"},
	{ID: "deepseek-v3-2-251201", Object: "model", Created: 1764547200, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek V3.2"},
	{ID: "deepseek-v4-pro", Object: "model", Created: 1764547200, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek V4 Pro"},
	{ID: "deepseek-v4-flash", Object: "model", Created: 1764547200, OwnedBy: "deepseek", Type: "model", DisplayName: "DeepSeek V4 Flash"},
}

func DefaultModelIDs() []string {
	ids := make([]string, len(DefaultModels))
	for i, m := range DefaultModels {
		ids[i] = m.ID
	}
	return ids
}
