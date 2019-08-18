
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

import "os"
import "fmt"
import "bufio"
import "regexp"
import "strings"


func getFileName(input string) string {
   endPos := strings.LastIndex(input, ".")
   if (endPos < 0) {
       endPos = len(input)
   }
   beginPos := 0
   for i := endPos - 1; i >= 0; i-- {
       r := input[i]
       if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
           beginPos = i
           break
       }
   }
   return input[beginPos+1:endPos]
}


func parseClass(fn string) {

file, err := os.Open(fn)
if err != nil {
	fmt.Println(err)
}

scanner := bufio.NewScanner(file)
re := regexp.MustCompile(`^([A-Za-z0-9_]+)\t([^\t]+)\t([^\t]+)\t(class|interface)`)
re2 := regexp.MustCompile(`inherits:([A-Za-z0-9_\:,]+)`)
for scanner.Scan() {
    match := re.FindStringSubmatch(scanner.Text())
    if len(match) != 0 {
         newclass := classinfo_st{}
         newclass.name = match[1]
         filepath := match[2]
         filebasename := getFileName(filepath)
         if (newclass.name != filebasename) {
             newclass.name = filebasename + "." + match[1]
         } else {
             newclass.name = match[1]
         }

         match2 := re2.FindStringSubmatch(scanner.Text())
         if len(match2) != 0 {
             newclass.parents = strings.Split(match2[1], ",")
         }
         newclass.id = idcounter
         idcounter++
         classmap[newclass.name] = newclass
    }
}

if err := scanner.Err(); err != nil {
    fmt.Println(err)
}

}


func parseMembersMethods(fn string) {

file, err := os.Open(fn)
if err != nil {
	fmt.Println(err)
}

scanner := bufio.NewScanner(file)
re := regexp.MustCompile(`^([A-Za-z0-9_]+)\t([^\t]+)\t([^\t]+)\t([A-Za-z]+)`)
rea := regexp.MustCompile(`access:([A-Za-z0-9_]+)`)
rec := regexp.MustCompile(`(interface|class):([A-Za-z0-9_\.]+)`)
rel := regexp.MustCompile(`language:([A-Za-z0-9_\#]+)`)
ret := regexp.MustCompile(`\/\^([ ]*)([A-Za-z0-9_\.]+)([^A-Za-z0-9_]+)(.*)\$\/`)
for scanner.Scan() {
    match := re.FindStringSubmatch(scanner.Text())
    if (len(match) == 0) {continue}
    matchc := rec.FindStringSubmatch(scanner.Text())
    ci := classinfo_st{}
    var cn string
    if (len(matchc) != 0) {cn = matchc[2]}

    //cnsep := strings.LastIndex(cn, ".")
    //if (cnsep != -1) {
        //cn = cn[cnsep+1:]
    //}
    if (len(cn) != 0) {ci = classmap[cn]}
    if (len(ci.name) == 0) {continue}
    matcha := rea.FindStringSubmatch(scanner.Text())
    matchl := rel.FindStringSubmatch(scanner.Text())
    access1 := " "
    if (len(matcha) != 0) {
        if matcha[1] == "public" {
           access1 = "+"
        } else if matcha[1] == "private" {
           access1 = "-"
        } else if matcha[1] == "protected" {
           access1 = "#"
        }
    }
    thetype := match[4]
    matcht := ret.FindStringSubmatch(remove_keywords(scanner.Text()))
    datatype := ""
    if (len(matcht) != 0)&&(datatype_supported(matchl[1])) {
        datatype = matcht[2]
    }
    if (thetype == "member")||(thetype == "field") {
        memberinfo := memberinfo_st{match[1], access1, datatype}
        ci.members = append(ci.members, memberinfo)
    } else if (thetype == "function")||(thetype == "method") {
        methodinfo := methodinfo_st{match[1], access1, datatype}
        ci.methods = append(ci.methods, methodinfo)
    }
    classmap[ci.name] = ci
    
}

if err := scanner.Err(); err != nil {
    fmt.Println(err)
}

}

func remove_keywords(txt string) string {
    str1 := txt
    str1 = strings.Replace(str1, "private", "", 1)
    str1 = strings.Replace(str1, "public", "", 1)
    str1 = strings.Replace(str1, "protected", "", 1)
    str1 = strings.Replace(str1, "static", "", 1)
    str1 = strings.Replace(str1, "volatile", "", 1)
    str1 = strings.Replace(str1, "synchronized", "", 1)
    str1 = strings.Replace(str1, "final", "", 1)
    str1 = strings.Replace(str1, "const", "", 1)
    str1 = strings.Replace(str1, "abstract", "", 1)
    str1 = strings.Replace(str1, "struct", "", 1)
    str1 = strings.Replace(str1, "union", "", 1)
    str1 = strings.Replace(str1, "enum", "", 1)
    str1 = strings.Replace(str1, "override", "", 1)
    str1 = strings.Replace(str1, "internal", "", 1)
    str1 = strings.Replace(str1, "extern", "", 1)
    str1 = strings.Replace(str1, "readonly", "", 1)
    str1 = strings.Replace(str1, "*", "", -1)
    str1 = strings.Replace(str1, ":", "", -1)
    return str1
}

func datatype_supported(lang string) bool {
    return (lang == "C++")||(lang == "C#")||(lang == "Java")
}

