# Envoyer Action
A simple action that triggers a deployment on [Envoyer](https://envoyer.io).

## Requirements
You need to create a secret for you action called `ENVOYER_TOKEN` that holds the Bearer Token that you can generate on the Envoyer dashboard.

## Example
```yaml
uses: httpoz/envoyer-action@v1
with:
  branch: "main"
  projectId: 12345
```

## Dependencies
The MVP depends on a cocktail of two actions.
- [fjogeleit/http-request-action](https://github.com/fjogeleit/http-request-action) - To trigger requests on Envoyer API.
- [mydea/action-wait-for-api](https://github.com/mydea/action-wait-for-api) - To poll the Envoyer API until the deployment is finished.