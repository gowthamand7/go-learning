package main

import (
	"fmt"
	"strconv"
)

type Group struct {
	Name string
}

type User struct {
	Name  string
	Group Group
}

type Permission struct {
	Read    bool
	Write   bool
	Execute bool
}

type Permissions struct {
	Owner  Permission
	Group  Permission
	Others Permission
}

type File struct {
	Name        string
	Size        int
	Parent      *File
	Owner       User
	Permissions Permissions
	Type        string
}

func main() {

	files := map[string]*File{}

	systemOwner := User{
		Name: "System",
	}

	gowthamn := User{
		Name: "Gowtham",
		Group: Group{
			"Leads",
		},
	}

	p := Permission{
		Read:    true,
		Write:   true,
		Execute: true,
	}

	FullPermission := Permissions{
		Owner:  p,
		Group:  p,
		Others: p,
	}

	files["root"] = &File{
		Name:        "Desktop",
		Size:        12,
		Parent:      nil,
		Owner:       systemOwner,
		Type:        "Directory",
		Permissions: FullPermission,
	}

	files["Documents"] = &File{
		Name:        "Documents",
		Size:        12,
		Parent:      nil,
		Owner:       systemOwner,
		Type:        "Directory",
		Permissions: FullPermission,
	}

	files["Todos"] = &File{
		Name:        "Todos.txt",
		Size:        6585,
		Owner:       gowthamn,
		Type:        "File",
		Permissions: FullPermission,
		Parent:      files["root"],
	}

	files["events"] = &File{
		Name:        "events.vcf",
		Size:        98535,
		Parent:      files["Documents"],
		Owner:       gowthamn,
		Type:        "File",
		Permissions: FullPermission,
	}

	for _, v := range files {
		fmt.Println(v.String())
		//fmt.Println("---------------------------")
	}

}

func (f File) String() string {
	var file string

	if f.Type == "Directory" {
		file += "d"
	} else {
		file += "-"
	}

	file += f.Permissions.String()

	file += "\t"
	file += strconv.Itoa(f.Size)

	file += "\t"
	file += f.Owner.Name

	file += "\t"
	file += f.Owner.Group.Name

	file += "\t"
	file += f.Name

	file += "\t"
	file += f.GetParentName()

	return file
}

func (f File) GetParentName() string {
	if f.Parent == nil {
		return ""
	}
	return f.Parent.Name
}

func (p Permissions) String() string {
	var permission string

	permission += p.Owner.String()
	permission += p.Group.String()
	permission += p.Others.String()

	return permission
}

func (p Permission) String() string {
	var permission string

	if (p).Read {
		permission += "r"
	} else {
		permission += "-"
	}

	if (p).Write {
		permission += "w"
	} else {
		permission += "-"
	}

	if (p).Execute {
		permission += "x"
	} else {
		permission += "-"
	}

	return permission
}
