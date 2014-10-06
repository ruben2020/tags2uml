
//   tags2uml
//   Copyright 2014 ruben2020 https://github.com/ruben2020/ 
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import "flag"

const ( 
    NONE       = 0
    ALL        = 1
    ONLYPUBLIC = 2
)

// options and their default values
var opt_blackbox bool = false
var opt_inherit bool = true
var opt_relationship bool = true
var opt_methods int = ONLYPUBLIC
var opt_members int = ONLYPUBLIC

var input_file string = "tags"
var output_file string = "-"
var help bool = false
var ver bool = false

func checkRange() bool {
    retval := true
    if (opt_methods < 0)||(opt_methods > 2) {retval = false}
    if (opt_members < 0)||(opt_members > 2) {retval = false}
    return retval
}

func parse_options() {
    flag.BoolVar(&opt_blackbox, "blackbox", false, "true for blackbox model, false for whitebox model (default=false)")
    flag.BoolVar(&opt_inherit, "inherit", true, "true to display inheritance info, false to not display (default=true)")
    flag.BoolVar(&opt_relationship, "relations", true, "true to display relationship info, false to not display (default=true)")
    flag.IntVar(&opt_methods, "methods", 2, "0=methods not displayed, 1=all methods displayed, 2=only public methods displayed (default)")
    flag.IntVar(&opt_members, "members", 2, "0=members not displayed, 1=all members displayed, 2=only public members displayed (default)")
    flag.StringVar(&input_file, "infile", "tags", "path to input file (default=\"tags\")")
    flag.StringVar(&output_file, "outfile", "-", "path to output file, use \"-\" for stdout (default=\"-\")")
    flag.BoolVar(&help, "help", false, "print help message")
    flag.BoolVar(&ver, "ver", false, "print version")
    flag.Parse()
}

