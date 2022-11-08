## Quick start

In your code

    import _ "github.com/gregoryv/another"

in another project, vendor dependencies

    $ go mod vendor

then generate swagger from the dependencies

    $ swag init --parseVendor 