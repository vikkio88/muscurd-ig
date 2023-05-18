# muscurd-ig
Portable multiplatform password manager.

Already done it in C# and Avalonia in [here](https://github.com/vikkio88/muscurd-i), attempting to do it in Golang and [Fyne](https://fyne.io/).

## Roadmap
- [x] TDD all entities
- [x] Design UI with Password Gating
- [x] Use [Clover](https://github.com/ostafen/clover) to persist data
- [x] Implement Search in List / Show All
- [ ] Add confirm modal to delete
- [ ] Implement NukeDb action to clean everything
- [x] Add Bind to List View to handle delete better (or remove Delete and move it to the details)
- [ ] Fix jumping error message on login failure
- [ ] Add version and link to repo on main screen or about view
- [ ] Add `[Enter]` and `SetFocus` on first Screen (may need to do a post render action after the route component is mounted)
- [ ] Hack on theme
    - [ ] Change the font colour of disable Entry
- [ ] Remove console opening/config stop logs on build
- [ ] Handle permission for file creation on os
- [ ] Github Action build/release for:
    - [ ] Linux
    - [ ] Mac
    - [ ] Windows


## Debt
- [ ] Handle `panic`s on app bootstrap

