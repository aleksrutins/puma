package packagefile

type Target struct {
	Name string `hcl:"name,label"`
	Src  string `hcl:"src"`
	Main string `hcl:"main"`
}

type Package struct {
	Name      string   `hcl:"name"`
	Libraries []Target `hcl:"library,block"`
	Bundles   []Target `hcl:"bundle,block"`
	Packages  []string `hcl:"packages"`
}

func (p *Package) GetLibrary(name string) *Target {
	for _, lib := range p.Libraries {
		if lib.Name == name {
			return &lib
		}
	}
	return nil
}

func (p *Package) GetBundle(name string) *Target {
	for _, bundle := range p.Bundles {
		if bundle.Name == name {
			return &bundle
		}
	}
	return nil
}
