package main

import "testing"

var data Sha256Data

func DoTest(input, expected string, t *testing.T) {
	data.original = []byte(input)
	data.CalculateSha256()
	hashString := data.GetHashString()
	if hashString != expected {
		t.Fatalf("Input: %v, Expected: %v, Got: %v", input, expected, hashString)
	}
}

func TestSha(t *testing.T) {
	data.Initialize()
	DoTest("abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", t)
	DoTest("aaaaaaaaaa", "bf2cb58a68f684d95a3b78ef8f661c9a4e5b09e82cc8f9cc88cce90528caeb27", t)
	DoTest("aaaaaaaaaa", "bf2cb58a68f684d95a3b78ef8f661c9a4e5b09e82cc8f9cc88cce90528caeb27", t)
	DoTest("abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq", "248d6a61d20638b8e5c026930c3e6039a33ce45964ff2167f6ecedd419db06c1", t)
	DoTest("Hello, World!", "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f", t)
}
