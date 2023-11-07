# Portfolio

- Static site, built with go templates into build folder
- Served locally with go net/http for testing

## git

- Allow list used in gitignore

## Serve locally

To run build and run locally:

```shell
go run ./ --dev
```

using go run main.go would only run the single file and miss the imports from other files in the main package

## Command line arguments

| Argument | Description                                                               |
| -------- | ------------------------------------------------------------------------- |
| `--dev`  | If set will serve the static site on the port defined in the config file. |

## Config settings

| Setting     | Description                                                                      |
| ----------- | -------------------------------------------------------------------------------- |
| `inputDir`  | the program will loop through this folder to find html templates and build them. |
| `outputDir` | the output directory for the built template pages.                               |
| `port`      | the port used by the dev server for serving the static pages.                    |
| `pageData`  | object sent to templates to display.                                             |

## Deploy

- Will deploy automatically the build folder to Azure from main via github actions
- Static web app
- Storage Account
