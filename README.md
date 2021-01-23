# `dprgstd`

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)
[![Check the NOTICE](https://img.shields.io/badge/Check%20the-NOTICE-420C3B.svg)](./NOTICE)

## Overview

This is a library intended to be consumed by [my r/DailyProgrammer project](https://github.com/thecjharries/r-daily-programmer) doing the exercises from [r/DailyProgrammer](https://old.reddit.com/r/dailyprogrammer). I got tired of copypasta between days. For now, this will only get updated with things I write one day and find I need another day. I don't forsee lots of activity here but who knows. It's currently only near the end of January and I've only done a slew of easy exercises.

I'm trying to follow [good Go package naming](https://blog.golang.org/package-names). `dpstd` isn't a great package name for a few reasons so I'm running with `dprgstd`. I'm also keeping `dprg{,std}` out of subpackage names which means I'll have to rename them if I import a normal `std` package of the same name.
