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
}
