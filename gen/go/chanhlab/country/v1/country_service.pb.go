// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: chanhlab/country/v1/country_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_chanhlab_country_v1_country_service_proto protoreflect.FileDescriptor

var file_chanhlab_country_v1_country_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x68, 0x61,
	0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x1a, 0x21, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x9a, 0x05, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x79, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x26, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x6e,
	0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x80,
	0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x85, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x1a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x82, 0x01, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x29, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c, 0x61,
	0x62, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x3a,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x68, 0x6c,
	0x61, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_chanhlab_country_v1_country_service_proto_goTypes = []interface{}{
	(*ListCountriesRequest)(nil),  // 0: chanhlab.country.v1.ListCountriesRequest
	(*GetCountryRequest)(nil),     // 1: chanhlab.country.v1.GetCountryRequest
	(*CreateCountryRequest)(nil),  // 2: chanhlab.country.v1.CreateCountryRequest
	(*UpdateCountryRequest)(nil),  // 3: chanhlab.country.v1.UpdateCountryRequest
	(*DeleteCountryRequest)(nil),  // 4: chanhlab.country.v1.DeleteCountryRequest
	(*ListCountriesResponse)(nil), // 5: chanhlab.country.v1.ListCountriesResponse
	(*GetCountryResponse)(nil),    // 6: chanhlab.country.v1.GetCountryResponse
	(*CreateCountryResponse)(nil), // 7: chanhlab.country.v1.CreateCountryResponse
	(*UpdateCountryResponse)(nil), // 8: chanhlab.country.v1.UpdateCountryResponse
	(*DeleteCountryResponse)(nil), // 9: chanhlab.country.v1.DeleteCountryResponse
}
var file_chanhlab_country_v1_country_service_proto_depIdxs = []int32{
	0, // 0: chanhlab.country.v1.CountryService.ListCountries:input_type -> chanhlab.country.v1.ListCountriesRequest
	1, // 1: chanhlab.country.v1.CountryService.GetCountry:input_type -> chanhlab.country.v1.GetCountryRequest
	2, // 2: chanhlab.country.v1.CountryService.CreateCountry:input_type -> chanhlab.country.v1.CreateCountryRequest
	3, // 3: chanhlab.country.v1.CountryService.UpdateCountry:input_type -> chanhlab.country.v1.UpdateCountryRequest
	4, // 4: chanhlab.country.v1.CountryService.DeleteCountry:input_type -> chanhlab.country.v1.DeleteCountryRequest
	5, // 5: chanhlab.country.v1.CountryService.ListCountries:output_type -> chanhlab.country.v1.ListCountriesResponse
	6, // 6: chanhlab.country.v1.CountryService.GetCountry:output_type -> chanhlab.country.v1.GetCountryResponse
	7, // 7: chanhlab.country.v1.CountryService.CreateCountry:output_type -> chanhlab.country.v1.CreateCountryResponse
	8, // 8: chanhlab.country.v1.CountryService.UpdateCountry:output_type -> chanhlab.country.v1.UpdateCountryResponse
	9, // 9: chanhlab.country.v1.CountryService.DeleteCountry:output_type -> chanhlab.country.v1.DeleteCountryResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chanhlab_country_v1_country_service_proto_init() }
func file_chanhlab_country_v1_country_service_proto_init() {
	if File_chanhlab_country_v1_country_service_proto != nil {
		return
	}
	file_chanhlab_country_v1_country_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chanhlab_country_v1_country_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chanhlab_country_v1_country_service_proto_goTypes,
		DependencyIndexes: file_chanhlab_country_v1_country_service_proto_depIdxs,
	}.Build()
	File_chanhlab_country_v1_country_service_proto = out.File
	file_chanhlab_country_v1_country_service_proto_rawDesc = nil
	file_chanhlab_country_v1_country_service_proto_goTypes = nil
	file_chanhlab_country_v1_country_service_proto_depIdxs = nil
}
