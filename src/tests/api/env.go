package api

import "discord-clone/src/utils"

var BASE_URL = utils.GetEnv("BASE_URL", "../../../../.env")
var EXAMPLE_USER_ID = utils.GetEnv("EXAMPLE_USER_ID", "../../../../.env")
