/*
Endringsservice

Dette er et api som har i ansvar å ta i mot endringer fra diverse kilder som mener noe om produksjonsplasser

API version: 1.0.0
Contact: no-spam@mattilsynet.no
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endringsservice-api

import (
	"encoding/json"
)

// Produksjonsplass struct for Produksjonsplass
type Produksjonsplass struct {
	// en identifikasjon som du bruker for å finne fram denne enheten/identiteten i ditt system, også kalt identifikasjonsnummer i andre systemer
	Identifikasjonsnummer string `json:"identifikasjonsnummer"`
	// for å kunne støtte flere kilder må man kunne gjennbruke informasjonsnummer for hver kilde, også kalt identifikasjontypeKode i andre systemer
	Identifikasjonstypekode string `json:"identifikasjonstypekode"`
	// string representasjon av SDO_UTIL.TO_WKTGEOMETRY
	Point *string `json:"point,omitempty"`
	// koordinatsystem som skal brukes for å representere point, mest typisk er 25833
	Koordinatsystem *string `json:"koordinatsystem,omitempty"`
	// hvor dataene fra, ikke hvor kilden er, men f.eks mats selvom kilden er matrikkel
	GeomMaster *string `json:"geomMaster,omitempty"`
	// hvor kommer dataene opprinnelig fra, i.e., matrikkel, slf, brreg, osv
	GeomKilde *string `json:"geomKilde,omitempty"`
}

// NewProduksjonsplass instantiates a new Produksjonsplass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduksjonsplass(identifikasjonsnummer string, identifikasjonstypekode string) *Produksjonsplass {
	this := Produksjonsplass{}
	this.Identifikasjonsnummer = identifikasjonsnummer
	this.Identifikasjonstypekode = identifikasjonstypekode
	return &this
}

// NewProduksjonsplassWithDefaults instantiates a new Produksjonsplass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduksjonsplassWithDefaults() *Produksjonsplass {
	this := Produksjonsplass{}
	return &this
}

// GetIdentifikasjonsnummer returns the Identifikasjonsnummer field value
func (o *Produksjonsplass) GetIdentifikasjonsnummer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifikasjonsnummer
}

// GetIdentifikasjonsnummerOk returns a tuple with the Identifikasjonsnummer field value
// and a boolean to check if the value has been set.
func (o *Produksjonsplass) GetIdentifikasjonsnummerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Identifikasjonsnummer, true
}

// SetIdentifikasjonsnummer sets field value
func (o *Produksjonsplass) SetIdentifikasjonsnummer(v string) {
	o.Identifikasjonsnummer = v
}

// GetIdentifikasjonstypekode returns the Identifikasjonstypekode field value
func (o *Produksjonsplass) GetIdentifikasjonstypekode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifikasjonstypekode
}

// GetIdentifikasjonstypekodeOk returns a tuple with the Identifikasjonstypekode field value
// and a boolean to check if the value has been set.
func (o *Produksjonsplass) GetIdentifikasjonstypekodeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Identifikasjonstypekode, true
}

// SetIdentifikasjonstypekode sets field value
func (o *Produksjonsplass) SetIdentifikasjonstypekode(v string) {
	o.Identifikasjonstypekode = v
}

// GetPoint returns the Point field value if set, zero value otherwise.
func (o *Produksjonsplass) GetPoint() string {
	if o == nil || isNil(o.Point) {
		var ret string
		return ret
	}
	return *o.Point
}

// GetPointOk returns a tuple with the Point field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Produksjonsplass) GetPointOk() (*string, bool) {
	if o == nil || isNil(o.Point) {
    return nil, false
	}
	return o.Point, true
}

// HasPoint returns a boolean if a field has been set.
func (o *Produksjonsplass) HasPoint() bool {
	if o != nil && !isNil(o.Point) {
		return true
	}

	return false
}

// SetPoint gets a reference to the given string and assigns it to the Point field.
func (o *Produksjonsplass) SetPoint(v string) {
	o.Point = &v
}

// GetKoordinatsystem returns the Koordinatsystem field value if set, zero value otherwise.
func (o *Produksjonsplass) GetKoordinatsystem() string {
	if o == nil || isNil(o.Koordinatsystem) {
		var ret string
		return ret
	}
	return *o.Koordinatsystem
}

// GetKoordinatsystemOk returns a tuple with the Koordinatsystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Produksjonsplass) GetKoordinatsystemOk() (*string, bool) {
	if o == nil || isNil(o.Koordinatsystem) {
    return nil, false
	}
	return o.Koordinatsystem, true
}

// HasKoordinatsystem returns a boolean if a field has been set.
func (o *Produksjonsplass) HasKoordinatsystem() bool {
	if o != nil && !isNil(o.Koordinatsystem) {
		return true
	}

	return false
}

// SetKoordinatsystem gets a reference to the given string and assigns it to the Koordinatsystem field.
func (o *Produksjonsplass) SetKoordinatsystem(v string) {
	o.Koordinatsystem = &v
}

// GetGeomMaster returns the GeomMaster field value if set, zero value otherwise.
func (o *Produksjonsplass) GetGeomMaster() string {
	if o == nil || isNil(o.GeomMaster) {
		var ret string
		return ret
	}
	return *o.GeomMaster
}

// GetGeomMasterOk returns a tuple with the GeomMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Produksjonsplass) GetGeomMasterOk() (*string, bool) {
	if o == nil || isNil(o.GeomMaster) {
    return nil, false
	}
	return o.GeomMaster, true
}

// HasGeomMaster returns a boolean if a field has been set.
func (o *Produksjonsplass) HasGeomMaster() bool {
	if o != nil && !isNil(o.GeomMaster) {
		return true
	}

	return false
}

// SetGeomMaster gets a reference to the given string and assigns it to the GeomMaster field.
func (o *Produksjonsplass) SetGeomMaster(v string) {
	o.GeomMaster = &v
}

// GetGeomKilde returns the GeomKilde field value if set, zero value otherwise.
func (o *Produksjonsplass) GetGeomKilde() string {
	if o == nil || isNil(o.GeomKilde) {
		var ret string
		return ret
	}
	return *o.GeomKilde
}

// GetGeomKildeOk returns a tuple with the GeomKilde field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Produksjonsplass) GetGeomKildeOk() (*string, bool) {
	if o == nil || isNil(o.GeomKilde) {
    return nil, false
	}
	return o.GeomKilde, true
}

// HasGeomKilde returns a boolean if a field has been set.
func (o *Produksjonsplass) HasGeomKilde() bool {
	if o != nil && !isNil(o.GeomKilde) {
		return true
	}

	return false
}

// SetGeomKilde gets a reference to the given string and assigns it to the GeomKilde field.
func (o *Produksjonsplass) SetGeomKilde(v string) {
	o.GeomKilde = &v
}

func (o Produksjonsplass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identifikasjonsnummer"] = o.Identifikasjonsnummer
	}
	if true {
		toSerialize["identifikasjonstypekode"] = o.Identifikasjonstypekode
	}
	if !isNil(o.Point) {
		toSerialize["point"] = o.Point
	}
	if !isNil(o.Koordinatsystem) {
		toSerialize["koordinatsystem"] = o.Koordinatsystem
	}
	if !isNil(o.GeomMaster) {
		toSerialize["geomMaster"] = o.GeomMaster
	}
	if !isNil(o.GeomKilde) {
		toSerialize["geomKilde"] = o.GeomKilde
	}
	return json.Marshal(toSerialize)
}

type NullableProduksjonsplass struct {
	value *Produksjonsplass
	isSet bool
}

func (v NullableProduksjonsplass) Get() *Produksjonsplass {
	return v.value
}

func (v *NullableProduksjonsplass) Set(val *Produksjonsplass) {
	v.value = val
	v.isSet = true
}

func (v NullableProduksjonsplass) IsSet() bool {
	return v.isSet
}

func (v *NullableProduksjonsplass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduksjonsplass(val *Produksjonsplass) *NullableProduksjonsplass {
	return &NullableProduksjonsplass{value: val, isSet: true}
}

func (v NullableProduksjonsplass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduksjonsplass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


