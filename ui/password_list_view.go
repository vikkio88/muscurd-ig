package ui

import (
	"fmt"
	c "muscurdig/context"
	"muscurdig/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

func GetPasswordListView(ctx *c.AppContext) *fyne.Container {
	passwordCount := ctx.Db.GetPasswordCount()
	baseContentString := fmt.Sprintf("%d Password Entries", passwordCount)
	content := getBaseContent(baseContentString)
	dataList := binding.NewUntypedList()
	listAllBtn := widget.NewButtonWithIcon("", theme.ListIcon(), func() {
		populateList(content, ctx.Db.GetAllPasswords(), dataList, ctx)

	})
	searchEntry := widget.NewEntry()
	searchBtn := widget.NewButtonWithIcon("Search", theme.SearchIcon(), func() {
		searchTerm := searchEntry.Text
		populateList(content, ctx.Db.FilterPasswords(searchTerm), dataList, ctx)
	})
	searchBtn.Disable()
	searchEntry.PlaceHolder = "Website or username..."
	searchEntry.OnChanged = func(s string) {
		searchBtn.Disable()
		if len(s) > 2 {
			searchBtn.Enable()
		}

		if len(s) == 0 {
			content.RemoveAll()
			content.Add(getBaseContent(baseContentString))
		}
	}

	return container.NewMax(
		container.NewBorder(
			container.NewBorder(nil, nil, nil,
				container.NewHBox(searchBtn, listAllBtn),
				searchEntry,
			),
			container.NewBorder(nil, nil, nil,
				widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
					ctx.NavigateTo(c.AddUpdate)
				},
				)),
			nil, nil,
			content,
		))
}

// DataList Items
// Renders a row for the result list
func getPasswordEntryRow() fyne.CanvasObject {
	return container.NewGridWithColumns(3,
		widget.NewLabel(""),
		container.NewMax(),
		rightAligned(
			container.NewHBox(
				widget.NewButtonWithIcon("", theme.ZoomInIcon(), func() {}),
				widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {}),
			),
		),
	)
}

// DataList Items
// gives you items in the rendered row list
func getListWidgetsFromContainer(co *fyne.CanvasObject) (*widget.Label, *widget.Button, *widget.Button) {
	container := (*co).(*fyne.Container)
	label := container.Objects[0].(*widget.Label)
	btnBoxWrapper := container.Objects[2].(*fyne.Container)
	btnBox := btnBoxWrapper.Objects[0].(*fyne.Container)
	btnView := btnBox.Objects[0].(*widget.Button)
	btnDelete := btnBox.Objects[1].(*widget.Button)

	return label, btnView, btnDelete
}

// DataList Items
// Converts untyped list item to PasswordEntry
func getPasswordEntryFromDI(item binding.DataItem) models.PasswordEntry {
	v, _ := item.(binding.Untyped).Get()
	return v.(models.PasswordEntry)
}

// populates the result list
func populateList(content *fyne.Container, passwords []models.PasswordEntry, dataList binding.UntypedList, ctx *c.AppContext) {
	// Resetting the results
	list, _ := dataList.Get()
	list = list[:0]
	dataList.Set(list)
	//
	for _, p := range passwords {
		dataList.Append(p)
	}

	content.RemoveAll()
	if dataList.Length() < 1 {
		content.Add(getBaseContent("No Password..."))
		return
	}

	content.Add(
		widget.NewListWithData(
			dataList,
			getPasswordEntryRow,
			func(di binding.DataItem, co fyne.CanvasObject) {
				label, btnView, btnDelete := getListWidgetsFromContainer(&co)
				pe := getPasswordEntryFromDI(di)
				label.SetText(pe.Website)

				btnView.OnTapped = func() {
					ctx.NavigateToWithParam(c.Details, pe.Id)
				}

				btnDelete.OnTapped = func() {
					dialog.ShowConfirm(
						fmt.Sprintf("Deleting password for \"%s\"", pe.Website),
						"Are you sure?",
						func(b bool) {
							if b {
								lst, _ := dataList.Get()

								idx := slices.IndexFunc(lst, func(pei interface{}) bool {
									peii, ok := pei.(models.PasswordEntry)
									return ok && pe.Id == peii.Id
								})

								// this should never happen
								if idx < 0 {
									return
								}

								lst = append(lst[:idx], lst[idx+1:]...)
								dataList.Set(lst)
								ctx.Db.DeletePasswordEntry(pe.Id)

							}
						},
						ctx.GetWindow(),
					)
				}
			},
		),
	)

}

// gives us the basic centered text result
func getBaseContent(text string) *fyne.Container {
	return container.NewMax(container.NewCenter(widget.NewLabel(text)))
}
