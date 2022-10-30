<!--
order: 7
-->

# Clients

A user can query the `x/claims` module using the CLI, gRPC or REST.

## CLI

Find below a list of `vinced` commands added with the `x/claims` module. You can obtain the full list by using the `vinced -h` command.

### Queries

The `query` commands allow users to query `claims` state.

**`total-unclaimed`**

Allows users to query total amount of unclaimed tokens from the airdrop.

```bash
vinced query claims total-unclaimed [flags]
```

**`claims-records`**

Allows users to query all the claims records available.

```bash
vinced query claims claims-records [flags]
```

**`claims-record`**

Allows users to query a claims record for a given user.

```bash
vinced query claims claims-record [address] [flags]
```

**`params`**

Allows users to query claims params.

```bash
vinced query claims params [flags]
```

## gRPC

### Queries

| Verb   | Method                                     | Description                                      |
|--------|--------------------------------------------|--------------------------------------------------|
| `gRPC` | `vince.claims.v1.Query/TotalUnclaimed`     | Gets the total unclaimed tokens from the airdrop |
| `gRPC` | `vince.claims.v1.Query/ClaimsRecords`      | Gets all registered claims records               |
| `gRPC` | `vince.claims.v1.Query/ClaimsRecord`       | Get the claims record for a given user            |
| `gRPC` | `vince.claims.v1.Query/Params`             | Gets claims params                               |
| `GET`  | `/vince/claims/v1/total_unclaimed`         | Gets the total unclaimed tokens from the airdrop |
| `GET`  | `/vince/claims/v1/claims_records`          | Gets all registered claims records               |
| `GET`  | `/vince/claims/v1/claims_record/{address}` | Gets a claims record for a given user            |
| `GET`  | `/vince/claims/v1/params`                  | Gets claims params                               |