package helpers

import "strings"

type FlexBox struct {
	contents  []string
	separator string
	gap       int
	padding   int
}

func NewFlexBox(contents ...string) *FlexBox {
	return &FlexBox{
		contents:  contents,
		separator: " ",
	}
}

func (f *FlexBox) WithSeparator(separator string) *FlexBox {
	f.separator = separator
	return f
}

func (f *FlexBox) WithGap(gap int) *FlexBox {
	f.gap = gap
	return f
}

func (f *FlexBox) WithPadding(padding int) *FlexBox {
	f.padding = padding
	return f
}

func (f *FlexBox) String() string {
	s := strings.Repeat("\n", f.padding)
	for i, c := range f.contents {
		s += c
		if i < len(f.contents)-1 {
			s += f.separator
			for i := 0; i < f.gap; i++ {
				s += " "
			}
		}
	}
	s += "\n"
	s += strings.Repeat("\n", f.padding)
	return s
}

func FlexBoxRow(contents ...string) *FlexBox {
	return NewFlexBox(contents...).WithSeparator(" ")
}

func FlexBoxColumn(contents ...string) *FlexBox {
	return NewFlexBox(contents...).WithSeparator("\n")
}
