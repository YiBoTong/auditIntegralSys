// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_StylePaneFilter struct {
	// Display All Styles
	AllStylesAttr *sharedTypes.ST_OnOff
	// Display Only Custom Styles
	CustomStylesAttr *sharedTypes.ST_OnOff
	// Display Latent Styles
	LatentStylesAttr *sharedTypes.ST_OnOff
	// Display Styles in Use
	StylesInUseAttr *sharedTypes.ST_OnOff
	// Display Heading Styles
	HeadingStylesAttr *sharedTypes.ST_OnOff
	// Display Numbering Styles
	NumberingStylesAttr *sharedTypes.ST_OnOff
	// Display Table Styles
	TableStylesAttr *sharedTypes.ST_OnOff
	// Display Run Level Direct Formatting
	DirectFormattingOnRunsAttr *sharedTypes.ST_OnOff
	// Display Paragraph Level Direct Formatting
	DirectFormattingOnParagraphsAttr *sharedTypes.ST_OnOff
	// Display Direct Formatting on Numbering Data
	DirectFormattingOnNumberingAttr *sharedTypes.ST_OnOff
	// Display Direct Formatting on Tables
	DirectFormattingOnTablesAttr *sharedTypes.ST_OnOff
	// Display Styles to Remove Formatting
	ClearFormattingAttr *sharedTypes.ST_OnOff
	// Display Heading 1 through 3
	Top3HeadingStylesAttr *sharedTypes.ST_OnOff
	// Only Show Visible Styles
	VisibleStylesAttr *sharedTypes.ST_OnOff
	// Use the Alternate Style Name
	AlternateStyleNamesAttr *sharedTypes.ST_OnOff
	// Bitmask of Suggested Filtering Options
	ValAttr *string
}

func NewCT_StylePaneFilter() *CT_StylePaneFilter {
	ret := &CT_StylePaneFilter{}
	return ret
}

func (m *CT_StylePaneFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AllStylesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:allStyles"},
			Value: fmt.Sprintf("%v", *m.AllStylesAttr)})
	}
	if m.CustomStylesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:customStyles"},
			Value: fmt.Sprintf("%v", *m.CustomStylesAttr)})
	}
	if m.LatentStylesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:latentStyles"},
			Value: fmt.Sprintf("%v", *m.LatentStylesAttr)})
	}
	if m.StylesInUseAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:stylesInUse"},
			Value: fmt.Sprintf("%v", *m.StylesInUseAttr)})
	}
	if m.HeadingStylesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:headingStyles"},
			Value: fmt.Sprintf("%v", *m.HeadingStylesAttr)})
	}
	if m.NumberingStylesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:numberingStyles"},
			Value: fmt.Sprintf("%v", *m.NumberingStylesAttr)})
	}
	if m.TableStylesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:tableStyles"},
			Value: fmt.Sprintf("%v", *m.TableStylesAttr)})
	}
	if m.DirectFormattingOnRunsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:directFormattingOnRuns"},
			Value: fmt.Sprintf("%v", *m.DirectFormattingOnRunsAttr)})
	}
	if m.DirectFormattingOnParagraphsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:directFormattingOnParagraphs"},
			Value: fmt.Sprintf("%v", *m.DirectFormattingOnParagraphsAttr)})
	}
	if m.DirectFormattingOnNumberingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:directFormattingOnNumbering"},
			Value: fmt.Sprintf("%v", *m.DirectFormattingOnNumberingAttr)})
	}
	if m.DirectFormattingOnTablesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:directFormattingOnTables"},
			Value: fmt.Sprintf("%v", *m.DirectFormattingOnTablesAttr)})
	}
	if m.ClearFormattingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:clearFormatting"},
			Value: fmt.Sprintf("%v", *m.ClearFormattingAttr)})
	}
	if m.Top3HeadingStylesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:top3HeadingStyles"},
			Value: fmt.Sprintf("%v", *m.Top3HeadingStylesAttr)})
	}
	if m.VisibleStylesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:visibleStyles"},
			Value: fmt.Sprintf("%v", *m.VisibleStylesAttr)})
	}
	if m.AlternateStyleNamesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:alternateStyleNames"},
			Value: fmt.Sprintf("%v", *m.AlternateStyleNamesAttr)})
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_StylePaneFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "directFormattingOnParagraphs" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DirectFormattingOnParagraphsAttr = &parsed
			continue
		}
		if attr.Name.Local == "allStyles" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.AllStylesAttr = &parsed
			continue
		}
		if attr.Name.Local == "directFormattingOnNumbering" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DirectFormattingOnNumberingAttr = &parsed
			continue
		}
		if attr.Name.Local == "latentStyles" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.LatentStylesAttr = &parsed
			continue
		}
		if attr.Name.Local == "headingStyles" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.HeadingStylesAttr = &parsed
			continue
		}
		if attr.Name.Local == "numberingStyles" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.NumberingStylesAttr = &parsed
			continue
		}
		if attr.Name.Local == "tableStyles" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.TableStylesAttr = &parsed
			continue
		}
		if attr.Name.Local == "directFormattingOnRuns" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DirectFormattingOnRunsAttr = &parsed
			continue
		}
		if attr.Name.Local == "customStyles" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.CustomStylesAttr = &parsed
			continue
		}
		if attr.Name.Local == "stylesInUse" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.StylesInUseAttr = &parsed
			continue
		}
		if attr.Name.Local == "directFormattingOnTables" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DirectFormattingOnTablesAttr = &parsed
			continue
		}
		if attr.Name.Local == "clearFormatting" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.ClearFormattingAttr = &parsed
			continue
		}
		if attr.Name.Local == "top3HeadingStyles" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.Top3HeadingStylesAttr = &parsed
			continue
		}
		if attr.Name.Local == "visibleStyles" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.VisibleStylesAttr = &parsed
			continue
		}
		if attr.Name.Local == "alternateStyleNames" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.AlternateStyleNamesAttr = &parsed
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_StylePaneFilter: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_StylePaneFilter and its children
func (m *CT_StylePaneFilter) Validate() error {
	return m.ValidateWithPath("CT_StylePaneFilter")
}

// ValidateWithPath validates the CT_StylePaneFilter and its children, prefixing error messages with path
func (m *CT_StylePaneFilter) ValidateWithPath(path string) error {
	if m.AllStylesAttr != nil {
		if err := m.AllStylesAttr.ValidateWithPath(path + "/AllStylesAttr"); err != nil {
			return err
		}
	}
	if m.CustomStylesAttr != nil {
		if err := m.CustomStylesAttr.ValidateWithPath(path + "/CustomStylesAttr"); err != nil {
			return err
		}
	}
	if m.LatentStylesAttr != nil {
		if err := m.LatentStylesAttr.ValidateWithPath(path + "/LatentStylesAttr"); err != nil {
			return err
		}
	}
	if m.StylesInUseAttr != nil {
		if err := m.StylesInUseAttr.ValidateWithPath(path + "/StylesInUseAttr"); err != nil {
			return err
		}
	}
	if m.HeadingStylesAttr != nil {
		if err := m.HeadingStylesAttr.ValidateWithPath(path + "/HeadingStylesAttr"); err != nil {
			return err
		}
	}
	if m.NumberingStylesAttr != nil {
		if err := m.NumberingStylesAttr.ValidateWithPath(path + "/NumberingStylesAttr"); err != nil {
			return err
		}
	}
	if m.TableStylesAttr != nil {
		if err := m.TableStylesAttr.ValidateWithPath(path + "/TableStylesAttr"); err != nil {
			return err
		}
	}
	if m.DirectFormattingOnRunsAttr != nil {
		if err := m.DirectFormattingOnRunsAttr.ValidateWithPath(path + "/DirectFormattingOnRunsAttr"); err != nil {
			return err
		}
	}
	if m.DirectFormattingOnParagraphsAttr != nil {
		if err := m.DirectFormattingOnParagraphsAttr.ValidateWithPath(path + "/DirectFormattingOnParagraphsAttr"); err != nil {
			return err
		}
	}
	if m.DirectFormattingOnNumberingAttr != nil {
		if err := m.DirectFormattingOnNumberingAttr.ValidateWithPath(path + "/DirectFormattingOnNumberingAttr"); err != nil {
			return err
		}
	}
	if m.DirectFormattingOnTablesAttr != nil {
		if err := m.DirectFormattingOnTablesAttr.ValidateWithPath(path + "/DirectFormattingOnTablesAttr"); err != nil {
			return err
		}
	}
	if m.ClearFormattingAttr != nil {
		if err := m.ClearFormattingAttr.ValidateWithPath(path + "/ClearFormattingAttr"); err != nil {
			return err
		}
	}
	if m.Top3HeadingStylesAttr != nil {
		if err := m.Top3HeadingStylesAttr.ValidateWithPath(path + "/Top3HeadingStylesAttr"); err != nil {
			return err
		}
	}
	if m.VisibleStylesAttr != nil {
		if err := m.VisibleStylesAttr.ValidateWithPath(path + "/VisibleStylesAttr"); err != nil {
			return err
		}
	}
	if m.AlternateStyleNamesAttr != nil {
		if err := m.AlternateStyleNamesAttr.ValidateWithPath(path + "/AlternateStyleNamesAttr"); err != nil {
			return err
		}
	}
	return nil
}
