package goncrete

import "git.apache.org/thrift.git/lib/go/thrift"

// DefaultTransportFactory returns the default implementation of
// thrift.TTransportFactory for goncrete projects,
// a buffered transport wrapped by a framed transport
func DefaultTransportFactory() thrift.TTransportFactory {
	return thrift.NewTFramedTransportFactory(thrift.NewTBufferedTransportFactory(8192))
}

// DefaultProtocolFactory returns the default implementation of
// thrift.TProtocolFactory for goncrete projects,
// the CompactProtocol
func DefaultProtocolFactory() thrift.TProtocolFactory {
	return thrift.NewTCompactProtocolFactory()
}
