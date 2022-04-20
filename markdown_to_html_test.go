package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHeader(t *testing.T) {

}

func TestIsLink(t *testing.T) {

}

func TestConvertMarkdown2HTML(t *testing.T) {
	testCases:= []struct{
		name string
		markdown string
		expectedHTML string
	}{
		{
			markdown: test1,
			name: "sample 1",
			expectedHTML: expected1,
		},
		{
			markdown: test2,
			name: "sample 2",
			expectedHTML: expected2,
		},
	}

	for _, tc := range testCases {
		var output bytes.Buffer
		t.Run(tc.name, func(t *testing.T) {
			ConvertMarkdown2HTML(strings.NewReader(tc.markdown), &output)
			assert.Equal(t, tc.expectedHTML, output.String())
		})
	}
}

func TestConvertLineMarkdown2HTML(t *testing.T) {

}

const test1 = `
# Sample Document

Hello!

This is sample markdown for the [Linktext](https://www.foo.com) code exercise.
`

const expected1 = `<h1>Sample Document</h1>
<p>Hello!</p>
<p>This is sample markdown for the <a href="https://www.foo.com">Linktext</a> code exercise.</p>
`

const test2 = `
# Header one

Hello there

How are you?
What's going on?

## Another Header

This is a paragraph [with an inline link](http://google.com). Neat, eh?

## This is a header [with a link](http://yahoo.com)
`

const expected2 = `<h1>Header one</h1>
<p>Hello there</p>
<p>How are you?
What's going on?</p>
<h2>Another Header</h2>
<p>This is a paragraph <a href="http://google.com">with an inline link</a>. Neat, eh?</p>
<h2>This is a header <a href="http://yahoo.com">with a link</a></h2>
`