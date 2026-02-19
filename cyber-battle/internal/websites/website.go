package websites

type Website struct {
	Name            string
	URL             string
	Vulnerabilities []string
	Owner           string
}

func (w *Website) SimulateAttack(vuln string) bool {
	for _, v := range w.Vulnerabilities {
		if v == vuln {
			return true
		}
	}
	return false
}

func (w *Website) Info() string {
	return w.Name + " (" + w.URL + ") owned by " + w.Owner
}

func NewWebsite(name, url string, vulns []string, owner string) *Website {
	return &Website{
		Name:            name,
		URL:             url,
		Vulnerabilities: vulns,
		Owner:           owner,
	}
}
