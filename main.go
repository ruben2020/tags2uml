
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

import "fmt"
import "flag"
import "os"
import "log"
import "io/ioutil"

const sw_version string = "v0.01"

func main() {
    init_datastore()
    parse_options()
    if help {
        fmt.Println("The tags file must be created using the ctags option --fields=+latinK")
        fmt.Println("Usage of ", os.Args[0], ":")
        flag.PrintDefaults()
        return
    } else if ver {
        print_version()
        return
    } else if checkRange() == false {
        log.Println("Value of members or methods is out of range")
        flag.PrintDefaults()
        return
    } else if fileExists(input_file) == false {
        log.Fatal("File ", input_file, " does not exist!\nPlease use --help for help\n")
    } else {
        parseClass(input_file)
        parseMembersMethods(input_file)
        outs := dotmake()
        if (output_file == "-") {
            fmt.Println(outs)
        } else {
            err := ioutil.WriteFile(output_file, []byte(outs), 0644)
            if err != nil {
                panic(err)
            }
        }
    }
}

func fileExists(fn string) bool {
    retval := true
    if _, err := os.Stat(fn); os.IsNotExist(err) {
        retval = false
    }
    return retval
}

func print_version() {
fmt.Println("")
fmt.Println("   tags2uml ", sw_version)
fmt.Println("   Copyright 2014 ruben2020 https://github.com/ruben2020/")
fmt.Println("")
fmt.Println("   Licensed under the Apache License, Version 2.0 (the \"License\");")
fmt.Println("   you may not use this file except in compliance with the License.")
fmt.Println("   You may obtain a copy of the License at")
fmt.Println("")
fmt.Println("       http://www.apache.org/licenses/LICENSE-2.0")
fmt.Println("")
fmt.Println("   Unless required by applicable law or agreed to in writing, software")
fmt.Println("   distributed under the License is distributed on an \"AS IS\" BASIS,")
fmt.Println("   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.")
fmt.Println("   See the License for the specific language governing permissions and")
fmt.Println("   limitations under the License.")
fmt.Println("")
}

