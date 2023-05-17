# muscurd-ig
Portable multiplatform password manager.

Already done it in C# and Avalonia in [here](https://github.com/vikkio88/muscurd-i), attempting to do it in Golang and [Fyne](https://fyne.io/).

## Roadmap
- [x] TDD all entities
- [x] Design UI with Password Gating
- [x] Use [Clover](https://github.com/ostafen/clover) to persist data
- [ ] Implement Search in List / Show All
- [ ] Add Bind to List View to handle delete better (or remove Delete and move it to the details)
- [ ] Add `[Enter]` and `SetFocus` on first Screen
- [ ] Add confirm modal to delete
- [ ] Hack on theme
- [ ] Github Action build/release for:
    - [ ] Linux
    - [ ] Mac
    - [ ] Windows


## Debt
- [ ] Handle `panic`s on app bootstrap

