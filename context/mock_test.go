package context_test

import (
	"muscurdig/libs"
	"muscurdig/models"

	"fyne.io/fyne/v2"
)

type WindowMock struct{}

// Title returns the current window title.
// This is typically displayed in the window decorations.
func (w WindowMock) Title() string {
	panic("not implemented") // TODO: Implement
}

// SetTitle updates the current title of the window.
func (w WindowMock) SetTitle(_ string) {
	panic("not implemented") // TODO: Implement
}

// FullScreen returns whether or not this window is currently full screen.
func (w WindowMock) FullScreen() bool {
	panic("not implemented") // TODO: Implement
}

// SetFullScreen changes the requested fullScreen property
// true for a fullScreen window and false to unset this.
func (w WindowMock) SetFullScreen(_ bool) {
	panic("not implemented") // TODO: Implement
}

// Resize this window to the requested content size.
// The result may not be exactly as desired due to various desktop or
// platform constraints.
func (w WindowMock) Resize(_ fyne.Size) {
	panic("not implemented") // TODO: Implement
}

// RequestFocus attempts to raise and focus this window.
// This should only be called when you are sure the user would want this window
// to steal focus from any current focused window.
func (w WindowMock) RequestFocus() {
	panic("not implemented") // TODO: Implement
}

// FixedSize returns whether or not this window should disable resizing.
func (w WindowMock) FixedSize() bool {
	panic("not implemented") // TODO: Implement
}

// SetFixedSize sets a hint that states whether the window should be a fixed
// size or allow resizing.
func (w WindowMock) SetFixedSize(_ bool) {
	panic("not implemented") // TODO: Implement
}

// CenterOnScreen places a window at the center of the monitor
// the Window object is currently positioned on.
func (w WindowMock) CenterOnScreen() {
	panic("not implemented") // TODO: Implement
}

// Padded, normally true, states whether the window should have inner
// padding so that components do not touch the window edge.
func (w WindowMock) Padded() bool {
	panic("not implemented") // TODO: Implement
}

// SetPadded allows applications to specify that a window should have
// no inner padding. Useful for fullscreen or graphic based applications.
func (w WindowMock) SetPadded(_ bool) {
	panic("not implemented") // TODO: Implement
}

// Icon returns the window icon, this is used in various ways
// depending on operating system.
// Most commonly this is displayed on the window border or task switcher.
func (w WindowMock) Icon() fyne.Resource {
	panic("not implemented") // TODO: Implement
}

// SetIcon sets the icon resource used for this window.
// If none is set should return the application icon.
func (w WindowMock) SetIcon(_ fyne.Resource) {
	panic("not implemented") // TODO: Implement
}

// SetMaster indicates that closing this window should exit the app
func (w WindowMock) SetMaster() {
	panic("not implemented") // TODO: Implement
}

// MainMenu gets the content of the window's top level menu.
func (w WindowMock) MainMenu() *fyne.MainMenu {
	panic("not implemented") // TODO: Implement
}

// SetMainMenu adds a top level menu to this window.
// The way this is rendered will depend on the loaded driver.
func (w WindowMock) SetMainMenu(_ *fyne.MainMenu) {
	panic("not implemented") // TODO: Implement
}

// SetOnClosed sets a function that runs when the window is closed.
func (w WindowMock) SetOnClosed(_ func()) {
	panic("not implemented") // TODO: Implement
}

// SetCloseIntercept sets a function that runs instead of closing if defined.
// Close() should be called explicitly in the interceptor to close the window.
//
// Since: 1.4
func (w WindowMock) SetCloseIntercept(_ func()) {
	panic("not implemented") // TODO: Implement
}

// Show the window on screen.
func (w WindowMock) Show() {
	panic("not implemented") // TODO: Implement
}

// Hide the window from the user.
// This will not destroy the window or cause the app to exit.
func (w WindowMock) Hide() {
	panic("not implemented") // TODO: Implement
}

// Close the window.
// If it is he "master" window the app will Quit.
// If it is the only open window and no menu is set via [desktop.App]
// SetSystemTrayMenu the app will also Quit.
func (w WindowMock) Close() {
	panic("not implemented") // TODO: Implement
}

// ShowAndRun is a shortcut to show the window and then run the application.
// This should be called near the end of a main() function as it will block.
func (w WindowMock) ShowAndRun() {
	panic("not implemented") // TODO: Implement
}

// Content returns the content of this window.
func (w WindowMock) Content() fyne.CanvasObject {
	panic("not implemented") // TODO: Implement
}

// SetContent sets the content of this window.
func (w WindowMock) SetContent(_ fyne.CanvasObject) {
	panic("not implemented") // TODO: Implement
}

// Canvas returns the canvas context to render in the window.
// This can be useful to set a key handler for the window, for example.
func (w WindowMock) Canvas() fyne.Canvas {
	panic("not implemented") // TODO: Implement
}

// Clipboard returns the system clipboard
func (w WindowMock) Clipboard() fyne.Clipboard {
	panic("not implemented") // TODO: Implement
}

type MockDb struct{}

func (m MockDb) Close() error {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) Drop() {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) SaveMasterPassword(mp models.MasterPassword) (models.MasterPassword, error) {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) GetMasterPassword() (models.MasterPassword, error) {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) GetCryptoInstance() *libs.Crypto {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) InsertPasswordEntry(password models.PasswordEntry) error {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) GetPasswordCount() int {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) GetPasswordById(id string) models.PasswordEntry {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) DeletePasswordEntry(id string) {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) UpdatePasswordEntry(pe models.PasswordEntry) error {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) FilterPasswords(search string) []models.PasswordEntry {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) GetAllPasswords() []models.PasswordEntry {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) GenerateDump(baseFolder string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (m MockDb) ImportDump(password string, dumpFileLocation string) error {
	panic("not implemented") // TODO: Implement
}
