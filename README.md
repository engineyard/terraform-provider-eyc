# Terraform Provider Hashicups

Run the following command to build the provider

```shell
go build -o terraform-provider-eyc
```

## Test sample configuration

At the root path, build and install the provider.

```shell
make install
```



Then, cd to the examples folder, run the following command to initialize the workspace and apply the sample configuration.

```shell
cd /examples && terraform init && terraform apply
```
