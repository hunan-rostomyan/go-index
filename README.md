Offline indexing for documentation to make it searchable.

### Install

- Clone this repository and `cd` into it.

- Download [Bosh](http://bosh.io) documentation from [Cloud Foundry](https://pivotal.io/platform) by issuing:
  ```bash
  go get -v github.com/cloudfoundry/docs-bosh`
  ```
  Alternatively you can modify *main.go* to point to your documentation.

### Index

```bash
make
```

If successful, this will yield an *index.json* file containing an inverted index of your documentation.

### Test

```bash
make test
```
