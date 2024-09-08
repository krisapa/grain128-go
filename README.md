## Grain-128 Stream Cipher in Go

This repository implements the Grain-128 stream cipher.

For more details on the Grain cipher, see [A Stream Cipher Proposal: Grain-128](https://www.ecrypt.eu.org/stream/p3ciphers/grain/Grain_p3.pdf).

### Performance
goos: darwin,
goarch: arm64,
cpu: Apple M1 Pro
| Benchmark              | Iterations | Time per Op |
|------------------------|------------|-------------|
| **Grain128Keystream**   | 89,611     | 13,413 ns/op |
| **Grain128Encrypt**     | 86,306     | 13,317 ns/op |
| **Grain128Decrypt**     | 30,188     | 39,706 ns/op |

---

To run the benchmarks yourself, use the following command:

```bash
go test -bench=.
