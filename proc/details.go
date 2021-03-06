package proc

import (
	"context"
	"encoding/json"
	"github.com/gotrino/fusion-rt-wasmjs/pkg/web/i18n"
	"github.com/gotrino/fusion/spec/app"
	"github.com/gotrino/fusion/spec/form"
	"github.com/gotrino/fusion/spec/http"
	"github.com/gotrino/fusion/spec/rest"
	"log"
)

type Details struct {
	Route string `route:"/devices/:id/details"`
	ID    string `path:"id"`
}

func (a *Details) Compose(ctx context.Context) app.Activity {
	return app.Activity{
		Title: "Details",
		Fragments: []app.Fragment{
			form.Form{
				Title:       "Stammdaten",
				Description: "Bearbeiten Sie hier die Stammdaten des Scripts",
				ResourceID:  a.ID,
				Repository: rest.Repository[ProcInfo]{
					Path: "/api/v1/procs",
				},
				CanCancel: true,
				CanWrite:  true,
				Fields: []form.Field{
					form.Text[ProcInfo]{
						Label:       i18n.Text(ctx, "Name"),
						Placeholder: i18n.Text(ctx, "Name of the script"),
						Description: i18n.Text(ctx, "A short but memorable name"),
						ToModel: func(src string, state ProcInfo) (ProcInfo, error) {
							if src == "" {
								return state, app.ValidationError{Message: i18n.Text(ctx, "Name must not be empty")}
							}

							state.Name = src
							return state, nil
						},
						FromModel: func(src ProcInfo) string {
							return src.Name
						},
					},

					form.Text[ProcInfo]{
						Label:       i18n.Text(ctx, "Description"),
						Placeholder: i18n.Text(ctx, "..."),
						Description: i18n.Text(ctx, "A more verbose description of the script."),
						ToModel: func(src string, state ProcInfo) (ProcInfo, error) {
							state.Description = src
							return state, nil
						},
						FromModel: func(src ProcInfo) string {
							return src.Description
						},
					},
				},
			},

			form.Form{
				Title:       "Beispiel",
				Description: "Beispiel für Ein- und Ausgabe-Formate für das Script",
				ResourceID:  a.ID + "/parameter",
				Repository: rest.Repository[ParamInfo]{
					Path: "/api/v1/procs",
				},
				Fields: []form.Field{
					form.Label[ParamInfo]{
						FromModel: func(src ParamInfo) string {
							return "Beispiel für die vom Script vorgegebenen Body Request-Parameter:"
						},
					},
					form.CodeEditor[ParamInfo]{
						ReadOnly: true,
						Lang:     "json",
						FromModel: func(src ParamInfo) string {
							buf, err := json.MarshalIndent(src.Example.Request, " ", " ")
							if err != nil {
								panic(err)
							}
							return string(buf)
						},
					},
					form.Label[ParamInfo]{
						Text: "Beispiel für die vom Script vorgegebene Response:",
					},
					form.CodeEditor[ParamInfo]{
						ReadOnly: true,
						Lang:     "json",
						FromModel: func(src ParamInfo) string {
							buf, err := json.MarshalIndent(src.Example.Response, " ", " ")
							if err != nil {
								panic(err)
							}
							return string(buf)
						},
					},
				},
			},

			form.Form{
				Title:       "Script",
				Description: "Der definierte Rechenkernel",
				CanWrite:    true,
				Repository: http.Repository[string]{
					OnSave: func(t string) error {
						log.Println("!!!!!!!!!!!!!!")
						_, err := http.Do(ctx, "PUT", http.URL(ctx, "/api/v1/procs", a.ID, "src"),
							http.Params{
								ContentType: "text/x-go",
								Body:        []byte(t),
							},
							http.StatusOK, http.StatusCreated, http.StatusAccepted,
						)

						log.Println("save done", err)

						return err
					},
					OnLoad: func(id string) (string, error) {
						buf, err := http.Do(ctx, "GET", http.URL(ctx, "/api/v1/procs", a.ID, "src"),
							http.Params{},
							http.StatusOK,
						)

						return string(buf), err
					},
				},
				Fields: []form.Field{
					form.Label[any]{
						Text: "Der Rechenkernel ist in MiEL zu schreiben. Einem Go-Dialekt, in dem die Standardbibliothek nicht zur Verfügung steht.",
					},
					form.CodeEditor[string]{
						Lang: "go",
						FromModel: func(src string) string {
							return src
						},
						ToModel: func(src, dst string) (string, error) {
							return src, nil
						},
					},
				},
			},
		},
	}
}
