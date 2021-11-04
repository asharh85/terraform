// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tfresource

import (
	"log"
)

const Version = "4.49.0"
const ReleaseDate = "2021-10-20"

func PrintVersion() {
	log.Printf("[INFO] terraform-provider-oci %s\n", Version)
}