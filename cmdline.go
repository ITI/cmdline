// package cmdline provides an interface for defining and accessing
// arguments to a go program that are specified on the command
// line of the program execution
package cmdline

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FlagArgType is the type basis for an enumerated type of command-line flags
type FlagArgType int

// IntFlag, Int64Flag, FloatFlag, StringFlag, and BoolFlag are the enumerated
// types of scalar types of arguments declared on the command line
const (
	IntFlag FlagArgType = iota
	Int64Flag
	FloatFlag
	StringFlag
	BoolFlag
	None
)

// FlagTypeString converts a command line argument enumerated type into
// a string representation
func FlagTypeString(type_name FlagArgType) string {
	switch type_name {
	case IntFlag:
		return "IntFlag"
	case Int64Flag:
		return "Int64Flag"
	case FloatFlag:
		return "FloatFlag"
	case StringFlag:
		return "StringFlag"
	case BoolFlag:
		return "BoolFlag"
	default:
		return "None"
	}
}

// The arg interface defines what is needed for a type to
// be used as a command line argument
type arg interface {
	ArgType() FlagArgType // what kind of argument is represented
	Name() string         // name of the argument
	Set(string)           // save the argument in the type's structure, extracted as a string from the command line
	Get() any             // return the argument in its native form, which means the return type for the interface is 'any'
	Loaded() bool         // has a flag with the specified name been set
	Required() bool       // is this argument required
}

// Below we have definitions for types intVar, int64Var, floatVar, stringVar, and boolVar.
// They are identical, save that the v_value attribute is type specific.
// For each
//	- v_name saves the name declared for the variable
//	- v_req flags whether a command must declare this flag and value
//  - v_loaded flags whether the command was recognized on the command line and loaded

// intVar represents a command variable whose type is an integer of default length
type intVar struct {
	v_name   string
	v_value  int
	v_req    bool
	v_loaded bool
}

// createIntVar is a constructor whose arguments give the argument a name and indicate whether it is required.
func createIntVar(name string, req bool) *intVar {
	vs := &intVar{v_name: name,
		v_req:    req,
		v_loaded: false}
	return vs
}

// ArgType returns the enumerated type IntFlag
func (vs *intVar) ArgType() FlagArgType {
	return IntFlag
}

// Name returns the name of the command line variable
func (vs *intVar) Name() string {
	return vs.v_name
}

// Set saves the type-specific represention of the command variable's string extracted from the command line
func (vs *intVar) Set(value string) {
	sv, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("Error setting integer flag variable")
		return
	}
	vs.v_value = int(sv)
	vs.v_loaded = true
}

// Get returns the command variable's value with unspecified type
func (vs *intVar) Get() any {
	return vs.v_value
}

// Loaded indicates whether this command variable was extracted from the command line
func (vs *intVar) Loaded() bool {
	return vs.v_loaded
}

// Required indicates whether this command variable must appear on the command line
func (vs *intVar) Required() bool {
	return vs.v_req

}

// int64Var represents a command variable whose type is an integer of 64 bits
type int64Var struct {
	v_name   string
	v_value  int64
	v_req    bool
	v_loaded bool
}

// createInt64Var is a constructor whose arguments give the argument a name and indicate whether it is required.
func createInt64Var(name string, req bool) *int64Var {
	vs := &int64Var{v_name: name,
		v_req:    req,
		v_loaded: false}
	return vs
}

// ArgType returns the enumerated type Int64Flag
func (vs *int64Var) ArgType() FlagArgType {
	return Int64Flag
}

// Name returns the name of the command line variable
func (vs *int64Var) Name() string {
	return vs.v_name
}

// Set saves the type-specific represention of the command value's string extracted from the command line
func (vs *int64Var) Set(value string) {
	sv, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("Error seting integer flag variable")
		return
	}
	vs.v_value = int64(sv)
	vs.v_loaded = true
}

// Get returns the command variable's value with unspecified type
func (vs *int64Var) Get() any {
	return vs.v_value
}

// Loaded indicates whether this command variable was extracted from the command line
func (vs *int64Var) Loaded() bool {
	return vs.v_loaded
}

// Required indicates whether this command variable must appear on the command line
func (vs *int64Var) Required() bool {
	return vs.v_req

}

// floatVar represents a command variable whose type is a float with 64 bits
type floatVar struct {
	v_name   string
	v_value  float64
	v_req    bool
	v_loaded bool
}

// createFloatVar is a constructor whose arguments give the argument a name and indicate whether it is required.
func createFloatVar(name string, req bool) *floatVar {
	vs := &floatVar{v_name: name,
		v_req:    req,
		v_loaded: false}
	return vs
}

// ArgType returns the enumerated type FloatFlag
func (vs *floatVar) ArgType() FlagArgType {
	return FloatFlag
}

// Name returns the name of the command line variable
func (vs *floatVar) Name() string {
	return vs.v_name
}

// Set saves the type-specific represention of the command value's string extracted from the command line
func (vs *floatVar) Set(value string) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error setting float in cmdline")
		return
	}
	vs.v_value = v
	vs.v_loaded = true
}

// Get returns the command variable's value with unspecified type
func (vs *floatVar) Get() any {
	return vs.v_value
}

// Loaded indicates whether this command variable was extracted from the command line
func (vs *floatVar) Loaded() bool {
	return vs.v_loaded
}

// Required indicates whether this command variable must appear on the command line
func (vs *floatVar) Required() bool {
	return vs.v_req

}

// stringVar represents a command variable whose type is a string
type stringVar struct {
	v_name   string
	v_value  string
	v_req    bool
	v_loaded bool
}

// createStringVar is a constructor whose arguments give the argument a name and indicate whether it is required.
func createStringVar(name string, req bool) *stringVar {
	vs := &stringVar{v_name: name,
		v_req:    req,
		v_loaded: false}
	return vs
}

// ArgType returns the enumerated type StringFlag
func (vs *stringVar) ArgType() FlagArgType {
	return StringFlag
}

// Name returns the name of the command line variable
func (vs *stringVar) Name() string {
	return vs.v_name
}

// Set saves the type-specific represention of the command value's string extracted from the command line
func (vs *stringVar) Set(value string) {
	vs.v_value = value
	vs.v_loaded = true
}

// Get returns the command variable's value with unspecified type
func (vs *stringVar) Get() any {
	return vs.v_value
}

// Loaded indicates whether this command variable was extracted from the command line
func (vs *stringVar) Loaded() bool {
	return vs.v_loaded
}

// Required indicates whether this command variable must appear on the command line
func (vs *stringVar) Required() bool {
	return vs.v_req

}

// boolVar represents a command variable whose type is a Boolean flag
type boolVar struct {
	v_name   string
	v_value  bool
	v_req    bool
	v_loaded bool
}

// createBoolVar is a constructor whose arguments give the argument a name and indicate whether it is required.
func createBoolVar(name string, req bool) *boolVar {
	vs := &boolVar{v_name: name,
		v_req:    req,
		v_loaded: false}
	return vs
}

// ArgType returns the enumerated type BoolFlag
func (vs *boolVar) ArgType() FlagArgType {
	return BoolFlag
}

// Name returns the name of the command line variable
func (vs *boolVar) Name() string {
	return vs.v_name
}

// Set saves the type-specific represention of the command value's string extracted from the command line
func (vs *boolVar) Set(value string) {
	v := false
	if value == "T" || value == "t" || value == "True" || value == "true" {
		v = true
	}
	vs.v_value = v
	vs.v_loaded = true
}

// Get returns the command variable's value with unspecified type
func (vs *boolVar) Get() any {
	return vs.v_value
}

// Loaded indicates whether this command variable was extracted from the command line
func (vs *boolVar) Loaded() bool {
	return vs.v_loaded
}

// Required indicates whether this command variable must appear on the command line
func (vs *boolVar) Required() bool {
	return vs.v_req

}

// A CmdParser struct maps the flag names of command variables to their type specific representations
type CmdParser struct {
	vars map[string]arg
}

// NewCmdParser is a constructor, initializes an empty CmdParser data structure
func NewCmdParser() *CmdParser {
	empty_vars := make(map[string]arg)
	cp := &CmdParser{vars: empty_vars}
	return cp
}

// AddFlag includes a new command flag to the parser.  The arguments give
// the type of the flag in enumerated type form, the name of the flag, and whether the flag is required
func (cp *CmdParser) AddFlag(arg_type FlagArgType, arg_name string, arg_req bool) {

	// for each type of command argument call the constructor for that type and save the
	// result (indexed by command argument name) in the CmdParser's 'vars' map
	switch arg_type {
	case IntFlag:
		v := createIntVar(arg_name, arg_req)
		cp.vars[arg_name] = v
		break

	case Int64Flag:
		v := createInt64Var(arg_name, arg_req)
		cp.vars[arg_name] = v
		break

	case FloatFlag:
		v := createFloatVar(arg_name, arg_req)
		cp.vars[arg_name] = v
		break

	case StringFlag:
		v := createStringVar(arg_name, arg_req)
		cp.vars[arg_name] = v
		break

	case BoolFlag:
		v := createBoolVar(arg_name, arg_req)
		cp.vars[arg_name] = v
		break
	}
}

// SetVar calls an arg interface function with a command variable name and string-encoded value
// from the command line to set the value in the type-specific struct.
func (cp *CmdParser) SetVar(name string, value string) {
	cp.vars[name].Set(value)
}

// GetVar returns the type-unspecified value of a command variable that was created in the CmdParser,
// and given a value from the command line.  It is assumed that the application calls IsLoaded
// before GetVar to ascertain that a value is indeed present
func (cp *CmdParser) GetVar(name string) any {
	gv, present := cp.vars[name]
	if present {
		return gv.Get()
	}
	msg := fmt.Sprintf("CmdParser.GetVar given unrecognized variable name %s\n", name)
	panic(msg)
}

// IsFlag returns a bool indicating whether the input argument string 'name'
// has been used to create a command variable in the CmdParser
func (cp *CmdParser) IsFlag(name string) bool {
	_, present := cp.vars[name]
	return present
}

// IsLoaded returns a bool indicating whether a command variable with the input argument
// string 'name' was recognized on the command line and so had a value stored
func (cp *CmdParser) IsLoaded(name string) bool {
	if !cp.IsFlag(name) {
		return false
	}
	return cp.vars[name].Loaded()
}

// IsRequired returns a bool indicating whether a command variable with the input argument
// 'name' was declared to be required
func (cp *CmdParser) IsRequired(name string) bool {
	if !cp.IsFlag(name) {
		return false
	}
	return cp.vars[name].Required()
}

type flagValue struct {
	flag  string
	value string
}

// ParseFromString separates the command line string into individual command statements
// and stores them in the CmdParser
func (cp *CmdParser) ParseFromString(cmd_string string) bool {

	// break up the input string by white space
	pieces := strings.Fields(cmd_string)

	// some of the arguments may be only flags (indicating value true), so
	// scan the list first to create flag-value pairs
	cmdVar := make([]flagValue, 0)

	idx := 0
	for idx < len(pieces) {
		if strings.HasPrefix(pieces[idx], "-") && (idx == len(pieces)-1 || strings.HasPrefix(pieces[idx+1], "-")) {
			// position idx is a flag
			fv := flagValue{flag: strings.Replace(pieces[idx], "-", "", 1), value: "true"}
			cmdVar = append(cmdVar, fv)
			idx += 1
		} else if strings.HasPrefix(pieces[idx], "-") {
			fv := flagValue{flag: strings.Replace(pieces[idx], "-", "", 1), value: pieces[idx+1]}
			cmdVar = append(cmdVar, fv)
			idx += 2
		} else {
			fmt.Printf("formatting problem in command line from %s\n", strings.Join(pieces[idx:], " "))
			return false
		}
	}

	// check that all the flags obtained have been declared for the CmdParser
	errMsg := []string{}
	for _, fv := range cmdVar {
		_, present := cp.vars[fv.flag]
		if !present {
			errMsg = append(errMsg, "-"+fv.flag)
		}
	}

	if len(errMsg) > 0 {
		msg := fmt.Sprintf("Flags not declared in CmdParser: %s, ignored", strings.Join(errMsg, ","))
		fmt.Println(msg)
		// return false
	}

	// now set the variables
	for _, fv := range cmdVar {
		_, present := cp.vars[fv.flag]
		if present {
			cp.SetVar(fv.flag, fv.value)
		}
	}

	// and finally, ensure that every variable that is required is present
	errMsg = []string{}
	for name, value := range cp.vars {
		if value.Required() && !value.Loaded() {
			errMsg = append(errMsg, "-"+name)
		}
	}

	if len(errMsg) > 0 {
		msg := fmt.Sprint("Flags required but missing: %s", strings.Join(errMsg, ","))
		fmt.Println(msg)
		return false
	}
	return true
}

// ParseFromCmdLine gets the command line string from os.Args, i.e., the run-time command line
func (cp *CmdParser) ParseFromCmdLine() bool {

	// join the already parsed command line pieces with white space to create a single string
	cmd_str := strings.Join(os.Args[1:], " ")

	// parse that string
	return cp.ParseFromString(cmd_str)
}

// ParseFromFile gets the command line flags from a file. This enables separation across lines
// and comments
func (cp *CmdParser) ParseFromFile(filename string) bool {

	// open the file
	inFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open command line file")
		return false
	}
	defer inFile.Close()

	// read the file line by line, skipping empty lines and commented lines
	cmd_string := ""
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {

		// line by line
		nxt_line := scanner.Text()

		// skip empty lines
		if string(nxt_line) == "" {
			continue
		}

		// remove anything after a '#'
		whitespace := 0
		for idx := 0; idx < len(nxt_line); idx++ {

			// move the whitespace marker up
			if string(nxt_line[idx]) == "" || string(nxt_line[idx]) == "\t" {
				whitespace += 1
			} else if string(nxt_line[idx]) == "#" {
				// hit a comment character.  Is there nothing but white space to the left?
				if whitespace == idx {

					// yes, so this line is empty
					nxt_line = ""
					break
				} else {
					// there is stuff to parse before the comment character
					nxt_line = nxt_line[whitespace:idx]
					break
				}
			}
		}

		if nxt_line != "" {
			// get rid of "\n" if present
			nxt_line = strings.Replace(nxt_line, "\n", "", 1)
			cmd_string = cmd_string + " " + nxt_line
		}
	}
	return cp.ParseFromString(cmd_string)
}

// Parse looks for a leading "-is" on the command line to determine whether to
// parse from a file (e.g., "-is" is present), or get the arguments from the command line itself
func (cp *CmdParser) Parse() bool {

	// see if the command line is empty and if so flag the error
	if len(os.Args) == 1 {
		fmt.Println("call requires command line arguments")
		os.Exit(1)
	}

	// see if the command line points to a file
	parsedOK := true
	if len(os.Args) > 1 && os.Args[1] == "-is" {
		// parse from the file
		cmdfile := os.Args[2]
		parsedOK = cp.ParseFromFile(cmdfile)
	} else {
		parsedOK = cp.ParseFromCmdLine()
	}

	if !parsedOK {
		panic("Command line parsing error")
	}
	return true
}
