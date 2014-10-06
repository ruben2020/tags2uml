
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

import "strings"
import "strconv"

func getrelations(membrs []memberinfo_st) []int {
    var clss classinfo_st
    var retval []int
    retmap := make(map[int]int)
    for membrsi := range membrs {
        clss = classmap[membrs[membrsi].datatype]
        if len(clss.name) == 0 {continue}
        retmap[clss.id] = 0
    }
    for idz, _ := range retmap {
        retval = append(retval, idz)
    }
    return retval
}

func dotmake() string {
var outs []string
var parentstr classinfo_st
var parenti int
outs = append(outs, "digraph G {\n")
for  _, value := range classmap {
    outs = append(outs, "n")
    outs = append(outs, strconv.Itoa(value.id))
    outs = append(outs, " [label=\"{")
    if opt_blackbox == true {
        outs = append(outs, value.name)
    } else {
        outs = append(outs, value.name)
        outs = append(outs, "|")
        outs = append(outs, add_members(value.members))
        outs = append(outs, "|")
        outs = append(outs, add_methods(value.methods, value.name))
    }
    outs = append(outs, "}\" style=filled fillcolor=\"#ffffff\" shape=\"record\"];\n")
}
idpairs := make(map[string]int)
parentmap := make(map[int][]int)
var tempintslice []int
for  _, value := range classmap {

    if opt_inherit {
        for parenti = range value.parents {
            parentstr = classmap[value.parents[parenti]]
            if len(parentstr.name) == 0 {continue}
            tempintslice = parentmap[parentstr.id]
            tempintslice = append(tempintslice, value.id)
            parentmap[parentstr.id] = tempintslice
        }
    }
    if opt_relationship {
        intlst := getrelations(value.members)
        for inti := range intlst {
            if intlst[inti] == value.id {continue}
            if idpairs[createPairOfIds(intlst[inti], value.id)] == 99 {continue}
            outs = append(outs, buildArrowLine(intlst[inti], value.id, "none"))
            idpairs[createPairOfIds(intlst[inti], value.id)] = 99
        }
    }
}

for par, children := range parentmap {
    outs = append(outs, "{ ")
    for childi := range children {
        outs = append(outs, "n")
        outs = append(outs, strconv.Itoa(children[childi]))
        outs = append(outs, " ")
    }
    outs = append(outs, "} -> n")
    outs = append(outs, strconv.Itoa(par))
    outs = append(outs, " [arrowhead=\"empty\"];\n")
}

outs = append(outs, "}\n")
return strings.Join(outs, "")
}

func createPairOfIds(id1 int, id2 int) string {
    var outs []string
    if (id1 < id2) {
        outs = append(outs, strconv.Itoa(id1))
        outs = append(outs, ",")
        outs = append(outs, strconv.Itoa(id2))
    } else {
        outs = append(outs, strconv.Itoa(id2))
        outs = append(outs, ",")
        outs = append(outs, strconv.Itoa(id1))
    }
    return strings.Join(outs, "")
}

func buildArrowLine(id1 int, id2 int, arrowtype string) string {
    var outs []string
    outs = append(outs, "n")
    outs = append(outs, strconv.Itoa(id1))
    outs = append(outs, " -> n")
    outs = append(outs, strconv.Itoa(id2))
    outs = append(outs, " [arrowhead=\"")
    outs = append(outs, arrowtype)
    outs = append(outs, "\"];\n")
    return strings.Join(outs, "")
}

func add_members(arr []memberinfo_st) string {
    var outs []string
    if opt_members == NONE {return ""}
    if opt_members == ONLYPUBLIC {
        for idx := range arr {
            if arr[idx].access == "+" {
                outs = append(outs, arr[idx].access)
                outs = append(outs, " ")
                outs = append(outs, arr[idx].name)
                if (len(arr[idx].datatype) > 0) {
                    outs = append(outs, " : ", arr[idx].datatype)
                }
                outs = append(outs, "\\l")
            }
        }
    } else {
        for idx := range arr {
            outs = append(outs, arr[idx].access)
            outs = append(outs, " ")
            outs = append(outs, arr[idx].name)
            if (len(arr[idx].datatype) > 0) {
                outs = append(outs, " : ", arr[idx].datatype)
            }
            outs = append(outs, "\\l")
        }
    }
    return strings.Join(outs, "")
}

func add_methods(arr []methodinfo_st, classname string) string {
    var outs []string
    dupmap := make(map[string]int)
    if opt_methods == NONE {return ""}
    if opt_methods == ONLYPUBLIC {
        for idx := range arr {
            if arr[idx].name == classname {continue}
            if dupmap[arr[idx].name] == 99 {continue}
            if arr[idx].access == "+" {
                outs = append(outs, arr[idx].access)
                outs = append(outs, " ")
                outs = append(outs, arr[idx].name)
                outs = append(outs, "()")
                if (len(arr[idx].returntype) > 0) {
                    outs = append(outs, " : ", arr[idx].returntype)
                }
                outs = append(outs, "\\l")
                dupmap[arr[idx].name] = 99
            }
        }
    } else {
        for idx := range arr {
            if arr[idx].name == classname {continue}
            if dupmap[arr[idx].name] == 99 {continue}
            outs = append(outs, arr[idx].access)
            outs = append(outs, " ")
            outs = append(outs, arr[idx].name)
            outs = append(outs, "()")
            if (len(arr[idx].returntype) > 0) {
                outs = append(outs, " : ", arr[idx].returntype)
            }
            outs = append(outs, "\\l")
            dupmap[arr[idx].name] = 99
        }
    }
    return strings.Join(outs, "")
}

