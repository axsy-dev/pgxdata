erb pgtype_array_type=Int2Array pgtype_element_type=Int2 go_array_types=[]int16,[]uint16 element_type_name=int2 text_null=NULL binary_format=true typed_array.go.erb > int2_array.go
erb pgtype_array_type=Int4Array pgtype_element_type=Int4 go_array_types=[]int32,[]uint32 element_type_name=int4 text_null=NULL binary_format=true typed_array.go.erb > int4_array.go
erb pgtype_array_type=Int8Array pgtype_element_type=Int8 go_array_types=[]int64,[]uint64 element_type_name=int8 text_null=NULL binary_format=true typed_array.go.erb > int8_array.go
erb pgtype_array_type=BoolArray pgtype_element_type=Bool go_array_types=[]bool element_type_name=bool text_null=NULL binary_format=true typed_array.go.erb > bool_array.go
erb pgtype_array_type=DateArray pgtype_element_type=Date go_array_types=[]time.Time element_type_name=date text_null=NULL binary_format=true typed_array.go.erb > date_array.go
erb pgtype_array_type=TimestamptzArray pgtype_element_type=Timestamptz go_array_types=[]time.Time element_type_name=timestamptz text_null=NULL binary_format=true typed_array.go.erb > timestamptz_array.go
erb pgtype_array_type=TimestampArray pgtype_element_type=Timestamp go_array_types=[]time.Time element_type_name=timestamp text_null=NULL binary_format=true typed_array.go.erb > timestamp_array.go
erb pgtype_array_type=Float4Array pgtype_element_type=Float4 go_array_types=[]float32 element_type_name=float4 text_null=NULL binary_format=true typed_array.go.erb > float4_array.go
erb pgtype_array_type=Float8Array pgtype_element_type=Float8 go_array_types=[]float64 element_type_name=float8 text_null=NULL binary_format=true typed_array.go.erb > float8_array.go
erb pgtype_array_type=InetArray pgtype_element_type=Inet go_array_types=[]*net.IPNet,[]net.IP element_type_name=inet text_null=NULL binary_format=true typed_array.go.erb > inet_array.go
erb pgtype_array_type=CidrArray pgtype_element_type=Cidr go_array_types=[]*net.IPNet,[]net.IP element_type_name=cidr text_null=NULL binary_format=true typed_array.go.erb > cidr_array.go
erb pgtype_array_type=TextArray pgtype_element_type=Text go_array_types=[]string element_type_name=text text_null='"NULL"' binary_format=true typed_array.go.erb > text_array.go
erb pgtype_array_type=VarcharArray pgtype_element_type=Varchar go_array_types=[]string element_type_name=varchar text_null='"NULL"' binary_format=true typed_array.go.erb > varchar_array.go
erb pgtype_array_type=ByteaArray pgtype_element_type=Bytea go_array_types=[][]byte element_type_name=bytea text_null=NULL binary_format=true typed_array.go.erb > bytea_array.go
erb pgtype_array_type=AclitemArray pgtype_element_type=Aclitem go_array_types=[]string element_type_name=aclitem text_null=NULL binary_format=false typed_array.go.erb > aclitem_array.go
erb pgtype_array_type=HstoreArray pgtype_element_type=Hstore go_array_types=[]map[string]string element_type_name=hstore text_null=NULL binary_format=true typed_array.go.erb > hstore_array.go
erb pgtype_array_type=NumericArray pgtype_element_type=Numeric go_array_types=[]float32,[]float64 element_type_name=numeric text_null=NULL binary_format=true typed_array.go.erb > numeric_array.go
goimports -w *_array.go
