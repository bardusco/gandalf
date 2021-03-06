// Copyright 2014 gandalf authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hook

import (
	"github.com/globocom/config"
	"github.com/globocom/gandalf/fs"
	"strings"
	"io"
	"os"
)

// Adds a hook script.
func Add(name string, body io.ReadCloser) error {
	path, err := config.GetString("git:bare:template")
	if err != nil {
	  return err
	}
	s := []string{path, "hooks", name}
	scriptPath := strings.Join(s, "/")
	file, err := fs.Filesystem().OpenFile(scriptPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
	  return err
	}
	defer file.Close()
	_, err = io.Copy(file, body)
	if err != nil {
	  return err
	}
	return nil
}
