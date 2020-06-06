module git.sr.ht/~whereswaldon/sprig

go 1.14

require (
	gioui.org v0.0.0-20200602143738-d489c20b8421
	gioui.org/cmd v0.0.0-20200602143738-d489c20b8421 // indirect
	git.sr.ht/~whereswaldon/colorpicker v0.0.0-20200606004755-a710a970ef7c
	git.sr.ht/~whereswaldon/forest-go v0.0.0-20200517003538-529ac9248d93
	git.sr.ht/~whereswaldon/sprout-go v0.0.0-20200517010141-a4188845a9a8
	golang.org/x/crypto v0.0.0-20191122220453-ac88ee75c92c
	golang.org/x/exp v0.0.0-20200331195152-e8c3332aa8e5
	golang.org/x/text v0.3.2 // indirect
)

replace (
	gioui.org => git.sr.ht/~whereswaldon/gio v0.0.0-20200605203753-d456e9629aa1
	golang.org/x/crypto => github.com/ProtonMail/crypto v0.0.0-20200416114516-1fa7f403fb9c
)
