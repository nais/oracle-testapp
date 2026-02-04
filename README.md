# Oracle Test App

En enkel Go-applikasjon som tester tilkobling til Oracle-database fra GCP.

## Hva gjør appen?

- Kobler til Oracle-database via go-ora driver
- Enkel webfrontend for å vise og legge til verdier i databasen
- Health-endpoint som verifiserer database-tilkobling

## Miljøvariabler

| Variabel | Beskrivelse |
|----------|-------------|
| `ORACLE_URL` | JDBC URL, f.eks. `jdbc:oracle:thin:@host:1521/database` |
| `ORACLE_USERNAME` | Database-bruker |
| `ORACLE_PASSWORD` | Database-passord |
| `PORT` | HTTP-port (default: `8080`) |

## Kjøre lokalt

```bash
export ORACLE_URL=jdbc:oracle:thin:@host:1521/database
export ORACLE_USERNAME=bruker
export ORACLE_PASSWORD=passord
go run .
```

Åpne http://localhost:8080

## Deploy til NAIS

Appen deployes til `dev-gcp` i `nais`-namespacet.

### Forutsetninger

1. Opprett secret med Oracle-credentials:

```bash
kubectl -n nais create secret generic oracleverk-testapp-oracle \
  --from-literal=ORACLE_URL=jdbc:oracle:thin:@host:1521/database \
  --from-literal=ORACLE_USERNAME=bruker \
  --from-literal=ORACLE_PASSWORD=passord
```

2. Legg til NetworkPolicy for å tillate trafikk til Oracle (håndteres av oracleverk-cronjob)

3. Legg til Kyverno-exception for Oracle IP-er (håndteres av oracleverk-cronjob)

### Manuell deploy

```bash
gh workflow run deploy.yaml
```

## Endepunkter

| Metode | Path | Beskrivelse |
|--------|------|-------------|
| GET | `/` | Frontend |
| GET | `/api/values` | Hent alle verdier |
| POST | `/api/values` | Legg til verdi (`{"value": "..."}`) |
| GET | `/api/health` | Health check |