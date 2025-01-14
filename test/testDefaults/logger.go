package testDefaults

import "bitbucket.org/ryanair/gofrlib/log"

var LoggerConfig = log.NewConfiguration("DEBUG", "TEST-APP", "TEST-PROJECT", "TEST-PROJECT-GROUP", "TEST-VERSION", "TEST-PREFIX", "TEST-ENV")
