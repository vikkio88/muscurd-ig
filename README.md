# muscurd-ig
Portable multiplatform password manager.

Already done it in C# and Avalonia in [here](https://github.com/vikkio88/muscurd-i), attempting to do it in Golang and [Fyne](https://fyne.io/).

![muscurdig](https://github.com/vikkio88/muscurd-ig/assets/248805/65588440-c687-4fe8-bd9c-ecefe7e0a89b)


## Roadmap
- [x] TDD all entities
- [x] Design UI with Password Gating
- [x] Use [Clover](https://github.com/ostafen/clover) to persist data
- [x] Implement Search in List / Show All
- [x] Add confirm modal to delete
- [x] Implement NukeDb action to clean everything
- [x] Add Bind to List View to handle delete better (or remove Delete and move it to the details)
- [x] Fix jumping error message on login failure
- [x] Add version on main screen
- [x] Add link to repo on a About page/modal
- [ ] Add `[Enter]` and `SetFocus` on first Screen (may need to do a post render action after the route component is mounted)
- [ ] Hack on theme
    - [ ] Change the font colour of disable Entry
- [x] Remove console opening/config stop logs on build
- [x] Handle permission for file creation on os
- [x] Github Action build/release for:
    - [x] Linux
    - [x] Mac
    - [x] Windows


## Debt
- [ ] Handle `panic`s on app bootstrap
- [x] use clover annotation on `PasswordEntry` to lowercase field
- [x] Implement case insensitive search

