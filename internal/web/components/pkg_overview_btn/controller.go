package pkg_overview_btn

import (
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/glasskube/glasskube/api/v1alpha1"
	"github.com/glasskube/glasskube/pkg/client"
	"github.com/glasskube/glasskube/pkg/list"
)

const TemplateId = "pkg-overview-btn"

type pkgOverviewBtnInput struct {
	ButtonId    string
	Swap        string
	PackageName string
	Status      *client.PackageStatus
	Manifest    *v1alpha1.PackageManifest
}

func getButtonId(pkgName string) string {
	return fmt.Sprintf("%v-%v", TemplateId, pkgName)
}

func Render(w io.Writer, tmpl *template.Template, pkgName string, status *client.PackageStatus, manifest *v1alpha1.PackageManifest) {
	buttonId := getButtonId(pkgName)
	err := tmpl.ExecuteTemplate(w, TemplateId, &pkgOverviewBtnInput{
		ButtonId:    buttonId,
		Swap:        fmt.Sprintf("outerHTML:#%s", buttonId),
		PackageName: pkgName,
		Status:      status,
		Manifest:    manifest,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred rendering %v for %v: \n%v\n"+
			"This is most likely a BUG!", TemplateId, pkgName, err)
	}
}

func ForPkgOverviewBtn(pkgTeaser *list.PackageTeaserWithStatus) *pkgOverviewBtnInput {
	buttonId := getButtonId(pkgTeaser.PackageName)
	return &pkgOverviewBtnInput{
		ButtonId:    buttonId,
		Swap:        "",
		PackageName: pkgTeaser.PackageName,
		Status:      pkgTeaser.Status,
		Manifest:    pkgTeaser.InstalledManifest,
	}
}
