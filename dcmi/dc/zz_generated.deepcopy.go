//go:build !ignore_autogenerated

/*
Copyright 2021-2024 The nagare media authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package dc

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Any) DeepCopyInto(out *Any) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Any.
func (in *Any) DeepCopy() *Any {
	if in == nil {
		return nil
	}
	out := new(Any)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Contributor) DeepCopyInto(out *Contributor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Contributor.
func (in *Contributor) DeepCopy() *Contributor {
	if in == nil {
		return nil
	}
	out := new(Contributor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Coverage) DeepCopyInto(out *Coverage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Coverage.
func (in *Coverage) DeepCopy() *Coverage {
	if in == nil {
		return nil
	}
	out := new(Coverage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Creator) DeepCopyInto(out *Creator) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Creator.
func (in *Creator) DeepCopy() *Creator {
	if in == nil {
		return nil
	}
	out := new(Creator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Date) DeepCopyInto(out *Date) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Date.
func (in *Date) DeepCopy() *Date {
	if in == nil {
		return nil
	}
	out := new(Date)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Description) DeepCopyInto(out *Description) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Description.
func (in *Description) DeepCopy() *Description {
	if in == nil {
		return nil
	}
	out := new(Description)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elements) DeepCopyInto(out *Elements) {
	*out = *in
	out.XMLName = in.XMLName
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = make([]Title, len(*in))
		copy(*out, *in)
	}
	if in.Creator != nil {
		in, out := &in.Creator, &out.Creator
		*out = make([]Creator, len(*in))
		copy(*out, *in)
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = make([]Subject, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = make([]Description, len(*in))
		copy(*out, *in)
	}
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = make([]Publisher, len(*in))
		copy(*out, *in)
	}
	if in.Contributor != nil {
		in, out := &in.Contributor, &out.Contributor
		*out = make([]Contributor, len(*in))
		copy(*out, *in)
	}
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = make([]Date, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = make([]Type, len(*in))
		copy(*out, *in)
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = make([]Format, len(*in))
		copy(*out, *in)
	}
	if in.Identifier != nil {
		in, out := &in.Identifier, &out.Identifier
		*out = make([]Identifier, len(*in))
		copy(*out, *in)
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]Source, len(*in))
		copy(*out, *in)
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = make([]Language, len(*in))
		copy(*out, *in)
	}
	if in.Relation != nil {
		in, out := &in.Relation, &out.Relation
		*out = make([]Relation, len(*in))
		copy(*out, *in)
	}
	if in.Coverage != nil {
		in, out := &in.Coverage, &out.Coverage
		*out = make([]Coverage, len(*in))
		copy(*out, *in)
	}
	if in.Rights != nil {
		in, out := &in.Rights, &out.Rights
		*out = make([]Rights, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elements.
func (in *Elements) DeepCopy() *Elements {
	if in == nil {
		return nil
	}
	out := new(Elements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Format) DeepCopyInto(out *Format) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Format.
func (in *Format) DeepCopy() *Format {
	if in == nil {
		return nil
	}
	out := new(Format)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identifier) DeepCopyInto(out *Identifier) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identifier.
func (in *Identifier) DeepCopy() *Identifier {
	if in == nil {
		return nil
	}
	out := new(Identifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Language) DeepCopyInto(out *Language) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Language.
func (in *Language) DeepCopy() *Language {
	if in == nil {
		return nil
	}
	out := new(Language)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Publisher) DeepCopyInto(out *Publisher) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Publisher.
func (in *Publisher) DeepCopy() *Publisher {
	if in == nil {
		return nil
	}
	out := new(Publisher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Relation) DeepCopyInto(out *Relation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Relation.
func (in *Relation) DeepCopy() *Relation {
	if in == nil {
		return nil
	}
	out := new(Relation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rights) DeepCopyInto(out *Rights) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rights.
func (in *Rights) DeepCopy() *Rights {
	if in == nil {
		return nil
	}
	out := new(Rights)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimpleLiteral) DeepCopyInto(out *SimpleLiteral) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimpleLiteral.
func (in *SimpleLiteral) DeepCopy() *SimpleLiteral {
	if in == nil {
		return nil
	}
	out := new(SimpleLiteral)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subject) DeepCopyInto(out *Subject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subject.
func (in *Subject) DeepCopy() *Subject {
	if in == nil {
		return nil
	}
	out := new(Subject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Title) DeepCopyInto(out *Title) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Title.
func (in *Title) DeepCopy() *Title {
	if in == nil {
		return nil
	}
	out := new(Title)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Type) DeepCopyInto(out *Type) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Type.
func (in *Type) DeepCopy() *Type {
	if in == nil {
		return nil
	}
	out := new(Type)
	in.DeepCopyInto(out)
	return out
}