/*
 * LF: Global Fully Replicated Key/Value Store
 * Copyright (C) 2018-2019  ZeroTier, Inc.  https://www.zerotier.com/
 *
 * Licensed under the terms of the MIT license (see LICENSE.txt).
 */

package lf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"sort"
	"strings"
)

// ClientConfigName is the default name of the client config file
const ClientConfigName = "client.json"

// ClientConfigOwner is a locally configured owner with private key information.
type ClientConfigOwner struct {
	Owner        Blob
	OwnerPrivate Blob
	Default      bool
}

// GetOwner gets an Owner object (including private key) from this ClientConfigOwner
func (co *ClientConfigOwner) GetOwner() (o *Owner, err error) {
	o, err = NewOwnerFromPrivateBytes(co.OwnerPrivate)
	return
}

// ClientConfig is the JSON format for the client configuration file.
type ClientConfig struct {
	Urls   []string                      ``         // URLs of full nodes and/or proxies
	Owners map[string]*ClientConfigOwner ``         // Owners by name
	Dirty  bool                          `json:"-"` // Non-persisted flag that can be used to indicate the config should be saved on client exit
}

// Load loads this client config from disk or initializes it with defaults if load fails.
func (c *ClientConfig) Load(path string) error {
	d, err := ioutil.ReadFile(path)
	if err == nil && len(d) > 0 {
		err = json.Unmarshal(d, c)
	}
	if c.Owners == nil {
		c.Owners = make(map[string]*ClientConfigOwner)
	}
	c.Dirty = false

	// If the file didn't exist, init config with defaults.
	if err != nil && os.IsNotExist(err) {
		c.Urls = nil // TODO: use acceptable defaults
		owner, _ := NewOwner(OwnerTypeEd25519)
		dflName := "user"
		u, _ := user.Current()
		if u != nil && len(u.Username) > 0 {
			dflName = strings.ReplaceAll(u.Username, " ", "") // use the current login user name if it can be determined
		}
		c.Owners[dflName] = &ClientConfigOwner{
			Owner:        owner.Bytes(),
			OwnerPrivate: owner.PrivateBytes(),
			Default:      true,
		}
		c.Dirty = true
		err = nil
	}

	return err
}

// Save writes this client config to disk and reset the dirty flag.
func (c *ClientConfig) Save(path string) error {
	// Make sure there is one and only one default owner.
	if len(c.Owners) > 0 {
		haveDfl := false
		var names []string
		for n, o := range c.Owners {
			names = append(names, n)
			if haveDfl {
				if o.Default {
					o.Default = false
				}
			} else if o.Default {
				haveDfl = true
			}
		}
		if !haveDfl {
			sort.Strings(names)
			c.Owners[names[0]].Default = true
		}
	}

	d, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, d, 0600) // 0600 since this file contains secrets
	if err == nil {
		c.Dirty = false
		return nil
	}
	return err
}
