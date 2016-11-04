// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/servicecontrol/v1/distribution.proto
// DO NOT EDIT!

package google_api_servicecontrol_v1 // import "google.golang.org/genproto/googleapis/api/servicecontrol/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Distribution represents a frequency distribution of double-valued sample
// points. It contains the size of the population of sample points plus
// additional optional information:
//
//   - the arithmetic mean of the samples
//   - the minimum and maximum of the samples
//   - the sum-squared-deviation of the samples, used to compute variance
//   - a histogram of the values of the sample points
type Distribution struct {
	// The total number of samples in the distribution. Must be >= 0.
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	// The arithmetic mean of the samples in the distribution. If `count` is
	// zero then this field must be zero.
	Mean float64 `protobuf:"fixed64,2,opt,name=mean" json:"mean,omitempty"`
	// The minimum of the population of values. Ignored if `count` is zero.
	Minimum float64 `protobuf:"fixed64,3,opt,name=minimum" json:"minimum,omitempty"`
	// The maximum of the population of values. Ignored if `count` is zero.
	Maximum float64 `protobuf:"fixed64,4,opt,name=maximum" json:"maximum,omitempty"`
	// The sum of squared deviations from the mean:
	//   Sum[i=1..count]((x_i - mean)^2)
	// where each x_i is a sample values. If `count` is zero then this field
	// must be zero, otherwise validation of the request fails.
	SumOfSquaredDeviation float64 `protobuf:"fixed64,5,opt,name=sum_of_squared_deviation,json=sumOfSquaredDeviation" json:"sum_of_squared_deviation,omitempty"`
	// The number of samples in each histogram bucket. `bucket_counts` are
	// optional. If present, they must sum to the `count` value.
	//
	// The buckets are defined below in `bucket_option`. There are N buckets.
	// `bucket_counts[0]` is the number of samples in the underflow bucket.
	// `bucket_counts[1]` to `bucket_counts[N-1]` are the numbers of samples
	// in each of the finite buckets. And `bucket_counts[N] is the number
	// of samples in the overflow bucket. See the comments of `bucket_option`
	// below for more details.
	//
	// Any suffix of trailing zeros may be omitted.
	BucketCounts []int64 `protobuf:"varint,6,rep,packed,name=bucket_counts,json=bucketCounts" json:"bucket_counts,omitempty"`
	// Defines the buckets in the histogram. `bucket_option` and `bucket_counts`
	// must be both set, or both unset.
	//
	// Buckets are numbered the the range of [0, N], with a total of N+1 buckets.
	// There must be at least two buckets (a single-bucket histogram gives
	// no information that isn't already provided by `count`).
	//
	// The first bucket is the underflow bucket which has a lower bound
	// of -inf. The last bucket is the overflow bucket which has an
	// upper bound of +inf. All other buckets (if any) are called "finite"
	// buckets because they have finite lower and upper bounds. As described
	// below, there are three ways to define the finite buckets.
	//
	//   (1) Buckets with constant width.
	//   (2) Buckets with exponentially growing widths.
	//   (3) Buckets with arbitrary user-provided widths.
	//
	// In all cases, the buckets cover the entire real number line (-inf,
	// +inf). Bucket upper bounds are exclusive and lower bounds are
	// inclusive. The upper bound of the underflow bucket is equal to the
	// lower bound of the smallest finite bucket; the lower bound of the
	// overflow bucket is equal to the upper bound of the largest finite
	// bucket.
	//
	// Types that are valid to be assigned to BucketOption:
	//	*Distribution_LinearBuckets_
	//	*Distribution_ExponentialBuckets_
	//	*Distribution_ExplicitBuckets_
	BucketOption isDistribution_BucketOption `protobuf_oneof:"bucket_option"`
}

func (m *Distribution) Reset()                    { *m = Distribution{} }
func (m *Distribution) String() string            { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()               {}
func (*Distribution) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isDistribution_BucketOption interface {
	isDistribution_BucketOption()
}

type Distribution_LinearBuckets_ struct {
	LinearBuckets *Distribution_LinearBuckets `protobuf:"bytes,7,opt,name=linear_buckets,json=linearBuckets,oneof"`
}
type Distribution_ExponentialBuckets_ struct {
	ExponentialBuckets *Distribution_ExponentialBuckets `protobuf:"bytes,8,opt,name=exponential_buckets,json=exponentialBuckets,oneof"`
}
type Distribution_ExplicitBuckets_ struct {
	ExplicitBuckets *Distribution_ExplicitBuckets `protobuf:"bytes,9,opt,name=explicit_buckets,json=explicitBuckets,oneof"`
}

func (*Distribution_LinearBuckets_) isDistribution_BucketOption()      {}
func (*Distribution_ExponentialBuckets_) isDistribution_BucketOption() {}
func (*Distribution_ExplicitBuckets_) isDistribution_BucketOption()    {}

func (m *Distribution) GetBucketOption() isDistribution_BucketOption {
	if m != nil {
		return m.BucketOption
	}
	return nil
}

func (m *Distribution) GetLinearBuckets() *Distribution_LinearBuckets {
	if x, ok := m.GetBucketOption().(*Distribution_LinearBuckets_); ok {
		return x.LinearBuckets
	}
	return nil
}

func (m *Distribution) GetExponentialBuckets() *Distribution_ExponentialBuckets {
	if x, ok := m.GetBucketOption().(*Distribution_ExponentialBuckets_); ok {
		return x.ExponentialBuckets
	}
	return nil
}

func (m *Distribution) GetExplicitBuckets() *Distribution_ExplicitBuckets {
	if x, ok := m.GetBucketOption().(*Distribution_ExplicitBuckets_); ok {
		return x.ExplicitBuckets
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Distribution) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Distribution_OneofMarshaler, _Distribution_OneofUnmarshaler, _Distribution_OneofSizer, []interface{}{
		(*Distribution_LinearBuckets_)(nil),
		(*Distribution_ExponentialBuckets_)(nil),
		(*Distribution_ExplicitBuckets_)(nil),
	}
}

func _Distribution_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Distribution)
	// bucket_option
	switch x := m.BucketOption.(type) {
	case *Distribution_LinearBuckets_:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LinearBuckets); err != nil {
			return err
		}
	case *Distribution_ExponentialBuckets_:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExponentialBuckets); err != nil {
			return err
		}
	case *Distribution_ExplicitBuckets_:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExplicitBuckets); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Distribution.BucketOption has unexpected type %T", x)
	}
	return nil
}

func _Distribution_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Distribution)
	switch tag {
	case 7: // bucket_option.linear_buckets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Distribution_LinearBuckets)
		err := b.DecodeMessage(msg)
		m.BucketOption = &Distribution_LinearBuckets_{msg}
		return true, err
	case 8: // bucket_option.exponential_buckets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Distribution_ExponentialBuckets)
		err := b.DecodeMessage(msg)
		m.BucketOption = &Distribution_ExponentialBuckets_{msg}
		return true, err
	case 9: // bucket_option.explicit_buckets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Distribution_ExplicitBuckets)
		err := b.DecodeMessage(msg)
		m.BucketOption = &Distribution_ExplicitBuckets_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Distribution_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Distribution)
	// bucket_option
	switch x := m.BucketOption.(type) {
	case *Distribution_LinearBuckets_:
		s := proto.Size(x.LinearBuckets)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Distribution_ExponentialBuckets_:
		s := proto.Size(x.ExponentialBuckets)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Distribution_ExplicitBuckets_:
		s := proto.Size(x.ExplicitBuckets)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Describing buckets with constant width.
type Distribution_LinearBuckets struct {
	// The number of finite buckets. With the underflow and overflow buckets,
	// the total number of buckets is `num_finite_buckets` + 2.
	// See comments on `bucket_options` for details.
	NumFiniteBuckets int32 `protobuf:"varint,1,opt,name=num_finite_buckets,json=numFiniteBuckets" json:"num_finite_buckets,omitempty"`
	// The i'th linear bucket covers the interval
	//   [offset + (i-1) * width, offset + i * width)
	// where i ranges from 1 to num_finite_buckets, inclusive.
	// Must be strictly positive.
	Width float64 `protobuf:"fixed64,2,opt,name=width" json:"width,omitempty"`
	// The i'th linear bucket covers the interval
	//   [offset + (i-1) * width, offset + i * width)
	// where i ranges from 1 to num_finite_buckets, inclusive.
	Offset float64 `protobuf:"fixed64,3,opt,name=offset" json:"offset,omitempty"`
}

func (m *Distribution_LinearBuckets) Reset()                    { *m = Distribution_LinearBuckets{} }
func (m *Distribution_LinearBuckets) String() string            { return proto.CompactTextString(m) }
func (*Distribution_LinearBuckets) ProtoMessage()               {}
func (*Distribution_LinearBuckets) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// Describing buckets with exponentially growing width.
type Distribution_ExponentialBuckets struct {
	// The number of finite buckets. With the underflow and overflow buckets,
	// the total number of buckets is `num_finite_buckets` + 2.
	// See comments on `bucket_options` for details.
	NumFiniteBuckets int32 `protobuf:"varint,1,opt,name=num_finite_buckets,json=numFiniteBuckets" json:"num_finite_buckets,omitempty"`
	// The i'th exponential bucket covers the interval
	//   [scale * growth_factor^(i-1), scale * growth_factor^i)
	// where i ranges from 1 to num_finite_buckets inclusive.
	// Must be larger than 1.0.
	GrowthFactor float64 `protobuf:"fixed64,2,opt,name=growth_factor,json=growthFactor" json:"growth_factor,omitempty"`
	// The i'th exponential bucket covers the interval
	//   [scale * growth_factor^(i-1), scale * growth_factor^i)
	// where i ranges from 1 to num_finite_buckets inclusive.
	// Must be > 0.
	Scale float64 `protobuf:"fixed64,3,opt,name=scale" json:"scale,omitempty"`
}

func (m *Distribution_ExponentialBuckets) Reset()         { *m = Distribution_ExponentialBuckets{} }
func (m *Distribution_ExponentialBuckets) String() string { return proto.CompactTextString(m) }
func (*Distribution_ExponentialBuckets) ProtoMessage()    {}
func (*Distribution_ExponentialBuckets) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 1}
}

// Describing buckets with arbitrary user-provided width.
type Distribution_ExplicitBuckets struct {
	// 'bound' is a list of strictly increasing boundaries between
	// buckets. Note that a list of length N-1 defines N buckets because
	// of fenceposting. See comments on `bucket_options` for details.
	//
	// The i'th finite bucket covers the interval
	//   [bound[i-1], bound[i])
	// where i ranges from 1 to bound_size() - 1. Note that there are no
	// finite buckets at all if 'bound' only contains a single element; in
	// that special case the single bound defines the boundary between the
	// underflow and overflow buckets.
	//
	// bucket number                   lower bound    upper bound
	//  i == 0 (underflow)              -inf           bound[i]
	//  0 < i < bound_size()            bound[i-1]     bound[i]
	//  i == bound_size() (overflow)    bound[i-1]     +inf
	Bounds []float64 `protobuf:"fixed64,1,rep,packed,name=bounds" json:"bounds,omitempty"`
}

func (m *Distribution_ExplicitBuckets) Reset()                    { *m = Distribution_ExplicitBuckets{} }
func (m *Distribution_ExplicitBuckets) String() string            { return proto.CompactTextString(m) }
func (*Distribution_ExplicitBuckets) ProtoMessage()               {}
func (*Distribution_ExplicitBuckets) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 2} }

func init() {
	proto.RegisterType((*Distribution)(nil), "google.api.servicecontrol.v1.Distribution")
	proto.RegisterType((*Distribution_LinearBuckets)(nil), "google.api.servicecontrol.v1.Distribution.LinearBuckets")
	proto.RegisterType((*Distribution_ExponentialBuckets)(nil), "google.api.servicecontrol.v1.Distribution.ExponentialBuckets")
	proto.RegisterType((*Distribution_ExplicitBuckets)(nil), "google.api.servicecontrol.v1.Distribution.ExplicitBuckets")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/servicecontrol/v1/distribution.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x31, 0x6e, 0x52, 0x18, 0x1a, 0x52, 0x96, 0x82, 0xac, 0x88, 0x83, 0x45, 0x2f, 0x41,
	0x42, 0xb6, 0x0a, 0x07, 0x50, 0x25, 0x2e, 0xa1, 0x54, 0x1c, 0x10, 0x54, 0xe6, 0x01, 0xac, 0x8d,
	0xbd, 0x76, 0x57, 0xf5, 0xee, 0x98, 0xfd, 0x93, 0xe6, 0xc2, 0x3b, 0xf1, 0x78, 0x1c, 0x91, 0x77,
	0x9d, 0x7f, 0x54, 0xaa, 0x94, 0x9b, 0xbf, 0xef, 0xf3, 0xcc, 0x6f, 0x3c, 0x1a, 0xc3, 0xf7, 0x1a,
	0xb1, 0x6e, 0x58, 0x52, 0x63, 0x43, 0x65, 0x9d, 0xa0, 0xaa, 0xd3, 0x9a, 0xc9, 0x56, 0xa1, 0xc1,
	0xd4, 0x47, 0xb4, 0xe5, 0x3a, 0xa5, 0x2d, 0x4f, 0x35, 0x53, 0x0b, 0x5e, 0xb0, 0x02, 0xa5, 0x51,
	0xd8, 0xa4, 0x8b, 0xb3, 0xb4, 0xe4, 0xda, 0x28, 0x3e, 0xb7, 0x86, 0xa3, 0x4c, 0x5c, 0x0d, 0x79,
	0xd5, 0xf7, 0xa3, 0x2d, 0x4f, 0x76, 0x0b, 0x92, 0xc5, 0xd9, 0xeb, 0x3f, 0x43, 0x38, 0xba, 0xd8,
	0x2a, 0x22, 0x27, 0x30, 0x28, 0xd0, 0x4a, 0x13, 0x05, 0x71, 0x30, 0x0d, 0x33, 0x2f, 0x08, 0x81,
	0x03, 0xc1, 0xa8, 0x8c, 0x1e, 0xc6, 0xc1, 0x34, 0xc8, 0xdc, 0x33, 0x89, 0xe0, 0x50, 0x70, 0xc9,
	0x85, 0x15, 0x51, 0xe8, 0xec, 0x95, 0x74, 0x09, 0x5d, 0xba, 0xe4, 0xa0, 0x4f, 0xbc, 0x24, 0x1f,
	0x20, 0xd2, 0x56, 0xe4, 0x58, 0xe5, 0xfa, 0x97, 0xa5, 0x8a, 0x95, 0x79, 0xc9, 0x16, 0x9c, 0x76,
	0xe4, 0x68, 0xe0, 0x5e, 0x7d, 0xa1, 0xad, 0xf8, 0x51, 0xfd, 0xf4, 0xe9, 0xc5, 0x2a, 0x24, 0xa7,
	0x30, 0x9a, 0xdb, 0xe2, 0x86, 0x99, 0xdc, 0x0d, 0xa4, 0xa3, 0x61, 0x1c, 0x4e, 0xc3, 0xec, 0xc8,
	0x9b, 0x9f, 0x9d, 0x47, 0x28, 0x3c, 0x6d, 0xb8, 0x64, 0x54, 0xe5, 0xde, 0xd6, 0xd1, 0x61, 0x1c,
	0x4c, 0x9f, 0xbc, 0xfb, 0x98, 0xdc, 0xb7, 0x83, 0x64, 0xfb, 0xfb, 0x93, 0x6f, 0xae, 0xc1, 0xcc,
	0xd7, 0x7f, 0x7d, 0x90, 0x8d, 0x9a, 0x6d, 0x83, 0xb4, 0xf0, 0x9c, 0x2d, 0x5b, 0x94, 0x4c, 0x1a,
	0x4e, 0x9b, 0x35, 0xe7, 0x91, 0xe3, 0x7c, 0xda, 0x83, 0xf3, 0x65, 0xd3, 0x65, 0x03, 0x23, 0xec,
	0x8e, 0x4b, 0x6a, 0x38, 0x66, 0xcb, 0xb6, 0xe1, 0x05, 0x37, 0x6b, 0xdc, 0x63, 0x87, 0x3b, 0xdf,
	0x0f, 0xe7, 0x5a, 0x6c, 0x58, 0x63, 0xb6, 0x6b, 0x4d, 0x6e, 0x60, 0xb4, 0xf3, 0xf1, 0xe4, 0x2d,
	0x10, 0x69, 0x45, 0x5e, 0x71, 0xc9, 0x0d, 0x5b, 0xb3, 0xbb, 0xbb, 0x18, 0x64, 0xc7, 0xd2, 0x8a,
	0x4b, 0x17, 0xac, 0xde, 0x3e, 0x81, 0xc1, 0x2d, 0x2f, 0xcd, 0x75, 0x7f, 0x23, 0x5e, 0x90, 0x97,
	0x30, 0xc4, 0xaa, 0xd2, 0xcc, 0xf4, 0x37, 0xd2, 0xab, 0xc9, 0x6f, 0x20, 0x77, 0x37, 0xb0, 0x27,
	0xf1, 0x14, 0x46, 0xb5, 0xc2, 0x5b, 0x73, 0x9d, 0x57, 0xb4, 0x30, 0xa8, 0x7a, 0xf2, 0x91, 0x37,
	0x2f, 0x9d, 0xd7, 0x8d, 0xa5, 0x0b, 0xda, 0xb0, 0x9e, 0xef, 0xc5, 0xe4, 0x0d, 0x8c, 0xff, 0xdb,
	0x48, 0x37, 0xe9, 0x1c, 0xad, 0x2c, 0x3b, 0x5e, 0xd8, 0x4d, 0xea, 0xd5, 0x6c, 0xbc, 0xbe, 0x3c,
	0x6c, 0xbb, 0x55, 0xce, 0xce, 0x21, 0x2e, 0x50, 0xdc, 0xbb, 0xfb, 0xd9, 0xb3, 0xed, 0xe5, 0x5f,
	0x75, 0xff, 0xe1, 0x55, 0xf0, 0x37, 0x08, 0xe6, 0x43, 0xf7, 0x4f, 0xbe, 0xff, 0x17, 0x00, 0x00,
	0xff, 0xff, 0xfb, 0xfd, 0x43, 0xc4, 0xe5, 0x03, 0x00, 0x00,
}