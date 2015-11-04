// Copyright 2015 The Tops'l Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

type MyBox struct {
	views.BoxLayout
}

func (m *MyBox) HandleEvent(ev tcell.Event) bool {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape {
			views.AppQuit()
			return true
		}
	}
	return m.BoxLayout.HandleEvent(ev)
}

func main() {

	if e := views.AppInit(); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(1)
	}

	outer := &MyBox{}
	outer.SetOrientation(views.Vertical)

	title := &views.TextBar{}
	title.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorYellow).
		Foreground(tcell.ColorBlack))
	title.SetCenter("Horizontal Boxes", tcell.StyleDefault)
	title.SetLeft("ESC to exit", tcell.StyleDefault.
		Background(tcell.ColorBlue).
		Foreground(tcell.ColorWhite))
	title.SetRight("==>X", tcell.StyleDefault)

	box := views.NewBoxLayout(views.Horizontal)

	l := views.NewText()
	m := views.NewText()
	r := views.NewText()

	l.SetText("Left (0.0)")
	m.SetText("Middle (0.7)")
	r.SetText("Right (0.3)")
	l.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).
		Background(tcell.ColorRed))
	m.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).
		Background(tcell.ColorLime))
	r.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlue))
	l.SetAlignment(views.AlignBegin)
	m.SetAlignment(views.AlignMiddle)
	r.SetAlignment(views.AlignEnd)

	box.AddWidget(l, 0)
	box.AddWidget(m, 0.7)
	box.AddWidget(r, 0.3)

	outer.AddWidget(title, 0)
	outer.AddWidget(box, 1)
	views.SetApplication(outer)
	views.RunApplication()
}
