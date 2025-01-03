// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package hanariu

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/hanariu/hanariu/pkg/hanariu"
)

type FormInputBosons struct {
	Size         string `default:"md"`
	Autocomplete string `default:"off"`
	Type         string `default:"text"`
	Attrs        hanariu.Attrs
}

type FormInputFieldBosons struct {
	Error bool
	Attrs hanariu.Attrs
}

var formInputFieldStyles = hanariu.Boson{
	"default": "appearance-none flex w-full shadow-sm focus-within:outline-none border focus-within:border-1 focus-within:ring-1 rounded-md focus-within:rounded-md text-secondary-700 dark:text-primary-200 dark:bg-primary-900/40",
	"noerror": "border-secondary-400/80 dark:border-primary-700/80 focus-within:ring-primary-500 focus-within:border-primary-500",
	"error":   "border-danger-500 dark:border-danger-400 focus-within:ring-danger-500 focus-within:border-danger-500 dark:focus-within:ring-danger-500/80 dark:focus-within:border-danger-400/80",
}

var formInputStyles = hanariu.Boson{
	"input": "w-full rounded-md border-none placeholder-secondary-400 dark:placeholder-primary-800 focus:border-transparent focus:outline-none focus:ring-0 bg-transparent",
}

var formInputSizeStyles = hanariu.Boson{
	"sm": "px-2 py-2 text-base md:px-2.5 md:py-1.5 md:text-sm",
	"md": "px-3 py-2.5 text-lg md:px-3 md:py-2 md:text-base",
	"lg": "px-4 py-3 text-xl md:px-4 md:py-3 md:text-lg",
}

func (attr *FormInputBosons) makeFormInputAttrs() {
	attrs := make(hanariu.Attrs)
	for k, v := range attr.Attrs {
		attrs[k] = v
	}
	attr.Attrs = attrs
}

func (attr *FormInputBosons) addFormInputAttr(k, v string) {
	attr.Attrs[k] = v
}
func FormInputField(props *FormInputFieldBosons) templ.Component {
	classes := getFormInputFieldClasses(props.Error)
	return hanariu.CreateComponent("div", classes, props.Attrs, false)
}

func FormInput(props *FormInputBosons) templ.Component {
	props.Size = hanariu.GetTagDefault("Size", props.Size, FormInputBosons{})
	classes := getFormInputClasses(props.Size)
	props.makeFormInputAttrs()
	props.addFormInputAttr("type", hanariu.GetTagDefault("Type", props.Type, FormInputBosons{}))
	props.addFormInputAttr("autocomplete", hanariu.GetTagDefault("Autocomplete", props.Autocomplete, FormInputBosons{}))
	return hanariu.CreateComponent("input", classes, props.Attrs, true)
}

func getFormInputClasses(size string) string {
	return formInputStyles["input"] + " " + formInputSizeStyles[size]
}

func getFormInputFieldClasses(err bool) string {
	var formInputFieldClasses = formInputFieldStyles["default"]
	if err {
		formInputFieldClasses += " " + formInputFieldStyles["error"]
	} else {
		formInputFieldClasses += " " + formInputFieldStyles["noerror"]
	}
	return formInputFieldClasses
}

var _ = templruntime.GeneratedTemplate
