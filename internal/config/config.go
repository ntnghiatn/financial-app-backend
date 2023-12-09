package config

import "flag"

var DataDirectory = flag.String("data-directory", "", "Path for loading template and migration script")
