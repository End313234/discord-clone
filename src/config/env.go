package config

import "discord-clone/src/utils"

var BASE_URL = utils.GetEnv("BASE_URL")
var EXAMPLE_USER_ID = utils.GetEnv("EXAMPLE_USER_ID")
var PORT = utils.GetEnv("PORT")
