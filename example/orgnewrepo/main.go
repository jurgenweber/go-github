// Copyright 2018 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The newrepo command utilizes go-github as a cli tool for
// creating new repositories. It takes an auth token as
// an environment variable and creates the new repo under
// the account affiliated with that token.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

var (
	name        = flag.String("name", "", "Name of repo to create in authenticated user's GitHub account.")
	description = flag.String("description", "", "Description of created repo.")
	org         = flag.String("organization", "", "the organization to add the repo to.")
	private     = flag.Bool("private", false, "Will created repo be private.")
	visibility  = flag.String("visibility", "internal", "the visiblity level: public, private, internal")
)

func main() {
	flag.Parse()
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	if *name == "" {
		log.Fatal("No name: New repos must be given a name")
	}
	if *org == "" {
		log.Fatal("No org: what org do I install this in?")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	r := &github.Repository{Name: name, Private: private, Description: description, Visibility: visibility}
	repo, _, err := client.Repositories.Create(ctx, *org, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}
