package packagefile

import (
	"fmt"
	"strings"

	"github.com/aleksrutins/puma/log"
)

type fetchType int

const (
	git fetchType = iota
	local
)

type fetchPackage struct {
	Source     string
	SourceType fetchType
}

/**
 * Fetch all dependencies of a package, and extract them into ./_packages.
 */
func (p *Package) GetDependencies() error {
	var allPackages = []fetchPackage{}
	var spinner = log.StageSpinner{}
	for _, dependency := range p.Packages {
		if dependency == p.Name {
			return fmt.Errorf("Package %s depends on itself", p.Name)
		}
		var method fetchType = git
		var source string
		if strings.HasPrefix(dependency, "git:") || strings.HasPrefix(dependency, "http:") || strings.HasPrefix(dependency, "https:") {
			if strings.HasPrefix(dependency, "http:") {
				log.Warn(fmt.Sprintf("%s: HTTP is not recommended. Please use HTTPS instead.", dependency))
			}
			method = git
			source = dependency
		}
		if strings.HasPrefix(dependency, "local:") {
			method = local
			source = dependency[7:]
		}
		allPackages = append(allPackages, fetchPackage{
			Source:     source,
			SourceType: method,
		})
	}
	spinner.Total = len(allPackages)
	for _, packageToFetch := range allPackages {
		go func() error {
			spinner.Message = fmt.Sprintf("Fetching %s", packageToFetch.Source)
			spinner.Print()
			if packageToFetch.SourceType == git {
				if err := p.FetchGit(packageToFetch.Source); err != nil {
					return err
				}
			}
			if packageToFetch.SourceType == local {
				if err := p.FetchLocal(packageToFetch.Source); err != nil {
					return err
				}
			}
			spinner.Finished++
			spinner.Print()
			return nil
		}()
	}
	return nil
}

func (p *Package) FetchGit(source string) error {
	return fmt.Errorf("Not implemented")
}

func (p *Package) FetchLocal(source string) error {
	return fmt.Errorf("Not implemented")
}
