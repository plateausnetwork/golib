package sendmail

type Templates map[string]string

func (t Templates) CheckIDExists(id string) bool {
	if _, ok := t[id]; ok {
		return true
	}
	return false
}

func (t Templates) GetID(key string) string {
	if id, ok := t[key]; ok {
		return id
	}
	return ""
}

type DynamicData map[string]interface{}
