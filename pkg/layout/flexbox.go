package layout

import "strings"

type FlexBox struct {
	contents []string // The contents

	separator string // The separator between contents
	gap       int    // The number of separators between contents

	pad     string // The padding around the container
	padding int    // The amount of padding
}

// NewFlexBox creates a new FlexBox with the given contents
func NewFlexBox(contents ...string) *FlexBox {
	return &FlexBox{
		contents:  contents,
		separator: " ",
		pad:       "\n",
	}
}

// Modify the separator between the contents
func (f *FlexBox) WithSeparator(separator string) *FlexBox {
	f.separator = separator
	return f
}

// Modify the amount of gap between the contents
func (f *FlexBox) WithGap(gap int) *FlexBox {
	f.gap = gap
	return f
}

// Modify the padding character around the container
func (f *FlexBox) WithPad(pad string) *FlexBox {
	f.pad = pad
	return f
}

// Modify the amount of padding around the container
func (f *FlexBox) WithPadding(padding int) *FlexBox {
	f.padding = padding
	return f
}

// Converts the FlexBox to a string
func (f *FlexBox) String() string {
	padding := strings.Repeat(f.pad, f.padding)
	gap := strings.Repeat(f.separator, f.gap)

	s := padding

	for i, c := range f.contents {
		s += c
		if i < len(f.contents)-1 {
			s += gap
		}
	}
	s += "\n"

	s += padding

	return s
}

// Creates a Row
func Row(contents ...string) *FlexBox {
	return NewFlexBox(contents...).WithSeparator(" ")
}

// Creates a Column
func Column(contents ...string) *FlexBox {
	return NewFlexBox(contents...).WithSeparator("\n")
}
