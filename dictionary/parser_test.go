package dictionary_test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	dict "github.com/bobrovde/radius/dictionary"
)

func TestParser(t *testing.T) {
	parser := dict.Parser{
		Opener: files,
	}

	d, err := parser.ParseFile("simple.dict")
	if err != nil {
		t.Fatal(err)
	}

	expected := &dict.Dictionary{
		Attributes: []*dict.Attribute{
			{
				Name: "User-Name",
				OID:  "1",
				Type: dict.AttributeString,
			},
			{
				Name:        "User-Password",
				OID:         "2",
				Type:        dict.AttributeOctets,
				FlagEncrypt: newIntPtr(1),
			},
			{
				Name: "Mode",
				OID:  "127",
				Type: dict.AttributeInteger,
			},
			{
				Name: "ARAP-Challenge-Response",
				OID:  "84",
				Type: dict.AttributeOctets,
				Size: newIntPtr(8),
			},
		},
		Values: []*dict.Value{
			{
				Attribute: "Mode",
				Name:      "Full",
				Number:    1,
			},
			{
				Attribute: "Mode",
				Name:      "Half",
				Number:    2,
			},
		},
	}

	if !reflect.DeepEqual(d, expected) {
		t.Fatalf("got %s, expected %s", dictString(d), dictString(expected))
	}
}

func TestParser_recursiveinclude(t *testing.T) {
	parser := dict.Parser{
		Opener: files,
	}

	d, err := parser.ParseFile("recursive_1.dict")
	pErr, ok := err.(*dict.ParseError)
	if !ok || pErr == nil || d != nil {
		t.Fatalf("got %v, expected *ParseError", pErr)
	}
	if _, ok := pErr.Inner.(*dict.RecursiveIncludeError); !ok {
		t.Fatalf("got %v, expected *RecursiveIncludeError", pErr.Inner)
	}
}

func newIntPtr(i int) *int {
	return &i
}

func dictString(d *dict.Dictionary) string {
	var b bytes.Buffer
	b.WriteString("dictionary.Dictionary\n")

	b.WriteString("\tAttributes:\n")
	for _, attr := range d.Attributes {
		b.WriteString(fmt.Sprintf("\t\t%q %q %q %#v %#v\n", attr.Name, attr.OID, attr.Type, attr.FlagHasTag, attr.FlagEncrypt))
	}

	b.WriteString("\tValues:\n")
	for _, value := range d.Values {
		b.WriteString(fmt.Sprintf("\t\t%q %q %d\n", value.Attribute, value.Name, value.Number))
	}

	b.WriteString("\tVendors:\n")
	for _, vendor := range d.Vendors {
		b.WriteString(fmt.Sprintf("\t\t%q %d\n", vendor.Name, vendor.Number))

		b.WriteString("\t\tAttributes:\n")
		for _, attr := range vendor.Attributes {
			b.WriteString(fmt.Sprintf("\t\t%q %q %q %#v %#v\n", attr.Name, attr.OID, attr.Type, attr.FlagHasTag, attr.FlagEncrypt))
		}

		b.WriteString("\t\tValues:\n")
		for _, value := range vendor.Values {
			b.WriteString(fmt.Sprintf("\t\t%q %q %d\n", value.Attribute, value.Name, value.Number))
		}
	}

	return b.String()
}
