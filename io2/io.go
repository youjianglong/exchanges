package io2

import "io"

// Writer io.Writer
type Writer = io.Writer

// Reader io.Reader
type Reader = io.Reader

// Closer io.Closer
type Closer = io.Closer

// WriteCloser io.WriteCloser
type WriteCloser = io.WriteCloser

// ReadCloser io.ReadCloser
type ReadCloser = io.ReadCloser

// MultiWriter io.MultiWriter
var MultiWriter = io.MultiWriter

// MultiReader io.MultiReader
var MultiReader = io.MultiReader
