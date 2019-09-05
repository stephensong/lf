/*
 * Copyright (c)2019 ZeroTier, Inc.
 *
 * Use of this software is governed by the Business Source License included
 * in the LICENSE.TXT file in the project's root directory.
 *
 * Change Date: 2023-01-01
 *
 * On the date above, in accordance with the Business Source License, use
 * of this software will be governed by version 2.0 of the Apache License.
 */
/****/

package lf

import "net"

// This contains the defaults for the Sol LF network, the global shared LF
// database for Earth and its neighbors in the Sol system. It should be good
// up to Kardashev type II civilization scale. Latency across interstellar
// links might be too high so a galactic civilization may require one data
// store per ~4 cubic light hours.

/*
{
  "ID": [23, 85, 34, 46, 124, 51, 168, 95, 201, 112, 89, 91, 250, 91, 70, 59, 42, 169, 53, 238, 62, 70, 190, 211, 59, 20, 20, 141, 227, 216, 141, 35],
  "AmendableFields": ["authcertificates", "seedpeers"],
  "Name": "Sol",
  "Contact": "https://www.zerotier.com/lf",
  "Comment": "Global Public LF Data Store",
  "AuthCertificates": "\bHDoiqf4e1SSCehz025bTodmFahEpF7m0ic0B3jnVUKJ4ZlORVDlHyWRAgCcWw0laC4Nkn8b8yYQ7e46bBebcUyTY062TSaPSBvBTJDHmbGN8ETo3Kt563VnKBtzOAn7uugwtjBq6ZjbVAWkhwkKhzGINaNxp7IK5PMeq1Sax3xK1vp1uc52BaRXkyPIHKiaEmgraKLw9NwMqeJhGnyzRXYm3n0YhvJSDcgvktcNU9amM3AUNThqDDOixWUBUhQFzegUjI5WxXNHj8P6fOiKDalMFP1xC1bVHCiEXWMpKPsLkGsatfcr5FOcTPeAIFmDSc0Fv9A0PDWY7uLIheD4pDS30z0gyfgXmH6HZ3eUNUQ7hW6rKbz3GzhsgqS9hjISg7huI9XGzXFZpJdfk9002H3z5i6J52LuXpcfd8GQ6sn5BC0HlgyWI7MIq0flIGGEkVRmgoZi5303mdq8PrGEKqBJLekSAIEDD4OGSfMBqwMzE8R1esFzCKX9CdnXnstFwANzrmuoxI93Uqk1cULQaU1lDwORAkf8JzKIs8uggNattBhF9V6YRXQ8hsOJaEBtuod2pLMyokqC11B2KWlQxjdbo1zpfbkbDX6xJ0RRREd5mCWXN9AD9uMijC9qdTt1OLTkMF9v0Nkbl6DFTKSaZ4t042FeVcHxtXZbOesmx2pCGdwFwxMuEYRO1i0e5zQfkEHSitoq1mz0LkGQ4qt41roGXGBvcrmbnLNlXSqW7nqhAhqQ8RJkBM9Wk0UWGrRmlFuQOM15EYTig4iL1IxSnSMu7Q00FaGey6nL97WkN9zxStTFGkeWaWCUDYbrW1w6NFlGeC0u479hBSrTg6HWs4vRnmXFU0WyKSuldAuANdE5ou2seZlt68G4BzP83WJDdhYMBHOrzsOKqGO6OKsz1CBEDDBEoeJpEaE4jHUgq8TwYUc2xpLgFQqvdWd7x0S6gYBveZAslXAEQcSGbu4MZaaFhTOAjGr9v1YfOnr3xkFJONPYO1wCWS0sBpSWAFfcLNIzyabWrKhTqa1LJY6Wc64CQqFXLx4knHg5R",
  "AuthRequired": false,
  "RecordMinLinks": 2,
  "RecordMaxValueSize": 1024,
  "RecordMaxTimeDrift": 60
}
*/

// SolNetworkID is the value of the ID field in the Sol network's genesis paramters.
var SolNetworkID = [32]byte{0x17, 0x55, 0x22, 0x2e, 0x7c, 0x33, 0xa8, 0x5f, 0xc9, 0x70, 0x59, 0x5b, 0xfa, 0x5b, 0x46, 0x3b, 0x2a, 0xa9, 0x35, 0xee, 0x3e, 0x46, 0xbe, 0xd3, 0x3b, 0x14, 0x14, 0x8d, 0xe3, 0xd8, 0x8d, 0x23}

// SolSeedPeers is an array of peers that can be contacted to start synchronizing nodes on Sol.
var SolSeedPeers = []Peer{
	Peer{
		IP:       net.ParseIP("185.180.13.82"), // Los Angeles
		Port:     9908,
		Identity: Base62Decode("FHZAb5KAmPKzUySSau9VDiugJaUtIJMW3yKZHcjyFyCEPiBQMk6QxlHiGoBEkHtQf"),
	},
	Peer{
		IP:       net.ParseIP("195.181.173.159"), // Amsterdam
		Port:     9908,
		Identity: Base62Decode("JZlvKdV12jt8rM3552LaThpTnsC9FfOZGRvGvbN3GImKnaw4wR1CCBu80k1ohc6H6"),
	},
	Peer{
		IP:       net.ParseIP("174.136.102.98"), // Los Angeles (ZeroTier office)
		Port:     9908,
		Identity: Base62Decode("Hcsqi4GP24UhaJL9poDM35k7KwvgvYzt1fMrYDr5EEAhTJ1ZnHu61xpDctypw66fh"),
	},
}

// SolDefaultNodeURLs is the default URL for clients to access Sol (servers operated by ZeroTier, Inc.)
var SolDefaultNodeURLs = []RemoteNode{RemoteNode("https://lf.zerotier.com")}

// SolGenesisRecords are the genesis records for initializing a Sol member node.
var SolGenesisRecords = []byte{0x81, 0x2, 0xc0, 0x9, 0x8b, 0x5c, 0x34, 0x97, 0xf4, 0xfd, 0xa2, 0x95, 0x16, 0xc2, 0x90, 0x54, 0x0, 0x99, 0xf3, 0xd9, 0x60, 0xdd, 0xde, 0xa7, 0x16, 0xaa, 0x79, 0x73, 0x8d, 0xa2, 0x87, 0xad, 0xce, 0xde, 0xc5, 0x73, 0xcd, 0xde, 0xa7, 0x8f, 0x32, 0x6c, 0xae, 0x4, 0xfc, 0x2a, 0x66, 0xc8, 0xf1, 0x71, 0xd3, 0x5b, 0xdf, 0x68, 0xdc, 0xf6, 0x5b, 0x97, 0x89, 0xa, 0xd, 0x44, 0x75, 0xc3, 0x32, 0x13, 0x5a, 0x50, 0x41, 0xfd, 0xa1, 0x36, 0x30, 0xde, 0x8e, 0x9f, 0xed, 0x6d, 0xe6, 0x97, 0x20, 0x3f, 0x21, 0x41, 0x11, 0x49, 0x78, 0xc6, 0x8b, 0xe3, 0xb6, 0x7f, 0x3f, 0x67, 0x6a, 0x67, 0xdd, 0x87, 0x9a, 0xf2, 0xb8, 0xb5, 0x6b, 0x70, 0xf, 0x1f, 0xc5, 0x3d, 0x54, 0x2c, 0x9a, 0x4a, 0x61, 0x55, 0xc5, 0xf, 0xb8, 0xed, 0x71, 0x80, 0xd9, 0xde, 0x5f, 0x7f, 0xae, 0x1, 0x53, 0xc, 0x47, 0xb5, 0x3c, 0xb8, 0x32, 0x85, 0xb1, 0xdc, 0x88, 0xed, 0x84, 0x81, 0x51, 0x40, 0xdc, 0xc1, 0x1c, 0xef, 0x20, 0x15, 0x30, 0xa9, 0xdf, 0x39, 0xc6, 0xbd, 0x92, 0x5, 0x6e, 0x4e, 0xdd, 0x48, 0x4b, 0x12, 0xc1, 0x42, 0x6b, 0xef, 0x72, 0x69, 0x2, 0xfa, 0xf7, 0x6b, 0xe3, 0x23, 0xa2, 0x3e, 0xd8, 0xe2, 0xf8, 0xf5, 0x83, 0xac, 0x75, 0x6e, 0x6c, 0x99, 0xbc, 0x98, 0xc9, 0xb2, 0xb9, 0x8d, 0xaf, 0xe9, 0xb5, 0x31, 0xad, 0xc1, 0x43, 0x15, 0x46, 0x2a, 0x80, 0x51, 0xb7, 0x8d, 0x58, 0xbb, 0xe1, 0xcd, 0x6c, 0x45, 0x6, 0x4, 0xe2, 0xf1, 0xec, 0xea, 0xdf, 0xac, 0x51, 0xb8, 0xda, 0xe7, 0x31, 0x64, 0xb2, 0x7, 0x6a, 0x8e, 0x95, 0x9d, 0x2f, 0x4f, 0x3, 0x59, 0x2, 0x44, 0x9c, 0x59, 0xd7, 0xd0, 0x78, 0x19, 0xb4, 0xc1, 0x72, 0x72, 0x28, 0x5a, 0x6c, 0xd8, 0xa4, 0xbe, 0xd9, 0xde, 0x54, 0x3d, 0xe6, 0x81, 0x28, 0xa9, 0xf5, 0xb5, 0xba, 0x10, 0xda, 0xa, 0x84, 0x83, 0x4, 0x71, 0x4a, 0xf5, 0xc, 0xf9, 0xf2, 0xc2, 0x5c, 0xc2, 0x23, 0x48, 0x44, 0xa2, 0x15, 0xfe, 0x9f, 0xf5, 0x90, 0x21, 0xae, 0x8d, 0xd8, 0xc, 0x8b, 0x20, 0x1d, 0xcf, 0x24, 0xe3, 0xc1, 0x94, 0x34, 0x3c, 0xf5, 0x6f, 0x65, 0x55, 0xb0, 0x86, 0x66, 0x97, 0x8, 0x3d, 0x41, 0xb4, 0xd9, 0x7f, 0x50, 0xd4, 0xd7, 0x55, 0x16, 0x7, 0x5e, 0x5b, 0xfa, 0xfe, 0xec, 0x9a, 0xee, 0x5b, 0x95, 0xc6, 0xbe, 0xc6, 0x68, 0x7f, 0xe5, 0xc7, 0xd0, 0x33, 0xbf, 0x5, 0xb7, 0x12, 0xb, 0x15, 0x18, 0xf9, 0x39, 0x3d, 0x56, 0xfc, 0x74, 0x2e, 0xa9, 0x72, 0x2b, 0xea, 0x61, 0xd1, 0x1d, 0x28, 0xe6, 0x8c, 0x98, 0x2f, 0xb2, 0xf4, 0x31, 0x2a, 0xa0, 0x0, 0x31, 0xb4, 0x93, 0x3b, 0x6d, 0xd6, 0x29, 0xb9, 0xc6, 0xd3, 0x80, 0x5, 0xe9, 0x61, 0x12, 0xae, 0x46, 0xc0, 0x6e, 0x3, 0x97, 0x8b, 0xde, 0xc, 0x66, 0x40, 0x22, 0x59, 0xff, 0x5b, 0xec, 0x8b, 0x9e, 0x78, 0xd7, 0xbf, 0x8f, 0xa6, 0x96, 0x4d, 0x2f, 0xc4, 0xba, 0xa0, 0xb0, 0x71, 0xfc, 0x37, 0x61, 0xe0, 0x2c, 0x7c, 0x7e, 0xda, 0xe6, 0x9c, 0x80, 0x79, 0xf3, 0x8c, 0x6d, 0x76, 0x3b, 0x8f, 0x89, 0xf6, 0x76, 0x5e, 0x3d, 0xd7, 0xc2, 0xca, 0x20, 0x93, 0xe8, 0xe7, 0xb2, 0xa1, 0xfc, 0x77, 0x67, 0xe3, 0x25, 0xf8, 0x1f, 0x3c, 0x52, 0x38, 0xc3, 0x39, 0x9d, 0x40, 0x8, 0x42, 0x67, 0xbf, 0x6e, 0x4b, 0x6, 0x5b, 0x2c, 0xa6, 0x86, 0xa5, 0xed, 0x8f, 0x9c, 0x88, 0x7a, 0x3d, 0x2a, 0xfa, 0x39, 0x3f, 0x12, 0x5, 0xc5, 0xed, 0xf7, 0x24, 0x0, 0xfe, 0xce, 0xf1, 0x9e, 0x1c, 0xbe, 0xfc, 0x5a, 0xc6, 0xf2, 0x2, 0xe5, 0x8, 0x3, 0x9c, 0x43, 0x1, 0x39, 0x9f, 0x6a, 0xd9, 0x36, 0xa3, 0x53, 0xdc, 0xea, 0x0, 0x3f, 0x2, 0x2a, 0x4d, 0xab, 0x4d, 0xc7, 0x7c, 0xb, 0x4d, 0xda, 0x40, 0x26, 0x9c, 0x43, 0xd7, 0x93, 0xd7, 0x81, 0x87, 0xc0, 0x48, 0x86, 0x3f, 0x79, 0x1c, 0x45, 0xa, 0x31, 0xe6, 0x2e, 0x13, 0x24, 0xc6, 0x85, 0xad, 0xd4, 0xc4, 0x5f, 0xbd, 0x90, 0x22, 0xc0, 0xc6, 0x14, 0xe4, 0x3a, 0x4e, 0x77, 0x27, 0xb4, 0xb7, 0xcc, 0xc5, 0x7c, 0x1f, 0x54, 0x17, 0x25, 0x11, 0xb8, 0xb5, 0xb8, 0xc, 0xc8, 0xb3, 0xea, 0x4a, 0x89, 0xdd, 0x36, 0xc6, 0x6c, 0xd7, 0xe7, 0xe, 0x78, 0xe5, 0xf, 0x2a, 0x5a, 0xee, 0xc6, 0x8e, 0x63, 0x1c, 0xb5, 0x82, 0x79, 0x23, 0x0, 0x79, 0x10, 0x3c, 0x17, 0x73, 0x1d, 0x53, 0x6c, 0x9a, 0xbe, 0x63, 0xad, 0x42, 0x5b, 0x4d, 0x8, 0x6c, 0x25, 0x38, 0x34, 0x7, 0x9b, 0xf8, 0xe9, 0x92, 0xad, 0x44, 0x91, 0x31, 0x53, 0xcb, 0xb2, 0x6a, 0x9c, 0xcd, 0x30, 0x7, 0xc8, 0x35, 0xc5, 0x25, 0x71, 0xd4, 0x30, 0x7f, 0x34, 0xae, 0xcc, 0x78, 0xa5, 0xfe, 0x6b, 0xaa, 0x21, 0xea, 0x1d, 0x2d, 0x4b, 0x81, 0x9b, 0x51, 0xdb, 0x2b, 0x87, 0x2b, 0x36, 0xaf, 0x68, 0xd3, 0x3a, 0x14, 0xe9, 0x52, 0xc, 0xcc, 0x73, 0xab, 0x48, 0xe8, 0x15, 0xd, 0x2b, 0xf5, 0xeb, 0xee, 0x9b, 0x12, 0x3d, 0x80, 0x8e, 0xe3, 0x41, 0x39, 0x82, 0xea, 0xa7, 0x5, 0x77, 0xad, 0x56, 0xce, 0x7f, 0xbd, 0x4b, 0x30, 0x48, 0x8, 0x35, 0xd0, 0x66, 0xbb, 0x68, 0x16, 0x25, 0x2f, 0xbf, 0xa8, 0x9c, 0x5f, 0xff, 0xa0, 0x33, 0xd2, 0xae, 0xcd, 0x2e, 0x8c, 0x28, 0xb9, 0x6b, 0xcf, 0x4, 0x2a, 0xf8, 0xef, 0xe5, 0x52, 0x33, 0xe, 0xec, 0x1, 0x19, 0x4f, 0xda, 0xe4, 0x7, 0xeb, 0xa3, 0x85, 0xaf, 0x6f, 0x25, 0xc1, 0x4, 0xac, 0x7, 0x24, 0x22, 0x85, 0x5b, 0x13, 0x91, 0xc3, 0x9c, 0x42, 0x4c, 0xa2, 0xed, 0xc1, 0xa6, 0x5d, 0x5, 0x48, 0xdc, 0x7, 0x6, 0x12, 0x2a, 0x5b, 0xce, 0x26, 0xb2, 0xd6, 0x14, 0x2a, 0xb7, 0xab, 0xac, 0xf1, 0x6c, 0x97, 0xd9, 0xfd, 0x11, 0xde, 0xa7, 0xbb, 0x36, 0x39, 0x17, 0xbe, 0xc2, 0x34, 0xbb, 0x9b, 0x40, 0x56, 0xad, 0x4f, 0x7d, 0xc1, 0x6f, 0xe9, 0x31, 0x67, 0x4f, 0xd, 0x58, 0x8a, 0xff, 0x1d, 0xf1, 0xe5, 0x96, 0xe1, 0x0, 0x38, 0x2c, 0x36, 0xb6, 0x3b, 0x59, 0xe1, 0x73, 0xe9, 0x63, 0xb1, 0xec, 0xf6, 0xf5, 0x50, 0x7c, 0x32, 0xea, 0x2a, 0x3c, 0x1a, 0x60, 0xe, 0x8c, 0xde, 0xaf, 0x9f, 0x8c, 0x51, 0x2a, 0x9, 0xa8, 0x5a, 0x50, 0xb0, 0x8a, 0x97, 0xb4, 0x8c, 0xca, 0xbe, 0x2b, 0xee, 0xfa, 0x88, 0x76, 0x5, 0xd0, 0x84, 0xb1, 0x2, 0x94, 0x18, 0xb4, 0x0, 0x6c, 0xa0, 0x92, 0x55, 0xe4, 0x0, 0xe8, 0x34, 0xc1, 0x82, 0xd6, 0x6e, 0x5c, 0x15, 0xa2, 0xe6, 0x1d, 0x1f, 0xec, 0x9b, 0xe3, 0x13, 0xb9, 0xc1, 0xec, 0x92, 0x1f, 0x81, 0xed, 0xb3, 0x5a, 0xdb, 0x3e, 0xbb, 0x48, 0x76, 0xc6, 0xc5, 0xf5, 0x76, 0xf2, 0xcb, 0x6f, 0x78, 0xe5, 0x2a, 0x44, 0x5b, 0x57, 0x44, 0x8a, 0x93, 0x61, 0x7c, 0x81, 0x39, 0x70, 0x9, 0xd4, 0xb6, 0x2d, 0x72, 0x67, 0xc9, 0x85, 0xdb, 0xbf, 0x9a, 0xd, 0x40, 0x5e, 0x2b, 0x4d, 0xa0, 0xb0, 0xa0, 0x3a, 0xaa, 0x8e, 0x3e, 0x74, 0x95, 0x28, 0x18, 0x89, 0x7f, 0x8e, 0xd0, 0x39, 0x18, 0x6d, 0x20, 0xad, 0xab, 0xb4, 0x55, 0xe1, 0x89, 0x68, 0x39, 0xd6, 0x12, 0xeb, 0xe6, 0x81, 0x5d, 0x52, 0x7f, 0x1e, 0xd5, 0xc4, 0x84, 0x2f, 0xad, 0x2b, 0xfd, 0xa6, 0x7a, 0xa8, 0x43, 0xd5, 0xbb, 0x71, 0xff, 0xde, 0x4d, 0xc4, 0x93, 0xa7, 0xdc, 0x31, 0xe6, 0xad, 0x28, 0xc1, 0xe1, 0x88, 0x24, 0x2c, 0x55, 0xdf, 0x9, 0x60, 0xf7, 0x59, 0x66, 0xdb, 0x10, 0x70, 0x67, 0x7c, 0x3e, 0x70, 0x73, 0x21, 0xf6, 0x35, 0x72, 0xcc, 0x52, 0xfe, 0x9c, 0x30, 0xda, 0x76, 0xf4, 0x3b, 0xb2, 0x92, 0xf, 0x60, 0xe8, 0x2, 0x67, 0x6e, 0x7e, 0x15, 0xed, 0xa, 0x79, 0x8b, 0xf6, 0x50, 0x6e, 0x29, 0x44, 0x51, 0xb3, 0x9b, 0xaf, 0x4a, 0x96, 0x36, 0xc8, 0xd5, 0xf1, 0xbe, 0x2c, 0xbb, 0x28, 0x9c, 0xaf, 0xd7, 0xe8, 0x3a, 0x67, 0x8c, 0xca, 0x1d, 0xa7, 0xc0, 0x7, 0x51, 0x99, 0xb0, 0x48, 0x14, 0x22, 0x0, 0x78, 0x98, 0x9c, 0xa7, 0x3a, 0x8b, 0xc5, 0xb1, 0xf9, 0xc6, 0x4a, 0x27, 0xb4, 0x5c, 0x38, 0x25, 0x53, 0xed, 0x76, 0x52, 0x5c, 0xe1, 0x62, 0x40, 0x81, 0xca, 0x62, 0x16, 0x95, 0xf6, 0xa, 0x20, 0xc8, 0x21, 0xbc, 0xdc, 0xad, 0x4e, 0x5a, 0xb5, 0xce, 0x2e, 0x1b, 0x58, 0xd, 0x88, 0x4a, 0x68, 0x83, 0xc2, 0x56, 0x9d, 0x3d, 0x2f, 0xaa, 0xdf, 0xff, 0x5d, 0xf8, 0xf, 0x8d, 0xa8, 0x40, 0x50, 0x82, 0xc8, 0x3a, 0x6d, 0xe2, 0x4, 0x58, 0xc7, 0x15, 0xb2, 0xe6, 0x18, 0x86, 0xd6, 0x3, 0xd2, 0x9c, 0xa, 0xa9, 0xd5, 0x2, 0xd9, 0x64, 0x4a, 0x21, 0x16, 0x4a, 0x2b, 0xb3, 0xa6, 0x93, 0x58, 0xd2, 0xf1, 0x7, 0x94, 0x8b, 0xc, 0x82, 0x4, 0x20, 0xaa, 0xe4, 0x18, 0xd5, 0x71, 0xf2, 0xda, 0xc5, 0x1c, 0x83, 0xc6, 0x8, 0xad, 0x7, 0xdc, 0xb, 0x97, 0xd8, 0x4a, 0xb8, 0xd5, 0x6e, 0xa7, 0x1e, 0xcf, 0x4c, 0x48, 0xf6, 0xc8, 0x90, 0xe7, 0x5, 0x10, 0xf, 0x5f, 0xac, 0xa1, 0x2d, 0xf, 0x5f, 0xab, 0xa5, 0x46, 0x0, 0xa, 0x3, 0x7a, 0x60, 0x4c, 0x42, 0x20, 0x83, 0xaa, 0xde, 0xda, 0xf, 0x49, 0x42, 0x14, 0x11, 0xe8, 0xa7, 0x48, 0x1e, 0x28, 0x6e, 0xdb, 0x1a, 0x77, 0x5a, 0x40, 0xb6, 0xde, 0x64, 0xa4, 0x1f, 0x96, 0x4c, 0x78, 0x8d, 0x62, 0x91, 0x47, 0xd5, 0x66, 0x97, 0x92, 0xec, 0xda, 0xc0, 0x7c, 0x86, 0x1f, 0x9e, 0x53, 0x28, 0xca, 0x2d, 0x7f, 0xb, 0xc8, 0xcc, 0x8d, 0x49, 0x9e, 0x8f, 0xbe, 0x82, 0xf2, 0xe7, 0xef, 0x18, 0xd4, 0xb5, 0x3c, 0x17, 0xda, 0xa9, 0xa4, 0x5, 0xe, 0x8d, 0x1d, 0x13, 0x3e, 0xe8, 0x2f, 0x91, 0x75, 0x41, 0xf2, 0x8b, 0x16, 0xe5, 0x1f, 0x68, 0x77, 0x58, 0x90, 0x4e, 0xd7, 0x58, 0xbd, 0xb, 0x90, 0x2, 0x18, 0xd5, 0x71, 0xf2, 0xda, 0xc5, 0x1c, 0x83, 0xc6, 0x8, 0xad, 0x7, 0xdc, 0xb, 0x97, 0xd8, 0x4a, 0xb8, 0xd5, 0x6e, 0xa7, 0x1e, 0xcf, 0x4c, 0x48, 0x5a, 0x77, 0xfc, 0x81, 0x46, 0x8b, 0xdb, 0x97, 0x4d, 0x3d, 0xd5, 0x34, 0x77, 0x74, 0xd2, 0x6c, 0x3c, 0xb6, 0xa9, 0xc8, 0xe7, 0x39, 0x2a, 0x40, 0x8e, 0x6c, 0x0, 0x7, 0x1f, 0xcb, 0x8c, 0x1a, 0xf7, 0xc8, 0x90, 0xe7, 0x5, 0x10, 0x86, 0xd, 0x37, 0x4f, 0xe8, 0x86, 0xd, 0x34, 0xae, 0xd2, 0x0, 0x0, 0x19, 0x80, 0x60, 0x5f, 0x7e, 0x8, 0x6a, 0xf7, 0x9b, 0xb4, 0x1b, 0x4f, 0xe4, 0xb7, 0x11, 0x77, 0x7f, 0xf5, 0x26, 0xc4, 0xb3, 0x25, 0xe3, 0xc8, 0x18, 0xe8, 0x10, 0xb0, 0xc0, 0x5a, 0xb1, 0x62, 0xd4, 0x6b, 0x6d, 0x28, 0xfb, 0x41, 0x22, 0xca, 0x2f, 0xfb, 0xc2, 0x6, 0x2a, 0x93, 0x32, 0x57, 0x80, 0xc6, 0xc9, 0x20, 0xa2, 0xeb, 0x7c, 0x95, 0x9c, 0x14, 0xf1, 0x67, 0xcc, 0x21, 0x71, 0xd7, 0x45, 0xe3, 0x17, 0x86, 0x8f, 0x8a, 0x79, 0x2, 0xaf, 0xf5, 0x39, 0x72, 0x8d, 0x60, 0xb3, 0x17, 0xd0, 0xc2, 0xd2, 0xc3, 0xae, 0x41, 0x6, 0xc9, 0x2e, 0x73, 0x6e, 0x40, 0x87, 0xd9, 0xe8, 0xbd, 0xe1, 0xee, 0x8b}
