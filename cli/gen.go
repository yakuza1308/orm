package main 

import (
    "errors"
    "strings"
    "github.com/eaciit/toolkit"
)

type StructModel struct{
    PkgName string
    Name string
    TableName string
    Fields []FieldModel
}

type FieldModel struct{
    Name string
    FieldType string
    FieldDefault string
    FieldTag string
}

func (sm *StructModel) Write(path string)error{
    if sm.PkgName=="" || sm.Name=="" {
        return toolkit.Errorf("Both package name and struct name should be defined")
    }
    return toolkit.Errorf("Fail to write %s.%s : method Write is not yet implemented", sm.PkgName, sm.Name)
}

func getPackage(txt string)(string,error){
    if !strings.HasPrefix(txt,"package"){
        return "",errors.New("No package has been defined") 
    }
    packages := strings.Split(txt," ")
    if len(packages)<2{
        return "",errors.New("No package has been defined")
    }
    return packages[1],nil
}

func getStructName(s string)string{
    txts := strings.Split(s," ")
    hasStruct := false
    for _,txt:=range txts{
        if !hasStruct && txt=="struct"{
            hasStruct=true
        } else if hasStruct && txt!="" {
            return txt
        }
    }
    return ""
}

func (sm *StructModel) makeGetFn(s string)error{
    return nil
}

func (sm *StructModel) makeFindFn(s string)error{
    return nil
}

func (sm *StructModel) makeComment(s string)error{
    return nil
}