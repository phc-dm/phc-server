package util

type H map[string]interface{}

func (target H) Apply(sources ...H) H {
	for _, source := range sources {
		for k, v := range source {
			target[k] = v
		}
	}

	return target
}
