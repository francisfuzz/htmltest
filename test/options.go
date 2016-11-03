package test

import (
  "strings"
  "github.com/imdario/mergo"
  "github.com/wjdp/htmltest/issues"
)

type Options struct {
  DirectoryPath string
  FilePath string

  CheckExternal bool
  CheckInternal bool
  CheckMailto bool
  CheckTel bool
  EnforceHTTPS bool

  TestFilesConcurrently bool
  LogLevel int

  DirectoryIndex string

  ExternalTimeout int
  StripQueryString bool
  StripQueryExcludes []string
}

var Opts Options

func DefaultOptions() map[string]interface{} {
  // Specify defaults here
  return map[string]interface{}{
    "CheckExternal": true,
    "CheckInternal": true,
    "CheckMailto": true,
    "CheckTel": true,
    "EnforceHTTPS": false,

    "TestFilesConcurrently": false,
    "LogLevel": issues.INFO,

    "DirectoryIndex": "index.html",

    "ExternalTimeout": 3,
    "StripQueryString": true,
    "StripQueryExcludes": []string{"fonts.googleapis.com"},
  }
}

func SetOptions(optsUser map[string]interface{}) {
  // Merge user and default options, set Opts var
  opts := DefaultOptions()
  mergo.MergeWithOverwrite(&opts, optsUser)
  Opts = Options{}
  mergo.MapWithOverwrite(&Opts, opts)
}

func InList(list []string, key string) bool {
  for _, item := range list {
    if strings.Contains(key, item) { return true }
  }
  return false
}